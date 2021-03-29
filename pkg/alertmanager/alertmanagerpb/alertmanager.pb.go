// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alertmanager.proto

package alertmanagerpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	clusterpb "github.com/prometheus/alertmanager/cluster/clusterpb"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UpdateStateStatus int32

const (
	OK             UpdateStateStatus = 0
	MERGE_ERROR    UpdateStateStatus = 2
	USER_NOT_FOUND UpdateStateStatus = 3
)

var UpdateStateStatus_name = map[int32]string{
	0: "OK",
	2: "MERGE_ERROR",
	3: "USER_NOT_FOUND",
}

var UpdateStateStatus_value = map[string]int32{
	"OK":             0,
	"MERGE_ERROR":    2,
	"USER_NOT_FOUND": 3,
}

func (UpdateStateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}

type ReadStateStatus int32

const (
	READ_UNSPECIFIED    ReadStateStatus = 0
	READ_OK             ReadStateStatus = 1
	READ_ERROR          ReadStateStatus = 2
	READ_USER_NOT_FOUND ReadStateStatus = 3
)

var ReadStateStatus_name = map[int32]string{
	0: "READ_UNSPECIFIED",
	1: "READ_OK",
	2: "READ_ERROR",
	3: "READ_USER_NOT_FOUND",
}

var ReadStateStatus_value = map[string]int32{
	"READ_UNSPECIFIED":    0,
	"READ_OK":             1,
	"READ_ERROR":          2,
	"READ_USER_NOT_FOUND": 3,
}

func (ReadStateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{1}
}

type UpdateStateResponse struct {
	Status UpdateStateStatus `protobuf:"varint,1,opt,name=status,proto3,enum=alertmanagerpb.UpdateStateStatus" json:"status,omitempty"`
	Error  string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *UpdateStateResponse) Reset()      { *m = UpdateStateResponse{} }
