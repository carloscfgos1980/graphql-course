package resolvers

import (
	"context"
	"fmt"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
)

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	// Retrieve all categories from the database
	categories, err := r.CategoryDB.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve categories: %w", err)
	}
	// For each category, retrieve the associated courses and populate the Courses field
	for _, category := range categories {
		courses, err := r.CourseDB.GetCoursesByCategoryID(category.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve courses for category %s: %w", category.ID, err)
		}
		category.Courses = courses
	}
	// Return the list of categories with their associated courses
	return categories, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	// Retrieve a category by its ID from the database
	category, err := r.CategoryDB.GetCategoryByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve category: %w", err)
	}
	// If the category is not found, return an error
	if category == nil {
		return nil, fmt.Errorf("category not found")
	}
	// Retrieve the courses associated with the category and populate the Courses field
	courses, err := r.CourseDB.GetCoursesByCategoryID(category.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve courses for category %s: %w", category.ID, err)
	}
	// Populate the Courses field of the category with the associated courses
	category.Courses = courses
	// Return the category with its associated courses (if any)
	return category, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.GetAllCourses()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve courses: %w", err)
	}
	return courses, nil
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	course, err := r.CourseDB.GetCourseByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve course: %w", err)
	}
	if course == nil {
		return nil, fmt.Errorf("course not found")
	}
	return course, nil
}
