// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package apiserver does all the work necessary to create a iam APIServer.
package apiserver

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"iamgeek/internal/apiserver/config"
	"iamgeek/internal/apiserver/options"
	"iamgeek/pkg/log"
)

// NewApp creates an App object with default parameters.
/*func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server", basename, app.WithRunFunc(run(opts)))
	return application
}*/

func NewApp(config string) {
	opts := options.NewOptions()
	Setup(config, opts)
	e := run(opts)
	if e != nil {
		log.Error("start err!")
	}

}
func run(opts *options.Options) error {
	log.Init(opts.Log)
	defer log.Flush()

	cfg, err := config.CreateConfigFromOptions(opts)
	if err != nil {
		return err
	}

	return Run(cfg)
}

// Setup 载入配置文件
func Setup(configFile string, opts *options.Options) {
	v := viper.New()
	//自动获取全部的env加入到viper中。默认别名和环境变量名一致。（如果环境变量多就全部加进来）
	v.AutomaticEnv()

	//替换读取格式。默认a.b.c.d格式读取env，改为a_b_c_d格式读取
	//v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// 本地配置文件位置
	v.SetConfigFile(configFile)

	//支持 "yaml", "yml", "json", "toml", "hcl", "tfvars", "ini", "properties", "props", "prop", "dotenv", "env":
	v.SetConfigType("yml")

	//读配置文件到viper配置中
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.Get("max-connection-life-time")

	// 系列化成config对象
	if err = v.Unmarshal(&opts); err != nil {
		panic(err)
	}

	config.ExtendConf = opts.ExtendOptions

	log.Info("!!! config init success")
}
