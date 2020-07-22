// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/DatasetModels.proto

package datasetmodels

import (
	commonmodels "github.com/ag-computational-bio/BioDataDBModels/go/commonmodels"
	datasetentrymodels "github.com/ag-computational-bio/BioDataDBModels/go/datasetentrymodels"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type DatasetObjectLinks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Entites []*DatasetObjectLink `protobuf:"bytes,2,rep,name=Entites,proto3" json:"Entites,omitempty"`
}

func (x *DatasetObjectLinks) Reset() {
	*x = DatasetObjectLinks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetObjectLinks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetObjectLinks) ProtoMessage() {}

func (x *DatasetObjectLinks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetObjectLinks.ProtoReflect.Descriptor instead.
func (*DatasetObjectLinks) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetObjectLinks) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DatasetObjectLinks) GetEntites() []*DatasetObjectLink {
	if x != nil {
		return x.Entites
	}
	return nil
}

type DatasetObjectLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link           string                                    `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	ObjectID       string                                    `protobuf:"bytes,2,opt,name=ObjectID,proto3" json:"ObjectID,omitempty"`
	Objectmetadata *datasetentrymodels.DatasetObjectMetaData `protobuf:"bytes,3,opt,name=Objectmetadata,proto3" json:"Objectmetadata,omitempty"`
}

func (x *DatasetObjectLink) Reset() {
	*x = DatasetObjectLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetObjectLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetObjectLink) ProtoMessage() {}

func (x *DatasetObjectLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetObjectLink.ProtoReflect.Descriptor instead.
func (*DatasetObjectLink) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{1}
}

func (x *DatasetObjectLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *DatasetObjectLink) GetObjectID() string {
	if x != nil {
		return x.ObjectID
	}
	return ""
}

func (x *DatasetObjectLink) GetObjectmetadata() *datasetentrymodels.DatasetObjectMetaData {
	if x != nil {
		return x.Objectmetadata
	}
	return nil
}

type CreateDatasetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetName string `protobuf:"bytes,1,opt,name=DatasetName,proto3" json:"DatasetName,omitempty"` // Name of the dataset
	Datatype    string `protobuf:"bytes,2,opt,name=Datatype,proto3" json:"Datatype,omitempty"`       //Datatype of the dataset, e.g. json, gbff, fasta
}

func (x *CreateDatasetRequest) Reset() {
	*x = CreateDatasetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatasetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatasetRequest) ProtoMessage() {}

func (x *CreateDatasetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatasetRequest.ProtoReflect.Descriptor instead.
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDatasetRequest) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

func (x *CreateDatasetRequest) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

type UpdateDatasetVersionObjectCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetVersionID string `protobuf:"bytes,1,opt,name=DatasetVersionID,proto3" json:"DatasetVersionID,omitempty"`
	Value            int64  `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	DatasetID        string `protobuf:"bytes,3,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"`
}

func (x *UpdateDatasetVersionObjectCountRequest) Reset() {
	*x = UpdateDatasetVersionObjectCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDatasetVersionObjectCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatasetVersionObjectCountRequest) ProtoMessage() {}

func (x *UpdateDatasetVersionObjectCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatasetVersionObjectCountRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatasetVersionObjectCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDatasetVersionObjectCountRequest) GetDatasetVersionID() string {
	if x != nil {
		return x.DatasetVersionID
	}
	return ""
}

func (x *UpdateDatasetVersionObjectCountRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *UpdateDatasetVersionObjectCountRequest) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

type DatasetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datasets []*datasetentrymodels.DatasetEntry `protobuf:"bytes,1,rep,name=Datasets,proto3" json:"Datasets,omitempty"`
}

func (x *DatasetList) Reset() {
	*x = DatasetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetList) ProtoMessage() {}

func (x *DatasetList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetList.ProtoReflect.Descriptor instead.
func (*DatasetList) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{4}
}

func (x *DatasetList) GetDatasets() []*datasetentrymodels.DatasetEntry {
	if x != nil {
		return x.Datasets
	}
	return nil
}

type DatasetVersionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetVersions []*datasetentrymodels.DatasetVersionEntry `protobuf:"bytes,2,rep,name=DatasetVersions,proto3" json:"DatasetVersions,omitempty"`
}

func (x *DatasetVersionList) Reset() {
	*x = DatasetVersionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetVersionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetVersionList) ProtoMessage() {}

func (x *DatasetVersionList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetVersionList.ProtoReflect.Descriptor instead.
func (*DatasetVersionList) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{5}
}

