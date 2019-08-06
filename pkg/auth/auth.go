package auth

import (
	"errors"
	"fmt"

	"pkims.io/pkims/pkg/db"
	"pkims.io/pkims/pkg/model"
)

type authDelegate interface {
	Login(username string, password string) *model.User
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
func (service *AuthService) Login(username string, password string) (model.User, string, error) {
	user := service.delegate.Login(username, password)
	if user == nil {
		service.userDAO.DeactivateIfExists(user.Username)
		return model.User{}, "", errors.New("Bad credentials")
	}
	userData, err := service.userDAO.GetOrCreateUser(*user)
	if err != nil {
		fmt.Println("error getting/creating user:", err)
		return model.User{}, "", err
	}
	fmt.Println("Got userId", userData.UserId)
	token, err := generateToken(userData.UserId, service.key)
	if err != nil {
		fmt.Println("error generating token:", err)
		return model.User{}, "", err
	}
	return userData, token, nil
}

// CheckToken checks an existing token by validating it and refreshing it if necessary
func (service *AuthService) CheckToken(token string) (int, string, error) {
	userId, expired := decryptToken(token, service.key)
	if userId <= 0 {
		return 0, "", Unauthorized
	}
	if !expired {
		return userId, "", nil
	}

	if !service.userDAO.IsActive(userId) {
		fmt.Println("User was no longer active")
		return 0, "", Unauthorized
	}

	newToken, err := generateToken(userId, service.key)
	if err != nil {
		fmt.Println("Error when refreshing token:", err)
		return 0, "", err
	}
	return userId, newToken, nil
}
