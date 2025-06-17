package request

import (
	"time"

	"github.com/brunobotter/mercado-livre/internal/domain/entity"
)

type RegisterNewUserRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r RegisterNewUserRequest) ToEntity(hashedPassword string) entity.User {
	return entity.User{
		Username:     r.Username,
		Password:     hashedPassword,
		Registration: time.Now(),
	}
}
