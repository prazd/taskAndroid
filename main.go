package main

import (
	"net/http"

	"./mongo"

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

	r.POST("/signup", func(c *gin.Context) {
		type User struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		var user User
		c.BindJSON(&user)

		if len(user.Login) == 0 {
			c.JSON(http.StatusOK, gin.H{"resp": "empty login"})
		} else if len(user.Password) == 0 {
			c.JSON(http.StatusOK, gin.H{"resp": "empty password"})
		} else {
			resp := mongo.SetInfo(user.Login, user.Password)

			c.JSON(http.StatusOK, gin.H{"resp": resp})
		}
	})

	r.POST("/pass", func(c *gin.Context) {
		type Log struct {
			Login string `json:"login"`
		}
		var uLog Log
		c.BindJSON(&uLog)
		resp := mongo.GetPassword(uLog.Login)
		c.JSON(http.StatusOK, gin.H{"hashpass": resp})
	})

	r.Run()
}
