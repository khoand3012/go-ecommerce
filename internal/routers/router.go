package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/khoand3012/go-ecommerce/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	pc := c.NewPongController()
	uc := c.NewUserController()

	v1 := r.Group("/v1/2025")
	{
		v1.GET("/ping", pc.Pong)
		v1.GET("/user/:id", uc.GetUserById)
	}
	return r
}
