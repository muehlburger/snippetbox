package mysql

import (
	"database/sql"

	"github.com/muehlburger/snippetbox/pkg/models"
)

// UserModel wraps the sql.DB connection pool
type UserModel struct {
	DB *sql.DB
}

// Insert adds a new record to the user table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate verifies that a user exists with the given email address
// and that the password matches the stored hash.
// Returns the User ID on success.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get returns details on a specific user based on the user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
