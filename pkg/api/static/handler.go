package static

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Handler = (func() echo.HandlerFunc {
	if assets != nil {
		return echo.WrapHandler(http.FileServer(assets))
	}
	return echo.NotFoundHandler
})()
