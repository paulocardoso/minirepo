package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func listUser(c *gin.Context) {
	var users []User
	c.JSON(http.StatusOK, users)

}



func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}