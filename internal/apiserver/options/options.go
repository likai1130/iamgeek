// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package options contains flags and options for initializing an apiserver
package options

import (
	"encoding/json"
	genericoptions "iamgeek/internal/pkg/options"
	"iamgeek/internal/pkg/server"
	"iamgeek/pkg/log"
)

// Options runs an iam api server.
type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
	ExtendOptions           *ExtendOptions                         `json:"extend"  mapstructure:"extend"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		Log:                     log.NewOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
		ExtendOptions:           &ExtendOptions{},
	}

	return &o
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *Options) ApplyTo(c *server.Config) error {
	return nil
}

// Flags returns flags for a specific APIServer by section name.

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

// Complete set default Options.
func (o *Options) Complete() error {
	return o.SecureServing.Complete()
}
