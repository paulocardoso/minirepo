package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paulocardoso/repo-lite/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	apiv1 := r.Group("/api/v1")
	apiv1.GET("/users", v1.GetUser)
	apiv1.GET("/repo", v1.GetRe)

	return r
}