// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ydsxiong/summingservice/gRPC/proto-files/service/numbers-service.proto

package service

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	"github.com/ydsxiong/summingservice/gRPC/domain"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/ydsxiong/summingservice/gRPC/proto-files/service/numbers-service.proto", fileDescriptor_258ccc95220a90d6)
}

var fileDescriptor_258ccc95220a90d6 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xbd, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0x11, 0x41, 0x21, 0x9d, 0xb1, 0xdb, 0x5a, 0xec, 0x36, 0x11, 0x05, 0x7b, 0x15, 0xac,
	0x2c, 0x64, 0x83, 0x8d, 0xdd, 0x7e, 0x8c, 0x71, 0x20, 0x93, 0xac, 0x99, 0x8c, 0xdc, 0xfd, 0xf7,
	0x07, 0xfb, 0x01, 0x77, 0x70, 0xd7, 0x5c, 0xf9, 0x92, 0xdf, 0xfb, 0xf1, 0x18, 0xf5, 0xe5, 0xb1,
	0xfc, 0x4a, 0x67, 0xfa, 0x44, 0x76, 0x3b, 0xf0, 0x06, 0x53, 0xf4, 0x96, 0x85, 0x08, 0xa3, 0x67,
	0xc8, 0xff, 0xd8, 0x83, 0xc5, 0x58, 0x20, 0xc7, 0x36, 0xd8, 0x31, 0xa7, 0x92, 0xea, 0x1f, 0x0c,
	0xc0, 0x76, 0xfd, 0x8c, 0x42, 0x1d, 0x64, 0xae, 0x97, 0x6c, 0x26, 0x46, 0x5f, 0x2f, 0xb1, 0xfa,
	0x38, 0xd7, 0x3f, 0x24, 0x6a, 0x31, 0xae, 0xfa, 0x59, 0xfb, 0x58, 0x94, 0x72, 0x42, 0x6e, 0xee,
	0x69, 0xa3, 0x2e, 0x59, 0x48, 0x6b, 0x33, 0xb3, 0xc6, 0x09, 0x35, 0xf0, 0x27, 0xc0, 0xa5, 0xba,
	0x3d, 0x78, 0xe3, 0x31, 0x45, 0x06, 0xfd, 0xac, 0x94, 0x87, 0xf2, 0x12, 0x82, 0x13, 0x62, 0x7d,
	0xb3, 0x87, 0xbc, 0x63, 0x28, 0x90, 0x8f, 0xb6, 0x1e, 0x2e, 0x5e, 0xef, 0xbf, 0xef, 0x4e, 0x2d,
	0xf6, 0xcd, 0xe7, 0xdb, 0x7a, 0x8a, 0xee, 0x6a, 0x1a, 0xf9, 0xb4, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0xfb, 0xac, 0x17, 0x54, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *domain.SumRequest, opts ...grpc.CallOption) (*domain.SumResponse, error)
	GetAllSums(ctx context.Context, in *domain.SumFilter, opts ...grpc.CallOption) (SumService_GetAllSumsClient, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *domain.SumRequest, opts ...grpc.CallOption) (*domain.SumResponse, error) {
	out := new(domain.SumResponse)
	err := c.cc.Invoke(ctx, "/service.SumService/sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) GetAllSums(ctx context.Context, in *domain.SumFilter, opts ...grpc.CallOption) (SumService_GetAllSumsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/service.SumService/getAllSums", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceGetAllSumsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_GetAllSumsClient interface {
	Recv() (*domain.SumResponse, error)
	grpc.ClientStream
}

type sumServiceGetAllSumsClient struct {
	grpc.ClientStream
}

func (x *sumServiceGetAllSumsClient) Recv() (*domain.SumResponse, error) {
	m := new(domain.SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Sum(context.Context, *domain.SumRequest) (*domain.SumResponse, error)
	GetAllSums(*domain.SumFilter, SumService_GetAllSumsServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(ctx context.Context, req *domain.SumRequest) (*domain.SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedSumServiceServer) GetAllSums(req *domain.SumFilter, srv SumService_GetAllSumsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllSums not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*domain.SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_GetAllSums_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(domain.SumFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).GetAllSums(m, &sumServiceGetAllSumsServer{0, stream})
}

type SumService_GetAllSumsServer interface {
	Send(*domain.SumResponse) error
	grpc.ServerStream
	NumOfMessagesSent() int
}

type sumServiceGetAllSumsServer struct {
	count int
	grpc.ServerStream
}

func (x *sumServiceGetAllSumsServer) Send(m *domain.SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceGetAllSumsServer) NumOfMessagesSent() int {
	return x.count
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAllSums",
			Handler:       _SumService_GetAllSums_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/ydsxiong/summingservice/internal/proto-files/service/numbers-service.proto",
}
