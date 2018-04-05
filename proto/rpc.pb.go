// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	GetValueRequest
	GetValueResponse
	SetValueRequest
	SetValueResponse
	GetKeysRequest
	GetKeysResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetValueRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetValueRequest) Reset()                    { *m = GetValueRequest{} }
func (m *GetValueRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetValueRequest) ProtoMessage()               {}
func (*GetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetValueResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetValueResponse) Reset()                    { *m = GetValueResponse{} }
func (m *GetValueResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetValueResponse) ProtoMessage()               {}
func (*GetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetValueResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetValueRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Ttl   string `protobuf:"bytes,3,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetValueRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *SetValueRequest) GetTtl() string {
	if m != nil {
		return m.Ttl
	}
	return ""
}

type SetValueResponse struct {
}

func (m *SetValueResponse) Reset()                    { *m = SetValueResponse{} }
func (m *SetValueResponse) String() string            { return proto1.CompactTextString(m) }
func (*SetValueResponse) ProtoMessage()               {}
func (*SetValueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetKeysRequest struct {
}

func (m *GetKeysRequest) Reset()                    { *m = GetKeysRequest{} }
func (m *GetKeysRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetKeysRequest) ProtoMessage()               {}
func (*GetKeysRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetKeysResponse struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetKeysResponse) Reset()                    { *m = GetKeysResponse{} }
func (m *GetKeysResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetKeysResponse) ProtoMessage()               {}
func (*GetKeysResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetKeysResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto1.RegisterType((*GetValueRequest)(nil), "proto.GetValueRequest")
	proto1.RegisterType((*GetValueResponse)(nil), "proto.GetValueResponse")
	proto1.RegisterType((*SetValueRequest)(nil), "proto.SetValueRequest")
	proto1.RegisterType((*SetValueResponse)(nil), "proto.SetValueResponse")
	proto1.RegisterType((*GetKeysRequest)(nil), "proto.GetKeysRequest")
	proto1.RegisterType((*GetKeysResponse)(nil), "proto.GetKeysResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
	GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (Storage_GetKeysClient, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	out := new(GetValueResponse)
	err := grpc.Invoke(ctx, "/proto.Storage/GetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := grpc.Invoke(ctx, "/proto.Storage/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetKeys(ctx context.Context, in *GetKeysRequest, opts ...grpc.CallOption) (Storage_GetKeysClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[0], c.cc, "/proto.Storage/GetKeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetKeysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetKeysClient interface {
	Recv() (*GetKeysResponse, error)
	grpc.ClientStream
}

type storageGetKeysClient struct {
	grpc.ClientStream
}

func (x *storageGetKeysClient) Recv() (*GetKeysResponse, error) {
	m := new(GetKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Storage service

type StorageServer interface {
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
	GetKeys(*GetKeysRequest, Storage_GetKeysServer) error
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Storage/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Storage/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetKeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetKeysRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetKeys(m, &storageGetKeysServer{stream})
}

type Storage_GetKeysServer interface {
	Send(*GetKeysResponse) error
	grpc.ServerStream
}

type storageGetKeysServer struct {
	grpc.ServerStream
}

func (x *storageGetKeysServer) Send(m *GetKeysResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _Storage_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _Storage_SetValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetKeys",
			Handler:       _Storage_GetKeys_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto1.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c, 0xfc, 0xee, 0xa9, 0x25,
	0x61, 0x89, 0x39, 0xa5, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x06, 0x97, 0x00,
	0x42, 0x51, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x19, 0x48, 0x00, 0xaa,
	0x0e, 0xc2, 0x51, 0xf2, 0xe6, 0xe2, 0x0f, 0x26, 0x64, 0x1c, 0x42, 0x2b, 0x13, 0x92, 0x56, 0x90,
	0xba, 0x92, 0x92, 0x1c, 0x09, 0x66, 0x88, 0xba, 0x92, 0x92, 0x1c, 0x25, 0x21, 0x2e, 0x81, 0x60,
	0x34, 0x6b, 0x95, 0x04, 0xb8, 0xf8, 0xdc, 0x53, 0x4b, 0xbc, 0x53, 0x2b, 0x8b, 0xa1, 0xe6, 0x43,
	0x7d, 0x00, 0x11, 0x81, 0xba, 0x0d, 0xc3, 0x4a, 0xa3, 0xa3, 0x8c, 0x5c, 0xec, 0xc1, 0x25, 0xf9,
	0x45, 0x89, 0xe9, 0xa9, 0x42, 0xb6, 0x5c, 0x1c, 0x30, 0xdf, 0x08, 0x89, 0x41, 0x42, 0x43, 0x0f,
	0x2d, 0x0c, 0xa4, 0xc4, 0x31, 0xc4, 0xa1, 0xf6, 0x33, 0x80, 0xb4, 0x07, 0xa3, 0x6b, 0x0f, 0xc6,
	0xa1, 0x3d, 0x18, 0x53, 0xbb, 0x0d, 0x17, 0x3b, 0xd4, 0xb9, 0x42, 0xa2, 0x08, 0x4b, 0x90, 0x3c,
	0x24, 0x25, 0x86, 0x2e, 0x0c, 0xd3, 0x6b, 0xc0, 0x98, 0xc4, 0x06, 0x96, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x73, 0x95, 0xa1, 0x11, 0xc9, 0x01, 0x00, 0x00,
}
