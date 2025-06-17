package datasql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/domain/entity"
)

type userRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *userRepository) Save(ctx context.Context, user entity.User) error {
	query := `
		INSERT INTO user (username, password, registration)
		VALUES (?, ?, ?)
	`
	_, err := r.conn.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Registration,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (r *userRepository) FindByUsername(ctx context.Context, user string) (exist bool, err error) {
	query := `
		SELECT id FROM user WHERE username = ? LIMIT 1;
	`
	var id int
	err = r.conn.QueryRowContext(ctx, query, user).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
