package router

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(router *gin.Engine, deps *configs.Deps) {
	handler.InitializeHandler(deps)
	v1 := router.Group("api/v1/users")
	{
		v1.POST("/register", handler.UserControleler.RegisterNewUser)
	}
}
