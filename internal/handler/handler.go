package handler

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/handler/controller"
)

var (
	UserControleler *controller.UserControleler
)

func InitializeHandler(deps *configs.Deps) {
	UserControleler = controller.NewUserController(deps.Svc)
}
