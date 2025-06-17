package routers

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/routers/router"
	"github.com/gin-gonic/gin"
)

func Initialize(deps *configs.Deps) {
	gin := gin.Default()
	router.InitializeUserRouter(gin, deps)
	router.InitializeCategoryRouter(gin, deps)
	gin.Run(":8080")
}
