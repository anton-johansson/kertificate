package v1

import (
	"fmt"
	"net/http"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

var ignoredPaths = [...]string{"/v1/authentication/authenticate", "/v1/status", "/v1/version"}

type AuthAPI struct {
	authService *auth.AuthService
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthAPI(authService *auth.AuthService) *AuthAPI {
	return &AuthAPI{authService}
}

func (api *AuthAPI) Register(group *echo.Group) {
	group.POST("/authenticate", api.authenticate)
}

func (api *AuthAPI) GetAuthMiddleware() echo.MiddlewareFunc {
	return api.checkRequest
}

func (api *AuthAPI) authenticate(context echo.Context) error {
	data := authRequest{}
	if err := context.Bind(&data); err != nil {
		return err
	}
	token, err := api.authService.Login(data.Username, data.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return nil
	}
	sendNewToken(context, token)
	context.Response().WriteHeader(http.StatusOK)
	return nil
}

func (api *AuthAPI) checkRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		requestPath := context.Request().URL.RequestURI()
		if func(path string) bool {
			for _, ignoredPath := range ignoredPaths {
				if ignoredPath == path {
					return true
				}
			}
			return false
		}(requestPath) {
			fmt.Println("auth mdlwr: ignoring path", requestPath)
			return next(context)
		}

		token := context.Request().Header.Get("Authorization")
		userId, newToken, err := api.authService.CheckToken(token)
		if err != nil {
			fmt.Println("auth mdlwr: denied")
			context.Error(err)
			return nil
		}
		context.Set("userId", userId)
		fmt.Println("auth mdlwr: got user", userId)
		if len(newToken) > 0 {
			fmt.Println("Refreshed token")
			sendNewToken(context, newToken)
		}
		return next(context)
	}
}

func sendNewToken(context echo.Context, token string) {
	context.Response().Header().Add("X-Set-Authorization", token)
}
