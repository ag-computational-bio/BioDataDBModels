// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: proto/CommonModels.proto

package commonmodels

import (
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

type Right int32

const (
	Right_Read  Right = 0
	Right_Write Right = 1
)

// Enum value maps for Right.
var (
	Right_name = map[int32]string{
		0: "Read",
		1: "Write",
	}
	Right_value = map[string]int32{
		"Read":  0,
		"Write": 1,
	}
)

func (x Right) Enum() *Right {
	p := new(Right)
	*p = x
	return p
}

func (x Right) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Right) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_CommonModels_proto_enumTypes[0].Descriptor()
}

func (Right) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[0]
}

func (x Right) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Right.Descriptor instead.
func (Right) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{0}
}

type Resource int32

const (
	Resource_Project                    Resource = 0
	Resource_Dataset                    Resource = 1
	Resource_DatasetVersion             Resource = 2
	Resource_DatasetObject              Resource = 3
	Resource_DatasetObjectGroupResource Resource = 4
)

// Enum value maps for Resource.
var (
	Resource_name = map[int32]string{
		0: "Project",
		1: "Dataset",
		2: "DatasetVersion",
		3: "DatasetObject",
		4: "DatasetObjectGroupResource",
	}
	Resource_value = map[string]int32{
		"Project":                    0,
		"Dataset":                    1,
		"DatasetVersion":             2,
		"DatasetObject":              3,
		"DatasetObjectGroupResource": 4,
	}
)

func (x Resource) Enum() *Resource {
	p := new(Resource)
	*p = x
	return p
}

func (x Resource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resource) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_CommonModels_proto_enumTypes[1].Descriptor()
}

func (Resource) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[1]
}

func (x Resource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resource.Descriptor instead.
func (Resource) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{1}
}

type Stage int32

const (
	Stage_Stable Stage = 0
	Stage_Beta   Stage = 1
	Stage_Alpha  Stage = 2
)

// Enum value maps for Stage.
var (
	Stage_name = map[int32]string{
		0: "Stable",
		1: "Beta",
		2: "Alpha",
	}
	Stage_value = map[string]int32{
		"Stable": 0,
		"Beta":   1,
		"Alpha":  2,
	}
)

func (x Stage) Enum() *Stage {
	p := new(Stage)
	*p = x
	return p
}

func (x Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_CommonModels_proto_enumTypes[2].Descriptor()
}

func (Stage) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[2]
}

func (x Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stage.Descriptor instead.
func (Stage) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{2}
}

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
	return file_proto_CommonModels_proto_enumTypes[3].Descriptor()
}

func (Origin_OriginTypeEnum) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[3]
}

func (x Origin_OriginTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Origin_OriginTypeEnum.Descriptor instead.
func (Origin_OriginTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{3, 0}
}

type Version_VersionStage int32

const (
	Version_Stable           Version_VersionStage = 0
	Version_ReleaseCandidate Version_VersionStage = 1
	Version_Beta             Version_VersionStage = 2
	Version_Alpha            Version_VersionStage = 3
)

// Enum value maps for Version_VersionStage.
var (
	Version_VersionStage_name = map[int32]string{
		0: "Stable",
		1: "ReleaseCandidate",
		2: "Beta",
		3: "Alpha",
	}
	Version_VersionStage_value = map[string]int32{
		"Stable":           0,
		"ReleaseCandidate": 1,
		"Beta":             2,
		"Alpha":            3,
	}
)

func (x Version_VersionStage) Enum() *Version_VersionStage {
	p := new(Version_VersionStage)
	*p = x
	return p
}

func (x Version_VersionStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_VersionStage) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_CommonModels_proto_enumTypes[4].Descriptor()
}

func (Version_VersionStage) Type() protoreflect.EnumType {
	return &file_proto_CommonModels_proto_enumTypes[4]
}

