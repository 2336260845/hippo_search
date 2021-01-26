package router

import (
	"github.com/2336260845/hippo_search/handler"
	"github.com/gin-gonic/gin"
)

func DebugRegister(engine *gin.Engine) {
	engine.Any("/ping", handler.Ping)
}
