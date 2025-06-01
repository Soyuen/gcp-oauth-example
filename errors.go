package gcpoauthexample

import "errors"

type ErrorResponse struct {
	Code int `json:"code"`
}

const (
	ErrCodeOAuthConfigMissing = 10001
	ErrCodeGenerateOAuthState = 10002
)

var (
	ErrOAuthConfigMissing = errors.New("oauth config is missing or not initialized")
	ErrGenerateOAuthState = errors.New("failed to generate OAuth state")
)