func (x Version_VersionStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_VersionStage.Descriptor instead.
func (Version_VersionStage) EnumDescriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{5, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Rights   []Right  `protobuf:"varint,2,rep,packed,name=Rights,proto3,enum=Right" json:"Rights,omitempty"`
	Resource Resource `protobuf:"varint,3,opt,name=Resource,proto3,enum=Resource" json:"Resource,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetRights() []Right {
	if x != nil {
		return x.Rights
	}
	return nil
}

func (x *User) GetResource() Resource {
	if x != nil {
		return x.Resource
	}
	return Resource_Project
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{1}
}

// A location in S3
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string `protobuf:"bytes,1,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	URL    string `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"` // Object storage endpoint
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Location) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Location) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link                 string                `protobuf:"bytes,1,opt,name=Link,proto3" json:"Link,omitempty"`
	ObjectStorageLocatio *Location             `protobuf:"bytes,2,opt,name=ObjectStorageLocatio,proto3" json:"ObjectStorageLocatio,omitempty"`
	OriginType           Origin_OriginTypeEnum `protobuf:"varint,4,opt,name=OriginType,proto3,enum=Origin_OriginTypeEnum" json:"OriginType,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[3]
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
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{3}
}

func (x *Origin) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Origin) GetObjectStorageLocatio() *Location {
	if x != nil {
		return x.ObjectStorageLocatio
	}
	return nil
}

func (x *Origin) GetOriginType() Origin_OriginTypeEnum {
	if x != nil {
		return x.OriginType
	}
	return Origin_ObjectStorage
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"` // An arbitrary ID
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[4]
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
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{4}
}

func (x *ID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major    int32                `protobuf:"varint,1,opt,name=Major,proto3" json:"Major,omitempty"`
	Minor    int32                `protobuf:"varint,2,opt,name=Minor,proto3" json:"Minor,omitempty"`
	Patch    int32                `protobuf:"varint,3,opt,name=Patch,proto3" json:"Patch,omitempty"`
	Revision int32                `protobuf:"varint,4,opt,name=Revision,proto3" json:"Revision,omitempty"`
	Stage    Version_VersionStage `protobuf:"varint,5,opt,name=Stage,proto3,enum=Version_VersionStage" json:"Stage,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{5}
}

func (x *Version) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() int32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *Version) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *Version) GetStage() Version_VersionStage {
	if x != nil {
		return x.Stage
	}
	return Version_Stable
}

type Int64Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Int64Wrapper) Reset() {
	*x = Int64Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Wrapper) ProtoMessage() {}

func (x *Int64Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Wrapper.ProtoReflect.Descriptor instead.
func (*Int64Wrapper) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{6}
}

func (x *Int64Wrapper) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type IDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []string `protobuf:"bytes,1,rep,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *IDList) Reset() {
	*x = IDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CommonModels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDList) ProtoMessage() {}

