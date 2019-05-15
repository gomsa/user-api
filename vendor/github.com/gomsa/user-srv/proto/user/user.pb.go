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
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	Valid                bool     `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
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

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Error)(nil), "user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xab, 0x4f, 0xcb, 0x63, 0xea, 0xc3, 0xd0, 0x96, 0xc5, 0xd0, 0x22, 0x64, 0x5a, 0xda,
	0x8b, 0x0b, 0xce, 0x39, 0x07, 0x13, 0x8c, 0xef, 0x02, 0x9f, 0x83, 0xec, 0x1d, 0xc2, 0x82, 0xac,
	0x55, 0x76, 0xd7, 0x49, 0xde, 0x20, 0x6f, 0x92, 0x87, 0xc9, 0x53, 0x85, 0x9d, 0x95, 0x83, 0x49,
	0x0e, 0xbe, 0x48, 0xf3, 0xff, 0x60, 0xd9, 0xf9, 0xb1, 0xf0, 0xbd, 0x37, 0xda, 0xe9, 0xff, 0x47,
	0x4b, 0x86, 0x3f, 0x0b, 0xd6, 0x98, 0xfa, 0xb9, 0x7a, 0x8d, 0x20, 0xdd, 0x5a, 0x32, 0x38, 0x85,
	0x58, 0x49, 0x11, 0x95, 0xd1, 0xdf, 0x71, 0x1d, 0x2b, 0x89, 0x08, 0x69, 0xd7, 0x1c, 0x48, 0xc4,
	0xec, 0xf0, 0x8c, 0x3f, 0x20, 0x3f, 0xe8, 0x9d, 0x6a, 0x49, 0x24, 0xec, 0x0e, 0x0a, 0xbf, 0x41,
	0x46, 0x87, 0x46, 0xb5, 0x22, 0x65, 0x3b, 0x08, 0x9c, 0x41, 0xd1, 0x37, 0xd6, 0x3e, 0x6a, 0x23,
	0x45, 0xc6, 0xc1, 0xbb, 0xf6, 0x27, 0x69, 0xa3, 0xee, 0x54, 0x27, 0xf2, 0x70, 0x52, 0x50, 0xf8,
	0x13, 0x60, 0x6f, 0xa8, 0x71, 0x24, 0x6f, 0x1b, 0x27, 0x46, 0x9c, 0x8d, 0x07, 0x67, 0xe5, 0x7c,
	0x7c, 0xec, 0xe5, 0x29, 0x2e, 0x42, 0x3c, 0x38, 0x2b, 0x57, 0x8d, 0x61, 0x54, 0xd3, 0xfd, 0x91,
	0xac, 0xab, 0x9e, 0x23, 0x28, 0x6a, 0xb2, 0xbd, 0xee, 0x2c, 0xe1, 0x2f, 0xe0, 0x65, 0x79, 0xbb,
	0xc9, 0x12, 0x16, 0x4c, 0xc1, 0x6f, 0x5d, 0xb3, 0x8f, 0x25, 0x64, 0xfe, 0x6f, 0x45, 0x5c, 0x26,
	0x1f, 0x0a, 0x21, 0xc0, 0x39, 0xe4, 0x64, 0x8c, 0x36, 0x56, 0x24, 0x5c, 0x99, 0x84, 0xca, 0xda,
	0x7b, 0xf5, 0x10, 0x79, 0x0c, 0x0f, 0x4d, 0xab, 0x24, 0x63, 0x28, 0xea, 0x20, 0xaa, 0x6b, 0xc8,
	0xb8, 0xe6, 0x89, 0xee, 0xb5, 0xa4, 0x81, 0x31, 0xcf, 0x58, 0xc2, 0x44, 0x92, 0xdd, 0x1b, 0xd5,
	0x3b, 0xa5, 0xbb, 0x01, 0xf6, 0xb9, 0xb5, 0x7c, 0x89, 0x20, 0xdb, 0xf2, 0x1d, 0xfe, 0x40, 0x7e,
	0xc3, 0x24, 0xf0, 0xec, 0x82, 0xb3, 0x69, 0x98, 0x4f, 0xbb, 0x56, 0x5f, 0xf0, 0x37, 0x64, 0xeb,
	0x27, 0x65, 0xdd, 0x85, 0xda, 0x1c, 0x92, 0x0d, 0x5d, 0x2a, 0xfd, 0x83, 0x7c, 0x43, 0x6e, 0xd5,
	0xb6, 0xf8, 0xf5, 0x94, 0x31, 0xdf, 0xcf, 0xd5, 0x5d, 0xce, 0xcf, 0xea, 0xea, 0x2d, 0x00, 0x00,
	0xff, 0xff, 0x11, 0x81, 0x19, 0xc5, 0x6f, 0x02, 0x00, 0x00,
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
	Exist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取用户
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	// 获取全部用户
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
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

func (c *usersClient) Exist(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.Exist", in)
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

func (c *usersClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Users.GetAll", in)
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
	Exist(context.Context, *User, *Response) error
	// 根据 唯一 获取用户
	Get(context.Context, *User, *Response) error
	// 获取全部用户
	GetAll(context.Context, *Request, *Response) error
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

func (h *Users) Exist(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Exist(ctx, in, out)
}

func (h *Users) Get(ctx context.Context, in *User, out *Response) error {
	return h.UsersHandler.Get(ctx, in, out)
}

func (h *Users) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.GetAll(ctx, in, out)
}
