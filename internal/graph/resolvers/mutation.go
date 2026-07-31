package resolvers

import (
	"context"
	"fmt"

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

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, name *string, description *string) (*model.Category, error) {
	category, err := r.CategoryDB.UpdateCategory(id, name, description)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}
	if category == nil {
		return nil, fmt.Errorf("category not found")
	}
	return category, nil
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (bool, error) {
	deleted, err := r.CategoryDB.DeleteCategory(id)
	if err != nil {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}
	if !deleted {
		return false, fmt.Errorf("category not found")
	}
	return true, nil
}
