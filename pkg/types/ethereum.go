package types

import (
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// EthereumSigningRequest contains all necessary data for Bitcoin protocol (BTC, BCH, LTC, DOGE, etc).
type EthereumSigningRequest struct {
	ChainId []byte `json:"chain_id,omitempty"`
	// Nonce (256-bit number)
	Nonce []byte `json:"nonce,omitempty"`
	// Gas price (256-bit number)
	// If > 0, legacy fee scheme is used; if 0, EIP1559 fee scheme is used
	GasPrice []byte `json:"gas_price,omitempty"`
	// Gas limit (256-bit number)
	GasLimit []byte `json:"gas_limit,omitempty"`
	// Maximum optional inclusion fee (aka tip) (256-bit number)
	// Used only for EIP1559 fee, disregarded for legacy
	MaxInclusionFeePerGas []byte `json:"max_inclusion_fee_per_gas,omitempty"`
	// Maximum fee (256-bit number)
	// Used only for EIP1559 fee, disregarded for legacy
	MaxFeePerGas []byte `json:"max_fee_per_gas,omitempty"`
	// Recipient's address.
	ToAddress string `json:"to_address,omitempty"`
	// ERC smart-contract address
	ContractAddress string `json:"contract_address,omitempty"`
	// Types that are assignable to Transaction:
	//	*transfer
	//	*erc20_transfer
	//	*erc20_approve
	//	*erc721_transfer
	//	*erc1155_transfer
	//	*contract
	TxType string `json:"tx_type,omitempty"`
	// Transfer
	// Amount to send in wei (256-bit number)
	Amount []byte `json:"amount,omitempty"`
	// Optional payload data
	Data []byte `json:"data,omitempty"`
	// ID of the token (256-bit number) - ERC721 NFT transfer transaction
	TokenId []byte `json:"token_id,omitempty"`
}

func (*EthereumSigningRequest) isSigningRequest_() {}

func (r *EthereumSigningRequest) UnmarshalJSON(data []byte) error {
	var rawMap map[string]string

	if err := json.Unmarshal(data, &rawMap); err != nil {
		return errors.Wrap(err, "unmarshal json to temporary container")
	}

	logrus.Debug("Unmarshal EthereumSigningRequest to rawMap map[string]string")
	logrus.Debug(rawMap)

	for _, field := range []string{
		"chain_id", "nonce",
		"gas_price", "gas_limit",
		"max_inclusion_fee_per_gas", "max_fee_per_gas",
		"amount", "data", "token_id",
	} {
		if valStr, exist := rawMap[field]; exist {
			logrus.Debugf("decode field [%s]", field)

			val, err := hex.DecodeString(valStr)
			if err != nil {
				return errors.Wrapf(err, "decode hex '%s", valStr)
			}

			logrus.Debug(val)

			switch field {
			case "chain_id":
				r.ChainId = val
			case "nonce":
				r.Nonce = val
			case "gas_price":
				r.GasPrice = val
			case "gas_limit":
				r.GasLimit = val
			case "max_inclusion_fee_per_gas":
				r.MaxInclusionFeePerGas = val
			case "max_fee_per_gas":
				r.MaxFeePerGas = val
			case "amount":
				r.Amount = val
			case "data":
				r.Data = val
			case "token_id":
				r.TokenId = val
			}
		}
	}

	for _, field := range []string{"to_address", "contract_address", "tx_type"} {
		if valStr, exist := rawMap[field]; exist {
			switch field {
			case "to_address":
				r.ToAddress = valStr
			case "contract_address":
				r.ContractAddress = valStr
			case "tx_type":
				r.TxType = valStr
			}
		}
	}

	return nil
}

func (r *EthereumSigningRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ChainId               string `json:"chain_id,omitempty"`
		Nonce                 string `json:"nonce,omitempty"`
		GasPrice              string `json:"gas_price,omitempty"`
		GasLimit              string `json:"gas_limit,omitempty"`
		MaxInclusionFeePerGas string `json:"max_inclusion_fee_per_gas,omitempty"`
		MaxFeePerGas          string `json:"max_fee_per_gas,omitempty"`
		ToAddress             string `json:"to_address,omitempty"`
		ContractAddress       string `json:"contract_address,omitempty"`
		TxType                string `json:"tx_type,omitempty"`
		Amount                string `json:"amount,omitempty"`
		Data                  string `json:"data,omitempty"`
		TokenId               string `json:"token_id,omitempty"`
	}{
		ChainId:               hex.EncodeToString(r.ChainId),
		Nonce:                 hex.EncodeToString(r.Nonce),
		GasPrice:              hex.EncodeToString(r.GasPrice),
		GasLimit:              hex.EncodeToString(r.GasLimit),
		MaxInclusionFeePerGas: hex.EncodeToString(r.MaxInclusionFeePerGas),
		MaxFeePerGas:          hex.EncodeToString(r.MaxFeePerGas),
		ToAddress:             r.ToAddress,
		ContractAddress:       r.ContractAddress,
		TxType:                r.TxType,
		Amount:                hex.EncodeToString(r.Amount),
		Data:                  hex.EncodeToString(r.Data),
		TokenId:               hex.EncodeToString(r.TokenId),
	})
}
