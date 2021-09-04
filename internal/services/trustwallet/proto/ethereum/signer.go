package ethereum

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

func SignTransaction(wallet unsafe.Pointer, req *types.EthereumSigningRequest, coinType uint32) (*SigningOutput, error) {
	if req == nil {
		return nil, twtypes.ErrSigningInputRequired
	}

	if req.ChainId == nil {
		defChain, err := hex.DecodeString("01")
		if err != nil {
			return nil, errors.Wrap(err, "get default chain id")
		}

		req.ChainId = defChain
	}

	fromAddr := address.GetAddressForCoin(wallet, coinType)
	tx, err := createTransaction(req, fromAddr)
	if err != nil {
		return nil, errors.Wrap(err, "create transaction to sign")
	}

	input := SigningInput{
		ChainId:               req.ChainId,
		Nonce:                 req.Nonce,
		GasPrice:              req.GasPrice,
		GasLimit:              req.GasLimit,
		MaxInclusionFeePerGas: req.MaxInclusionFeePerGas,
		MaxFeePerGas:          req.MaxFeePerGas,
		ToAddress:             req.ToAddress,
		PrivateKey:            keys.GeneratePrivateKey(wallet, coinType),
		Transaction:           tx,
	}

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

//createTransaction creates Transaction to sign.
func createTransaction(req *types.EthereumSigningRequest, fromAddr string) (*Transaction, error) {
	tx := &Transaction{}

	switch req.TxType {
	case "transfer":
		tx.TransactionOneof = &Transaction_Transfer_{
			Transfer: &Transaction_Transfer{
				Amount: req.Amount,
				Data:   req.Data,
			},
		}
	case "erc20_transfer":
		tx.TransactionOneof = &Transaction_Erc20Transfer{
			Erc20Transfer: &Transaction_ERC20Transfer{
				Amount: req.Amount,
				To:     req.ToAddress,
			},
		}
	case "erc20_approve":
		tx.TransactionOneof = &Transaction_Erc20Approve{
			Erc20Approve: &Transaction_ERC20Approve{
				Spender: req.ToAddress,
				Amount:  req.Amount,
			},
		}
	case "erc721_transfer":
		tx.TransactionOneof = &Transaction_Erc721Transfer{
			Erc721Transfer: &Transaction_ERC721Transfer{
				From:    fromAddr,
				To:      req.ToAddress,
				TokenId: req.TokenId,
			},
		}
	case "erc1155_transfer":
		tx.TransactionOneof = &Transaction_Erc1155Transfer{
			Erc1155Transfer: &Transaction_ERC1155Transfer{
				From:    fromAddr,
				To:      req.ToAddress,
				TokenId: req.TokenId,
				Value:   req.Amount,
				Data:    req.Data,
			},
		}
	case "contract":
		tx.TransactionOneof = &Transaction_ContractGeneric_{
			ContractGeneric: &Transaction_ContractGeneric{
				Amount: req.Amount,
				Data:   req.Data,
			},
		}
	default:
		return nil, errors.New("invalid tx_type")
	}

	return tx, nil
}
