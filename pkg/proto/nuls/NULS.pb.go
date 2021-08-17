// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: NULS.proto

package nuls

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionCoinFrom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress   string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	AssetsChainid uint32 `protobuf:"varint,2,opt,name=assets_chainid,json=assetsChainid,proto3" json:"assets_chainid,omitempty"`
	AssetsId      uint32 `protobuf:"varint,3,opt,name=assets_id,json=assetsId,proto3" json:"assets_id,omitempty"`
	//tranaction out amount (256-bit number)
	IdAmount []byte `protobuf:"bytes,4,opt,name=id_amount,json=idAmount,proto3" json:"id_amount,omitempty"`
	//8 bytes
	Nonce []byte `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	//lock status: 1 locked; 0 unlocked
	Locked uint32 `protobuf:"varint,6,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *TransactionCoinFrom) Reset() {
	*x = TransactionCoinFrom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCoinFrom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCoinFrom) ProtoMessage() {}

func (x *TransactionCoinFrom) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCoinFrom.ProtoReflect.Descriptor instead.
func (*TransactionCoinFrom) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionCoinFrom) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TransactionCoinFrom) GetAssetsChainid() uint32 {
	if x != nil {
		return x.AssetsChainid
	}
	return 0
}

func (x *TransactionCoinFrom) GetAssetsId() uint32 {
	if x != nil {
		return x.AssetsId
	}
	return 0
}

func (x *TransactionCoinFrom) GetIdAmount() []byte {
	if x != nil {
		return x.IdAmount
	}
	return nil
}

func (x *TransactionCoinFrom) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *TransactionCoinFrom) GetLocked() uint32 {
	if x != nil {
		return x.Locked
	}
	return 0
}

type TransactionCoinTo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToAddress     string `protobuf:"bytes,1,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetsChainid uint32 `protobuf:"varint,2,opt,name=assets_chainid,json=assetsChainid,proto3" json:"assets_chainid,omitempty"`
	AssetsId      uint32 `protobuf:"varint,3,opt,name=assets_id,json=assetsId,proto3" json:"assets_id,omitempty"`
	// tranaction amount (256-bit number)
	IdAmount []byte `protobuf:"bytes,4,opt,name=id_amount,json=idAmount,proto3" json:"id_amount,omitempty"`
	LockTime uint32 `protobuf:"varint,5,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
}

func (x *TransactionCoinTo) Reset() {
	*x = TransactionCoinTo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCoinTo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCoinTo) ProtoMessage() {}

func (x *TransactionCoinTo) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCoinTo.ProtoReflect.Descriptor instead.
func (*TransactionCoinTo) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionCoinTo) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionCoinTo) GetAssetsChainid() uint32 {
	if x != nil {
		return x.AssetsChainid
	}
	return 0
}

func (x *TransactionCoinTo) GetAssetsId() uint32 {
	if x != nil {
		return x.AssetsId
	}
	return 0
}

func (x *TransactionCoinTo) GetIdAmount() []byte {
	if x != nil {
		return x.IdAmount
	}
	return nil
}

func (x *TransactionCoinTo) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PkeyLen   uint32 `protobuf:"varint,1,opt,name=pkey_len,json=pkeyLen,proto3" json:"pkey_len,omitempty"`
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SigLen    uint32 `protobuf:"varint,3,opt,name=sig_len,json=sigLen,proto3" json:"sig_len,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetPkeyLen() uint32 {
	if x != nil {
		return x.PkeyLen
	}
	return 0
}

func (x *Signature) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Signature) GetSigLen() uint32 {
	if x != nil {
		return x.SigLen
	}
	return 0
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp uint32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Remark    string `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	TxData    []byte `protobuf:"bytes,4,opt,name=tx_data,json=txData,proto3" json:"tx_data,omitempty"`
	//CoinFrom
	Input *TransactionCoinFrom `protobuf:"bytes,5,opt,name=input,proto3" json:"input,omitempty"`
	//CoinTo
	Output *TransactionCoinTo `protobuf:"bytes,6,opt,name=output,proto3" json:"output,omitempty"`
	TxSigs *Signature         `protobuf:"bytes,7,opt,name=tx_sigs,json=txSigs,proto3" json:"tx_sigs,omitempty"`
	Hash   uint32             `protobuf:"varint,8,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Transaction) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Transaction) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Transaction) GetTxData() []byte {
	if x != nil {
		return x.TxData
	}
	return nil
}

func (x *Transaction) GetInput() *TransactionCoinFrom {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Transaction) GetOutput() *TransactionCoinTo {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *Transaction) GetTxSigs() *Signature {
	if x != nil {
		return x.TxSigs
	}
	return nil
}

func (x *Transaction) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

// Input data necessary to create a signed order.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey []byte `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	From       string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount     []byte `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	ChainId    uint32 `protobuf:"varint,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	IdassetsId uint32 `protobuf:"varint,6,opt,name=idassets_id,json=idassetsId,proto3" json:"idassets_id,omitempty"`
	//The last 8 bytes of latest transaction hash
	Nonce  []byte `protobuf:"bytes,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	// Account balance
	Balance []byte `protobuf:"bytes,9,opt,name=balance,proto3" json:"balance,omitempty"`
	// time, accurate to the second
	Timestamp uint32 `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningInput.ProtoReflect.Descriptor instead.
