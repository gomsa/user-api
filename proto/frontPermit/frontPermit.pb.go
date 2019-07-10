// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/frontPermit/frontPermit.proto

package frontPermit

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

type Permission struct {
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_e01d8dcd8aa06d26, []int{0}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Permission) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type FrontPermit struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	App                  string        `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	Service              string        `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Method               string        `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Name                 string        `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Permissions          []*Permission `protobuf:"bytes,7,rep,name=permissions,proto3" json:"permissions,omitempty"`
	CreatedAt            string        `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FrontPermit) Reset()         { *m = FrontPermit{} }
func (m *FrontPermit) String() string { return proto.CompactTextString(m) }
func (*FrontPermit) ProtoMessage()    {}
func (*FrontPermit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e01d8dcd8aa06d26, []int{1}
}

func (m *FrontPermit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontPermit.Unmarshal(m, b)
}
func (m *FrontPermit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontPermit.Marshal(b, m, deterministic)
}
func (m *FrontPermit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontPermit.Merge(m, src)
}
func (m *FrontPermit) XXX_Size() int {
	return xxx_messageInfo_FrontPermit.Size(m)
}
func (m *FrontPermit) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontPermit.DiscardUnknown(m)
}

var xxx_messageInfo_FrontPermit proto.InternalMessageInfo

func (m *FrontPermit) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FrontPermit) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *FrontPermit) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *FrontPermit) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *FrontPermit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrontPermit) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FrontPermit) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *FrontPermit) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FrontPermit) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Service              string   `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_e01d8dcd8aa06d26, []int{2}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ListQuery) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ListQuery) GetName() string {
	if m != nil {
		return m.Name
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
	return fileDescriptor_e01d8dcd8aa06d26, []int{3}
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
	FrontPermit          *FrontPermit   `protobuf:"bytes,1,opt,name=frontPermit,proto3" json:"frontPermit,omitempty"`
	FrontPermits         []*FrontPermit `protobuf:"bytes,2,rep,name=frontPermits,proto3" json:"frontPermits,omitempty"`
	Total                int64          `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Valid                bool           `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e01d8dcd8aa06d26, []int{4}
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

func (m *Response) GetFrontPermit() *FrontPermit {
	if m != nil {
		return m.FrontPermit
	}
	return nil
}

func (m *Response) GetFrontPermits() []*FrontPermit {
	if m != nil {
		return m.FrontPermits
	}
	return nil
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Permission)(nil), "frontPermit.Permission")
	proto.RegisterType((*FrontPermit)(nil), "frontPermit.FrontPermit")
	proto.RegisterType((*ListQuery)(nil), "frontPermit.ListQuery")
	proto.RegisterType((*Request)(nil), "frontPermit.Request")
	proto.RegisterType((*Response)(nil), "frontPermit.Response")
}

func init() {
	proto.RegisterFile("proto/frontPermit/frontPermit.proto", fileDescriptor_e01d8dcd8aa06d26)
}

