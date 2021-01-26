package router

import (
	"github.com/2336260845/hippo_search/handler"
	"github.com/gin-gonic/gin"
)

func SearchRegister(engine *gin.Engine) {
	engine.GET("/search", handler.QuerySearch)
}
