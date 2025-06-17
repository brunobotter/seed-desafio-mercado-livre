package service

import (
	"context"
	"errors"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/request"
	"github.com/brunobotter/mercado-livre/internal/response"
)

type categorySerice struct {
	svc contract.ServiceManager
}

func NewCategoryService(svc contract.ServiceManager) contract.CategoryService {
	return &categorySerice{
		svc: svc,
	}
}

func (s *categorySerice) SaveCategory(ctx context.Context, categoryRequest request.SaveCategoryRequest) (response response.SaveCategoryResponse, err error) {
	if categoryRequest.Name == "" {
		return response, errors.New("invalid category name")
	}
	exist, err := s.svc.DB().CategoryRepo().FindByCategory(ctx, categoryRequest.Name)
	if err != nil {
		return response, errors.New("error to find category")
	}
	if exist {
		return response, errors.New("category alread exists")
	}
	var idCategory *int64
	if categoryRequest.Parent != nil {
		idCategory, err = s.svc.DB().CategoryRepo().FindByCategoryParent(ctx, *categoryRequest.Parent)
		if err != nil {
			return response, errors.New("error to find category")
		}
	}
	category := categoryRequest.ToEntity()
	response, err = s.svc.DB().CategoryRepo().Save(ctx, category, idCategory)
	if err != nil {
		return response, errors.New("error to save category")
	}
	return response, nil
}
