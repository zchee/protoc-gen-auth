// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: auth.proto

package auth

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FieldBehaviour int32

const (
	FieldBehaviour_ID  FieldBehaviour = 0
	FieldBehaviour_IDS FieldBehaviour = 1
)

// Enum value maps for FieldBehaviour.
var (
	FieldBehaviour_name = map[int32]string{
		0: "ID",
		1: "IDS",
	}
	FieldBehaviour_value = map[string]int32{
		"ID":  0,
		"IDS": 1,
	}
)

func (x FieldBehaviour) Enum() *FieldBehaviour {
	p := new(FieldBehaviour)
	*p = x
	return p
}

func (x FieldBehaviour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldBehaviour) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (FieldBehaviour) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x FieldBehaviour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldBehaviour.Descriptor instead.
func (FieldBehaviour) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

var file_auth_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldBehaviour)(nil),
		Field:         1072,
		Name:          "auth.field_behaviour",
		Tag:           "varint,1072,opt,name=field_behaviour,enum=auth.FieldBehaviour",
		Filename:      "auth.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1072,
		Name:          "auth.message_permission",
		Tag:           "bytes,1072,opt,name=message_permission",
		Filename:      "auth.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional auth.FieldBehaviour field_behaviour = 1072;
	E_FieldBehaviour = &file_auth_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional string message_permission = 1072;
	E_MessagePermission = &file_auth_proto_extTypes[1]
)

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x21, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x44, 0x53, 0x10, 0x01, 0x3a, 0x5d, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x3a, 0x4f, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_goTypes = []interface{}{
	(FieldBehaviour)(0),               // 0: auth.FieldBehaviour
	(*descriptor.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: auth.field_behaviour:extendee -> google.protobuf.FieldOptions
	2, // 1: auth.message_permission:extendee -> google.protobuf.MessageOptions
	0, // 2: auth.field_behaviour:type_name -> auth.FieldBehaviour
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		ExtensionInfos:    file_auth_proto_extTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
