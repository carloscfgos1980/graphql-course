package resolvers

import (
	"context"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.CreateCategory(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}
	return category, nil
}
