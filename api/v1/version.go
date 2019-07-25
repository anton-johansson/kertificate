package v1

import (
	"net/http"
	"pkims/version"

	echo "github.com/labstack/echo/v4"
)

type VersionAPI struct {
	GET func(echo.Context) error
}

func NewVersionAPI() *VersionAPI {
	return &VersionAPI{
		GET: func(context echo.Context) error {
			return context.JSON(http.StatusOK, version.Info())
		},
	}
}
