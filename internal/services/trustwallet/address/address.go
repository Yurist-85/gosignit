package address

// #cgo CFLAGS: -I../../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../../lib/trustwallet/build -L../../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWAnyAddress.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWString.h>
import "C"
import (
	"unsafe"

	twtypes "github.com/yurist-85/gosignit/internal/services/trustwallet/types"
)

// GetAddressForCoin returns address for a given wallet and coin type.
func GetAddressForCoin(wallet unsafe.Pointer, coinType uint32) string {
	wptr := (*C.struct_TWHDWallet)(wallet)
	address := C.TWHDWalletGetAddressForCoin(wptr, coinType)
	defer C.TWStringDelete(address)

	return twtypes.TWStringGoString(address)
}

// ValidateAddress TODO Implement
func ValidateAddress(address string, coinType uint32) bool {
	cAddr := twtypes.TWStringCreateWithGoString(address)
	defer C.TWStringDelete(cAddr)

	res := C.TWAnyAddressIsValid(cAddr, coinType)

	return bool(res)
}
