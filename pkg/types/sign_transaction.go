package types

import (
	"github.com/yurist-85/gosignit/pkg/proto/bitcoin"
)

// SigningRequest is common interface for all signing requests.
type SigningRequest interface {
	isSigningRequest_()
}

// SigningResponse is common interface for all signing responses.
type SigningResponse interface {
	GetSignedTransaction() interface{}
}

// BitcoinSigningRequest contains all necessary data for Bitcoin protocol (BTC, BCH, LTC, DOGE, etc).
type BitcoinSigningRequest struct {
	Amount        int64                         `json:"amount,omitempty"`
	ByteFee       int64                         `json:"byte_fee,omitempty"`
	ToAddress     string                        `json:"to_address,omitempty"`
	Scripts       map[string][]byte             `json:"scripts,omitempty"`
	Utxo          []*bitcoin.UnspentTransaction `json:"utxo,omitempty"`
	ChangeAddress string                        `json:"change_address,omitempty"`
	UseMaxAmount  bool                          `json:"use_max_amount,omitempty"`
	Plan          *bitcoin.TransactionPlan      `json:"plan,omitempty"`
}

func (*BitcoinSigningRequest) isSigningRequest_() {}

// EthereumSigningRequest contains all necessary data for Bitcoin protocol (BTC, BCH, LTC, DOGE, etc).
type EthereumSigningRequest struct {
	Amount    int64  `json:"amount,omitempty"`
	ToAddress string `json:"to_address,omitempty"`
}

func (*EthereumSigningRequest) isSigningRequest_() {}

// CommonSigningResponse is common response struct. Contains original request and signed result.
type CommonSigningResponse struct {
	Gate   string         `json:"gate"`
	Tx     SigningRequest `json:"tx"`
	Signed interface{}    `json:"signed"`
}

func (r *CommonSigningResponse) GetSignedTransaction() interface{} {
	return r.Signed
}
