package persistence

import (
	"sync"
	"zgo/pkg/domain"
)

// MemoryUserPersistence is used for crud based operations with users in memory
type MemoryUserPersistence struct {
	mu    sync.Mutex
	users []domain.User
}

// FindUserByEmail will find a user by it's email
func (m *MemoryUserPersistence) FindUserByEmail(email string) (*domain.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, nil
}

// NewMemoryUserPersistence is a constructor function for MemoryUserPersistence
func NewMemoryUserPersistence() *MemoryUserPersistence {
	return &MemoryUserPersistence{
		users: []domain.User{
			{
				ID:       1,
				Email:    "dev@dev.com",
				Password: "$2a$10$Y6qzzswYdv3/V8WF8vp7b.MIS9VSsPPuhTHB/gPIht7wDMRC9P1mm", // dev
			},
			{
				ID:       2,
				Email:    "demo@demo.com",
				Password: "$2a$04$pkwn/t8n1LHXnWRo3unZ7.r3NC41fYnQHl/VADl6jbAmf3AwXa0G.", // demo
			},
		},
	}
}
