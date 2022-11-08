package service

import (
	"context"
	"regexp"
	metav1 "iamgeek/internal/apiserver/meta/v1"
	v1 "iamgeek/internal/apiserver/model/v1"
	"iamgeek/internal/apiserver/store"
	"iamgeek/internal/pkg/code"
	"iamgeek/pkg/errors"
	"iamgeek/pkg/log"
)

type BlockchainSrv interface {
	Create(ctx context.Context, chain *v1.BlockChain) error
	Update(ctx context.Context, chain *v1.BlockChain) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*v1.BlockChain, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error)
}

type blockchainService struct {
	store store.Factory
}

var _ BlockchainSrv = (*blockchainService)(nil)

func newBlockchains(srv *service) *blockchainService {
	return &blockchainService{store: srv.store}
}

// List returns user list in the storage. This function has a good performance.
func (u *blockchainService) List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error) {
	blockchains, err := u.store.BlockChains().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	/*wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	var m sync.Map

	// Improve query efficiency in parallel
	for _, blockchain := range blockchains.Items {
		wg.Add(1)

		go func(blockchain *v1.BlockChain) {
			defer wg.Done()

			// some cost time process
			policies, err := u.store.Policies().List(ctx, user.Name, metav1.ListOptions{})
			if err != nil {
				errChan <- errors.WithCode(code.ErrDatabase, err.Error())

				return
			}

			m.Store(blockchain.ID, &v1.BlockChain{
				ObjectMeta: metav1.ObjectMeta{
					ID:         user.ID,
					InstanceID: user.InstanceID,
					Name:       user.Name,
					Extend:     user.Extend,
					CreatedAt:  user.CreatedAt,
					UpdatedAt:  user.UpdatedAt,
				},
				Nickname:    user.Nickname,
				Email:       user.Email,
				Phone:       user.Phone,
				TotalPolicy: policies.TotalCount,
				LoginedAt:   user.LoginedAt,
			})
		}(blockchain)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*v1.BlockChain, 0, len(blockchains.Items))
	for _, chain := range blockchains.Items {
		info, _ := m.Load(chain.ID)
		infos = append(infos, info.(*v1.BlockChain))
	}*/

	log.L(ctx).Debugf("get %d chains from backend storage.", len(blockchains.Items))

	return blockchains, nil
}

func (u *blockchainService) Create(ctx context.Context, blockchain *v1.BlockChain) error {
	if err := u.store.BlockChains().Create(ctx, blockchain); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error()); match {
			return errors.WithCode(code.ErrUserAlreadyExist, err.Error())
		}

		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (u *blockchainService) Delete(ctx context.Context, id int64) error {
	if err := u.store.BlockChains().Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u *blockchainService) Get(ctx context.Context, id int64) (*v1.BlockChain, error) {
	chain, err := u.store.BlockChains().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (u *blockchainService) Update(ctx context.Context, blockchain *v1.BlockChain) error {
	if err := u.store.BlockChains().Update(ctx, blockchain); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}
