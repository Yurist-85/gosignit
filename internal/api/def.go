package api

import (
	"github.com/labstack/echo/v4"
	echomiddle "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"

	"github.com/yurist-85/gosignit/internal/handlers/ping"
	"github.com/yurist-85/gosignit/internal/middleware"
)

const (
	DefinitionName     = "api-handler"
	DefinitionNameEcho = "api-echo"
)

var (
	Definition = di.Def{
		Name: DefinitionName,
		Build: func(ctn di.Container) (i interface{}, err error) {
			handler := NewHandler(ctn)

			return handler, nil
		},
	}
	DefinitionEcho = di.Def{
		Name:  DefinitionNameEcho,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			e := echo.New()
			e.DisableHTTP2 = true
			e.HidePort = true
			e.HideBanner = true
			e.Logger.SetLevel(log.WARN)
			// TODO: Default error handler required
			// e.HTTPErrorHandler = handler.DefaultErrorsHandler
			// Define middlewares
			e.Pre(echomiddle.RemoveTrailingSlash())
			e.Use(middleware.EnsureContentType([]string{echo.MIMEApplicationJSON}))
			e.Use(echomiddle.Recover())
			e.Use(echomiddle.BodyLimitWithConfig(echomiddle.BodyLimitConfig{Limit: "1M"}))
			e.Use(echomiddle.RequestIDWithConfig(echomiddle.RequestIDConfig{}))
			e.Use(echomiddle.Gzip())
			e.Use(middleware.RequestLogger(logrus.StandardLogger()))
			e.Use(middleware.DefaultLogger())
			// Default ping-pong endpoint
			e.GET(ping.Endpoint, ping.Handler, middleware.NoCache)

			return e, nil
		},
		Close: func(obj interface{}) error {
			return obj.(*echo.Echo).Close()
		},
	}
)
