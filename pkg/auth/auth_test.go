package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"zgo/pkg/domain"
)

type mockedDB struct {
	err             error
	findByEmailUser *domain.User
}

func (s mockedDB) FindUserByEmail(email string) (*domain.User, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.findByEmailUser, nil
}

type mockedHasher struct {
	check bool
}

func (h mockedHasher) Check(plainText, hash string) bool {
	return h.check
}

func (mockedHasher) Make(plainText string) (string, error) {
	return "", nil
}

type mockedTokenGen struct {
	err      error
	token    string
	expireAt time.Time
}

func (m mockedTokenGen) GenerateToken(user *domain.User) (string, time.Time, error) {
	if m.err != nil {
		return "", time.Now(), m.err
	}

	return m.token, m.expireAt, nil
}

func (mockedTokenGen) GetTokenType() string { return "type" }

func TestTokenAuthenticator_Authenticate(t *testing.T) {
	t.Run("Authentication should succeed in the happy flow", func(t *testing.T) {
		now := time.Now().UTC()

		a := TokenAuthenticator{
			dbStorage: &mockedDB{findByEmailUser: &domain.User{}},
			hasher:    &mockedHasher{check: true},
			tokenGen:  &mockedTokenGen{token: "token", expireAt: now},
		}

		tk, err := a.Authenticate("ion", "pass")
		assert.Nil(t, err)
		assert.NotNil(t, tk)

		assert.Equal(t, tk.AccessToken, "token")
		assert.Equal(t, tk.ExpiresAt, now)
		assert.Equal(t, tk.TokenType, "type")
	})

	t.Run("Authentication return error when user is not found", func(t *testing.T) {
		a := TokenAuthenticator{
			dbStorage: &mockedDB{},
			hasher:    &mockedHasher{},
		}

		tk, err := a.Authenticate("ion", "pass")
		assert.NotNil(t, err)
		assert.Nil(t, tk)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("Authentication return error when password is invalid", func(t *testing.T) {
		a := TokenAuthenticator{
			dbStorage: &mockedDB{findByEmailUser: &domain.User{}},
			hasher:    &mockedHasher{check: false},
		}

		tk, err := a.Authenticate("ion", "pass")
		assert.NotNil(t, err)
		assert.Nil(t, tk)
		assert.Contains(t, err.Error(), "invalid password")
	})
}
