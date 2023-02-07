// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: MiniTikTok-Proto/relation.proto

package relation

import (
	context "context"
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

// 0-关注，1-取消关注
type FollowActionRequest_FollowActionType int32

const (
	FollowActionRequest_FOLLOW   FollowActionRequest_FollowActionType = 0
	FollowActionRequest_UNFOLLOW FollowActionRequest_FollowActionType = 1
)

// Enum value maps for FollowActionRequest_FollowActionType.
var (
	FollowActionRequest_FollowActionType_name = map[int32]string{
		0: "FOLLOW",
		1: "UNFOLLOW",
	}
	FollowActionRequest_FollowActionType_value = map[string]int32{
		"FOLLOW":   0,
		"UNFOLLOW": 1,
	}
)

func (x FollowActionRequest_FollowActionType) Enum() *FollowActionRequest_FollowActionType {
	p := new(FollowActionRequest_FollowActionType)
	*p = x
	return p
}

func (x FollowActionRequest_FollowActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FollowActionRequest_FollowActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_MiniTikTok_Proto_relation_proto_enumTypes[0].Descriptor()
}

func (FollowActionRequest_FollowActionType) Type() protoreflect.EnumType {
	return &file_MiniTikTok_Proto_relation_proto_enumTypes[0]
}

func (x FollowActionRequest_FollowActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FollowActionRequest_FollowActionType.Descriptor instead.
func (FollowActionRequest_FollowActionType) EnumDescriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{1, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 关注总数
	FollowCount uint32 `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	// 粉丝总数
	FollowerCount uint32 `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	// true-已关注，false-未关注
	IsFollow bool `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() uint32 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() uint32 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type FollowActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 对方用户 id
	ToUserId uint32 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	// 关系操作码
	ActionType FollowActionRequest_FollowActionType `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=mini_tiktok.proto.relation.FollowActionRequest_FollowActionType" json:"action_type,omitempty"`
}

func (x *FollowActionRequest) Reset() {
	*x = FollowActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowActionRequest) ProtoMessage() {}

func (x *FollowActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowActionRequest.ProtoReflect.Descriptor instead.
func (*FollowActionRequest) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{1}
}

func (x *FollowActionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FollowActionRequest) GetToUserId() uint32 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *FollowActionRequest) GetActionType() FollowActionRequest_FollowActionType {
	if x != nil {
		return x.ActionType
	}
	return FollowActionRequest_FOLLOW
}

type FollowActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-成功，非0-失败
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// 状态描述
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *FollowActionResponse) Reset() {
	*x = FollowActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowActionResponse) ProtoMessage() {}

func (x *FollowActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowActionResponse.ProtoReflect.Descriptor instead.
func (*FollowActionResponse) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{2}
}

func (x *FollowActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FollowActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type ListFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 执行操作的用户 id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 目标用户 id
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListFollowRequest) Reset() {
	*x = ListFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowRequest) ProtoMessage() {}

func (x *ListFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowRequest.ProtoReflect.Descriptor instead.
func (*ListFollowRequest) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{3}
}

func (x *ListFollowRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListFollowRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-成功，非0-失败
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// 状态描述
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// 用户信息列表
	UserList []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *ListFollowResponse) Reset() {
	*x = ListFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowResponse) ProtoMessage() {}

func (x *ListFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowResponse.ProtoReflect.Descriptor instead.
func (*ListFollowResponse) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{4}
}

func (x *ListFollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListFollowResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListFollowResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type ListFansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 执行操作的用户 id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 目标用户 id
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListFansRequest) Reset() {
	*x = ListFansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFansRequest) ProtoMessage() {}

func (x *ListFansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFansRequest.ProtoReflect.Descriptor instead.
func (*ListFansRequest) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{5}
}

func (x *ListFansRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListFansRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListFansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-成功，非0-失败
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// 状态描述
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// 用户信息列表
	UserList []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *ListFansResponse) Reset() {
	*x = ListFansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFansResponse) ProtoMessage() {}

func (x *ListFansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFansResponse.ProtoReflect.Descriptor instead.
func (*ListFansResponse) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{6}
}

func (x *ListFansResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListFansResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListFansResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type ListFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 执行操作的用户 id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 目标用户 id
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListFriendsRequest) Reset() {
	*x = ListFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendsRequest) ProtoMessage() {}

func (x *ListFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendsRequest.ProtoReflect.Descriptor instead.
func (*ListFriendsRequest) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{7}
}

func (x *ListFriendsRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListFriendsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListFriendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 0-成功，非0-失败
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// 状态描述
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// 用户信息列表
	UserList []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *ListFriendsResponse) Reset() {
	*x = ListFriendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendsResponse) ProtoMessage() {}

