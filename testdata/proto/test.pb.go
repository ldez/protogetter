// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: test.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Optional enum
type Test_OEnum int32

const (
	Test_O_ENUM1 Test_OEnum = 0
	Test_O_ENUM2 Test_OEnum = 1
)

// Enum value maps for Test_OEnum.
var (
	Test_OEnum_name = map[int32]string{
		0: "O_ENUM1",
		1: "O_ENUM2",
	}
	Test_OEnum_value = map[string]int32{
		"O_ENUM1": 0,
		"O_ENUM2": 1,
	}
)

func (x Test_OEnum) Enum() *Test_OEnum {
	p := new(Test_OEnum)
	*p = x
	return p
}

func (x Test_OEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test_OEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (Test_OEnum) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Test_OEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test_OEnum.Descriptor instead.
func (Test_OEnum) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0}
}

type Test struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	D                 float64                `protobuf:"fixed64,1,opt,name=d,proto3" json:"d,omitempty"`
	F                 float32                `protobuf:"fixed32,2,opt,name=f,proto3" json:"f,omitempty"`
	I32               int32                  `protobuf:"varint,3,opt,name=i32,proto3" json:"i32,omitempty"`
	I64               int64                  `protobuf:"varint,4,opt,name=i64,proto3" json:"i64,omitempty"`
	U32               uint32                 `protobuf:"varint,5,opt,name=u32,proto3" json:"u32,omitempty"`
	U64               uint64                 `protobuf:"varint,6,opt,name=u64,proto3" json:"u64,omitempty"`
	T                 bool                   `protobuf:"varint,7,opt,name=t,proto3" json:"t,omitempty"`
	B                 []byte                 `protobuf:"bytes,8,opt,name=b,proto3" json:"b,omitempty"`
	S                 string                 `protobuf:"bytes,9,opt,name=s,proto3" json:"s,omitempty"`
	Embedded          *Embedded              `protobuf:"bytes,10,opt,name=embedded,proto3" json:"embedded,omitempty"`
	RepeatedEmbeddeds []*Embedded            `protobuf:"bytes,11,rep,name=repeated_embeddeds,json=repeatedEmbeddeds,proto3" json:"repeated_embeddeds,omitempty"`
	// issue #4
	OptBool       *bool             `protobuf:"varint,12,opt,name=opt_bool,json=optBool,proto3,oneof" json:"opt_bool,omitempty"`
	OptEnum       *Test_OEnum       `protobuf:"varint,13,opt,name=opt_enum,json=optEnum,proto3,enum=Test_OEnum,oneof" json:"opt_enum,omitempty"`
	Map           map[string]string `protobuf:"bytes,14,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Test) Reset() {
	*x = Test{}
	mi := &file_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetD() float64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *Test) GetF() float32 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *Test) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *Test) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *Test) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

func (x *Test) GetU64() uint64 {
	if x != nil {
		return x.U64
	}
	return 0
}

func (x *Test) GetT() bool {
	if x != nil {
		return x.T
	}
	return false
}

func (x *Test) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Test) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *Test) GetEmbedded() *Embedded {
	if x != nil {
		return x.Embedded
	}
	return nil
}

func (x *Test) GetRepeatedEmbeddeds() []*Embedded {
	if x != nil {
		return x.RepeatedEmbeddeds
	}
	return nil
}

func (x *Test) GetOptBool() bool {
	if x != nil && x.OptBool != nil {
		return *x.OptBool
	}
	return false
}

func (x *Test) GetOptEnum() Test_OEnum {
	if x != nil && x.OptEnum != nil {
		return *x.OptEnum
	}
	return Test_O_ENUM1
}

func (x *Test) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

type Embedded struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	S             string                 `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	Embedded      *Embedded              `protobuf:"bytes,2,opt,name=embedded,proto3" json:"embedded,omitempty"`
	OptBool       *bool                  `protobuf:"varint,3,opt,name=opt_bool,json=optBool,proto3,oneof" json:"opt_bool,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Embedded) Reset() {
	*x = Embedded{}
	mi := &file_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embedded) ProtoMessage() {}

func (x *Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embedded.ProtoReflect.Descriptor instead.
func (*Embedded) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Embedded) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *Embedded) GetEmbedded() *Embedded {
	if x != nil {
		return x.Embedded
	}
	return nil
}

func (x *Embedded) GetOptBool() bool {
	if x != nil && x.OptBool != nil {
		return *x.OptBool
	}
	return false
}

// Issue #9
type Foo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Bar:
	//
	//	*Foo_Baz
	//	*Foo_Quix
	Bar           isFoo_Bar `protobuf_oneof:"bar"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Foo) Reset() {
	*x = Foo{}
	mi := &file_test_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *Foo) GetBar() isFoo_Bar {
	if x != nil {
		return x.Bar
	}
	return nil
}

func (x *Foo) GetBaz() string {
	if x != nil {
		if x, ok := x.Bar.(*Foo_Baz); ok {
			return x.Baz
		}
	}
	return ""
}

func (x *Foo) GetQuix() string {
	if x != nil {
		if x, ok := x.Bar.(*Foo_Quix); ok {
			return x.Quix
		}
	}
	return ""
}

type isFoo_Bar interface {
	isFoo_Bar()
}

type Foo_Baz struct {
	Baz string `protobuf:"bytes,1,opt,name=baz,proto3,oneof"`
}

type Foo_Quix struct {
	Quix string `protobuf:"bytes,2,opt,name=quix,proto3,oneof"`
}

func (*Foo_Baz) isFoo_Bar() {}

func (*Foo_Quix) isFoo_Bar() {}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a,
	0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x66, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x69, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x75, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x52, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x12, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x64, 0x52, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6c,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x42, 0x6f, 0x6f,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x45,
	0x6e, 0x75, 0x6d, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x88, 0x01,
	0x01, 0x12, 0x20, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03,
	0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x21, 0x0a, 0x05, 0x4f,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x31, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x32, 0x10, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6f, 0x70, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6f, 0x70, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x22, 0x6c, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52,
	0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x70, 0x74,
	0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x70,
	0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x22, 0x36, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x12, 0x0a,
	0x03, 0x62, 0x61, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x62, 0x61,
	0x7a, 0x12, 0x14, 0x0a, 0x04, 0x71, 0x75, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x71, 0x75, 0x69, 0x78, 0x42, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x32, 0x1f,
	0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x61, 0x6c,
	0x6c, 0x12, 0x05, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData []byte
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_test_proto_rawDesc), len(file_test_proto_rawDesc)))
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_proto_goTypes = []any{
	(Test_OEnum)(0),  // 0: Test.OEnum
	(*Test)(nil),     // 1: Test
	(*Embedded)(nil), // 2: Embedded
	(*Foo)(nil),      // 3: Foo
	nil,              // 4: Test.MapEntry
}
var file_test_proto_depIdxs = []int32{
	2, // 0: Test.embedded:type_name -> Embedded
	2, // 1: Test.repeated_embeddeds:type_name -> Embedded
	0, // 2: Test.opt_enum:type_name -> Test.OEnum
	4, // 3: Test.map:type_name -> Test.MapEntry
	2, // 4: Embedded.embedded:type_name -> Embedded
	1, // 5: Testing.call:input_type -> Test
	1, // 6: Testing.call:output_type -> Test
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	file_test_proto_msgTypes[0].OneofWrappers = []any{}
	file_test_proto_msgTypes[1].OneofWrappers = []any{}
	file_test_proto_msgTypes[2].OneofWrappers = []any{
		(*Foo_Baz)(nil),
		(*Foo_Quix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_test_proto_rawDesc), len(file_test_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
