// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: Tezos.proto

package tezos

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

type Operation_OperationKind int32

const (
	// Note: Proto3 semantics require a zero value.
	Operation_ENDORSEMENT Operation_OperationKind = 0
	Operation_REVEAL      Operation_OperationKind = 107
	Operation_TRANSACTION Operation_OperationKind = 108
	Operation_DELEGATION  Operation_OperationKind = 110
)

// Enum value maps for Operation_OperationKind.
var (
	Operation_OperationKind_name = map[int32]string{
		0:   "ENDORSEMENT",
		107: "REVEAL",
		108: "TRANSACTION",
		110: "DELEGATION",
	}
	Operation_OperationKind_value = map[string]int32{
		"ENDORSEMENT": 0,
		"REVEAL":      107,
		"TRANSACTION": 108,
		"DELEGATION":  110,
	}
)

func (x Operation_OperationKind) Enum() *Operation_OperationKind {
	p := new(Operation_OperationKind)
	*p = x
	return p
}

func (x Operation_OperationKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_OperationKind) Descriptor() protoreflect.EnumDescriptor {
	return file_Tezos_proto_enumTypes[0].Descriptor()
}

func (Operation_OperationKind) Type() protoreflect.EnumType {
	return &file_Tezos_proto_enumTypes[0]
}

func (x Operation_OperationKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_OperationKind.Descriptor instead.
func (Operation_OperationKind) EnumDescriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{3, 0}
}

// Input data necessary to create a signed Tezos transaction.
// Next field: 3
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationList *OperationList `protobuf:"bytes,1,opt,name=operation_list,json=operationList,proto3" json:"operation_list,omitempty"`
	PrivateKey    []byte         `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[0]
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
	return file_Tezos_proto_rawDescGZIP(), []int{0}
}

func (x *SigningInput) GetOperationList() *OperationList {
	if x != nil {
		return x.OperationList
	}
	return nil
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

// Transaction signing output.
// Next field: 2
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[1]
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
	return file_Tezos_proto_rawDescGZIP(), []int{1}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

// A list of operations and a branch.
// Next field: 3
type OperationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch     string       `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	Operations []*Operation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *OperationList) Reset() {
	*x = OperationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationList) ProtoMessage() {}

func (x *OperationList) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationList.ProtoReflect.Descriptor instead.
func (*OperationList) Descriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{2}
}

func (x *OperationList) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *OperationList) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

// An operation that can be applied to the Tezos blockchain.
// Next field: 12
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter      int64                   `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
	Source       string                  `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Fee          int64                   `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	GasLimit     int64                   `protobuf:"varint,4,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	StorageLimit int64                   `protobuf:"varint,5,opt,name=storage_limit,json=storageLimit,proto3" json:"storage_limit,omitempty"`
	Kind         Operation_OperationKind `protobuf:"varint,7,opt,name=kind,proto3,enum=TW.Tezos.Proto.Operation_OperationKind" json:"kind,omitempty"`
	// Operation specific data depending on the type of the operation.
	//
	// Types that are assignable to OperationData:
	//	*Operation_RevealOperationData
	//	*Operation_TransactionOperationData
	//	*Operation_DelegationOperationData
	OperationData isOperation_OperationData `protobuf_oneof:"operation_data"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{3}
}

func (x *Operation) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *Operation) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Operation) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Operation) GetGasLimit() int64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Operation) GetStorageLimit() int64 {
	if x != nil {
		return x.StorageLimit
	}
	return 0
}

func (x *Operation) GetKind() Operation_OperationKind {
	if x != nil {
		return x.Kind
	}
	return Operation_ENDORSEMENT
}

func (m *Operation) GetOperationData() isOperation_OperationData {
	if m != nil {
		return m.OperationData
	}
	return nil
}

func (x *Operation) GetRevealOperationData() *RevealOperationData {
	if x, ok := x.GetOperationData().(*Operation_RevealOperationData); ok {
		return x.RevealOperationData
	}
	return nil
}

func (x *Operation) GetTransactionOperationData() *TransactionOperationData {
	if x, ok := x.GetOperationData().(*Operation_TransactionOperationData); ok {
		return x.TransactionOperationData
	}
	return nil
}

func (x *Operation) GetDelegationOperationData() *DelegationOperationData {
	if x, ok := x.GetOperationData().(*Operation_DelegationOperationData); ok {
		return x.DelegationOperationData
	}
	return nil
}

type isOperation_OperationData interface {
	isOperation_OperationData()
}

type Operation_RevealOperationData struct {
	RevealOperationData *RevealOperationData `protobuf:"bytes,8,opt,name=reveal_operation_data,json=revealOperationData,proto3,oneof"`
}

type Operation_TransactionOperationData struct {
	TransactionOperationData *TransactionOperationData `protobuf:"bytes,9,opt,name=transaction_operation_data,json=transactionOperationData,proto3,oneof"`
}

type Operation_DelegationOperationData struct {
	DelegationOperationData *DelegationOperationData `protobuf:"bytes,11,opt,name=delegation_operation_data,json=delegationOperationData,proto3,oneof"`
}

func (*Operation_RevealOperationData) isOperation_OperationData() {}

func (*Operation_TransactionOperationData) isOperation_OperationData() {}

func (*Operation_DelegationOperationData) isOperation_OperationData() {}

// Transaction operation specific data.
// Next field: 3
type TransactionOperationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination string `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	Amount      int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionOperationData) Reset() {
	*x = TransactionOperationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOperationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOperationData) ProtoMessage() {}

