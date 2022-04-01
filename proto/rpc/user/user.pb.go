// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/rpc/user/user.proto

package rpcv3

import (
	v31 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	v3 "github.com/RafayLabs/rcloud-base/proto/types/userpb/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApiKeyRequest) Reset() {
	*x = ApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyRequest) ProtoMessage() {}

func (x *ApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyRequest.ProtoReflect.Descriptor instead.
func (*ApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKeyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ApiKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id         string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Key        string                 `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name       string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ApiKeyResponse) Reset() {
	*x = ApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyResponse) ProtoMessage() {}

func (x *ApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyResponse.ProtoReflect.Descriptor instead.
func (*ApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *ApiKeyResponse) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *ApiKeyResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiKeyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiKeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiKeyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ApiKeyResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ApiKeyResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApiKeyResponseList) Reset() {
	*x = ApiKeyResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyResponseList) ProtoMessage() {}

func (x *ApiKeyResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyResponseList.ProtoReflect.Descriptor instead.
func (*ApiKeyResponseList) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *ApiKeyResponseList) GetItems() []*ApiKeyResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{3}
}

type CliConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CliConfigRequest) Reset() {
	*x = CliConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliConfigRequest) ProtoMessage() {}

func (x *CliConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliConfigRequest.ProtoReflect.Descriptor instead.
func (*CliConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{4}
}

var File_proto_rpc_user_user_proto protoreflect.FileDescriptor

var file_proto_rpc_user_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62,
	0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4c, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xa9, 0x09, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x9e, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x52,
	0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x69, 0x73, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22,
	0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x27,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x6e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x21, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x74,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x72, 0x61,
	0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x1a, 0x1d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xb1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x24, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a,
	0x03, 0x32, 0x30, 0x34, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x79, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1f, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1f,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0xa1, 0x04, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66, 0x61, 0x79, 0x4c, 0x61, 0x62, 0x73,
	0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x72, 0x70, 0x63, 0x76,
	0x33, 0xa2, 0x02, 0x03, 0x52, 0x44, 0x52, 0xaa, 0x02, 0x10, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e,
	0x44, 0x65, 0x76, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x10, 0x52, 0x61, 0x66,
	0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x1c,
	0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x56, 0x33,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a,
	0x56, 0x33, 0x92, 0x41, 0xe2, 0x02, 0x12, 0x2b, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x61, 0x79, 0x20, 0x44, 0x65, 0x76, 0x32, 0x03,
	0x32, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x52, 0x50,
	0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x38, 0x0a,
	0x25, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x08,
	0x02, 0x1a, 0x11, 0x58, 0x2d, 0x52, 0x41, 0x46, 0x41, 0x59, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b,
	0x45, 0x59, 0x49, 0x44, 0x20, 0x02, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_user_user_proto_rawDescOnce sync.Once
	file_proto_rpc_user_user_proto_rawDescData = file_proto_rpc_user_user_proto_rawDesc
)

func file_proto_rpc_user_user_proto_rawDescGZIP() []byte {
	file_proto_rpc_user_user_proto_rawDescOnce.Do(func() {
		file_proto_rpc_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_user_user_proto_rawDescData)
	})
	return file_proto_rpc_user_user_proto_rawDescData
}

var file_proto_rpc_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_rpc_user_user_proto_goTypes = []interface{}{
	(*ApiKeyRequest)(nil),         // 0: rafay.dev.rpc.v3.ApiKeyRequest
	(*ApiKeyResponse)(nil),        // 1: rafay.dev.rpc.v3.ApiKeyResponse
	(*ApiKeyResponseList)(nil),    // 2: rafay.dev.rpc.v3.ApiKeyResponseList
	(*DeleteUserResponse)(nil),    // 3: rafay.dev.rpc.v3.DeleteUserResponse
	(*CliConfigRequest)(nil),      // 4: rafay.dev.rpc.v3.CliConfigRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*v3.User)(nil),               // 6: rafay.dev.types.user.v3.User
	(*v31.QueryOptions)(nil),      // 7: rafay.dev.types.common.v3.QueryOptions
	(*v3.UserList)(nil),           // 8: rafay.dev.types.user.v3.UserList
	(*v3.UserInfo)(nil),           // 9: rafay.dev.types.user.v3.UserInfo
	(*v31.HttpBody)(nil),          // 10: rafay.dev.types.common.v3.HttpBody
}
var file_proto_rpc_user_user_proto_depIdxs = []int32{
	5,  // 0: rafay.dev.rpc.v3.ApiKeyResponse.modifiedAt:type_name -> google.protobuf.Timestamp
	5,  // 1: rafay.dev.rpc.v3.ApiKeyResponse.createdAt:type_name -> google.protobuf.Timestamp
	1,  // 2: rafay.dev.rpc.v3.ApiKeyResponseList.items:type_name -> rafay.dev.rpc.v3.ApiKeyResponse
	6,  // 3: rafay.dev.rpc.v3.User.CreateUser:input_type -> rafay.dev.types.user.v3.User
	7,  // 4: rafay.dev.rpc.v3.User.GetUsers:input_type -> rafay.dev.types.common.v3.QueryOptions
	6,  // 5: rafay.dev.rpc.v3.User.GetUser:input_type -> rafay.dev.types.user.v3.User
	6,  // 6: rafay.dev.rpc.v3.User.GetUserInfo:input_type -> rafay.dev.types.user.v3.User
	6,  // 7: rafay.dev.rpc.v3.User.UpdateUser:input_type -> rafay.dev.types.user.v3.User
	6,  // 8: rafay.dev.rpc.v3.User.DeleteUser:input_type -> rafay.dev.types.user.v3.User
	4,  // 9: rafay.dev.rpc.v3.User.DownloadCliConfig:input_type -> rafay.dev.rpc.v3.CliConfigRequest
	0,  // 10: rafay.dev.rpc.v3.User.UserListApiKeys:input_type -> rafay.dev.rpc.v3.ApiKeyRequest
	0,  // 11: rafay.dev.rpc.v3.User.UserDeleteApiKeys:input_type -> rafay.dev.rpc.v3.ApiKeyRequest
	6,  // 12: rafay.dev.rpc.v3.User.CreateUser:output_type -> rafay.dev.types.user.v3.User
	8,  // 13: rafay.dev.rpc.v3.User.GetUsers:output_type -> rafay.dev.types.user.v3.UserList
	6,  // 14: rafay.dev.rpc.v3.User.GetUser:output_type -> rafay.dev.types.user.v3.User
	9,  // 15: rafay.dev.rpc.v3.User.GetUserInfo:output_type -> rafay.dev.types.user.v3.UserInfo
	6,  // 16: rafay.dev.rpc.v3.User.UpdateUser:output_type -> rafay.dev.types.user.v3.User
	3,  // 17: rafay.dev.rpc.v3.User.DeleteUser:output_type -> rafay.dev.rpc.v3.DeleteUserResponse
	10, // 18: rafay.dev.rpc.v3.User.DownloadCliConfig:output_type -> rafay.dev.types.common.v3.HttpBody
	2,  // 19: rafay.dev.rpc.v3.User.UserListApiKeys:output_type -> rafay.dev.rpc.v3.ApiKeyResponseList
	3,  // 20: rafay.dev.rpc.v3.User.UserDeleteApiKeys:output_type -> rafay.dev.rpc.v3.DeleteUserResponse
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rpc_user_user_proto_init() }
func file_proto_rpc_user_user_proto_init() {
	if File_proto_rpc_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyRequest); i {
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
		file_proto_rpc_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyResponseList); i {
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
		file_proto_rpc_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliConfigRequest); i {
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
			RawDescriptor: file_proto_rpc_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_user_user_proto_goTypes,
		DependencyIndexes: file_proto_rpc_user_user_proto_depIdxs,
		MessageInfos:      file_proto_rpc_user_user_proto_msgTypes,
	}.Build()
	File_proto_rpc_user_user_proto = out.File
	file_proto_rpc_user_user_proto_rawDesc = nil
	file_proto_rpc_user_user_proto_goTypes = nil
	file_proto_rpc_user_user_proto_depIdxs = nil
}
