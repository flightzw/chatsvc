// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: chatsvc/model/friend.proto

package model

import (
	_ "github.com/google/gnostic/openapiv3"
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

type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FriendId  int32  `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatsvc_model_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_chatsvc_model_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_chatsvc_model_friend_proto_rawDescGZIP(), []int{0}
}

func (x *Friend) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Friend) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *Friend) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Friend) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Friend) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_chatsvc_model_friend_proto protoreflect.FileDescriptor

var file_chatsvc_model_friend_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x76, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x06, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xba, 0x47, 0x0b, 0x4a, 0x09, 0xe5, 0xa5,
	0xbd, 0xe5, 0x8f, 0x8b, 0x75, 0x69, 0x64, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x47, 0x08, 0x4a, 0x06, 0xe6, 0x98, 0xb5, 0xe7, 0xa7, 0xb0,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0xba, 0x47, 0x0b, 0x4a, 0x09, 0xe5, 0xa4, 0xb4, 0xe5, 0x83, 0x8f, 0x75, 0x72, 0x6c, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xba,
	0x47, 0x0e, 0x4a, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x39, 0xba, 0x47, 0x36,
	0xba, 0x01, 0x02, 0x69, 0x64, 0xba, 0x01, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0xba, 0x01, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0xba, 0x01, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0xba, 0x01, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x7a, 0x77, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x76,
	0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatsvc_model_friend_proto_rawDescOnce sync.Once
	file_chatsvc_model_friend_proto_rawDescData = file_chatsvc_model_friend_proto_rawDesc
)

func file_chatsvc_model_friend_proto_rawDescGZIP() []byte {
	file_chatsvc_model_friend_proto_rawDescOnce.Do(func() {
		file_chatsvc_model_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatsvc_model_friend_proto_rawDescData)
	})
	return file_chatsvc_model_friend_proto_rawDescData
}

var file_chatsvc_model_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chatsvc_model_friend_proto_goTypes = []interface{}{
	(*Friend)(nil), // 0: chatsvc.model.Friend
}
var file_chatsvc_model_friend_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatsvc_model_friend_proto_init() }
func file_chatsvc_model_friend_proto_init() {
	if File_chatsvc_model_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatsvc_model_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Friend); i {
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
			RawDescriptor: file_chatsvc_model_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chatsvc_model_friend_proto_goTypes,
		DependencyIndexes: file_chatsvc_model_friend_proto_depIdxs,
		MessageInfos:      file_chatsvc_model_friend_proto_msgTypes,
	}.Build()
	File_chatsvc_model_friend_proto = out.File
	file_chatsvc_model_friend_proto_rawDesc = nil
	file_chatsvc_model_friend_proto_goTypes = nil
	file_chatsvc_model_friend_proto_depIdxs = nil
}
