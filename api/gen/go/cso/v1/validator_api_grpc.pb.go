// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cso/v1/validator_api.proto

package csov1

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
	ValidatorService_Validate_FullMethodName = "/cso.v1.ValidatorService/Validate"
)

// ValidatorServiceClient is the client API for ValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidatorServiceClient interface {
	// Validate given path according to CSO sepcification.
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type validatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorServiceClient(cc grpc.ClientConnInterface) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, ValidatorService_Validate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServiceServer is the server API for ValidatorService service.
// All implementations must embed UnimplementedValidatorServiceServer
// for forward compatibility
type ValidatorServiceServer interface {
	// Validate given path according to CSO sepcification.
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedValidatorServiceServer()
}

// UnimplementedValidatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedValidatorServiceServer struct {
}

func (UnimplementedValidatorServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedValidatorServiceServer) mustEmbedUnimplementedValidatorServiceServer() {}

// UnsafeValidatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidatorServiceServer will
// result in compilation errors.
type UnsafeValidatorServiceServer interface {
	mustEmbedUnimplementedValidatorServiceServer()
}

func RegisterValidatorServiceServer(s grpc.ServiceRegistrar, srv ValidatorServiceServer) {
	s.RegisterService(&ValidatorService_ServiceDesc, srv)
}

func _ValidatorService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ValidatorService_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidatorService_ServiceDesc is the grpc.ServiceDesc for ValidatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cso.v1.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _ValidatorService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cso/v1/validator_api.proto",
}
