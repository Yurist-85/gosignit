package coins

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWBlockchain.h>
import "C"
import (
	"net/http"

	"github.com/jonboulle/clockwork"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"

	"github.com/yurist-85/gosignit/internal/services/clock"
	"github.com/yurist-85/gosignit/internal/services/trustwallet"
)

const Endpoint = "/coins"

type HandlerInterface interface {
	SetupRoutes(server *echo.Echo, prefix string)
	GetCoinsList(c echo.Context) error
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

// SetupRoutes registers endpoints within given Echo server.
func (h *Handler) SetupRoutes(server *echo.Echo, prefix string) {
	server.GET(prefix+Endpoint, h.GetCoinsList)
}

// GetCoinsList returns list of supported coins. Depends on TrustWallet.
func (h *Handler) GetCoinsList(c echo.Context) error {
	cList := trustwallet.GetCoins()
	var coins []*trustwallet.CoinInfo

	for i := 0; i < len(cList); i++ {
		coin, err := trustwallet.CoinInfoByIdString(cList[i])
		if err != nil {
			logrus.Errorf("unknown coin: %s", cList[i])
			continue
		}

		if coin.Blockchain != C.TWBlockchainBitcoin && coin.Blockchain != C.TWBlockchainEthereum {
			logrus.Debugf("not bitcoin or ethereum blockchain, skip: %v", coin.Name)
			continue
		}

		coins = append(coins, coin)
	}

	return c.JSON(http.StatusOK, &coins)
}