func (*UpdateStateResponse) ProtoMessage() {}
func (*UpdateStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{0}
}
func (m *UpdateStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStateResponse.Merge(m, src)
}
func (m *UpdateStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStateResponse proto.InternalMessageInfo

func (m *UpdateStateResponse) GetStatus() UpdateStateStatus {
	if m != nil {
		return m.Status
	}
	return OK
}

func (m *UpdateStateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReadStateRequest struct {
}

func (m *ReadStateRequest) Reset()      { *m = ReadStateRequest{} }
func (*ReadStateRequest) ProtoMessage() {}
func (*ReadStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{1}
}
func (m *ReadStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStateRequest.Merge(m, src)
}
func (m *ReadStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReadStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStateRequest proto.InternalMessageInfo

type ReadStateResponse struct {
	Status ReadStateStatus      `protobuf:"varint,1,opt,name=status,proto3,enum=alertmanagerpb.ReadStateStatus" json:"status,omitempty"`
	Error  string               `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	State  *clusterpb.FullState `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *ReadStateResponse) Reset()      { *m = ReadStateResponse{} }
func (*ReadStateResponse) ProtoMessage() {}
func (*ReadStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e60437b6e0c74c9a, []int{2}
}
func (m *ReadStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadStateResponse.Merge(m, src)
}
func (m *ReadStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *ReadStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadStateResponse proto.InternalMessageInfo

func (m *ReadStateResponse) GetStatus() ReadStateStatus {
	if m != nil {
		return m.Status
	}
	return READ_UNSPECIFIED
}

func (m *ReadStateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ReadStateResponse) GetState() *clusterpb.FullState {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterEnum("alertmanagerpb.UpdateStateStatus", UpdateStateStatus_name, UpdateStateStatus_value)
	proto.RegisterEnum("alertmanagerpb.ReadStateStatus", ReadStateStatus_name, ReadStateStatus_value)
	proto.RegisterType((*UpdateStateResponse)(nil), "alertmanagerpb.UpdateStateResponse")
	proto.RegisterType((*ReadStateRequest)(nil), "alertmanagerpb.ReadStateRequest")
	proto.RegisterType((*ReadStateResponse)(nil), "alertmanagerpb.ReadStateResponse")
}

func init() { proto.RegisterFile("alertmanager.proto", fileDescriptor_e60437b6e0c74c9a) }

var fileDescriptor_e60437b6e0c74c9a = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0x12, 0x41,
	0x18, 0x9d, 0xa1, 0x16, 0xd3, 0x0f, 0x85, 0xed, 0x14, 0x95, 0x70, 0x98, 0x52, 0xbc, 0x10, 0x0e,
	0xbb, 0x09, 0x9a, 0x18, 0x3d, 0xb5, 0x95, 0xc5, 0x36, 0x8d, 0x40, 0x06, 0xb8, 0x98, 0x18, 0x32,
	0xc0, 0x08, 0x46, 0x60, 0xd6, 0xd9, 0x59, 0x7b, 0xf5, 0x27, 0x78, 0xf0, 0x07, 0x78, 0xf4, 0xa7,
	0x78, 0xe4, 0xd8, 0xa3, 0x2c, 0x97, 0x26, 0x5e, 0xfa, 0x13, 0x4c, 0x59, 0x76, 0x5d, 0xd7, 0xd8,
	0xf4, 0xb4, 0xdf, 0xbc, 0xf9, 0xde, 0x7b, 0xf3, 0xbd, 0x99, 0x05, 0xc2, 0xa7, 0x42, 0xe9, 0x19,
	0x9f, 0xf3, 0xb1, 0x50, 0xa6, 0xa3, 0xa4, 0x96, 0x24, 0x1b, 0xc7, 0x9c, 0x41, 0x31, 0x3f, 0x96,
	0x63, 0xb9, 0xde, 0xb2, 0xae, 0xab, 0xa0, 0xab, 0xf8, 0x74, 0xfc, 0x5e, 0x4f, 0xbc, 0x81, 0x39,
	0x94, 0x33, 0xeb, 0x5c, 0xf0, 0x4f, 0xe2, 0x5c, 0xaa, 0x0f, 0xae, 0x35, 0x94, 0xb3, 0x99, 0x9c,
	0x5b, 0x13, 0xad, 0x9d, 0xb1, 0x72, 0x86, 0x51, 0xb1, 0x61, 0x1d, 0xc7, 0x58, 0x8e, 0x92, 0x33,
	0xa1, 0x27, 0xc2, 0x73, 0xad, 0xb8, 0xa3, 0x35, 0x9c, 0x7a, 0xae, 0xfe, 0xf3, 0x75, 0x06, 0x61,
	0x15, 0x68, 0x94, 0xdf, 0xc1, 0x5e, 0xcf, 0x19, 0x71, 0x2d, 0x3a, 0x9a, 0x6b, 0xc1, 0x84, 0xeb,
	0xc8, 0xb9, 0x2b, 0xc8, 0x73, 0x48, 0xbb, 0x9a, 0x6b, 0xcf, 0x2d, 0xe0, 0x12, 0xae, 0x64, 0x6b,
	0x07, 0xe6, 0xdf, 0x73, 0x98, 0x31, 0x52, 0x67, 0xdd, 0xc8, 0x36, 0x04, 0x92, 0x87, 0x6d, 0xa1,
	0x94, 0x54, 0x85, 0x54, 0x09, 0x57, 0x76, 0x58, 0xb0, 0x28, 0x13, 0x30, 0x98, 0xe0, 0xa3, 0x8d,
	0xcb, 0x47, 0x4f, 0xb8, 0xba, 0xfc, 0x15, 0xc3, 0x6e, 0x0c, 0xdc, 0x58, 0x3f, 0x4b, 0x58, 0xef,
	0x27, 0xad, 0x23, 0xca, 0x6d, 0x8c, 0x49, 0x15, 0xb6, 0xaf, 0xf7, 0x45, 0x61, 0xab, 0x84, 0x2b,
	0x99, 0x5a, 0xde, 0x8c, 0x92, 0x30, 0x1b, 0xde, 0x74, 0x1a, 0x78, 0x07, 0x2d, 0x2f, 0xee, 0x5c,
	0x7e, 0xdb, 0x47, 0xd5, 0x43, 0xd8, 0xfd, 0x67, 0x3a, 0x92, 0x86, 0x54, 0xeb, 0xcc, 0x40, 0x24,
	0x07, 0x99, 0xd7, 0x36, 0x7b, 0x65, 0xf7, 0x6d, 0xc6, 0x5a, 0xcc, 0x48, 0x11, 0x02, 0xd9, 0x5e,
	0xc7, 0x66, 0xfd, 0x66, 0xab, 0xdb, 0x6f, 0xb4, 0x7a, 0xcd, 0xba, 0xb1, 0x55, 0x7d, 0x0b, 0xb9,
	0xc4, 0x21, 0x49, 0x1e, 0x0c, 0x66, 0x1f, 0xd5, 0xfb, 0xbd, 0x66, 0xa7, 0x6d, 0xbf, 0x3c, 0x6d,
	0x9c, 0xda, 0x75, 0x03, 0x91, 0x0c, 0xdc, 0x5d, 0xa3, 0xad, 0x33, 0x03, 0x93, 0x2c, 0xc0, 0x7a,
	0x11, 0x2a, 0x3f, 0x82, 0xbd, 0x80, 0x92, 0x90, 0xaf, 0xfd, 0xc2, 0x70, 0xef, 0x28, 0x96, 0x09,
	0x39, 0x84, 0xfb, 0x27, 0x7c, 0x3e, 0x9a, 0x86, 0xc9, 0x92, 0x07, 0x66, 0xf4, 0x54, 0x4e, 0xba,
	0xdd, 0xf6, 0x06, 0x2e, 0x3e, 0x4c, 0xc2, 0x41, 0xe4, 0x65, 0x44, 0x6c, 0xc8, 0xc4, 0x66, 0x26,
	0xb9, 0x58, 0x4a, 0x6d, 0xae, 0x74, 0xf1, 0xf1, 0x0d, 0xf7, 0x1f, 0x93, 0x61, 0xb0, 0x13, 0x0d,
	0x4e, 0x4a, 0xff, 0xbd, 0xb8, 0xf0, 0x3c, 0x07, 0x37, 0x74, 0x84, 0x9a, 0xc7, 0xf5, 0xc5, 0x92,
	0xa2, 0x8b, 0x25, 0x45, 0x57, 0x4b, 0x8a, 0x3f, 0xfb, 0x14, 0x7f, 0xf7, 0x29, 0xfe, 0xe1, 0x53,
	0xbc, 0xf0, 0x29, 0xfe, 0xe9, 0x53, 0x7c, 0xe9, 0x53, 0x74, 0xe5, 0x53, 0xfc, 0x65, 0x45, 0xd1,
	0x62, 0x45, 0xd1, 0xc5, 0x8a, 0xa2, 0x37, 0x89, 0xff, 0x6e, 0x90, 0x5e, 0x3f, 0xf7, 0x27, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x5b, 0x6b, 0x33, 0xa4, 0x03, 0x00, 0x00,
}

func (x UpdateStateStatus) String() string {
	s, ok := UpdateStateStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ReadStateStatus) String() string {
	s, ok := ReadStateStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *UpdateStateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateStateResponse)
	if !ok {
		that2, ok := that.(UpdateStateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *ReadStateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReadStateRequest)
	if !ok {
		that2, ok := that.(ReadStateRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UpdateStateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&alertmanagerpb.UpdateStateResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReadStateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&alertmanagerpb.ReadStateRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReadStateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&alertmanagerpb.ReadStateResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	if this.State != nil {
		s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAlertmanager(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertmanagerClient is the client API for Alertmanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertmanagerClient interface {
	HandleRequest(ctx context.Context, in *httpgrpc.HTTPRequest, opts ...grpc.CallOption) (*httpgrpc.HTTPResponse, error)
	UpdateState(ctx context.Context, in *clusterpb.Part, opts ...grpc.CallOption) (*UpdateStateResponse, error)
	ReadState(ctx context.Context, in *ReadStateRequest, opts ...grpc.CallOption) (*ReadStateResponse, error)
}

type alertmanagerClient struct {
	cc *grpc.ClientConn
}

func NewAlertmanagerClient(cc *grpc.ClientConn) AlertmanagerClient {
	return &alertmanagerClient{cc}
}

func (c *alertmanagerClient) HandleRequest(ctx context.Context, in *httpgrpc.HTTPRequest, opts ...grpc.CallOption) (*httpgrpc.HTTPResponse, error) {
	out := new(httpgrpc.HTTPResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/HandleRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertmanagerClient) UpdateState(ctx context.Context, in *clusterpb.Part, opts ...grpc.CallOption) (*UpdateStateResponse, error) {
	out := new(UpdateStateResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertmanagerClient) ReadState(ctx context.Context, in *ReadStateRequest, opts ...grpc.CallOption) (*ReadStateResponse, error) {
	out := new(ReadStateResponse)
	err := c.cc.Invoke(ctx, "/alertmanagerpb.Alertmanager/ReadState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertmanagerServer is the server API for Alertmanager service.
type AlertmanagerServer interface {
	HandleRequest(context.Context, *httpgrpc.HTTPRequest) (*httpgrpc.HTTPResponse, error)
	UpdateState(context.Context, *clusterpb.Part) (*UpdateStateResponse, error)
	ReadState(context.Context, *ReadStateRequest) (*ReadStateResponse, error)
}

// UnimplementedAlertmanagerServer can be embedded to have forward compatible implementations.
type UnimplementedAlertmanagerServer struct {
}

func (*UnimplementedAlertmanagerServer) HandleRequest(ctx context.Context, req *httpgrpc.HTTPRequest) (*httpgrpc.HTTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRequest not implemented")
}
func (*UnimplementedAlertmanagerServer) UpdateState(ctx context.Context, req *clusterpb.Part) (*UpdateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (*UnimplementedAlertmanagerServer) ReadState(ctx context.Context, req *ReadStateRequest) (*ReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadState not implemented")
}

func RegisterAlertmanagerServer(s *grpc.Server, srv AlertmanagerServer) {
	s.RegisterService(&_Alertmanager_serviceDesc, srv)
}

func _Alertmanager_HandleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpgrpc.HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).HandleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/HandleRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).HandleRequest(ctx, req.(*httpgrpc.HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alertmanager_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(clusterpb.Part)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).UpdateState(ctx, req.(*clusterpb.Part))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alertmanager_ReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertmanagerServer).ReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertmanagerpb.Alertmanager/ReadState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertmanagerServer).ReadState(ctx, req.(*ReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alertmanager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alertmanagerpb.Alertmanager",
	HandlerType: (*AlertmanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleRequest",
			Handler:    _Alertmanager_HandleRequest_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Alertmanager_UpdateState_Handler,
		},
		{
			MethodName: "ReadState",
			Handler:    _Alertmanager_ReadState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alertmanager.proto",
}

func (m *UpdateStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintAlertmanager(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintAlertmanager(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReadStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ReadStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != nil {
		{
			size, err := m.State.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAlertmanager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintAlertmanager(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintAlertmanager(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlertmanager(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlertmanager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovAlertmanager(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	return n
}

func (m *ReadStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReadStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovAlertmanager(uint64(m.Status))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	if m.State != nil {
		l = m.State.Size()
		n += 1 + l + sovAlertmanager(uint64(l))
	}
	return n
}

func sovAlertmanager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlertmanager(x uint64) (n int) {
	return sovAlertmanager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UpdateStateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateStateResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReadStateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReadStateRequest{`,
		`}`,
	}, "")
	return s
}
func (this *ReadStateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReadStateResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`State:` + strings.Replace(fmt.Sprintf("%v", this.State), "FullState", "clusterpb.FullState", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAlertmanager(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UpdateStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= UpdateStateStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ReadStateStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAlertmanager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.State == nil {
				m.State = &clusterpb.FullState{}
			}
			if err := m.State.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlertmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAlertmanager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAlertmanager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlertmanager
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAlertmanager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAlertmanager
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAlertmanager
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAlertmanager
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAlertmanager(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAlertmanager
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAlertmanager = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlertmanager   = fmt.Errorf("proto: integer overflow")
)