// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: Decred.proto

package decred

import (
	proto "github.com/yurist-85/gosignit/pkg/proto"
	bitcoin "github.com/yurist-85/gosignit/pkg/proto/bitcoin"
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Serialization format
	SerializeType uint32 `protobuf:"varint,1,opt,name=serializeType,proto3" json:"serializeType,omitempty"`
	/// Transaction data format version
	Version uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// A list of 1 or more transaction inputs or sources for coins.
	Inputs []*TransactionInput `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// A list of 1 or more transaction outputs or destinations for coins
	Outputs []*TransactionOutput `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty"`
	/// The time when a transaction can be spent (usually zero, in which case it has no effect).
	LockTime uint32 `protobuf:"varint,5,opt,name=lockTime,proto3" json:"lockTime,omitempty"`
	/// The block height at which the transaction expires and is no longer valid.
	Expiry uint32 `protobuf:"varint,6,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decred_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_Decred_proto_msgTypes[0]
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
	return file_Decred_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetSerializeType() uint32 {
	if x != nil {
		return x.SerializeType
	}
	return 0
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*TransactionOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Transaction) GetLockTime() uint32 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *Transaction) GetExpiry() uint32 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

// Decred transaction input.
type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to the previous transaction's output.
	PreviousOutput *bitcoin.OutPoint `protobuf:"bytes,1,opt,name=previousOutput,proto3" json:"previousOutput,omitempty"`
	// Transaction version as defined by the sender.
	Sequence    uint32 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	ValueIn     int64  `protobuf:"varint,3,opt,name=valueIn,proto3" json:"valueIn,omitempty"`
	BlockHeight uint32 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BlockIndex  uint32 `protobuf:"varint,5,opt,name=blockIndex,proto3" json:"blockIndex,omitempty"`
	// Computational script for confirming transaction authorization.
	Script []byte `protobuf:"bytes,6,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decred_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_Decred_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInput.ProtoReflect.Descriptor instead.
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return file_Decred_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionInput) GetPreviousOutput() *bitcoin.OutPoint {
	if x != nil {
		return x.PreviousOutput
	}
	return nil
}

func (x *TransactionInput) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *TransactionInput) GetValueIn() int64 {
	if x != nil {
		return x.ValueIn
	}
	return 0
}

func (x *TransactionInput) GetBlockHeight() uint32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *TransactionInput) GetBlockIndex() uint32 {
	if x != nil {
		return x.BlockIndex
	}
	return 0
}

func (x *TransactionInput) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

// Decred transaction output.
type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction amount.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	/// Transaction output version.
	Version uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Usually contains the public key as a Decred script setting up conditions to claim this output.
	Script []byte `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decred_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Decred_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOutput.ProtoReflect.Descriptor instead.
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return file_Decred_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionOutput) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionOutput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TransactionOutput) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

// Transaction signing output.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resulting transaction. Note that the amount may be different than the requested amount to account for fees and available funds.
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,2,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// Transaction id
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Optional error
	Error proto.SigningError `protobuf:"varint,4,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Decred_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Decred_proto_msgTypes[3]
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
	return file_Decred_proto_rawDescGZIP(), []int{3}
}

func (x *SigningOutput) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *SigningOutput) GetError() proto.SigningError {
	if x != nil {
		return x.Error
	}
	return proto.SigningError(0)
}

var File_Decred_proto protoreflect.FileDescriptor

var file_Decred_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x44, 0x65, 0x63, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x54, 0x57, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x06,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x54,
	0x57, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x54, 0x57, 0x2e, 0x44, 0x65,
	0x63, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0xe6, 0x01, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x42,
	0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x54, 0x57, 0x2e, 0x42, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x22, 0x5b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22,
	0xc5, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x54, 0x57, 0x2e, 0x44, 0x65, 0x63, 0x72,
	0x65, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x3b, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x22, 0x61, 0x6c, 0x66, 0x61, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x63, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Decred_proto_rawDescOnce sync.Once
	file_Decred_proto_rawDescData = file_Decred_proto_rawDesc
)

func file_Decred_proto_rawDescGZIP() []byte {
	file_Decred_proto_rawDescOnce.Do(func() {
		file_Decred_proto_rawDescData = protoimpl.X.CompressGZIP(file_Decred_proto_rawDescData)
	})
	return file_Decred_proto_rawDescData
}

var file_Decred_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Decred_proto_goTypes = []interface{}{
	(*Transaction)(nil),       // 0: TW.Decred.Proto.Transaction
	(*TransactionInput)(nil),  // 1: TW.Decred.Proto.TransactionInput
	(*TransactionOutput)(nil), // 2: TW.Decred.Proto.TransactionOutput
	(*SigningOutput)(nil),     // 3: TW.Decred.Proto.SigningOutput
	(*bitcoin.OutPoint)(nil),  // 4: TW.Bitcoin.Proto.OutPoint
	(proto.SigningError)(0),   // 5: TW.Common.Proto.SigningError
}
var file_Decred_proto_depIdxs = []int32{
	1, // 0: TW.Decred.Proto.Transaction.inputs:type_name -> TW.Decred.Proto.TransactionInput
	2, // 1: TW.Decred.Proto.Transaction.outputs:type_name -> TW.Decred.Proto.TransactionOutput
	4, // 2: TW.Decred.Proto.TransactionInput.previousOutput:type_name -> TW.Bitcoin.Proto.OutPoint
	0, // 3: TW.Decred.Proto.SigningOutput.transaction:type_name -> TW.Decred.Proto.Transaction
	5, // 4: TW.Decred.Proto.SigningOutput.error:type_name -> TW.Common.Proto.SigningError
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Decred_proto_init() }
func file_Decred_proto_init() {
	if File_Decred_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Decred_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Decred_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInput); i {
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
		file_Decred_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOutput); i {
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
		file_Decred_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Decred_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Decred_proto_goTypes,
		DependencyIndexes: file_Decred_proto_depIdxs,
		MessageInfos:      file_Decred_proto_msgTypes,
	}.Build()
	File_Decred_proto = out.File
	file_Decred_proto_rawDesc = nil
	file_Decred_proto_goTypes = nil
	file_Decred_proto_depIdxs = nil
}
