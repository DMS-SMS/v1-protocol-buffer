// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: auth-parent.proto

package authproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoginParentAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentID string `protobuf:"bytes,1,opt,name=ParentID,proto3" json:"ParentID,omitempty"` // 학부모용 계정 아이디 (4~16자 사이)
	ParentPW string `protobuf:"bytes,2,opt,name=ParentPW,proto3" json:"ParentPW,omitempty"` // 학부모용 계정 비밀번호 (4~16자 사이)
}

func (x *LoginParentAuthRequest) Reset() {
	*x = LoginParentAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParentAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParentAuthRequest) ProtoMessage() {}

func (x *LoginParentAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParentAuthRequest.ProtoReflect.Descriptor instead.
func (*LoginParentAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{0}
}

func (x *LoginParentAuthRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *LoginParentAuthRequest) GetParentPW() string {
	if x != nil {
		return x.ParentPW
	}
	return ""
}

type LoginParentAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`                        // 상태 코드
	Code               int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`                          // 세부 코드
	Message            string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`                       // 세부 설명
	AccessToken        string `protobuf:"bytes,4,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`               // 학부모 계정 인증을 위한 JWT
	LoggedInParentUUID string `protobuf:"bytes,5,opt,name=LoggedInParentUUID,proto3" json:"LoggedInParentUUID,omitempty"` // 로그인한 학부모 계정의 uuid
}

func (x *LoginParentAuthResponse) Reset() {
	*x = LoginParentAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParentAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParentAuthResponse) ProtoMessage() {}

func (x *LoginParentAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParentAuthResponse.ProtoReflect.Descriptor instead.
func (*LoginParentAuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{1}
}

func (x *LoginParentAuthResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginParentAuthResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginParentAuthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginParentAuthResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginParentAuthResponse) GetLoggedInParentUUID() string {
	if x != nil {
		return x.LoggedInParentUUID
	}
	return ""
}

type ChangeParentPWRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID       string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`             // API를 호출한 학부모 계정의 uuid (20자)
	ParentUUID string `protobuf:"bytes,2,opt,name=ParentUUID,proto3" json:"ParentUUID,omitempty"` // 변경할 선생님 계정의 uuid (20자)
	CurrentPW  string `protobuf:"bytes,3,opt,name=CurrentPW,proto3" json:"CurrentPW,omitempty"`   // 해당 사용자의 현재 PW (4~16 사이)
	RevisionPW string `protobuf:"bytes,4,opt,name=RevisionPW,proto3" json:"RevisionPW,omitempty"` // 변경할 PW (4~16 사이)
}

func (x *ChangeParentPWRequest) Reset() {
	*x = ChangeParentPWRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeParentPWRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeParentPWRequest) ProtoMessage() {}

func (x *ChangeParentPWRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeParentPWRequest.ProtoReflect.Descriptor instead.
func (*ChangeParentPWRequest) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeParentPWRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *ChangeParentPWRequest) GetParentUUID() string {
	if x != nil {
		return x.ParentUUID
	}
	return ""
}

func (x *ChangeParentPWRequest) GetCurrentPW() string {
	if x != nil {
		return x.CurrentPW
	}
	return ""
}

func (x *ChangeParentPWRequest) GetRevisionPW() string {
	if x != nil {
		return x.RevisionPW
	}
	return ""
}

type ChangeParentPWResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`  // 상태 코드
	Code    int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`    // 세부 코드
	Message string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"` // 세부 설명
}

func (x *ChangeParentPWResponse) Reset() {
	*x = ChangeParentPWResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeParentPWResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeParentPWResponse) ProtoMessage() {}

func (x *ChangeParentPWResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeParentPWResponse.ProtoReflect.Descriptor instead.
func (*ChangeParentPWResponse) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeParentPWResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ChangeParentPWResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangeParentPWResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetParentInformWithUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID       string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`             // API를 호출한 학부모 계정의 uuid (19자)
	ParentUUID string `protobuf:"bytes,2,opt,name=ParentUUID,proto3" json:"ParentUUID,omitempty"` // 조회할 학부모 계정의 uuid (19자)
}

func (x *GetParentInformWithUUIDRequest) Reset() {
	*x = GetParentInformWithUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParentInformWithUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParentInformWithUUIDRequest) ProtoMessage() {}

func (x *GetParentInformWithUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParentInformWithUUIDRequest.ProtoReflect.Descriptor instead.
func (*GetParentInformWithUUIDRequest) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{4}
}

func (x *GetParentInformWithUUIDRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *GetParentInformWithUUIDRequest) GetParentUUID() string {
	if x != nil {
		return x.ParentUUID
	}
	return ""
}

type GetParentInformWithUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`          // 상태 코드
	Code        int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`            // 세부 코드
	Message     string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`         // 세부 설명
	Name        string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`               // 학부모 이름
	PhoneNumber string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"` // 전화번호 정보
}

