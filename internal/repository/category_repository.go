package repository

import (
	"database/sql"
	"fmt"
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
