package auth

import (
	"errors"
	"fmt"
	"pkims/db"
)

type authDelegate interface {
	Login(username string, password string) *authResult
}

type authResult struct {
	username string
	active   bool
}

type AuthService struct {
	userDAO  *db.UserDAO
	delegate authDelegate
	key      []byte
}

func NewAuthService(userDAO *db.UserDAO) *AuthService {
	delegate := &dummy{}
	key := []byte("f4b92153521a438795eb853454242bba")
	return &AuthService{userDAO, delegate, key}
}

// Login attempts to log a user in. It returns either a login token back, or an error
func (service *AuthService) Login(username string, password string) (string, error) {
	result := service.delegate.Login(username, password)
	if result == nil {
		service.userDAO.DeactivateIfExists(username)
		return "", errors.New("Bad credentials")
	}
	userId := service.userDAO.GetOrCreateId(result.username)
	fmt.Println("Got userId", userId)
	token, err := generateToken(userId, service.key)
	if err != nil {
		fmt.Println(err)
	}
	return token, err
}

// CheckToken checks an existing token by validating it and refreshing it if necessary
func (service *AuthService) CheckToken(token string) (int64, error) {
	userId, err := decryptToken(token, service.key)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return userId, nil
}
