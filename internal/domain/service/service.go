package service

import (
	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/domain/contract"
)

type serviceManager struct {
	cfg             *mapping.Config
	db              contract.DataManager
	internalService contract.InternalService
}

func (s *serviceManager) Config() *mapping.Config {
	return s.cfg
}

func (s *serviceManager) DB() contract.DataManager {
	return s.db
}
func (s *serviceManager) InternalService() contract.InternalService {
	return s.internalService
}

type internalServices struct {
	userService     contract.UserService
	categoryService contract.CategoryService
}

type ServiceDeps struct {
	Cfg *mapping.Config
	DB  contract.DataManager
}

func New(deps ServiceDeps) (contract.ServiceManager, error) {
	instance := &serviceManager{
		db:  deps.DB,
		cfg: deps.Cfg,
	}

	internalServices := internalServices{}
	internalServices.userService = NewUserService(instance)
	internalServices.categoryService = NewCategoryService(instance)

	instance.internalService = &internalServices
	return instance, nil
}

func (s *internalServices) UserService() contract.UserService {
	return s.userService
}

func (s *internalServices) CategoryService() contract.CategoryService {
	return s.categoryService
}
