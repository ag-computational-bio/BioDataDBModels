// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/APIs.proto

package api

import (
	context "context"
	commonmodels "github.com/ag-computational-bio/BioDataDBModels/go/commonmodels"
	datasetmodels "github.com/ag-computational-bio/BioDataDBModels/go/datasetmodels"
	loadmodels "github.com/ag-computational-bio/BioDataDBModels/go/loadmodels"
	metadatamodels "github.com/ag-computational-bio/BioDataDBModels/go/metadatamodels"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_proto_APIs_proto protoreflect.FileDescriptor

var file_proto_APIs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x50, 0x49, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x70, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x3a, 0x01, 0x2a, 0x32, 0xb6, 0x04, 0x0a, 0x0e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x41, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x74,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x4e, 0x65, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x15, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x3a, 0x01, 0x2a, 0x32, 0xdc, 0x03, 0x0a, 0x16, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x5e, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44,
	0x42, 0x12, 0x16, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x44, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x64, 0x62, 0x3a, 0x01, 0x2a, 0x12,
	0x70, 0x0a, 0x18, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44,
	0x42, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44, 0x42,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x69, 0x6e, 0x69, 0x74, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01,
	0x2a, 0x12, 0x54, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x01, 0x2a,
	0x12, 0x41, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x3a, 0x01, 0x2a, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_APIs_proto_goTypes = []interface{}{
	(*loadmodels.CreateLoadLinkSetRequest)(nil),    // 0: CreateLoadLinkSetRequest
	(*datasetmodels.CreateDatasetRequest)(nil),     // 1: CreateDatasetRequest
	(*commonmodels.Empty)(nil),                     // 2: Empty
	(*datasetmodels.NewDatasetVersionRequest)(nil), // 3: NewDatasetVersionRequest
	(*commonmodels.ID)(nil),                        // 4: ID
	(*metadatamodels.InitMetadataDBRequest)(nil),   // 5: InitMetadataDBRequest
	(*metadatamodels.InsertMetadataRequest)(nil),   // 6: InsertMetadataRequest
	(*metadatamodels.AddMetadataIndexRequest)(nil), // 7: AddMetadataIndexRequest
	(*metadatamodels.QueryRequest)(nil),            // 8: QueryRequest
	(*loadmodels.UploadLinks)(nil),                 // 9: UploadLinks
	(*datasetmodels.DatasetEntry)(nil),             // 10: DatasetEntry
	(*datasetmodels.DatasetList)(nil),              // 11: DatasetList
	(*datasetmodels.DatasetVersionEntry)(nil),      // 12: DatasetVersionEntry
	(*datasetmodels.DatasetVersionList)(nil),       // 13: DatasetVersionList
	(*datasetmodels.DatasetObjectList)(nil),        // 14: DatasetObjectList
	(*datasetmodels.DatasetObjectLinks)(nil),       // 15: DatasetObjectLinks
	(*metadatamodels.MetadataDBEntry)(nil),         // 16: MetadataDBEntry
	(*metadatamodels.Field)(nil),                   // 17: Field
}
var file_proto_APIs_proto_depIdxs = []int32{
	0,  // 0: LoadService.CreateLoadLinkSet:input_type -> CreateLoadLinkSetRequest
	1,  // 1: DatasetService.CreateNewDataset:input_type -> CreateDatasetRequest
	2,  // 2: DatasetService.Datasets:input_type -> Empty
	3,  // 3: DatasetService.CreateNewDatasetVersion:input_type -> NewDatasetVersionRequest
	4,  // 4: DatasetService.DatasetVersions:input_type -> ID
	4,  // 5: DatasetService.DatasetVersionObjects:input_type -> ID
	4,  // 6: DatasetService.DatasetVersionObjectsLinks:input_type -> ID
	5,  // 7: MetadataCompositeStore.InitMetadataDB:input_type -> InitMetadataDBRequest
	5,  // 8: MetadataCompositeStore.InitMetadataDBCollection:input_type -> InitMetadataDBRequest
	6,  // 9: MetadataCompositeStore.InsertMetadata:input_type -> InsertMetadataRequest
	7,  // 10: MetadataCompositeStore.AddMetadataIndex:input_type -> AddMetadataIndexRequest
	8,  // 11: MetadataCompositeStore.Query:input_type -> QueryRequest
	9,  // 12: LoadService.CreateLoadLinkSet:output_type -> UploadLinks
	10, // 13: DatasetService.CreateNewDataset:output_type -> DatasetEntry
	11, // 14: DatasetService.Datasets:output_type -> DatasetList
	12, // 15: DatasetService.CreateNewDatasetVersion:output_type -> DatasetVersionEntry
	13, // 16: DatasetService.DatasetVersions:output_type -> DatasetVersionList
	14, // 17: DatasetService.DatasetVersionObjects:output_type -> DatasetObjectList
	15, // 18: DatasetService.DatasetVersionObjectsLinks:output_type -> DatasetObjectLinks
	16, // 19: MetadataCompositeStore.InitMetadataDB:output_type -> MetadataDBEntry
	16, // 20: MetadataCompositeStore.InitMetadataDBCollection:output_type -> MetadataDBEntry
	2,  // 21: MetadataCompositeStore.InsertMetadata:output_type -> Empty
	2,  // 22: MetadataCompositeStore.AddMetadataIndex:output_type -> Empty
	17, // 23: MetadataCompositeStore.Query:output_type -> Field
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_APIs_proto_init() }
func file_proto_APIs_proto_init() {
	if File_proto_APIs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_APIs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proto_APIs_proto_goTypes,
		DependencyIndexes: file_proto_APIs_proto_depIdxs,
	}.Build()
	File_proto_APIs_proto = out.File
	file_proto_APIs_proto_rawDesc = nil
	file_proto_APIs_proto_goTypes = nil
	file_proto_APIs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoadServiceClient is the client API for LoadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadServiceClient interface {
	// Creates a list of upload links to place dataset entities in object storage
	// and adds corresponding metadata objects
	CreateLoadLinkSet(ctx context.Context, in *loadmodels.CreateLoadLinkSetRequest, opts ...grpc.CallOption) (*loadmodels.UploadLinks, error)
}

type loadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadServiceClient(cc grpc.ClientConnInterface) LoadServiceClient {
	return &loadServiceClient{cc}
}

func (c *loadServiceClient) CreateLoadLinkSet(ctx context.Context, in *loadmodels.CreateLoadLinkSetRequest, opts ...grpc.CallOption) (*loadmodels.UploadLinks, error) {
	out := new(loadmodels.UploadLinks)
	err := c.cc.Invoke(ctx, "/LoadService/CreateLoadLinkSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServiceServer is the server API for LoadService service.
type LoadServiceServer interface {
	// Creates a list of upload links to place dataset entities in object storage
	// and adds corresponding metadata objects
	CreateLoadLinkSet(context.Context, *loadmodels.CreateLoadLinkSetRequest) (*loadmodels.UploadLinks, error)
}

// UnimplementedLoadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadServiceServer struct {
}

func (*UnimplementedLoadServiceServer) CreateLoadLinkSet(context.Context, *loadmodels.CreateLoadLinkSetRequest) (*loadmodels.UploadLinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoadLinkSet not implemented")
}

func RegisterLoadServiceServer(s *grpc.Server, srv LoadServiceServer) {
	s.RegisterService(&_LoadService_serviceDesc, srv)
}

func _LoadService_CreateLoadLinkSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(loadmodels.CreateLoadLinkSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServiceServer).CreateLoadLinkSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoadService/CreateLoadLinkSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServiceServer).CreateLoadLinkSet(ctx, req.(*loadmodels.CreateLoadLinkSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LoadService",
	HandlerType: (*LoadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLoadLinkSet",
			Handler:    _LoadService_CreateLoadLinkSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/APIs.proto",
}

// DatasetServiceClient is the client API for DatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasetServiceClient interface {
	// Creates a new dataset
	CreateNewDataset(ctx context.Context, in *datasetmodels.CreateDatasetRequest, opts ...grpc.CallOption) (*datasetmodels.DatasetEntry, error)
	// Lists all datasets
	Datasets(ctx context.Context, in *commonmodels.Empty, opts ...grpc.CallOption) (*datasetmodels.DatasetList, error)
	// Creates a new dataset version based on an existing dataset
	CreateNewDatasetVersion(ctx context.Context, in *datasetmodels.NewDatasetVersionRequest, opts ...grpc.CallOption) (*datasetmodels.DatasetVersionEntry, error)
	// Lists Versions of a dataset
	DatasetVersions(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetVersionList, error)
	// Lists all entities of a dataset
	DatasetVersionObjects(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetObjectList, error)
	// Creates preauth download links for all entities of a dataset
	DatasetVersionObjectsLinks(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetObjectLinks, error)
}

type datasetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetServiceClient(cc grpc.ClientConnInterface) DatasetServiceClient {
	return &datasetServiceClient{cc}
}

func (c *datasetServiceClient) CreateNewDataset(ctx context.Context, in *datasetmodels.CreateDatasetRequest, opts ...grpc.CallOption) (*datasetmodels.DatasetEntry, error) {
	out := new(datasetmodels.DatasetEntry)
	err := c.cc.Invoke(ctx, "/DatasetService/CreateNewDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) Datasets(ctx context.Context, in *commonmodels.Empty, opts ...grpc.CallOption) (*datasetmodels.DatasetList, error) {
	out := new(datasetmodels.DatasetList)
	err := c.cc.Invoke(ctx, "/DatasetService/Datasets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) CreateNewDatasetVersion(ctx context.Context, in *datasetmodels.NewDatasetVersionRequest, opts ...grpc.CallOption) (*datasetmodels.DatasetVersionEntry, error) {
	out := new(datasetmodels.DatasetVersionEntry)
	err := c.cc.Invoke(ctx, "/DatasetService/CreateNewDatasetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) DatasetVersions(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetVersionList, error) {
	out := new(datasetmodels.DatasetVersionList)
	err := c.cc.Invoke(ctx, "/DatasetService/DatasetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) DatasetVersionObjects(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetObjectList, error) {
	out := new(datasetmodels.DatasetObjectList)
	err := c.cc.Invoke(ctx, "/DatasetService/DatasetVersionObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) DatasetVersionObjectsLinks(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetObjectLinks, error) {
	out := new(datasetmodels.DatasetObjectLinks)
	err := c.cc.Invoke(ctx, "/DatasetService/DatasetVersionObjectsLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetServiceServer is the server API for DatasetService service.
type DatasetServiceServer interface {
	// Creates a new dataset
	CreateNewDataset(context.Context, *datasetmodels.CreateDatasetRequest) (*datasetmodels.DatasetEntry, error)
	// Lists all datasets
	Datasets(context.Context, *commonmodels.Empty) (*datasetmodels.DatasetList, error)
	// Creates a new dataset version based on an existing dataset
	CreateNewDatasetVersion(context.Context, *datasetmodels.NewDatasetVersionRequest) (*datasetmodels.DatasetVersionEntry, error)
	// Lists Versions of a dataset
	DatasetVersions(context.Context, *commonmodels.ID) (*datasetmodels.DatasetVersionList, error)
	// Lists all entities of a dataset
	DatasetVersionObjects(context.Context, *commonmodels.ID) (*datasetmodels.DatasetObjectList, error)
	// Creates preauth download links for all entities of a dataset
	DatasetVersionObjectsLinks(context.Context, *commonmodels.ID) (*datasetmodels.DatasetObjectLinks, error)
}

// UnimplementedDatasetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDatasetServiceServer struct {
}

func (*UnimplementedDatasetServiceServer) CreateNewDataset(context.Context, *datasetmodels.CreateDatasetRequest) (*datasetmodels.DatasetEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewDataset not implemented")
}
func (*UnimplementedDatasetServiceServer) Datasets(context.Context, *commonmodels.Empty) (*datasetmodels.DatasetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Datasets not implemented")
}
func (*UnimplementedDatasetServiceServer) CreateNewDatasetVersion(context.Context, *datasetmodels.NewDatasetVersionRequest) (*datasetmodels.DatasetVersionEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewDatasetVersion not implemented")
}
func (*UnimplementedDatasetServiceServer) DatasetVersions(context.Context, *commonmodels.ID) (*datasetmodels.DatasetVersionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatasetVersions not implemented")
}
func (*UnimplementedDatasetServiceServer) DatasetVersionObjects(context.Context, *commonmodels.ID) (*datasetmodels.DatasetObjectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatasetVersionObjects not implemented")
}
func (*UnimplementedDatasetServiceServer) DatasetVersionObjectsLinks(context.Context, *commonmodels.ID) (*datasetmodels.DatasetObjectLinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatasetVersionObjectsLinks not implemented")
}

func RegisterDatasetServiceServer(s *grpc.Server, srv DatasetServiceServer) {
	s.RegisterService(&_DatasetService_serviceDesc, srv)
}

func _DatasetService_CreateNewDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datasetmodels.CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).CreateNewDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/CreateNewDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).CreateNewDataset(ctx, req.(*datasetmodels.CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_Datasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Datasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/Datasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Datasets(ctx, req.(*commonmodels.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_CreateNewDatasetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(datasetmodels.NewDatasetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).CreateNewDatasetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/CreateNewDatasetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).CreateNewDatasetVersion(ctx, req.(*datasetmodels.NewDatasetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_DatasetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).DatasetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/DatasetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).DatasetVersions(ctx, req.(*commonmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_DatasetVersionObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).DatasetVersionObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/DatasetVersionObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).DatasetVersionObjects(ctx, req.(*commonmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_DatasetVersionObjectsLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).DatasetVersionObjectsLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DatasetService/DatasetVersionObjectsLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).DatasetVersionObjectsLinks(ctx, req.(*commonmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DatasetService",
	HandlerType: (*DatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewDataset",
			Handler:    _DatasetService_CreateNewDataset_Handler,
		},
		{
			MethodName: "Datasets",
			Handler:    _DatasetService_Datasets_Handler,
		},
		{
			MethodName: "CreateNewDatasetVersion",
			Handler:    _DatasetService_CreateNewDatasetVersion_Handler,
		},
		{
			MethodName: "DatasetVersions",
			Handler:    _DatasetService_DatasetVersions_Handler,
		},
		{
			MethodName: "DatasetVersionObjects",
			Handler:    _DatasetService_DatasetVersionObjects_Handler,
		},
		{
			MethodName: "DatasetVersionObjectsLinks",
			Handler:    _DatasetService_DatasetVersionObjectsLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/APIs.proto",
}

// MetadataCompositeStoreClient is the client API for MetadataCompositeStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataCompositeStoreClient interface {
	InitMetadataDB(ctx context.Context, in *metadatamodels.InitMetadataDBRequest, opts ...grpc.CallOption) (*metadatamodels.MetadataDBEntry, error)
	InitMetadataDBCollection(ctx context.Context, in *metadatamodels.InitMetadataDBRequest, opts ...grpc.CallOption) (*metadatamodels.MetadataDBEntry, error)
	InsertMetadata(ctx context.Context, in *metadatamodels.InsertMetadataRequest, opts ...grpc.CallOption) (*commonmodels.Empty, error)
	AddMetadataIndex(ctx context.Context, in *metadatamodels.AddMetadataIndexRequest, opts ...grpc.CallOption) (*commonmodels.Empty, error)
	Query(ctx context.Context, in *metadatamodels.QueryRequest, opts ...grpc.CallOption) (*metadatamodels.Field, error)
}

type metadataCompositeStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataCompositeStoreClient(cc grpc.ClientConnInterface) MetadataCompositeStoreClient {
	return &metadataCompositeStoreClient{cc}
}

func (c *metadataCompositeStoreClient) InitMetadataDB(ctx context.Context, in *metadatamodels.InitMetadataDBRequest, opts ...grpc.CallOption) (*metadatamodels.MetadataDBEntry, error) {
	out := new(metadatamodels.MetadataDBEntry)
	err := c.cc.Invoke(ctx, "/MetadataCompositeStore/InitMetadataDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataCompositeStoreClient) InitMetadataDBCollection(ctx context.Context, in *metadatamodels.InitMetadataDBRequest, opts ...grpc.CallOption) (*metadatamodels.MetadataDBEntry, error) {
	out := new(metadatamodels.MetadataDBEntry)
	err := c.cc.Invoke(ctx, "/MetadataCompositeStore/InitMetadataDBCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataCompositeStoreClient) InsertMetadata(ctx context.Context, in *metadatamodels.InsertMetadataRequest, opts ...grpc.CallOption) (*commonmodels.Empty, error) {
	out := new(commonmodels.Empty)
	err := c.cc.Invoke(ctx, "/MetadataCompositeStore/InsertMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataCompositeStoreClient) AddMetadataIndex(ctx context.Context, in *metadatamodels.AddMetadataIndexRequest, opts ...grpc.CallOption) (*commonmodels.Empty, error) {
	out := new(commonmodels.Empty)
	err := c.cc.Invoke(ctx, "/MetadataCompositeStore/AddMetadataIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataCompositeStoreClient) Query(ctx context.Context, in *metadatamodels.QueryRequest, opts ...grpc.CallOption) (*metadatamodels.Field, error) {
	out := new(metadatamodels.Field)
	err := c.cc.Invoke(ctx, "/MetadataCompositeStore/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataCompositeStoreServer is the server API for MetadataCompositeStore service.
type MetadataCompositeStoreServer interface {
	InitMetadataDB(context.Context, *metadatamodels.InitMetadataDBRequest) (*metadatamodels.MetadataDBEntry, error)
	InitMetadataDBCollection(context.Context, *metadatamodels.InitMetadataDBRequest) (*metadatamodels.MetadataDBEntry, error)
	InsertMetadata(context.Context, *metadatamodels.InsertMetadataRequest) (*commonmodels.Empty, error)
	AddMetadataIndex(context.Context, *metadatamodels.AddMetadataIndexRequest) (*commonmodels.Empty, error)
	Query(context.Context, *metadatamodels.QueryRequest) (*metadatamodels.Field, error)
}

// UnimplementedMetadataCompositeStoreServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataCompositeStoreServer struct {
}

func (*UnimplementedMetadataCompositeStoreServer) InitMetadataDB(context.Context, *metadatamodels.InitMetadataDBRequest) (*metadatamodels.MetadataDBEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitMetadataDB not implemented")
}
func (*UnimplementedMetadataCompositeStoreServer) InitMetadataDBCollection(context.Context, *metadatamodels.InitMetadataDBRequest) (*metadatamodels.MetadataDBEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitMetadataDBCollection not implemented")
}
func (*UnimplementedMetadataCompositeStoreServer) InsertMetadata(context.Context, *metadatamodels.InsertMetadataRequest) (*commonmodels.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMetadata not implemented")
}
func (*UnimplementedMetadataCompositeStoreServer) AddMetadataIndex(context.Context, *metadatamodels.AddMetadataIndexRequest) (*commonmodels.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetadataIndex not implemented")
}
func (*UnimplementedMetadataCompositeStoreServer) Query(context.Context, *metadatamodels.QueryRequest) (*metadatamodels.Field, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterMetadataCompositeStoreServer(s *grpc.Server, srv MetadataCompositeStoreServer) {
	s.RegisterService(&_MetadataCompositeStore_serviceDesc, srv)
}

func _MetadataCompositeStore_InitMetadataDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatamodels.InitMetadataDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataCompositeStoreServer).InitMetadataDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetadataCompositeStore/InitMetadataDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataCompositeStoreServer).InitMetadataDB(ctx, req.(*metadatamodels.InitMetadataDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataCompositeStore_InitMetadataDBCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatamodels.InitMetadataDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataCompositeStoreServer).InitMetadataDBCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetadataCompositeStore/InitMetadataDBCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataCompositeStoreServer).InitMetadataDBCollection(ctx, req.(*metadatamodels.InitMetadataDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataCompositeStore_InsertMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatamodels.InsertMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataCompositeStoreServer).InsertMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetadataCompositeStore/InsertMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataCompositeStoreServer).InsertMetadata(ctx, req.(*metadatamodels.InsertMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataCompositeStore_AddMetadataIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatamodels.AddMetadataIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataCompositeStoreServer).AddMetadataIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetadataCompositeStore/AddMetadataIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataCompositeStoreServer).AddMetadataIndex(ctx, req.(*metadatamodels.AddMetadataIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataCompositeStore_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatamodels.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataCompositeStoreServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MetadataCompositeStore/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataCompositeStoreServer).Query(ctx, req.(*metadatamodels.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataCompositeStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MetadataCompositeStore",
	HandlerType: (*MetadataCompositeStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitMetadataDB",
			Handler:    _MetadataCompositeStore_InitMetadataDB_Handler,
		},
		{
			MethodName: "InitMetadataDBCollection",
			Handler:    _MetadataCompositeStore_InitMetadataDBCollection_Handler,
		},
		{
			MethodName: "InsertMetadata",
			Handler:    _MetadataCompositeStore_InsertMetadata_Handler,
		},
		{
			MethodName: "AddMetadataIndex",
			Handler:    _MetadataCompositeStore_AddMetadataIndex_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _MetadataCompositeStore_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/APIs.proto",
}
