// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/CommonModels.proto

package commonmodels

import (
	proto "github.com/golang/protobuf/proto"
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

type Origin_OriginTypeEnum int32

const (
	Origin_ObjectStorage Origin_OriginTypeEnum = 0
	Origin_OriginLink    Origin_OriginTypeEnum = 1
)

// Enum value maps for Origin_OriginTypeEnum.
var (
	Origin_OriginTypeEnum_name = map[int32]string{
		0: "ObjectStorage",
		1: "OriginLink",
	}
	Origin_OriginTypeEnum_value = map[string]int32{
		"ObjectStorage": 0,
		"OriginLink":    1,
	}
)

func (x Origin_OriginTypeEnum) Enum() *Origin_OriginTypeEnum {
	p := new(Origin_OriginTypeEnum)
	*p = x
	return p
}

func (x Origin_OriginTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Origin_OriginTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_CommonModels_proto_enumTypes[0].Descriptor()
}

func (Origin_OriginTypeEnum) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[0]
}

func (x Origin_OriginTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Origin_OriginTypeEnum.Descriptor instead.
func (Origin_OriginTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{1, 0}
}

type DatasetEntityMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename   string  `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Filetype   string  `protobuf:"bytes,2,opt,name=Filetype,proto3" json:"Filetype,omitempty"`
	Name       string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Version    string  `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
	Origin     *Origin `protobuf:"bytes,5,opt,name=Origin,proto3" json:"Origin,omitempty"`
	ContentLen int64   `protobuf:"varint,6,opt,name=ContentLen,proto3" json:"ContentLen,omitempty"`
}

func (x *DatasetEntityMetaData) Reset() {
	*x = DatasetEntityMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetEntityMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetEntityMetaData) ProtoMessage() {}

func (x *DatasetEntityMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetEntityMetaData.ProtoReflect.Descriptor instead.
func (*DatasetEntityMetaData) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetEntityMetaData) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DatasetEntityMetaData) GetFiletype() string {
	if x != nil {
		return x.Filetype
	}
	return ""
}

func (x *DatasetEntityMetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatasetEntityMetaData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DatasetEntityMetaData) GetOrigin() *Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DatasetEntityMetaData) GetContentLen() int64 {
	if x != nil {
		return x.ContentLen
	}
	return 0
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link       string                `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	Bucket     string                `protobuf:"bytes,2,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Key        string                `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	OriginType Origin_OriginTypeEnum `protobuf:"varint,4,opt,name=OriginType,proto3,enum=Origin_OriginTypeEnum" json:"OriginType,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{1}
}

func (x *Origin) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Origin) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Origin) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Origin) GetOriginType() Origin_OriginTypeEnum {
	if x != nil {
		return x.OriginType
	}
	return Origin_ObjectStorage
}

var File_proto_CommonModels_proto protoreflect.FileDescriptor

var file_proto_CommonModels_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x15, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x06,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x0e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x11,
	0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x10,
	0x01, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_CommonModels_proto_rawDescOnce sync.Once
	file_proto_CommonModels_proto_rawDescData = file_proto_CommonModels_proto_rawDesc
)

func file_proto_CommonModels_proto_rawDescGZIP() []byte {
	file_proto_CommonModels_proto_rawDescOnce.Do(func() {
		file_proto_CommonModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CommonModels_proto_rawDescData)
	})
	return file_proto_CommonModels_proto_rawDescData
}

var file_proto_CommonModels_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_CommonModels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_CommonModels_proto_goTypes = []interface{}{
	(Origin_OriginTypeEnum)(0),    // 0: Origin.OriginTypeEnum
	(*DatasetEntityMetaData)(nil), // 1: DatasetEntityMetaData
	(*Origin)(nil),                // 2: Origin
}
var file_proto_CommonModels_proto_depIdxs = []int32{
	2, // 0: DatasetEntityMetaData.Origin:type_name -> Origin
	0, // 1: Origin.OriginType:type_name -> Origin.OriginTypeEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_CommonModels_proto_init() }
func file_proto_CommonModels_proto_init() {
	if File_proto_CommonModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_CommonModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetEntityMetaData); i {
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
		file_proto_CommonModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
			RawDescriptor: file_proto_CommonModels_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CommonModels_proto_goTypes,
		DependencyIndexes: file_proto_CommonModels_proto_depIdxs,
		EnumInfos:         file_proto_CommonModels_proto_enumTypes,
		MessageInfos:      file_proto_CommonModels_proto_msgTypes,
	}.Build()
	File_proto_CommonModels_proto = out.File
	file_proto_CommonModels_proto_rawDesc = nil
	file_proto_CommonModels_proto_goTypes = nil
	file_proto_CommonModels_proto_depIdxs = nil
}
