package service

import (
	"context"
	"errors"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/util"
)

type userService struct {
	svc contract.ServiceManager
}

func NewUserService(svc contract.ServiceManager) contract.UserService {
	return &userService{
		svc: svc,
	}
}

func (s *userService) Register(ctx context.Context, registerRequest request.RegisterNewUserRequest) error {
	if registerRequest.Password == "" || len(registerRequest.Password) < 6 {
		return errors.New("invalid password")
	}
	if registerRequest.Username == "" {
		return errors.New("invalid username")
	}
	exist, err := s.svc.DB().UserRepo().FindByUsername(ctx, registerRequest.Username)
	if err != nil {
		return errors.New("error to find username")
	}
	if exist {
		return errors.New("username alread register")
	}
	if !util.IsValidEmail(registerRequest.Username) {
		return errors.New("invalid username format. ex: aaa@email.com")
	}
	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		return errors.New("error in encripty data")
	}
	register := registerRequest.ToEntity(hashedPassword)
	err = s.svc.DB().UserRepo().Save(ctx, register)
	if err != nil {
		return errors.New("error to save user")
	}
	return nil
}
