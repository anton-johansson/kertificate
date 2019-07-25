package api

import (
	echo "github.com/labstack/echo/v4"
)

// ApiServer hosts the RESTful API
type ApiServer struct {
	api *echo.Echo
	V1  *echo.Group
}

// NewApiServer initializes a new API server
func NewApiServer() *ApiServer {
	api := echo.New()
	api.HideBanner = true
	v1 := api.Group("/v1")
	return &ApiServer{
		api: api,
		V1:  v1,
	}
}

// Start starts the API server
func (server *ApiServer) Start() error {
	return server.api.Start(":8080")
}
