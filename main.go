package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"github.com/paulocardoso/repo-lite/routers"
)

var html = template.Must(template.New("http").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>`))

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetHTMLTemplate(html)

	r.GET("/user", listUser)

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "http", gin.H{
			"status": "success",
		})
	})

	return r
}

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

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}