func (x *ListFriendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MiniTikTok_Proto_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendsResponse.ProtoReflect.Descriptor instead.
func (*ListFriendsResponse) Descriptor() ([]byte, []int) {
	return file_MiniTikTok_Proto_relation_proto_rawDescGZIP(), []int{8}
}

func (x *ListFriendsResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListFriendsResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListFriendsResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

var File_MiniTikTok_Proto_relation_proto protoreflect.FileDescriptor

var file_MiniTikTok_Proto_relation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x69, 0x6e, 0x69, 0x54, 0x69, 0x6b, 0x54, 0x6f, 0x6b, 0x2d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x61, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e,
	0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x22, 0x56, 0x0a, 0x14, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x22, 0x3c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x93,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f,
	0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x91, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f,
	0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xd0, 0x03, 0x0a, 0x0f, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73,
	0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x67, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x12, 0x2b,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x61, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x2e, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x64, 0x6c,
	0x73, 0x73, 0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x54, 0x69, 0x6b, 0x54, 0x6f, 0x6b, 0x2d, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MiniTikTok_Proto_relation_proto_rawDescOnce sync.Once
	file_MiniTikTok_Proto_relation_proto_rawDescData = file_MiniTikTok_Proto_relation_proto_rawDesc
)

func file_MiniTikTok_Proto_relation_proto_rawDescGZIP() []byte {
	file_MiniTikTok_Proto_relation_proto_rawDescOnce.Do(func() {
		file_MiniTikTok_Proto_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_MiniTikTok_Proto_relation_proto_rawDescData)
	})
	return file_MiniTikTok_Proto_relation_proto_rawDescData
}

var file_MiniTikTok_Proto_relation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MiniTikTok_Proto_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_MiniTikTok_Proto_relation_proto_goTypes = []interface{}{
	(FollowActionRequest_FollowActionType)(0), // 0: mini_tiktok.proto.relation.FollowActionRequest.FollowActionType
	(*User)(nil),                 // 1: mini_tiktok.proto.relation.User
	(*FollowActionRequest)(nil),  // 2: mini_tiktok.proto.relation.FollowActionRequest
	(*FollowActionResponse)(nil), // 3: mini_tiktok.proto.relation.FollowActionResponse
	(*ListFollowRequest)(nil),    // 4: mini_tiktok.proto.relation.ListFollowRequest
	(*ListFollowResponse)(nil),   // 5: mini_tiktok.proto.relation.ListFollowResponse
	(*ListFansRequest)(nil),      // 6: mini_tiktok.proto.relation.ListFansRequest
	(*ListFansResponse)(nil),     // 7: mini_tiktok.proto.relation.ListFansResponse
	(*ListFriendsRequest)(nil),   // 8: mini_tiktok.proto.relation.ListFriendsRequest
	(*ListFriendsResponse)(nil),  // 9: mini_tiktok.proto.relation.ListFriendsResponse
}
var file_MiniTikTok_Proto_relation_proto_depIdxs = []int32{
	0, // 0: mini_tiktok.proto.relation.FollowActionRequest.action_type:type_name -> mini_tiktok.proto.relation.FollowActionRequest.FollowActionType
	1, // 1: mini_tiktok.proto.relation.ListFollowResponse.user_list:type_name -> mini_tiktok.proto.relation.User
	1, // 2: mini_tiktok.proto.relation.ListFansResponse.user_list:type_name -> mini_tiktok.proto.relation.User
	1, // 3: mini_tiktok.proto.relation.ListFriendsResponse.user_list:type_name -> mini_tiktok.proto.relation.User
	2, // 4: mini_tiktok.proto.relation.RelationService.FollowAction:input_type -> mini_tiktok.proto.relation.FollowActionRequest
	4, // 5: mini_tiktok.proto.relation.RelationService.ListFollow:input_type -> mini_tiktok.proto.relation.ListFollowRequest
	6, // 6: mini_tiktok.proto.relation.RelationService.ListFans:input_type -> mini_tiktok.proto.relation.ListFansRequest
	8, // 7: mini_tiktok.proto.relation.RelationService.ListFriends:input_type -> mini_tiktok.proto.relation.ListFriendsRequest
	3, // 8: mini_tiktok.proto.relation.RelationService.FollowAction:output_type -> mini_tiktok.proto.relation.FollowActionResponse
	5, // 9: mini_tiktok.proto.relation.RelationService.ListFollow:output_type -> mini_tiktok.proto.relation.ListFollowResponse
	7, // 10: mini_tiktok.proto.relation.RelationService.ListFans:output_type -> mini_tiktok.proto.relation.ListFansResponse
	9, // 11: mini_tiktok.proto.relation.RelationService.ListFriends:output_type -> mini_tiktok.proto.relation.ListFriendsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_MiniTikTok_Proto_relation_proto_init() }
func file_MiniTikTok_Proto_relation_proto_init() {
	if File_MiniTikTok_Proto_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MiniTikTok_Proto_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowActionRequest); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowActionResponse); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowRequest); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowResponse); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFansRequest); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFansResponse); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendsRequest); i {
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
		file_MiniTikTok_Proto_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendsResponse); i {
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
			RawDescriptor: file_MiniTikTok_Proto_relation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MiniTikTok_Proto_relation_proto_goTypes,
		DependencyIndexes: file_MiniTikTok_Proto_relation_proto_depIdxs,
		EnumInfos:         file_MiniTikTok_Proto_relation_proto_enumTypes,
		MessageInfos:      file_MiniTikTok_Proto_relation_proto_msgTypes,
	}.Build()
	File_MiniTikTok_Proto_relation_proto = out.File
	file_MiniTikTok_Proto_relation_proto_rawDesc = nil
	file_MiniTikTok_Proto_relation_proto_goTypes = nil
	file_MiniTikTok_Proto_relation_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type RelationService interface {
	FollowAction(ctx context.Context, req *FollowActionRequest) (res *FollowActionResponse, err error)
	ListFollow(ctx context.Context, req *ListFollowRequest) (res *ListFollowResponse, err error)
	ListFans(ctx context.Context, req *ListFansRequest) (res *ListFansResponse, err error)
	ListFriends(ctx context.Context, req *ListFriendsRequest) (res *ListFriendsResponse, err error)
}