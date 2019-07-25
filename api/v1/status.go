package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type StatusAPI struct {
	GET func(echo.Context) error
}

func NewStatusAPI() *StatusAPI {
	return &StatusAPI{
		GET: func(context echo.Context) error {
			return context.String(http.StatusInternalServerError, "")
		},
	}
}
