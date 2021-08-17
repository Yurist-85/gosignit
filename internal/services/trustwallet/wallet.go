package trustwallet

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWMnemonic.h>
// #include <TrustWalletCore/TWHDWallet.h>
import "C"
import "unsafe"

// MnemonicIsValid checks if given string is a valid mnemonic phrase.
func MnemonicIsValid(mnemonic string) bool {
	cValid := C.TWMnemonicIsValid(TWStringCreateWithGoString(mnemonic))

	return bool(cValid)
}

// WalletCreateWithMnemonic creates wallet with a given mnemonic phrase.
func WalletCreateWithMnemonic(mnemonic string) (unsafe.Pointer, error) {
	str := TWStringCreateWithGoString(mnemonic)
	empty := TWStringCreateWithGoString("")
	wallet, err := C.TWHDWalletCreateWithMnemonic(str, empty)

	return unsafe.Pointer(wallet), err
}

// WalletDelete deletes wallet from memory.
func WalletDelete(wallet unsafe.Pointer) error {
	cWallet := (*C.struct_TWHDWallet)(wallet)
	_, err := C.TWHDWalletDelete(cWallet)

	return err
}

// WalletGetKeyForCoin retrieves wallet key for a given coin
func WalletGetKeyForCoin(wallet, coin unsafe.Pointer) (unsafe.Pointer, error) {
	ct := *(*C.enum_TWCoinType)(coin)
	key := C.TWHDWalletGetKeyForCoin((*C.struct_TWHDWallet)(wallet), ct)

	return unsafe.Pointer(key), nil
}
