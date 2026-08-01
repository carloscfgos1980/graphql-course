package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/carloscfgos1980/graphql-course/internal/graph/model"
	"github.com/google/uuid"
)

type CourseRepository struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

// CreateCourse creates a new course in the database and returns the created course.
func (c *CourseRepository) CreateCourse(title string, description string, categoryId string) (*model.Course, error) {
	// Generate a new UUID for the course and set the current time for created_at and updated_at
	now := time.Now()
	id := uuid.New().String()
	// Execute the SQL query to insert the new course into the database
	_, err := c.DB.Exec("INSERT INTO courses (id, title, description, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", id, title, description, categoryId, now, now)
	if err != nil {
		return nil, err
	}
	// Retrieve the newly created course from the database to return it
	var course model.Course
	course.Category = &model.Category{}
	err = c.DB.QueryRow(
		`SELECT c.id, c.title, c.description, c.created_at, c.updated_at,
		        cat.id, cat.name, cat.description, cat.created_at, cat.updated_at
		 FROM courses c
		 JOIN categories cat ON cat.id = c.category_id
		 WHERE c.id = $1`,
		id,
	).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.Category.ID,
		&course.Category.Name,
		&course.Category.Description,
		&course.Category.CreatedAt,
		&course.Category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	// Return the newly created course
	return &course, nil
}

func (c *CourseRepository) GetCoursesByCategoryID(categoryId string) ([]*model.Course, error) {
	rows, err := c.DB.Query(
		`SELECT c.id, c.title, c.description, c.created_at, c.updated_at,
		        cat.id, cat.name, cat.description, cat.created_at, cat.updated_at
		 FROM courses c
		 JOIN categories cat ON cat.id = c.category_id
		 WHERE c.category_id = $1`,
		categoryId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		var course model.Course
		course.Category = &model.Category{}
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.Category.ID,
			&course.Category.Name,
			&course.Category.Description,
			&course.Category.CreatedAt,
			&course.Category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *CourseRepository) GetAllCourses() ([]*model.Course, error) {
	rows, err := c.DB.Query(
		`SELECT c.id, c.title, c.description, c.created_at, c.updated_at,
		        cat.id, cat.name, cat.description, cat.created_at, cat.updated_at
		 FROM courses c
		 JOIN categories cat ON cat.id = c.category_id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course
	for rows.Next() {
		var course model.Course
		course.Category = &model.Category{}
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.Category.ID,
			&course.Category.Name,
			&course.Category.Description,
			&course.Category.CreatedAt,
			&course.Category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *CourseRepository) GetCourseByID(id string) (*model.Course, error) {
	var course model.Course
	course.Category = &model.Category{}
	err := c.DB.QueryRow(
		`SELECT c.id, c.title, c.description, c.created_at, c.updated_at,
		        cat.id, cat.name, cat.description, cat.created_at, cat.updated_at
		 FROM courses c
		 JOIN categories cat ON cat.id = c.category_id
		 WHERE c.id = $1`,
		id,
	).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.Category.ID,
		&course.Category.Name,
		&course.Category.Description,
		&course.Category.CreatedAt,
		&course.Category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No course found with the given ID
		}
		return nil, err
	}
	return &course, nil
}

func (c *CourseRepository) UpdateCourse(id string, title *string, description *string, categoryID *string) (*model.Course, error) {
	// SQL Fragment
	var setClauses []string
	// args is a slice to hold the values for the SQL query
	var args []interface{}
	// Validate and prepare the fields to update
	if title != nil {
		setClauses = append(setClauses, "title = ?")
		args = append(args, *title)
	}
	// Validate and prepare the fields to update
	if description != nil {
		setClauses = append(setClauses, "description = ?")
		args = append(args, *description)
	}
	if categoryID != nil {
		setClauses = append(setClauses, "category_id = ?")
		args = append(args, *categoryID)
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
	query := fmt.Sprintf("UPDATE courses SET %s WHERE id = ?", setClause)
	args = append(args, id)
	// Execute the update query
	_, err := c.DB.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update course: %w", err)
	}
	// Retrieve the updated course
	return c.GetCourseByID(id)
}

func (c *CourseRepository) DeleteCourse(id string) (bool, error) {
	result, err := c.DB.Exec("DELETE FROM courses WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete course: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to retrieve rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return false, nil // No course found with the given ID
	}
	return true, nil
}
