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
}

func NewAuthService(userDAO *db.UserDAO) *AuthService {
	delegate := &dummy{}
	return &AuthService{
		userDAO,
		delegate,
	}
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
	return "s3cr3t-t0k3n", nil
}
