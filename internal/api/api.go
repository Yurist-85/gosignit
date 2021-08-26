package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di/v2"

	"github.com/yurist-85/gosignit/internal/handlers/coins"
	"github.com/yurist-85/gosignit/internal/handlers/txs"
)

const EnvAllowOrigin = "API_ALLOW_ORIGIN"

type HandlerInterface interface {
	GetEcho() *echo.Echo
}

type Handler struct {
	server *echo.Echo
	coins  coins.HandlerInterface
	txs    txs.HandlerInterface
}

// NewHandler creates new API handler with routes and middlewares.
func NewHandler(ctn di.Container) HandlerInterface {
	api := &Handler{
		server: ctn.Get(DefinitionNameEcho).(*echo.Echo),
		coins:  ctn.Get(coins.DefinitionName).(coins.HandlerInterface),
		txs:    ctn.Get(txs.DefinitionName).(txs.HandlerInterface),
	}

	// CORS middleware
	api.server.Use(middleware.CORSWithConfig(api.getCORSConfig()))
	// Setup routes
	prefix := "/api/v1"
	api.coins.SetupRoutes(api.server, prefix)
	api.txs.SetupRoutes(api.server, prefix)
	// TODO Swagger route
	//api.server.GET(PathPublicSwagger, echoutils.SwaggerHandler(echoutils.SwaggerConfig{DocsPath: "/docs"}))

	return api
}

// getCORSConfig returns CORS configuration.
func (h *Handler) getCORSConfig() (cors middleware.CORSConfig) {
	cors = middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: h.getOriginsFromEnv(),
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	}

	return cors
}

// getOriginsFromEnv returns Origins settings.
func (h *Handler) getOriginsFromEnv() (origins []string) {
	originsStr, ok := os.LookupEnv(EnvAllowOrigin)
	if ok && len(originsStr) > 5 {
		origins = strings.Split(originsStr, ",")
	}

	if len(origins) == 0 {
		origins = []string{
			"http://localhost:3000",
		}
	}

	return
}

func (h *Handler) GetEcho() *echo.Echo {
	return h.server
}
