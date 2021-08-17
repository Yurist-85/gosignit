// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: NEO.proto

package neo

import (
	proto "github.com/yurist-85/gosignit/pkg/proto"
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

type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevHash  []byte `protobuf:"bytes,1,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	PrevIndex uint32 `protobuf:"fixed32,2,opt,name=prev_index,json=prevIndex,proto3" json:"prev_index,omitempty"`
	// unspent value of UTXO
	Value   int64  `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	AssetId string `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[0]
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
	return file_NEO_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionInput) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *TransactionInput) GetPrevIndex() uint32 {
	if x != nil {
		return x.PrevIndex
	}
	return 0
}

func (x *TransactionInput) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionInput) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId       string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount        int64  `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ToAddress     string `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	ChangeAddress string `protobuf:"bytes,4,opt,name=change_address,json=changeAddress,proto3" json:"change_address,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[1]
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
	return file_NEO_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionOutput) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *TransactionOutput) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionOutput) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionOutput) GetChangeAddress() string {
	if x != nil {
		return x.ChangeAddress
	}
	return ""
}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs           []*TransactionInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs          []*TransactionOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	PrivateKey       []byte               `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Fee              int64                `protobuf:"varint,4,opt,name=fee,proto3" json:"fee,omitempty"`
	GasAssetId       string               `protobuf:"bytes,5,opt,name=gas_asset_id,json=gasAssetId,proto3" json:"gas_asset_id,omitempty"`
	GasChangeAddress string               `protobuf:"bytes,6,opt,name=gas_change_address,json=gasChangeAddress,proto3" json:"gas_change_address,omitempty"`
	Plan             *TransactionPlan     `protobuf:"bytes,7,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[2]
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
	return file_NEO_proto_rawDescGZIP(), []int{2}
}

func (x *SigningInput) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *SigningInput) GetOutputs() []*TransactionOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *SigningInput) GetGasAssetId() string {
	if x != nil {
		return x.GasAssetId
	}
	return ""
}

func (x *SigningInput) GetGasChangeAddress() string {
	if x != nil {
		return x.GasChangeAddress
	}
	return ""
}

func (x *SigningInput) GetPlan() *TransactionPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

// Transaction signing output.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
	// Optional error
	Error proto.SigningError `protobuf:"varint,2,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[3]
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
	return file_NEO_proto_rawDescGZIP(), []int{3}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

func (x *SigningOutput) GetError() proto.SigningError {
	if x != nil {
		return x.Error
	}
	return proto.SigningError(0)
}

// Describes a preliminary transaction output plan.
type TransactionOutputPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to be received at the other end.
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Maximum available amount.
	AvailableAmount int64  `protobuf:"varint,2,opt,name=available_amount,json=availableAmount,proto3" json:"available_amount,omitempty"`
	Change          int64  `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	AssetId         string `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	ToAddress       string `protobuf:"bytes,5,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	ChangeAddress   string `protobuf:"bytes,6,opt,name=change_address,json=changeAddress,proto3" json:"change_address,omitempty"`
}

func (x *TransactionOutputPlan) Reset() {
	*x = TransactionOutputPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutputPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutputPlan) ProtoMessage() {}

func (x *TransactionOutputPlan) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOutputPlan.ProtoReflect.Descriptor instead.
func (*TransactionOutputPlan) Descriptor() ([]byte, []int) {
	return file_NEO_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionOutputPlan) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionOutputPlan) GetAvailableAmount() int64 {
	if x != nil {
		return x.AvailableAmount
	}
	return 0
}

func (x *TransactionOutputPlan) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *TransactionOutputPlan) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *TransactionOutputPlan) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionOutputPlan) GetChangeAddress() string {
	if x != nil {
		return x.ChangeAddress
	}
	return ""
}

// Describes a preliminary transaction plan.
type TransactionPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used assets
	Outputs []*TransactionOutputPlan `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Selected unspent transaction outputs.
	Inputs []*TransactionInput `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// GAS used
	Fee int64 `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	// Optional error
	Error proto.SigningError `protobuf:"varint,4,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
}

func (x *TransactionPlan) Reset() {
	*x = TransactionPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEO_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionPlan) ProtoMessage() {}

