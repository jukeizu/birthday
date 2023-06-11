// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: birthday.proto

package birthdaypb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetBirthdayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Month    int32  `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
	Day      int32  `protobuf:"varint,4,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *SetBirthdayRequest) Reset() {
	*x = SetBirthdayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBirthdayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBirthdayRequest) ProtoMessage() {}

func (x *SetBirthdayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBirthdayRequest.ProtoReflect.Descriptor instead.
func (*SetBirthdayRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{0}
}

func (x *SetBirthdayRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *SetBirthdayRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetBirthdayRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *SetBirthdayRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type UserBirthday struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Month    int32  `protobuf:"varint,4,opt,name=month,proto3" json:"month,omitempty"`
	Day      int32  `protobuf:"varint,5,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *UserBirthday) Reset() {
	*x = UserBirthday{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBirthday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBirthday) ProtoMessage() {}

func (x *UserBirthday) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBirthday.ProtoReflect.Descriptor instead.
func (*UserBirthday) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{1}
}

func (x *UserBirthday) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserBirthday) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UserBirthday) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserBirthday) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *UserBirthday) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type SetBirthdayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Birthday *UserBirthday `protobuf:"bytes,1,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *SetBirthdayResponse) Reset() {
	*x = SetBirthdayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBirthdayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBirthdayResponse) ProtoMessage() {}

func (x *SetBirthdayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBirthdayResponse.ProtoReflect.Descriptor instead.
func (*SetBirthdayResponse) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{2}
}

func (x *SetBirthdayResponse) GetBirthday() *UserBirthday {
	if x != nil {
		return x.Birthday
	}
	return nil
}

type QueryBirthdaysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Month    int32  `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day      int32  `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *QueryBirthdaysRequest) Reset() {
	*x = QueryBirthdaysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBirthdaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBirthdaysRequest) ProtoMessage() {}

func (x *QueryBirthdaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBirthdaysRequest.ProtoReflect.Descriptor instead.
func (*QueryBirthdaysRequest) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{3}
}

func (x *QueryBirthdaysRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *QueryBirthdaysRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *QueryBirthdaysRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type QueryBirthdaysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Birthdays []*UserBirthday `protobuf:"bytes,1,rep,name=birthdays,proto3" json:"birthdays,omitempty"`
}

func (x *QueryBirthdaysResponse) Reset() {
	*x = QueryBirthdaysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_birthday_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBirthdaysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBirthdaysResponse) ProtoMessage() {}

func (x *QueryBirthdaysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_birthday_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBirthdaysResponse.ProtoReflect.Descriptor instead.
func (*QueryBirthdaysResponse) Descriptor() ([]byte, []int) {
	return file_birthday_proto_rawDescGZIP(), []int{4}
}

func (x *QueryBirthdaysResponse) GetBirthdays() []*UserBirthday {
	if x != nil {
		return x.Birthdays
	}
	return nil
}

var File_birthday_proto protoreflect.FileDescriptor

var file_birthday_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x72, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x22, 0x7c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22,
	0x4c, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x5c, 0x0a,
	0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x51, 0x0a, 0x16, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x73, 0x32, 0xc2,
	0x01, 0x0a, 0x0f, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x12, 0x1f, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_birthday_proto_rawDescOnce sync.Once
	file_birthday_proto_rawDescData = file_birthday_proto_rawDesc
)

func file_birthday_proto_rawDescGZIP() []byte {
	file_birthday_proto_rawDescOnce.Do(func() {
		file_birthday_proto_rawDescData = protoimpl.X.CompressGZIP(file_birthday_proto_rawDescData)
	})
	return file_birthday_proto_rawDescData
}