var fileDescriptor_e01d8dcd8aa06d26 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0xd9, 0x36, 0x6d, 0x26, 0xc7, 0x21, 0xc3, 0x79, 0x2e, 0x82, 0x50, 0xe2, 0x4b, 0x9f,
	0x4e, 0xa8, 0x82, 0x3f, 0x11, 0xca, 0xc9, 0xf9, 0x22, 0xa8, 0x0b, 0x3e, 0x4b, 0x6c, 0x46, 0x5d,
	0x48, 0xb3, 0x71, 0x77, 0x7a, 0xe0, 0x5f, 0xa1, 0x7f, 0x8a, 0xe0, 0x3f, 0x28, 0xd9, 0xe4, 0xda,
	0xad, 0x98, 0x97, 0xbc, 0xcd, 0x7c, 0x33, 0xdf, 0x4e, 0xbe, 0x6f, 0x86, 0xc0, 0x83, 0xc6, 0x1a,
	0x36, 0x0f, 0xbf, 0x58, 0x53, 0xf3, 0x7b, 0xb2, 0x5b, 0xcd, 0x61, 0x7c, 0xe1, 0xab, 0x98, 0x05,
	0x50, 0xfe, 0x0a, 0xc0, 0x47, 0xce, 0x69, 0x53, 0xa3, 0x84, 0x99, 0x23, 0x7b, 0xad, 0x37, 0x24,
	0xe3, 0x45, 0xb4, 0x4c, 0xd5, 0x4d, 0x8a, 0xe7, 0x90, 0x6c, 0x89, 0xbf, 0x99, 0x52, 0x0a, 0x5f,
	0xe8, 0xb3, 0xfc, 0x57, 0x0c, 0xd9, 0xd5, 0xe1, 0x3d, 0x3c, 0x85, 0x58, 0x97, 0x32, 0x5a, 0x44,
	0x4b, 0xa1, 0x62, 0x5d, 0xe2, 0x6d, 0x10, 0x45, 0xd3, 0xf4, 0xaf, 0xb5, 0x61, 0x38, 0x43, 0x0c,
	0xcd, 0x98, 0x84, 0x33, 0x10, 0x61, 0x52, 0x17, 0x5b, 0x92, 0x53, 0x8f, 0xfa, 0x18, 0x17, 0x90,
	0x95, 0xe4, 0x36, 0x56, 0x37, 0xac, 0x4d, 0x2d, 0x13, 0x5f, 0x0a, 0x21, 0x7c, 0x06, 0x59, 0xb3,
	0x57, 0xe6, 0xe4, 0x6c, 0x21, 0x96, 0xd9, 0xea, 0xee, 0x45, 0xe8, 0xc7, 0x41, 0xb9, 0x0a, 0x7b,
	0xf1, 0x3e, 0xc0, 0xc6, 0x52, 0xc1, 0x54, 0x7e, 0x2a, 0x58, 0xce, 0xfd, 0xdb, 0x69, 0x8f, 0xac,
	0xb9, 0x2d, 0xef, 0x9a, 0xf2, 0xa6, 0x9c, 0x76, 0xe5, 0x1e, 0x59, 0x73, 0xfe, 0x33, 0x82, 0xf4,
	0xad, 0x76, 0xfc, 0x61, 0x47, 0xf6, 0x07, 0x9e, 0xc1, 0xb4, 0xd2, 0x5b, 0xcd, 0xbd, 0x27, 0x5d,
	0xd2, 0x4a, 0x6a, 0x8a, 0xaf, 0x9d, 0xcb, 0x42, 0xf9, 0xb8, 0xc5, 0x9c, 0xb1, 0xdc, 0xbb, 0xe2,
	0xe3, 0xd0, 0xac, 0xc9, 0x90, 0x59, 0xd3, 0xff, 0x9a, 0x95, 0x1c, 0xcc, 0xca, 0x53, 0x98, 0x29,
	0xfa, 0xbe, 0x23, 0xc7, 0xf9, 0xef, 0x08, 0xe6, 0x8a, 0x5c, 0x63, 0x6a, 0x47, 0xf8, 0x1c, 0xc2,
	0x5b, 0xf0, 0x5f, 0x98, 0xad, 0xe4, 0x91, 0x45, 0xc1, 0x6e, 0x55, 0xd8, 0x8c, 0x2f, 0xe1, 0x24,
	0x48, 0x9d, 0x8c, 0xbd, 0xbf, 0xc3, 0xe4, 0xa3, 0xee, 0xd6, 0x15, 0x36, 0x5c, 0x54, 0x5e, 0xac,
	0x50, 0x5d, 0xd2, 0xa2, 0xd7, 0x45, 0xa5, 0xbb, 0xfd, 0xcf, 0x55, 0x97, 0xac, 0xfe, 0x08, 0x38,
	0xb9, 0x0a, 0xc9, 0x8f, 0x41, 0xac, 0xab, 0x0a, 0xcf, 0x8e, 0x66, 0xf5, 0x02, 0xef, 0xdd, 0xf9,
	0x07, 0xed, 0xa4, 0xe6, 0xb7, 0xf0, 0x09, 0x4c, 0xda, 0xad, 0xe0, 0xf9, 0x51, 0xc3, 0x7e, 0x51,
	0xc3, 0xc4, 0xa7, 0x20, 0xde, 0x10, 0xe3, 0xa0, 0xb4, 0x61, 0xe6, 0x0b, 0x48, 0x2e, 0xfd, 0xd5,
	0x8c, 0x24, 0x7f, 0xf4, 0x37, 0x35, 0x92, 0xfc, 0x9a, 0x2a, 0x1a, 0x47, 0xbe, 0x84, 0xd3, 0x6e,
	0xf2, 0x3b, 0x3b, 0xfa, 0xf3, 0x3f, 0x27, 0xfe, 0x67, 0xf3, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbd, 0xc0, 0x44, 0xfc, 0x93, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FrontPermits service

