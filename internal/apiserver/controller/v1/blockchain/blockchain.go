package blockchain

import (
	srvv1 "iamgeek/internal/apiserver/service"
	"iamgeek/internal/apiserver/store"
)

// BlockchainController create a user handler used to handle request for user resource.
type BlockchainController struct {
	srv srvv1.Service
}

// NewBlockchainController creates a user handler.
func NewBlockchainController(store store.Factory) *BlockchainController {
	return &BlockchainController{
		srv: srvv1.NewService(store),
	}
}
