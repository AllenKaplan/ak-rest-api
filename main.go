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
	r := gin.Default()

	userSrv = &user.UserService{
		Database: "sql",
	}

	authSrv = &auth.AuthService{
		Database: "sql",
	}

	r.GET("/", homeHandler)
	r.GET("/user", getUsers)
	r.POST("/user", createUser)
	r.POST("/login", login)
	r.POST("/auth", validate)

	r.Run(":8080")
}
