package v1

import "iamgeek/internal/apiserver/meta/v1"

type BlockChain struct {
	v1.ObjectMeta
	ChainId        uint32     `json:"chain_id" gorm:"unique;column:chain_id;type:int(32);not null"`
	ChainName      string     `json:"chain_name" gorm:"type:varchar(20)"`
	InfraHttp      string     `json:"infra_http" gorm:"type:varchar(200)"`
	InfraWebsocket string     `json:"infra_websocket" gorm:"type:varchar(200)"`
	ERC20Tokens    []Contract `json:"erc_20_tokens,omitempty" gorm:"-"`
}

type BlockChainList struct {
	v1.ListMeta
	Items []*BlockChain `json:"items"`
}

func (BlockChain) TableName() string {
	return "blockchain"
}
