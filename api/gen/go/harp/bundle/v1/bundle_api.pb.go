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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: harp/bundle/v1/bundle_api.proto

package bundlev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetSecretRequest describes information required to retrieve secret from
// container server.
type GetSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namepace name.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Secret path.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetSecretRequest) Reset() {
	*x = GetSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRequest) ProtoMessage() {}

func (x *GetSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRequest.ProtoReflect.Descriptor instead.
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return file_harp_bundle_v1_bundle_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetSecretRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetSecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace name.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Secret path.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Secret content returned by mapped engine.
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetSecretResponse) Reset() {
	*x = GetSecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponse) ProtoMessage() {}

func (x *GetSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponse.ProtoReflect.Descriptor instead.
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return file_harp_bundle_v1_bundle_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetSecretResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetSecretResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetSecretResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_harp_bundle_v1_bundle_api_proto protoreflect.FileDescriptor

var file_harp_bundle_v1_bundle_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x61, 0x0a, 0x0d, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbc, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x6e, 0x74, 0x72, 0x69, 0x6f, 0x2f, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x61, 0x72, 0x70, 0x2f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x42, 0x58, 0xaa, 0x02, 0x0e, 0x48, 0x61, 0x72, 0x70, 0x2e,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x48, 0x61, 0x72, 0x70,
	0x5c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x48, 0x61, 0x72,
	0x70, 0x5c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x48, 0x61, 0x72, 0x70, 0x3a, 0x3a,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_harp_bundle_v1_bundle_api_proto_rawDescOnce sync.Once
	file_harp_bundle_v1_bundle_api_proto_rawDescData = file_harp_bundle_v1_bundle_api_proto_rawDesc
)

func file_harp_bundle_v1_bundle_api_proto_rawDescGZIP() []byte {
	file_harp_bundle_v1_bundle_api_proto_rawDescOnce.Do(func() {
		file_harp_bundle_v1_bundle_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_harp_bundle_v1_bundle_api_proto_rawDescData)
	})
	return file_harp_bundle_v1_bundle_api_proto_rawDescData
}

var file_harp_bundle_v1_bundle_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_harp_bundle_v1_bundle_api_proto_goTypes = []interface{}{
	(*GetSecretRequest)(nil),  // 0: harp.bundle.v1.GetSecretRequest
	(*GetSecretResponse)(nil), // 1: harp.bundle.v1.GetSecretResponse
}
var file_harp_bundle_v1_bundle_api_proto_depIdxs = []int32{
	0, // 0: harp.bundle.v1.BundleService.GetSecret:input_type -> harp.bundle.v1.GetSecretRequest
	1, // 1: harp.bundle.v1.BundleService.GetSecret:output_type -> harp.bundle.v1.GetSecretResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_harp_bundle_v1_bundle_api_proto_init() }
func file_harp_bundle_v1_bundle_api_proto_init() {
	if File_harp_bundle_v1_bundle_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_harp_bundle_v1_bundle_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_harp_bundle_v1_bundle_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_harp_bundle_v1_bundle_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_harp_bundle_v1_bundle_api_proto_goTypes,
		DependencyIndexes: file_harp_bundle_v1_bundle_api_proto_depIdxs,
		MessageInfos:      file_harp_bundle_v1_bundle_api_proto_msgTypes,
	}.Build()
	File_harp_bundle_v1_bundle_api_proto = out.File
	file_harp_bundle_v1_bundle_api_proto_rawDesc = nil
	file_harp_bundle_v1_bundle_api_proto_goTypes = nil
	file_harp_bundle_v1_bundle_api_proto_depIdxs = nil
}
