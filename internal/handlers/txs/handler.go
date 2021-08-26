package txs

import (
	"net/http"

	"github.com/jonboulle/clockwork"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"

	"github.com/yurist-85/gosignit/internal/services/clock"
)

const (
	EndpointRoot = "/tx"
	EndpointSign = "/tx/sign"
)

type HandlerInterface interface {
	SetupRoutes(server *echo.Echo, prefix string)
	SignTransaction(c echo.Context) error
}

type Handler struct {
	clock clockwork.Clock
}

// NewHandler creates and returns new instance of Handler.
func NewHandler(ctn di.Container) HandlerInterface {
	h := &Handler{
		clock: ctn.Get(clock.DefinitionName).(clockwork.Clock),
	}

	return h
}

// SetupRoutes registers its endpoints within given Echo server.
func (h *Handler) SetupRoutes(server *echo.Echo, prefix string) {
	server.POST(prefix+EndpointSign, h.SignTransaction)
}

// SignTransaction returns signed transaction. Depends on TrustWallet.
func (h *Handler) SignTransaction(c echo.Context) error {
	// TODO Return signed transaction
	tx := map[string]interface{}{}

	return c.JSON(http.StatusOK, &tx)
}
