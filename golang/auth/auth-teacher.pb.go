// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: auth-teacher.proto

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

type LoginTeacherAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 선생님용 계정 아이디 (4~16자 사이)
	Pw string `protobuf:"bytes,2,opt,name=pw,proto3" json:"pw,omitempty"` // 선생님용 계정 비밀번호 (4~16자 사이)
}

func (x *LoginTeacherAuthRequest) Reset() {
	*x = LoginTeacherAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTeacherAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTeacherAuthRequest) ProtoMessage() {}

func (x *LoginTeacherAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTeacherAuthRequest.ProtoReflect.Descriptor instead.
func (*LoginTeacherAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *LoginTeacherAuthRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginTeacherAuthRequest) GetPw() string {
	if x != nil {
		return x.Pw
	}
	return ""
}

type LoginTeacherAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`                             // 상태 코드
	Code        int32  `protobuf:"zigzag32,2,opt,name=code,proto3" json:"code,omitempty"`                               // 세부 코드
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`                            // 세부 설명
	AccessToken string `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"` // 선생님 계정 인증을 위한 JWT
	Tid         string `protobuf:"bytes,5,opt,name=tid,proto3" json:"tid,omitempty"`                                    // 로그인한 선생님 계정의 uuid
}

func (x *LoginTeacherAuthResponse) Reset() {
	*x = LoginTeacherAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginTeacherAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginTeacherAuthResponse) ProtoMessage() {}

func (x *LoginTeacherAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginTeacherAuthResponse.ProtoReflect.Descriptor instead.
func (*LoginTeacherAuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *LoginTeacherAuthResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginTeacherAuthResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginTeacherAuthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginTeacherAuthResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginTeacherAuthResponse) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

type ChangeTeacherAuthPwRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                            // API를 호출한 선생님 계정의 uuid (20자)
	Tid       string `protobuf:"bytes,2,opt,name=tid,proto3" json:"tid,omitempty"`                              // 변경할 선생님 계정의 uuid (20자)
	CurrentPw string `protobuf:"bytes,3,opt,name=current_pw,json=currentPw,proto3" json:"current_pw,omitempty"` // 해당 사용자의 현재 Pw (4~16 사이)
	NewPw     string `protobuf:"bytes,4,opt,name=new_pw,json=newPw,proto3" json:"new_pw,omitempty"`             // 변경할 Pw (4~16 사이)
}

func (x *ChangeTeacherAuthPwRequest) Reset() {
	*x = ChangeTeacherAuthPwRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeTeacherAuthPwRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeTeacherAuthPwRequest) ProtoMessage() {}

func (x *ChangeTeacherAuthPwRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeTeacherAuthPwRequest.ProtoReflect.Descriptor instead.
func (*ChangeTeacherAuthPwRequest) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeTeacherAuthPwRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ChangeTeacherAuthPwRequest) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *ChangeTeacherAuthPwRequest) GetCurrentPw() string {
	if x != nil {
		return x.CurrentPw
	}
	return ""
}

func (x *ChangeTeacherAuthPwRequest) GetNewPw() string {
	if x != nil {
		return x.NewPw
	}
	return ""
}

type ChangeTeacherAuthPwResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`  // 상태 코드
	Code    int32  `protobuf:"zigzag32,2,opt,name=code,proto3" json:"code,omitempty"`    // 세부 코드
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // 세부 설명
}

func (x *ChangeTeacherAuthPwResponse) Reset() {
	*x = ChangeTeacherAuthPwResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeTeacherAuthPwResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeTeacherAuthPwResponse) ProtoMessage() {}

func (x *ChangeTeacherAuthPwResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeTeacherAuthPwResponse.ProtoReflect.Descriptor instead.
func (*ChangeTeacherAuthPwResponse) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeTeacherAuthPwResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ChangeTeacherAuthPwResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangeTeacherAuthPwResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTeacherUserInformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // API를 호출한 학생 계정의 uuid (20자)
	Tid  string `protobuf:"bytes,2,opt,name=tid,proto3" json:"tid,omitempty"`   // 조회할 선생님 계정의 uuid (20자)
}

func (x *GetTeacherUserInformRequest) Reset() {
	*x = GetTeacherUserInformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeacherUserInformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherUserInformRequest) ProtoMessage() {}

func (x *GetTeacherUserInformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherUserInformRequest.ProtoReflect.Descriptor instead.
func (*GetTeacherUserInformRequest) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *GetTeacherUserInformRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetTeacherUserInformRequest) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

type GetTeacherUserInformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`  // 상태 코드
	Code    int32  `protobuf:"zigzag32,2,opt,name=code,proto3" json:"code,omitempty"`    // 세부 코드
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"` // 세부 설명
	Grade   uint32 `protobuf:"varint,4,opt,name=grade,proto3" json:"grade,omitempty"`    // 학년 정보
	Class   uint32 `protobuf:"varint,5,opt,name=class,proto3" json:"class,omitempty"`    // 반 정보
	Name    string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`       // 이름 정보
}

func (x *GetTeacherUserInformResponse) Reset() {
	*x = GetTeacherUserInformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_teacher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeacherUserInformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherUserInformResponse) ProtoMessage() {}

func (x *GetTeacherUserInformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_teacher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherUserInformResponse.ProtoReflect.Descriptor instead.
func (*GetTeacherUserInformResponse) Descriptor() ([]byte, []int) {
	return file_auth_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *GetTeacherUserInformResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetTeacherUserInformResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTeacherUserInformResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTeacherUserInformResponse) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *GetTeacherUserInformResponse) GetClass() uint32 {
	if x != nil {
		return x.Class
	}
	return 0
}

func (x *GetTeacherUserInformResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_auth_teacher_proto protoreflect.FileDescriptor

var file_auth_teacher_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x70, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x77, 0x22,
	0x95, 0x01, 0x0a, 0x18, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x50, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x77, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x65,
	0x77, 0x5f, 0x70, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x77, 0x50,
	0x77, 0x22, 0x63, 0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x50, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x1c,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0xfc, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x49, 0x0a, 0x10, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x18, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x50, 0x77, 0x12, 0x1b, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x50, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_teacher_proto_rawDescOnce sync.Once
	file_auth_teacher_proto_rawDescData = file_auth_teacher_proto_rawDesc
)

func file_auth_teacher_proto_rawDescGZIP() []byte {
	file_auth_teacher_proto_rawDescOnce.Do(func() {
		file_auth_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_teacher_proto_rawDescData)
	})
	return file_auth_teacher_proto_rawDescData
}

var file_auth_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_teacher_proto_goTypes = []interface{}{
	(*LoginTeacherAuthRequest)(nil),      // 0: LoginTeacherAuthRequest
	(*LoginTeacherAuthResponse)(nil),     // 1: LoginTeacherAuthResponse
	(*ChangeTeacherAuthPwRequest)(nil),   // 2: ChangeTeacherAuthPwRequest
	(*ChangeTeacherAuthPwResponse)(nil),  // 3: ChangeTeacherAuthPwResponse
	(*GetTeacherUserInformRequest)(nil),  // 4: GetTeacherUserInformRequest
	(*GetTeacherUserInformResponse)(nil), // 5: GetTeacherUserInformResponse
}
var file_auth_teacher_proto_depIdxs = []int32{
	0, // 0: Auth.LoginTeacherAuth:input_type -> LoginTeacherAuthRequest
	2, // 1: Auth.ChangeTeacherAuthPw:input_type -> ChangeTeacherAuthPwRequest
	4, // 2: Auth.GetTeacherUserInform:input_type -> GetTeacherUserInformRequest
	1, // 3: Auth.LoginTeacherAuth:output_type -> LoginTeacherAuthResponse
	3, // 4: Auth.ChangeTeacherAuthPw:output_type -> ChangeTeacherAuthPwResponse
	5, // 5: Auth.GetTeacherUserInform:output_type -> GetTeacherUserInformResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_teacher_proto_init() }
func file_auth_teacher_proto_init() {
	if File_auth_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTeacherAuthRequest); i {
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
		file_auth_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginTeacherAuthResponse); i {
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
		file_auth_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeTeacherAuthPwRequest); i {
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
		file_auth_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeTeacherAuthPwResponse); i {
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
		file_auth_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeacherUserInformRequest); i {
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
		file_auth_teacher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeacherUserInformResponse); i {
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
			RawDescriptor: file_auth_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_teacher_proto_goTypes,
		DependencyIndexes: file_auth_teacher_proto_depIdxs,
		MessageInfos:      file_auth_teacher_proto_msgTypes,
	}.Build()
	File_auth_teacher_proto = out.File
	file_auth_teacher_proto_rawDesc = nil
	file_auth_teacher_proto_goTypes = nil
	file_auth_teacher_proto_depIdxs = nil
}
