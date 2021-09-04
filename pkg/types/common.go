package types

// SigningRequest is common interface for all signing requests.
type SigningRequest interface {
	isSigningRequest_()
}

// SigningResponse is common interface for all signing responses.
type SigningResponse interface {
	GetSignedTransaction() interface{}
}

// CommonSigningResponse is common response struct. Contains original request and signed result.
type CommonSigningResponse struct {
	Gate   string         `json:"gate,omitempty"`
	Tx     SigningRequest `json:"tx,omitempty"`
	Signed interface{}    `json:"signed,omitempty"`
	Error  string         `json:"error,omitempty"`
}

func (r *CommonSigningResponse) GetSignedTransaction() interface{} {
	return r.Signed
}
