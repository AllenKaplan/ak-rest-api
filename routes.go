package main

import (
	"fmt"
	"strconv"

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
func getAllUsers(c *gin.Context) {
	resp, err := userSrv.GetAllUsers()

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

// GET /users/:id
func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := userSrv.Get(id)

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

	resp, err := authSrv.Create(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}

//PUT /user
func updateUser(c *gin.Context) {
	var user *user.User

	err := c.ShouldBind(&user)

	resp, err := userSrv.Update(user)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return

	}

	c.JSON(200, &resp)
}

//PUT /login
func updateLogin(c *gin.Context) {
	var login *auth.Login

	err := c.ShouldBind(&login)

	resp, err := authSrv.Update(login)

	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
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
	var request *auth.Token
	c.ShouldBind(&request)

	resp, err := authSrv.ValidateToken(request.Email, request.Token)
	if err != nil {
		c.JSON(500, fmt.Sprintf("%v", err))
		return
	}

	c.JSON(200, &resp)
}
