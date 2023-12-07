// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: handler/campaign/grpc/campaign.proto

package grpc

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

// CampaignHandlerClient is the client API for CampaignHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignHandlerClient interface {
	GetCampaignByID(ctx context.Context, in *GetCampaignByIDRequest, opts ...grpc.CallOption) (*Campaign, error)
	GetCampaignByListID(ctx context.Context, opts ...grpc.CallOption) (CampaignHandler_GetCampaignByListIDClient, error)
}

type campaignHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignHandlerClient(cc grpc.ClientConnInterface) CampaignHandlerClient {
	return &campaignHandlerClient{cc}
}

func (c *campaignHandlerClient) GetCampaignByID(ctx context.Context, in *GetCampaignByIDRequest, opts ...grpc.CallOption) (*Campaign, error) {
	out := new(Campaign)
	err := c.cc.Invoke(ctx, "/CampaignHandler/GetCampaignByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignHandlerClient) GetCampaignByListID(ctx context.Context, opts ...grpc.CallOption) (CampaignHandler_GetCampaignByListIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &CampaignHandler_ServiceDesc.Streams[0], "/CampaignHandler/GetCampaignByListID", opts...)
	if err != nil {
		return nil, err
	}
	x := &campaignHandlerGetCampaignByListIDClient{stream}
	return x, nil
}

type CampaignHandler_GetCampaignByListIDClient interface {
	Send(*GetCampaignByIDRequest) error
	CloseAndRecv() (*GetCampaignByIDResponse, error)
	grpc.ClientStream
}

type campaignHandlerGetCampaignByListIDClient struct {
	grpc.ClientStream
}

func (x *campaignHandlerGetCampaignByListIDClient) Send(m *GetCampaignByIDRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *campaignHandlerGetCampaignByListIDClient) CloseAndRecv() (*GetCampaignByIDResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetCampaignByIDResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CampaignHandlerServer is the server API for CampaignHandler service.
// All implementations must embed UnimplementedCampaignHandlerServer
// for forward compatibility
type CampaignHandlerServer interface {
	GetCampaignByID(context.Context, *GetCampaignByIDRequest) (*Campaign, error)
	GetCampaignByListID(CampaignHandler_GetCampaignByListIDServer) error
	mustEmbedUnimplementedCampaignHandlerServer()
}

// UnimplementedCampaignHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignHandlerServer struct {
}

func (UnimplementedCampaignHandlerServer) GetCampaignByID(context.Context, *GetCampaignByIDRequest) (*Campaign, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignByID not implemented")
}
func (UnimplementedCampaignHandlerServer) GetCampaignByListID(CampaignHandler_GetCampaignByListIDServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCampaignByListID not implemented")
}
func (UnimplementedCampaignHandlerServer) mustEmbedUnimplementedCampaignHandlerServer() {}

// UnsafeCampaignHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignHandlerServer will
// result in compilation errors.
type UnsafeCampaignHandlerServer interface {
	mustEmbedUnimplementedCampaignHandlerServer()
}

func RegisterCampaignHandlerServer(s grpc.ServiceRegistrar, srv CampaignHandlerServer) {
	s.RegisterService(&CampaignHandler_ServiceDesc, srv)
}

func _CampaignHandler_GetCampaignByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignHandlerServer).GetCampaignByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CampaignHandler/GetCampaignByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignHandlerServer).GetCampaignByID(ctx, req.(*GetCampaignByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignHandler_GetCampaignByListID_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CampaignHandlerServer).GetCampaignByListID(&campaignHandlerGetCampaignByListIDServer{stream})
}

type CampaignHandler_GetCampaignByListIDServer interface {
	SendAndClose(*GetCampaignByIDResponse) error
	Recv() (*GetCampaignByIDRequest, error)
	grpc.ServerStream
}

type campaignHandlerGetCampaignByListIDServer struct {
	grpc.ServerStream
}

func (x *campaignHandlerGetCampaignByListIDServer) SendAndClose(m *GetCampaignByIDResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *campaignHandlerGetCampaignByListIDServer) Recv() (*GetCampaignByIDRequest, error) {
	m := new(GetCampaignByIDRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CampaignHandler_ServiceDesc is the grpc.ServiceDesc for CampaignHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CampaignHandler",
	HandlerType: (*CampaignHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignByID",
			Handler:    _CampaignHandler_GetCampaignByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCampaignByListID",
			Handler:       _CampaignHandler_GetCampaignByListID_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "handler/campaign/grpc/campaign.proto",
}