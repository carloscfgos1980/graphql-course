package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
	"github.com/google/uuid"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (c *CategoryRepository) CreateCategory(name string, description string) (*model.Category, error) {
	now := time.Now()
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO categories (id, name, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", id, name, description, now, now)
	if err != nil {
		return nil, err
	}
	var category model.Category
	err = c.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created category: %w", err)
	}
	return &category, nil
}

func (c *CategoryRepository) GetCategories() ([]*model.Category, error) {
	rows, err := c.DB.Query("SELECT id, name, description, created_at, updated_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepository) GetCategoryByID(id string) (*model.Category, error) {
	var category model.Category
	err := c.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No category found with the given ID
		}
		return nil, fmt.Errorf("failed to retrieve category by ID: %w", err)
	}
	return &category, nil
}

func (c *CategoryRepository) UpdateCategory(id string, name *string, description *string) (*model.Category, error) {
	// SQL Fragment
	var setClauses []string
	// args is a slice to hold the values for the SQL query
	var args []interface{}
	// Validate and prepare the fields to update
	// Validate and prepare the fields to update
	if name != nil {
		setClauses = append(setClauses, "name = ?")
		args = append(args, *name)
	}
	// Validate and prepare the fields to update
	if description != nil {
		setClauses = append(setClauses, "description = ?")
		args = append(args, *description)
	}
	// If no fields are provided to update, return an error
	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields provided to update")
	}
	// Add the updated_at timestamp
	now := time.Now()
	setClauses = append(setClauses, "updated_at = ?")
	args = append(args, now)
	// setClause is a string that joins the setClauses with commas, forming the SET part of the SQL query
	setClause := strings.Join(setClauses, ", ")
	// query is the final SQL query string that will be executed to update the habit
	query := fmt.Sprintf("UPDATE categories SET %s WHERE id = ?", setClause)
	args = append(args, id)
	// Execute the update query
	_, err := c.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}
	// Retrieve the updated category
	return c.GetCategoryByID(id)
}

func (c *CategoryRepository) DeleteCategory(id string) (bool, error) {
	result, err := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
