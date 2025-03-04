// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: chatsvc/v1/user.api.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/google/gnostic/openapiv3"
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

type UserVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Nickname  string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender    int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Signature string `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	IsOnline  bool   `protobuf:"varint,7,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
}

func (x *UserVO) Reset() {
	*x = UserVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVO) ProtoMessage() {}

func (x *UserVO) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVO.ProtoReflect.Descriptor instead.
func (*UserVO) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserVO) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserVO) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserVO) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserVO) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserVO) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserVO) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UserVO) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

type GetUserSelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserSelfRequest) Reset() {
	*x = GetUserSelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSelfRequest) ProtoMessage() {}

func (x *GetUserSelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSelfRequest.ProtoReflect.Descriptor instead.
func (*GetUserSelfRequest) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{1}
}

type GetUserSelfReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *UserVO `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserSelfReply) Reset() {
	*x = GetUserSelfReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSelfReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSelfReply) ProtoMessage() {}

func (x *GetUserSelfReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSelfReply.ProtoReflect.Descriptor instead.
func (*GetUserSelfReply) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserSelfReply) GetData() *UserVO {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname  string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender    int32  `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserInfoRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UpdateUserInfoRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type UpdateUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserInfoReply) Reset() {
	*x = UpdateUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoReply) ProtoMessage() {}

func (x *UpdateUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoReply.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoReply) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{4}
}

type GetUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIDRequest) Reset() {
	*x = GetUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRequest) ProtoMessage() {}

