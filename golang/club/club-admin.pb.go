// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: club-admin.proto

package clubproto

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

type CreateNewClubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID        string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`               // API를 호출한 관리자 계정의 UUID (18자)
	Name        string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`               // 동아리 이름 (1~20자 사이)
	LeaderUUID  string   `protobuf:"bytes,3,opt,name=LeaderUUID,proto3" json:"LeaderUUID,omitempty"`   // 동아리 장의 계정 UUID (20자)
	MemberUUIDs []string `protobuf:"bytes,4,rep,name=MemberUUIDs,proto3" json:"MemberUUIDs,omitempty"` // 동아리 부원의 계정 UUID List (20자)
	Field       string   `protobuf:"bytes,5,opt,name=Field,proto3" json:"Field,omitempty"`             // 동아리 분야 (1~20자 이내)
	Location    string   `protobuf:"bytes,6,opt,name=Location,proto3" json:"Location,omitempty"`       // 동아리 위치 (1~20자 이내)
	Floor       string   `protobuf:"bytes,7,opt,name=Floor,proto3" json:"Floor,omitempty"`             // 동아리실 층 (1~5 사이 값)
	Logo        []byte   `protobuf:"bytes,8,opt,name=Logo,proto3" json:"Logo,omitempty"`               // 동아리 로고 바이트 배열
}

func (x *CreateNewClubRequest) Reset() {
	*x = CreateNewClubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewClubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewClubRequest) ProtoMessage() {}

func (x *CreateNewClubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_club_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewClubRequest.ProtoReflect.Descriptor instead.
func (*CreateNewClubRequest) Descriptor() ([]byte, []int) {
	return file_club_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNewClubRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CreateNewClubRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNewClubRequest) GetLeaderUUID() string {
	if x != nil {
		return x.LeaderUUID
	}
	return ""
}

func (x *CreateNewClubRequest) GetMemberUUIDs() []string {
	if x != nil {
		return x.MemberUUIDs
	}
	return nil
}

func (x *CreateNewClubRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *CreateNewClubRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateNewClubRequest) GetFloor() string {
	if x != nil {
		return x.Floor
	}
	return ""
}

func (x *CreateNewClubRequest) GetLogo() []byte {
	if x != nil {
		return x.Logo
	}
	return nil
}

type CreateNewClubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   uint32 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`    // 상태 코드
	Code     int32  `protobuf:"zigzag32,2,opt,name=Code,proto3" json:"Code,omitempty"`      // 세부 코드
	Message  string `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`   // 세부 설명
	ClubUUID string `protobuf:"bytes,4,opt,name=ClubUUID,proto3" json:"ClubUUID,omitempty"` // 생성된 동아리의 UUID
}

func (x *CreateNewClubResponse) Reset() {
	*x = CreateNewClubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_club_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewClubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewClubResponse) ProtoMessage() {}

func (x *CreateNewClubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_club_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewClubResponse.ProtoReflect.Descriptor instead.
func (*CreateNewClubResponse) Descriptor() ([]byte, []int) {
	return file_club_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNewClubResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateNewClubResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateNewClubResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateNewClubResponse) GetClubUUID() string {
	if x != nil {
		return x.ClubUUID
	}
	return ""
}

var File_club_admin_proto protoreflect.FileDescriptor

var file_club_admin_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x75, 0x62, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xdc,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x75, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x22, 0x79, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x75, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6c, 0x75, 0x62, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6c, 0x75, 0x62, 0x55, 0x55, 0x49, 0x44, 0x32, 0x63, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x62,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x43, 0x6c, 0x75, 0x62, 0x12, 0x20, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x62, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43,
	0x6c, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x3b, 0x63, 0x6c, 0x75, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_club_admin_proto_rawDescOnce sync.Once
	file_club_admin_proto_rawDescData = file_club_admin_proto_rawDesc
)

func file_club_admin_proto_rawDescGZIP() []byte {
	file_club_admin_proto_rawDescOnce.Do(func() {
		file_club_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_club_admin_proto_rawDescData)
	})
	return file_club_admin_proto_rawDescData
}

var file_club_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_club_admin_proto_goTypes = []interface{}{
	(*CreateNewClubRequest)(nil),  // 0: club.admin.CreateNewClubRequest
	(*CreateNewClubResponse)(nil), // 1: club.admin.CreateNewClubResponse
}
var file_club_admin_proto_depIdxs = []int32{
	0, // 0: club.admin.ClubAdmin.CreateNewClub:input_type -> club.admin.CreateNewClubRequest
	1, // 1: club.admin.ClubAdmin.CreateNewClub:output_type -> club.admin.CreateNewClubResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_club_admin_proto_init() }
func file_club_admin_proto_init() {
	if File_club_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_club_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewClubRequest); i {
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
		file_club_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewClubResponse); i {
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
			RawDescriptor: file_club_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_club_admin_proto_goTypes,
		DependencyIndexes: file_club_admin_proto_depIdxs,
		MessageInfos:      file_club_admin_proto_msgTypes,
	}.Build()
	File_club_admin_proto = out.File
	file_club_admin_proto_rawDesc = nil
	file_club_admin_proto_goTypes = nil
	file_club_admin_proto_depIdxs = nil
}
