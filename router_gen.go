package main

import (
	"github.com/2336260845/hippo_search/router"
	"github.com/gin-gonic/gin"
)

func register(engine *gin.Engine) {
	//debug接口
	router.DebugRegister(engine)

	//搜索主接口
	router.SearchRegister(engine)
}
