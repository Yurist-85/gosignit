// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: Zilliqa.proto

package zilliqa

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageOneof:
	//	*Transaction_Transfer_
	//	*Transaction_RawTransaction
	MessageOneof isTransaction_MessageOneof `protobuf_oneof:"message_oneof"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zilliqa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_Zilliqa_proto_msgTypes[0]
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
	return file_Zilliqa_proto_rawDescGZIP(), []int{0}
}

func (m *Transaction) GetMessageOneof() isTransaction_MessageOneof {
	if m != nil {
		return m.MessageOneof
	}
	return nil
}

func (x *Transaction) GetTransfer() *Transaction_Transfer {
	if x, ok := x.GetMessageOneof().(*Transaction_Transfer_); ok {
		return x.Transfer
	}
	return nil
}

func (x *Transaction) GetRawTransaction() *Transaction_Raw {
	if x, ok := x.GetMessageOneof().(*Transaction_RawTransaction); ok {
		return x.RawTransaction
	}
	return nil
}

type isTransaction_MessageOneof interface {
	isTransaction_MessageOneof()
}

type Transaction_Transfer_ struct {
	Transfer *Transaction_Transfer `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

type Transaction_RawTransaction struct {
	RawTransaction *Transaction_Raw `protobuf:"bytes,2,opt,name=raw_transaction,json=rawTransaction,proto3,oneof"`
}

func (*Transaction_Transfer_) isTransaction_MessageOneof() {}

func (*Transaction_RawTransaction) isTransaction_MessageOneof() {}

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction version
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Nonce
	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Recipient's address.
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	// GasPrice (256-bit number)
	GasPrice []byte `protobuf:"bytes,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// GasLimit
	GasLimit uint64 `protobuf:"varint,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Private Key
	PrivateKey  []byte       `protobuf:"bytes,6,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Transaction *Transaction `protobuf:"bytes,7,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zilliqa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Zilliqa_proto_msgTypes[1]
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
	return file_Zilliqa_proto_rawDescGZIP(), []int{1}
}

func (x *SigningInput) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SigningInput) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *SigningInput) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SigningInput) GetGasPrice() []byte {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

func (x *SigningInput) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// Transaction signing output.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signed signature bytes.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// JSON transaction with signature
	Json string `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zilliqa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Zilliqa_proto_msgTypes[2]
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
	return file_Zilliqa_proto_rawDescGZIP(), []int{2}
}

func (x *SigningOutput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SigningOutput) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type Transaction_Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send (256-bit number)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transaction_Transfer) Reset() {
	*x = Transaction_Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zilliqa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_Transfer) ProtoMessage() {}

func (x *Transaction_Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Zilliqa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_Transfer.ProtoReflect.Descriptor instead.
func (*Transaction_Transfer) Descriptor() ([]byte, []int) {
	return file_Zilliqa_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Transaction_Transfer) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

type Transaction_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Amount to send (256-bit number)
	Amount []byte `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Smart contract code
	Code []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// String-ified JSON object specifying the transition parameter
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Transaction_Raw) Reset() {
	*x = Transaction_Raw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zilliqa_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction_Raw) ProtoMessage() {}

func (x *Transaction_Raw) ProtoReflect() protoreflect.Message {
	mi := &file_Zilliqa_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction_Raw.ProtoReflect.Descriptor instead.
func (*Transaction_Raw) Descriptor() ([]byte, []int) {
	return file_Zilliqa_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Transaction_Raw) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Transaction_Raw) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Transaction_Raw) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Zilliqa_proto protoreflect.FileDescriptor

var file_Zilliqa_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x5a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x54, 0x57, 0x2e, 0x5a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x54, 0x57, 0x2e, 0x5a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x61, 0x77, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x54, 0x57, 0x2e, 0x5a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x61, 0x77, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x45, 0x0a, 0x03, 0x52, 0x61, 0x77,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x22, 0xea, 0x01, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x3f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54, 0x57, 0x2e, 0x5a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41,
	0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f,
	0x6e, 0x42, 0x3c, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x23, 0x61, 0x6c, 0x66, 0x61,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x7a, 0x69, 0x6c, 0x6c, 0x69, 0x71, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Zilliqa_proto_rawDescOnce sync.Once
	file_Zilliqa_proto_rawDescData = file_Zilliqa_proto_rawDesc
)

func file_Zilliqa_proto_rawDescGZIP() []byte {
	file_Zilliqa_proto_rawDescOnce.Do(func() {
		file_Zilliqa_proto_rawDescData = protoimpl.X.CompressGZIP(file_Zilliqa_proto_rawDescData)
	})
	return file_Zilliqa_proto_rawDescData
}

var file_Zilliqa_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Zilliqa_proto_goTypes = []interface{}{
	(*Transaction)(nil),          // 0: TW.Zilliqa.Proto.Transaction
	(*SigningInput)(nil),         // 1: TW.Zilliqa.Proto.SigningInput
	(*SigningOutput)(nil),        // 2: TW.Zilliqa.Proto.SigningOutput
	(*Transaction_Transfer)(nil), // 3: TW.Zilliqa.Proto.Transaction.Transfer
	(*Transaction_Raw)(nil),      // 4: TW.Zilliqa.Proto.Transaction.Raw
}
var file_Zilliqa_proto_depIdxs = []int32{
	3, // 0: TW.Zilliqa.Proto.Transaction.transfer:type_name -> TW.Zilliqa.Proto.Transaction.Transfer
	4, // 1: TW.Zilliqa.Proto.Transaction.raw_transaction:type_name -> TW.Zilliqa.Proto.Transaction.Raw
	0, // 2: TW.Zilliqa.Proto.SigningInput.transaction:type_name -> TW.Zilliqa.Proto.Transaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Zilliqa_proto_init() }
func file_Zilliqa_proto_init() {
	if File_Zilliqa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Zilliqa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Zilliqa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Zilliqa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Zilliqa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_Transfer); i {
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
		file_Zilliqa_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction_Raw); i {
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
	file_Zilliqa_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Transaction_Transfer_)(nil),
		(*Transaction_RawTransaction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Zilliqa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Zilliqa_proto_goTypes,
		DependencyIndexes: file_Zilliqa_proto_depIdxs,
		MessageInfos:      file_Zilliqa_proto_msgTypes,
	}.Build()
	File_Zilliqa_proto = out.File
	file_Zilliqa_proto_rawDesc = nil
	file_Zilliqa_proto_goTypes = nil
	file_Zilliqa_proto_depIdxs = nil
}