func (x *IDList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CommonModels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDList.ProtoReflect.Descriptor instead.
func (*IDList) Descriptor() ([]byte, []int) {
	return file_proto_CommonModels_proto_rawDescGZIP(), []int{7}
}

func (x *IDList) GetIDs() []string {
	if x != nil {
		return x.IDs
	}
	return nil
}

var File_proto_CommonModels_proto protoreflect.FileDescriptor

var file_proto_CommonModels_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x06, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x3a, 0x1d, 0x92, 0x41, 0x1a, 0x0a, 0x18, 0x2a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x0f, 0x41, 0x6e, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x20, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x78, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x3a, 0x30, 0x92, 0x41, 0x2d,
	0x0a, 0x2b, 0x2a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x1f, 0x41, 0x20,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xec, 0x03,
	0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3d, 0x0a, 0x14,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x0e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x10, 0x01, 0x3a, 0xa1, 0x02, 0x92, 0x41, 0x9d, 0x02, 0x0a,
	0xc6, 0x01, 0x2a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x32, 0xbb, 0x01, 0x54, 0x68, 0x65,
	0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x2e, 0x20, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x20, 0x61, 0x72,
	0x65, 0x20, 0x73, 0x65, 0x74, 0x20, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x20, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x20, 0x20, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x20,
	0x6d, 0x65, 0x61, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x6e,
	0x6b, 0x20, 0x69, 0x73, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x2c, 0x20, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2c, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6c,
	0x69, 0x6e, 0x6b, 0x20, 0x61, 0x72, 0x65, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x2e,
	0x20, 0x4c, 0x69, 0x6e, 0x6b, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x33, 0x20,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0x52, 0x12, 0x50, 0x7b, 0x20, 0x22, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x61, 0x77, 0x73, 0x6f, 0x6d, 0x65, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x62, 0x66, 0x66, 0x22, 0x2c, 0x20, 0x22, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x3a, 0x20, 0x22, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x20, 0x7d, 0x22, 0x30, 0x0a, 0x02,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15, 0x2a, 0x02, 0x49, 0x44, 0x32, 0x0f, 0x41,
	0x6e, 0x20, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72, 0x61, 0x72, 0x79, 0x20, 0x49, 0x44, 0x22, 0xef,
	0x02, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61,
	0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x65, 0x74, 0x61, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x10, 0x03, 0x3a, 0x91, 0x01, 0x92,
	0x41, 0x8d, 0x01, 0x0a, 0x40, 0x2a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x35,
	0x41, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x62, 0x61, 0x73, 0x65, 0x64, 0x20, 0x6f,
	0x6e, 0x20, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x32, 0x49, 0x12, 0x47, 0x7b, 0x20, 0x22, 0x4d, 0x61, 0x6a, 0x6f,
	0x72, 0x22, 0x3a, 0x20, 0x30, 0x2c, 0x20, 0x22, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x22, 0x3a, 0x20,
	0x31, 0x2c, 0x20, 0x22, 0x50, 0x61, 0x74, 0x63, 0x68, 0x22, 0x3a, 0x20, 0x30, 0x2c, 0x20, 0x22,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x22, 0x20, 0x7d,
	0x22, 0x51, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x2b, 0x92, 0x41, 0x28, 0x0a, 0x26, 0x2a, 0x0c, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x32, 0x16, 0x41, 0x20, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x22, 0x5a, 0x0a, 0x06, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x49, 0x44, 0x73, 0x3a,
	0x3e, 0x92, 0x41, 0x3b, 0x0a, 0x1a, 0x2a, 0x07, 0x49, 0x44, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x32,
	0x0f, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x64, 0x65, 0x61, 0x73,
	0x32, 0x1d, 0x12, 0x1b, 0x7b, 0x20, 0x22, 0x49, 0x44, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x31,
	0x22, 0x2c, 0x22, 0x32, 0x22, 0x2c, 0x22, 0x33, 0x22, 0x2c, 0x22, 0x34, 0x22, 0x5d, 0x7d, 0x2a,
	0x1c, 0x0a, 0x05, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x10, 0x01, 0x2a, 0x6b, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x04, 0x2a, 0x28, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x42, 0x65, 0x74, 0x61, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x10, 0x02, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x67, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2d, 0x62, 0x69, 0x6f, 0x2f, 0x42, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x44,
	0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_CommonModels_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_proto_CommonModels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_CommonModels_proto_goTypes = []interface{}{
	(Right)(0),                 // 0: Right
	(Resource)(0),              // 1: Resource
	(Stage)(0),                 // 2: Stage
	(Origin_OriginTypeEnum)(0), // 3: Origin.OriginTypeEnum
	(Version_VersionStage)(0),  // 4: Version.VersionStage
	(*User)(nil),               // 5: User
	(*Empty)(nil),              // 6: Empty
	(*Location)(nil),           // 7: Location
	(*Origin)(nil),             // 8: Origin
	(*ID)(nil),                 // 9: ID
	(*Version)(nil),            // 10: Version
	(*Int64Wrapper)(nil),       // 11: Int64Wrapper
	(*IDList)(nil),             // 12: IDList
}
var file_proto_CommonModels_proto_depIdxs = []int32{
	0, // 0: User.Rights:type_name -> Right
	1, // 1: User.Resource:type_name -> Resource
	7, // 2: Origin.ObjectStorageLocatio:type_name -> Location
	3, // 3: Origin.OriginType:type_name -> Origin.OriginTypeEnum
	4, // 4: Version.Stage:type_name -> Version.VersionStage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_CommonModels_proto_init() }
func file_proto_CommonModels_proto_init() {
	if File_proto_CommonModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_CommonModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_CommonModels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_proto_CommonModels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_CommonModels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_CommonModels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_proto_CommonModels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Wrapper); i {
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
		file_proto_CommonModels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDList); i {
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
			NumEnums:      5,
			NumMessages:   8,
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
