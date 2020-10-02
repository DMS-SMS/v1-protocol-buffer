// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: auth-admin.proto

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

type CreateNewStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID          string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`                    // API를 호출한 관리자 uuid (18자)
	StudentID     string `protobuf:"bytes,2,opt,name=StudentID,proto3" json:"StudentID,omitempty"`          // 학생 계정 아이디 (4~16자 사이)
	StudentPW     string `protobuf:"bytes,3,opt,name=StudentPW,proto3" json:"StudentPW,omitempty"`          // 학생 계정 비밀번호 (4~16자 사이)
	ParentUUID    string `protobuf:"bytes,4,opt,name=ParentUUID,proto3" json:"ParentUUID,omitempty"`        // 부모님 설정을 위한 부모님 계정의 uuid (19자)
	Grade         uint32 `protobuf:"varint,5,opt,name=Grade,proto3" json:"Grade,omitempty"`                 // 학년 정보 (1~3 사이)
	Group         uint32 `protobuf:"varint,6,opt,name=Group,proto3" json:"Group,omitempty"`                 // 반 정보 (1~4 사이)
	StudentNumber uint32 `protobuf:"varint,7,opt,name=StudentNumber,proto3" json:"StudentNumber,omitempty"` // 번호 정보 (1~21 사이)
	Name          string `protobuf:"bytes,8,opt,name=Name,proto3" json:"Name,omitempty"`                    // 이름 정보 (2~4자 사이, 한글)
	PhoneNumber   string `protobuf:"bytes,9,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`      // 전화번호 정보 (11자)
	Image         []byte `protobuf:"bytes,10,opt,name=Image,proto3" json:"Image,omitempty"`                 // 학생증 사진 바이트 배열 (?)
}

func (x *CreateNewStudentRequest) Reset() {
	*x = CreateNewStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewStudentRequest) ProtoMessage() {}

func (x *CreateNewStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateNewStudentRequest) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNewStudentRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateNewStudentRequest) GetStudentID() string {
	if x != nil {
		return x.StudentID
	}
	return ""
}

func (x *CreateNewStudentRequest) GetStudentPW() string {
	if x != nil {
		return x.StudentPW
	}
	return ""
}

func (x *CreateNewStudentRequest) GetParentUUID() string {
	if x != nil {
		return x.ParentUUID
	}
	return ""
}

func (x *CreateNewStudentRequest) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *CreateNewStudentRequest) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *CreateNewStudentRequest) GetStudentNumber() uint32 {
	if x != nil {
		return x.StudentNumber
	}
	return 0
}

func (x *CreateNewStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNewStudentRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateNewStudentRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type CreateNewStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`                        // 상태 코드
	Code               int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`                          // 세부 코드
	Message            string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`                       // 세부 설명
	CreatedStudentUUID string `protobuf:"bytes,4,opt,name=CreatedStudentUUID,proto3" json:"CreatedStudentUUID,omitempty"` // 생성된 학생 계정의 uuid
}

func (x *CreateNewStudentResponse) Reset() {
	*x = CreateNewStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewStudentResponse) ProtoMessage() {}

func (x *CreateNewStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewStudentResponse.ProtoReflect.Descriptor instead.
func (*CreateNewStudentResponse) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNewStudentResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateNewStudentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateNewStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateNewStudentResponse) GetCreatedStudentUUID() string {
	if x != nil {
		return x.CreatedStudentUUID
	}
	return ""
}

type CreateNewTeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID        string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`               // API를 호출한 관리자 uuid (18자)
	TeacherID   string `protobuf:"bytes,2,opt,name=TeacherID,proto3" json:"TeacherID,omitempty"`     // 선생님 계정 아이디 (4~16자 사이)
	TeacherPW   string `protobuf:"bytes,3,opt,name=TeacherPW,proto3" json:"TeacherPW,omitempty"`     // 선생님 계정 비밀번호 (4~16자 사이)
	Grade       uint32 `protobuf:"varint,4,opt,name=Grade,proto3" json:"Grade,omitempty"`            // 담임 선생님 학년 정보 (1~3 사이)
	Group       uint32 `protobuf:"varint,5,opt,name=Group,proto3" json:"Group,omitempty"`            // 담임 선생님 반 정보 (1~4 사이)
	Name        string `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`               // 이름 정보 (2~4자 사이, 한글)
	PhoneNumber string `protobuf:"bytes,8,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"` // 전화번호 정보 (11자)
}

func (x *CreateNewTeacherRequest) Reset() {
	*x = CreateNewTeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewTeacherRequest) ProtoMessage() {}

