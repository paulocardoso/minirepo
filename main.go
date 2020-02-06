package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paulocardoso/minirepo/config"
	"github.com/paulocardoso/minirepo/routers"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)

	var cfg = config.GetConfig()
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Port),
		Handler: routersInit,
	}

	var err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[info] start http server listening %s", cfg.Port)

}
