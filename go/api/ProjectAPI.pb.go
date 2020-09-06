// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/ProjectAPI.proto

package api

import (
	context "context"
	commonmodels "github.com/ag-computational-bio/BioDataDBModels/go/commonmodels"
	datasetmodels "github.com/ag-computational-bio/BioDataDBModels/go/datasetmodels"
	projectmodels "github.com/ag-computational-bio/BioDataDBModels/go/projectmodels"
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

var File_proto_ProjectAPI_proto protoreflect.FileDescriptor

var file_proto_ProjectAPI_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x41, 0x50, 0x49, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb5, 0x03, 0x0a, 0x0a, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x50, 0x49, 0x12, 0x5f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x61, 0x64, 0x64, 0x75, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0c,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x3f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x42, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_ProjectAPI_proto_goTypes = []interface{}{
	(*projectmodels.CreateProjectRequest)(nil),    // 0: CreateProjectRequest
	(*projectmodels.AddUserToProjectRequest)(nil), // 1: AddUserToProjectRequest
	(*commonmodels.ID)(nil),                       // 2: ID
	(*commonmodels.Empty)(nil),                    // 3: Empty
	(*projectmodels.ProjectEntry)(nil),            // 4: ProjectEntry
	(*datasetmodels.DatasetList)(nil),             // 5: DatasetList
	(*projectmodels.ProjectEntryList)(nil),        // 6: ProjectEntryList
}
var file_proto_ProjectAPI_proto_depIdxs = []int32{
	0, // 0: ProjectAPI.CreateProject:input_type -> CreateProjectRequest
	1, // 1: ProjectAPI.AddUserToProject:input_type -> AddUserToProjectRequest
	2, // 2: ProjectAPI.GetProjectDatasets:input_type -> ID
	3, // 3: ProjectAPI.GetUserProjects:input_type -> Empty
	2, // 4: ProjectAPI.DeleteProject:input_type -> ID
	4, // 5: ProjectAPI.CreateProject:output_type -> ProjectEntry
	4, // 6: ProjectAPI.AddUserToProject:output_type -> ProjectEntry
	5, // 7: ProjectAPI.GetProjectDatasets:output_type -> DatasetList
	6, // 8: ProjectAPI.GetUserProjects:output_type -> ProjectEntryList
	3, // 9: ProjectAPI.DeleteProject:output_type -> Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ProjectAPI_proto_init() }
func file_proto_ProjectAPI_proto_init() {
	if File_proto_ProjectAPI_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ProjectAPI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ProjectAPI_proto_goTypes,
		DependencyIndexes: file_proto_ProjectAPI_proto_depIdxs,
	}.Build()
	File_proto_ProjectAPI_proto = out.File
	file_proto_ProjectAPI_proto_rawDesc = nil
	file_proto_ProjectAPI_proto_goTypes = nil
	file_proto_ProjectAPI_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectAPIClient is the client API for ProjectAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectAPIClient interface {
	CreateProject(ctx context.Context, in *projectmodels.CreateProjectRequest, opts ...grpc.CallOption) (*projectmodels.ProjectEntry, error)
	AddUserToProject(ctx context.Context, in *projectmodels.AddUserToProjectRequest, opts ...grpc.CallOption) (*projectmodels.ProjectEntry, error)
	GetProjectDatasets(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetList, error)
	GetUserProjects(ctx context.Context, in *commonmodels.Empty, opts ...grpc.CallOption) (*projectmodels.ProjectEntryList, error)
	DeleteProject(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*commonmodels.Empty, error)
}

type projectAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAPIClient(cc grpc.ClientConnInterface) ProjectAPIClient {
	return &projectAPIClient{cc}
}

func (c *projectAPIClient) CreateProject(ctx context.Context, in *projectmodels.CreateProjectRequest, opts ...grpc.CallOption) (*projectmodels.ProjectEntry, error) {
	out := new(projectmodels.ProjectEntry)
	err := c.cc.Invoke(ctx, "/ProjectAPI/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) AddUserToProject(ctx context.Context, in *projectmodels.AddUserToProjectRequest, opts ...grpc.CallOption) (*projectmodels.ProjectEntry, error) {
	out := new(projectmodels.ProjectEntry)
	err := c.cc.Invoke(ctx, "/ProjectAPI/AddUserToProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetProjectDatasets(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*datasetmodels.DatasetList, error) {
	out := new(datasetmodels.DatasetList)
	err := c.cc.Invoke(ctx, "/ProjectAPI/GetProjectDatasets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) GetUserProjects(ctx context.Context, in *commonmodels.Empty, opts ...grpc.CallOption) (*projectmodels.ProjectEntryList, error) {
	out := new(projectmodels.ProjectEntryList)
	err := c.cc.Invoke(ctx, "/ProjectAPI/GetUserProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAPIClient) DeleteProject(ctx context.Context, in *commonmodels.ID, opts ...grpc.CallOption) (*commonmodels.Empty, error) {
	out := new(commonmodels.Empty)
	err := c.cc.Invoke(ctx, "/ProjectAPI/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAPIServer is the server API for ProjectAPI service.
type ProjectAPIServer interface {
	CreateProject(context.Context, *projectmodels.CreateProjectRequest) (*projectmodels.ProjectEntry, error)
	AddUserToProject(context.Context, *projectmodels.AddUserToProjectRequest) (*projectmodels.ProjectEntry, error)
	GetProjectDatasets(context.Context, *commonmodels.ID) (*datasetmodels.DatasetList, error)
	GetUserProjects(context.Context, *commonmodels.Empty) (*projectmodels.ProjectEntryList, error)
	DeleteProject(context.Context, *commonmodels.ID) (*commonmodels.Empty, error)
}

// UnimplementedProjectAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProjectAPIServer struct {
}

func (*UnimplementedProjectAPIServer) CreateProject(context.Context, *projectmodels.CreateProjectRequest) (*projectmodels.ProjectEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (*UnimplementedProjectAPIServer) AddUserToProject(context.Context, *projectmodels.AddUserToProjectRequest) (*projectmodels.ProjectEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToProject not implemented")
}
func (*UnimplementedProjectAPIServer) GetProjectDatasets(context.Context, *commonmodels.ID) (*datasetmodels.DatasetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectDatasets not implemented")
}
func (*UnimplementedProjectAPIServer) GetUserProjects(context.Context, *commonmodels.Empty) (*projectmodels.ProjectEntryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProjects not implemented")
}
func (*UnimplementedProjectAPIServer) DeleteProject(context.Context, *commonmodels.ID) (*commonmodels.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}

func RegisterProjectAPIServer(s *grpc.Server, srv ProjectAPIServer) {
	s.RegisterService(&_ProjectAPI_serviceDesc, srv)
}

func _ProjectAPI_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(projectmodels.CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectAPI/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).CreateProject(ctx, req.(*projectmodels.CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_AddUserToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(projectmodels.AddUserToProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).AddUserToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectAPI/AddUserToProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).AddUserToProject(ctx, req.(*projectmodels.AddUserToProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetProjectDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetProjectDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectAPI/GetProjectDatasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetProjectDatasets(ctx, req.(*commonmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_GetUserProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).GetUserProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectAPI/GetUserProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).GetUserProjects(ctx, req.(*commonmodels.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAPI_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonmodels.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAPIServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectAPI/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAPIServer).DeleteProject(ctx, req.(*commonmodels.ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProjectAPI",
	HandlerType: (*ProjectAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectAPI_CreateProject_Handler,
		},
		{
			MethodName: "AddUserToProject",
			Handler:    _ProjectAPI_AddUserToProject_Handler,
		},
		{
			MethodName: "GetProjectDatasets",
			Handler:    _ProjectAPI_GetProjectDatasets_Handler,
		},
		{
			MethodName: "GetUserProjects",
			Handler:    _ProjectAPI_GetUserProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectAPI_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ProjectAPI.proto",
}
