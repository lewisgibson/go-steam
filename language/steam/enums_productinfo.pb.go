// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: enums_productinfo.proto

package steam

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

type EContentDescriptorID int32

const (
	EContentDescriptorID_k_EContentDescriptor_NudityOrSexualContent   EContentDescriptorID = 1
	EContentDescriptorID_k_EContentDescriptor_FrequentViolenceOrGore  EContentDescriptorID = 2
	EContentDescriptorID_k_EContentDescriptor_AdultOnlySexualContent  EContentDescriptorID = 3
	EContentDescriptorID_k_EContentDescriptor_GratuitousSexualContent EContentDescriptorID = 4
	EContentDescriptorID_k_EContentDescriptor_AnyMatureContent        EContentDescriptorID = 5
	EContentDescriptorID_k_EContentDescriptorMAX                      EContentDescriptorID = 6
)

// Enum value maps for EContentDescriptorID.
var (
	EContentDescriptorID_name = map[int32]string{
		1: "k_EContentDescriptor_NudityOrSexualContent",
		2: "k_EContentDescriptor_FrequentViolenceOrGore",
		3: "k_EContentDescriptor_AdultOnlySexualContent",
		4: "k_EContentDescriptor_GratuitousSexualContent",
		5: "k_EContentDescriptor_AnyMatureContent",
		6: "k_EContentDescriptorMAX",
	}
	EContentDescriptorID_value = map[string]int32{
		"k_EContentDescriptor_NudityOrSexualContent":   1,
		"k_EContentDescriptor_FrequentViolenceOrGore":  2,
		"k_EContentDescriptor_AdultOnlySexualContent":  3,
		"k_EContentDescriptor_GratuitousSexualContent": 4,
		"k_EContentDescriptor_AnyMatureContent":        5,
		"k_EContentDescriptorMAX":                      6,
	}
)

func (x EContentDescriptorID) Enum() *EContentDescriptorID {
	p := new(EContentDescriptorID)
	*p = x
	return p
}

func (x EContentDescriptorID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EContentDescriptorID) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_productinfo_proto_enumTypes[0].Descriptor()
}

func (EContentDescriptorID) Type() protoreflect.EnumType {
	return &file_enums_productinfo_proto_enumTypes[0]
}

func (x EContentDescriptorID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EContentDescriptorID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EContentDescriptorID(num)
	return nil
}

// Deprecated: Use EContentDescriptorID.Descriptor instead.
func (EContentDescriptorID) EnumDescriptor() ([]byte, []int) {
	return file_enums_productinfo_proto_rawDescGZIP(), []int{0}
}

var File_enums_productinfo_proto protoreflect.FileDescriptor

var file_enums_productinfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xa2, 0x02, 0x0a, 0x14, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x2a,
	0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x5f, 0x4e, 0x75, 0x64, 0x69, 0x74, 0x79, 0x4f, 0x72, 0x53, 0x65, 0x78,
	0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x2f, 0x0a, 0x2b,
	0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x5f, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x6f,
	0x6c, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x72, 0x47, 0x6f, 0x72, 0x65, 0x10, 0x02, 0x12, 0x2f, 0x0a,
	0x2b, 0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x41, 0x64, 0x75, 0x6c, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x53,
	0x65, 0x78, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x30,
	0x0a, 0x2c, 0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x47, 0x72, 0x61, 0x74, 0x75, 0x69, 0x74, 0x6f, 0x75,
	0x73, 0x53, 0x65, 0x78, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x04,
	0x12, 0x29, 0x0a, 0x25, 0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x41, 0x6e, 0x79, 0x4d, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x6b,
	0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x4d, 0x41, 0x58, 0x10, 0x06, 0x42, 0x09, 0x80, 0xb5, 0x18, 0x01, 0x48, 0x01,
	0x80, 0x01, 0x01,
}

var (
	file_enums_productinfo_proto_rawDescOnce sync.Once
	file_enums_productinfo_proto_rawDescData = file_enums_productinfo_proto_rawDesc
)

func file_enums_productinfo_proto_rawDescGZIP() []byte {
	file_enums_productinfo_proto_rawDescOnce.Do(func() {
		file_enums_productinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_productinfo_proto_rawDescData)
	})
	return file_enums_productinfo_proto_rawDescData
}

var file_enums_productinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_productinfo_proto_goTypes = []any{
	(EContentDescriptorID)(0), // 0: EContentDescriptorID
}
var file_enums_productinfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_productinfo_proto_init() }
func file_enums_productinfo_proto_init() {
	if File_enums_productinfo_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_productinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_productinfo_proto_goTypes,
		DependencyIndexes: file_enums_productinfo_proto_depIdxs,
		EnumInfos:         file_enums_productinfo_proto_enumTypes,
	}.Build()
	File_enums_productinfo_proto = out.File
	file_enums_productinfo_proto_rawDesc = nil
	file_enums_productinfo_proto_goTypes = nil
	file_enums_productinfo_proto_depIdxs = nil
}
