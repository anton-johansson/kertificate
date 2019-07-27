package api

import (
	"pkims/api/common"
	"pkims/auth"

	echo "github.com/labstack/echo/v4"
)

// ApiServer hosts the RESTful API
type ApiServer struct {
	api *echo.Echo
	V1  *echo.Group
}

// NewApiServer initializes a new API server
func NewApiServer(authService *auth.AuthService) *ApiServer {
	api := echo.New()
	api.HideBanner = true

	v1AuthMiddleware := common.NewAuthenticationMiddleware(authService, "/v1/authenticate", "/v1/status", "/v1/version")
	v1 := api.Group("/v1")
	v1.Use(v1AuthMiddleware.Process)
	return &ApiServer{
		api: api,
		V1:  v1,
	}
}

// Start starts the API server
func (server *ApiServer) Start() error {
	return server.api.Start(":8080")
}
