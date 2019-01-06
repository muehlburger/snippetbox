package mysql

import (
	"database/sql"

	"github.com/muehlburger/snippetbox/pkg/models"
)

// SnippetModel wraps the sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a Snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get will return a specific Snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