func (x *DatasetVersionList) GetDatasetVersions() []*datasetentrymodels.DatasetVersionEntry {
	if x != nil {
		return x.DatasetVersions
	}
	return nil
}

type DatasetObjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetVersionObjects []*datasetentrymodels.DatasetObjectEntry `protobuf:"bytes,1,rep,name=DatasetVersionObjects,proto3" json:"DatasetVersionObjects,omitempty"`
}

func (x *DatasetObjectList) Reset() {
	*x = DatasetObjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetObjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetObjectList) ProtoMessage() {}

func (x *DatasetObjectList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetObjectList.ProtoReflect.Descriptor instead.
func (*DatasetObjectList) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{6}
}

func (x *DatasetObjectList) GetDatasetVersionObjects() []*datasetentrymodels.DatasetObjectEntry {
	if x != nil {
		return x.DatasetVersionObjects
	}
	return nil
}

type NewDatasetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetID                          string                `protobuf:"bytes,1,opt,name=DatasetID,proto3" json:"DatasetID,omitempty"`                                                   // ID of the parent Dataset
	Version                            *commonmodels.Version `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`                                                       // Version of the Dataset
	AdditionalMetadataMessageRef       string                `protobuf:"bytes,3,opt,name=AdditionalMetadataMessageRef,proto3" json:"AdditionalMetadataMessageRef,omitempty"`             // Message reference for the metadata
	AdditionalObjectMetadataMessageRef string                `protobuf:"bytes,4,opt,name=AdditionalObjectMetadataMessageRef,proto3" json:"AdditionalObjectMetadataMessageRef,omitempty"` // Message reference for the metadata of the objects associated with this DatasetVersion
	AdditionalMetadata                 *_struct.Struct       `protobuf:"bytes,5,opt,name=AdditionalMetadata,proto3" json:"AdditionalMetadata,omitempty"`                                 // Additional metadata for the dataset version
	ExpectedObjectCount                int64                 `protobuf:"varint,6,opt,name=ExpectedObjectCount,proto3" json:"ExpectedObjectCount,omitempty"`                              // Expected objects
}

func (x *NewDatasetVersionRequest) Reset() {
	*x = NewDatasetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_DatasetModels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewDatasetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewDatasetVersionRequest) ProtoMessage() {}

func (x *NewDatasetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_DatasetModels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewDatasetVersionRequest.ProtoReflect.Descriptor instead.
func (*NewDatasetVersionRequest) Descriptor() ([]byte, []int) {
	return file_proto_DatasetModels_proto_rawDescGZIP(), []int{7}
}

func (x *NewDatasetVersionRequest) GetDatasetID() string {
	if x != nil {
		return x.DatasetID
	}
	return ""
}

func (x *NewDatasetVersionRequest) GetVersion() *commonmodels.Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *NewDatasetVersionRequest) GetAdditionalMetadataMessageRef() string {
	if x != nil {
		return x.AdditionalMetadataMessageRef
	}
	return ""
}

func (x *NewDatasetVersionRequest) GetAdditionalObjectMetadataMessageRef() string {
	if x != nil {
		return x.AdditionalObjectMetadataMessageRef
	}
	return ""
}

func (x *NewDatasetVersionRequest) GetAdditionalMetadata() *_struct.Struct {
	if x != nil {
		return x.AdditionalMetadata
	}
	return nil
}

func (x *NewDatasetVersionRequest) GetExpectedObjectCount() int64 {
	if x != nil {
		return x.ExpectedObjectCount
	}
	return 0
}

var File_proto_DatasetModels_proto protoreflect.FileDescriptor

var file_proto_DatasetModels_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01,
	0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x07, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x65, 0x73, 0x3a, 0x48, 0x92, 0x41, 0x45, 0x0a, 0x43, 0x2a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x32, 0x33, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x20,
	0x6f, 0x66, 0x20, 0x61, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0xcf, 0x01, 0x0a,
	0x11, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x3e, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x3a, 0x4a, 0x92, 0x41, 0x47, 0x0a, 0x45, 0x2a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x30, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x20,
	0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0xdd,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x86, 0x01, 0x92, 0x41, 0x82, 0x01, 0x0a, 0x4e, 0x2a, 0x0f,
	0x49, 0x6e, 0x69, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x32,
	0x3b, 0x44, 0x61, 0x74, 0x61, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x20,
	0x61, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x32, 0x30, 0x12, 0x2e,
	0x7b, 0x20, 0x22, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x73, 0x74, 0x32, 0x22, 0x20, 0x7d, 0x22, 0x88,
	0x01, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x44, 0x22, 0x60, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x3a, 0x26, 0x92, 0x41, 0x23, 0x0a, 0x21, 0x2a, 0x0b, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x12, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x12,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x2d, 0x92, 0x41, 0x2a, 0x0a, 0x28, 0x2a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x19, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x3a, 0x26, 0x92, 0x41, 0x23, 0x0a, 0x21, 0x2a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x12, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x22, 0xab, 0x03, 0x0a, 0x18, 0x4e,
	0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x1c, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1c, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x12, 0x4e, 0x0a,
	0x22, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x66, 0x12, 0x47, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x12, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x13, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3e, 0x92, 0x41, 0x3b, 0x0a, 0x39, 0x2a,
	0x18, 0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44,
	0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_DatasetModels_proto_rawDescOnce sync.Once
	file_proto_DatasetModels_proto_rawDescData = file_proto_DatasetModels_proto_rawDesc
)

