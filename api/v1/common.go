package v1

import (
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func userId(context echo.Context) int {
	return context.Get("userId").(int)
}

func location(context echo.Context, identifier int) string {
	return context.Request().RequestURI + "/" + strconv.Itoa(identifier)
}
