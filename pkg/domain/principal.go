package domain

// Principal is the authenticated user for a request
// This may come from a JWT token, a session entry and so on
type Principal struct {
	ID    int64    `json:"id"`
	Roles []string `json:"roles"`
}