func file_proto_DatasetModels_proto_rawDescGZIP() []byte {
	file_proto_DatasetModels_proto_rawDescOnce.Do(func() {
		file_proto_DatasetModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_DatasetModels_proto_rawDescData)
	})
	return file_proto_DatasetModels_proto_rawDescData
}

var file_proto_DatasetModels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_DatasetModels_proto_goTypes = []interface{}{
	(*DatasetObjectLinks)(nil),                       // 0: DatasetObjectLinks
	(*DatasetObjectLink)(nil),                        // 1: DatasetObjectLink
	(*CreateDatasetRequest)(nil),                     // 2: CreateDatasetRequest
	(*UpdateDatasetVersionObjectCountRequest)(nil),   // 3: UpdateDatasetVersionObjectCountRequest
	(*DatasetList)(nil),                              // 4: DatasetList
	(*DatasetVersionList)(nil),                       // 5: DatasetVersionList
	(*DatasetObjectList)(nil),                        // 6: DatasetObjectList
	(*NewDatasetVersionRequest)(nil),                 // 7: NewDatasetVersionRequest
	(*datasetentrymodels.DatasetObjectMetaData)(nil), // 8: DatasetObjectMetaData
	(*datasetentrymodels.DatasetEntry)(nil),          // 9: DatasetEntry
	(*datasetentrymodels.DatasetVersionEntry)(nil),   // 10: DatasetVersionEntry
	(*datasetentrymodels.DatasetObjectEntry)(nil),    // 11: DatasetObjectEntry
	(*commonmodels.Version)(nil),                     // 12: Version
	(*_struct.Struct)(nil),                           // 13: google.protobuf.Struct
}
var file_proto_DatasetModels_proto_depIdxs = []int32{
	1,  // 0: DatasetObjectLinks.Entites:type_name -> DatasetObjectLink
	8,  // 1: DatasetObjectLink.Objectmetadata:type_name -> DatasetObjectMetaData
	9,  // 2: DatasetList.Datasets:type_name -> DatasetEntry
	10, // 3: DatasetVersionList.DatasetVersions:type_name -> DatasetVersionEntry
	11, // 4: DatasetObjectList.DatasetVersionObjects:type_name -> DatasetObjectEntry
	12, // 5: NewDatasetVersionRequest.Version:type_name -> Version
	13, // 6: NewDatasetVersionRequest.AdditionalMetadata:type_name -> google.protobuf.Struct
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_DatasetModels_proto_init() }
func file_proto_DatasetModels_proto_init() {
	if File_proto_DatasetModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_DatasetModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetObjectLinks); i {
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
		file_proto_DatasetModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetObjectLink); i {
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
		file_proto_DatasetModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatasetRequest); i {
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
		file_proto_DatasetModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDatasetVersionObjectCountRequest); i {
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
		file_proto_DatasetModels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetList); i {
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
		file_proto_DatasetModels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetVersionList); i {
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
		file_proto_DatasetModels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetObjectList); i {
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
		file_proto_DatasetModels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewDatasetVersionRequest); i {
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
			RawDescriptor: file_proto_DatasetModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_DatasetModels_proto_goTypes,
		DependencyIndexes: file_proto_DatasetModels_proto_depIdxs,
		MessageInfos:      file_proto_DatasetModels_proto_msgTypes,
	}.Build()
	File_proto_DatasetModels_proto = out.File
	file_proto_DatasetModels_proto_rawDesc = nil
	file_proto_DatasetModels_proto_goTypes = nil
	file_proto_DatasetModels_proto_depIdxs = nil
}
