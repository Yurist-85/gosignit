package middleware

import (
	"bytes"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const LoggerName = "logger"

// RequestLogger proceeds Request's logs
func RequestLogger(l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := c.Response().Header().Get(echo.HeaderXRequestID)
			if id == "" {
				id = uuid.New().String()
			}

			logger := l.WithField("requestId", id)
			c.Set(LoggerName, logger)
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), LoggerName, logger)))

			return next(c)
		}
	}
}

// DefaultLogger defines default request's logger
func DefaultLogger() echo.MiddlewareFunc {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(new(RequestFormatter))

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:  logger.Writer(),
		Format:  middleware.DefaultLoggerConfig.Format,
		Skipper: DefaultSkipper,
	})
}

type RequestFormatter struct{}

// Format defines logs' format
func (f *RequestFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	if _, err := b.WriteString(entry.Message); err != nil {
		return b.Bytes(), err
	}

	if err := b.WriteByte('\n'); err != nil {
		return b.Bytes(), err
	}

	return b.Bytes(), nil
}
