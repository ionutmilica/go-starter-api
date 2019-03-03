package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"zgo/pkg/domain"
)

// JWTTokenGenerator is a service that will create a jwt token for a given user
type JWTTokenGenerator struct {
	SecretKey     string
	TokenLifetime time.Duration
}

type customClaims struct {
	*jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

// GenerateToken will construct a temporary jwt token
func (j *JWTTokenGenerator) GenerateToken(user *domain.User) (string, time.Time, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(j.TokenLifetime)

	generator := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
		},
		UserID: user.ID,
	})

	tk, err := generator.SignedString([]byte(j.SecretKey))

	return tk, expiresAt, err
}

// GetTokenType allows us to identify the algorithm of the token generator
func (j *JWTTokenGenerator) GetTokenType() string {
	return "JWT"
}

// NewJwtTokenGenerator is a constructor function for JWTTokenGenerator
func NewJwtTokenGenerator(key string, lifetime time.Duration) *JWTTokenGenerator {
	return &JWTTokenGenerator{
		SecretKey:     key,
		TokenLifetime: lifetime,
	}
}
