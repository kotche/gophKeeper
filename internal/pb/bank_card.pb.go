// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal/proto/bank_card.proto

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

type BankCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Number   string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	MetaInfo string `protobuf:"bytes,4,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *BankCardRequest) Reset() {
	*x = BankCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardRequest) ProtoMessage() {}

func (x *BankCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardRequest.ProtoReflect.Descriptor instead.
func (*BankCardRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{0}
}

func (x *BankCardRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BankCardRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *BankCardRequest) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BankCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BankCardResponse) Reset() {
	*x = BankCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardResponse) ProtoMessage() {}

func (x *BankCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardResponse.ProtoReflect.Descriptor instead.
func (*BankCardResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{1}
}

func (x *BankCardResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BankCardUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Number   string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	MetaInfo string `protobuf:"bytes,4,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *BankCardUpdateRequest) Reset() {
	*x = BankCardUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardUpdateRequest) ProtoMessage() {}

func (x *BankCardUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardUpdateRequest.ProtoReflect.Descriptor instead.
func (*BankCardUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{2}
}

func (x *BankCardUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BankCardUpdateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BankCardUpdateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *BankCardUpdateRequest) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BankCardUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BankCardUpdateResponse) Reset() {
	*x = BankCardUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardUpdateResponse) ProtoMessage() {}

func (x *BankCardUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardUpdateResponse.ProtoReflect.Descriptor instead.
func (*BankCardUpdateResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{3}
}

type BankCardDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BankCardDeleteRequest) Reset() {
	*x = BankCardDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardDeleteRequest) ProtoMessage() {}

