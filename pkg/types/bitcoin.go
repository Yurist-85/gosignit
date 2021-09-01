package types

// BitcoinUnspentTransaction holds UTXO (unspent input).
type BitcoinUnspentTransaction struct {
	Hash     string `json:"hash,omitempty"`
	Index    uint32 `json:"index,omitempty"`
	Sequence uint32 `json:"sequence,omitempty"`
	Script   string `json:"script,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
}

// BitcoinTransactionPlan holds estimated transaction.
type BitcoinTransactionPlan struct {
	Amount          int64                        `json:"amount,omitempty"`
	AvailableAmount int64                        `json:"available_amount,omitempty"`
	Fee             int64                        `json:"fee,omitempty"`
	Change          int64                        `json:"change,omitempty"`
	Utxos           []*BitcoinUnspentTransaction `json:"utxos,omitempty"`
	// Zcash branch id
	BranchId string `json:"branch_id,omitempty"`
}

// BitcoinSigningRequest holds all necessary data for Bitcoin protocol (BTC, BCH, LTC, DOGE, etc).
type BitcoinSigningRequest struct {
	Amount        int64                        `json:"amount,omitempty"`
	ByteFee       int64                        `json:"byte_fee,omitempty"`
	ToAddress     string                       `json:"to_address,omitempty"`
	Scripts       map[string][]byte            `json:"scripts,omitempty"`
	Utxo          []*BitcoinUnspentTransaction `json:"utxo,omitempty"`
	ChangeAddress string                       `json:"change_address,omitempty"`
	UseMaxAmount  bool                         `json:"use_max_amount,omitempty"`
	Plan          *BitcoinTransactionPlan      `json:"plan,omitempty"`
}

func (*BitcoinSigningRequest) isSigningRequest_() {}
