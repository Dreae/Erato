// Code generated by protoc-gen-go.
// source: protobuf/register.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	protobuf/register.proto

It has these top-level messages:
	RegisterRequest
	RegisterResult
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type RegisterRequest struct {
	ApiKey string `protobuf:"bytes,1,opt,name=apiKey" json:"apiKey,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterResult struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RegisterResult) Reset()                    { *m = RegisterResult{} }
func (m *RegisterResult) String() string            { return proto.CompactTextString(m) }
func (*RegisterResult) ProtoMessage()               {}
func (*RegisterResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "protobuf.RegisterRequest")
	proto.RegisterType((*RegisterResult)(nil), "protobuf.RegisterResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Master service

type MasterClient interface {
	DoRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResult, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) DoRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResult, error) {
	out := new(RegisterResult)
	err := grpc.Invoke(ctx, "/protobuf.Master/DoRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	DoRegister(context.Context, *RegisterRequest) (*RegisterResult, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_DoRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DoRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Master/DoRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DoRegister(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRegister",
			Handler:    _Master_DoRegister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x03, 0x8b,
	0x08, 0x71, 0xc0, 0x24, 0x94, 0x34, 0xb9, 0xf8, 0x83, 0xa0, 0x72, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0x05, 0x99, 0xde, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x16, 0x17, 0x1f, 0x42, 0x69, 0x71, 0x69, 0x4e, 0x89, 0x90,
	0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0x2a, 0x54, 0x29, 0x8c, 0x6b, 0xe4, 0xcb,
	0xc5, 0xe6, 0x9b, 0x08, 0x52, 0x29, 0xe4, 0xcc, 0xc5, 0xe5, 0x92, 0x0f, 0xd3, 0x27, 0x24, 0xa9,
	0x07, 0xb3, 0x59, 0x0f, 0xcd, 0x5a, 0x29, 0x09, 0x6c, 0x52, 0x20, 0x6b, 0x94, 0x18, 0x92, 0xd8,
	0xc0, 0x52, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x07, 0x80, 0x0f, 0xd1, 0x00, 0x00,
	0x00,
}