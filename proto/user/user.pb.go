// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Origin               string   `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Valid                bool     `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x6d, 0x9a, 0xa4, 0xe9, 0x08, 0x3d, 0x0c, 0x2a, 0x4b, 0x41, 0x29, 0x11, 0xd4, 0x53,
	0x85, 0xfa, 0x04, 0x45, 0x44, 0xbc, 0x06, 0x7a, 0x96, 0x8d, 0x3b, 0xc8, 0x42, 0x92, 0x8d, 0xbb,
	0x1b, 0xf5, 0xe2, 0x0b, 0xfa, 0x54, 0xb2, 0xb3, 0xa9, 0x88, 0x97, 0x5e, 0x92, 0xf9, 0xff, 0x6f,
	0x76, 0x98, 0x7f, 0x17, 0x4e, 0x7b, 0x6b, 0xbc, 0xb9, 0x1d, 0x1c, 0x59, 0xfe, 0xac, 0x59, 0x63,
	0x1a, 0xea, 0xf2, 0x7b, 0x02, 0xe9, 0xce, 0x91, 0xc5, 0x05, 0x24, 0x5a, 0x89, 0xc9, 0x6a, 0x72,
	0x33, 0xaf, 0x12, 0xad, 0x10, 0x21, 0xed, 0x64, 0x4b, 0x22, 0x61, 0x87, 0x6b, 0x3c, 0x83, 0xbc,
	0x35, 0xb5, 0x6e, 0x48, 0x4c, 0xd9, 0x1d, 0x15, 0x9e, 0x40, 0x46, 0xad, 0xd4, 0x8d, 0x48, 0xd9,
	0x8e, 0x02, 0x97, 0x50, 0xf4, 0xd2, 0xb9, 0x0f, 0x63, 0x95, 0xc8, 0x18, 0xfc, 0xea, 0x30, 0xc9,
	0x58, 0xfd, 0xaa, 0x3b, 0x91, 0xc7, 0x49, 0x51, 0xe1, 0x39, 0xc0, 0x8b, 0x25, 0xe9, 0x49, 0x3d,
	0x4b, 0x2f, 0x66, 0xcc, 0xe6, 0xa3, 0xb3, 0xf5, 0x01, 0x0f, 0xbd, 0xda, 0xe3, 0x22, 0xe2, 0xd1,
	0xd9, 0xfa, 0x72, 0x0e, 0xb3, 0x8a, 0xde, 0x06, 0x72, 0xbe, 0xac, 0xa1, 0xa8, 0xc8, 0xf5, 0xa6,
	0x73, 0x84, 0x17, 0xc0, 0x59, 0x39, 0xdc, 0xf1, 0x06, 0xd6, 0x7c, 0x09, 0x21, 0x74, 0xc5, 0x3e,
	0xae, 0x20, 0x0b, 0x7f, 0x27, 0x92, 0xd5, 0xf4, 0x5f, 0x43, 0x04, 0x21, 0xe0, 0xbb, 0x6c, 0xb4,
	0xe2, 0xdc, 0x45, 0x15, 0xc5, 0xe6, 0x0b, 0xb2, 0x1d, 0xe3, 0x2b, 0xc8, 0xef, 0x79, 0x47, 0xfc,
	0x73, 0x76, 0xb9, 0x88, 0xf5, 0x7e, 0x8d, 0xf2, 0x08, 0xaf, 0x61, 0xf6, 0xe4, 0x1e, 0x3e, 0xb5,
	0xf3, 0x07, 0x1a, 0x2f, 0x61, 0xfa, 0x48, 0x07, 0x9a, 0xea, 0x9c, 0xdf, 0xf1, 0xee, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x73, 0xf3, 0x1e, 0x9f, 0xe0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersClient interface {
	// 创建用户
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 用过 用户名、邮箱、手机 查询用户是否存在
	IsExist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取用户
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type usersClient struct {
	c           client.Client
	serviceName string
}

func NewUsersClient(serviceName string, c client.Client) UsersClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "user"
	}
	return &usersClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *usersClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) IsExist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.IsExist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	// 创建用户
	Create(context.Context, *User, *Response) error
	// 用过 用户名、邮箱、手机 查询用户是否存在
	IsExist(context.Context, *User, *Response) error
	// 根据 唯一 获取用户
	Get(context.Context, *User, *Response) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Users{hdlr}, opts...))
}

type Users struct {
	UsersHandler
}

func (h *Users) Create(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Create(ctx, in, out)
}

func (h *Users) IsExist(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.IsExist(ctx, in, out)
}

func (h *Users) Get(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Get(ctx, in, out)
}
