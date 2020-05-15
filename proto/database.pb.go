// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.7.1
// source: database.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OperationSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OperationSetRequest) Reset() {
	*x = OperationSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationSetRequest) ProtoMessage() {}

func (x *OperationSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationSetRequest.ProtoReflect.Descriptor instead.
func (*OperationSetRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{0}
}

func (x *OperationSetRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *OperationSetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type OperationSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *OperationSetResponse) Reset() {
	*x = OperationSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationSetResponse) ProtoMessage() {}

func (x *OperationSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationSetResponse.ProtoReflect.Descriptor instead.
func (*OperationSetResponse) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{1}
}

func (x *OperationSetResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *OperationSetResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type OperationGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *OperationGetRequest) Reset() {
	*x = OperationGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationGetRequest) ProtoMessage() {}

func (x *OperationGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationGetRequest.ProtoReflect.Descriptor instead.
func (*OperationGetRequest) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{2}
}

func (x *OperationGetRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type OperationGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok           bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Value        []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OperationGetResponse) Reset() {
	*x = OperationGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationGetResponse) ProtoMessage() {}

func (x *OperationGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationGetResponse.ProtoReflect.Descriptor instead.
func (*OperationGetResponse) Descriptor() ([]byte, []int) {
	return file_database_proto_rawDescGZIP(), []int{3}
}

func (x *OperationGetResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *OperationGetResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *OperationGetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_database_proto protoreflect.FileDescriptor

var file_database_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x4a, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x13, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x60, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x72, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_proto_rawDescOnce sync.Once
	file_database_proto_rawDescData = file_database_proto_rawDesc
)

func file_database_proto_rawDescGZIP() []byte {
	file_database_proto_rawDescOnce.Do(func() {
		file_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_proto_rawDescData)
	})
	return file_database_proto_rawDescData
}

var file_database_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_database_proto_goTypes = []interface{}{
	(*OperationSetRequest)(nil),  // 0: OperationSetRequest
	(*OperationSetResponse)(nil), // 1: OperationSetResponse
	(*OperationGetRequest)(nil),  // 2: OperationGetRequest
	(*OperationGetResponse)(nil), // 3: OperationGetResponse
}
var file_database_proto_depIdxs = []int32{
	0, // 0: Database.Set:input_type -> OperationSetRequest
	2, // 1: Database.Get:input_type -> OperationGetRequest
	1, // 2: Database.Set:output_type -> OperationSetResponse
	3, // 3: Database.Get:output_type -> OperationGetResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_database_proto_init() }
func file_database_proto_init() {
	if File_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationSetRequest); i {
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
		file_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationSetResponse); i {
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
		file_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationGetRequest); i {
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
		file_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationGetResponse); i {
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
			RawDescriptor: file_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_database_proto_goTypes,
		DependencyIndexes: file_database_proto_depIdxs,
		MessageInfos:      file_database_proto_msgTypes,
	}.Build()
	File_database_proto = out.File
	file_database_proto_rawDesc = nil
	file_database_proto_goTypes = nil
	file_database_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseClient interface {
	Set(ctx context.Context, in *OperationSetRequest, opts ...grpc.CallOption) (*OperationSetResponse, error)
	Get(ctx context.Context, in *OperationGetRequest, opts ...grpc.CallOption) (*OperationGetResponse, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Set(ctx context.Context, in *OperationSetRequest, opts ...grpc.CallOption) (*OperationSetResponse, error) {
	out := new(OperationSetResponse)
	err := c.cc.Invoke(ctx, "/Database/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Get(ctx context.Context, in *OperationGetRequest, opts ...grpc.CallOption) (*OperationGetResponse, error) {
	out := new(OperationGetResponse)
	err := c.cc.Invoke(ctx, "/Database/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
type DatabaseServer interface {
	Set(context.Context, *OperationSetRequest) (*OperationSetResponse, error)
	Get(context.Context, *OperationGetRequest) (*OperationGetResponse, error)
}

// UnimplementedDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedDatabaseServer struct {
}

func (*UnimplementedDatabaseServer) Set(context.Context, *OperationSetRequest) (*OperationSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedDatabaseServer) Get(context.Context, *OperationGetRequest) (*OperationGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Set(ctx, req.(*OperationSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*OperationGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Database_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database.proto",
}
