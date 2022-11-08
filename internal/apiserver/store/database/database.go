// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	v1 "iamgeek/internal/apiserver/model/v1"
	"iamgeek/internal/apiserver/options"
	"iamgeek/pkg/errors"
	mongodb "iamgeek/pkg/mongdb"

	"gorm.io/gorm"

	"iamgeek/internal/apiserver/store"
	"iamgeek/internal/pkg/logger"
	genericoptions "iamgeek/internal/pkg/options"
	"iamgeek/pkg/db"
)

type datastore struct {
	mysqldb *gorm.DB
	mongodb *mongo.Client

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) BlockChains() store.BlockChainStore {
	return newBlockchains(ds)
}

func (ds *datastore) Close() error {
	ds.mongodb.Disconnect(context.TODO())
	mysqlDB, err := ds.mysqldb.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return mysqlDB.Close()
}

var (
	databaseFactory store.Factory
	once            sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetDatabaseFactoryOr(mysqlOpts *genericoptions.MySQLOptions, mongoOpts *options.MongoDBOptions) (store.Factory, error) {
	if (mysqlOpts == nil || mongoOpts == nil) && databaseFactory == nil {
		return nil, fmt.Errorf("failed to get database store fatory")
	}

	var mysqlErr, mongoErr error
	var mysqlIns *gorm.DB

	var mongoIns *mongo.Client

	once.Do(func() {
		mysqlIns, mysqlErr = mysqlInstance(mysqlOpts)
		mongoIns, mongoErr = mongodbInstance(mongoOpts)
		databaseFactory = &datastore{mysqlIns, mongoIns}
	})

	if databaseFactory == nil || mysqlErr != nil || mongoErr != nil {
		return nil, fmt.Errorf("failed to get database store fatory, databaseFactory: %+v, mysqlError: %w,mongoError: %w", databaseFactory, mysqlErr, mongoErr)
	}

	return databaseFactory, nil
}

// cleanDatabase tear downs the database tables.
// nolint:unused // may be reused in the feature, or just show a migrate usage.
func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&v1.User{}); err != nil {
		return errors.Wrap(err, "drop user table failed")
	}

	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
// won't delete/change current data.
// nolint:unused // may be reused in the feature, or just show a migrate usage.
func migrateDatabase(db *gorm.DB) error {

	if err := db.AutoMigrate(&v1.BlockChain{}); err != nil {
		return errors.Wrap(err, "migrate blockchain model failed")
	}

	return nil
}

// resetDatabase resets the database tables.
// nolint:unused,deadcode // may be reused in the feature, or just show a migrate usage.
func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}

func mysqlInstance(opts *genericoptions.MySQLOptions) (*gorm.DB, error) {
	options := &db.Options{
		Host:                  opts.Host,
		Username:              opts.Username,
		Password:              opts.Password,
		Database:              opts.Database,
		MaxIdleConnections:    opts.MaxIdleConnections,
		MaxOpenConnections:    opts.MaxOpenConnections,
		MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
		LogLevel:              opts.LogLevel,
		Logger:                logger.New(opts.LogLevel),
	}
	dbIns, err := db.New(options)
	if err != nil {
		return nil, err
	}
	// uncomment the following line if you need auto migration the given models
	// not suggested in production environment.
	err = migrateDatabase(dbIns)
	if err != nil {
		return nil, err
	}
	return dbIns, nil
}

// GetMySQLFactoryOr create mysql factory with the given config.
func mongodbInstance(opts *options.MongoDBOptions) (*mongo.Client, error) {
	mongoOpts := &mongodb.Options{
		Hosts:                   opts.Hosts,
		UserName:                opts.UserName,
		Password:                opts.Password,
		MaxPoolSize:             opts.MaxPoolSize,
		DbName:                  opts.DbName,
		ReplicaSet:              opts.ReplicaSet,
		ReadPreference:          opts.ReadPreference,
		ServerSelectionTimeoutS: opts.ServerSelectionTimeoutS,
	}
	return mongodb.New(mongoOpts)
}
