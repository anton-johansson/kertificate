package v1

import (
	"pkims/api/common"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

// InitializeV1 initializes version 1 of the API
func InitializeV1(group *echo.Group, authService *auth.AuthService, consumerTypeAPI *ConsumerTypeAPI) {
	authMiddleware := common.NewAuthenticationMiddleware(authService, "/v1/authenticate", "/v1/status", "/v1/version")
	versionAPI := NewVersionAPI()
	statusAPI := NewStatusAPI()
	authAPI := NewAuthAPI(authService)
	testAPI := NewTestAPI()

	group.Use(authMiddleware.Process)
	group.GET("/version", versionAPI.GET)
	group.GET("/status", statusAPI.GET)
	group.POST("/authenticate", authAPI.POST)
	group.GET("/test", testAPI)
	group.POST("/consumer-types", consumerTypeAPI.Create)
	group.GET("/consumer-types", consumerTypeAPI.List)
	group.GET("/consumer-types/:typeId", consumerTypeAPI.Get)
	group.PUT("/consumer-types/:typeId", consumerTypeAPI.Update)
	group.DELETE("/consumer-types/:typeId", consumerTypeAPI.Delete)
}
