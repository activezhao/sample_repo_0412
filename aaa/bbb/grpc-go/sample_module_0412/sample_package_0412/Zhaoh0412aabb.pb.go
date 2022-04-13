// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: Zhaoh0412aabb.proto

package sample_package_0412

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthcheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HealthcheckRequest) Reset() {
	*x = HealthcheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zhaoh0412aabb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckRequest) ProtoMessage() {}

func (x *HealthcheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Zhaoh0412aabb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckRequest.ProtoReflect.Descriptor instead.
func (*HealthcheckRequest) Descriptor() ([]byte, []int) {
	return file_Zhaoh0412aabb_proto_rawDescGZIP(), []int{0}
}

func (x *HealthcheckRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HealthcheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HealthcheckResponse) Reset() {
	*x = HealthcheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zhaoh0412aabb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthcheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthcheckResponse) ProtoMessage() {}

func (x *HealthcheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Zhaoh0412aabb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthcheckResponse.ProtoReflect.Descriptor instead.
func (*HealthcheckResponse) Descriptor() ([]byte, []int) {
	return file_Zhaoh0412aabb_proto_rawDescGZIP(), []int{1}
}

func (x *HealthcheckResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zhaoh0412aabb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Zhaoh0412aabb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_Zhaoh0412aabb_proto_rawDescGZIP(), []int{2}
}

func (x *HelloRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Zhaoh0412aabb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_Zhaoh0412aabb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_Zhaoh0412aabb_proto_rawDescGZIP(), []int{3}
}

func (x *HelloReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *HelloReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_Zhaoh0412aabb_proto protoreflect.FileDescriptor

var file_Zhaoh0412aabb_proto_rawDesc = []byte{
	0x0a, 0x13, 0x5a, 0x68, 0x61, 0x6f, 0x68, 0x30, 0x34, 0x31, 0x32, 0x61, 0x61, 0x62, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x1a, 0x92, 0x41,
	0x17, 0x32, 0x15, 0x7b, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x3a, 0x20, 0x22, 0x68, 0x65, 0x61, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x7d, 0x22, 0x44, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x3a, 0x15, 0x92, 0x41, 0x12, 0x32, 0x10, 0x7b, 0x22,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x6f, 0x6b, 0x22, 0x7d, 0x22, 0x46,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x03, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x1b, 0x92, 0x41, 0x18, 0x32, 0x16,
	0x7b, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x3a, 0x20, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x22, 0x7d, 0x22, 0x67, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x33, 0x92, 0x41, 0x30, 0x32,
	0x2e, 0x7b, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x3a, 0x20, 0x22, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x22, 0x30, 0x22, 0x7d, 0x32,
	0xc6, 0x02, 0x0a, 0x0d, 0x5a, 0x68, 0x61, 0x6f, 0x68, 0x30, 0x34, 0x31, 0x32, 0x61, 0x61, 0x62,
	0x62, 0x12, 0x9f, 0x01, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x3a, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x30, 0x34, 0x31,
	0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x92, 0x01, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x12, 0x34, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x30, 0x34, 0x31,
	0x32, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x7a, 0x68, 0x61,
	0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x30, 0x34,
	0x31, 0x32, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x62, 0x62, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x30, 0x34, 0x31, 0x32, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x30, 0x34, 0x31, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Zhaoh0412aabb_proto_rawDescOnce sync.Once
	file_Zhaoh0412aabb_proto_rawDescData = file_Zhaoh0412aabb_proto_rawDesc
)

func file_Zhaoh0412aabb_proto_rawDescGZIP() []byte {
	file_Zhaoh0412aabb_proto_rawDescOnce.Do(func() {
		file_Zhaoh0412aabb_proto_rawDescData = protoimpl.X.CompressGZIP(file_Zhaoh0412aabb_proto_rawDescData)
	})
	return file_Zhaoh0412aabb_proto_rawDescData
}

var file_Zhaoh0412aabb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Zhaoh0412aabb_proto_goTypes = []interface{}{
	(*HealthcheckRequest)(nil),  // 0: sample_module_0412.sample_package_0412.HealthcheckRequest
	(*HealthcheckResponse)(nil), // 1: sample_module_0412.sample_package_0412.HealthcheckResponse
	(*HelloRequest)(nil),        // 2: sample_module_0412.sample_package_0412.HelloRequest
	(*HelloReply)(nil),          // 3: sample_module_0412.sample_package_0412.HelloReply
}
var file_Zhaoh0412aabb_proto_depIdxs = []int32{
	0, // 0: sample_module_0412.sample_package_0412.Zhaoh0412aabb.Heathcheck:input_type -> sample_module_0412.sample_package_0412.HealthcheckRequest
	2, // 1: sample_module_0412.sample_package_0412.Zhaoh0412aabb.Helloworld:input_type -> sample_module_0412.sample_package_0412.HelloRequest
	1, // 2: sample_module_0412.sample_package_0412.Zhaoh0412aabb.Heathcheck:output_type -> sample_module_0412.sample_package_0412.HealthcheckResponse
	3, // 3: sample_module_0412.sample_package_0412.Zhaoh0412aabb.Helloworld:output_type -> sample_module_0412.sample_package_0412.HelloReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Zhaoh0412aabb_proto_init() }
func file_Zhaoh0412aabb_proto_init() {
	if File_Zhaoh0412aabb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Zhaoh0412aabb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckRequest); i {
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
		file_Zhaoh0412aabb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthcheckResponse); i {
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
		file_Zhaoh0412aabb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_Zhaoh0412aabb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_Zhaoh0412aabb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Zhaoh0412aabb_proto_goTypes,
		DependencyIndexes: file_Zhaoh0412aabb_proto_depIdxs,
		MessageInfos:      file_Zhaoh0412aabb_proto_msgTypes,
	}.Build()
	File_Zhaoh0412aabb_proto = out.File
	file_Zhaoh0412aabb_proto_rawDesc = nil
	file_Zhaoh0412aabb_proto_goTypes = nil
	file_Zhaoh0412aabb_proto_depIdxs = nil
}
