package txs

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWBlockchain.h>
// #include <TrustWalletCore/TWCoinType.h>
import "C"
import (
	"github.com/pkg/errors"
	"net/http"
	"strings"

	"github.com/jonboulle/clockwork"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di/v2"

	"github.com/yurist-85/gosignit/internal/services/clock"
	"github.com/yurist-85/gosignit/internal/services/trustwallet"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/coins"
	"github.com/yurist-85/gosignit/pkg/types"
)

const (
	EndpointRoot = "/tx"
	EndpointSign = "/:gate/sign"
)

type HandlerInterface interface {
	SetupRoutes(server *echo.Echo, prefix string)
	SignTransaction(c echo.Context) error
}

type Handler struct {
	clock clockwork.Clock
	tw trustwallet.TrustWallet
}

// NewHandler creates and returns new instance of Handler.
func NewHandler(ctn di.Container) HandlerInterface {
	h := &Handler{
		clock: ctn.Get(clock.DefinitionName).(clockwork.Clock),
		tw: ctn.Get(trustwallet.DefinitionName).(trustwallet.TrustWallet),
	}

	return h
}

// SetupRoutes registers its endpoints within given Echo server.
func (h *Handler) SetupRoutes(server *echo.Echo, prefix string) {
	server.POST(prefix+EndpointRoot+EndpointSign, h.SignTransaction)
}

// SignTransaction returns signed transaction. Depends on TrustWallet.
func (h *Handler) SignTransaction(c echo.Context) error {
	gate := strings.ToLower(strings.TrimSpace(c.Param("gate")))

	if gate == "" {
		// Should never happen
		return c.JSON(http.StatusBadRequest, types.ErrBadRequestCoinRequired)
	}

	blockchain, err := coins.BlockchainByIdString(gate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrBadRequestUnknownCoin)
	}

	var req types.SigningRequest

	switch *blockchain {
	case C.TWCoinTypeBitcoin: req = new(types.BitcoinSigningRequest)
	case C.TWBlockchainEthereum: req = new(types.EthereumSigningRequest)
	//case C.TWBlockchainVechain:
	//case C.TWBlockchainTron:
	//case C.TWBlockchainIcon:
	//case C.TWBlockchainBinance:
	//case C.TWBlockchainRipple:
	//case C.TWBlockchainTezos:
	//case C.TWBlockchainNimiq:
	//case C.TWBlockchainStellar:
	//case C.TWBlockchainAion:
	//case C.TWBlockchainCosmos:
	//case C.TWBlockchainTheta:
	//case C.TWBlockchainOntology:
	//case C.TWBlockchainZilliqa:
	//case C.TWBlockchainIoTeX:
	//case C.TWBlockchainEOS:
	//case C.TWBlockchainNano:
	//case C.TWBlockchainNULS:
	//case C.TWBlockchainWaves:
	//case C.TWBlockchainAeternity:
	//case C.TWBlockchainNebulas:
	//case C.TWBlockchainFIO:
	//case C.TWBlockchainSolana:
	//case C.TWBlockchainHarmony:
	//case C.TWBlockchainNEAR:
	//case C.TWBlockchainAlgorand:
	//case C.TWBlockchainPolkadot:
	//case C.TWBlockchainCardano:
	//case C.TWBlockchainNEO:
	//case C.TWBlockchainFilecoin:
	//case C.TWBlockchainElrondNetwork:
	//case C.TWBlockchainOasisNetwork:
	default:
		return c.JSON(http.StatusNotImplemented, types.ErrNotImplemented)
	}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "bind request data"))
	}

	signed, err := h.tw.SignTransaction(req, gate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "sign tx"))
	}

	resp := types.CommonSigningResponse{
		Gate:   gate,
		Tx:     req,
		Signed: signed,
	}

	return c.JSON(http.StatusOK, &resp)
}
