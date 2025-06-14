package routers

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/gin-gonic/gin"
)

func Initialize(deps *configs.Deps) {
	gin := gin.Default()

	gin.Run(":8080")
}
