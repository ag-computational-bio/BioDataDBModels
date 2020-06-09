// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/APIs.proto

package api

import (
	context "context"
	loadmodels "github.com/ag-computational-bio/BioDataDBModels/go/loadmodels"
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
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe0, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74,
	0x4c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x69, 0x6e,
	0x69, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x35, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4c, 0x6f,
	0x61, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42,
	0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_APIs_proto_goTypes = []interface{}{
	(*loadmodels.InitLoadDataset)(nil),          // 0: InitLoadDataset
	(*loadmodels.CreateLoadLinkSetRequest)(nil), // 1: CreateLoadLinkSetRequest
	(*loadmodels.ID)(nil),                       // 2: ID
	(*loadmodels.UploadLinks)(nil),              // 3: UploadLinks
	(*loadmodels.Empty)(nil),                    // 4: Empty
}
var file_proto_APIs_proto_depIdxs = []int32{
	0, // 0: LoadService.InitLoad:input_type -> InitLoadDataset
	1, // 1: LoadService.CreateLoadLinkSet:input_type -> CreateLoadLinkSetRequest
	2, // 2: LoadService.FinishLoad:input_type -> ID
	2, // 3: LoadService.InitLoad:output_type -> ID
	3, // 4: LoadService.CreateLoadLinkSet:output_type -> UploadLinks
	4, // 5: LoadService.FinishLoad:output_type -> Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumServices:   1,
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
	InitLoad(ctx context.Context, in *loadmodels.InitLoadDataset, opts ...grpc.CallOption) (*loadmodels.ID, error)
	CreateLoadLinkSet(ctx context.Context, in *loadmodels.CreateLoadLinkSetRequest, opts ...grpc.CallOption) (*loadmodels.UploadLinks, error)
	FinishLoad(ctx context.Context, in *loadmodels.ID, opts ...grpc.CallOption) (*loadmodels.Empty, error)
}

type loadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadServiceClient(cc grpc.ClientConnInterface) LoadServiceClient {
	return &loadServiceClient{cc}
}

func (c *loadServiceClient) InitLoad(ctx context.Context, in *loadmodels.InitLoadDataset, opts ...grpc.CallOption) (*loadmodels.ID, error) {
	out := new(loadmodels.ID)
	err := c.cc.Invoke(ctx, "/LoadService/InitLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadServiceClient) CreateLoadLinkSet(ctx context.Context, in *loadmodels.CreateLoadLinkSetRequest, opts ...grpc.CallOption) (*loadmodels.UploadLinks, error) {
	out := new(loadmodels.UploadLinks)
	err := c.cc.Invoke(ctx, "/LoadService/CreateLoadLinkSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadServiceClient) FinishLoad(ctx context.Context, in *loadmodels.ID, opts ...grpc.CallOption) (*loadmodels.Empty, error) {
	out := new(loadmodels.Empty)
	err := c.cc.Invoke(ctx, "/LoadService/FinishLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServiceServer is the server API for LoadService service.
type LoadServiceServer interface {
	InitLoad(context.Context, *loadmodels.InitLoadDataset) (*loadmodels.ID, error)
	CreateLoadLinkSet(context.Context, *loadmodels.CreateLoadLinkSetRequest) (*loadmodels.UploadLinks, error)
	FinishLoad(context.Context, *loadmodels.ID) (*loadmodels.Empty, error)
}

// UnimplementedLoadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadServiceServer struct {
}

func (*UnimplementedLoadServiceServer) InitLoad(context.Context, *loadmodels.InitLoadDataset) (*loadmodels.ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitLoad not implemented")
}
func (*UnimplementedLoadServiceServer) CreateLoadLinkSet(context.Context, *loadmodels.CreateLoadLinkSetRequest) (*loadmodels.UploadLinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoadLinkSet not implemented")
}
func (*UnimplementedLoadServiceServer) FinishLoad(context.Context, *loadmodels.ID) (*loadmodels.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishLoad not implemented")
}

func RegisterLoadServiceServer(s *grpc.Server, srv LoadServiceServer) {
	s.RegisterService(&_LoadService_serviceDesc, srv)
}

func _LoadService_InitLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(loadmodels.InitLoadDataset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServiceServer).InitLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoadService/InitLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServiceServer).InitLoad(ctx, req.(*loadmodels.InitLoadDataset))
	}
	return interceptor(ctx, in, info, handler)
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

func _LoadService_FinishLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(loadmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServiceServer).FinishLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoadService/FinishLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServiceServer).FinishLoad(ctx, req.(*loadmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LoadService",
	HandlerType: (*LoadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitLoad",
			Handler:    _LoadService_InitLoad_Handler,
		},
		{
			MethodName: "CreateLoadLinkSet",
			Handler:    _LoadService_CreateLoadLinkSet_Handler,
		},
		{
			MethodName: "FinishLoad",
			Handler:    _LoadService_FinishLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/APIs.proto",
}
