package blockchain

import (
	"github.com/gin-gonic/gin"
	"strconv"
	v1 "iamgeek/internal/apiserver/model/v1"
	"iamgeek/internal/pkg/code"
	"iamgeek/internal/pkg/core"
	"iamgeek/pkg/errors"
	"iamgeek/pkg/log"
)

// Update 更新区块链信息
// @Summary 更新区块链信息
// @Description 更新区块链信息
// @Tags 区块链管理
// @Param id path int true "主键id"
// @Param data body v1.BlockChain true "链数据"
// @Success 200 {object} core.Response{data=v1.BlockChain} "{"code": 200, "message":"Ok", "data": [...]}"
// @Success 400 {object} core.Response
// @Success 404 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/iamgeek/chains/{id} [put]
func (u *BlockchainController) Update(c *gin.Context) {
	log.L(c).Info("update user function called.")

	var r v1.BlockChain

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, err.Error()), nil)
	}

	chain, err := u.srv.Blockchains().Get(c, id)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	chain.ChainId = r.ChainId
	chain.InfraHttp = r.InfraHttp
	chain.ChainName = r.ChainName
	chain.InfraWebsocket = r.InfraWebsocket

	// Save changed fields.
	if err = u.srv.Blockchains().Update(c, chain); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, chain)
}
