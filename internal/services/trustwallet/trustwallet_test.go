package trustwallet

import (
	"encoding/hex"
	"encoding/json"
	"github.com/sarulabs/di/v2"
	"github.com/stretchr/testify/suite"
	"github.com/yurist-85/gosignit/internal/config"
	"github.com/yurist-85/gosignit/internal/services/clock"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/keys"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/proto/bitcoin"
	"github.com/yurist-85/gosignit/pkg/types"
	"os"
	"testing"
)

type TestSignerSuite struct {
	suite.Suite
	mnemonic string
	cases    map[string]testSignerCase
	ctn      di.Container
	tw       TrustWallet
}

type testSignerCase struct {
	json string
}

func TestSigner(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestSignerSuite))
}

func (t *TestSignerSuite) SetupTest() {
	t.mnemonic = "observe drum fault concert analyst old short plunge loan essence symbol invite"
	t.Require().NoError(os.Setenv(config.EnvSeed, t.mnemonic))

	t.cases = map[string]testSignerCase{
		"empty":       {"{}"},
		"btc_valid":   {"{\"amount\":1020304,\"byte_fee\":1,\"to_address\":\"1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx\",\"change_address\":\"1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU\",\"use_max_amount\": false,\"utxo\":[{\"hash\":\"fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f\",\"index\":0,\"sequence\":4294967295,\"amount\":625000000}]}"},
		"btc_invalid": {"{\"amount\":1020304,\"byte_fee\":1,\"to_address\":\"1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx\",\"change_address\":\"1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU\",\"use_max_amount\":false}"},
	}

	builder, err := di.NewBuilder()
	t.Require().NoError(err)
	t.Require().NoError(builder.Add(
		clock.DefinitionTest,
		config.Definition,
		Definition,
	))

	t.ctn = builder.Build()
	t.tw = t.ctn.Get(DefinitionName).(TrustWallet)
}

func (t *TestSignerSuite) TearDownTest() {
	t.Require().NoError(t.ctn.DeleteWithSubContainers())
}

func (t *TestSignerSuite) TestSigner() {
	t.Run("TestGeneratePrivateKeyBitcoin", t.testGeneratePrivateKeyBitcoin)
	t.Run("TestGeneratePrivateKeyEthereum", t.testGeneratePrivateKeyEthereum)
	t.Run("TestSignBitcoin", t.testSignBitcoin)
}

func (t *TestSignerSuite) testGeneratePrivateKeyBitcoin() {
	coinType := uint32(0) // Bitcoin
	pk := keys.GeneratePrivateKey(t.tw.Wallet(), coinType)
	t.Require().NotEmpty(pk)

	str := hex.EncodeToString(pk)
	t.Require().Equal("f31bc13d989e3d5fbd87a871541f19cbc69a5f67fe14be259629943b911c52b0", str)
}

func (t *TestSignerSuite) testGeneratePrivateKeyEthereum() {
	coinType := uint32(60) // Ethereum
	pk := keys.GeneratePrivateKey(t.tw.Wallet(), coinType)
	t.Require().NotEmpty(pk)

	str := hex.EncodeToString(pk)
	t.Require().Equal("a9d7689f4f53227871f84865009503a27f5ba4fdec443751a906c55d341e3d12", str)
}

func (t *TestSignerSuite) testSignBitcoin() {
	input := t.getBitcoinSigningInput()
	t.Require().NotEmpty(input)

	coinType := uint32(0) // Bitcoin
	pk := keys.GeneratePrivateKey(t.tw.Wallet(), coinType)
	t.Require().NotEmpty(pk)

	signed, err := bitcoin.SignTransaction(t.tw.Wallet(), &input, coinType)
	t.Require().NoError(err)
	t.Require().NotNil(signed)

	//refSigned := bitcoin.SigningOutput{
	//	Transaction:   nil,
	//	Encoded:       nil,
	//	TransactionId: "",
	//	Error:         0,
	//}

	t.Require().NotNil(signed.Encoded)
	t.Require().NotNil(signed.Transaction)
	t.Require().NotNil(signed.TransactionId)
	t.Require().Empty(signed.Error)
}

func (t *TestSignerSuite) getBitcoinSigningInput() types.BitcoinSigningRequest {
	var req types.BitcoinSigningRequest

	tCase, exist := t.cases["btc_valid"]
	t.Require().True(exist)
	t.Require().NotEmpty(tCase)

	err := json.Unmarshal([]byte(tCase.json), &req)
	t.Require().NoError(err)
	t.Require().Equal("1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx", req.ToAddress)
	t.Require().Equal("1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU", req.ChangeAddress)
	t.Require().Equal(int64(1), req.ByteFee)
	t.Require().Equal(int64(1020304), req.Amount)
	t.Require().False(req.UseMaxAmount)
	t.Require().Equal(1, len(req.Utxo))
	t.Require().Equal(int64(625000000), req.Utxo[0].Amount)
	t.Require().NotEmpty(req.Utxo[0])
	t.Require().Equal("fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f", req.Utxo[0].Hash)
	t.Require().Equal(uint32(4294967295), req.Utxo[0].Sequence)

	return req
}
