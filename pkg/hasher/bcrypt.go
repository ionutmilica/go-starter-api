package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// BCryptHasher is used to hash secrets
type BCryptHasher struct {
	Cost int
}

// Check will compare a hash against the plain text value
func (h BCryptHasher) Check(hash, plainText string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText)) == nil
}

// Make will create a new hash for a secret that's in plain text
func (h BCryptHasher) Make(plainText string) (string, error) {
	cost := bcrypt.DefaultCost
	if h.Cost != 0 {
		cost = h.Cost
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// NewBCryptHasher is a constructor function for BCryptHasher
func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{
		Cost: 10,
	}
}
