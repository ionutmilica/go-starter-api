package auth

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"zgo/pkg/domain"
)

type mockedAuthenticator struct{}

func (mockedAuthenticator) Authenticate(username, password string) (*domain.AuthToken, error) {
	tm, _ := time.Parse(time.RFC3339, "2019-03-03T20:57:22.758160766Z")

	return &domain.AuthToken{
		TokenType:   "JWT",
		AccessToken: "blah",
		ExpiresAt:   tm,
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

	assert.Equal(t, w.Code, http.StatusBadRequest)
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

	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
	assert.JSONEq(t, issueMissingFieldExpected, w.Body.String())
}

const issueTokenOKReq = `
{
	"email": "dev@dev.com",
	"password": "password"
}
`

const issueTokenOKResp = `
{
    "status": "ok",
    "data": {
        "token_type": "JWT",
        "access_token": "blah",
        "expires_at": "2019-03-03T20:57:22.758160766Z"
    }
}
`

func TestIssueToken_Success(t *testing.T) {
	req := httptest.NewRequest("POST", "/tokens", bytes.NewReader([]byte(issueTokenOKReq)))
	w := httptest.NewRecorder()

	IssueToken(&mockedAuthenticator{}).ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.JSONEq(t, issueTokenOKResp, w.Body.String())
}
