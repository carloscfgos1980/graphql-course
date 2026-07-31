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
	for _, category := range categories {
		courses, err := r.CourseDB.GetCoursesByCategoryID(category.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve courses for category %s: %w", category.ID, err)
		}
		category.Courses = courses
	}
	return categories, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	category, err := r.CategoryDB.GetCategoryByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve category: %w", err)
	}
	if category == nil {
		return nil, fmt.Errorf("category not found")
	}
	return category, nil
}
