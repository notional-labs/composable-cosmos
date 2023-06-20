// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: centauri/transfermiddleware/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// message QueryEscrowAddressRequest
type QueryEscrowAddressRequest struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
}

func (m *QueryEscrowAddressRequest) Reset()         { *m = QueryEscrowAddressRequest{} }
func (m *QueryEscrowAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowAddressRequest) ProtoMessage()    {}
func (*QueryEscrowAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2bc06232cb63aa, []int{0}
}
func (m *QueryEscrowAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowAddressRequest.Merge(m, src)
}
func (m *QueryEscrowAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowAddressRequest proto.InternalMessageInfo

func (m *QueryEscrowAddressRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

// QueryEscrowAddressResponse
type QueryEscrowAddressResponse struct {
	EscrowAddress string `protobuf:"bytes,1,opt,name=escrow_address,json=escrowAddress,proto3" json:"escrow_address,omitempty" yaml:"escrow_address"`
}

func (m *QueryEscrowAddressResponse) Reset()         { *m = QueryEscrowAddressResponse{} }
func (m *QueryEscrowAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowAddressResponse) ProtoMessage()    {}
func (*QueryEscrowAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2bc06232cb63aa, []int{1}
}
func (m *QueryEscrowAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowAddressResponse.Merge(m, src)
}
func (m *QueryEscrowAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowAddressResponse proto.InternalMessageInfo

func (m *QueryEscrowAddressResponse) GetEscrowAddress() string {
	if m != nil {
		return m.EscrowAddress
	}
	return ""
}

// QueryParaTokenInfoRequest is the request type for the Query/Params RPC method.
type QueryParaTokenInfoRequest struct {
	NativeDenom string `protobuf:"bytes,1,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
}

func (m *QueryParaTokenInfoRequest) Reset()         { *m = QueryParaTokenInfoRequest{} }
func (m *QueryParaTokenInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParaTokenInfoRequest) ProtoMessage()    {}
func (*QueryParaTokenInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2bc06232cb63aa, []int{2}
}
func (m *QueryParaTokenInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParaTokenInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParaTokenInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParaTokenInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParaTokenInfoRequest.Merge(m, src)
}
func (m *QueryParaTokenInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParaTokenInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParaTokenInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParaTokenInfoRequest proto.InternalMessageInfo

func (m *QueryParaTokenInfoRequest) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

// QueryParaTokenInfoResponse is the response type for the Query/ParaTokenInfo RPC method.
type QueryParaTokenInfoResponse struct {
	IbcDenom    string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty" yaml:"ibc_denom"`
	ChannelId   string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	NativeDenom string `protobuf:"bytes,3,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
	AssetId     string `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *QueryParaTokenInfoResponse) Reset()         { *m = QueryParaTokenInfoResponse{} }
func (m *QueryParaTokenInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParaTokenInfoResponse) ProtoMessage()    {}
func (*QueryParaTokenInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc2bc06232cb63aa, []int{3}
}
func (m *QueryParaTokenInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParaTokenInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParaTokenInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParaTokenInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParaTokenInfoResponse.Merge(m, src)
}
func (m *QueryParaTokenInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParaTokenInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParaTokenInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParaTokenInfoResponse proto.InternalMessageInfo

func (m *QueryParaTokenInfoResponse) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *QueryParaTokenInfoResponse) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *QueryParaTokenInfoResponse) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

func (m *QueryParaTokenInfoResponse) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryEscrowAddressRequest)(nil), "centauri.transfermiddleware.v1beta1.QueryEscrowAddressRequest")
	proto.RegisterType((*QueryEscrowAddressResponse)(nil), "centauri.transfermiddleware.v1beta1.QueryEscrowAddressResponse")
	proto.RegisterType((*QueryParaTokenInfoRequest)(nil), "centauri.transfermiddleware.v1beta1.QueryParaTokenInfoRequest")
	proto.RegisterType((*QueryParaTokenInfoResponse)(nil), "centauri.transfermiddleware.v1beta1.QueryParaTokenInfoResponse")
}

func init() {
	proto.RegisterFile("centauri/transfermiddleware/v1beta1/query.proto", fileDescriptor_cc2bc06232cb63aa)
}

