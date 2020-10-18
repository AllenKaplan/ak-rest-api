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
	type createMessage struct {
		userID   int    `json:"userID"`
		name     string `json:"name"`
		email    string `json:"email"`
		password string `json:"password"`
	}

	var message *createMessage

	c.ShouldBind(&message)

	user := &user.User{
		UserID: message.userID,
		Name:   message.name,
		Email:  message.email,
	}

	login := &auth.Login{
		UserID:   message.userID,
		Email:    message.email,
		Password: message.password,
	}

	_, _ = userSrv.Create(user)
	token, _ := authSrv.Create(login)

	c.JSON(200, &token)
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
	resp, _ := authSrv.ValidateToken(token)
	c.JSON(200, &resp)
}
