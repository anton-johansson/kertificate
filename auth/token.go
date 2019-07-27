package auth

import (
	"strconv"
	"time"

	"github.com/o1egl/paseto"
)

var v2 = paseto.NewV2()

func generateToken(userId int64, key []byte) (string, error) {
	now := time.Now()
	exp := now.Add(time.Hour)

	jsonToken := paseto.JSONToken{
		Audience:   "test",
		Issuer:     "test_service",
		Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  now,
	}
	jsonToken.Set("userId", strconv.FormatInt(userId, 10))
	footer := ""

	return v2.Encrypt(key, jsonToken, footer)
}

func decryptToken(token string, key []byte) (int64, error) {
	var (
		jsonToken paseto.JSONToken
		footer    string
	)

	if err := v2.Decrypt(token, key, &jsonToken, &footer); err != nil {
		return 0, err
	}

	value := jsonToken.Get("userId")
	userId, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
