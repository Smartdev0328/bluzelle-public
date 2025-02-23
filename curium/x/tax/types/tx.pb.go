// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tax/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgSetGasTaxBp struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Bp      int64  `protobuf:"varint,2,opt,name=bp,proto3" json:"bp,omitempty"`
}

func (m *MsgSetGasTaxBp) Reset()         { *m = MsgSetGasTaxBp{} }
func (m *MsgSetGasTaxBp) String() string { return proto.CompactTextString(m) }
func (*MsgSetGasTaxBp) ProtoMessage()    {}
func (*MsgSetGasTaxBp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{0}
}
func (m *MsgSetGasTaxBp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetGasTaxBp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetGasTaxBp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetGasTaxBp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetGasTaxBp.Merge(m, src)
}
func (m *MsgSetGasTaxBp) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetGasTaxBp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetGasTaxBp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetGasTaxBp proto.InternalMessageInfo

func (m *MsgSetGasTaxBp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetGasTaxBp) GetBp() int64 {
	if m != nil {
		return m.Bp
	}
	return 0
}

type MsgSetGasTaxBpResponse struct {
}

func (m *MsgSetGasTaxBpResponse) Reset()         { *m = MsgSetGasTaxBpResponse{} }
func (m *MsgSetGasTaxBpResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetGasTaxBpResponse) ProtoMessage()    {}
func (*MsgSetGasTaxBpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{1}
}
func (m *MsgSetGasTaxBpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetGasTaxBpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetGasTaxBpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetGasTaxBpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetGasTaxBpResponse.Merge(m, src)
}
func (m *MsgSetGasTaxBpResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetGasTaxBpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetGasTaxBpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetGasTaxBpResponse proto.InternalMessageInfo

type MsgSetTransferTaxBp struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Bp      int64  `protobuf:"varint,2,opt,name=bp,proto3" json:"bp,omitempty"`
}

func (m *MsgSetTransferTaxBp) Reset()         { *m = MsgSetTransferTaxBp{} }
func (m *MsgSetTransferTaxBp) String() string { return proto.CompactTextString(m) }
func (*MsgSetTransferTaxBp) ProtoMessage()    {}
func (*MsgSetTransferTaxBp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{2}
}
func (m *MsgSetTransferTaxBp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetTransferTaxBp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetTransferTaxBp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetTransferTaxBp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetTransferTaxBp.Merge(m, src)
}
func (m *MsgSetTransferTaxBp) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetTransferTaxBp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetTransferTaxBp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetTransferTaxBp proto.InternalMessageInfo

func (m *MsgSetTransferTaxBp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetTransferTaxBp) GetBp() int64 {
	if m != nil {
		return m.Bp
	}
	return 0
}

type MsgSetTransferTaxBpResponse struct {
}

func (m *MsgSetTransferTaxBpResponse) Reset()         { *m = MsgSetTransferTaxBpResponse{} }
func (m *MsgSetTransferTaxBpResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetTransferTaxBpResponse) ProtoMessage()    {}
func (*MsgSetTransferTaxBpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{3}
}
func (m *MsgSetTransferTaxBpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetTransferTaxBpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetTransferTaxBpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetTransferTaxBpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetTransferTaxBpResponse.Merge(m, src)
}
func (m *MsgSetTransferTaxBpResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetTransferTaxBpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetTransferTaxBpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetTransferTaxBpResponse proto.InternalMessageInfo

type MsgSetTaxCollector struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TaxCollector string `protobuf:"bytes,2,opt,name=taxCollector,proto3" json:"taxCollector,omitempty"`
}

func (m *MsgSetTaxCollector) Reset()         { *m = MsgSetTaxCollector{} }
func (m *MsgSetTaxCollector) String() string { return proto.CompactTextString(m) }
func (*MsgSetTaxCollector) ProtoMessage()    {}
func (*MsgSetTaxCollector) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{4}
}
func (m *MsgSetTaxCollector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetTaxCollector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetTaxCollector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetTaxCollector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetTaxCollector.Merge(m, src)
}
func (m *MsgSetTaxCollector) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetTaxCollector) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetTaxCollector.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetTaxCollector proto.InternalMessageInfo

func (m *MsgSetTaxCollector) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetTaxCollector) GetTaxCollector() string {
	if m != nil {
		return m.TaxCollector
	}
	return ""
}

type MsgSetTaxCollectorResponse struct {
}

