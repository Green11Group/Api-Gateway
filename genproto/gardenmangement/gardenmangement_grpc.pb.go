// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.18.0
// source: protos/gardenmangement.proto

package gardenmangement

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
	GardenService_CreateGarden_FullMethodName   = "/garden.GardenService/CreateGarden"
	GardenService_UpdateGarden_FullMethodName   = "/garden.GardenService/UpdateGarden"
	GardenService_GetGarden_FullMethodName      = "/garden.GardenService/GetGarden"
	GardenService_DeleteGareden_FullMethodName  = "/garden.GardenService/DeleteGareden"
	GardenService_GetUserGardens_FullMethodName = "/garden.GardenService/GetUserGardens"
	GardenService_CreatePlant_FullMethodName    = "/garden.GardenService/CreatePlant"
	GardenService_GetPlant_FullMethodName       = "/garden.GardenService/GetPlant"
	GardenService_UpdatePlant_FullMethodName    = "/garden.GardenService/UpdatePlant"
	GardenService_DeletePlant_FullMethodName    = "/garden.GardenService/DeletePlant"
	GardenService_CreateCareLog_FullMethodName  = "/garden.GardenService/CreateCareLog"
	GardenService_GetCareLog_FullMethodName     = "/garden.GardenService/GetCareLog"
)

// GardenServiceClient is the client API for GardenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GardenServiceClient interface {
	// Yangi bog' yaratish
	CreateGarden(ctx context.Context, in *CreateGardenRequest, opts ...grpc.CallOption) (*CreateGardenResponse, error)
	// Bog'ni yangilash
	UpdateGarden(ctx context.Context, in *UpdateGardenRequest, opts ...grpc.CallOption) (*UpdateGardenResponse, error)
	// Bog'  malumotini olish xizmati
	GetGarden(ctx context.Context, in *GetGardenRequest, opts ...grpc.CallOption) (*GetGardenResponse, error)
	// Bog'ni o'chirish xizmati
	DeleteGareden(ctx context.Context, in *DeleteGardenRequest, opts ...grpc.CallOption) (*DeleteGardenResponse, error)
	// Foydalanuvchining bog'larini ro'yxatlash
	GetUserGardens(ctx context.Context, in *GetUserGardensRequest, opts ...grpc.CallOption) (*GetUserGardensResponse, error)
	// Yangi o'simliklar yaratish
	CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantReponse, error)
	// O'simlik malumotlarini olish
	GetPlant(ctx context.Context, in *GetPlantRequest, opts ...grpc.CallOption) (*GetPlantResponse, error)
	// O'simlikni yangilash
	UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error)
	// O'simlikni o'chirish
	DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error)
	// Parvarish loglarini yaratish
	CreateCareLog(ctx context.Context, in *CreateCareLogRequest, opts ...grpc.CallOption) (*CreateCareLogResponse, error)
	// Parvarish loglarini ro'yxatlash
	GetCareLog(ctx context.Context, in *GetCareLogRequest, opts ...grpc.CallOption) (*GetCareLogResponse, error)
}

type gardenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGardenServiceClient(cc grpc.ClientConnInterface) GardenServiceClient {
	return &gardenServiceClient{cc}
}

func (c *gardenServiceClient) CreateGarden(ctx context.Context, in *CreateGardenRequest, opts ...grpc.CallOption) (*CreateGardenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGardenResponse)
	err := c.cc.Invoke(ctx, GardenService_CreateGarden_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) UpdateGarden(ctx context.Context, in *UpdateGardenRequest, opts ...grpc.CallOption) (*UpdateGardenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGardenResponse)
	err := c.cc.Invoke(ctx, GardenService_UpdateGarden_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) GetGarden(ctx context.Context, in *GetGardenRequest, opts ...grpc.CallOption) (*GetGardenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGardenResponse)
	err := c.cc.Invoke(ctx, GardenService_GetGarden_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) DeleteGareden(ctx context.Context, in *DeleteGardenRequest, opts ...grpc.CallOption) (*DeleteGardenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGardenResponse)
	err := c.cc.Invoke(ctx, GardenService_DeleteGareden_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) GetUserGardens(ctx context.Context, in *GetUserGardensRequest, opts ...grpc.CallOption) (*GetUserGardensResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserGardensResponse)
	err := c.cc.Invoke(ctx, GardenService_GetUserGardens_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantReponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePlantReponse)
	err := c.cc.Invoke(ctx, GardenService_CreatePlant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) GetPlant(ctx context.Context, in *GetPlantRequest, opts ...grpc.CallOption) (*GetPlantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlantResponse)
	err := c.cc.Invoke(ctx, GardenService_GetPlant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePlantResponse)
	err := c.cc.Invoke(ctx, GardenService_UpdatePlant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePlantResponse)
	err := c.cc.Invoke(ctx, GardenService_DeletePlant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) CreateCareLog(ctx context.Context, in *CreateCareLogRequest, opts ...grpc.CallOption) (*CreateCareLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCareLogResponse)
	err := c.cc.Invoke(ctx, GardenService_CreateCareLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gardenServiceClient) GetCareLog(ctx context.Context, in *GetCareLogRequest, opts ...grpc.CallOption) (*GetCareLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCareLogResponse)
	err := c.cc.Invoke(ctx, GardenService_GetCareLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GardenServiceServer is the server API for GardenService service.
// All implementations must embed UnimplementedGardenServiceServer
// for forward compatibility
type GardenServiceServer interface {
	// Yangi bog' yaratish
	CreateGarden(context.Context, *CreateGardenRequest) (*CreateGardenResponse, error)
	// Bog'ni yangilash
	UpdateGarden(context.Context, *UpdateGardenRequest) (*UpdateGardenResponse, error)
	// Bog'  malumotini olish xizmati
	GetGarden(context.Context, *GetGardenRequest) (*GetGardenResponse, error)
	// Bog'ni o'chirish xizmati
	DeleteGareden(context.Context, *DeleteGardenRequest) (*DeleteGardenResponse, error)
	// Foydalanuvchining bog'larini ro'yxatlash
	GetUserGardens(context.Context, *GetUserGardensRequest) (*GetUserGardensResponse, error)
	// Yangi o'simliklar yaratish
	CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantReponse, error)
	// O'simlik malumotlarini olish
	GetPlant(context.Context, *GetPlantRequest) (*GetPlantResponse, error)
	// O'simlikni yangilash
	UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error)
	// O'simlikni o'chirish
	DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error)
	// Parvarish loglarini yaratish
	CreateCareLog(context.Context, *CreateCareLogRequest) (*CreateCareLogResponse, error)
	// Parvarish loglarini ro'yxatlash
	GetCareLog(context.Context, *GetCareLogRequest) (*GetCareLogResponse, error)
	mustEmbedUnimplementedGardenServiceServer()
}

// UnimplementedGardenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGardenServiceServer struct {
}

func (UnimplementedGardenServiceServer) CreateGarden(context.Context, *CreateGardenRequest) (*CreateGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGarden not implemented")
}
func (UnimplementedGardenServiceServer) UpdateGarden(context.Context, *UpdateGardenRequest) (*UpdateGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGarden not implemented")
}
func (UnimplementedGardenServiceServer) GetGarden(context.Context, *GetGardenRequest) (*GetGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGarden not implemented")
}
func (UnimplementedGardenServiceServer) DeleteGareden(context.Context, *DeleteGardenRequest) (*DeleteGardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGareden not implemented")
}
func (UnimplementedGardenServiceServer) GetUserGardens(context.Context, *GetUserGardensRequest) (*GetUserGardensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGardens not implemented")
}
func (UnimplementedGardenServiceServer) CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlant not implemented")
}
func (UnimplementedGardenServiceServer) GetPlant(context.Context, *GetPlantRequest) (*GetPlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlant not implemented")
}
func (UnimplementedGardenServiceServer) UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlant not implemented")
}
func (UnimplementedGardenServiceServer) DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlant not implemented")
}
func (UnimplementedGardenServiceServer) CreateCareLog(context.Context, *CreateCareLogRequest) (*CreateCareLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCareLog not implemented")
}
func (UnimplementedGardenServiceServer) GetCareLog(context.Context, *GetCareLogRequest) (*GetCareLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCareLog not implemented")
}
func (UnimplementedGardenServiceServer) mustEmbedUnimplementedGardenServiceServer() {}

// UnsafeGardenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GardenServiceServer will
// result in compilation errors.
type UnsafeGardenServiceServer interface {
	mustEmbedUnimplementedGardenServiceServer()
}

func RegisterGardenServiceServer(s grpc.ServiceRegistrar, srv GardenServiceServer) {
	s.RegisterService(&GardenService_ServiceDesc, srv)
}

func _GardenService_CreateGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).CreateGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_CreateGarden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).CreateGarden(ctx, req.(*CreateGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_UpdateGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).UpdateGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_UpdateGarden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).UpdateGarden(ctx, req.(*UpdateGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_GetGarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).GetGarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_GetGarden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).GetGarden(ctx, req.(*GetGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_DeleteGareden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).DeleteGareden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_DeleteGareden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).DeleteGareden(ctx, req.(*DeleteGardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_GetUserGardens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGardensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).GetUserGardens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_GetUserGardens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).GetUserGardens(ctx, req.(*GetUserGardensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_CreatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).CreatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_CreatePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).CreatePlant(ctx, req.(*CreatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_GetPlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).GetPlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_GetPlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).GetPlant(ctx, req.(*GetPlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_UpdatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).UpdatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_UpdatePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).UpdatePlant(ctx, req.(*UpdatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_DeletePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).DeletePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_DeletePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).DeletePlant(ctx, req.(*DeletePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_CreateCareLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCareLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).CreateCareLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_CreateCareLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).CreateCareLog(ctx, req.(*CreateCareLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GardenService_GetCareLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCareLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GardenServiceServer).GetCareLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GardenService_GetCareLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GardenServiceServer).GetCareLog(ctx, req.(*GetCareLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GardenService_ServiceDesc is the grpc.ServiceDesc for GardenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GardenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "garden.GardenService",
	HandlerType: (*GardenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGarden",
			Handler:    _GardenService_CreateGarden_Handler,
		},
		{
			MethodName: "UpdateGarden",
			Handler:    _GardenService_UpdateGarden_Handler,
		},
		{
			MethodName: "GetGarden",
			Handler:    _GardenService_GetGarden_Handler,
		},
		{
			MethodName: "DeleteGareden",
			Handler:    _GardenService_DeleteGareden_Handler,
		},
		{
			MethodName: "GetUserGardens",
			Handler:    _GardenService_GetUserGardens_Handler,
		},
		{
			MethodName: "CreatePlant",
			Handler:    _GardenService_CreatePlant_Handler,
		},
		{
			MethodName: "GetPlant",
			Handler:    _GardenService_GetPlant_Handler,
		},
		{
			MethodName: "UpdatePlant",
			Handler:    _GardenService_UpdatePlant_Handler,
		},
		{
			MethodName: "DeletePlant",
			Handler:    _GardenService_DeletePlant_Handler,
		},
		{
			MethodName: "CreateCareLog",
			Handler:    _GardenService_CreateCareLog_Handler,
		},
		{
			MethodName: "GetCareLog",
			Handler:    _GardenService_GetCareLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/gardenmangement.proto",
}
