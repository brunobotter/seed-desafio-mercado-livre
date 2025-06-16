package contract

import (
	"context"

	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/response"
)

type ServiceManager interface {
	Config() *mapping.Config
	DB() DataManager
	InternalService() InternalService
}

type InternalService interface {
	UserService() UserService
	CategoryService() CategoryService
}

type UserService interface {
	Register(ctx context.Context, register request.RegisterNewUserRequest) error
}

type CategoryService interface {
	SaveCategory(ctx context.Context, category request.SaveCategoryRequest) (response.SaveCategoryResponse, error)
}
