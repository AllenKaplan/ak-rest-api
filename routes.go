package main

import (
	"fmt"

	auth "github.com/allenkaplan/ak-rest-api/auth"
	user "github.com/allenkaplan/ak-rest-api/user"
	"github.com/gin-gonic/gin"
)

// the '/' endpoint
func homeHandler(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "hello"},
	)
}

// GET /users
func getUsers(c *gin.Context) {
	resp, err := userSrv.Get()

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

//POST /users
func createUser(c *gin.Context) {
	type CreateMessage struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var message *CreateMessage

	err := c.ShouldBind(&message)

	user := &user.User{
		Name:  message.Name,
		Email: message.Email,
	}

	userCreated, err := userSrv.Create(user)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	login := &auth.Login{
		UserID:   userCreated.UserID,
		Email:    message.Email,
		Password: message.Password,
	}

	token, err := authSrv.Create(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &token)
}

func login(c *gin.Context) {
	var login *auth.LoginRequest

	c.ShouldBind(&login)
	resp, err := authSrv.Login(login)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}
	c.JSON(200, &resp)
}

func validate(c *gin.Context) {
	var token *auth.Token
	c.ShouldBind(&token)
	resp, err := authSrv.ValidateToken(token)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}
	c.JSON(200, &resp)
}
