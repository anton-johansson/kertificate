package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func NewTestAPI() func(echo.Context) error {
	return func(context echo.Context) error {
		return context.String(http.StatusLocked, "")
	}
}
