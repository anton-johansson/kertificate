package v1

import (
	"pkims/api/common"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

// InitializeV1 initializes version 1 of the API
func InitializeV1(group *echo.Group, authService *auth.AuthService) {
	authMiddleware := common.NewAuthenticationMiddleware(authService, "/v1/authenticate", "/v1/status", "/v1/version")
	versionAPI := NewVersionAPI()
	statusAPI := NewStatusAPI()
	authAPI := NewAuthAPI(authService)

	group.Use(authMiddleware.Process)
	group.GET("/version", versionAPI.GET)
	group.GET("/status", statusAPI.GET)
	group.POST("/authenticate", authAPI.POST)
}
