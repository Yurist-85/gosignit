package models

type CoinType string

type SignTransactionInput struct {
	Gate CoinType    `json:"gate"`
	Tx   interface{} `json:"tx"`
}

type SignTransactionOutput struct {
	Gate   CoinType    `json:"gate"`
	Tx     interface{} `json:"tx"`
	Signed interface{} `json:"signed"`
}
