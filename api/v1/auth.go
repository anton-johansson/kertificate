package v1

import (
	"net/http"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

type AuthenticationAPI struct {
	POST func(echo.Context) error
}

type authenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthenticationAPI(authenticationService *auth.AuthService) *AuthenticationAPI {
	return &AuthenticationAPI{
		POST: func(context echo.Context) error {
			data := authenticationRequest{}
			if err := context.Bind(&data); err != nil {
				return err
			}
			token, err := authenticationService.Login(data.Username, data.Password)
			if err != nil {
				context.JSON(http.StatusBadRequest, err)
				return nil
			}
			return context.JSON(http.StatusOK, token)
		},
	}
}