func (m *MsgSetTaxCollectorResponse) Reset()         { *m = MsgSetTaxCollectorResponse{} }
func (m *MsgSetTaxCollectorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetTaxCollectorResponse) ProtoMessage()    {}
func (*MsgSetTaxCollectorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce26c37199483c6a, []int{5}
}
func (m *MsgSetTaxCollectorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetTaxCollectorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetTaxCollectorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetTaxCollectorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetTaxCollectorResponse.Merge(m, src)
}
func (m *MsgSetTaxCollectorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetTaxCollectorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetTaxCollectorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetTaxCollectorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetGasTaxBp)(nil), "bluzelle.curium.tax.MsgSetGasTaxBp")
	proto.RegisterType((*MsgSetGasTaxBpResponse)(nil), "bluzelle.curium.tax.MsgSetGasTaxBpResponse")
	proto.RegisterType((*MsgSetTransferTaxBp)(nil), "bluzelle.curium.tax.MsgSetTransferTaxBp")
	proto.RegisterType((*MsgSetTransferTaxBpResponse)(nil), "bluzelle.curium.tax.MsgSetTransferTaxBpResponse")
	proto.RegisterType((*MsgSetTaxCollector)(nil), "bluzelle.curium.tax.MsgSetTaxCollector")
	proto.RegisterType((*MsgSetTaxCollectorResponse)(nil), "bluzelle.curium.tax.MsgSetTaxCollectorResponse")
}

func init() { proto.RegisterFile("tax/tx.proto", fileDescriptor_ce26c37199483c6a) }

