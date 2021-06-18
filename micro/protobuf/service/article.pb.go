// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/service/article.proto

package service // import "JupiterBlog/micro/protobuf/service"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import request "JupiterBlog/micro/protobuf/request"
import response "JupiterBlog/micro/protobuf/response"

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

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleClient interface {
	Info(ctx context.Context, in *request.ArticleInfo, opts ...grpc.CallOption) (*response.ArticleInfo, error)
}

type articleClient struct {
	cc *grpc.ClientConn
}

func NewArticleClient(cc *grpc.ClientConn) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) Info(ctx context.Context, in *request.ArticleInfo, opts ...grpc.CallOption) (*response.ArticleInfo, error) {
	out := new(response.ArticleInfo)
	err := c.cc.Invoke(ctx, "/service.Article/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
type ArticleServer interface {
	Info(context.Context, *request.ArticleInfo) (*response.ArticleInfo, error)
}

func RegisterArticleServer(s *grpc.Server, srv ArticleServer) {
	s.RegisterService(&_Article_serviceDesc, srv)
}

func _Article_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ArticleInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Article/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Info(ctx, req.(*request.ArticleInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Article_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Article_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/service/article.proto",
}

func init() {
	proto.RegisterFile("protobuf/service/article.proto", fileDescriptor_article_0ca4de55017eed56)
}

var fileDescriptor_article_0ca4de55017eed56 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x2a,
	0xc9, 0x4c, 0xce, 0x49, 0xd5, 0x03, 0x4b, 0x08, 0xb1, 0x43, 0x85, 0xa5, 0x10, 0x0a, 0x8b, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x50, 0x15, 0x4a, 0xc9, 0x23, 0xc9, 0x17, 0x17, 0xe4, 0xe7, 0x15,
	0xa3, 0x99, 0x64, 0x64, 0xc7, 0xc5, 0xee, 0x08, 0x11, 0x10, 0x32, 0xe6, 0x62, 0xf1, 0xcc, 0x4b,
	0xcb, 0x17, 0x12, 0xd1, 0x83, 0x9a, 0xa5, 0x07, 0x95, 0x01, 0x89, 0x4a, 0x89, 0xea, 0xc1, 0x4c,
	0x40, 0x16, 0x76, 0x52, 0x89, 0x52, 0xf2, 0x2a, 0x2d, 0xc8, 0x2c, 0x49, 0x2d, 0x72, 0xca, 0xc9,
	0x4f, 0xd7, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x47, 0x77, 0x7d, 0x12, 0x1b, 0x58, 0xc4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xaf, 0x57, 0x94, 0xd8, 0x00, 0x00, 0x00,
}