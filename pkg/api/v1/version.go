package v1

import (
	"net/http"

	"pkims.io/pkims/pkg/version"

	echo "github.com/labstack/echo/v4"
)

type VersionAPI struct {
}

func NewVersionAPI() *VersionAPI {
	return &VersionAPI{}
}

func (api *VersionAPI) Register(group *echo.Group) {
	group.GET("", getVersion)
}

func getVersion(context echo.Context) error {
	return context.JSON(http.StatusOK, version.Info())
}
