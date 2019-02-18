// Code generated by protoc-gen-go. DO NOT EDIT.
// source: birthday.proto

package birthdaypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SetBirthdayRequest struct {
	ServerId             string   `protobuf:"bytes,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Month                int32    `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"varint,4,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBirthdayRequest) Reset()         { *m = SetBirthdayRequest{} }
func (m *SetBirthdayRequest) String() string { return proto.CompactTextString(m) }
func (*SetBirthdayRequest) ProtoMessage()    {}
func (*SetBirthdayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_birthday_76c9f9a9b39dc045, []int{0}
}
func (m *SetBirthdayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBirthdayRequest.Unmarshal(m, b)
}
func (m *SetBirthdayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBirthdayRequest.Marshal(b, m, deterministic)
}
func (dst *SetBirthdayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBirthdayRequest.Merge(dst, src)
}
func (m *SetBirthdayRequest) XXX_Size() int {
	return xxx_messageInfo_SetBirthdayRequest.Size(m)
}
func (m *SetBirthdayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBirthdayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBirthdayRequest proto.InternalMessageInfo

func (m *SetBirthdayRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *SetBirthdayRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SetBirthdayRequest) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *SetBirthdayRequest) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type UserBirthday struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerId             string   `protobuf:"bytes,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Month                int32    `protobuf:"varint,4,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"varint,5,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBirthday) Reset()         { *m = UserBirthday{} }
func (m *UserBirthday) String() string { return proto.CompactTextString(m) }
func (*UserBirthday) ProtoMessage()    {}
func (*UserBirthday) Descriptor() ([]byte, []int) {
	return fileDescriptor_birthday_76c9f9a9b39dc045, []int{1}
}
func (m *UserBirthday) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBirthday.Unmarshal(m, b)
}
func (m *UserBirthday) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBirthday.Marshal(b, m, deterministic)
}
func (dst *UserBirthday) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBirthday.Merge(dst, src)
}
func (m *UserBirthday) XXX_Size() int {
	return xxx_messageInfo_UserBirthday.Size(m)
}
func (m *UserBirthday) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBirthday.DiscardUnknown(m)
}

var xxx_messageInfo_UserBirthday proto.InternalMessageInfo

func (m *UserBirthday) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserBirthday) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *UserBirthday) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserBirthday) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *UserBirthday) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type SetBirthdayReply struct {
	Birthday             *UserBirthday `protobuf:"bytes,1,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SetBirthdayReply) Reset()         { *m = SetBirthdayReply{} }
func (m *SetBirthdayReply) String() string { return proto.CompactTextString(m) }
func (*SetBirthdayReply) ProtoMessage()    {}
func (*SetBirthdayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_birthday_76c9f9a9b39dc045, []int{2}
}
func (m *SetBirthdayReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBirthdayReply.Unmarshal(m, b)
}
func (m *SetBirthdayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBirthdayReply.Marshal(b, m, deterministic)
}
func (dst *SetBirthdayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBirthdayReply.Merge(dst, src)
}
func (m *SetBirthdayReply) XXX_Size() int {
	return xxx_messageInfo_SetBirthdayReply.Size(m)
}
func (m *SetBirthdayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBirthdayReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetBirthdayReply proto.InternalMessageInfo

func (m *SetBirthdayReply) GetBirthday() *UserBirthday {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func init() {
	proto.RegisterType((*SetBirthdayRequest)(nil), "birthdaypb.SetBirthdayRequest")
	proto.RegisterType((*UserBirthday)(nil), "birthdaypb.UserBirthday")
	proto.RegisterType((*SetBirthdayReply)(nil), "birthdaypb.SetBirthdayReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BirthdayClient is the client API for Birthday service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BirthdayClient interface {
	SetBirthday(ctx context.Context, in *SetBirthdayRequest, opts ...grpc.CallOption) (*SetBirthdayReply, error)
}

type birthdayClient struct {
	cc *grpc.ClientConn
}

func NewBirthdayClient(cc *grpc.ClientConn) BirthdayClient {
	return &birthdayClient{cc}
}

func (c *birthdayClient) SetBirthday(ctx context.Context, in *SetBirthdayRequest, opts ...grpc.CallOption) (*SetBirthdayReply, error) {
	out := new(SetBirthdayReply)
	err := c.cc.Invoke(ctx, "/birthdaypb.Birthday/SetBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirthdayServer is the server API for Birthday service.
type BirthdayServer interface {
	SetBirthday(context.Context, *SetBirthdayRequest) (*SetBirthdayReply, error)
}

func RegisterBirthdayServer(s *grpc.Server, srv BirthdayServer) {
	s.RegisterService(&_Birthday_serviceDesc, srv)
}

func _Birthday_SetBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayServer).SetBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthdaypb.Birthday/SetBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayServer).SetBirthday(ctx, req.(*SetBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Birthday_serviceDesc = grpc.ServiceDesc{
	ServiceName: "birthdaypb.Birthday",
	HandlerType: (*BirthdayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBirthday",
			Handler:    _Birthday_SetBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "birthday.proto",
}

func init() { proto.RegisterFile("birthday.proto", fileDescriptor_birthday_76c9f9a9b39dc045) }

var fileDescriptor_birthday_76c9f9a9b39dc045 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xca, 0x2c, 0x2a,
	0xc9, 0x48, 0x49, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xf1, 0x0b, 0x92,
	0x94, 0x0a, 0xb8, 0x84, 0x82, 0x53, 0x4b, 0x9c, 0xa0, 0x02, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x52, 0x5c, 0x1c, 0xc5, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x9e, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x18, 0x17, 0x5b, 0x69, 0x31, 0x58, 0x86, 0x09, 0x2c,
	0x03, 0xe5, 0x09, 0x89, 0x70, 0xb1, 0xe6, 0xe6, 0xe7, 0x95, 0x64, 0x48, 0x30, 0x2b, 0x30, 0x6a,
	0xb0, 0x06, 0x41, 0x38, 0x42, 0x02, 0x5c, 0xcc, 0x29, 0x89, 0x95, 0x12, 0x2c, 0x60, 0x31, 0x10,
	0x53, 0xa9, 0x8a, 0x8b, 0x27, 0xb4, 0x38, 0xb5, 0x08, 0x66, 0xa5, 0x10, 0x1f, 0x17, 0x53, 0x26,
	0xcc, 0x16, 0xa6, 0xcc, 0x14, 0x14, 0xbb, 0x99, 0x70, 0xda, 0xcd, 0x8c, 0xdd, 0x6e, 0x16, 0x2c,
	0x76, 0xb3, 0x22, 0xec, 0xf6, 0xe0, 0x12, 0x40, 0xf1, 0x6d, 0x41, 0x4e, 0xa5, 0x90, 0x09, 0x17,
	0x07, 0x2c, 0x3c, 0xc0, 0xae, 0xe0, 0x36, 0x92, 0xd0, 0x43, 0x04, 0x90, 0x1e, 0xb2, 0x5b, 0x83,
	0xe0, 0x2a, 0x8d, 0x22, 0xb9, 0x38, 0xe0, 0x3e, 0xf0, 0xe5, 0xe2, 0x46, 0x32, 0x55, 0x48, 0x0e,
	0x59, 0x3b, 0x66, 0xe0, 0x4a, 0xc9, 0xe0, 0x94, 0x2f, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03,
	0xc7, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xad, 0xb2, 0xc2, 0xb7, 0x01, 0x00, 0x00,
}
