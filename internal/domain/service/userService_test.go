package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/mercado-livre/internal/domain/service"
	"github.com/brunobotter/mercado-livre/internal/mock"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserService_Register_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	// Mocks
	userRepoMock := mock.NewMockUserRepository(ctrl)
	dbManagerMock := mock.NewMockRepoManager(ctrl)
	serviceManagerMock := mock.NewMockServiceManager(ctrl)

	// Inputs
	req := request.RegisterNewUserRequest{
		Username: "test@email.com",
		Password: "123456",
	}

	// Configurar expectations dos mocks
	userRepoMock.EXPECT().FindByUsername(ctx, req.Username).Return(false, nil)
	userRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)

	// Configurar o retorno do DBManager -> UserRepo
	dbManagerMock.EXPECT().UserRepo().Return(userRepoMock).AnyTimes()

	// Configurar o retorno do ServiceManager -> DB
	serviceManagerMock.EXPECT().DB().Return(dbManagerMock).AnyTimes()

	// Criar service com mocks
	userService := service.NewUserService(serviceManagerMock)

	// Executar
	err := userService.Register(ctx, req)

	// Assert
	require.NoError(t, err)
}

func newMocks(t *testing.T) (*gomock.Controller, *mock.MockUserRepository, *mock.MockRepoManager, *mock.MockServiceManager) {
	ctrl := gomock.NewController(t)
	userRepoMock := mock.NewMockUserRepository(ctrl)
	dbManagerMock := mock.NewMockRepoManager(ctrl)
	serviceManagerMock := mock.NewMockServiceManager(ctrl)

	dbManagerMock.EXPECT().UserRepo().Return(userRepoMock).AnyTimes()
	serviceManagerMock.EXPECT().DB().Return(dbManagerMock).AnyTimes()

	return ctrl, userRepoMock, dbManagerMock, serviceManagerMock
}

func TestRegister_InvalidPassword(t *testing.T) {
	ctrl, _, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	// Teste com senha vazia
	req := request.RegisterNewUserRequest{
		Username: "test@email.com",
		Password: "",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "invalid password")

	// Teste com senha menor que 6 caracteres
	req.Password = "123"
	err = userService.Register(context.Background(), req)
	require.EqualError(t, err, "invalid password")
}

func TestRegister_InvalidUsername(t *testing.T) {
	ctrl, _, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	req := request.RegisterNewUserRequest{
		Username: "",
		Password: "123456",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "invalid username")
}

func TestRegister_InvalidEmailFormat(t *testing.T) {
	ctrl, userRepoMock, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	// Simula usuário não existente (para passar o FindByUsername)
	userRepoMock.EXPECT().FindByUsername(gomock.Any(), "invalid-email").Return(false, nil)

	req := request.RegisterNewUserRequest{
		Username: "invalid-email",
		Password: "123456",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "invalid username format. ex: aaa@email.com")
}

func TestRegister_UsernameAlreadyExists(t *testing.T) {
	ctrl, userRepoMock, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	userRepoMock.EXPECT().FindByUsername(gomock.Any(), "test@email.com").Return(true, nil)

	req := request.RegisterNewUserRequest{
		Username: "test@email.com",
		Password: "123456",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "username alread register")
}

func TestRegister_ErrorOnFindByUsername(t *testing.T) {
	ctrl, userRepoMock, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	userRepoMock.EXPECT().FindByUsername(gomock.Any(), "test@email.com").Return(false, errors.New("db error"))

	req := request.RegisterNewUserRequest{
		Username: "test@email.com",
		Password: "123456",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "error to find username")
}

func TestRegister_ErrorOnSave(t *testing.T) {
	ctrl, userRepoMock, _, serviceManagerMock := newMocks(t)
	defer ctrl.Finish()

	userService := service.NewUserService(serviceManagerMock)

	userRepoMock.EXPECT().FindByUsername(gomock.Any(), "test@email.com").Return(false, nil)
	userRepoMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("db save error"))

	req := request.RegisterNewUserRequest{
		Username: "test@email.com",
		Password: "123456",
	}

	err := userService.Register(context.Background(), req)
	require.EqualError(t, err, "error to save user")
}
