package database

import (
	"context"
	"gorm.io/gorm"
	metav1 "iamgeek/internal/apiserver/meta/v1"
	v1 "iamgeek/internal/apiserver/model/v1"
	"iamgeek/internal/apiserver/store"
	"iamgeek/internal/pkg/code"
	"iamgeek/internal/pkg/util/gormutil"
	"iamgeek/pkg/errors"
)

type blockchains struct {
	db *gorm.DB
}

func newBlockchains(ds *datastore) store.BlockChainStore {
	return &blockchains{ds.mysqldb}
}

// Create creates a new user account.
func (u *blockchains) Create(ctx context.Context, user *v1.BlockChain) error {
	return u.db.Create(&user).Error
}

// Update updates an user account information.
func (u *blockchains) Update(ctx context.Context, user *v1.BlockChain) error {
	return u.db.Save(user).Error
}

// Delete deletes the user by the user identifier.
func (u *blockchains) Delete(ctx context.Context, id int64) error {
	err := u.db.Unscoped().Where("id = ?", id).Delete(&v1.BlockChain{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Get return an user by the user identifier.
func (u *blockchains) Get(ctx context.Context, id int64) (*v1.BlockChain, error) {
	chain := &v1.BlockChain{}
	err := u.db.Where("id = ?", id).First(&chain).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrBlockChainNotFound, err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return chain, nil
}

// List return all blockchains.
func (u *blockchains) List(ctx context.Context, opts metav1.ListOptions) (*v1.BlockChainList, error) {
	ret := &v1.BlockChainList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)
	d := u.db.Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	ret.SetTotalPage(*opts.Limit)
	ret.SetCurrentPage(*opts.Offset)
	return ret, d.Error
}
