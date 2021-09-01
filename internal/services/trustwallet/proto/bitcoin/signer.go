package bitcoin

// #cgo CFLAGS: -I../../../../../lib/trustwallet/include
// #cgo LDFLAGS: -L../../../../../lib/trustwallet/build -L../../../../../lib/trustwallet/build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWAnySigner.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWBitcoinSigHashType.h>
// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWString.h>
import "C"
import (
	"encoding/hex"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/address"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/keys"
	twtypes "github.com/yurist-85/gosignit/internal/services/trustwallet/types"
	"github.com/yurist-85/gosignit/pkg/types"
	"unsafe"
)

func SignTransaction(wallet unsafe.Pointer, req *types.BitcoinSigningRequest, coinType uint32) (*SigningOutput, error) {
	if req == nil {
		return nil, twtypes.ErrSigningInputRequired
	}

	addressStr := address.GetAddressForCoin(wallet, coinType)
	addr := twtypes.TWStringCreateWithGoString(addressStr)
	defer C.TWStringDelete(addr)

	script := C.TWBitcoinScriptLockScriptForAddress(addr, coinType)
	scriptData := C.TWBitcoinScriptData(script)
	defer C.TWBitcoinScriptDelete(script)
	defer C.TWDataDelete(scriptData)

	utxos, err := createUnspentTransactions(req.Utxo, scriptData)
	if err != nil {
		return nil, errors.Wrap(err, "create utxo")
	}

	pk := keys.GeneratePrivateKey(wallet, coinType)

	input := SigningInput{
		HashType:      C.TWBitcoinSigHashTypeAll,
		Amount:        req.Amount,
		ByteFee:       req.ByteFee,
		ToAddress:     req.ToAddress,
		ChangeAddress: req.ChangeAddress,
		PrivateKey:    [][]byte{pk},
		Utxo:          utxos,
		CoinType:      coinType,
	}

	// TODO Add Transaction Plan

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return nil, errors.Wrap(err, "proto.marshal signing input")
	}

	inputData := twtypes.TWDataCreateWithGoBytes(inputBytes)
	defer C.TWDataDelete(inputData)

	outputData := C.TWAnySignerSign(inputData, coinType)
	defer C.TWDataDelete(outputData)

	var output SigningOutput
	if err := proto.Unmarshal(twtypes.TWDataGoBytes(outputData), &output); err != nil {
		return nil, errors.Wrap(err, "proto.unmarshal signing output")
	}

	return &output, nil
}

//createUnspentTransactions creates list of UnspentTransaction's.
func createUnspentTransactions(txs []*types.BitcoinUnspentTransaction, scriptData unsafe.Pointer) ([]*UnspentTransaction, error) {
	var utxos []*UnspentTransaction

	if len(txs) < 1 {
		return nil, twtypes.ErrBitcoinUtxoRequired
	}

	for i, utx := range txs {
		utxoHash, _ := hex.DecodeString(utx.Hash)

		index := utx.Index
		if index == 0 {
			index = uint32(i)
		}

		utxo := &UnspentTransaction{
			OutPoint: &OutPoint{
				Hash:     utxoHash,
				Index:    index,
				Sequence: utx.Sequence,
			},
			Script: twtypes.TWDataGoBytes(scriptData),
			Amount: utx.Amount,
		}

		utxos = append(utxos, utxo)
	}

	return utxos, nil
}
