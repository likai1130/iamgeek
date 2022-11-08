// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"iamgeek/internal/apiserver/config"
	"iamgeek/internal/apiserver/options"
	"iamgeek/internal/apiserver/store"
	"iamgeek/internal/apiserver/store/database"
	genericoptions "iamgeek/internal/pkg/options"
	genericapiserver "iamgeek/internal/pkg/server"
	"iamgeek/pkg/log"
	"iamgeek/pkg/shutdown"
	"iamgeek/pkg/shutdown/shutdownmanagers/posixsignal"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

// ExtraConfig defines extra configuration for the server
type ExtraConfig struct {
	mysqlOptions *genericoptions.MySQLOptions //数据库
	mongoOptions *options.MongoDBOptions
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	extraConfig.complete().New()
	server := &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		databaseStore, _ := database.GetDatabaseFactoryOr(nil, nil)
		if databaseStore != nil {
			_ = databaseStore.Close()
		}

		s.genericAPIServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	// start shutdown managers
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}

type completedExtraConfig struct {
	*ExtraConfig
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() *completedExtraConfig {
	return &completedExtraConfig{c}
}

func (c *completedExtraConfig) New() error {
	// 数据库初始化
	storeIns, err := database.GetDatabaseFactoryOr(c.mysqlOptions, c.mongoOptions)
	if err != nil {
		return err
	}
	//路由配置
	store.SetClient(storeIns)
	return nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		mysqlOptions: cfg.MySQLOptions,
	}, nil
}
