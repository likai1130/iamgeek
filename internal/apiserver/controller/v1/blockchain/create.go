package blockchain

import (
	"github.com/gin-gonic/gin"
	v1 "iamgeek/internal/apiserver/model/v1"
	"iamgeek/internal/pkg/code"
	"iamgeek/internal/pkg/core"
	"iamgeek/pkg/errors"
	"iamgeek/pkg/log"
)

// Create 创建一个链
// @Summary 创建一个链
// @Description 创建一个链
// @Tags 区块链管理
// @Param data body v1.BlockChain true "链数据"
// @Success 200 {object} core.Response{data=v1.BlockChain} "{"code": 200, "message":"Ok","data": [...]}"
// @Success 400 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/iamgeek/chains [post]
func (u *BlockchainController) Create(c *gin.Context) {
	log.L(c).Info("chain create function called.")

	var r v1.BlockChain

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if err := u.srv.Blockchains().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, r)
}
