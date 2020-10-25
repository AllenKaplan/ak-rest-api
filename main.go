package main

import (
	auth "github.com/allenkaplan/ak-rest-api/auth"
	user "github.com/allenkaplan/ak-rest-api/user"

	"github.com/gin-gonic/gin"
)

var (
	userSrv *user.UserService
	authSrv *auth.AuthService
)

func main() {
	router := gin.Default()

	userSrv = user.NewService()
	authSrv = auth.NewService()

	router.GET("/", homeHandler)
	router.GET("/user", getUsers)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.POST("/auth", validate)

	router.Run(":8080")
}
