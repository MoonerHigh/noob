package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	return engine
}
