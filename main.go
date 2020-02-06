package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paulocardoso/repo-lite/routers"
	"net/http"
)



func main() {
	gin.SetMode(gin.DebugMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", 8080)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
