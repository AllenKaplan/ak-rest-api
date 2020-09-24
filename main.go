package main

import (
	"github.com/gin-gonic/gin"
)

var users []*User

func main() {
	users = append(users, &User{Name: "Eli"})

	r := gin.Default()
	r.GET("/", homeHandler)
	r.POST("/user", create)
	r.GET("/user", get)

	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}

func create(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	users = append(users, &user)
}

func get(c *gin.Context) {
	c.JSON(200, users)
}
