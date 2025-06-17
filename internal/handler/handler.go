package handler

import (
	"github.com/brunobotter/mercado-livre/configs"
	"github.com/brunobotter/mercado-livre/internal/handler/controller"
)

var (
	UserController     *controller.UserControleler
	CategoryController *controller.CategoryController
)

func InitializeHandler(deps *configs.Deps) {
	UserController = controller.NewUserController(deps.Svc)
	CategoryController = controller.NewCategoryController(deps.Svc)
}
