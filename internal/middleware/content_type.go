package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// EnsureContentType validates request's Content-type header
func EnsureContentType(allowed []string) echo.MiddlewareFunc {
	allowedContentTypes := map[string]bool{}
	for _, contentType := range allowed {
		allowedContentTypes[strings.ToLower(contentType)] = true
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			if method == http.MethodPost || method == http.MethodPut {
				contentTypeHeader := c.Request().Header.Get(echo.HeaderContentType)
				contentType := strings.TrimSpace(strings.ToLower(strings.SplitN(contentTypeHeader, ";", 2)[0]))

				if !allowedContentTypes[contentType] {
					return c.JSON(http.StatusUnsupportedMediaType, map[string]string{})
				}
			}

			return next(c)
		}
	}
}
