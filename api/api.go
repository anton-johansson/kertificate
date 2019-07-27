package api

import (
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

	return &ApiServer{
		api: api,
		V1:  api.Group("/v1"),
	}
}

// Start starts the API server
func (server *ApiServer) Start() error {
	return server.api.Start(":8080")
}
