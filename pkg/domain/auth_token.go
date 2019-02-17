package domain

import "time"

// AuthToken contains fields that are useful for a stateless authentication
type AuthToken struct {
	TokenType   string    `json:"token_type"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}
