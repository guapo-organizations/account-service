// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//数据库,基本的账户信息
type UserBaseInfo struct {
	//id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//账号
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//邮箱
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	//手机号
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	//用户的真实姓名
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBaseInfo) Reset()         { *m = UserBaseInfo{} }
func (m *UserBaseInfo) String() string { return proto.CompactTextString(m) }
func (*UserBaseInfo) ProtoMessage()    {}
func (*UserBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
}

func (m *UserBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBaseInfo.Unmarshal(m, b)
}
func (m *UserBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBaseInfo.Marshal(b, m, deterministic)
}
func (m *UserBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBaseInfo.Merge(m, src)
}
func (m *UserBaseInfo) XXX_Size() int {
	return xxx_messageInfo_UserBaseInfo.Size(m)
}
func (m *UserBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserBaseInfo proto.InternalMessageInfo

func (m *UserBaseInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserBaseInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserBaseInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserBaseInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserBaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//数据库,用户绑定的三方账号
type UserPlatformInfo struct {
	//账户类型
	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	//账户id
	AccountId int64 `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	//三方平台标识的用户id
	PlatformId           string   `protobuf:"bytes,3,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPlatformInfo) Reset()         { *m = UserPlatformInfo{} }
func (m *UserPlatformInfo) String() string { return proto.CompactTextString(m) }
func (*UserPlatformInfo) ProtoMessage()    {}
func (*UserPlatformInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *UserPlatformInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPlatformInfo.Unmarshal(m, b)
}
func (m *UserPlatformInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPlatformInfo.Marshal(b, m, deterministic)
}
func (m *UserPlatformInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPlatformInfo.Merge(m, src)
}
func (m *UserPlatformInfo) XXX_Size() int {
	return xxx_messageInfo_UserPlatformInfo.Size(m)
}
func (m *UserPlatformInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPlatformInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserPlatformInfo proto.InternalMessageInfo

func (m *UserPlatformInfo) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UserPlatformInfo) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *UserPlatformInfo) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

//获取用户的信息
type UserInfo struct {
	UserBaseInfo         *UserBaseInfo       `protobuf:"bytes,1,opt,name=user_base_info,json=userBaseInfo,proto3" json:"user_base_info,omitempty"`
	UserPlatformInfo     []*UserPlatformInfo `protobuf:"bytes,2,rep,name=user_platform_info,json=userPlatformInfo,proto3" json:"user_platform_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserBaseInfo() *UserBaseInfo {
	if m != nil {
		return m.UserBaseInfo
	}
	return nil
}

func (m *UserInfo) GetUserPlatformInfo() []*UserPlatformInfo {
	if m != nil {
		return m.UserPlatformInfo
	}
	return nil
}

//邮件或手机号请求
type EmailOrPhoneRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailOrPhoneRequest) Reset()         { *m = EmailOrPhoneRequest{} }
func (m *EmailOrPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*EmailOrPhoneRequest) ProtoMessage()    {}
func (*EmailOrPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{3}
}

func (m *EmailOrPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailOrPhoneRequest.Unmarshal(m, b)
}
func (m *EmailOrPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailOrPhoneRequest.Marshal(b, m, deterministic)
}
func (m *EmailOrPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailOrPhoneRequest.Merge(m, src)
}
func (m *EmailOrPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_EmailOrPhoneRequest.Size(m)
}
func (m *EmailOrPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailOrPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailOrPhoneRequest proto.InternalMessageInfo

func (m *EmailOrPhoneRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailOrPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*UserBaseInfo)(nil), "account.UserBaseInfo")
	proto.RegisterType((*UserPlatformInfo)(nil), "account.UserPlatformInfo")
	proto.RegisterType((*UserInfo)(nil), "account.UserInfo")
	proto.RegisterType((*EmailOrPhoneRequest)(nil), "account.EmailOrPhoneRequest")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0x87, 0xdf, 0xb6, 0xf0, 0x22, 0x03, 0x21, 0x64, 0xd5, 0xa4, 0xfe, 0x21, 0x92, 0x9e, 0x38,
	0x61, 0x82, 0x47, 0x4f, 0x90, 0x18, 0xc3, 0x49, 0xb2, 0x89, 0x5e, 0xc9, 0x42, 0xa7, 0xba, 0x09,
	0x74, 0xeb, 0x6e, 0x9b, 0xc8, 0xb7, 0xf0, 0x23, 0x9b, 0x1d, 0xb6, 0xba, 0x9a, 0x9e, 0xba, 0xf3,
	0xb4, 0x9d, 0x67, 0xe6, 0xd7, 0xc2, 0x55, 0xa1, 0x55, 0xa9, 0x6e, 0xc5, 0x76, 0xab, 0xaa, 0xbc,
	0xac, 0xaf, 0x53, 0xa2, 0xac, 0xe3, 0xca, 0xe4, 0x03, 0xfa, 0xcf, 0x06, 0xf5, 0x42, 0x18, 0x5c,
	0xe6, 0x99, 0x62, 0x03, 0x08, 0x65, 0x1a, 0x07, 0xe3, 0x60, 0x12, 0xf1, 0x50, 0xa6, 0x2c, 0x86,
	0xfa, 0xd1, 0x38, 0x1c, 0x07, 0x93, 0x2e, 0xaf, 0x4b, 0x76, 0x06, 0x6d, 0xdc, 0x0b, 0xb9, 0x8b,
	0x23, 0xe2, 0xc7, 0xc2, 0xd2, 0xe2, 0x4d, 0xe5, 0x18, 0xb7, 0x8e, 0x94, 0x0a, 0xc6, 0xa0, 0x95,
	0x8b, 0x3d, 0xc6, 0x6d, 0x82, 0x74, 0x4e, 0x32, 0x18, 0x5a, 0xf3, 0x6a, 0x27, 0xca, 0x4c, 0xe9,
	0x3d, 0xd9, 0x19, 0xb4, 0xca, 0x43, 0x81, 0xce, 0x4f, 0x67, 0x36, 0x02, 0x70, 0xca, 0xb5, 0x4c,
	0x69, 0x88, 0x88, 0x77, 0x1d, 0x59, 0xa6, 0xec, 0x06, 0x7a, 0x85, 0x6b, 0x61, 0xef, 0x1f, 0x87,
	0x81, 0x1a, 0x2d, 0xd3, 0xe4, 0x33, 0x80, 0x13, 0x2b, 0x22, 0xc1, 0x3d, 0x0c, 0x2a, 0x83, 0x7a,
	0xbd, 0x11, 0x06, 0xd7, 0x32, 0xcf, 0x14, 0xa9, 0x7a, 0xb3, 0xf3, 0x69, 0x9d, 0x8f, 0x9f, 0x06,
	0xef, 0x57, 0x7e, 0x36, 0x8f, 0xc0, 0xe8, 0xe5, 0x1f, 0x9f, 0x6d, 0x10, 0x8e, 0xa3, 0x49, 0x6f,
	0x76, 0xf1, 0xab, 0x81, 0xbf, 0x14, 0x1f, 0x56, 0x7f, 0x48, 0x32, 0x87, 0xd3, 0x07, 0x9b, 0xd6,
	0x93, 0x5e, 0xd9, 0x78, 0x38, 0xbe, 0x57, 0x68, 0xbc, 0x44, 0x83, 0xc6, 0x44, 0x43, 0x2f, 0xd1,
	0x99, 0x80, 0xce, 0xdc, 0x7d, 0x88, 0x17, 0x18, 0x71, 0x7c, 0x95, 0xa6, 0x44, 0xed, 0xd0, 0xe2,
	0xe0, 0xb7, 0x67, 0xd7, 0xdf, 0xb3, 0x35, 0x58, 0x2f, 0x9b, 0x57, 0x4f, 0xfe, 0x6d, 0xfe, 0xd3,
	0xaf, 0x72, 0xf7, 0x15, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xac, 0x2e, 0xab, 0x49, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	//通过email或者手机号注册一个用户
	RegisterAccountByEmailOrPhone(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserBaseInfo, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) RegisterAccountByEmailOrPhone(ctx context.Context, in *EmailOrPhoneRequest, opts ...grpc.CallOption) (*UserBaseInfo, error) {
	out := new(UserBaseInfo)
	err := c.cc.Invoke(ctx, "/account.Account/RegisterAccountByEmailOrPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	//通过email或者手机号注册一个用户
	RegisterAccountByEmailOrPhone(context.Context, *EmailOrPhoneRequest) (*UserBaseInfo, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) RegisterAccountByEmailOrPhone(ctx context.Context, req *EmailOrPhoneRequest) (*UserBaseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccountByEmailOrPhone not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_RegisterAccountByEmailOrPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailOrPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RegisterAccountByEmailOrPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/RegisterAccountByEmailOrPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RegisterAccountByEmailOrPhone(ctx, req.(*EmailOrPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccountByEmailOrPhone",
			Handler:    _Account_RegisterAccountByEmailOrPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/account.proto",
}
