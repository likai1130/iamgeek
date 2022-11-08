// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

//go:generate mockgen -self_package=iamgeek/internal/apiserver/store -destination mock_store.go -package store iamgeek/internal/apiserver/store Factory,UserStore,SecretStore,PolicyStore

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	BlockChains() BlockChainStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
