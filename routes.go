package main

import (
	auth "github.com/allenkaplan/ak-rest-api/auth"
	user "github.com/allenkaplan/ak-rest-api/user"
	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}

func getUsers(c *gin.Context) {
	resp, _ := userSrv.Get()
	c.JSON(200, &resp)
}

func createUser(c *gin.Context) {
	var user *user.User
	c.ShouldBind(&user)
	resp, _ := userSrv.Create(user)
	c.JSON(200, &resp)
}

func login(c *gin.Context) {
	var login *auth.Login
	c.ShouldBind(&login)
	resp, _ := authSrv.Login(login)
	c.JSON(200, &resp)
}

func validate(c *gin.Context) {
	var token *auth.Token
	c.ShouldBind(&token)
	resp, _ := authSrv.Validate(token)
	c.JSON(200, &resp)
}
