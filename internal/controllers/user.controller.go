package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khoand3012/go-ecommerce/internal/services"
	"github.com/khoand3012/go-ecommerce/pkg/response"
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
	response.SuccessResponse(c, response.ErrCodeSuccess, []string{"Khoa", "Trang"})
}
