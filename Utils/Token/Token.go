package Token

import (
	"FinPro/Models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var blacklistedTokens = make(map[string]bool)

func TokenIsBlacklisted(tokenString string) bool {
	return blacklistedTokens[tokenString]
}

func BlacklistToken(tokenString string) {
	blacklistedTokens[tokenString] = true
}

func GenerateTokenString(form Models.LoginForm, secretKey string, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": form.Username,
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("Failed to create token")
	}

	return tokenString, nil
}

func InvalidateTokenString(tokenString string) error {
	blacklistedTokens[tokenString] = true
	return nil
}