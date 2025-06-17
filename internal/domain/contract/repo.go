package contract

import (
	"context"

	"github.com/brunobotter/mercado-livre/internal/domain/entity"
	"github.com/brunobotter/mercado-livre/internal/response"
)

type RepoManager interface {
	UserRepo() UserRepository
	CategoryRepo() CategoryRepository
}

type UserRepository interface {
	Save(ctx context.Context, register entity.User) error
	FindByUsername(ctx context.Context, username string) (exist bool, err error)
}

type CategoryRepository interface {
	Save(ctx context.Context, category entity.Category, id *int64) (categoryResponse response.SaveCategoryResponse, err error)
	FindByCategory(ctx context.Context, category string) (exist bool, err error)
	FindByCategoryParent(ctx context.Context, category string) (id *int64, err error)
}
