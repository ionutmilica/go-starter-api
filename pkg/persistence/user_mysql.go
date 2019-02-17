package persistence

import (
	"database/sql"
	"zgo/pkg/domain"
)

// MySQLUserPersistence is used for crud based operations with users in MySQL database
type MySQLUserPersistence struct {
	db *sql.DB
}

// FindUserByEmail will find a user by it's email
func (p *MySQLUserPersistence) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

// NewMySQLUserPersistence is a constructor function for MySQLUserPersistence
func NewMySQLUserPersistence(db *sql.DB) *MySQLUserPersistence {
	return &MySQLUserPersistence{
		db: db,
	}
}