var fileDescriptor_cc2bc06232cb63aa = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0xc3, 0xab, 0x19, 0x28, 0x0f, 0xb7, 0x55, 0x89, 0x85, 0x6c, 0x34, 0x6c, 0x58, 0xd9,
	0x0a, 0x74, 0xd5, 0x05, 0x8f, 0x0a, 0x16, 0xd9, 0x51, 0x0b, 0x09, 0x89, 0x05, 0xd1, 0x8d, 0x7d,
	0x1b, 0x46, 0x24, 0x33, 0xee, 0xcc, 0xa4, 0x25, 0x5b, 0xbe, 0x00, 0x89, 0x2f, 0xe1, 0x2f, 0x58,
	0x56, 0x62, 0xc3, 0xca, 0x42, 0x09, 0x7b, 0x84, 0xbf, 0x00, 0xd9, 0x33, 0xa1, 0x4d, 0x6b, 0xa4,
	0xb4, 0xbb, 0xb9, 0x3a, 0x73, 0xce, 0x9c, 0x7b, 0x7d, 0xae, 0x49, 0x94, 0x20, 0xd7, 0x30, 0x96,
	0x2c, 0xd2, 0x12, 0xb8, 0xda, 0x43, 0x39, 0x62, 0x69, 0x3a, 0xc4, 0x43, 0x90, 0x18, 0x1d, 0x74,
	0xfa, 0xa8, 0xa1, 0x13, 0xed, 0x8f, 0x51, 0x4e, 0xc2, 0x4c, 0x0a, 0x2d, 0xdc, 0x07, 0x73, 0x42,
	0x78, 0x96, 0x10, 0x5a, 0x82, 0xb7, 0x3e, 0x10, 0x03, 0x51, 0xdd, 0x8f, 0xca, 0x93, 0xa1, 0x7a,
	0xf7, 0x06, 0x42, 0x0c, 0x86, 0x18, 0x41, 0xc6, 0x22, 0xe0, 0x5c, 0x68, 0xd0, 0x4c, 0x70, 0x65,
	0x50, 0xba, 0x4b, 0xda, 0xbb, 0xe5, 0x3b, 0x2f, 0x55, 0x22, 0xc5, 0xe1, 0xf3, 0x34, 0x95, 0xa8,
	0x54, 0x8c, 0xfb, 0x63, 0x54, 0xda, 0xdd, 0x22, 0x24, 0x79, 0x0f, 0x9c, 0xe3, 0xb0, 0xc7, 0xd2,
	0xbb, 0xce, 0x7d, 0xe7, 0x61, 0x6b, 0x67, 0xa3, 0xc8, 0x83, 0x3b, 0x13, 0x18, 0x0d, 0xb7, 0xe9,
	0x31, 0x46, 0xe3, 0x96, 0x2d, 0xba, 0x29, 0x7d, 0x47, 0xbc, 0x3a, 0x49, 0x95, 0x09, 0xae, 0xd0,
	0x7d, 0x46, 0x6e, 0x62, 0x05, 0xf4, 0xc0, 0x20, 0x56, 0xb7, 0x5d, 0xe4, 0xc1, 0x86, 0xd1, 0x5d,
	0xc4, 0x69, 0xbc, 0x8a, 0x27, 0x95, 0xe8, 0x1b, 0x6b, 0xf9, 0x15, 0x48, 0x78, 0x2d, 0x3e, 0x20,
	0xef, 0xf2, 0x3d, 0x31, 0xb7, 0xbc, 0x4d, 0x6e, 0x70, 0xd0, 0xec, 0x00, 0x7b, 0x29, 0x72, 0x31,
	0xb2, 0xe2, 0x9b, 0x45, 0x1e, 0xac, 0x19, 0xf1, 0x93, 0x28, 0x8d, 0xaf, 0x9b, 0xf2, 0x45, 0x55,
	0xfd, 0x71, 0xac, 0xf3, 0x53, 0xca, 0xd6, 0x79, 0x87, 0xb4, 0x58, 0x3f, 0x59, 0xd0, 0x5d, 0x2f,
	0xf2, 0xe0, 0xb6, 0xd1, 0xfd, 0x07, 0xd1, 0x78, 0x85, 0xf5, 0x93, 0x4a, 0xf1, 0xd4, 0x00, 0x9b,
	0xcb, 0x0d, 0xf0, 0x4c, 0x0f, 0x97, 0x96, 0xef, 0xc1, 0x0d, 0xc9, 0x0a, 0x28, 0x85, 0xba, 0x7c,
	0xef, 0x72, 0xc5, 0x5b, 0x2b, 0xf2, 0xe0, 0x96, 0xe1, 0xcd, 0x11, 0x1a, 0x5f, 0xab, 0x8e, 0xdd,
	0xf4, 0xd1, 0xef, 0x26, 0xb9, 0x52, 0xf5, 0xec, 0x7e, 0x75, 0xc8, 0xea, 0x42, 0xe3, 0xee, 0x93,
	0x70, 0x89, 0xd4, 0x85, 0xff, 0xfd, 0x16, 0xde, 0xd3, 0x0b, 0xf3, 0xcd, 0xc4, 0x69, 0xf0, 0xe9,
	0xfb, 0xaf, 0x2f, 0xcd, 0xb6, 0xbb, 0x79, 0xbc, 0x2f, 0x19, 0x48, 0xd0, 0xe5, 0x45, 0x56, 0x3a,
	0x2c, 0x3d, 0x2f, 0xc4, 0xec, 0x3c, 0x9e, 0xeb, 0x22, 0x7f, 0x1e, 0xcf, 0xb5, 0xf9, 0xae, 0xf3,
	0x6c, 0xe2, 0x6b, 0xe3, 0xbc, 0xb3, 0xf5, 0x6d, 0xea, 0x3b, 0x47, 0x53, 0xdf, 0xf9, 0x39, 0xf5,
	0x9d, 0xcf, 0x33, 0xbf, 0x71, 0x34, 0xf3, 0x1b, 0x3f, 0x66, 0x7e, 0xe3, 0xad, 0xf7, 0xb1, 0xee,
	0x77, 0xa0, 0x27, 0x19, 0xaa, 0xfe, 0xd5, 0x6a, 0x5d, 0x1f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0xa5, 0xa0, 0xee, 0x3a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ParaTokenInfo queries all token info of a native denom.
	ParaTokenInfo(ctx context.Context, in *QueryParaTokenInfoRequest, opts ...grpc.CallOption) (*QueryParaTokenInfoResponse, error)
	EscrowAddress(ctx context.Context, in *QueryEscrowAddressRequest, opts ...grpc.CallOption) (*QueryEscrowAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ParaTokenInfo(ctx context.Context, in *QueryParaTokenInfoRequest, opts ...grpc.CallOption) (*QueryParaTokenInfoResponse, error) {
	out := new(QueryParaTokenInfoResponse)
	err := c.cc.Invoke(ctx, "/centauri.transfermiddleware.v1beta1.Query/ParaTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EscrowAddress(ctx context.Context, in *QueryEscrowAddressRequest, opts ...grpc.CallOption) (*QueryEscrowAddressResponse, error) {
	out := new(QueryEscrowAddressResponse)
	err := c.cc.Invoke(ctx, "/centauri.transfermiddleware.v1beta1.Query/EscrowAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ParaTokenInfo queries all token info of a native denom.
	ParaTokenInfo(context.Context, *QueryParaTokenInfoRequest) (*QueryParaTokenInfoResponse, error)
	EscrowAddress(context.Context, *QueryEscrowAddressRequest) (*QueryEscrowAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ParaTokenInfo(ctx context.Context, req *QueryParaTokenInfoRequest) (*QueryParaTokenInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParaTokenInfo not implemented")
}
func (*UnimplementedQueryServer) EscrowAddress(ctx context.Context, req *QueryEscrowAddressRequest) (*QueryEscrowAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscrowAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ParaTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParaTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ParaTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/centauri.transfermiddleware.v1beta1.Query/ParaTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ParaTokenInfo(ctx, req.(*QueryParaTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EscrowAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEscrowAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EscrowAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/centauri.transfermiddleware.v1beta1.Query/EscrowAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EscrowAddress(ctx, req.(*QueryEscrowAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "centauri.transfermiddleware.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParaTokenInfo",
			Handler:    _Query_ParaTokenInfo_Handler,
		},
		{
			MethodName: "EscrowAddress",
			Handler:    _Query_EscrowAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centauri/transfermiddleware/v1beta1/query.proto",
}

func (m *QueryEscrowAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEscrowAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EscrowAddress) > 0 {
		i -= len(m.EscrowAddress)
		copy(dAtA[i:], m.EscrowAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EscrowAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParaTokenInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParaTokenInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParaTokenInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParaTokenInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParaTokenInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParaTokenInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEscrowAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEscrowAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EscrowAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParaTokenInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParaTokenInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEscrowAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEscrowAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEscrowAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEscrowAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParaTokenInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParaTokenInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParaTokenInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParaTokenInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParaTokenInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParaTokenInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