func (x *TransactionOperationData) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOperationData.ProtoReflect.Descriptor instead.
func (*TransactionOperationData) Descriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionOperationData) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *TransactionOperationData) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Reveal operation specific data.
// Next field: 2
type RevealOperationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *RevealOperationData) Reset() {
	*x = RevealOperationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevealOperationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevealOperationData) ProtoMessage() {}

func (x *RevealOperationData) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevealOperationData.ProtoReflect.Descriptor instead.
func (*RevealOperationData) Descriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{5}
}

func (x *RevealOperationData) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Delegation operation specific data.
// Next field: 2
type DelegationOperationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delegate string `protobuf:"bytes,1,opt,name=delegate,proto3" json:"delegate,omitempty"`
}

func (x *DelegationOperationData) Reset() {
	*x = DelegationOperationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tezos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegationOperationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegationOperationData) ProtoMessage() {}

func (x *DelegationOperationData) ProtoReflect() protoreflect.Message {
	mi := &file_Tezos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelegationOperationData.ProtoReflect.Descriptor instead.
func (*DelegationOperationData) Descriptor() ([]byte, []int) {
	return file_Tezos_proto_rawDescGZIP(), []int{6}
}

func (x *DelegationOperationData) GetDelegate() string {
	if x != nil {
		return x.Delegate
	}
	return ""
}

var File_Tezos_proto protoreflect.FileDescriptor

var file_Tezos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x54,
	0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a,
	0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x44, 0x0a,
	0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x22,
	0x62, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x54,
	0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0xdb, 0x04, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x54, 0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x59, 0x0a, 0x15, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x54, 0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x13, 0x72, 0x65, 0x76, 0x65,
	0x61, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x68, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x54, 0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x65, 0x0a, 0x19, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x54,
	0x57, 0x2e, 0x54, 0x65, 0x7a, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x4d, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x53, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x56, 0x45, 0x41, 0x4c, 0x10, 0x6b, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x6c, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x4c, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x6e, 0x42,
	0x10, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x54, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x13, 0x52, 0x65, 0x76, 0x65, 0x61,
	0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x35, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x42, 0x3a, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x21, 0x61,
	0x6c, 0x66, 0x61, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x7a, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Tezos_proto_rawDescOnce sync.Once
	file_Tezos_proto_rawDescData = file_Tezos_proto_rawDesc
)

func file_Tezos_proto_rawDescGZIP() []byte {
	file_Tezos_proto_rawDescOnce.Do(func() {
		file_Tezos_proto_rawDescData = protoimpl.X.CompressGZIP(file_Tezos_proto_rawDescData)
	})
	return file_Tezos_proto_rawDescData
}

var file_Tezos_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Tezos_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Tezos_proto_goTypes = []interface{}{
	(Operation_OperationKind)(0),     // 0: TW.Tezos.Proto.Operation.OperationKind
	(*SigningInput)(nil),             // 1: TW.Tezos.Proto.SigningInput
	(*SigningOutput)(nil),            // 2: TW.Tezos.Proto.SigningOutput
	(*OperationList)(nil),            // 3: TW.Tezos.Proto.OperationList
	(*Operation)(nil),                // 4: TW.Tezos.Proto.Operation
	(*TransactionOperationData)(nil), // 5: TW.Tezos.Proto.TransactionOperationData
	(*RevealOperationData)(nil),      // 6: TW.Tezos.Proto.RevealOperationData
	(*DelegationOperationData)(nil),  // 7: TW.Tezos.Proto.DelegationOperationData
}
var file_Tezos_proto_depIdxs = []int32{
	3, // 0: TW.Tezos.Proto.SigningInput.operation_list:type_name -> TW.Tezos.Proto.OperationList
	4, // 1: TW.Tezos.Proto.OperationList.operations:type_name -> TW.Tezos.Proto.Operation
	0, // 2: TW.Tezos.Proto.Operation.kind:type_name -> TW.Tezos.Proto.Operation.OperationKind
	6, // 3: TW.Tezos.Proto.Operation.reveal_operation_data:type_name -> TW.Tezos.Proto.RevealOperationData
	5, // 4: TW.Tezos.Proto.Operation.transaction_operation_data:type_name -> TW.Tezos.Proto.TransactionOperationData
	7, // 5: TW.Tezos.Proto.Operation.delegation_operation_data:type_name -> TW.Tezos.Proto.DelegationOperationData
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Tezos_proto_init() }
func file_Tezos_proto_init() {
	if File_Tezos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Tezos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Tezos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Tezos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationList); i {
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
		file_Tezos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_Tezos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOperationData); i {
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
		file_Tezos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevealOperationData); i {
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
		file_Tezos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegationOperationData); i {
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
	file_Tezos_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Operation_RevealOperationData)(nil),
		(*Operation_TransactionOperationData)(nil),
		(*Operation_DelegationOperationData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Tezos_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Tezos_proto_goTypes,
		DependencyIndexes: file_Tezos_proto_depIdxs,
		EnumInfos:         file_Tezos_proto_enumTypes,
		MessageInfos:      file_Tezos_proto_msgTypes,
	}.Build()
	File_Tezos_proto = out.File
	file_Tezos_proto_rawDesc = nil
	file_Tezos_proto_goTypes = nil
	file_Tezos_proto_depIdxs = nil
}