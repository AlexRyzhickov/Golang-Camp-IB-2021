// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: portal.api.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type SetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetInfoRequest) Reset() {
	*x = SetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetInfoRequest) ProtoMessage() {}

func (x *SetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetInfoRequest.ProtoReflect.Descriptor instead.
func (*SetInfoRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{1}
}

func (x *SetInfoRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *SetInfoRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SetInfoResponse) Reset() {
	*x = SetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetInfoResponse) ProtoMessage() {}

func (x *SetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetInfoResponse.ProtoReflect.Descriptor instead.
func (*SetInfoResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{2}
}

func (x *SetInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetInfoResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetUptimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetUptimeRequest) Reset() {
	*x = GetUptimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUptimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUptimeRequest) ProtoMessage() {}

func (x *GetUptimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUptimeRequest.ProtoReflect.Descriptor instead.
func (*GetUptimeRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetUptimeRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetUptimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetUptimeResponse) Reset() {
	*x = GetUptimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUptimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUptimeResponse) ProtoMessage() {}

func (x *GetUptimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUptimeResponse.ProtoReflect.Descriptor instead.
func (*GetUptimeResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetUptimeResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetRequestsRequest) Reset() {
	*x = GetRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestsRequest) ProtoMessage() {}

func (x *GetRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestsRequest.ProtoReflect.Descriptor instead.
func (*GetRequestsRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetRequestsRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetRequestsResponse) Reset() {
	*x = GetRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestsResponse) ProtoMessage() {}

func (x *GetRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestsResponse.ProtoReflect.Descriptor instead.
func (*GetRequestsResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetRequestsResponse) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{9}
}

func (x *ResetRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type ResetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResetResponse) Reset() {
	*x = ResetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetResponse) ProtoMessage() {}

func (x *ResetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetResponse.ProtoReflect.Descriptor instead.
func (*ResetResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{10}
}

func (x *ResetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *GetModeRequest) Reset() {
	*x = GetModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModeRequest) ProtoMessage() {}

func (x *GetModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModeRequest.ProtoReflect.Descriptor instead.
func (*GetModeRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{11}
}

func (x *GetModeRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type GetModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *GetModeResponse) Reset() {
	*x = GetModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModeResponse) ProtoMessage() {}

func (x *GetModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModeResponse.ProtoReflect.Descriptor instead.
func (*GetModeResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{12}
}

func (x *GetModeResponse) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type SetModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Mode    bool   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *SetModeRequest) Reset() {
	*x = SetModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetModeRequest) ProtoMessage() {}

func (x *SetModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetModeRequest.ProtoReflect.Descriptor instead.
func (*SetModeRequest) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{13}
}

func (x *SetModeRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *SetModeRequest) GetMode() bool {
	if x != nil {
		return x.Mode
	}
	return false
}

type SetModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SetModeResponse) Reset() {
	*x = SetModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetModeResponse) ProtoMessage() {}

func (x *SetModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portal_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetModeResponse.ProtoReflect.Descriptor instead.
func (*SetModeResponse) Descriptor() ([]byte, []int) {
	return file_portal_api_proto_rawDescGZIP(), []int{14}
}

func (x *SetModeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_portal_api_proto protoreflect.FileDescriptor

var file_portal_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x21, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xcd, 0x05, 0x0a,
	0x06, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x52, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x3a,
	0x01, 0x2a, 0x12, 0x55, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x67, 0x65, 0x74, 0x55,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c,
	0x2f, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x4d, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x55,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x67, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22,
	0x08, 0x2f, 0x73, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x15, 0x5a, 0x13,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portal_api_proto_rawDescOnce sync.Once
	file_portal_api_proto_rawDescData = file_portal_api_proto_rawDesc
)

func file_portal_api_proto_rawDescGZIP() []byte {
	file_portal_api_proto_rawDescOnce.Do(func() {
		file_portal_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_api_proto_rawDescData)
	})
	return file_portal_api_proto_rawDescData
}

var file_portal_api_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_portal_api_proto_goTypes = []interface{}{
	(*VersionResponse)(nil),     // 0: portal.pb.VersionResponse
	(*SetInfoRequest)(nil),      // 1: portal.pb.SetInfoRequest
	(*SetInfoResponse)(nil),     // 2: portal.pb.SetInfoResponse
	(*GetInfoRequest)(nil),      // 3: portal.pb.GetInfoRequest
	(*GetInfoResponse)(nil),     // 4: portal.pb.GetInfoResponse
	(*GetUptimeRequest)(nil),    // 5: portal.pb.GetUptimeRequest
	(*GetUptimeResponse)(nil),   // 6: portal.pb.GetUptimeResponse
	(*GetRequestsRequest)(nil),  // 7: portal.pb.GetRequestsRequest
	(*GetRequestsResponse)(nil), // 8: portal.pb.GetRequestsResponse
	(*ResetRequest)(nil),        // 9: portal.pb.ResetRequest
	(*ResetResponse)(nil),       // 10: portal.pb.ResetResponse
	(*GetModeRequest)(nil),      // 11: portal.pb.GetModeRequest
	(*GetModeResponse)(nil),     // 12: portal.pb.GetModeResponse
	(*SetModeRequest)(nil),      // 13: portal.pb.SetModeRequest
	(*SetModeResponse)(nil),     // 14: portal.pb.SetModeResponse
	(*empty.Empty)(nil),         // 15: google.protobuf.Empty
}
var file_portal_api_proto_depIdxs = []int32{
	15, // 0: portal.pb.Portal.GetVersion:input_type -> google.protobuf.Empty
	3,  // 1: portal.pb.Portal.GetInfo:input_type -> portal.pb.GetInfoRequest
	1,  // 2: portal.pb.Portal.SetInfo:input_type -> portal.pb.SetInfoRequest
	5,  // 3: portal.pb.Portal.GetUptime:input_type -> portal.pb.GetUptimeRequest
	7,  // 4: portal.pb.Portal.GetRequests:input_type -> portal.pb.GetRequestsRequest
	9,  // 5: portal.pb.Portal.Reset:input_type -> portal.pb.ResetRequest
	11, // 6: portal.pb.Portal.GetMode:input_type -> portal.pb.GetModeRequest
	13, // 7: portal.pb.Portal.SetMode:input_type -> portal.pb.SetModeRequest
	0,  // 8: portal.pb.Portal.GetVersion:output_type -> portal.pb.VersionResponse
	4,  // 9: portal.pb.Portal.GetInfo:output_type -> portal.pb.GetInfoResponse
	2,  // 10: portal.pb.Portal.SetInfo:output_type -> portal.pb.SetInfoResponse
	6,  // 11: portal.pb.Portal.GetUptime:output_type -> portal.pb.GetUptimeResponse
	8,  // 12: portal.pb.Portal.GetRequests:output_type -> portal.pb.GetRequestsResponse
	10, // 13: portal.pb.Portal.Reset:output_type -> portal.pb.ResetResponse
	12, // 14: portal.pb.Portal.GetMode:output_type -> portal.pb.GetModeResponse
	14, // 15: portal.pb.Portal.SetMode:output_type -> portal.pb.SetModeResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_portal_api_proto_init() }
func file_portal_api_proto_init() {
	if File_portal_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portal_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_portal_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetInfoRequest); i {
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
		file_portal_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetInfoResponse); i {
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
		file_portal_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_portal_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_portal_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUptimeRequest); i {
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
		file_portal_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUptimeResponse); i {
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
		file_portal_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestsRequest); i {
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
		file_portal_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestsResponse); i {
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
		file_portal_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_portal_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetResponse); i {
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
		file_portal_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModeRequest); i {
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
		file_portal_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModeResponse); i {
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
		file_portal_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetModeRequest); i {
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
		file_portal_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetModeResponse); i {
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
			RawDescriptor: file_portal_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portal_api_proto_goTypes,
		DependencyIndexes: file_portal_api_proto_depIdxs,
		MessageInfos:      file_portal_api_proto_msgTypes,
	}.Build()
	File_portal_api_proto = out.File
	file_portal_api_proto_rawDesc = nil
	file_portal_api_proto_goTypes = nil
	file_portal_api_proto_depIdxs = nil
}