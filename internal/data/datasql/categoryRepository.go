package datasql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/brunobotter/mercado-livre/internal/domain/entity"
	"github.com/brunobotter/mercado-livre/internal/response"
)

type categoryRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *categoryRepository) Save(ctx context.Context, category entity.Category, id *int64) (categoryResponse response.SaveCategoryResponse, err error) {
	created := time.Now()
	query := `
		INSERT INTO categories (name, parent_id, created_at)
		VALUES (?, ?, ?)
	`
	_, err = r.conn.ExecContext(
		ctx,
		query,
		category.Name,
		id,
		created,
	)
	if err != nil {
		return categoryResponse, fmt.Errorf("failed to insert category: %v", err)
	}
	categoryResponse = response.SaveCategoryResponse{
		Name:       category.Name,
		ParentName: category.ParentName,
	}
	return categoryResponse, nil
}

func (r *categoryRepository) FindByCategory(ctx context.Context, category string) (exist bool, err error) {
	query := `
		SELECT id FROM categories WHERE name = ? LIMIT 1;
	`
	var id int
	err = r.conn.QueryRowContext(ctx, query, category).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *categoryRepository) FindByCategoryParent(ctx context.Context, category string) (id *int64, err error) {
	query := `SELECT id FROM categories WHERE name = ? LIMIT 1;`
	err = r.conn.QueryRowContext(ctx, query, category).Scan(&id)
	if err == sql.ErrNoRows {
		return id, err
	}
	if err != nil {
		return id, err
	}
	return id, nil
}
