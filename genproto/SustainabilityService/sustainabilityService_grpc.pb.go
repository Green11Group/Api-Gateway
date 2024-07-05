// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.18.0
// source: protos/sustainabilityService.proto

package SustainabilityService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SustainabilityService_LogImpact_FullMethodName               = "/SustainabilityService.SustainabilityService/LogImpact"
	SustainabilityService_GetUserImpact_FullMethodName           = "/SustainabilityService.SustainabilityService/GetUserImpact"
	SustainabilityService_GetCommunityImpact_FullMethodName      = "/SustainabilityService.SustainabilityService/GetCommunityImpact"
	SustainabilityService_GetChallenges_FullMethodName           = "/SustainabilityService.SustainabilityService/GetChallenges"
	SustainabilityService_JoinChallenge_FullMethodName           = "/SustainabilityService.SustainabilityService/JoinChallenge"
	SustainabilityService_UpdateChallengeProgress_FullMethodName = "/SustainabilityService.SustainabilityService/UpdateChallengeProgress"
	SustainabilityService_GetUserChallenges_FullMethodName       = "/SustainabilityService.SustainabilityService/GetUserChallenges"
	SustainabilityService_GetUserLeaderboard_FullMethodName      = "/SustainabilityService.SustainabilityService/GetUserLeaderboard"
	SustainabilityService_GetCommunityLeaderboard_FullMethodName = "/SustainabilityService.SustainabilityService/GetCommunityLeaderboard"
	SustainabilityService_PostChallenges_FullMethodName          = "/SustainabilityService.SustainabilityService/PostChallenges"
)

// SustainabilityServiceClient is the client API for SustainabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SustainabilityServiceClient interface {
	// 1
	LogImpact(ctx context.Context, in *LogImpactRequest, opts ...grpc.CallOption) (*LogImpactResponse, error)
	// 2
	GetUserImpact(ctx context.Context, in *GetUserImpactRequest, opts ...grpc.CallOption) (*GetUserImpactResponse, error)
	// 3
	GetCommunityImpact(ctx context.Context, in *GetCommunityImpactRequest, opts ...grpc.CallOption) (*GetCommunityImpactResponse, error)
	// 4
	GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error)
	// 5
	JoinChallenge(ctx context.Context, in *JoinChallengeRequest, opts ...grpc.CallOption) (*JoinChallengeResponse, error)
	// 6
	UpdateChallengeProgress(ctx context.Context, in *UpdateChallengeProgressRequest, opts ...grpc.CallOption) (*UpdateChallengeProgressResponse, error)
	// 7
	GetUserChallenges(ctx context.Context, in *GetUserChallengesRequest, opts ...grpc.CallOption) (*GetUserChallengesResponse, error)
	// 8
	GetUserLeaderboard(ctx context.Context, in *GetUserLeaderboardRequest, opts ...grpc.CallOption) (*GetUserLeaderboardResponse, error)
	// 9
	GetCommunityLeaderboard(ctx context.Context, in *GetCommunityLeaderboardRequest, opts ...grpc.CallOption) (*GetCommunityLeaderboardResponse, error)
	// 10
	PostChallenges(ctx context.Context, in *PostChallengesRequest, opts ...grpc.CallOption) (*PostChallengesResponse, error)
}

type sustainabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSustainabilityServiceClient(cc grpc.ClientConnInterface) SustainabilityServiceClient {
	return &sustainabilityServiceClient{cc}
}

