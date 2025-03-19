package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/go-ecommerce/internal/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUserInfoService(),
		"id":      id,
		"users":   []string{"Khoa", "Nam", "Tu Anh"},
	})
}
