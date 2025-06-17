package request

import "github.com/brunobotter/mercado-livre/internal/domain/entity"

type SaveCategoryRequest struct {
	Name   string  `json:"name" validate:"required"`
	Parent *string `json:"parent_name,omitempty"`
}

func (r SaveCategoryRequest) ToEntity() entity.Category {
	return entity.Category{
		Name:       r.Name,
		ParentName: r.Parent,
	}
}