func (c *sustainabilityServiceClient) LogImpact(ctx context.Context, in *LogImpactRequest, opts ...grpc.CallOption) (*LogImpactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogImpactResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_LogImpact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetUserImpact(ctx context.Context, in *GetUserImpactRequest, opts ...grpc.CallOption) (*GetUserImpactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserImpactResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetUserImpact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetCommunityImpact(ctx context.Context, in *GetCommunityImpactRequest, opts ...grpc.CallOption) (*GetCommunityImpactResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommunityImpactResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetCommunityImpact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChallengesResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetChallenges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) JoinChallenge(ctx context.Context, in *JoinChallengeRequest, opts ...grpc.CallOption) (*JoinChallengeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinChallengeResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_JoinChallenge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) UpdateChallengeProgress(ctx context.Context, in *UpdateChallengeProgressRequest, opts ...grpc.CallOption) (*UpdateChallengeProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateChallengeProgressResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_UpdateChallengeProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetUserChallenges(ctx context.Context, in *GetUserChallengesRequest, opts ...grpc.CallOption) (*GetUserChallengesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserChallengesResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetUserChallenges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetUserLeaderboard(ctx context.Context, in *GetUserLeaderboardRequest, opts ...grpc.CallOption) (*GetUserLeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserLeaderboardResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetUserLeaderboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) GetCommunityLeaderboard(ctx context.Context, in *GetCommunityLeaderboardRequest, opts ...grpc.CallOption) (*GetCommunityLeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommunityLeaderboardResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_GetCommunityLeaderboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sustainabilityServiceClient) PostChallenges(ctx context.Context, in *PostChallengesRequest, opts ...grpc.CallOption) (*PostChallengesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostChallengesResponse)
	err := c.cc.Invoke(ctx, SustainabilityService_PostChallenges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SustainabilityServiceServer is the server API for SustainabilityService service.
// All implementations must embed UnimplementedSustainabilityServiceServer
// for forward compatibility
type SustainabilityServiceServer interface {
	// 1
	LogImpact(context.Context, *LogImpactRequest) (*LogImpactResponse, error)
	// 2
	GetUserImpact(context.Context, *GetUserImpactRequest) (*GetUserImpactResponse, error)
	// 3
	GetCommunityImpact(context.Context, *GetCommunityImpactRequest) (*GetCommunityImpactResponse, error)
	// 4
	GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error)
	// 5
	JoinChallenge(context.Context, *JoinChallengeRequest) (*JoinChallengeResponse, error)
	// 6
	UpdateChallengeProgress(context.Context, *UpdateChallengeProgressRequest) (*UpdateChallengeProgressResponse, error)
	// 7
	GetUserChallenges(context.Context, *GetUserChallengesRequest) (*GetUserChallengesResponse, error)
	// 8
	GetUserLeaderboard(context.Context, *GetUserLeaderboardRequest) (*GetUserLeaderboardResponse, error)
	// 9
	GetCommunityLeaderboard(context.Context, *GetCommunityLeaderboardRequest) (*GetCommunityLeaderboardResponse, error)
	// 10
	PostChallenges(context.Context, *PostChallengesRequest) (*PostChallengesResponse, error)
	mustEmbedUnimplementedSustainabilityServiceServer()
}

// UnimplementedSustainabilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSustainabilityServiceServer struct {
}

func (UnimplementedSustainabilityServiceServer) LogImpact(context.Context, *LogImpactRequest) (*LogImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogImpact not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetUserImpact(context.Context, *GetUserImpactRequest) (*GetUserImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserImpact not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetCommunityImpact(context.Context, *GetCommunityImpactRequest) (*GetCommunityImpactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunityImpact not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenges not implemented")
}
func (UnimplementedSustainabilityServiceServer) JoinChallenge(context.Context, *JoinChallengeRequest) (*JoinChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChallenge not implemented")
}
func (UnimplementedSustainabilityServiceServer) UpdateChallengeProgress(context.Context, *UpdateChallengeProgressRequest) (*UpdateChallengeProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChallengeProgress not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetUserChallenges(context.Context, *GetUserChallengesRequest) (*GetUserChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserChallenges not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetUserLeaderboard(context.Context, *GetUserLeaderboardRequest) (*GetUserLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLeaderboard not implemented")
}
func (UnimplementedSustainabilityServiceServer) GetCommunityLeaderboard(context.Context, *GetCommunityLeaderboardRequest) (*GetCommunityLeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunityLeaderboard not implemented")
}
func (UnimplementedSustainabilityServiceServer) PostChallenges(context.Context, *PostChallengesRequest) (*PostChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostChallenges not implemented")
}
func (UnimplementedSustainabilityServiceServer) mustEmbedUnimplementedSustainabilityServiceServer() {}

// UnsafeSustainabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SustainabilityServiceServer will
// result in compilation errors.
type UnsafeSustainabilityServiceServer interface {
	mustEmbedUnimplementedSustainabilityServiceServer()
}

func RegisterSustainabilityServiceServer(s grpc.ServiceRegistrar, srv SustainabilityServiceServer) {
	s.RegisterService(&SustainabilityService_ServiceDesc, srv)
}

func _SustainabilityService_LogImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).LogImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_LogImpact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).LogImpact(ctx, req.(*LogImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetUserImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetUserImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetUserImpact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetUserImpact(ctx, req.(*GetUserImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetCommunityImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityImpactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetCommunityImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetCommunityImpact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetCommunityImpact(ctx, req.(*GetCommunityImpactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetChallenges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetChallenges(ctx, req.(*GetChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_JoinChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).JoinChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_JoinChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).JoinChallenge(ctx, req.(*JoinChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_UpdateChallengeProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChallengeProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).UpdateChallengeProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_UpdateChallengeProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).UpdateChallengeProgress(ctx, req.(*UpdateChallengeProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetUserChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetUserChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetUserChallenges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetUserChallenges(ctx, req.(*GetUserChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetUserLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetUserLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetUserLeaderboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetUserLeaderboard(ctx, req.(*GetUserLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_GetCommunityLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityLeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).GetCommunityLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_GetCommunityLeaderboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).GetCommunityLeaderboard(ctx, req.(*GetCommunityLeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SustainabilityService_PostChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SustainabilityServiceServer).PostChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SustainabilityService_PostChallenges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SustainabilityServiceServer).PostChallenges(ctx, req.(*PostChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SustainabilityService_ServiceDesc is the grpc.ServiceDesc for SustainabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SustainabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SustainabilityService.SustainabilityService",
	HandlerType: (*SustainabilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogImpact",
			Handler:    _SustainabilityService_LogImpact_Handler,
		},
		{
			MethodName: "GetUserImpact",
			Handler:    _SustainabilityService_GetUserImpact_Handler,
		},
		{
			MethodName: "GetCommunityImpact",
			Handler:    _SustainabilityService_GetCommunityImpact_Handler,
		},
		{
			MethodName: "GetChallenges",
			Handler:    _SustainabilityService_GetChallenges_Handler,
		},
		{
			MethodName: "JoinChallenge",
			Handler:    _SustainabilityService_JoinChallenge_Handler,
		},
		{
			MethodName: "UpdateChallengeProgress",
			Handler:    _SustainabilityService_UpdateChallengeProgress_Handler,
		},
		{
			MethodName: "GetUserChallenges",
			Handler:    _SustainabilityService_GetUserChallenges_Handler,
		},
		{
			MethodName: "GetUserLeaderboard",
			Handler:    _SustainabilityService_GetUserLeaderboard_Handler,
		},
		{
			MethodName: "GetCommunityLeaderboard",
			Handler:    _SustainabilityService_GetCommunityLeaderboard_Handler,
		},
		{
			MethodName: "PostChallenges",
			Handler:    _SustainabilityService_PostChallenges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/sustainabilityService.proto",
}
