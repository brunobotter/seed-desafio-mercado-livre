package contract

import "github.com/brunobotter/mercado-livre/configs/mapping"

type ServiceManager interface {
	Config() *mapping.Config
	DB() DataManager
	InternalService() InternalService
}

type InternalService interface {
}
