package pkg

import (
	"crypto/rand"
	"encoding/base64"
)

type OAuthStatePayload struct {
	CSRFToken string `json:"csrf_token"`
	Sub       string `json:"sub"`
}

func GenerateSecurityToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
