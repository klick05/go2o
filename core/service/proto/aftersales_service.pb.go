// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aftersales_service.proto

package proto // import "."

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AfterSalesServiceClient is the client API for AfterSalesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AfterSalesServiceClient interface {
}

type afterSalesServiceClient struct {
	cc *grpc.ClientConn
}

func NewAfterSalesServiceClient(cc *grpc.ClientConn) AfterSalesServiceClient {
	return &afterSalesServiceClient{cc}
}

// AfterSalesServiceServer is the server API for AfterSalesService service.
type AfterSalesServiceServer interface {
}

func RegisterAfterSalesServiceServer(s *grpc.Server, srv AfterSalesServiceServer) {
	s.RegisterService(&_AfterSalesService_serviceDesc, srv)
}

var _AfterSalesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AfterSalesService",
	HandlerType: (*AfterSalesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "aftersales_service.proto",
}

func init() {
	proto.RegisterFile("aftersales_service.proto", fileDescriptor_aftersales_service_8a02d819d1bf34e3)
}

var fileDescriptor_aftersales_service_8a02d819d1bf34e3 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4c, 0x2b, 0x49,
	0x2d, 0x2a, 0x4e, 0xcc, 0x49, 0x2d, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x49, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x81, 0xf2, 0x24, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xf5, 0xc1, 0xea, 0xe3, 0xc1, 0x1a, 0x20, 0x52, 0x46, 0xc2,
	0x5c, 0x82, 0x8e, 0x20, 0xc1, 0x60, 0x90, 0x58, 0x30, 0xc4, 0x0c, 0x27, 0xce, 0x28, 0x76, 0x3d,
	0x6b, 0xb0, 0x7c, 0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x65, 0x0d, 0x79,
	0x0c, 0x6b, 0x00, 0x00, 0x00,
}