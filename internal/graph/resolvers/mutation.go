package resolvers

import (
	"context"
	"fmt"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	var description string
	if input.Description != nil {
		description = *input.Description
	}

	category, err := r.CategoryDB.CreateCategory(input.Name, description)
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

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	var description string
	if input.Description != nil {
		description = *input.Description
	}

	course, err := r.CourseDB.CreateCourse(input.Title, description, input.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to create course: %w", err)
	}
	return course, nil
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, id string, title *string, description *string, categoryID *string) (*model.Course, error) {
	course, err := r.CourseDB.UpdateCourse(id, title, description, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to update course: %w", err)
	}
	if course == nil {
		return nil, fmt.Errorf("course not found")
	}
	return course, nil
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, id string) (bool, error) {
	deleted, err := r.CourseDB.DeleteCourse(id)
	if err != nil {
		return false, fmt.Errorf("failed to delete course: %w", err)
	}
	if !deleted {
		return false, fmt.Errorf("course not found")
	}
	return true, nil
}
