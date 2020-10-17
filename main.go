package main

import (
	user "github.com/allenkaplan/ak-rest-api/user"

	"github.com/gin-gonic/gin"
)

var users []*user.User

func main() {
	users = append(users, &user.User{Name: "Eli"})

	r := gin.Default()

	userSrv := &user.UserService{
		Database: "sql",
	}

	r.GET("/", homeHandler)
	r.GET("/user", userSrv.Get)
	r.POST("/user", userSrv.Create)

	r.GET("/messages", func(c *gin.Context) {
		message.wshandler(c.Writer, c.Request)
	})

	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}
