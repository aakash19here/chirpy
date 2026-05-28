package auth

import (
	"errors"
	"net/http"
	"strings"
)

const prefix = "Bearer "

func GetBearerToken(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")

	if len(authorizationHeader) < 1 {
		return "", errors.New("missing authorization header")
	}

	if !strings.HasPrefix(authorizationHeader, prefix) {
		return "", errors.New("invalid authorization header")
	}

	token := strings.TrimPrefix(authorizationHeader, prefix)

	if token == "" {
		return "", errors.New("missing bearer token")
	}

	return token, nil
}
