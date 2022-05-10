// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: distributor.proto

package distributorpb

import (
	context "context"
	fmt "fmt"
	cortexpb "github.com/muhammadn/cortex/pkg/cortexpb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("distributor.proto", fileDescriptor_c518e33639ca565d) }

var fileDescriptor_c518e33639ca565d = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x57,
	0xca, 0x12, 0x49, 0x79, 0x72, 0x7e, 0x51, 0x49, 0x6a, 0x45, 0x41, 0x51, 0x7e, 0x56, 0x6a, 0x72,
	0x09, 0x94, 0xa7, 0x5f, 0x90, 0x9d, 0x0e, 0x93, 0x48, 0x82, 0x32, 0x20, 0x5a, 0x8d, 0x3c, 0xb8,
	0xb8, 0x5d, 0x10, 0x16, 0x0b, 0x59, 0x72, 0xb1, 0x04, 0x94, 0x16, 0x67, 0x08, 0x89, 0xe9, 0xc1,
	0x94, 0xeb, 0x85, 0x17, 0x65, 0x96, 0xa4, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x89,
	0x63, 0x88, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x38, 0x39, 0x5f, 0x78, 0x28, 0xc7,
	0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92,
	0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f, 0x3c,
	0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0xe2, 0x45, 0xf2, 0x76, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x55, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xcd, 0x1b, 0x85, 0x21, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistributorClient is the client API for Distributor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistributorClient interface {
	Push(ctx context.Context, in *cortexpb.WriteRequest, opts ...grpc.CallOption) (*cortexpb.WriteResponse, error)
}

type distributorClient struct {
	cc *grpc.ClientConn
}

func NewDistributorClient(cc *grpc.ClientConn) DistributorClient {
	return &distributorClient{cc}
}

func (c *distributorClient) Push(ctx context.Context, in *cortexpb.WriteRequest, opts ...grpc.CallOption) (*cortexpb.WriteResponse, error) {
	out := new(cortexpb.WriteResponse)
	err := c.cc.Invoke(ctx, "/distributor.Distributor/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributorServer is the server API for Distributor service.
type DistributorServer interface {
	Push(context.Context, *cortexpb.WriteRequest) (*cortexpb.WriteResponse, error)
}

// UnimplementedDistributorServer can be embedded to have forward compatible implementations.
type UnimplementedDistributorServer struct {
}

func (*UnimplementedDistributorServer) Push(ctx context.Context, req *cortexpb.WriteRequest) (*cortexpb.WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterDistributorServer(s *grpc.Server, srv DistributorServer) {
	s.RegisterService(&_Distributor_serviceDesc, srv)
}

func _Distributor_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cortexpb.WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributorServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distributor.Distributor/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributorServer).Push(ctx, req.(*cortexpb.WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Distributor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "distributor.Distributor",
	HandlerType: (*DistributorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Distributor_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distributor.proto",
}
