package middleware

import "github.com/labstack/echo/v4"

const (
	HeaderCacheControl = "Cache-Control"
	HeaderPragma       = "Pragma"
)

// NoCache sets Cache-Control: no-store and Pragma: no-cache headers for response.
func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(HeaderCacheControl, "no-store")
		c.Response().Header().Set(HeaderPragma, "no-cache")
		return next(c)
	}
}
