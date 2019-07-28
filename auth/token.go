package auth

import (
	"errors"
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
	fmt.Println("will refresh at:", refreshesAt.Format(refreshedAtLayout))
	jsonToken.Set("userId", strconv.FormatInt(userId, 10))
	jsonToken.Set("refreshedAt", refreshesAt.Format(refreshedAtLayout))
	footer := ""

	return v2.Encrypt(key, jsonToken, footer)
}

func decryptToken(token string, key []byte) (int64, bool, error) {
	var (
		jsonToken paseto.JSONToken
		footer    string
	)

	if err := v2.Decrypt(token, key, &jsonToken, &footer); err != nil {
		return 0, false, err
	}

	if time.Now().UTC().After(jsonToken.Expiration) {
		fmt.Println("The token was expired")
		return 0, false, errors.New("Token expired")
	}

	value := jsonToken.Get("userId")
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, false, err
	}

	refreshedAtValue := jsonToken.Get("refreshedAt")
	refreshedAt, err := time.Parse(refreshedAtLayout, refreshedAtValue)
	if err != nil {
		return 0, false, err
	}

	refresh := time.Now().UTC().After(refreshedAt)
	return userId, refresh, nil
}
