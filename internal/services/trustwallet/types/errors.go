package types

import "errors"

var (
	ErrCoinRequired          = errors.New("coin required")
	ErrUnknownCoin           = errors.New("unknown or unsupported coin")
	ErrUnsupportedBlockchain = errors.New("unsupported blockchain")
	ErrGeneratePrivateKey    = errors.New("generate private key failed")
	ErrSigningInputRequired  = errors.New("signing input data required")
	ErrSigningErrorOccurred  = errors.New("error occurred during signing process")
)

var (
	ErrBitcoinInvalidUtxoHash = errors.New("invalid utxo hash")
	ErrBitcoinUtxoRequired    = errors.New("at least one utxo required")
)

var (
	ErrNonBitcoinSigningRequest  = errors.New("non-bitcoin input data")
	ErrNonEthereumSigningRequest = errors.New("non-ethereum input data")
)
