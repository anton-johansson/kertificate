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

package v1

import (
	echo "github.com/labstack/echo/v4"
)

// ApiV1 holds routes for version 1 of the API
type ApiV1 struct {
	authAPI                *AuthAPI
	certificateTemplateAPI *CertificateTemplateAPI
	certificateAPI         *CertificateAPI
	commonAuthorityAPI     *CommonAuthorityAPI
	consumerTypeAPI        *ConsumerTypeAPI
	statusAPI              *StatusAPI
	versionAPI             *VersionAPI
}

func NewApiV1(
	authAPI *AuthAPI,
	certificateTemplateAPI *CertificateTemplateAPI,
	certificateAPI *CertificateAPI,
	commonAuthorityAPI *CommonAuthorityAPI,
	consumerTypeAPI *ConsumerTypeAPI,
	statusAPI *StatusAPI,
	versionAPI *VersionAPI) *ApiV1 {
	return &ApiV1{
		authAPI,
		certificateTemplateAPI,
		certificateAPI,
		commonAuthorityAPI,
		consumerTypeAPI,
		statusAPI,
		versionAPI,
	}
}

// Register registers the V1 API methods
func (api *ApiV1) Register(group *echo.Group) {
	api.authAPI.Register(group.Group("/authentication"))
	api.certificateTemplateAPI.Register(group.Group("/certificate-templates"))
	api.certificateAPI.Register(group.Group("/certificates"))
	api.commonAuthorityAPI.Register(group.Group("/common-authorities"))
	api.consumerTypeAPI.Register(group.Group("/consumer-types"))
	api.statusAPI.Register(group.Group("/status"))
	api.versionAPI.Register(group.Group("/version"))
}

// Middlewares returns all the middleware functions that should be used globally for the V1 API
func (api *ApiV1) Middlewares() []echo.MiddlewareFunc {
	middlewares := make([]echo.MiddlewareFunc, 0)
	middlewares = append(middlewares, api.authAPI.GetAuthMiddleware())
	return middlewares
}
