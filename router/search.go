package router

import (
	"github.com/2336260845/hippo_search/handler"
	"github.com/gin-gonic/gin"
)

func SearchRegister(engine *gin.Engine) {
	//query接口 目前为空
	engine.GET("/search", handler.QuerySearch)

	//切词debug接口
	engine.GET("/search/query_analysis/debug", handler.QueryAnalysisDebug)

	//召回debug接口
	engine.GET("/search/recall/debug", handler.RecallDebug)
}
