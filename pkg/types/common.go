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
	Gate   string         `json:"gate"`
	Tx     SigningRequest `json:"tx"`
	Signed interface{}    `json:"signed"`
}

func (r *CommonSigningResponse) GetSignedTransaction() interface{} {
	return r.Signed
}
