package contract

import (
	"context"

	"github.com/brunobotter/mercado-livre/internal/domain/entity"
)

type RepoManager interface {
	UserRepo() UserRepository
}

type UserRepository interface {
	Save(ctx context.Context, register entity.User) error
	FindByUsername(ctx context.Context, username string) (exist bool, err error)
}
