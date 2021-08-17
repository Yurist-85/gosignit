package trustwallet

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestCoinSuite struct {
	suite.Suite
}

func TestCoin(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestCoinSuite))
}

func (t *TestCoinSuite) TestGetCoins_ok() {
	ctArr := GetCoins()
	t.Require().Greater(len(ctArr), 0)
}

func (t *TestCoinSuite) TestCoinTypeByIdString_ok() {
	ct, err := CoinTypeByIdString("ethereum")
	t.Require().NoError(err)
	t.Require().Equal(uint32(60), *ct)
}

func (t *TestCoinSuite) TestCoinTypeByIdString_fail() {
	ct, err := CoinTypeByIdString("unknown")
	t.Require().Error(err)
	t.Require().Nil(ct)
}

func (t *TestCoinSuite) TestCoinInfoByIdString_ok() {
	c, err := CoinInfoByIdString("ripple")
	t.Require().NoError(err)
	t.Require().Equal("ripple", c.ID)
	t.Equal("XRP", c.Name)
	t.Equal("XRP", c.Symbol)
	t.Equal(uint32(7), c.Blockchain)
	t.Equal(uint32(44), c.Purpose)
	t.Equal(uint32(0), c.Curve)
	t.Equal(uint32(0), c.XpubVersion)
	t.Equal(uint32(0), c.XprvVersion)
	t.Equal("m/44'/144'/0'/0/0", c.DerivationPath)
	t.Equal(uint32(0), c.PublicKeyType)
	t.Equal(uint8(0), c.StaticPrefix)
	t.Equal(uint8(0), c.P2pkhPrefix)
	t.Equal(uint8(0), c.P2shPrefix)
	t.Equal(uint32(0), c.Hrp)
	t.Equal(6, c.Decimals)
	t.Equal("https://bithomp.com/explorer/", c.ExplorerTransactionUrl)
	t.Equal("https://bithomp.com/explorer/", c.ExplorerAccountUrl)
	t.Equal(uint32(144), c.CoinType)
}

func (t *TestCoinSuite) TestCoinInfoByIdString_fail() {
	ct, err := CoinInfoByIdString("unknown")
	t.Require().Error(err)
	t.Require().Nil(ct)
}
