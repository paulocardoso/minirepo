package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Page struct {
	title string
	repo string
	host string
	path string
	param string
	dirs []string
}

func GetRepository(c *gin.Context) {

	param := c.Param("regex")

	var fPath = "/Users/pauloc/.m2/repository" + param

	info, err := os.Stat(fPath)

	if err != nil {

		if info == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"notfound": "notfound",
			})

			return
		}
		log.Fatal(err)
	}



	if info.IsDir() {
		listDir(fPath, c, param)
	} else {
		downloadFile(fPath, c)
	}
}

func listDir(fPath string, c *gin.Context, param string) {
	dirs, err := ioutil.ReadDir(fPath)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "repo/index.tmpl", gin.H{
		"title": "Posts",
		"repo":  "Mini Repo",
		"host":  c.Request.Host,
		"path":  c.Request.RequestURI,
		"param": param,
		"dirs":  dirs,
	})
}

func downloadFile(fPath string, c *gin.Context) {
	file, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Fatal(err)
	}

	var s = strings.Split(fPath, "/")
	var name = s[len(s)-1]

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+ name)
	c.Data(http.StatusOK, "application/octet-stream", file)
}


func UploadRepository(c *gin.Context){

	outFile, err := os.Create("/Users/pauloc/.m2/repository/" + c.Param("regex"))
	if err != nil{
		log.Fatal(err)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, c.Request.Body)


	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", "asd"))
}