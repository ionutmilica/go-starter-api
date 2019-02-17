package auth

import (
	"log"
	"net/http"
	"zgo/pkg/api"
	"zgo/pkg/domain"
)

type issueTokenRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type issueTokenResponse struct {
	Status string           `json:"status"`
	Data   domain.AuthToken `json:"data"`
}

// Authenticator provides the shape that a service should have
// so the http handlers will perform the authentication phase
type Authenticator interface {
	Authenticate(username, password string) (*domain.AuthToken, error)
}

// IssueToken will create a http handler that can issue new authentication
// tokens based on email and password
func IssueToken(auth Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req issueTokenRequest

		validationErrors, err := api.BindJSON(r, &req)
		if err != nil {
			log.Printf("Error parsing the request: %s", err.Error())
			api.SendBadRequest(w)
			return
		}

		if len(validationErrors) > 0 {
			api.SendValidationErrors(w, validationErrors)
			return
		}

		token, err := auth.Authenticate(req.Email, req.Password)

		switch err := err.(type) {
		case *domain.Error:
			api.SendCustomError(w, err)
			return
		case error:
			log.Printf("Error authenticating: %s", err)
			api.SendInternalError(w)
			return
		}

		api.SendJSON(w, http.StatusOK, issueTokenResponse{
			Status: "ok",
			Data:   *token,
		})
	}
}

// CurrentUser will create a new handler that returns the current user
// Notice: This should be protected with an auth middleware
func CurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
