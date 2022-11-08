package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"iamgeek/internal/apiserver/options"
	"iamgeek/internal/apiserver/store"
	mongodb "iamgeek/pkg/mongdb"
)

/**
享元模式(连接池)：

	减少应用程序创建的对象，降低程序内存的占用，增强程序性能。
	将多数据源的数据库连接缓存起来，直接使用。
*/

type datastore struct {
	db *mongo.Client

	// can include two database instance if needed
	// docker *mongo.Client
	// db *mongo.Client
}

func (ds *datastore) BlockChains() store.BlockChainStore {
	//TODO implement me
	panic("implement me")
}

func (ds *datastore) Close() error {
	return ds.db.Disconnect(context.TODO())
}

var (
	mongoFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMongoFactoryOr(opts *options.MongoDBOptions) (store.Factory, error) {
	if opts == nil && mongoFactory == nil {
		return nil, fmt.Errorf("failed to get mongo store fatory")
	}

	var err error
	var dbIns *mongo.Client
	once.Do(func() {
		options := &mongodb.Options{
			Hosts:                   opts.Hosts,
			UserName:                opts.UserName,
			Password:                opts.Password,
			MaxPoolSize:             opts.MaxPoolSize,
			DbName:                  opts.DbName,
			ReplicaSet:              opts.ReplicaSet,
			ReadPreference:          opts.ReadPreference,
			ServerSelectionTimeoutS: opts.ServerSelectionTimeoutS,
		}
		dbIns, err = mongodb.New(options)

		mongoFactory = &datastore{dbIns}
	})

	if mongoFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mongo store fatory, mongoFactory: %+v, error: %w", mongoFactory, err)
	}

	return mongoFactory, nil
}
