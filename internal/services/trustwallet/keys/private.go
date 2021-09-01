package keys

// #cgo CFLAGS: -I../../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../../lib/trustwallet/build -L../../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWData.h>
import "C"
import (
	"unsafe"

	twtypes "github.com/yurist-85/gosignit/internal/services/trustwallet/types"
)

// GeneratePrivateKey returns private key for a given wallet and coin type.
func GeneratePrivateKey(wallet unsafe.Pointer, coinType uint32) []byte {
	w := (*C.struct_TWHDWallet)(wallet)
	key := C.TWHDWalletGetKeyForCoin(w, coinType)
	keyData := C.TWPrivateKeyData(key)
	defer C.TWDataDelete(keyData)

	return twtypes.TWDataGoBytes(keyData)
}
