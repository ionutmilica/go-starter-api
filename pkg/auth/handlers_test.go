package auth

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
	"zgo/pkg/domain"
)

type mockedAuthenticator struct{}

func (mockedAuthenticator) Authenticate(username, password string) (*domain.AuthToken, error) {
	return &domain.AuthToken{
		TokenType:   "jwt",
		AccessToken: "blah",
		ExpiresAt:   time.Now(),
	}, nil
}

const issueBadReqResponse = `
{
   "status":"error",
   "error":{
      "type":"bad_request",
      "message":"Bad request data provided"
   }
}
`

func TestIssueToken_InvalidBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/tokens", nil)
	w := httptest.NewRecorder()

	IssueToken(nil).ServeHTTP(w, req)

	assert.JSONEq(t, issueBadReqResponse, w.Body.String())
}

const issueMissingFieldReq = `
{
	"email": "dev@dev.com"
}
`

const issueMissingFieldExpected = `
{
	"status": "error",
	"errors": [
		{
			"field": "Password",
			"message": "Key: 'issueTokenRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"
		}
	]
}
`

func TestIssueToken_ValidationError(t *testing.T) {
	req := httptest.NewRequest("POST", "/tokens", bytes.NewReader([]byte(issueMissingFieldReq)))
	w := httptest.NewRecorder()

	IssueToken(nil).ServeHTTP(w, req)

	assert.JSONEq(t, issueMissingFieldExpected, w.Body.String())
}
