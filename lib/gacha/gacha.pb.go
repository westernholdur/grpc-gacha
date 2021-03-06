// Code generated by protoc-gen-go.
// source: gacha.proto
// DO NOT EDIT!

/*
Package gacha is a generated protocol buffer package.

It is generated from these files:
	gacha.proto

It has these top-level messages:
	Card
	Request
	Response
*/
package gacha

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Card struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	Cards []*Card `protobuf:"bytes,1,rep,name=cards" json:"cards,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

type Response struct {
	Card    *Card `protobuf:"bytes,1,opt,name=card" json:"card,omitempty"`
	RetCode int32 `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func init() {
	proto.RegisterType((*Card)(nil), "gacha.Card")
	proto.RegisterType((*Request)(nil), "gacha.Request")
	proto.RegisterType((*Response)(nil), "gacha.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Gacha service

type GachaClient interface {
	Lottery(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type gachaClient struct {
	cc *grpc.ClientConn
}

func NewGachaClient(cc *grpc.ClientConn) GachaClient {
	return &gachaClient{cc}
}

func (c *gachaClient) Lottery(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gacha.Gacha/Lottery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gacha service

type GachaServer interface {
	Lottery(context.Context, *Request) (*Response, error)
}

func RegisterGachaServer(s *grpc.Server, srv GachaServer) {
	s.RegisterService(&_Gacha_serviceDesc, srv)
}

func _Gacha_Lottery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GachaServer).Lottery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gacha.Gacha/Lottery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GachaServer).Lottery(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gacha_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gacha.Gacha",
	HandlerType: (*GachaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lottery",
			Handler:    _Gacha_Lottery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("gacha.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4f, 0x4c, 0xce,
	0x48, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xa4, 0xb8, 0x58, 0x9c,
	0x13, 0x8b, 0x52, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x1d, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x45, 0x2e, 0xd6, 0x64, 0xa0, 0xb2, 0x62, 0xa0, 0x3c, 0xb3, 0x06, 0xb7, 0x11, 0xb7, 0x1e, 0xc4,
	0x28, 0x90, 0xd6, 0x20, 0x88, 0x8c, 0x92, 0x1b, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x3c, 0x17, 0x0b, 0x48, 0x10, 0x6c, 0x1a, 0x9a, 0x6a, 0xb0, 0x84, 0x90, 0x24,
	0x17, 0x47, 0x51, 0x6a, 0x49, 0x7c, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0x13, 0x50, 0x11, 0x6b, 0x10,
	0x3b, 0x90, 0xef, 0x0c, 0xe4, 0x1a, 0x99, 0x72, 0xb1, 0xba, 0x83, 0x94, 0x0b, 0x01, 0xad, 0xf7,
	0xc9, 0x2f, 0x29, 0x49, 0x2d, 0xaa, 0x14, 0xe2, 0x83, 0x9a, 0x00, 0x75, 0x8e, 0x14, 0x3f, 0x9c,
	0x0f, 0xb1, 0x50, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x2d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x4f, 0x6d, 0x85, 0xe5, 0x00, 0x00, 0x00,
}
