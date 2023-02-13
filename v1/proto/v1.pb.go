// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/v1.proto

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

type Date int32

const (
	Date_UNSPECIFIED_VALUE Date = 0
	Date_DAY               Date = 1
	Date_MONTH             Date = 2
	Date_YEAR              Date = 3
)

// Enum value maps for Date.
var (
	Date_name = map[int32]string{
		0: "UNSPECIFIED_VALUE",
		1: "DAY",
		2: "MONTH",
		3: "YEAR",
	}
	Date_value = map[string]int32{
		"UNSPECIFIED_VALUE": 0,
		"DAY":               1,
		"MONTH":             2,
		"YEAR":              3,
	}
)

func (x Date) Enum() *Date {
	p := new(Date)
	*p = x
	return p
}

func (x Date) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Date) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1_proto_enumTypes[0].Descriptor()
}

func (Date) Type() protoreflect.EnumType {
	return &file_proto_v1_proto_enumTypes[0]
}

func (x Date) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Date.Descriptor instead.
func (Date) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{0}
}

type AboutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  uint32 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Date Date   `protobuf:"varint,4,opt,name=date,proto3,enum=Date" json:"date,omitempty"`
}

func (x *AboutRequest) Reset() {
	*x = AboutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AboutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AboutRequest) ProtoMessage() {}

func (x *AboutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AboutRequest.ProtoReflect.Descriptor instead.
func (*AboutRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{0}
}

func (x *AboutRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AboutRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AboutRequest) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *AboutRequest) GetDate() Date {
	if x != nil {
		return x.Date
	}
	return Date_UNSPECIFIED_VALUE
}

type AboutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *AboutResponse) Reset() {
	*x = AboutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AboutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AboutResponse) ProtoMessage() {}

func (x *AboutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AboutResponse.ProtoReflect.Descriptor instead.
func (*AboutResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AboutResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_proto_v1_proto protoreflect.FileDescriptor

var file_proto_v1_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x3b, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x59, 0x45, 0x41, 0x52, 0x10, 0x03, 0x32, 0x3c, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x75, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x41, 0x62, 0x6f, 0x75, 0x74,
	0x5f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x2f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_proto_rawDescOnce sync.Once
	file_proto_v1_proto_rawDescData = file_proto_v1_proto_rawDesc
)

func file_proto_v1_proto_rawDescGZIP() []byte {
	file_proto_v1_proto_rawDescOnce.Do(func() {
		file_proto_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_proto_rawDescData)
	})
	return file_proto_v1_proto_rawDescData
}

var file_proto_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_proto_goTypes = []interface{}{
	(Date)(0),             // 0: Date
	(*AboutRequest)(nil),  // 1: AboutRequest
	(*AboutResponse)(nil), // 2: AboutResponse
}
var file_proto_v1_proto_depIdxs = []int32{
	0, // 0: AboutRequest.date:type_name -> Date
	1, // 1: AboutService.About_Unary:input_type -> AboutRequest
	2, // 2: AboutService.About_Unary:output_type -> AboutResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1_proto_init() }
func file_proto_v1_proto_init() {
	if File_proto_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AboutRequest); i {
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
		file_proto_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AboutResponse); i {
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
			RawDescriptor: file_proto_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_proto_goTypes,
		DependencyIndexes: file_proto_v1_proto_depIdxs,
		EnumInfos:         file_proto_v1_proto_enumTypes,
		MessageInfos:      file_proto_v1_proto_msgTypes,
	}.Build()
	File_proto_v1_proto = out.File
	file_proto_v1_proto_rawDesc = nil
	file_proto_v1_proto_goTypes = nil
	file_proto_v1_proto_depIdxs = nil
}
