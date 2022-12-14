package store

import (
	"context"
	metav1 "iamgeek/internal/apiserver/meta/v1"
	v1 "iamgeek/internal/apiserver/model/v1"
)

type BlockChainStore interface {
	Create(ctx context.Context, user *v1.BlockChain) error
	Update(ctx context.Context, user *v1.BlockChain) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*v1.BlockChain, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error)
}