type FrontPermitsClient interface {
	// 权限验证授权
	// 全部权限
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取权限列表
	List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取权限
	Get(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error)
	// 创建权限
	Create(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error)
	// 更新权限
	Update(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error)
	// 删除权限
	Delete(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error)
	// 同步前端权限 允许外部权限  需要最高 root 权限  *********高危 调用慎重*********
	UpdateOrCreate(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error)
}

type frontPermitsClient struct {
	c           client.Client
	serviceName string
}

func NewFrontPermitsClient(serviceName string, c client.Client) FrontPermitsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "frontPermit"
	}
	return &frontPermitsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *frontPermitsClient) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) List(ctx context.Context, in *ListQuery, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) Get(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) Create(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) Update(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) Delete(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontPermitsClient) UpdateOrCreate(ctx context.Context, in *FrontPermit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FrontPermits.UpdateOrCreate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FrontPermits service

type FrontPermitsHandler interface {
	// 权限验证授权
	// 全部权限
	All(context.Context, *Request, *Response) error
	// 获取权限列表
	List(context.Context, *ListQuery, *Response) error
	// 根据 唯一 获取权限
	Get(context.Context, *FrontPermit, *Response) error
	// 创建权限
	Create(context.Context, *FrontPermit, *Response) error
	// 更新权限
	Update(context.Context, *FrontPermit, *Response) error
	// 删除权限
	Delete(context.Context, *FrontPermit, *Response) error
	// 同步前端权限 允许外部权限  需要最高 root 权限  *********高危 调用慎重*********
	UpdateOrCreate(context.Context, *FrontPermit, *Response) error
}

func RegisterFrontPermitsHandler(s server.Server, hdlr FrontPermitsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&FrontPermits{hdlr}, opts...))
}

type FrontPermits struct {
	FrontPermitsHandler
}

func (h *FrontPermits) All(ctx context.Context, in *Request, out *Response) error {
	return h.FrontPermitsHandler.All(ctx, in, out)
}

func (h *FrontPermits) List(ctx context.Context, in *ListQuery, out *Response) error {
	return h.FrontPermitsHandler.List(ctx, in, out)
}

func (h *FrontPermits) Get(ctx context.Context, in *FrontPermit, out *Response) error {
	return h.FrontPermitsHandler.Get(ctx, in, out)
}

func (h *FrontPermits) Create(ctx context.Context, in *FrontPermit, out *Response) error {
	return h.FrontPermitsHandler.Create(ctx, in, out)
}

func (h *FrontPermits) Update(ctx context.Context, in *FrontPermit, out *Response) error {
	return h.FrontPermitsHandler.Update(ctx, in, out)
}

func (h *FrontPermits) Delete(ctx context.Context, in *FrontPermit, out *Response) error {
	return h.FrontPermitsHandler.Delete(ctx, in, out)
}

func (h *FrontPermits) UpdateOrCreate(ctx context.Context, in *FrontPermit, out *Response) error {
	return h.FrontPermitsHandler.UpdateOrCreate(ctx, in, out)
}
