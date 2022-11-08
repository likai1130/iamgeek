// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// apiserver is the api server for iam-apiserver service.
// it is responsible for serving the platform RESTful resource management.
package main

import (
	_ "go.uber.org/automaxprocs"
	"iamgeek/internal/apiserver"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	apiserver.NewApp("configs/iam-apiserver.yaml")
}
