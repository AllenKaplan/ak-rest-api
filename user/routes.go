package user

import "github.com/gin-gonic/gin"

func (s UserService) Create(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	resp, _ := create(user)
	c.JSON(200, &resp)
}

func (s UserService) Get(c *gin.Context) {
	resp, _ := get()

	c.JSON(200, &resp)
}
