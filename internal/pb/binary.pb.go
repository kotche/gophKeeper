// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/proto/binary.proto

package pb

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

type BinaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Binary   string `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	MetaInfo string `protobuf:"bytes,3,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *BinaryRequest) Reset() {
	*x = BinaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryRequest) ProtoMessage() {}

func (x *BinaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryRequest.ProtoReflect.Descriptor instead.
func (*BinaryRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{0}
}

func (x *BinaryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BinaryRequest) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

func (x *BinaryRequest) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BinaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BinaryResponse) Reset() {
	*x = BinaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryResponse) ProtoMessage() {}

func (x *BinaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryResponse.ProtoReflect.Descriptor instead.
func (*BinaryResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BinaryUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Binary   string `protobuf:"bytes,3,opt,name=binary,proto3" json:"binary,omitempty"`
	MetaInfo string `protobuf:"bytes,4,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *BinaryUpdateRequest) Reset() {
	*x = BinaryUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryUpdateRequest) ProtoMessage() {}

func (x *BinaryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryUpdateRequest.ProtoReflect.Descriptor instead.
func (*BinaryUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{2}
}

func (x *BinaryUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BinaryUpdateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BinaryUpdateRequest) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

func (x *BinaryUpdateRequest) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BinaryUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinaryUpdateResponse) Reset() {
	*x = BinaryUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryUpdateResponse) ProtoMessage() {}

func (x *BinaryUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryUpdateResponse.ProtoReflect.Descriptor instead.
func (*BinaryUpdateResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{3}
}

type BinaryDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BinaryDeleteRequest) Reset() {
	*x = BinaryDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDeleteRequest) ProtoMessage() {}

func (x *BinaryDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDeleteRequest.ProtoReflect.Descriptor instead.
func (*BinaryDeleteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{4}
}

func (x *BinaryDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BinaryDeleteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BinaryDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinaryDeleteResponse) Reset() {
	*x = BinaryDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryDeleteResponse) ProtoMessage() {}

func (x *BinaryDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryDeleteResponse.ProtoReflect.Descriptor instead.
func (*BinaryDeleteResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{5}
}

type BinaryGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BinaryGetAllRequest) Reset() {
	*x = BinaryGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryGetAllRequest) ProtoMessage() {}

func (x *BinaryGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryGetAllRequest.ProtoReflect.Descriptor instead.
func (*BinaryGetAllRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{6}
}

func (x *BinaryGetAllRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllBinaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Binary   string `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	MetaInfo string `protobuf:"bytes,3,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *GetAllBinaryResponse) Reset() {
	*x = GetAllBinaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBinaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBinaryResponse) ProtoMessage() {}

func (x *GetAllBinaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBinaryResponse.ProtoReflect.Descriptor instead.
func (*GetAllBinaryResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllBinaryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAllBinaryResponse) GetBinary() string {
	if x != nil {
		return x.Binary
	}
	return ""
}

func (x *GetAllBinaryResponse) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BinaryGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binary []*GetAllBinaryResponse `protobuf:"bytes,1,rep,name=binary,proto3" json:"binary,omitempty"`
}

func (x *BinaryGetAllResponse) Reset() {
	*x = BinaryGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_binary_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryGetAllResponse) ProtoMessage() {}

func (x *BinaryGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_binary_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryGetAllResponse.ProtoReflect.Descriptor instead.
func (*BinaryGetAllResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_binary_proto_rawDescGZIP(), []int{8}
}

func (x *BinaryGetAllResponse) GetBinary() []*GetAllBinaryResponse {
	if x != nil {
		return x.Binary
	}
	return nil
}

var File_internal_proto_binary_proto protoreflect.FileDescriptor

var file_internal_proto_binary_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x13, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x13, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x13, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x14, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x32, 0xaf, 0x02, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x1b, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x2e,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_binary_proto_rawDescOnce sync.Once
	file_internal_proto_binary_proto_rawDescData = file_internal_proto_binary_proto_rawDesc
)

func file_internal_proto_binary_proto_rawDescGZIP() []byte {
	file_internal_proto_binary_proto_rawDescOnce.Do(func() {
		file_internal_proto_binary_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_binary_proto_rawDescData)
	})
	return file_internal_proto_binary_proto_rawDescData
}

var file_internal_proto_binary_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_binary_proto_goTypes = []interface{}{
	(*BinaryRequest)(nil),        // 0: keeper.BinaryRequest
	(*BinaryResponse)(nil),       // 1: keeper.BinaryResponse
	(*BinaryUpdateRequest)(nil),  // 2: keeper.BinaryUpdateRequest
	(*BinaryUpdateResponse)(nil), // 3: keeper.BinaryUpdateResponse
	(*BinaryDeleteRequest)(nil),  // 4: keeper.BinaryDeleteRequest
	(*BinaryDeleteResponse)(nil), // 5: keeper.BinaryDeleteResponse
	(*BinaryGetAllRequest)(nil),  // 6: keeper.BinaryGetAllRequest
	(*GetAllBinaryResponse)(nil), // 7: keeper.GetAllBinaryResponse
	(*BinaryGetAllResponse)(nil), // 8: keeper.BinaryGetAllResponse
}
var file_internal_proto_binary_proto_depIdxs = []int32{
	7, // 0: keeper.BinaryGetAllResponse.binary:type_name -> keeper.GetAllBinaryResponse
	0, // 1: keeper.BinaryService.CreateBinary:input_type -> keeper.BinaryRequest
	2, // 2: keeper.BinaryService.UpdateBinary:input_type -> keeper.BinaryUpdateRequest
	4, // 3: keeper.BinaryService.DeleteBinary:input_type -> keeper.BinaryDeleteRequest
	6, // 4: keeper.BinaryService.GetAllBinary:input_type -> keeper.BinaryGetAllRequest
	1, // 5: keeper.BinaryService.CreateBinary:output_type -> keeper.BinaryResponse
	3, // 6: keeper.BinaryService.UpdateBinary:output_type -> keeper.BinaryUpdateResponse
	5, // 7: keeper.BinaryService.DeleteBinary:output_type -> keeper.BinaryDeleteResponse
	8, // 8: keeper.BinaryService.GetAllBinary:output_type -> keeper.BinaryGetAllResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_binary_proto_init() }
func file_internal_proto_binary_proto_init() {
	if File_internal_proto_binary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_binary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryRequest); i {
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
		file_internal_proto_binary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryResponse); i {
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
		file_internal_proto_binary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryUpdateRequest); i {
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
		file_internal_proto_binary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryUpdateResponse); i {
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
		file_internal_proto_binary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryDeleteRequest); i {
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
		file_internal_proto_binary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryDeleteResponse); i {
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
		file_internal_proto_binary_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryGetAllRequest); i {
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
		file_internal_proto_binary_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBinaryResponse); i {
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
		file_internal_proto_binary_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryGetAllResponse); i {
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
			RawDescriptor: file_internal_proto_binary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_binary_proto_goTypes,
		DependencyIndexes: file_internal_proto_binary_proto_depIdxs,
		MessageInfos:      file_internal_proto_binary_proto_msgTypes,
	}.Build()
	File_internal_proto_binary_proto = out.File
	file_internal_proto_binary_proto_rawDesc = nil
	file_internal_proto_binary_proto_goTypes = nil
	file_internal_proto_binary_proto_depIdxs = nil
}
