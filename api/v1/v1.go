package v1

import (
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

// InitializeV1 initializes version 1 of the API
func InitializeV1(group *echo.Group, authenticationService *auth.AuthService) {
	versionAPI := NewVersionAPI()
	statusAPI := NewStatusAPI()
	authenticationAPI := NewAuthenticationAPI(authenticationService)

	group.GET("/version", versionAPI.GET)
	group.GET("/status", statusAPI.GET)
	group.POST("/authenticate", authenticationAPI.POST)
}
