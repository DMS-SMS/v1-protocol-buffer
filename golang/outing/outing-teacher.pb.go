// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: outing-teacher.proto

package outingproto

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

type GetOutingWithFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Grade  int32  `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Group  int32  `protobuf:"varint,4,opt,name=group,proto3" json:"group,omitempty"`
	Floor  int32  `protobuf:"varint,5,opt,name=floor,proto3" json:"floor,omitempty"`
	Start  int32  `protobuf:"varint,6,opt,name=start,proto3" json:"start,omitempty"`
	Count  int32  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetOutingWithFilterRequest) Reset() {
	*x = GetOutingWithFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outing_teacher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutingWithFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutingWithFilterRequest) ProtoMessage() {}

func (x *GetOutingWithFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outing_teacher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutingWithFilterRequest.ProtoReflect.Descriptor instead.
func (*GetOutingWithFilterRequest) Descriptor() ([]byte, []int) {
	return file_outing_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *GetOutingWithFilterRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetOutingWithFilterRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetOutingWithFilterRequest) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *GetOutingWithFilterRequest) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *GetOutingWithFilterRequest) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

func (x *GetOutingWithFilterRequest) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetOutingWithFilterRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OutingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   int32     `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string    `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Outing []*Outing `protobuf:"bytes,4,rep,name=outing,proto3" json:"outing,omitempty"`
}

func (x *OutingResponse) Reset() {
	*x = OutingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outing_teacher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutingResponse) ProtoMessage() {}

func (x *OutingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outing_teacher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutingResponse.ProtoReflect.Descriptor instead.
func (*OutingResponse) Descriptor() ([]byte, []int) {
	return file_outing_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *OutingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OutingResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OutingResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OutingResponse) GetOuting() []*Outing {
	if x != nil {
		return x.Outing
	}
	return nil
}

type Outing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutingId  string `protobuf:"bytes,1,opt,name=outing_id,json=outingId,proto3" json:"outing_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Grade     int32  `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Group     int32  `protobuf:"varint,4,opt,name=group,proto3" json:"group,omitempty"`
	Number    int32  `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	Place     string `protobuf:"bytes,6,opt,name=place,proto3" json:"place,omitempty"`
	Reason    string `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	StartTime int64  `protobuf:"varint,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status    string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Situation string `protobuf:"bytes,11,opt,name=situation,proto3" json:"situation,omitempty"`
	IsLate    bool   `protobuf:"varint,12,opt,name=is_late,json=isLate,proto3" json:"is_late,omitempty"`
}

func (x *Outing) Reset() {
	*x = Outing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outing_teacher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outing) ProtoMessage() {}

func (x *Outing) ProtoReflect() protoreflect.Message {
	mi := &file_outing_teacher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outing.ProtoReflect.Descriptor instead.
func (*Outing) Descriptor() ([]byte, []int) {
	return file_outing_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *Outing) GetOutingId() string {
	if x != nil {
		return x.OutingId
	}
	return ""
}

func (x *Outing) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Outing) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Outing) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Outing) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Outing) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Outing) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Outing) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Outing) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Outing) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Outing) GetSituation() string {
	if x != nil {
		return x.Situation
	}
	return ""
}

func (x *Outing) GetIsLate() bool {
	if x != nil {
		return x.IsLate
	}
	return false
}

type ConfirmOutingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OutingId string `protobuf:"bytes,2,opt,name=outing_id,json=outingId,proto3" json:"outing_id,omitempty"`
}

func (x *ConfirmOutingRequest) Reset() {
	*x = ConfirmOutingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outing_teacher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmOutingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmOutingRequest) ProtoMessage() {}

func (x *ConfirmOutingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outing_teacher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmOutingRequest.ProtoReflect.Descriptor instead.
func (*ConfirmOutingRequest) Descriptor() ([]byte, []int) {
	return file_outing_teacher_proto_rawDescGZIP(), []int{3}
}

func (x *ConfirmOutingRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ConfirmOutingRequest) GetOutingId() string {
	if x != nil {
		return x.OutingId
	}
	return ""
}

type ConfirmOutingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ConfirmOutingResponse) Reset() {
	*x = ConfirmOutingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outing_teacher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmOutingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmOutingResponse) ProtoMessage() {}

func (x *ConfirmOutingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outing_teacher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmOutingResponse.ProtoReflect.Descriptor instead.
func (*ConfirmOutingResponse) Descriptor() ([]byte, []int) {
	return file_outing_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmOutingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ConfirmOutingResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConfirmOutingResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_outing_teacher_proto protoreflect.FileDescriptor

var file_outing_teacher_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x6f, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x1f, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0xb4, 0x02, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x4c, 0x61, 0x74, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x55, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x9b, 0x02, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x74, 0x68,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outing_teacher_proto_rawDescOnce sync.Once
	file_outing_teacher_proto_rawDescData = file_outing_teacher_proto_rawDesc
)

func file_outing_teacher_proto_rawDescGZIP() []byte {
	file_outing_teacher_proto_rawDescOnce.Do(func() {
		file_outing_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(file_outing_teacher_proto_rawDescData)
	})
	return file_outing_teacher_proto_rawDescData
}

var file_outing_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_outing_teacher_proto_goTypes = []interface{}{
	(*GetOutingWithFilterRequest)(nil), // 0: GetOutingWithFilterRequest
	(*OutingResponse)(nil),             // 1: OutingResponse
	(*Outing)(nil),                     // 2: Outing
	(*ConfirmOutingRequest)(nil),       // 3: ConfirmOutingRequest
	(*ConfirmOutingResponse)(nil),      // 4: ConfirmOutingResponse
}
var file_outing_teacher_proto_depIdxs = []int32{
	2, // 0: OutingResponse.outing:type_name -> Outing
	0, // 1: OutingTeacher.GetOutingWithFilter:input_type -> GetOutingWithFilterRequest
	3, // 2: OutingTeacher.ApproveOuting:input_type -> ConfirmOutingRequest
	3, // 3: OutingTeacher.RejectOuting:input_type -> ConfirmOutingRequest
	3, // 4: OutingTeacher.CertifyOuting:input_type -> ConfirmOutingRequest
	1, // 5: OutingTeacher.GetOutingWithFilter:output_type -> OutingResponse
	4, // 6: OutingTeacher.ApproveOuting:output_type -> ConfirmOutingResponse
	4, // 7: OutingTeacher.RejectOuting:output_type -> ConfirmOutingResponse
	4, // 8: OutingTeacher.CertifyOuting:output_type -> ConfirmOutingResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_outing_teacher_proto_init() }
func file_outing_teacher_proto_init() {
	if File_outing_teacher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outing_teacher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutingWithFilterRequest); i {
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
		file_outing_teacher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutingResponse); i {
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
		file_outing_teacher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outing); i {
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
		file_outing_teacher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmOutingRequest); i {
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
		file_outing_teacher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmOutingResponse); i {
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
			RawDescriptor: file_outing_teacher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_outing_teacher_proto_goTypes,
		DependencyIndexes: file_outing_teacher_proto_depIdxs,
		MessageInfos:      file_outing_teacher_proto_msgTypes,
	}.Build()
	File_outing_teacher_proto = out.File
	file_outing_teacher_proto_rawDesc = nil
	file_outing_teacher_proto_goTypes = nil
	file_outing_teacher_proto_depIdxs = nil
}