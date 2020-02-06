package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocardoso/minirepo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/repo/*regex", v1.GetRepository)
	r.PUT("/repo/*regex", v1.UploadRepository)

	return r
}
