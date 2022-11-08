package blockchain

import (
	"github.com/gin-gonic/gin"
	metav1 "iamgeek/internal/apiserver/meta/v1"
	"iamgeek/internal/pkg/code"
	"iamgeek/internal/pkg/core"
	"iamgeek/pkg/errors"
	"iamgeek/pkg/log"
)

// List 获取区块链列表
// @Summary 获取区块链列表
// @Description 获取接口管理列表
// @Tags 区块链管理
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {object} core.Response{data=v1.BlockChainList} "{"code": 200, "message":"Ok","data": [...]}"
// @Success 500 {object} core.Response
// @Router /v1/iamgeek/chains [get]
func (u *BlockchainController) List(c *gin.Context) {
	log.L(c).Info("list chain function called.")

	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	chains, err := u.srv.Blockchains().List(c, r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, chains)
}
