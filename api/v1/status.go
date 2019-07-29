package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type StatusAPI struct {
}

func NewStatusAPI() *StatusAPI {
	return &StatusAPI{}
}

func (api *StatusAPI) Register(group *echo.Group) {
	group.GET("", getStatus)
}

func getStatus(context echo.Context) error {
	return context.String(http.StatusInternalServerError, "")
}
