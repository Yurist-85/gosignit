package trustwallet

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWString.h>
// #include <TrustWalletCore/TWData.h>
import "C"
import (
	"encoding/hex"
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

// TWDataGoBytes converts TrustWallet data entity to Go bytes.
func TWDataGoBytes(d unsafe.Pointer) []byte {
	cBytes := C.TWDataBytes(d)
	cSize := C.TWDataSize(d)

	return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}

// TWDataCreateWithGoBytes creates TrustWallet data entity with a given Go byte[].
func TWDataCreateWithGoBytes(d []byte) unsafe.Pointer {
	cBytes := C.CBytes(d)
	defer C.free(unsafe.Pointer(cBytes))

	data := C.TWDataCreateWithBytes((*C.uchar)(cBytes), C.ulong(len(d)))

	return data
}

// TWDataHexString converts TrustWallet data entity to HEX string.
func TWDataHexString(d unsafe.Pointer) string {
	return hex.EncodeToString(TWDataGoBytes(d))
}