func (*SigningInput) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{4}
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SigningInput) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SigningInput) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *SigningInput) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *SigningInput) GetIdassetsId() uint32 {
	if x != nil {
		return x.IdassetsId
	}
	return 0
}

func (x *SigningInput) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *SigningInput) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SigningInput) GetBalance() []byte {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *SigningInput) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NULS_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_NULS_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningOutput.ProtoReflect.Descriptor instead.
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return file_NULS_proto_rawDescGZIP(), []int{5}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

var File_NULS_proto protoreflect.FileDescriptor

var file_NULS_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x4e, 0x55, 0x4c, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x54, 0x57,
	0x2e, 0x4e, 0x55, 0x4c, 0x53, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x13,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6b, 0x65, 0x79, 0x4c, 0x65, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x69, 0x67, 0x4c, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x54, 0x57, 0x2e, 0x4e, 0x55,
	0x4c, 0x53, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x54, 0x57, 0x2e, 0x4e, 0x55, 0x4c, 0x53, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x31, 0x0a,
	0x07, 0x74, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x54, 0x57, 0x2e, 0x4e, 0x55, 0x4c, 0x53, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x06, 0x74, 0x78, 0x53, 0x69, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x8d, 0x02, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x64, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x69, 0x64, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x42,
	0x39, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a,
	0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x20, 0x61, 0x6c, 0x66, 0x61, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x75, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_NULS_proto_rawDescOnce sync.Once
	file_NULS_proto_rawDescData = file_NULS_proto_rawDesc
)

func file_NULS_proto_rawDescGZIP() []byte {
	file_NULS_proto_rawDescOnce.Do(func() {
		file_NULS_proto_rawDescData = protoimpl.X.CompressGZIP(file_NULS_proto_rawDescData)
	})
	return file_NULS_proto_rawDescData
}

var file_NULS_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_NULS_proto_goTypes = []interface{}{
	(*TransactionCoinFrom)(nil), // 0: TW.NULS.Proto.TransactionCoinFrom
	(*TransactionCoinTo)(nil),   // 1: TW.NULS.Proto.TransactionCoinTo
	(*Signature)(nil),           // 2: TW.NULS.Proto.Signature
	(*Transaction)(nil),         // 3: TW.NULS.Proto.Transaction
	(*SigningInput)(nil),        // 4: TW.NULS.Proto.SigningInput
	(*SigningOutput)(nil),       // 5: TW.NULS.Proto.SigningOutput
}
var file_NULS_proto_depIdxs = []int32{
	0, // 0: TW.NULS.Proto.Transaction.input:type_name -> TW.NULS.Proto.TransactionCoinFrom
	1, // 1: TW.NULS.Proto.Transaction.output:type_name -> TW.NULS.Proto.TransactionCoinTo
	2, // 2: TW.NULS.Proto.Transaction.tx_sigs:type_name -> TW.NULS.Proto.Signature
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_NULS_proto_init() }
func file_NULS_proto_init() {
	if File_NULS_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NULS_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionCoinFrom); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_NULS_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionCoinTo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_NULS_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_NULS_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_NULS_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_NULS_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NULS_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NULS_proto_goTypes,
		DependencyIndexes: file_NULS_proto_depIdxs,
		MessageInfos:      file_NULS_proto_msgTypes,
	}.Build()
	File_NULS_proto = out.File
	file_NULS_proto_rawDesc = nil
	file_NULS_proto_goTypes = nil
	file_NULS_proto_depIdxs = nil
}
