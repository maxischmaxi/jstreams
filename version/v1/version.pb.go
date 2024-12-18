// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: version/v1/version.proto

package versionv1

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

type GetVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVersionsRequest) Reset() {
	*x = GetVersionsRequest{}
	mi := &file_version_v1_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsRequest) ProtoMessage() {}

func (x *GetVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsRequest.ProtoReflect.Descriptor instead.
func (*GetVersionsRequest) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{0}
}

type GetVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []string `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *GetVersionsResponse) Reset() {
	*x = GetVersionsResponse{}
	mi := &file_version_v1_version_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsResponse) ProtoMessage() {}

func (x *GetVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsResponse.ProtoReflect.Descriptor instead.
func (*GetVersionsResponse) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{1}
}

func (x *GetVersionsResponse) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

type GetCurrentVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentVersionRequest) Reset() {
	*x = GetCurrentVersionRequest{}
	mi := &file_version_v1_version_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentVersionRequest) ProtoMessage() {}

func (x *GetCurrentVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentVersionRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentVersionRequest) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{2}
}

type GetCurrentVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetCurrentVersionResponse) Reset() {
	*x = GetCurrentVersionResponse{}
	mi := &file_version_v1_version_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentVersionResponse) ProtoMessage() {}

func (x *GetCurrentVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_version_v1_version_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentVersionResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentVersionResponse) Descriptor() ([]byte, []int) {
	return file_version_v1_version_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrentVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_version_v1_version_proto protoreflect.FileDescriptor

var file_version_v1_version_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1a, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32,
	0xba, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29,
	0x6d, 0x61, 0x78, 0x69, 0x73, 0x63, 0x68, 0x6d, 0x61, 0x78, 0x69, 0x2f, 0x6a, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_version_v1_version_proto_rawDescOnce sync.Once
	file_version_v1_version_proto_rawDescData = file_version_v1_version_proto_rawDesc
)

func file_version_v1_version_proto_rawDescGZIP() []byte {
	file_version_v1_version_proto_rawDescOnce.Do(func() {
		file_version_v1_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_v1_version_proto_rawDescData)
	})
	return file_version_v1_version_proto_rawDescData
}

var file_version_v1_version_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_version_v1_version_proto_goTypes = []any{
	(*GetVersionsRequest)(nil),        // 0: version.GetVersionsRequest
	(*GetVersionsResponse)(nil),       // 1: version.GetVersionsResponse
	(*GetCurrentVersionRequest)(nil),  // 2: version.GetCurrentVersionRequest
	(*GetCurrentVersionResponse)(nil), // 3: version.GetCurrentVersionResponse
}
var file_version_v1_version_proto_depIdxs = []int32{
	0, // 0: version.VersionService.GetVersions:input_type -> version.GetVersionsRequest
	2, // 1: version.VersionService.GetCurrentVersion:input_type -> version.GetCurrentVersionRequest
	1, // 2: version.VersionService.GetVersions:output_type -> version.GetVersionsResponse
	3, // 3: version.VersionService.GetCurrentVersion:output_type -> version.GetCurrentVersionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_v1_version_proto_init() }
func file_version_v1_version_proto_init() {
	if File_version_v1_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_version_v1_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_version_v1_version_proto_goTypes,
		DependencyIndexes: file_version_v1_version_proto_depIdxs,
		MessageInfos:      file_version_v1_version_proto_msgTypes,
	}.Build()
	File_version_v1_version_proto = out.File
	file_version_v1_version_proto_rawDesc = nil
	file_version_v1_version_proto_goTypes = nil
	file_version_v1_version_proto_depIdxs = nil
}