func (x *TransactionPlan) ProtoReflect() protoreflect.Message {
	mi := &file_NEO_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionPlan.ProtoReflect.Descriptor instead.
func (*TransactionPlan) Descriptor() ([]byte, []int) {
	return file_NEO_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionPlan) GetOutputs() []*TransactionOutputPlan {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *TransactionPlan) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *TransactionPlan) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TransactionPlan) GetError() proto.SigningError {
	if x != nil {
		return x.Error
	}
	return proto.SigningError(0)
}

var File_NEO_proto protoreflect.FileDescriptor

var file_NEO_proto_rawDesc = []byte{
	0x0a, 0x09, 0x4e, 0x45, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x54, 0x57, 0x2e,
	0x4e, 0x45, 0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x76, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x70, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x76,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x76, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xb7, 0x02, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x54, 0x57, 0x2e, 0x4e, 0x45,
	0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x57, 0x2e, 0x4e, 0x45, 0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x67, 0x61, 0x73, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x73, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x67, 0x61, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x61,
	0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31,
	0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54,
	0x57, 0x2e, 0x4e, 0x45, 0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61,
	0x6e, 0x22, 0x5e, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x54, 0x57,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0xd3, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x54,
	0x57, 0x2e, 0x4e, 0x45, 0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x54, 0x57, 0x2e,
	0x4e, 0x45, 0x4f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x66, 0x65, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x38, 0x0a, 0x15, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x1f, 0x61, 0x6c, 0x66, 0x61, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NEO_proto_rawDescOnce sync.Once
	file_NEO_proto_rawDescData = file_NEO_proto_rawDesc
)

func file_NEO_proto_rawDescGZIP() []byte {
	file_NEO_proto_rawDescOnce.Do(func() {
		file_NEO_proto_rawDescData = protoimpl.X.CompressGZIP(file_NEO_proto_rawDescData)
	})
	return file_NEO_proto_rawDescData
}

var file_NEO_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_NEO_proto_goTypes = []interface{}{
	(*TransactionInput)(nil),      // 0: TW.NEO.Proto.TransactionInput
	(*TransactionOutput)(nil),     // 1: TW.NEO.Proto.TransactionOutput
	(*SigningInput)(nil),          // 2: TW.NEO.Proto.SigningInput
	(*SigningOutput)(nil),         // 3: TW.NEO.Proto.SigningOutput
	(*TransactionOutputPlan)(nil), // 4: TW.NEO.Proto.TransactionOutputPlan
	(*TransactionPlan)(nil),       // 5: TW.NEO.Proto.TransactionPlan
	(proto.SigningError)(0),       // 6: TW.Common.Proto.SigningError
}
var file_NEO_proto_depIdxs = []int32{
	0, // 0: TW.NEO.Proto.SigningInput.inputs:type_name -> TW.NEO.Proto.TransactionInput
	1, // 1: TW.NEO.Proto.SigningInput.outputs:type_name -> TW.NEO.Proto.TransactionOutput
	5, // 2: TW.NEO.Proto.SigningInput.plan:type_name -> TW.NEO.Proto.TransactionPlan
	6, // 3: TW.NEO.Proto.SigningOutput.error:type_name -> TW.Common.Proto.SigningError
	4, // 4: TW.NEO.Proto.TransactionPlan.outputs:type_name -> TW.NEO.Proto.TransactionOutputPlan
	0, // 5: TW.NEO.Proto.TransactionPlan.inputs:type_name -> TW.NEO.Proto.TransactionInput
	6, // 6: TW.NEO.Proto.TransactionPlan.error:type_name -> TW.Common.Proto.SigningError
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_NEO_proto_init() }
func file_NEO_proto_init() {
	if File_NEO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NEO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_NEO_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_NEO_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_NEO_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_NEO_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOutputPlan); i {
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
		file_NEO_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionPlan); i {
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
			RawDescriptor: file_NEO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NEO_proto_goTypes,
		DependencyIndexes: file_NEO_proto_depIdxs,
		MessageInfos:      file_NEO_proto_msgTypes,
	}.Build()
	File_NEO_proto = out.File
	file_NEO_proto_rawDesc = nil
	file_NEO_proto_goTypes = nil
	file_NEO_proto_depIdxs = nil
}
