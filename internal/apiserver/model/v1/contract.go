package v1

import "iamgeek/internal/apiserver/meta/v1"

type Contract struct {
	v1.Model
	ChainId      uint32 `json:"chain_id" gorm:"type:int(20)"` //外键
	ContractAddr string `json:"contract_addr" gorm:"type:int(20)"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Decimals     string `json:"decimals"`
	Type         string `json:"type"`
}

func (Contract) TableName() string {
	return "contract"
}