func (x *GetParentInformWithUUIDResponse) Reset() {
	*x = GetParentInformWithUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParentInformWithUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParentInformWithUUIDResponse) ProtoMessage() {}

func (x *GetParentInformWithUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParentInformWithUUIDResponse.ProtoReflect.Descriptor instead.
func (*GetParentInformWithUUIDResponse) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{5}
}

func (x *GetParentInformWithUUIDResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetParentInformWithUUIDResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetParentInformWithUUIDResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetParentInformWithUUIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetParentInformWithUUIDResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetParentUUIDsWithInformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID        string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`               // API를 호출한 학부모 계정의 uuid (19자)
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`               // 학부모 이름
	PhoneNumber string `protobuf:"bytes,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"` // 전화번호 정보
}

func (x *GetParentUUIDsWithInformRequest) Reset() {
	*x = GetParentUUIDsWithInformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParentUUIDsWithInformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParentUUIDsWithInformRequest) ProtoMessage() {}

func (x *GetParentUUIDsWithInformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParentUUIDsWithInformRequest.ProtoReflect.Descriptor instead.
func (*GetParentUUIDsWithInformRequest) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{6}
}

func (x *GetParentUUIDsWithInformRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *GetParentUUIDsWithInformRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetParentUUIDsWithInformRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetParentUUIDsWithInformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      uint32   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Code        int32    `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Message     string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	ParentUUIDs []string `protobuf:"bytes,4,rep,name=ParentUUIDs,proto3" json:"ParentUUIDs,omitempty"`
}

func (x *GetParentUUIDsWithInformResponse) Reset() {
	*x = GetParentUUIDsWithInformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetParentUUIDsWithInformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParentUUIDsWithInformResponse) ProtoMessage() {}

func (x *GetParentUUIDsWithInformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParentUUIDsWithInformResponse.ProtoReflect.Descriptor instead.
func (*GetParentUUIDsWithInformResponse) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{7}
}

func (x *GetParentUUIDsWithInformResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetParentUUIDsWithInformResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetParentUUIDsWithInformResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetParentUUIDsWithInformResponse) GetParentUUIDs() []string {
	if x != nil {
		return x.ParentUUIDs
	}
	return nil
}

type GetChildrenInformsWithUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID       string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`             // API를 호출한 학부모 계정의 uuid (19자)
	ParentUUID string `protobuf:"bytes,2,opt,name=ParentUUID,proto3" json:"ParentUUID,omitempty"` // 자녀 정보를 조회할 학부모 계정의 uuid (19자)
}

func (x *GetChildrenInformsWithUUIDRequest) Reset() {
	*x = GetChildrenInformsWithUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChildrenInformsWithUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChildrenInformsWithUUIDRequest) ProtoMessage() {}

func (x *GetChildrenInformsWithUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChildrenInformsWithUUIDRequest.ProtoReflect.Descriptor instead.
func (*GetChildrenInformsWithUUIDRequest) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{8}
}

func (x *GetChildrenInformsWithUUIDRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *GetChildrenInformsWithUUIDRequest) GetParentUUID() string {
	if x != nil {
		return x.ParentUUID
	}
	return ""
}

type GetChildrenInformsWithUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         uint32           `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Code           int32            `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Message        string           `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	ChildrenInform []*StudentInform `protobuf:"bytes,4,rep,name=ChildrenInform,proto3" json:"ChildrenInform,omitempty"` // 자녀 정보 지르
}

func (x *GetChildrenInformsWithUUIDResponse) Reset() {
	*x = GetChildrenInformsWithUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_parent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChildrenInformsWithUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChildrenInformsWithUUIDResponse) ProtoMessage() {}

func (x *GetChildrenInformsWithUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_parent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChildrenInformsWithUUIDResponse.ProtoReflect.Descriptor instead.
func (*GetChildrenInformsWithUUIDResponse) Descriptor() ([]byte, []int) {
	return file_auth_parent_proto_rawDescGZIP(), []int{9}
}

func (x *GetChildrenInformsWithUUIDResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetChildrenInformsWithUUIDResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetChildrenInformsWithUUIDResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetChildrenInformsWithUUIDResponse) GetChildrenInform() []*StudentInform {
	if x != nil {
		return x.ChildrenInform
	}
	return nil
}

var File_auth_parent_proto protoreflect.FileDescriptor

var file_auth_parent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x1a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x22, 0xb1, 0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x57, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x57, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x57, 0x22, 0x5e, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x9d, 0x01, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x6b, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8a, 0x01, 0x0a, 0x20, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55,
	0x49, 0x44, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x22, 0x57, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22,
	0xaf, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x52, 0x0e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x32, 0xbd, 0x04, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x5e, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x57, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x57, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x7f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_parent_proto_rawDescOnce sync.Once
	file_auth_parent_proto_rawDescData = file_auth_parent_proto_rawDesc
)

func file_auth_parent_proto_rawDescGZIP() []byte {
	file_auth_parent_proto_rawDescOnce.Do(func() {
		file_auth_parent_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_parent_proto_rawDescData)
	})
	return file_auth_parent_proto_rawDescData
}

var file_auth_parent_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_auth_parent_proto_goTypes = []interface{}{
	(*LoginParentAuthRequest)(nil),             // 0: auth.parent.LoginParentAuthRequest
	(*LoginParentAuthResponse)(nil),            // 1: auth.parent.LoginParentAuthResponse
	(*ChangeParentPWRequest)(nil),              // 2: auth.parent.ChangeParentPWRequest
	(*ChangeParentPWResponse)(nil),             // 3: auth.parent.ChangeParentPWResponse
	(*GetParentInformWithUUIDRequest)(nil),     // 4: auth.parent.GetParentInformWithUUIDRequest
	(*GetParentInformWithUUIDResponse)(nil),    // 5: auth.parent.GetParentInformWithUUIDResponse
	(*GetParentUUIDsWithInformRequest)(nil),    // 6: auth.parent.GetParentUUIDsWithInformRequest
	(*GetParentUUIDsWithInformResponse)(nil),   // 7: auth.parent.GetParentUUIDsWithInformResponse
	(*GetChildrenInformsWithUUIDRequest)(nil),  // 8: auth.parent.GetChildrenInformsWithUUIDRequest
	(*GetChildrenInformsWithUUIDResponse)(nil), // 9: auth.parent.GetChildrenInformsWithUUIDResponse
	(*StudentInform)(nil),                      // 10: auth.student.StudentInform
}
var file_auth_parent_proto_depIdxs = []int32{
	10, // 0: auth.parent.GetChildrenInformsWithUUIDResponse.ChildrenInform:type_name -> auth.student.StudentInform
	0,  // 1: auth.parent.AuthParent.LoginParentAuth:input_type -> auth.parent.LoginParentAuthRequest
	2,  // 2: auth.parent.AuthParent.ChangeParentPW:input_type -> auth.parent.ChangeParentPWRequest
	4,  // 3: auth.parent.AuthParent.GetParentInformWithUUID:input_type -> auth.parent.GetParentInformWithUUIDRequest
	6,  // 4: auth.parent.AuthParent.GetParentUUIDsWithInform:input_type -> auth.parent.GetParentUUIDsWithInformRequest
	8,  // 5: auth.parent.AuthParent.GetChildrenInformsWithUUID:input_type -> auth.parent.GetChildrenInformsWithUUIDRequest
	1,  // 6: auth.parent.AuthParent.LoginParentAuth:output_type -> auth.parent.LoginParentAuthResponse
	3,  // 7: auth.parent.AuthParent.ChangeParentPW:output_type -> auth.parent.ChangeParentPWResponse
	5,  // 8: auth.parent.AuthParent.GetParentInformWithUUID:output_type -> auth.parent.GetParentInformWithUUIDResponse
	7,  // 9: auth.parent.AuthParent.GetParentUUIDsWithInform:output_type -> auth.parent.GetParentUUIDsWithInformResponse
	9,  // 10: auth.parent.AuthParent.GetChildrenInformsWithUUID:output_type -> auth.parent.GetChildrenInformsWithUUIDResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_auth_parent_proto_init() }
func file_auth_parent_proto_init() {
	if File_auth_parent_proto != nil {
		return
	}
	file_auth_student_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_parent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParentAuthRequest); i {
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
		file_auth_parent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParentAuthResponse); i {
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
		file_auth_parent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeParentPWRequest); i {
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
		file_auth_parent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeParentPWResponse); i {
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
		file_auth_parent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParentInformWithUUIDRequest); i {
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
		file_auth_parent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParentInformWithUUIDResponse); i {
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
		file_auth_parent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParentUUIDsWithInformRequest); i {
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
		file_auth_parent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetParentUUIDsWithInformResponse); i {
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
		file_auth_parent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChildrenInformsWithUUIDRequest); i {
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
		file_auth_parent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChildrenInformsWithUUIDResponse); i {
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
			RawDescriptor: file_auth_parent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_parent_proto_goTypes,
		DependencyIndexes: file_auth_parent_proto_depIdxs,
		MessageInfos:      file_auth_parent_proto_msgTypes,
	}.Build()
	File_auth_parent_proto = out.File
	file_auth_parent_proto_rawDesc = nil
	file_auth_parent_proto_goTypes = nil
	file_auth_parent_proto_depIdxs = nil
}
