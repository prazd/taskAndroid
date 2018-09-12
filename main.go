package main

import (
	"net/http"

	"./mongo"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/name", func(c *gin.Context) { // POST request - {"name":"..."} to http://localhost:8080/name
		type User struct {
			Name string `json:"name"`
		}
		var req User
		c.BindJSON(&req)
		c.JSON(http.StatusOK, gin.H{"yourname": req.Name}) // response - {"yourname":"..."}
	})

	r.POST("/count", func(c *gin.Context) { // POST request- {"number":"..."} to http://localhost:8080/count
		type Count struct {
			Num int `json:"number"`
		}
		var req Count
		c.BindJSON(&req)
		c.JSON(http.StatusOK, gin.H{"numberplus": req.Num + 1}) // response - {"numberplus":"...+1"}
	})

	r.POST("/signup", func(c *gin.Context) { // POST request- {"login":"..."} to http://localhost:8080/signup
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
			c.JSON(http.StatusOK, gin.H{"resp": resp}) // response - {"resp":true}
		}
	})

	r.POST("/pass", func(c *gin.Context) { // POST request - {"lgin":"..."} to http://localhost:8080/signup
		type Log struct {
			Login string `json:"login"`
		}
		var uLog Log

		c.BindJSON(&uLog)

		if len(uLog.Login) == 0 {
			c.JSON(http.StatusOK, gin.H{"resp": "empty login"})
		} else {
			resp := mongo.GetPassword(uLog.Login)
			c.JSON(http.StatusOK, gin.H{"hashpass": resp}) // response {"hashpass":"..."}
		}

	})

	r.Run()
}
