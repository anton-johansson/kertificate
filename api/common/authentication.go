package common

import (
	"fmt"
	"pkims/auth"

	"github.com/labstack/echo/v4"
)

type AuthenticationMiddleware struct {
	authService  *auth.AuthService
	ignoredPaths []string
}

func NewAuthenticationMiddleware(authService *auth.AuthService, ignoredPaths ...string) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{authService, ignoredPaths}
}

func (middleware *AuthenticationMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		requestPath := context.Request().URL.RequestURI()
		if func(paths []string, path string) bool {
			for _, ignoredPath := range paths {
				if ignoredPath == path {
					return true
				}
			}
			return false
		}(middleware.ignoredPaths, requestPath) {
			fmt.Println("auth mdlwr: ignoring path", requestPath)
			return next(context)
		}

		token := context.Request().Header.Get("Authorization")
		userId, newToken, err := middleware.authService.CheckToken(token)
		if err != nil {
			fmt.Println("auth mdlwr: denied")
			context.Error(err)
			return nil
		}
		context.Set("userId", userId)
		fmt.Println("auth mdlwr: got user", userId)
		if len(newToken) > 0 {
			fmt.Println("Refreshed token")
			SendNewToken(context, newToken)
		}
		return next(context)
	}
}

func SendNewToken(context echo.Context, token string) {
	context.Response().Header().Add("X-Set-Authorization", token)
}