func (x *GetUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *UserVO `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserByIDReply) Reset() {
	*x = GetUserByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDReply) ProtoMessage() {}

func (x *GetUserByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDReply.ProtoReflect.Descriptor instead.
func (*GetUserByIDReply) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByIDReply) GetData() *UserVO {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListUserInfoRequest) Reset() {
	*x = ListUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserInfoRequest) ProtoMessage() {}

func (x *ListUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserInfoRequest.ProtoReflect.Descriptor instead.
func (*ListUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListUserInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserVO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListUserInfoReply) Reset() {
	*x = ListUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserInfoReply) ProtoMessage() {}

func (x *ListUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserInfoReply.ProtoReflect.Descriptor instead.
func (*ListUserInfoReply) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListUserInfoReply) GetData() []*UserVO {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdatePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type UpdatePasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePasswordReply) Reset() {
	*x = UpdatePasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_v1_user_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordReply) ProtoMessage() {}

func (x *UpdatePasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_v1_user_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordReply.ProtoReflect.Descriptor instead.
func (*UpdatePasswordReply) Descriptor() ([]byte, []int) {
	return file_chatsvc_v1_user_api_proto_rawDescGZIP(), []int{10}
}

var File_chatsvc_v1_user_api_proto protoreflect.FileDescriptor

var file_chatsvc_v1_user_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x01, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xca, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0xa3, 0x01, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x86, 0x01, 0xba, 0x47, 0x09, 0x92, 0x02, 0x06, 0xe6, 0x98, 0xb5,
	0xe7, 0xa7, 0xb0, 0xba, 0x48, 0x77, 0xba, 0x01, 0x74, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x5c, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29, 0x20, 0x3c, 0x20,
	0x32, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29,
	0x20, 0x3e, 0x20, 0x31, 0x30, 0x29, 0x20, 0x3f, 0x20, 0x27, 0xe6, 0x98, 0xb5, 0xe7, 0xa7, 0xb0,
	0xe9, 0x95, 0xbf, 0xe5, 0xba, 0xa6, 0xe4, 0xb8, 0x8d, 0xe5, 0xbe, 0x97, 0xe5, 0xb0, 0x8f, 0xe4,
	0xba, 0x8e, 0x32, 0xe6, 0x88, 0x96, 0xe5, 0xa4, 0xa7, 0xe4, 0xba, 0x8e, 0x31, 0x30, 0xe4, 0xb8,
	0xaa, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x75, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x5d, 0xba, 0x47, 0x09, 0x92, 0x02, 0x06, 0xe6,
	0x80, 0xa7, 0xe5, 0x88, 0xab, 0xba, 0x48, 0x4e, 0xba, 0x01, 0x4b, 0x0a, 0x12, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a,
	0x35, 0x28, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3c, 0x20, 0x30, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x3e, 0x20, 0x32, 0x29, 0x20, 0x3f, 0x20, 0x27, 0xe6, 0x80, 0xa7, 0xe5, 0x88,
	0xab, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0xe5, 0x80, 0xbc, 0xe6, 0x97, 0xa0, 0xe6, 0x95, 0x88,
	0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x93,
	0x01, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x75, 0xba, 0x47, 0x0f, 0x92, 0x02, 0x0c, 0xe4, 0xb8, 0xaa, 0xe6, 0x80, 0xa7,
	0xe7, 0xad, 0xbe, 0xe5, 0x90, 0x8d, 0xba, 0x48, 0x60, 0xba, 0x01, 0x5d, 0x0a, 0x15, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x1a, 0x44, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29,
	0x20, 0x3e, 0x20, 0x32, 0x30, 0x30, 0x20, 0x3f, 0x20, 0x27, 0xe4, 0xb8, 0xaa, 0xe6, 0x80, 0xa7,
	0xe7, 0xad, 0xbe, 0xe5, 0x90, 0x8d, 0xe9, 0x95, 0xbf, 0xe5, 0xba, 0xa6, 0xe4, 0xb8, 0x8d, 0xe5,
	0xbe, 0x97, 0xe8, 0xb6, 0x85, 0xe8, 0xbf, 0x87, 0x20, 0x32, 0x30, 0x30, 0x20, 0xe4, 0xb8, 0xaa,
	0xe5, 0xad, 0x97, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x01,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x78, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x64, 0xba, 0x47, 0x13, 0x92, 0x02, 0x10, 0xe7, 0x94, 0xa8, 0xe6, 0x88,
	0xb7, 0xe5, 0x90, 0x8d, 0x2f, 0xe6, 0x98, 0xb5, 0xe7, 0xa7, 0xb0, 0xba, 0x48, 0x4b, 0xba, 0x01,
	0x48, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x1a, 0x36, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x3f, 0x20,
	0x27, 0xe5, 0xbf, 0x85, 0xe9, 0xa1, 0xbb, 0xe6, 0x8f, 0x90, 0xe4, 0xbe, 0x9b, 0xe6, 0x9f, 0xa5,
	0xe8, 0xaf, 0xa2, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0x90, 0x8d, 0x2f, 0xe6, 0x98, 0xb5,
	0xe7, 0xa7, 0xb0, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa9, 0x05, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xc6, 0x02, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xa2, 0x02,
	0xba, 0x47, 0x0c, 0x92, 0x02, 0x09, 0xe6, 0x97, 0xa7, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0xba,
	0x48, 0x8f, 0x02, 0xba, 0x01, 0x8f, 0x01, 0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x6f, 0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x21, 0x40, 0x23, 0x24, 0x25, 0x5e, 0x26, 0x2a, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x20, 0x3f,
	0x20, 0x27, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0xe4, 0xbb, 0x85, 0xe5, 0x85, 0x81, 0xe8, 0xae,
	0xb8, 0xe5, 0x8c, 0x85, 0xe5, 0x90, 0xab, 0xe5, 0xad, 0x97, 0xe6, 0xaf, 0x8d, 0xe3, 0x80, 0x81,
	0xe6, 0x95, 0xb0, 0xe5, 0xad, 0x97, 0xe6, 0x88, 0x96, 0xe7, 0x89, 0xb9, 0xe6, 0xae, 0x8a, 0xe5,
	0xad, 0x97, 0xe7, 0xac, 0xa6, 0xef, 0xbc, 0x9a, 0x21, 0x40, 0x23, 0x24, 0x25, 0x5e, 0x26, 0x2a,
	0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0xba, 0x01, 0x79, 0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x59, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73,
	0x69, 0x7a, 0x65, 0x28, 0x29, 0x20, 0x3c, 0x20, 0x31, 0x30, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29, 0x20, 0x3e, 0x20, 0x32, 0x30, 0x29, 0x20,
	0x3f, 0x20, 0x27, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0xe9, 0x95, 0xbf, 0xe5, 0xba, 0xa6, 0xe5,
	0xbf, 0x85, 0xe9, 0xa1, 0xbb, 0xe5, 0x9c, 0xa8, 0x31, 0x30, 0x2d, 0x32, 0x30, 0xe4, 0xb8, 0xaa,
	0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4, 0xb9, 0x8b, 0xe9, 0x97, 0xb4, 0x27, 0x20, 0x3a, 0x20,
	0x27, 0x27, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0xc6, 0x02, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xa2, 0x02, 0xba, 0x47, 0x0c, 0x92, 0x02, 0x09, 0xe6,
	0x96, 0xb0, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0xba, 0x48, 0x8f, 0x02, 0xba, 0x01, 0x8f, 0x01,
	0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x6f,
	0x21, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x21, 0x40, 0x23, 0x24, 0x25, 0x5e,
	0x26, 0x2a, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x20, 0x3f, 0x20, 0x27, 0xe5, 0xaf, 0x86, 0xe7, 0xa0,
	0x81, 0xe4, 0xbb, 0x85, 0xe5, 0x85, 0x81, 0xe8, 0xae, 0xb8, 0xe5, 0x8c, 0x85, 0xe5, 0x90, 0xab,
	0xe5, 0xad, 0x97, 0xe6, 0xaf, 0x8d, 0xe3, 0x80, 0x81, 0xe6, 0x95, 0xb0, 0xe5, 0xad, 0x97, 0xe6,
	0x88, 0x96, 0xe7, 0x89, 0xb9, 0xe6, 0xae, 0x8a, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xef, 0xbc,
	0x9a, 0x21, 0x40, 0x23, 0x24, 0x25, 0x5e, 0x26, 0x2a, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0xba,
	0x01, 0x79, 0x0a, 0x1c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x1a, 0x59, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x29, 0x20, 0x3c,
	0x20, 0x31, 0x30, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x69, 0x7a, 0x65,
	0x28, 0x29, 0x20, 0x3e, 0x20, 0x32, 0x30, 0x29, 0x20, 0x3f, 0x20, 0x27, 0xe5, 0xaf, 0x86, 0xe7,
	0xa0, 0x81, 0xe9, 0x95, 0xbf, 0xe5, 0xba, 0xa6, 0xe5, 0xbf, 0x85, 0xe9, 0xa1, 0xbb, 0xe5, 0x9c,
	0xa8, 0x31, 0x30, 0x2d, 0x32, 0x30, 0xe4, 0xb8, 0xaa, 0xe5, 0xad, 0x97, 0xe7, 0xac, 0xa6, 0xe4,
	0xb9, 0x8b, 0xe9, 0x97, 0xb4, 0x27, 0x20, 0x3a, 0x20, 0x27, 0x27, 0x52, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xc3, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x66, 0x12, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x72, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x1a, 0x11, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x7b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x7a, 0x77, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatsvc_v1_user_api_proto_rawDescOnce sync.Once
	file_chatsvc_v1_user_api_proto_rawDescData = file_chatsvc_v1_user_api_proto_rawDesc
)

func file_chatsvc_v1_user_api_proto_rawDescGZIP() []byte {
	file_chatsvc_v1_user_api_proto_rawDescOnce.Do(func() {
		file_chatsvc_v1_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatsvc_v1_user_api_proto_rawDescData)
	})
	return file_chatsvc_v1_user_api_proto_rawDescData
}

var file_chatsvc_v1_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chatsvc_v1_user_api_proto_goTypes = []interface{}{
	(*UserVO)(nil),                // 0: chatsvc.v1.UserVO
	(*GetUserSelfRequest)(nil),    // 1: chatsvc.v1.GetUserSelfRequest
	(*GetUserSelfReply)(nil),      // 2: chatsvc.v1.GetUserSelfReply
	(*UpdateUserInfoRequest)(nil), // 3: chatsvc.v1.UpdateUserInfoRequest
	(*UpdateUserInfoReply)(nil),   // 4: chatsvc.v1.UpdateUserInfoReply
	(*GetUserByIDRequest)(nil),    // 5: chatsvc.v1.GetUserByIDRequest
	(*GetUserByIDReply)(nil),      // 6: chatsvc.v1.GetUserByIDReply
	(*ListUserInfoRequest)(nil),   // 7: chatsvc.v1.ListUserInfoRequest
	(*ListUserInfoReply)(nil),     // 8: chatsvc.v1.ListUserInfoReply
	(*UpdatePasswordRequest)(nil), // 9: chatsvc.v1.UpdatePasswordRequest
	(*UpdatePasswordReply)(nil),   // 10: chatsvc.v1.UpdatePasswordReply
}
var file_chatsvc_v1_user_api_proto_depIdxs = []int32{
	0,  // 0: chatsvc.v1.GetUserSelfReply.data:type_name -> chatsvc.v1.UserVO
	0,  // 1: chatsvc.v1.GetUserByIDReply.data:type_name -> chatsvc.v1.UserVO
	0,  // 2: chatsvc.v1.ListUserInfoReply.data:type_name -> chatsvc.v1.UserVO
	1,  // 3: chatsvc.v1.UserService.GetUserSelf:input_type -> chatsvc.v1.GetUserSelfRequest
	3,  // 4: chatsvc.v1.UserService.UpdateUserInfo:input_type -> chatsvc.v1.UpdateUserInfoRequest
	5,  // 5: chatsvc.v1.UserService.GetUserByID:input_type -> chatsvc.v1.GetUserByIDRequest
	7,  // 6: chatsvc.v1.UserService.ListUserInfo:input_type -> chatsvc.v1.ListUserInfoRequest
	9,  // 7: chatsvc.v1.UserService.UpdatePassword:input_type -> chatsvc.v1.UpdatePasswordRequest
	2,  // 8: chatsvc.v1.UserService.GetUserSelf:output_type -> chatsvc.v1.GetUserSelfReply
	4,  // 9: chatsvc.v1.UserService.UpdateUserInfo:output_type -> chatsvc.v1.UpdateUserInfoReply
	6,  // 10: chatsvc.v1.UserService.GetUserByID:output_type -> chatsvc.v1.GetUserByIDReply
	8,  // 11: chatsvc.v1.UserService.ListUserInfo:output_type -> chatsvc.v1.ListUserInfoReply
	10, // 12: chatsvc.v1.UserService.UpdatePassword:output_type -> chatsvc.v1.UpdatePasswordReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_chatsvc_v1_user_api_proto_init() }
func file_chatsvc_v1_user_api_proto_init() {
	if File_chatsvc_v1_user_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatsvc_v1_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVO); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSelfRequest); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSelfReply); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoRequest); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoReply); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDRequest); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDReply); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserInfoRequest); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserInfoReply); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordRequest); i {
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
		file_chatsvc_v1_user_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasswordReply); i {
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
			RawDescriptor: file_chatsvc_v1_user_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatsvc_v1_user_api_proto_goTypes,
		DependencyIndexes: file_chatsvc_v1_user_api_proto_depIdxs,
		MessageInfos:      file_chatsvc_v1_user_api_proto_msgTypes,
	}.Build()
	File_chatsvc_v1_user_api_proto = out.File
	file_chatsvc_v1_user_api_proto_rawDesc = nil
	file_chatsvc_v1_user_api_proto_goTypes = nil
	file_chatsvc_v1_user_api_proto_depIdxs = nil
}
