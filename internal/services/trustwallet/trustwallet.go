package trustwallet

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWAnyAddress.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWString.h>
import "C"
import (
	"github.com/yurist-85/gosignit/internal/services/trustwallet/coins"
	"unsafe"

	"github.com/sirupsen/logrus"

	"github.com/yurist-85/gosignit/internal/config"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/proto/bitcoin"
	twtypes "github.com/yurist-85/gosignit/internal/services/trustwallet/types"
	"github.com/yurist-85/gosignit/pkg/types"
)

type TrustWallet interface {
	SignTransaction(in types.SigningRequest, coin string) (interface{}, error)
	Wallet() unsafe.Pointer
}

type trustWallet struct {
	cfg    config.ConfigInterface
	wallet unsafe.Pointer
}

func NewTrustWallet(cfg config.ConfigInterface) TrustWallet {
	tw := &trustWallet{}
	tw.cfg = cfg

	seed := twtypes.TWStringCreateWithGoString(tw.cfg.Seed())
	empty := twtypes.TWStringCreateWithGoString("")
	wptr := C.TWHDWalletCreateWithMnemonic(seed, empty)
	tw.wallet = unsafe.Pointer(wptr)

	return tw
}

func (s *trustWallet) walletDelete() {
	wptr := (*C.struct_TWHDWallet)(s.wallet)
	C.TWHDWalletDelete(wptr)
}

func (s *trustWallet) Wallet() unsafe.Pointer {
	return s.wallet
}

func (s *trustWallet) SignTransaction(in types.SigningRequest, coin string) (interface{}, error) {
	logrus.WithFields(logrus.Fields{"req": in, "coin": coin}).Debug("sign tx")

	if coin == "" {
		return nil, twtypes.ErrCoinRequired
	}

	coinType, err := coins.CoinTypeByIdString(coin)
	if err != nil {
		return nil, twtypes.ErrUnknownCoin
	}

	blockchain, err := coins.BlockchainByIdString(coin)
	if err != nil {
		return nil, twtypes.ErrUnknownCoin
	}

	switch *blockchain {
	case C.TWCoinTypeBitcoin:
		if in, ok := in.(*types.BitcoinSigningRequest); ok {
			return bitcoin.SignTransaction(s.wallet, in, *coinType)
		}

		return nil, twtypes.ErrNonBitcoinSigningRequest
		//case C.TWBlockchainEthereum:
		//	if in, ok := in.(*types.EthereumSigningRequest); !ok {
		//		return s.signEthereum(in, *coinType, pk)
		//	}
		//
		//	return nil, twtypes.ErrNonEthereumSigningRequest
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
	}

	return nil, twtypes.ErrUnsupportedBlockchain
}
