package bitcoin

import (
	"encoding/hex"
	"encoding/json"
	"github.com/yurist-85/gosignit/internal/services/trustwallet/proto"
)

// MarshalJSON temporary workaround to marshal bytes to string without encoding to base64.
func (x *SigningOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Transaction   *Transaction       `json:"transaction,omitempty"`
		Encoded       string             `json:"encoded,omitempty"`
		TransactionId string             `json:"transaction_id,omitempty"`
		Error         proto.SigningError `json:"error,omitempty"`
	}{
		Transaction:   x.Transaction,
		Encoded:       hex.EncodeToString(x.Encoded),
		TransactionId: x.TransactionId,
		Error:         x.Error,
	})
}

func (x *TransactionInput) MarshalJSON() ([]byte, error) {
	seq := x.Sequence
	if seq == 0 && x.PreviousOutput.Sequence != 0 {
		seq = x.PreviousOutput.Sequence
	}

	return json.Marshal(struct {
		Hash     string `json:"hash,omitempty"`
		Script   string `json:"script,omitempty"`
		Index    uint32 `json:"index,omitempty"`
		Sequence uint32 `json:"sequence,omitempty"`
	}{
		Hash:     hex.EncodeToString(x.PreviousOutput.Hash),
		Script:   hex.EncodeToString(x.Script),
		Index:    x.PreviousOutput.Index,
		Sequence: seq,
	})
}

func (x *TransactionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value  int64  `json:"value,omitempty"`
		Script string `json:"script,omitempty"`
	}{
		Value:  x.Value,
		Script: hex.EncodeToString(x.Script),
	})
}
