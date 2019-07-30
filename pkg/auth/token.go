package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/o1egl/paseto"
)

const refreshedAtLayout = "2006-01-02T15:04:05.000"

var v2 = paseto.NewV2()

func generateToken(userId int64, key []byte) (string, error) {
	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(time.Hour)
	refreshesAt := issuedAt.Add(time.Minute * 10)

	jsonToken := paseto.JSONToken{
		Audience:   "test",
		Issuer:     "test_service",
		Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   issuedAt,
		Expiration: expiresAt,
		NotBefore:  issuedAt,
	}
	jsonToken.Set("userId", strconv.FormatInt(userId, 10))
	jsonToken.Set("refreshedAt", refreshesAt.Format(refreshedAtLayout))
	footer := ""

	return v2.Encrypt(key, jsonToken, footer)
}

func decryptToken(token string, key []byte) (int64, bool) {
	var (
		jsonToken paseto.JSONToken
		footer    string
	)

	if err := v2.Decrypt(token, key, &jsonToken, &footer); err != nil {
		return 0, false
	}

	if time.Now().UTC().After(jsonToken.Expiration) {
		fmt.Println("The token was expired")
		return 0, false
	}

	value := jsonToken.Get("userId")
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Could not parse userId from token:", err)
		return 0, false
	}

	refreshedAtValue := jsonToken.Get("refreshedAt")
	refreshedAt, err := time.Parse(refreshedAtLayout, refreshedAtValue)
	if err != nil {
		fmt.Println("Could not parse refreshedAt from token:", err)
		return 0, false
	}

	refresh := time.Now().UTC().After(refreshedAt)
	return userId, refresh
}
