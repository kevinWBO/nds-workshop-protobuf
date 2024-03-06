// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: bangunruang.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BangunRuang_HitungVolumeKubus_FullMethodName     = "/proto.BangunRuang/HitungVolumeKubus"
	BangunRuang_HitungKelilingPersegi_FullMethodName = "/proto.BangunRuang/HitungKelilingPersegi"
	BangunRuang_HitungLingkaran_FullMethodName       = "/proto.BangunRuang/HitungLingkaran"
	BangunRuang_HitungSegitiga_FullMethodName        = "/proto.BangunRuang/HitungSegitiga"
)

// BangunRuangClient is the client API for BangunRuang service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BangunRuangClient interface {
	HitungVolumeKubus(ctx context.Context, in *RequestKubus, opts ...grpc.CallOption) (*ResponseVolumeKubus, error)
	HitungKelilingPersegi(ctx context.Context, in *RequestPersegi, opts ...grpc.CallOption) (*ResponseKelilingPersegi, error)
	HitungLingkaran(ctx context.Context, in *RequestLingkaran, opts ...grpc.CallOption) (*ResponseLingkaran, error)
	HitungSegitiga(ctx context.Context, in *RequestSegitiga, opts ...grpc.CallOption) (*ResponseSegitiga, error)
}

type bangunRuangClient struct {
	cc grpc.ClientConnInterface
}

func NewBangunRuangClient(cc grpc.ClientConnInterface) BangunRuangClient {
	return &bangunRuangClient{cc}
}

func (c *bangunRuangClient) HitungVolumeKubus(ctx context.Context, in *RequestKubus, opts ...grpc.CallOption) (*ResponseVolumeKubus, error) {
	out := new(ResponseVolumeKubus)
	err := c.cc.Invoke(ctx, BangunRuang_HitungVolumeKubus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bangunRuangClient) HitungKelilingPersegi(ctx context.Context, in *RequestPersegi, opts ...grpc.CallOption) (*ResponseKelilingPersegi, error) {
	out := new(ResponseKelilingPersegi)
	err := c.cc.Invoke(ctx, BangunRuang_HitungKelilingPersegi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bangunRuangClient) HitungLingkaran(ctx context.Context, in *RequestLingkaran, opts ...grpc.CallOption) (*ResponseLingkaran, error) {
	out := new(ResponseLingkaran)
	err := c.cc.Invoke(ctx, BangunRuang_HitungLingkaran_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bangunRuangClient) HitungSegitiga(ctx context.Context, in *RequestSegitiga, opts ...grpc.CallOption) (*ResponseSegitiga, error) {
	out := new(ResponseSegitiga)
	err := c.cc.Invoke(ctx, BangunRuang_HitungSegitiga_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BangunRuangServer is the server API for BangunRuang service.
// All implementations should embed UnimplementedBangunRuangServer
// for forward compatibility
type BangunRuangServer interface {
	HitungVolumeKubus(context.Context, *RequestKubus) (*ResponseVolumeKubus, error)
	HitungKelilingPersegi(context.Context, *RequestPersegi) (*ResponseKelilingPersegi, error)
	HitungLingkaran(context.Context, *RequestLingkaran) (*ResponseLingkaran, error)
	HitungSegitiga(context.Context, *RequestSegitiga) (*ResponseSegitiga, error)
}

// UnimplementedBangunRuangServer should be embedded to have forward compatible implementations.
type UnimplementedBangunRuangServer struct {
}

func (UnimplementedBangunRuangServer) HitungVolumeKubus(context.Context, *RequestKubus) (*ResponseVolumeKubus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungVolumeKubus not implemented")
}
func (UnimplementedBangunRuangServer) HitungKelilingPersegi(context.Context, *RequestPersegi) (*ResponseKelilingPersegi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungKelilingPersegi not implemented")
}
func (UnimplementedBangunRuangServer) HitungLingkaran(context.Context, *RequestLingkaran) (*ResponseLingkaran, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungLingkaran not implemented")
}
func (UnimplementedBangunRuangServer) HitungSegitiga(context.Context, *RequestSegitiga) (*ResponseSegitiga, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HitungSegitiga not implemented")
}

// UnsafeBangunRuangServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BangunRuangServer will
// result in compilation errors.
type UnsafeBangunRuangServer interface {
	mustEmbedUnimplementedBangunRuangServer()
}

func RegisterBangunRuangServer(s grpc.ServiceRegistrar, srv BangunRuangServer) {
	s.RegisterService(&BangunRuang_ServiceDesc, srv)
}

func _BangunRuang_HitungVolumeKubus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestKubus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BangunRuangServer).HitungVolumeKubus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BangunRuang_HitungVolumeKubus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BangunRuangServer).HitungVolumeKubus(ctx, req.(*RequestKubus))
	}
	return interceptor(ctx, in, info, handler)
}

func _BangunRuang_HitungKelilingPersegi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPersegi)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BangunRuangServer).HitungKelilingPersegi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BangunRuang_HitungKelilingPersegi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BangunRuangServer).HitungKelilingPersegi(ctx, req.(*RequestPersegi))
	}
	return interceptor(ctx, in, info, handler)
}

func _BangunRuang_HitungLingkaran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLingkaran)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BangunRuangServer).HitungLingkaran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BangunRuang_HitungLingkaran_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BangunRuangServer).HitungLingkaran(ctx, req.(*RequestLingkaran))
	}
	return interceptor(ctx, in, info, handler)
}

func _BangunRuang_HitungSegitiga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSegitiga)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BangunRuangServer).HitungSegitiga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BangunRuang_HitungSegitiga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BangunRuangServer).HitungSegitiga(ctx, req.(*RequestSegitiga))
	}
	return interceptor(ctx, in, info, handler)
}

// BangunRuang_ServiceDesc is the grpc.ServiceDesc for BangunRuang service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BangunRuang_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BangunRuang",
	HandlerType: (*BangunRuangServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HitungVolumeKubus",
			Handler:    _BangunRuang_HitungVolumeKubus_Handler,
		},
		{
			MethodName: "HitungKelilingPersegi",
			Handler:    _BangunRuang_HitungKelilingPersegi_Handler,
		},
		{
			MethodName: "HitungLingkaran",
			Handler:    _BangunRuang_HitungLingkaran_Handler,
		},
		{
			MethodName: "HitungSegitiga",
			Handler:    _BangunRuang_HitungSegitiga_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bangunruang.proto",
}