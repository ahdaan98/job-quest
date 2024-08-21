// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: auth/auth.proto

package auth

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
	Admin_AdminSignup_FullMethodName = "/auth.Admin/AdminSignup"
	Admin_AdminLogin_FullMethodName  = "/auth.Admin/AdminLogin"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	AdminSignup(ctx context.Context, in *AdminSignupRequest, opts ...grpc.CallOption) (*AdminSignupResponse, error)
	AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) AdminSignup(ctx context.Context, in *AdminSignupRequest, opts ...grpc.CallOption) (*AdminSignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminSignupResponse)
	err := c.cc.Invoke(ctx, Admin_AdminSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AdminLogin(ctx context.Context, in *AdminLoginInRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, Admin_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	AdminSignup(context.Context, *AdminSignupRequest) (*AdminSignupResponse, error)
	AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) AdminSignup(context.Context, *AdminSignupRequest) (*AdminSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminSignup not implemented")
}
func (UnimplementedAdminServer) AdminLogin(context.Context, *AdminLoginInRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_AdminSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminSignup(ctx, req.(*AdminSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AdminLogin(ctx, req.(*AdminLoginInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminSignup",
			Handler:    _Admin_AdminSignup_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _Admin_AdminLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	Employer_EmployerSignup_FullMethodName    = "/auth.Employer/EmployerSignup"
	Employer_EmployerLogin_FullMethodName     = "/auth.Employer/EmployerLogin"
	Employer_PostJobOpening_FullMethodName    = "/auth.Employer/PostJobOpening"
	Employer_GetCompanyDetails_FullMethodName = "/auth.Employer/GetCompanyDetails"
	Employer_UpdateCompany_FullMethodName     = "/auth.Employer/UpdateCompany"
)

// EmployerClient is the client API for Employer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployerClient interface {
	EmployerSignup(ctx context.Context, in *EmployerSignupRequest, opts ...grpc.CallOption) (*EmployerSignupResponse, error)
	EmployerLogin(ctx context.Context, in *EmployerLoginInRequest, opts ...grpc.CallOption) (*EmployerLoginResponse, error)
	PostJobOpening(ctx context.Context, in *PostJobOpeningRequest, opts ...grpc.CallOption) (*PostJobOpeningResponse, error)
	GetCompanyDetails(ctx context.Context, in *GetCompanyDetailsRequest, opts ...grpc.CallOption) (*EmployerDetailsResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
}

type employerClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployerClient(cc grpc.ClientConnInterface) EmployerClient {
	return &employerClient{cc}
}

func (c *employerClient) EmployerSignup(ctx context.Context, in *EmployerSignupRequest, opts ...grpc.CallOption) (*EmployerSignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployerSignupResponse)
	err := c.cc.Invoke(ctx, Employer_EmployerSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employerClient) EmployerLogin(ctx context.Context, in *EmployerLoginInRequest, opts ...grpc.CallOption) (*EmployerLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployerLoginResponse)
	err := c.cc.Invoke(ctx, Employer_EmployerLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employerClient) PostJobOpening(ctx context.Context, in *PostJobOpeningRequest, opts ...grpc.CallOption) (*PostJobOpeningResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostJobOpeningResponse)
	err := c.cc.Invoke(ctx, Employer_PostJobOpening_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employerClient) GetCompanyDetails(ctx context.Context, in *GetCompanyDetailsRequest, opts ...grpc.CallOption) (*EmployerDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmployerDetailsResponse)
	err := c.cc.Invoke(ctx, Employer_GetCompanyDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employerClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, Employer_UpdateCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployerServer is the server API for Employer service.
// All implementations must embed UnimplementedEmployerServer
// for forward compatibility
type EmployerServer interface {
	EmployerSignup(context.Context, *EmployerSignupRequest) (*EmployerSignupResponse, error)
	EmployerLogin(context.Context, *EmployerLoginInRequest) (*EmployerLoginResponse, error)
	PostJobOpening(context.Context, *PostJobOpeningRequest) (*PostJobOpeningResponse, error)
	GetCompanyDetails(context.Context, *GetCompanyDetailsRequest) (*EmployerDetailsResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	mustEmbedUnimplementedEmployerServer()
}

// UnimplementedEmployerServer must be embedded to have forward compatible implementations.
type UnimplementedEmployerServer struct {
}

func (UnimplementedEmployerServer) EmployerSignup(context.Context, *EmployerSignupRequest) (*EmployerSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmployerSignup not implemented")
}
func (UnimplementedEmployerServer) EmployerLogin(context.Context, *EmployerLoginInRequest) (*EmployerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmployerLogin not implemented")
}
func (UnimplementedEmployerServer) PostJobOpening(context.Context, *PostJobOpeningRequest) (*PostJobOpeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostJobOpening not implemented")
}
func (UnimplementedEmployerServer) GetCompanyDetails(context.Context, *GetCompanyDetailsRequest) (*EmployerDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyDetails not implemented")
}
func (UnimplementedEmployerServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedEmployerServer) mustEmbedUnimplementedEmployerServer() {}

// UnsafeEmployerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployerServer will
// result in compilation errors.
type UnsafeEmployerServer interface {
	mustEmbedUnimplementedEmployerServer()
}

func RegisterEmployerServer(s grpc.ServiceRegistrar, srv EmployerServer) {
	s.RegisterService(&Employer_ServiceDesc, srv)
}

func _Employer_EmployerSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployerSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).EmployerSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_EmployerSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).EmployerSignup(ctx, req.(*EmployerSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employer_EmployerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployerLoginInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).EmployerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_EmployerLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).EmployerLogin(ctx, req.(*EmployerLoginInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employer_PostJobOpening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostJobOpeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).PostJobOpening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_PostJobOpening_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).PostJobOpening(ctx, req.(*PostJobOpeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employer_GetCompanyDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).GetCompanyDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_GetCompanyDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).GetCompanyDetails(ctx, req.(*GetCompanyDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employer_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployerServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Employer_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployerServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Employer_ServiceDesc is the grpc.ServiceDesc for Employer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Employer",
	HandlerType: (*EmployerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmployerSignup",
			Handler:    _Employer_EmployerSignup_Handler,
		},
		{
			MethodName: "EmployerLogin",
			Handler:    _Employer_EmployerLogin_Handler,
		},
		{
			MethodName: "PostJobOpening",
			Handler:    _Employer_PostJobOpening_Handler,
		},
		{
			MethodName: "GetCompanyDetails",
			Handler:    _Employer_GetCompanyDetails_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Employer_UpdateCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	JobSeeker_JobSeekerSignup_FullMethodName = "/auth.JobSeeker/JobSeekerSignup"
	JobSeeker_JobSeekerLogin_FullMethodName  = "/auth.JobSeeker/JobSeekerLogin"
)

// JobSeekerClient is the client API for JobSeeker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobSeekerClient interface {
	JobSeekerSignup(ctx context.Context, in *JobSeekerSignupRequest, opts ...grpc.CallOption) (*JobSeekerSignupResponse, error)
	JobSeekerLogin(ctx context.Context, in *JobSeekerLoginRequest, opts ...grpc.CallOption) (*JobSeekerLoginResponse, error)
}

type jobSeekerClient struct {
	cc grpc.ClientConnInterface
}

func NewJobSeekerClient(cc grpc.ClientConnInterface) JobSeekerClient {
	return &jobSeekerClient{cc}
}

func (c *jobSeekerClient) JobSeekerSignup(ctx context.Context, in *JobSeekerSignupRequest, opts ...grpc.CallOption) (*JobSeekerSignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobSeekerSignupResponse)
	err := c.cc.Invoke(ctx, JobSeeker_JobSeekerSignup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobSeekerClient) JobSeekerLogin(ctx context.Context, in *JobSeekerLoginRequest, opts ...grpc.CallOption) (*JobSeekerLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobSeekerLoginResponse)
	err := c.cc.Invoke(ctx, JobSeeker_JobSeekerLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobSeekerServer is the server API for JobSeeker service.
// All implementations must embed UnimplementedJobSeekerServer
// for forward compatibility
type JobSeekerServer interface {
	JobSeekerSignup(context.Context, *JobSeekerSignupRequest) (*JobSeekerSignupResponse, error)
	JobSeekerLogin(context.Context, *JobSeekerLoginRequest) (*JobSeekerLoginResponse, error)
	mustEmbedUnimplementedJobSeekerServer()
}

// UnimplementedJobSeekerServer must be embedded to have forward compatible implementations.
type UnimplementedJobSeekerServer struct {
}

func (UnimplementedJobSeekerServer) JobSeekerSignup(context.Context, *JobSeekerSignupRequest) (*JobSeekerSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerSignup not implemented")
}
func (UnimplementedJobSeekerServer) JobSeekerLogin(context.Context, *JobSeekerLoginRequest) (*JobSeekerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobSeekerLogin not implemented")
}
func (UnimplementedJobSeekerServer) mustEmbedUnimplementedJobSeekerServer() {}

// UnsafeJobSeekerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobSeekerServer will
// result in compilation errors.
type UnsafeJobSeekerServer interface {
	mustEmbedUnimplementedJobSeekerServer()
}

func RegisterJobSeekerServer(s grpc.ServiceRegistrar, srv JobSeekerServer) {
	s.RegisterService(&JobSeeker_ServiceDesc, srv)
}

func _JobSeeker_JobSeekerSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSeekerServer).JobSeekerSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobSeeker_JobSeekerSignup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSeekerServer).JobSeekerSignup(ctx, req.(*JobSeekerSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobSeeker_JobSeekerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobSeekerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobSeekerServer).JobSeekerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobSeeker_JobSeekerLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobSeekerServer).JobSeekerLogin(ctx, req.(*JobSeekerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobSeeker_ServiceDesc is the grpc.ServiceDesc for JobSeeker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobSeeker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.JobSeeker",
	HandlerType: (*JobSeekerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JobSeekerSignup",
			Handler:    _JobSeeker_JobSeekerSignup_Handler,
		},
		{
			MethodName: "JobSeekerLogin",
			Handler:    _JobSeeker_JobSeekerLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
