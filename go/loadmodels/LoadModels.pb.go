// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/LoadModels.proto

package loadmodels

import (
	datasetmodels "github.com/ag-computational-bio/BioDataDBModels/go/datasetmodels"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type UploadLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*UploadLink `protobuf:"bytes,1,rep,name=Links,proto3" json:"Links,omitempty"` // List of Links to upload a dataset
}

func (x *UploadLinks) Reset() {
	*x = UploadLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLinks) ProtoMessage() {}

func (x *UploadLinks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadLinks.ProtoReflect.Descriptor instead.
func (*UploadLinks) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{0}
}

func (x *UploadLinks) GetLinks() []*UploadLink {
	if x != nil {
		return x.Links
	}
	return nil
}

type UploadLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`     // ID of the upload request
	Link string `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"` // Upload link (normally a presigned put URL for S3)
}

func (x *UploadLink) Reset() {
	*x = UploadLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLink) ProtoMessage() {}

func (x *UploadLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadLink.ProtoReflect.Descriptor instead.
func (*UploadLink) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{1}
}

func (x *UploadLink) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UploadLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateLoadLinkSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetID        string                                 `protobuf:"bytes,1,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"` // ID of the dataset upload links should be created for
	DatasetVersionID string                                 `protobuf:"bytes,2,opt,name=DatasetVersionID,proto3" json:"DatasetVersionID,omitempty"`
	Metadata         []*datasetmodels.DatasetObjectMetaData `protobuf:"bytes,3,rep,name=Metadata,proto3" json:"Metadata,omitempty"` // List of metadata of uploaded objects
}

func (x *CreateLoadLinkSetRequest) Reset() {
	*x = CreateLoadLinkSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoadLinkSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoadLinkSetRequest) ProtoMessage() {}

func (x *CreateLoadLinkSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoadLinkSetRequest.ProtoReflect.Descriptor instead.
func (*CreateLoadLinkSetRequest) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLoadLinkSetRequest) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

func (x *CreateLoadLinkSetRequest) GetDatasetVersionID() string {
	if x != nil {
		return x.DatasetVersionID
	}
	return ""
}

func (x *CreateLoadLinkSetRequest) GetMetadata() []*datasetmodels.DatasetObjectMetaData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_proto_LoadModels_proto protoreflect.FileDescriptor

var file_proto_LoadModels_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a,
	0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x21, 0x0a, 0x05,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x3a,
	0x40, 0x92, 0x41, 0x3d, 0x0a, 0x3b, 0x2a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x32, 0x2c, 0x41, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x70, 0x75, 0x74, 0x20, 0x55, 0x52, 0x4c, 0x20,
	0x74, 0x6f, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x74, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c,
	0x69, 0x6e, 0x6b, 0x3a, 0x42, 0x92, 0x41, 0x3f, 0x0a, 0x3d, 0x2a, 0x0a, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x2f, 0x41, 0x20, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x69, 0x74, 0x73, 0x20, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x20, 0x55, 0x52, 0x4c, 0x22, 0xf2, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x32,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x58, 0x92, 0x41, 0x55, 0x0a, 0x53, 0x2a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x37, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
	0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x20, 0x55, 0x52, 0x4c, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f,
	0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_LoadModels_proto_rawDescOnce sync.Once
	file_proto_LoadModels_proto_rawDescData = file_proto_LoadModels_proto_rawDesc
)

func file_proto_LoadModels_proto_rawDescGZIP() []byte {
	file_proto_LoadModels_proto_rawDescOnce.Do(func() {
		file_proto_LoadModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_LoadModels_proto_rawDescData)
	})
	return file_proto_LoadModels_proto_rawDescData
}

var file_proto_LoadModels_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_LoadModels_proto_goTypes = []interface{}{
	(*UploadLinks)(nil),                         // 0: UploadLinks
	(*UploadLink)(nil),                          // 1: UploadLink
	(*CreateLoadLinkSetRequest)(nil),            // 2: CreateLoadLinkSetRequest
	(*datasetmodels.DatasetObjectMetaData)(nil), // 3: DatasetObjectMetaData
}
var file_proto_LoadModels_proto_depIdxs = []int32{
	1, // 0: UploadLinks.Links:type_name -> UploadLink
	3, // 1: CreateLoadLinkSetRequest.Metadata:type_name -> DatasetObjectMetaData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_LoadModels_proto_init() }
func file_proto_LoadModels_proto_init() {
	if File_proto_LoadModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_LoadModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadLinks); i {
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
		file_proto_LoadModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadLink); i {
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
		file_proto_LoadModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoadLinkSetRequest); i {
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
			RawDescriptor: file_proto_LoadModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_LoadModels_proto_goTypes,
		DependencyIndexes: file_proto_LoadModels_proto_depIdxs,
		MessageInfos:      file_proto_LoadModels_proto_msgTypes,
	}.Build()
	File_proto_LoadModels_proto = out.File
	file_proto_LoadModels_proto_rawDesc = nil
	file_proto_LoadModels_proto_goTypes = nil
	file_proto_LoadModels_proto_depIdxs = nil
}