var file_birthday_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_birthday_proto_goTypes = []interface{}{
	(*SetBirthdayRequest)(nil),     // 0: birthday.v1.SetBirthdayRequest
	(*UserBirthday)(nil),           // 1: birthday.v1.UserBirthday
	(*SetBirthdayResponse)(nil),    // 2: birthday.v1.SetBirthdayResponse
	(*QueryBirthdaysRequest)(nil),  // 3: birthday.v1.QueryBirthdaysRequest
	(*QueryBirthdaysResponse)(nil), // 4: birthday.v1.QueryBirthdaysResponse
}
var file_birthday_proto_depIdxs = []int32{
	1, // 0: birthday.v1.SetBirthdayResponse.birthday:type_name -> birthday.v1.UserBirthday
	1, // 1: birthday.v1.QueryBirthdaysResponse.birthdays:type_name -> birthday.v1.UserBirthday
	0, // 2: birthday.v1.BirthdayService.SetBirthday:input_type -> birthday.v1.SetBirthdayRequest
	3, // 3: birthday.v1.BirthdayService.QueryBirthdays:input_type -> birthday.v1.QueryBirthdaysRequest
	2, // 4: birthday.v1.BirthdayService.SetBirthday:output_type -> birthday.v1.SetBirthdayResponse
	4, // 5: birthday.v1.BirthdayService.QueryBirthdays:output_type -> birthday.v1.QueryBirthdaysResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_birthday_proto_init() }
func file_birthday_proto_init() {
	if File_birthday_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_birthday_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBirthdayRequest); i {
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
		file_birthday_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBirthday); i {
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
		file_birthday_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBirthdayResponse); i {
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
		file_birthday_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBirthdaysRequest); i {
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
		file_birthday_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBirthdaysResponse); i {
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
			RawDescriptor: file_birthday_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_birthday_proto_goTypes,
		DependencyIndexes: file_birthday_proto_depIdxs,
		MessageInfos:      file_birthday_proto_msgTypes,
	}.Build()
	File_birthday_proto = out.File
	file_birthday_proto_rawDesc = nil
	file_birthday_proto_goTypes = nil
	file_birthday_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BirthdayServiceClient is the client API for BirthdayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BirthdayServiceClient interface {
	SetBirthday(ctx context.Context, in *SetBirthdayRequest, opts ...grpc.CallOption) (*SetBirthdayResponse, error)
	QueryBirthdays(ctx context.Context, in *QueryBirthdaysRequest, opts ...grpc.CallOption) (*QueryBirthdaysResponse, error)
}

type birthdayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBirthdayServiceClient(cc grpc.ClientConnInterface) BirthdayServiceClient {
	return &birthdayServiceClient{cc}
}

func (c *birthdayServiceClient) SetBirthday(ctx context.Context, in *SetBirthdayRequest, opts ...grpc.CallOption) (*SetBirthdayResponse, error) {
	out := new(SetBirthdayResponse)
	err := c.cc.Invoke(ctx, "/birthday.v1.BirthdayService/SetBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birthdayServiceClient) QueryBirthdays(ctx context.Context, in *QueryBirthdaysRequest, opts ...grpc.CallOption) (*QueryBirthdaysResponse, error) {
	out := new(QueryBirthdaysResponse)
	err := c.cc.Invoke(ctx, "/birthday.v1.BirthdayService/QueryBirthdays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirthdayServiceServer is the server API for BirthdayService service.
type BirthdayServiceServer interface {
	SetBirthday(context.Context, *SetBirthdayRequest) (*SetBirthdayResponse, error)
	QueryBirthdays(context.Context, *QueryBirthdaysRequest) (*QueryBirthdaysResponse, error)
}

// UnimplementedBirthdayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBirthdayServiceServer struct {
}

func (*UnimplementedBirthdayServiceServer) SetBirthday(context.Context, *SetBirthdayRequest) (*SetBirthdayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBirthday not implemented")
}
func (*UnimplementedBirthdayServiceServer) QueryBirthdays(context.Context, *QueryBirthdaysRequest) (*QueryBirthdaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBirthdays not implemented")
}

func RegisterBirthdayServiceServer(s *grpc.Server, srv BirthdayServiceServer) {
	s.RegisterService(&_BirthdayService_serviceDesc, srv)
}

func _BirthdayService_SetBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayServiceServer).SetBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.v1.BirthdayService/SetBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayServiceServer).SetBirthday(ctx, req.(*SetBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirthdayService_QueryBirthdays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBirthdaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayServiceServer).QueryBirthdays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.v1.BirthdayService/QueryBirthdays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayServiceServer).QueryBirthdays(ctx, req.(*QueryBirthdaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BirthdayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "birthday.v1.BirthdayService",
	HandlerType: (*BirthdayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBirthday",
			Handler:    _BirthdayService_SetBirthday_Handler,
		},
		{
			MethodName: "QueryBirthdays",
			Handler:    _BirthdayService_QueryBirthdays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "birthday.proto",
}
