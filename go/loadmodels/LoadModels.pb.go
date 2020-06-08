// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/LoadModels.proto

package loadmodels

import (
	commonmodels "github.com/ag-computational-bio/BioDataDBModels/go/commonmodels"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{0}
}

type InitLoadDataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetName string `protobuf:"bytes,1,opt,name=DatasetName,proto3" json:"DatasetName,omitempty"`
	Datatype    string `protobuf:"bytes,2,opt,name=Datatype,proto3" json:"Datatype,omitempty"`
}

func (x *InitLoadDataset) Reset() {
	*x = InitLoadDataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitLoadDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitLoadDataset) ProtoMessage() {}

func (x *InitLoadDataset) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use InitLoadDataset.ProtoReflect.Descriptor instead.
func (*InitLoadDataset) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{1}
}

func (x *InitLoadDataset) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

func (x *InitLoadDataset) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

type UploadLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links []*UploadLink `protobuf:"bytes,1,rep,name=Links,proto3" json:"Links,omitempty"`
}

func (x *UploadLinks) Reset() {
	*x = UploadLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLinks) ProtoMessage() {}

func (x *UploadLinks) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadLinks.ProtoReflect.Descriptor instead.
func (*UploadLinks) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{2}
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

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Link string `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *UploadLink) Reset() {
	*x = UploadLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLink) ProtoMessage() {}

func (x *UploadLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[3]
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
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{3}
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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{4}
}

func (x *ID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateLoadLinkSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetID string                                `protobuf:"bytes,1,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"`
	Metadata  []*commonmodels.DatasetEntityMetaData `protobuf:"bytes,2,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (x *CreateLoadLinkSetRequest) Reset() {
	*x = CreateLoadLinkSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoadLinkSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoadLinkSetRequest) ProtoMessage() {}

func (x *CreateLoadLinkSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[5]
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
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLoadLinkSetRequest) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

func (x *CreateLoadLinkSetRequest) GetMetadata() []*commonmodels.DatasetEntityMetaData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type DatasetLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DatasetLink) Reset() {
	*x = DatasetLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_LoadModels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetLink) ProtoMessage() {}

func (x *DatasetLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_LoadModels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetLink.ProtoReflect.Descriptor instead.
func (*DatasetLink) Descriptor() ([]byte, []int) {
	return file_proto_LoadModels_proto_rawDescGZIP(), []int{6}
}

func (x *DatasetLink) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_proto_LoadModels_proto protoreflect.FileDescriptor

var file_proto_LoadModels_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4f, 0x0a, 0x0f, 0x49,
	0x6e, 0x69, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x22, 0x30, 0x0a, 0x0b,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x30,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x6c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x32, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_LoadModels_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_LoadModels_proto_goTypes = []interface{}{
	(*Empty)(nil),                              // 0: Empty
	(*InitLoadDataset)(nil),                    // 1: InitLoadDataset
	(*UploadLinks)(nil),                        // 2: UploadLinks
	(*UploadLink)(nil),                         // 3: UploadLink
	(*ID)(nil),                                 // 4: ID
	(*CreateLoadLinkSetRequest)(nil),           // 5: CreateLoadLinkSetRequest
	(*DatasetLink)(nil),                        // 6: DatasetLink
	(*commonmodels.DatasetEntityMetaData)(nil), // 7: DatasetEntityMetaData
}
var file_proto_LoadModels_proto_depIdxs = []int32{
	3, // 0: UploadLinks.Links:type_name -> UploadLink
	7, // 1: CreateLoadLinkSetRequest.Metadata:type_name -> DatasetEntityMetaData
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
			switch v := v.(*Empty); i {
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
			switch v := v.(*InitLoadDataset); i {
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
		file_proto_LoadModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_LoadModels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_proto_LoadModels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_LoadModels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetLink); i {
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
			NumMessages:   7,
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
