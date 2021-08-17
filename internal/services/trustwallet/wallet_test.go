package trustwallet

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/suite"
)

type TestWalletSuite struct {
	suite.Suite
}

func TestWallet(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestWalletSuite))
}

func (t *TestWalletSuite) TestMnemonicIsValid_fail() {
	valid := MnemonicIsValid("invalid mnemonic phrase")
	t.Require().False(valid)
}

func (t *TestWalletSuite) TestMnemonicIsValid_ok() {
	valid := MnemonicIsValid("observe drum fault concert analyst old short plunge loan essence symbol invite")
	t.Require().True(valid)
}

func (t *TestWalletSuite) TestWalletCreateWithMnemonic_ok() {
	wallet, err := WalletCreateWithMnemonic("observe drum fault concert analyst old short plunge loan essence symbol invite")
	t.Require().NoError(err)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)
}

func (t *TestWalletSuite) TestWalletCreateWithMnemonic_fail() {
	wallet, err := WalletCreateWithMnemonic("invalid mnemonic phrase")
	t.Require().Error(err)
	t.Require().Equal(unsafe.Pointer(nil), wallet)
}

func (t *TestWalletSuite) TestWalletDelete_ok() {
	wallet, err := WalletCreateWithMnemonic("observe drum fault concert analyst old short plunge loan essence symbol invite")
	t.Require().NoError(err)
	t.Require().NotNil(wallet)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)
	errDel := WalletDelete(wallet)
	t.Require().NoError(errDel)
}
