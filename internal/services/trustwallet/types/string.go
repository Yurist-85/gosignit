package types

// #cgo CFLAGS: -I../../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../../lib/trustwallet/build -L../../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWString.h>
import "C"
import (
	"unsafe"
)

// TWStringGoString converts TrustWallet string to Go string.
func TWStringGoString(s unsafe.Pointer) string {
	return C.GoString(C.TWStringUTF8Bytes(s))
}

// TWStringCreateWithGoString creates new TrustWallet string entity with a given Go string.
func TWStringCreateWithGoString(s string) unsafe.Pointer {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	str := C.TWStringCreateWithUTF8Bytes(cStr)

	return str
}