func (x *BankCardDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardDeleteRequest.ProtoReflect.Descriptor instead.
func (*BankCardDeleteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{4}
}

func (x *BankCardDeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BankCardDeleteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BankCardDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BankCardDeleteResponse) Reset() {
	*x = BankCardDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardDeleteResponse) ProtoMessage() {}

func (x *BankCardDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardDeleteResponse.ProtoReflect.Descriptor instead.
func (*BankCardDeleteResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{5}
}

type BankCardGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BankCardGetAllRequest) Reset() {
	*x = BankCardGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardGetAllRequest) ProtoMessage() {}

func (x *BankCardGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardGetAllRequest.ProtoReflect.Descriptor instead.
func (*BankCardGetAllRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{6}
}

func (x *BankCardGetAllRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllBankCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number   string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	MetaInfo string `protobuf:"bytes,3,opt,name=meta_info,json=metaInfo,proto3" json:"meta_info,omitempty"`
}

func (x *GetAllBankCardResponse) Reset() {
	*x = GetAllBankCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBankCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBankCardResponse) ProtoMessage() {}

func (x *GetAllBankCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBankCardResponse.ProtoReflect.Descriptor instead.
func (*GetAllBankCardResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllBankCardResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAllBankCardResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *GetAllBankCardResponse) GetMetaInfo() string {
	if x != nil {
		return x.MetaInfo
	}
	return ""
}

type BankCardGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankCards []*GetAllBankCardResponse `protobuf:"bytes,1,rep,name=bank_cards,json=bankCards,proto3" json:"bank_cards,omitempty"`
}

func (x *BankCardGetAllResponse) Reset() {
	*x = BankCardGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_bank_card_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardGetAllResponse) ProtoMessage() {}

func (x *BankCardGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_bank_card_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardGetAllResponse.ProtoReflect.Descriptor instead.
func (*BankCardGetAllResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_bank_card_proto_rawDescGZIP(), []int{8}
}

func (x *BankCardGetAllResponse) GetBankCards() []*GetAllBankCardResponse {
	if x != nil {
		return x.BankCards
	}
	return nil
}

var File_internal_proto_bank_card_proto protoreflect.FileDescriptor

var file_internal_proto_bank_card_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x6b,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x42, 0x61, 0x6e,
	0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a,
	0x15, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40,
	0x0a, 0x15, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x15, 0x42, 0x61,
	0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57, 0x0a, 0x16, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x1d, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1d, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61,
	0x72, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x1d, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43,
	0x61, 0x72, 0x64, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61,
	0x72, 0x64, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_bank_card_proto_rawDescOnce sync.Once
	file_internal_proto_bank_card_proto_rawDescData = file_internal_proto_bank_card_proto_rawDesc
)

func file_internal_proto_bank_card_proto_rawDescGZIP() []byte {
	file_internal_proto_bank_card_proto_rawDescOnce.Do(func() {
		file_internal_proto_bank_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_bank_card_proto_rawDescData)
	})
	return file_internal_proto_bank_card_proto_rawDescData
}

var file_internal_proto_bank_card_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_bank_card_proto_goTypes = []interface{}{
	(*BankCardRequest)(nil),        // 0: keeper.BankCardRequest
	(*BankCardResponse)(nil),       // 1: keeper.BankCardResponse
	(*BankCardUpdateRequest)(nil),  // 2: keeper.BankCardUpdateRequest
	(*BankCardUpdateResponse)(nil), // 3: keeper.BankCardUpdateResponse
	(*BankCardDeleteRequest)(nil),  // 4: keeper.BankCardDeleteRequest
	(*BankCardDeleteResponse)(nil), // 5: keeper.BankCardDeleteResponse
	(*BankCardGetAllRequest)(nil),  // 6: keeper.BankCardGetAllRequest
	(*GetAllBankCardResponse)(nil), // 7: keeper.GetAllBankCardResponse
	(*BankCardGetAllResponse)(nil), // 8: keeper.BankCardGetAllResponse
}
var file_internal_proto_bank_card_proto_depIdxs = []int32{
	7, // 0: keeper.BankCardGetAllResponse.bank_cards:type_name -> keeper.GetAllBankCardResponse
	0, // 1: keeper.BankCardService.CreateBankCard:input_type -> keeper.BankCardRequest
	2, // 2: keeper.BankCardService.UpdateBankCard:input_type -> keeper.BankCardUpdateRequest
	4, // 3: keeper.BankCardService.DeleteBankCard:input_type -> keeper.BankCardDeleteRequest
	6, // 4: keeper.BankCardService.GetAllBankCard:input_type -> keeper.BankCardGetAllRequest
	1, // 5: keeper.BankCardService.CreateBankCard:output_type -> keeper.BankCardResponse
	3, // 6: keeper.BankCardService.UpdateBankCard:output_type -> keeper.BankCardUpdateResponse
	5, // 7: keeper.BankCardService.DeleteBankCard:output_type -> keeper.BankCardDeleteResponse
	8, // 8: keeper.BankCardService.GetAllBankCard:output_type -> keeper.BankCardGetAllResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_proto_bank_card_proto_init() }
func file_internal_proto_bank_card_proto_init() {
	if File_internal_proto_bank_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_bank_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardRequest); i {
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
		file_internal_proto_bank_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardResponse); i {
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
		file_internal_proto_bank_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardUpdateRequest); i {
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
		file_internal_proto_bank_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardUpdateResponse); i {
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
		file_internal_proto_bank_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardDeleteRequest); i {
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
		file_internal_proto_bank_card_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardDeleteResponse); i {
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
		file_internal_proto_bank_card_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardGetAllRequest); i {
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
		file_internal_proto_bank_card_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBankCardResponse); i {
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
		file_internal_proto_bank_card_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardGetAllResponse); i {
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
			RawDescriptor: file_internal_proto_bank_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_bank_card_proto_goTypes,
		DependencyIndexes: file_internal_proto_bank_card_proto_depIdxs,
		MessageInfos:      file_internal_proto_bank_card_proto_msgTypes,
	}.Build()
	File_internal_proto_bank_card_proto = out.File
	file_internal_proto_bank_card_proto_rawDesc = nil
	file_internal_proto_bank_card_proto_goTypes = nil
	file_internal_proto_bank_card_proto_depIdxs = nil
}
