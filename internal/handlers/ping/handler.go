package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const Endpoint = "/ping"

func Handler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{})
}
