package contract

import (
	"context"

	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/request"
)

type ServiceManager interface {
	Config() *mapping.Config
	DB() DataManager
	InternalService() InternalService
}

type InternalService interface {
	UserService() UserService
}

type UserService interface {
	Register(ctx context.Context, register request.RegisterNewUserRequest) error
}
