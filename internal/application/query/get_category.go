package query

import (
	"context"

	"github.com/neutrinocorp/life-track-api/internal/domain/model"
	"github.com/neutrinocorp/life-track-api/internal/domain/repository"
	"github.com/neutrinocorp/life-track-api/internal/domain/value"
)

// GetCategory request a single category
type GetCategory struct {
	repo repository.Category
}

// NewGetCategory get a new get category query
func NewGetCategory(r repository.Category) *GetCategory {
	return &GetCategory{
		repo: r,
	}
}

func (q GetCategory) Query(ctx context.Context, id string) (*model.Category, error) {
	idCUID := new(value.CUID)
	if err := idCUID.Set(id); err != nil {
		return nil, err
	}

	return q.repo.FetchByID(ctx, *idCUID)
}
