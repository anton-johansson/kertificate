package api

import (
	"database/sql"
	"fmt"
	"net/http"
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
	api.HTTPErrorHandler = HandleError

	return &ApiServer{
		api: api,
		V1:  api.Group("/v1"),
	}
}

// Start starts the API server
func (server *ApiServer) Start() error {
	return server.api.Start(":8080")
}

func HandleError(err error, context echo.Context) {
	if httpErr, ok := err.(*echo.HTTPError); ok {
		context.Response().WriteHeader(httpErr.Code)
	} else if err == sql.ErrNoRows {
		context.Response().WriteHeader(http.StatusNotFound)
	} else {
		fmt.Println("Unhandled API error:", err.Error())
		context.String(http.StatusInternalServerError, err.Error())
	}
}
