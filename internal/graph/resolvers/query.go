package resolvers

import (
	"context"
	"fmt"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
)

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve categories: %w", err)
	}
	return categories, nil
}
