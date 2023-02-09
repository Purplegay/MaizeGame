// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: RPCtest.proto

package myproto

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

type SumNumsREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nums int32 `protobuf:"varint,1,opt,name=nums,proto3" json:"nums,omitempty"`
}

func (x *SumNumsREQ) Reset() {
	*x = SumNumsREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPCtest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumNumsREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumNumsREQ) ProtoMessage() {}

func (x *SumNumsREQ) ProtoReflect() protoreflect.Message {
	mi := &file_RPCtest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumNumsREQ.ProtoReflect.Descriptor instead.
func (*SumNumsREQ) Descriptor() ([]byte, []int) {
	return file_RPCtest_proto_rawDescGZIP(), []int{0}
}

func (x *SumNumsREQ) GetNums() int32 {
	if x != nil {
		return x.Nums
	}
	return 0
}

type SumNumsACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumNumsACK) Reset() {
	*x = SumNumsACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPCtest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumNumsACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumNumsACK) ProtoMessage() {}

func (x *SumNumsACK) ProtoReflect() protoreflect.Message {
	mi := &file_RPCtest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumNumsACK.ProtoReflect.Descriptor instead.
func (*SumNumsACK) Descriptor() ([]byte, []int) {
	return file_RPCtest_proto_rawDescGZIP(), []int{1}
}

func (x *SumNumsACK) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_RPCtest_proto protoreflect.FileDescriptor

var file_RPCtest_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x52, 0x50, 0x43, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x4e,
	0x75, 0x6d, 0x73, 0x52, 0x45, 0x51, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x75,
	0x6d, 0x4e, 0x75, 0x6d, 0x73, 0x41, 0x43, 0x4b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x3d, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75,
	0x6d, 0x4e, 0x75, 0x6d, 0x73, 0x52, 0x45, 0x51, 0x1a, 0x13, 0x2e, 0x6d, 0x79, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x73, 0x41, 0x43, 0x4b, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RPCtest_proto_rawDescOnce sync.Once
	file_RPCtest_proto_rawDescData = file_RPCtest_proto_rawDesc
)

func file_RPCtest_proto_rawDescGZIP() []byte {
	file_RPCtest_proto_rawDescOnce.Do(func() {
		file_RPCtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_RPCtest_proto_rawDescData)
	})
	return file_RPCtest_proto_rawDescData
}

var file_RPCtest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RPCtest_proto_goTypes = []interface{}{
	(*SumNumsREQ)(nil), // 0: myproto.SumNumsREQ
	(*SumNumsACK)(nil), // 1: myproto.SumNumsACK
}
var file_RPCtest_proto_depIdxs = []int32{
	0, // 0: myproto.Sum.SayHello:input_type -> myproto.SumNumsREQ
	1, // 1: myproto.Sum.SayHello:output_type -> myproto.SumNumsACK
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RPCtest_proto_init() }
func file_RPCtest_proto_init() {
	if File_RPCtest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RPCtest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumNumsREQ); i {
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
		file_RPCtest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumNumsACK); i {
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
			RawDescriptor: file_RPCtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RPCtest_proto_goTypes,
		DependencyIndexes: file_RPCtest_proto_depIdxs,
		MessageInfos:      file_RPCtest_proto_msgTypes,
	}.Build()
	File_RPCtest_proto = out.File
	file_RPCtest_proto_rawDesc = nil
	file_RPCtest_proto_goTypes = nil
	file_RPCtest_proto_depIdxs = nil
}
