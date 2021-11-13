package pkg

import (
	"math/rand"
	"strconv"
	"strings"
)

func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}

func RandomID() string {
	return strconv.Itoa(rand.Int())
}
