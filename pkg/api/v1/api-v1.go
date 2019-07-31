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
