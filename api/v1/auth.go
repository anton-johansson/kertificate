package v1

import (
	"net/http"
	"pkims/api/common"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

type AuthAPI struct {
	POST func(echo.Context) error
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthAPI(authService *auth.AuthService) *AuthAPI {
	return &AuthAPI{
		POST: func(context echo.Context) error {
			data := authRequest{}
			if err := context.Bind(&data); err != nil {
				return err
			}
			token, err := authService.Login(data.Username, data.Password)
			if err != nil {
				context.JSON(http.StatusBadRequest, err)
				return nil
			}
			common.SendNewToken(context, token)
			context.Response().WriteHeader(http.StatusOK)
			return nil
		},
	}
}
