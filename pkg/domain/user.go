package domain

import "time"

// User contains user information
type User struct {
	ID        int64
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