func (x *CreateNewTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewTeacherRequest.ProtoReflect.Descriptor instead.
func (*CreateNewTeacherRequest) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNewTeacherRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateNewTeacherRequest) GetTeacherID() string {
	if x != nil {
		return x.TeacherID
	}
	return ""
}

func (x *CreateNewTeacherRequest) GetTeacherPW() string {
	if x != nil {
		return x.TeacherPW
	}
	return ""
}

func (x *CreateNewTeacherRequest) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *CreateNewTeacherRequest) GetGroup() uint32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *CreateNewTeacherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNewTeacherRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type CreateNewTeacherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`                        // 상태 코드
	Code               int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`                          // 세부 코드
	Message            string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`                       // 세부 설명
	CreatedTeacherUUID string `protobuf:"bytes,4,opt,name=CreatedTeacherUUID,proto3" json:"CreatedTeacherUUID,omitempty"` // 생성된 선생님 계정의 uuid
}

func (x *CreateNewTeacherResponse) Reset() {
	*x = CreateNewTeacherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewTeacherResponse) ProtoMessage() {}

func (x *CreateNewTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewTeacherResponse.ProtoReflect.Descriptor instead.
func (*CreateNewTeacherResponse) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNewTeacherResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateNewTeacherResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateNewTeacherResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateNewTeacherResponse) GetCreatedTeacherUUID() string {
	if x != nil {
		return x.CreatedTeacherUUID
	}
	return ""
}

type CreateNewParentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID        string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`               // API를 호출한 관리자 uuid (18자)
	ParentID    string `protobuf:"bytes,2,opt,name=ParentID,proto3" json:"ParentID,omitempty"`       // 선생님 계정 아이디 (4~16자 사이)
	ParentPW    string `protobuf:"bytes,3,opt,name=ParentPW,proto3" json:"ParentPW,omitempty"`       // 선생님 계정 비밀번호 (4~16자 사이)
	Name        string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`               // 이름 정보 (2~4자 사이, 한글)
	PhoneNumber string `protobuf:"bytes,6,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"` // 전화번호 정보 (11자)
}

func (x *CreateNewParentRequest) Reset() {
	*x = CreateNewParentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewParentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewParentRequest) ProtoMessage() {}

func (x *CreateNewParentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewParentRequest.ProtoReflect.Descriptor instead.
func (*CreateNewParentRequest) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNewParentRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateNewParentRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *CreateNewParentRequest) GetParentPW() string {
	if x != nil {
		return x.ParentPW
	}
	return ""
}

func (x *CreateNewParentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNewParentRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type CreateNewParentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`                      // 상태 코드
	Code              int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`                        // 세부 코드
	Message           string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`                     // 세부 설명
	CreatedParentUUID string `protobuf:"bytes,4,opt,name=CreatedParentUUID,proto3" json:"CreatedParentUUID,omitempty"` // 생성된 학부모 계정의 uuid
}

func (x *CreateNewParentResponse) Reset() {
	*x = CreateNewParentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewParentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewParentResponse) ProtoMessage() {}

func (x *CreateNewParentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewParentResponse.ProtoReflect.Descriptor instead.
func (*CreateNewParentResponse) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{5}
}

func (x *CreateNewParentResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateNewParentResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateNewParentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateNewParentResponse) GetCreatedParentUUID() string {
	if x != nil {
		return x.CreatedParentUUID
	}
	return ""
}

type LoginAdminAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminID string `protobuf:"bytes,1,opt,name=AdminID,proto3" json:"AdminID,omitempty"` // 관리자 계정 아이디 (4~16자 사이)
	AdminPW string `protobuf:"bytes,2,opt,name=AdminPW,proto3" json:"AdminPW,omitempty"` // 관리자 계정 비밀번호 (4~16자 사이)
}

func (x *LoginAdminAuthRequest) Reset() {
	*x = LoginAdminAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAdminAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAdminAuthRequest) ProtoMessage() {}

func (x *LoginAdminAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAdminAuthRequest.ProtoReflect.Descriptor instead.
func (*LoginAdminAuthRequest) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{6}
}

func (x *LoginAdminAuthRequest) GetAdminID() string {
	if x != nil {
		return x.AdminID
	}
	return ""
}

func (x *LoginAdminAuthRequest) GetAdminPW() string {
	if x != nil {
		return x.AdminPW
	}
	return ""
}

type LoginAdminAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`                      // 상태 코드
	Code              int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`                        // 세부 코드
	Message           string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`                     // 세부 설명
	AccessToken       string `protobuf:"bytes,4,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`             // 학생 계정 인증을 위한 JWT
	LoggedInAdminUUID string `protobuf:"bytes,5,opt,name=LoggedInAdminUUID,proto3" json:"LoggedInAdminUUID,omitempty"` // 로그인한 관리자 계정의 uuid
}

func (x *LoginAdminAuthResponse) Reset() {
	*x = LoginAdminAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAdminAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAdminAuthResponse) ProtoMessage() {}

func (x *LoginAdminAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAdminAuthResponse.ProtoReflect.Descriptor instead.
func (*LoginAdminAuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_admin_proto_rawDescGZIP(), []int{7}
}

func (x *LoginAdminAuthResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginAdminAuthResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginAdminAuthResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginAdminAuthResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginAdminAuthResponse) GetLoggedInAdminUUID() string {
	if x != nil {
		return x.LoggedInAdminUUID
	}
	return ""
}

var File_auth_admin_proto protoreflect.FileDescriptor

var file_auth_admin_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x90, 0x01, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22,
	0xcb, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x57, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x57, 0x12, 0x14, 0x0a, 0x05, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x90, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44,
	0x22, 0x9a, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x57, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x8d, 0x01,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44, 0x22, 0x4b, 0x0a,
	0x15, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x57, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x57, 0x22, 0xae, 0x01, 0x0a, 0x16, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a,
	0x11, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x32, 0xae, 0x02, 0x0a, 0x09,
	0x41, 0x75, 0x74, 0x68, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x49, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_admin_proto_rawDescOnce sync.Once
	file_auth_admin_proto_rawDescData = file_auth_admin_proto_rawDesc
)

func file_auth_admin_proto_rawDescGZIP() []byte {
	file_auth_admin_proto_rawDescOnce.Do(func() {
		file_auth_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_admin_proto_rawDescData)
	})
	return file_auth_admin_proto_rawDescData
}

var file_auth_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_admin_proto_goTypes = []interface{}{
	(*CreateNewStudentRequest)(nil),  // 0: CreateNewStudentRequest
	(*CreateNewStudentResponse)(nil), // 1: CreateNewStudentResponse
	(*CreateNewTeacherRequest)(nil),  // 2: CreateNewTeacherRequest
	(*CreateNewTeacherResponse)(nil), // 3: CreateNewTeacherResponse
	(*CreateNewParentRequest)(nil),   // 4: CreateNewParentRequest
	(*CreateNewParentResponse)(nil),  // 5: CreateNewParentResponse
	(*LoginAdminAuthRequest)(nil),    // 6: LoginAdminAuthRequest
	(*LoginAdminAuthResponse)(nil),   // 7: LoginAdminAuthResponse
}
var file_auth_admin_proto_depIdxs = []int32{
	0, // 0: AuthAdmin.CreateNewStudent:input_type -> CreateNewStudentRequest
	2, // 1: AuthAdmin.CreateNewTeacher:input_type -> CreateNewTeacherRequest
	4, // 2: AuthAdmin.CreateNewParent:input_type -> CreateNewParentRequest
	6, // 3: AuthAdmin.LoginAdminAuth:input_type -> LoginAdminAuthRequest
	1, // 4: AuthAdmin.CreateNewStudent:output_type -> CreateNewStudentResponse
	3, // 5: AuthAdmin.CreateNewTeacher:output_type -> CreateNewTeacherResponse
	5, // 6: AuthAdmin.CreateNewParent:output_type -> CreateNewParentResponse
	7, // 7: AuthAdmin.LoginAdminAuth:output_type -> LoginAdminAuthResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_admin_proto_init() }
func file_auth_admin_proto_init() {
	if File_auth_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewStudentRequest); i {
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
		file_auth_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewStudentResponse); i {
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
		file_auth_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewTeacherRequest); i {
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
		file_auth_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewTeacherResponse); i {
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
		file_auth_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewParentRequest); i {
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
		file_auth_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewParentResponse); i {
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
		file_auth_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAdminAuthRequest); i {
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
		file_auth_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAdminAuthResponse); i {
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
			RawDescriptor: file_auth_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_admin_proto_goTypes,
		DependencyIndexes: file_auth_admin_proto_depIdxs,
		MessageInfos:      file_auth_admin_proto_msgTypes,
	}.Build()
	File_auth_admin_proto = out.File
	file_auth_admin_proto_rawDesc = nil
	file_auth_admin_proto_goTypes = nil
	file_auth_admin_proto_depIdxs = nil
}
