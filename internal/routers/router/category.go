package router

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeCategoryRouter(router *gin.Engine, deps *configs.Deps) {
	handler.InitializeHandler(deps)
	v1 := router.Group("api/v1/categories")
	{
		v1.POST("/save", handler.CategoryController.SaveNewCategory)
	}
}
