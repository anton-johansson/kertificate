// Copyright 2019 Anton Johansson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"kertificate.io/kertificate/pkg/api/static"
	v1 "kertificate.io/kertificate/pkg/api/v1"
	"kertificate.io/kertificate/pkg/auth"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ApiServer hosts the RESTful API
type ApiServer struct {
	api *echo.Echo
	v1  *v1.ApiV1
}

// NewApiServer initializes a new API server
func NewApiServer(v1 *v1.ApiV1) *ApiServer {
	api := echo.New()
	api.HideBanner = true
	api.HTTPErrorHandler = HandleError

	return &ApiServer{api, v1}
}

// Start starts the API server
func (server *ApiServer) Start() error {
	v1Group := server.api.Group("/v1", server.v1.Middlewares()...)

	server.api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{"X-Set-Authorization"},
	}))

	server.v1.Register(v1Group)
	server.api.GET("/*", static.Handler)
	return server.api.Start(":8080")
}

func HandleError(err error, context echo.Context) {
	if httpErr, ok := err.(*echo.HTTPError); ok {
		context.Response().WriteHeader(httpErr.Code)
	} else if err == sql.ErrNoRows {
		context.Response().WriteHeader(http.StatusNotFound)
	} else if err == auth.Unauthorized {
		context.Response().WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Println("Unhandled API error:", err.Error())
		context.String(http.StatusInternalServerError, err.Error())
	}
}
