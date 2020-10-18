package main

import (
	"fmt"

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
	type CreateMessage struct {
		name     string `json:"name"`
		email    string `json:"email"`
		password string `json:"password"`
	}

	var message *CreateMessage

	c.ShouldBind(&message)

	user := &user.User{
		Name:  message.name,
		Email: message.email,
	}

	userCreated, _ := userSrv.Create(user)

	login := &auth.Login{
		UserID:   userCreated.UserID,
		Email:    message.email,
		Password: message.password,
	}

	token, _ := authSrv.Create(login)

	c.JSON(200, &token)
}

func login(c *gin.Context) {
	//currently login requires userID --> change with SQL querries
	// type LoginRequest struct {
	// 	email    string `json:"email"`
	// 	password string `json:"password"`
	// }
	var login *auth.Login
	c.ShouldBind(&login)
	resp, err := authSrv.Login(login)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	c.JSON(200, &resp)
}

func validate(c *gin.Context) {
	var token *auth.Token
	c.ShouldBind(&token)
	resp, _ := authSrv.ValidateToken(token)
	c.JSON(200, &resp)
}
