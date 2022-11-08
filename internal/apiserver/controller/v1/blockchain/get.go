package blockchain

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"iamgeek/internal/pkg/code"
	"iamgeek/internal/pkg/core"
	"iamgeek/pkg/errors"
	"iamgeek/pkg/log"
)

// Get 获取区块链详情
// @Summary 获取区块链详情
// @Description 获取区块链详情
// @Tags 区块链管理
// @Param id path int true "主键Id"
// @Success 200 {object} core.Response{data=v1.BlockChain} "{"code": 200, message:"ok", "data": [...]}"
// @Success 400 {object} core.Response
// @Success 404 {object} core.Response
// @Success 500 {object} core.Response
// @Router /v1/iamgeek/chains/{id} [get]
func (u *BlockchainController) Get(c *gin.Context) {
	log.L(c).Info("get chain function called.")
	param := c.Param("id") // 转int
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, err.Error()), nil)
	}

	user, err := u.srv.Blockchains().Get(c, id)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)
}
