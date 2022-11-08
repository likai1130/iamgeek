// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "iamgeek/cmd/wallet-doc-gen/docs"
	"iamgeek/internal/apiserver/controller/v1/blockchain"
	"iamgeek/internal/apiserver/store/database"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	gr := g.Group("")
	gr.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	storeIns, _ := database.GetDatabaseFactoryOr(nil, nil)
	v1 := g.Group("/v1/iamgeek")
	{
		// user RESTful resource
		chainv1 := v1.Group("/chains")
		{
			chainController := blockchain.NewBlockchainController(storeIns)
			chainv1.POST("", chainController.Create)
			chainv1.DELETE(":id", chainController.Delete)
			chainv1.PUT(":id", chainController.Update)
			chainv1.GET(":id", chainController.Get)
			chainv1.GET("", chainController.List)

		}
	}
	return g
}
