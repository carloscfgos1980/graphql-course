package graph

import (
	"github.com/carloscfgos1980/graphql-course/internal/database"
)

type Resolver struct {
	CategoryDB *database.Category
}
