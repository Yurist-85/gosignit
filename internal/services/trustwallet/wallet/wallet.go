package wallet

// #cgo CFLAGS: -I../../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../../lib/trustwallet/build -L../../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWMnemonic.h>
// #include <TrustWalletCore/TWHDWallet.h>
import "C"
import (
	"unsafe"

	"github.com/yurist-85/gosignit/internal/services/trustwallet/types"
)

// MnemonicIsValid checks if given string is a valid mnemonic phrase.
func MnemonicIsValid(mnemonic string) bool {
	cValid := C.TWMnemonicIsValid(types.TWStringCreateWithGoString(mnemonic))

	return bool(cValid)
}

// WalletCreateWithMnemonic creates wallet with a given mnemonic phrase.
func WalletCreateWithMnemonic(mnemonic string) (unsafe.Pointer, error) {
	str := types.TWStringCreateWithGoString(mnemonic)
	empty := types.TWStringCreateWithGoString("")
	wallet, err := C.TWHDWalletCreateWithMnemonic(str, empty)

	return unsafe.Pointer(wallet), err
}

// WalletDelete deletes wallet from memory.
func WalletDelete(wallet unsafe.Pointer) error {
	cWallet := (*C.struct_TWHDWallet)(wallet)
	_, err := C.TWHDWalletDelete(cWallet)

	return err
}

// WalletGetKeyForCoin retrieves wallet key for a given coin.
func WalletGetKeyForCoin(wallet, coin unsafe.Pointer) (unsafe.Pointer, error) {
	ct := *(*C.enum_TWCoinType)(coin)
	key := C.TWHDWalletGetKeyForCoin((*C.struct_TWHDWallet)(wallet), ct)

	return unsafe.Pointer(key), nil
}

// WalletGetAddressForCoin generates address for a given wallet and coin.
func WalletGetAddressForCoin(wallet, coin unsafe.Pointer) (string, error) {
	ct := *(*C.enum_TWCoinType)(coin)
	addr := C.TWHDWalletGetAddressForCoin((*C.struct_TWHDWallet)(wallet), ct)

	return types.TWStringGoString(addr), nil
}
