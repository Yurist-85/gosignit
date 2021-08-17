package trustwallet

// #cgo CFLAGS: -I../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../lib/trustwallet/build -L../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWAnySigner.h>
// #include <TrustWalletCore/TWMnemonic.h>
import "C"
import "github.com/yurist-85/gosignit/pkg/models"

type SignerInterface interface {
	SignTransaction(in models.SignTransactionInput) (*models.SignTransactionOutput, error)
}

type Signer struct{}

func NewSigner() SignerInterface {
	tw := &Signer{}

	return tw
}

func (s *Signer) SignTransaction(in models.SignTransactionInput) (*models.SignTransactionOutput, error) {
	return nil, nil
}

func (s *Signer) GeneratePrivateKey() (string, error) {
	return "", nil
}

func (s *Signer) GeneratePublicKey() (string, error) {
	return "", nil
}

func (s *Signer) GenerateAddressFromKey() (string, error) {
	return "", nil
}

func (s *Signer) ValidateAddress(coin, address string) error {
	return nil
}
