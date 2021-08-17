// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: Oasis.proto

package oasis

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

type TransferMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To       string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	GasPrice uint64 `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Amount values strings prefixed by zero. e.g. "\u000010000000"
	GasAmount string `protobuf:"bytes,3,opt,name=gas_amount,json=gasAmount,proto3" json:"gas_amount,omitempty"`
	// Amount values strings prefixed by zero
	Amount  string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Nonce   uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Context string `protobuf:"bytes,6,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *TransferMessage) Reset() {
	*x = TransferMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Oasis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferMessage) ProtoMessage() {}

func (x *TransferMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Oasis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferMessage.ProtoReflect.Descriptor instead.
func (*TransferMessage) Descriptor() ([]byte, []int) {
	return file_Oasis_proto_rawDescGZIP(), []int{0}
}

func (x *TransferMessage) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransferMessage) GetGasPrice() uint64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *TransferMessage) GetGasAmount() string {
	if x != nil {
		return x.GasAmount
	}
	return ""
}

func (x *TransferMessage) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TransferMessage) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *TransferMessage) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey []byte `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Types that are assignable to Message:
	//	*SigningInput_Transfer
	Message isSigningInput_Message `protobuf_oneof:"message"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Oasis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Oasis_proto_msgTypes[1]
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
	return file_Oasis_proto_rawDescGZIP(), []int{1}
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (m *SigningInput) GetMessage() isSigningInput_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *SigningInput) GetTransfer() *TransferMessage {
	if x, ok := x.GetMessage().(*SigningInput_Transfer); ok {
		return x.Transfer
	}
	return nil
}

type isSigningInput_Message interface {
	isSigningInput_Message()
}

type SigningInput_Transfer struct {
	Transfer *TransferMessage `protobuf:"bytes,2,opt,name=transfer,proto3,oneof"`
}

func (*SigningInput_Transfer) isSigningInput_Message() {}

// Transaction signing output.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed and encoded transaction bytes.
	Encoded []byte `protobuf:"bytes,1,opt,name=encoded,proto3" json:"encoded,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Oasis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Oasis_proto_msgTypes[2]
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
	return file_Oasis_proto_rawDescGZIP(), []int{2}
}

func (x *SigningOutput) GetEncoded() []byte {
	if x != nil {
		return x.Encoded
	}
	return nil
}

var File_Oasis_proto protoreflect.FileDescriptor

var file_Oasis_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x4f, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x54,
	0x57, 0x2e, 0x4f, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x61, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x79, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x54, 0x57, 0x2e, 0x4f, 0x61,
	0x73, 0x69, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x08, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x29, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x42, 0x3a, 0x0a, 0x15, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x21, 0x61, 0x6c, 0x66, 0x61, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x61, 0x73, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Oasis_proto_rawDescOnce sync.Once
	file_Oasis_proto_rawDescData = file_Oasis_proto_rawDesc
)

func file_Oasis_proto_rawDescGZIP() []byte {
	file_Oasis_proto_rawDescOnce.Do(func() {
		file_Oasis_proto_rawDescData = protoimpl.X.CompressGZIP(file_Oasis_proto_rawDescData)
	})
	return file_Oasis_proto_rawDescData
}

var file_Oasis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Oasis_proto_goTypes = []interface{}{
	(*TransferMessage)(nil), // 0: TW.Oasis.Proto.TransferMessage
	(*SigningInput)(nil),    // 1: TW.Oasis.Proto.SigningInput
	(*SigningOutput)(nil),   // 2: TW.Oasis.Proto.SigningOutput
}
var file_Oasis_proto_depIdxs = []int32{
	0, // 0: TW.Oasis.Proto.SigningInput.transfer:type_name -> TW.Oasis.Proto.TransferMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Oasis_proto_init() }
func file_Oasis_proto_init() {
	if File_Oasis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Oasis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferMessage); i {
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
		file_Oasis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Oasis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_Oasis_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SigningInput_Transfer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Oasis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Oasis_proto_goTypes,
		DependencyIndexes: file_Oasis_proto_depIdxs,
		MessageInfos:      file_Oasis_proto_msgTypes,
	}.Build()
	File_Oasis_proto = out.File
	file_Oasis_proto_rawDesc = nil
	file_Oasis_proto_goTypes = nil
	file_Oasis_proto_depIdxs = nil
}
