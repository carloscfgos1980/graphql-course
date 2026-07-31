package repository

import (
	"database/sql"
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

func (c *CourseRepository) CreateCourse(title string, description string, categoryId string) (*model.Course, error) {
	now := time.Now()
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO courses (id, title, description, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", id, title, description, categoryId, now, now)
	if err != nil {
		return nil, err
	}
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
	return &course, nil
}
