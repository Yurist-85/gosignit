package types

import "errors"

var (
	ErrBadRequestCoinRequired = errors.New("coin required")
	ErrBadRequestUnknownCoin  = errors.New("unknown coin")
	ErrBadRequestInvalidData  = errors.New("invalid data given")
	ErrInternalSigningFailed  = errors.New("signing failed")
	ErrNotImplemented         = errors.New("not implemented")
)
