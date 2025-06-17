package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/mercado-livre/internal/domain/service"
	"github.com/brunobotter/mercado-livre/internal/mock"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/response"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestSaveCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	categoryName := "Smartphones"

	repo := mock.NewMockCategoryRepository(ctrl)
	db := mock.NewMockRepoManager(ctrl)
	svc := mock.NewMockServiceManager(ctrl)

	req := request.SaveCategoryRequest{
		Name: categoryName,
	}

	repo.EXPECT().FindByCategory(ctx, categoryName).Return(false, nil)
	repo.EXPECT().Save(ctx, gomock.Any(), nil).Return(response.SaveCategoryResponse{Name: categoryName}, nil)
	db.EXPECT().CategoryRepo().Return(repo).AnyTimes()
	svc.EXPECT().DB().Return(db).AnyTimes()

	service := service.NewCategoryService(svc)
	res, err := service.SaveCategory(ctx, req)

	require.NoError(t, err)
	require.Equal(t, categoryName, res.Name)
}

func TestSaveCategory_InvalidName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mock.NewMockServiceManager(ctrl)
	service := service.NewCategoryService(svc)

	req := request.SaveCategoryRequest{
		Name: "",
	}

	_, err := service.SaveCategory(context.Background(), req)
	require.EqualError(t, err, "invalid category name")
}

func TestSaveCategory_AlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	name := "Smartphones"

	repo := mock.NewMockCategoryRepository(ctrl)
	db := mock.NewMockRepoManager(ctrl)
	svc := mock.NewMockServiceManager(ctrl)

	repo.EXPECT().FindByCategory(ctx, name).Return(true, nil)
	db.EXPECT().CategoryRepo().Return(repo).AnyTimes()
	svc.EXPECT().DB().Return(db).AnyTimes()

	service := service.NewCategoryService(svc)
	_, err := service.SaveCategory(ctx, request.SaveCategoryRequest{Name: name})
	require.EqualError(t, err, "category alread exists")
}

func TestSaveCategory_FindByCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	name := "Smartphones"

	repo := mock.NewMockCategoryRepository(ctrl)
	db := mock.NewMockRepoManager(ctrl)
	svc := mock.NewMockServiceManager(ctrl)

	repo.EXPECT().FindByCategory(ctx, name).Return(false, errors.New("db error"))
	db.EXPECT().CategoryRepo().Return(repo).AnyTimes()
	svc.EXPECT().DB().Return(db).AnyTimes()

	service := service.NewCategoryService(svc)
	_, err := service.SaveCategory(ctx, request.SaveCategoryRequest{Name: name})
	require.EqualError(t, err, "error to find category")
}

func TestSaveCategory_ParentCategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	name := "Android"
	parent := "Smartphones"

	repo := mock.NewMockCategoryRepository(ctrl)
	db := mock.NewMockRepoManager(ctrl)
	svc := mock.NewMockServiceManager(ctrl)

	repo.EXPECT().FindByCategory(ctx, name).Return(false, nil)
	repo.EXPECT().FindByCategoryParent(ctx, parent).Return(nil, errors.New("not found"))
	db.EXPECT().CategoryRepo().Return(repo).AnyTimes()
	svc.EXPECT().DB().Return(db).AnyTimes()

	service := service.NewCategoryService(svc)
	_, err := service.SaveCategory(ctx, request.SaveCategoryRequest{Name: name, Parent: &parent})
	require.EqualError(t, err, "error to find category")
}

func TestSaveCategory_ErrorOnSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	name := "Android"

	repo := mock.NewMockCategoryRepository(ctrl)
	db := mock.NewMockRepoManager(ctrl)
	svc := mock.NewMockServiceManager(ctrl)

	repo.EXPECT().FindByCategory(ctx, name).Return(false, nil)
	repo.EXPECT().Save(ctx, gomock.Any(), nil).Return(response.SaveCategoryResponse{}, errors.New("insert failed"))
	db.EXPECT().CategoryRepo().Return(repo).AnyTimes()
	svc.EXPECT().DB().Return(db).AnyTimes()

	service := service.NewCategoryService(svc)
	_, err := service.SaveCategory(ctx, request.SaveCategoryRequest{Name: name})
	require.EqualError(t, err, "error to save category")
}
