package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/name", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
		}
		var req User
		c.BindJSON(&req)
		c.JSON(http.StatusOK, gin.H{"yourname": req.Name})
	})

	r.POST("/count", func(c *gin.Context) {
		type Count struct {
			Num int `json:"number"`
		}
		var req Count
		c.BindJSON(&req)
		c.JSON(http.StatusOK, gin.H{"numberplus": req.Num + 1})
	})

	r.Run()
}
