package auth

import (
	"time"
	"zgo/pkg/domain"
)

// Storage provides the shape that a service should have to manipulate an user entity
// Implementation may use in memory, MySQL, Mongo and so on.
type Storage interface {
	FindUserByEmail(email string) (*domain.User, error)
}

// PasswordHasher provides the shape that a service should have to handle passwords in a safe manner
type PasswordHasher interface {
	Check(plainText, hash string) bool
	Make(plainText string) (string, error)
}

// TokenGenerator provides the shape that a service should have to generate
// an authentication token
type TokenGenerator interface {
	GenerateToken(user *domain.User) (string, time.Time, error)
	GetTokenType() string
}

// TokenAuthenticator is a service that performs the authentication
type TokenAuthenticator struct {
	dbStorage Storage
	hasher    PasswordHasher
	tokenGen  TokenGenerator
}

// Authenticate will authenticate a user by returning a token that can be used
// in subsequent requests
func (j *TokenAuthenticator) Authenticate(username, password string) (*domain.AuthToken, error) {
	user, err := j.dbStorage.FindUserByEmail(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.NewError("user_not_found", "user not found")
	}

	if !j.hasher.Check(user.Password, password) {
		return nil, domain.NewError("invalid_password", "invalid password")
	}

	token, expiresAt, err := j.tokenGen.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthToken{
		TokenType:   j.tokenGen.GetTokenType(),
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}

// NewTokenAuthenticator will create a new token based authenticator
func NewTokenAuthenticator(s Storage, h PasswordHasher, gen TokenGenerator) *TokenAuthenticator {
	return &TokenAuthenticator{
		dbStorage: s,
		hasher:    h,
		tokenGen:  gen,
	}
}
