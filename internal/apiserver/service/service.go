package service

import "iamgeek/internal/apiserver/store"

// Service defines functions used to return resource interface.
type Service interface {
	Blockchains() BlockchainSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Blockchains() BlockchainSrv {
	return newBlockchains(s)
}
