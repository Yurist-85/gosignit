package middleware

import (
	"github.com/labstack/echo/v4"

	"github.com/yurist-85/gosignit/internal/handlers/ping"
)

func DefaultSkipper(ctx echo.Context) bool {
	// By default skip all requests to ping handler.
	return ctx.Request().RequestURI == ping.Endpoint
}