var fileDescriptor_ce26c37199483c6a = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x49, 0xac, 0xd0,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xca, 0x29, 0xad, 0x4a, 0xcd,
	0xc9, 0x49, 0xd5, 0x4b, 0x2e, 0x2d, 0xca, 0x2c, 0xcd, 0xd5, 0x2b, 0x49, 0xac, 0x50, 0xb2, 0xe2,
	0xe2, 0xf3, 0x2d, 0x4e, 0x0f, 0x4e, 0x2d, 0x71, 0x4f, 0x2c, 0x0e, 0x49, 0xac, 0x70, 0x2a, 0x10,
	0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x92, 0x0a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83,
	0x98, 0x92, 0x0a, 0x94, 0x24, 0xb8, 0xc4, 0x50, 0xf5, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15,
	0xa7, 0x2a, 0xd9, 0x73, 0x09, 0x43, 0x64, 0x42, 0x8a, 0x12, 0xf3, 0x8a, 0xd3, 0x52, 0x8b, 0x48,
	0x35, 0x5a, 0x96, 0x4b, 0x1a, 0x8b, 0x01, 0x70, 0xf3, 0x83, 0xb8, 0x84, 0xa0, 0xd2, 0x89, 0x15,
	0xce, 0xf9, 0x39, 0x39, 0xa9, 0xc9, 0x20, 0x43, 0x70, 0x1b, 0xaf, 0x04, 0x0e, 0x0a, 0xb8, 0x4a,
	0xb0, 0x45, 0x9c, 0x41, 0x28, 0x62, 0x4a, 0x32, 0x5c, 0x52, 0x98, 0x66, 0xc2, 0x6c, 0x34, 0x3a,
	0xcc, 0xc4, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x14, 0xcf, 0xc5, 0x8d, 0x1c, 0x58, 0xca, 0x7a, 0x58,
	0x02, 0x55, 0x0f, 0x35, 0x54, 0xa4, 0xb4, 0x89, 0x50, 0x04, 0xb3, 0x48, 0x28, 0x8f, 0x4b, 0x00,
	0x23, 0xdc, 0x34, 0xf0, 0x18, 0x80, 0xa2, 0x52, 0xca, 0x80, 0x58, 0x95, 0x70, 0xfb, 0xb2, 0xb9,
	0xf8, 0xd1, 0xc3, 0x51, 0x1d, 0x9f, 0x21, 0x48, 0x0a, 0xa5, 0xf4, 0x89, 0x54, 0x08, 0xb3, 0xcc,
	0xc9, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61, 0x86, 0xea, 0x43, 0x0c, 0xd5, 0xaf, 0xd0,
	0x07, 0xa7, 0xe3, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x5a, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x8b, 0x53, 0x9a, 0xdb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SetGasTaxBp(ctx context.Context, in *MsgSetGasTaxBp, opts ...grpc.CallOption) (*MsgSetGasTaxBpResponse, error)
	SetTransferTaxBp(ctx context.Context, in *MsgSetTransferTaxBp, opts ...grpc.CallOption) (*MsgSetTransferTaxBpResponse, error)
	SetTaxCollector(ctx context.Context, in *MsgSetTaxCollector, opts ...grpc.CallOption) (*MsgSetTaxCollectorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetGasTaxBp(ctx context.Context, in *MsgSetGasTaxBp, opts ...grpc.CallOption) (*MsgSetGasTaxBpResponse, error) {
	out := new(MsgSetGasTaxBpResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.tax.Msg/SetGasTaxBp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetTransferTaxBp(ctx context.Context, in *MsgSetTransferTaxBp, opts ...grpc.CallOption) (*MsgSetTransferTaxBpResponse, error) {
	out := new(MsgSetTransferTaxBpResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.tax.Msg/SetTransferTaxBp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetTaxCollector(ctx context.Context, in *MsgSetTaxCollector, opts ...grpc.CallOption) (*MsgSetTaxCollectorResponse, error) {
	out := new(MsgSetTaxCollectorResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.tax.Msg/SetTaxCollector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetGasTaxBp(context.Context, *MsgSetGasTaxBp) (*MsgSetGasTaxBpResponse, error)
	SetTransferTaxBp(context.Context, *MsgSetTransferTaxBp) (*MsgSetTransferTaxBpResponse, error)
	SetTaxCollector(context.Context, *MsgSetTaxCollector) (*MsgSetTaxCollectorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetGasTaxBp(ctx context.Context, req *MsgSetGasTaxBp) (*MsgSetGasTaxBpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGasTaxBp not implemented")
}
func (*UnimplementedMsgServer) SetTransferTaxBp(ctx context.Context, req *MsgSetTransferTaxBp) (*MsgSetTransferTaxBpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTransferTaxBp not implemented")
}
func (*UnimplementedMsgServer) SetTaxCollector(ctx context.Context, req *MsgSetTaxCollector) (*MsgSetTaxCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTaxCollector not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetGasTaxBp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetGasTaxBp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetGasTaxBp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.tax.Msg/SetGasTaxBp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetGasTaxBp(ctx, req.(*MsgSetGasTaxBp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetTransferTaxBp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetTransferTaxBp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetTransferTaxBp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.tax.Msg/SetTransferTaxBp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetTransferTaxBp(ctx, req.(*MsgSetTransferTaxBp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetTaxCollector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetTaxCollector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetTaxCollector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.tax.Msg/SetTaxCollector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetTaxCollector(ctx, req.(*MsgSetTaxCollector))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bluzelle.curium.tax.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGasTaxBp",
			Handler:    _Msg_SetGasTaxBp_Handler,
		},
		{
			MethodName: "SetTransferTaxBp",
			Handler:    _Msg_SetTransferTaxBp_Handler,
		},
		{
			MethodName: "SetTaxCollector",
			Handler:    _Msg_SetTaxCollector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tax/tx.proto",
}

func (m *MsgSetGasTaxBp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetGasTaxBp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetGasTaxBp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Bp))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetGasTaxBpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetGasTaxBpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetGasTaxBpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetTransferTaxBp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetTransferTaxBp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetTransferTaxBp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Bp))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetTransferTaxBpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetTransferTaxBpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetTransferTaxBpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetTaxCollector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetTaxCollector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetTaxCollector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaxCollector) > 0 {
		i -= len(m.TaxCollector)
		copy(dAtA[i:], m.TaxCollector)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TaxCollector)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetTaxCollectorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetTaxCollectorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetTaxCollectorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetGasTaxBp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bp != 0 {
		n += 1 + sovTx(uint64(m.Bp))
	}
	return n
}

func (m *MsgSetGasTaxBpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetTransferTaxBp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bp != 0 {
		n += 1 + sovTx(uint64(m.Bp))
	}
	return n
}

func (m *MsgSetTransferTaxBpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetTaxCollector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TaxCollector)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetTaxCollectorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetGasTaxBp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetGasTaxBp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetGasTaxBp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bp", wireType)
			}
			m.Bp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetGasTaxBpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetGasTaxBpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetGasTaxBpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetTransferTaxBp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetTransferTaxBp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetTransferTaxBp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bp", wireType)
			}
			m.Bp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetTransferTaxBpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetTransferTaxBpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetTransferTaxBpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetTaxCollector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetTaxCollector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetTaxCollector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCollector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxCollector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetTaxCollectorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetTaxCollectorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetTaxCollectorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
