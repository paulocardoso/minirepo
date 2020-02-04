package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func getRepository(c *gin.Context) {
	dir, err := ioutil.ReadDir("~/.m2/repository")
	if err != nil {
		log.Fatal(err)
	}

	var dirs []string

	for _, f := range dir {
		fmt.Println(f.Name())
		dirs = append(dirs, f.Name())
	}


	c.JSON(http.StatusOK,dirs)

}
