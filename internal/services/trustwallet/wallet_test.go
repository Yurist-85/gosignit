package trustwallet

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/suite"
)

type TestWalletSuite struct {
	suite.Suite
	mnemonic string
}

type testCase struct {
	coin    string
	address string
}

func TestWallet(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestWalletSuite))
}

func (t *TestWalletSuite) SetupTest() {
	t.mnemonic = "observe drum fault concert analyst old short plunge loan essence symbol invite"
}

func (t *TestWalletSuite) TestMnemonicIsValid_fail() {
	valid := MnemonicIsValid("invalid mnemonic phrase")
	t.Require().False(valid)
}

func (t *TestWalletSuite) TestMnemonicIsValid_ok() {
	valid := MnemonicIsValid(t.mnemonic)
	t.Require().True(valid)
}

func (t *TestWalletSuite) TestWalletCreateWithMnemonic_ok() {
	wallet, err := WalletCreateWithMnemonic(t.mnemonic)
	t.Require().NoError(err)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)
}

func (t *TestWalletSuite) TestWalletCreateWithMnemonic_fail() {
	wallet, err := WalletCreateWithMnemonic("invalid mnemonic phrase")
	t.Require().Error(err)
	t.Require().Equal(unsafe.Pointer(nil), wallet)
}

func (t *TestWalletSuite) TestWalletDelete_ok() {
	wallet, err := WalletCreateWithMnemonic(t.mnemonic)
	t.Require().NoError(err)
	t.Require().NotNil(wallet)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)

	errDel := WalletDelete(wallet)
	t.Require().NoError(errDel)
}

func (t *TestWalletSuite) TestWalletGetKeyForCoin_ok() {
	wallet, err := WalletCreateWithMnemonic(t.mnemonic)
	t.Require().NoError(err)
	t.Require().NotNil(wallet)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)

	coin, err := CoinInfoByIdString("bitcoin")
	t.Require().NoError(err)
	t.Require().NotNil(coin)

	key, err := WalletGetKeyForCoin(wallet, unsafe.Pointer(&coin.CoinType))
	t.Require().NoError(err)
	t.Require().NotNil(key)
}

func (t *TestWalletSuite) TestWalletGetAddressForCoin_ok() {
	wallet, err := WalletCreateWithMnemonic(t.mnemonic)
	t.Require().NoError(err)
	t.Require().NotNil(wallet)
	t.Require().NotEqual(unsafe.Pointer(nil), wallet)

	cases := []testCase{
		{"bitcoin", "bc1q0tum480myw4kn6y94n5phvkpxqnducaqqmxp76"},
		{"ethereum", "0xd7B3500AB740d84F4A4e7aa36ecBAb343020A23B"},
	}

	for _, c := range cases {
		coin, err := CoinInfoByIdString(c.coin)
		t.Require().NoError(err)
		t.Require().NotNil(coin)

		addr, err := WalletGetAddressForCoin(wallet, unsafe.Pointer(&coin.CoinType))
		t.Require().NoError(err)
		t.Require().NotNil(addr)
		t.Require().Equal(c.address, addr)
	}
}
