package ethereum

import (
	"encoding/hex"
	"encoding/json"
)

func (x *SigningOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		// Signed and encoded transaction bytes.
		Encoded string `json:"encoded,omitempty"`
		V       string `json:"v,omitempty"`
		R       string `json:"r,omitempty"`
		S       string `json:"s,omitempty"`
		// The payload part, supplied in the input or assembled from input parameters
		Data string `json:"data,omitempty"`
	}{
		Encoded: hex.EncodeToString(x.Encoded),
		V:       hex.EncodeToString(x.V),
		R:       hex.EncodeToString(x.R),
		S:       hex.EncodeToString(x.S),
		// The payload part, supplied in the input or assembled from input parameters
		Data: hex.EncodeToString(x.Data),
	})
}
