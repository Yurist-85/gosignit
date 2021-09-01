package types

// EthereumSigningRequest contains all necessary data for Bitcoin protocol (BTC, BCH, LTC, DOGE, etc).
type EthereumSigningRequest struct {
	Amount    int64  `json:"amount,omitempty"`
	ToAddress string `json:"to_address,omitempty"`
}

func (*EthereumSigningRequest) isSigningRequest_() {}
