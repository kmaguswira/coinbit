// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: coinbit.proto

package proto

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

type DepositAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp int64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DepositAmount) Reset() {
	*x = DepositAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositAmount) ProtoMessage() {}

func (x *DepositAmount) ProtoReflect() protoreflect.Message {
	mi := &file_coinbit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositAmount.ProtoReflect.Descriptor instead.
func (*DepositAmount) Descriptor() ([]byte, []int) {
	return file_coinbit_proto_rawDescGZIP(), []int{0}
}

func (x *DepositAmount) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DepositAmount) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletID      string           `protobuf:"bytes,1,opt,name=walletID,proto3" json:"walletID,omitempty"`
	DepositAmount []*DepositAmount `protobuf:"bytes,2,rep,name=DepositAmount,proto3" json:"DepositAmount,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_coinbit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_coinbit_proto_rawDescGZIP(), []int{1}
}

func (x *Wallet) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *Wallet) GetDepositAmount() []*DepositAmount {
	if x != nil {
		return x.DepositAmount
	}
	return nil
}

type AboveThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletID       string `protobuf:"bytes,1,opt,name=walletID,proto3" json:"walletID,omitempty"`
	AboveThreshold bool   `protobuf:"varint,2,opt,name=aboveThreshold,proto3" json:"aboveThreshold,omitempty"`
}

func (x *AboveThreshold) Reset() {
	*x = AboveThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AboveThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AboveThreshold) ProtoMessage() {}

func (x *AboveThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_coinbit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AboveThreshold.ProtoReflect.Descriptor instead.
func (*AboveThreshold) Descriptor() ([]byte, []int) {
	return file_coinbit_proto_rawDescGZIP(), []int{2}
}

func (x *AboveThreshold) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *AboveThreshold) GetAboveThreshold() bool {
	if x != nil {
		return x.AboveThreshold
	}
	return false
}

var File_coinbit_proto protoreflect.FileDescriptor

var file_coinbit_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x60, 0x0a,
	0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x54, 0x0a, 0x0e, 0x41, 0x62, 0x6f, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x61, 0x62, 0x6f, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x62, 0x6f, 0x76, 0x65, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbit_proto_rawDescOnce sync.Once
	file_coinbit_proto_rawDescData = file_coinbit_proto_rawDesc
)

func file_coinbit_proto_rawDescGZIP() []byte {
	file_coinbit_proto_rawDescOnce.Do(func() {
		file_coinbit_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbit_proto_rawDescData)
	})
	return file_coinbit_proto_rawDescData
}

var file_coinbit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbit_proto_goTypes = []interface{}{
	(*DepositAmount)(nil),  // 0: proto.DepositAmount
	(*Wallet)(nil),         // 1: proto.Wallet
	(*AboveThreshold)(nil), // 2: proto.AboveThreshold
}
var file_coinbit_proto_depIdxs = []int32{
	0, // 0: proto.Wallet.DepositAmount:type_name -> proto.DepositAmount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coinbit_proto_init() }
func file_coinbit_proto_init() {
	if File_coinbit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositAmount); i {
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
		file_coinbit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
		file_coinbit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AboveThreshold); i {
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
			RawDescriptor: file_coinbit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbit_proto_goTypes,
		DependencyIndexes: file_coinbit_proto_depIdxs,
		MessageInfos:      file_coinbit_proto_msgTypes,
	}.Build()
	File_coinbit_proto = out.File
	file_coinbit_proto_rawDesc = nil
	file_coinbit_proto_goTypes = nil
	file_coinbit_proto_depIdxs = nil
}
