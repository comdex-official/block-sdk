// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/blocksdk/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryLaneRequest is the request type for the Query/Lane RPC method.
type QueryLaneRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryLaneRequest) Reset()         { *m = QueryLaneRequest{} }
func (m *QueryLaneRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLaneRequest) ProtoMessage()    {}
func (*QueryLaneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{0}
}
func (m *QueryLaneRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLaneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLaneRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLaneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLaneRequest.Merge(m, src)
}
func (m *QueryLaneRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLaneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLaneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLaneRequest proto.InternalMessageInfo

func (m *QueryLaneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryLaneResponse is the response type for the Query/Lane RPC method.
type QueryLaneResponse struct {
	Lane Lane `protobuf:"bytes,1,opt,name=lane,proto3" json:"lane"`
}

func (m *QueryLaneResponse) Reset()         { *m = QueryLaneResponse{} }
func (m *QueryLaneResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLaneResponse) ProtoMessage()    {}
func (*QueryLaneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{1}
}
func (m *QueryLaneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLaneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLaneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLaneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLaneResponse.Merge(m, src)
}
func (m *QueryLaneResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLaneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLaneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLaneResponse proto.InternalMessageInfo

func (m *QueryLaneResponse) GetLane() Lane {
	if m != nil {
		return m.Lane
	}
	return Lane{}
}

// QueryLaneRequest is the request type for the Query/Lanes RPC method.
type QueryLanesRequest struct {
}

func (m *QueryLanesRequest) Reset()         { *m = QueryLanesRequest{} }
func (m *QueryLanesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLanesRequest) ProtoMessage()    {}
func (*QueryLanesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{2}
}
func (m *QueryLanesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLanesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLanesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLanesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLanesRequest.Merge(m, src)
}
func (m *QueryLanesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLanesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLanesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLanesRequest proto.InternalMessageInfo

// QueryLaneResponse is the response type for the Query/Lanes RPC method.
type QueryLanesResponse struct {
	Lanes []Lane `protobuf:"bytes,1,rep,name=lanes,proto3" json:"lanes"`
}

func (m *QueryLanesResponse) Reset()         { *m = QueryLanesResponse{} }
func (m *QueryLanesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLanesResponse) ProtoMessage()    {}
func (*QueryLanesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{3}
}
func (m *QueryLanesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLanesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLanesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLanesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLanesResponse.Merge(m, src)
}
func (m *QueryLanesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLanesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLanesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLanesResponse proto.InternalMessageInfo

func (m *QueryLanesResponse) GetLanes() []Lane {
	if m != nil {
		return m.Lanes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryLaneRequest)(nil), "sdk.blocksdk.v1.QueryLaneRequest")
	proto.RegisterType((*QueryLaneResponse)(nil), "sdk.blocksdk.v1.QueryLaneResponse")
	proto.RegisterType((*QueryLanesRequest)(nil), "sdk.blocksdk.v1.QueryLanesRequest")
	proto.RegisterType((*QueryLanesResponse)(nil), "sdk.blocksdk.v1.QueryLanesResponse")
}

func init() { proto.RegisterFile("sdk/blocksdk/v1/query.proto", fileDescriptor_0c631e2e97c81d34) }

var fileDescriptor_0c631e2e97c81d34 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x93, 0xdc, 0xf6, 0xc2, 0x9d, 0x0b, 0xfe, 0x19, 0x15, 0x4a, 0x5a, 0xc6, 0x3a, 0x6e,
	0x54, 0xe8, 0x0c, 0xad, 0x6f, 0x50, 0x04, 0x5d, 0xb8, 0xb1, 0x4b, 0x77, 0x69, 0x33, 0xc4, 0x21,
	0x6d, 0x26, 0xed, 0xa4, 0xc5, 0x52, 0xdd, 0x74, 0xe5, 0x52, 0xf0, 0x45, 0x7c, 0x8c, 0x2e, 0x0b,
	0x6e, 0x5c, 0x89, 0xb4, 0x82, 0xaf, 0x21, 0x33, 0x49, 0x6d, 0x88, 0x10, 0x77, 0x87, 0x39, 0xdf,
	0x7c, 0xe7, 0x77, 0x86, 0x01, 0x65, 0xe9, 0xfa, 0xb4, 0xdd, 0x15, 0x1d, 0x5f, 0x15, 0xa3, 0x3a,
	0xed, 0x0f, 0xd9, 0x60, 0x4c, 0xc2, 0x81, 0x88, 0x04, 0xdc, 0x94, 0xae, 0x4f, 0x56, 0x4d, 0x32,
	0xaa, 0xdb, 0x15, 0x4f, 0x08, 0xaf, 0xcb, 0xa8, 0x13, 0x72, 0xea, 0x04, 0x81, 0x88, 0x9c, 0x88,
	0x8b, 0x40, 0xc6, 0xb8, 0xbd, 0xeb, 0x09, 0x4f, 0xe8, 0x92, 0xaa, 0x2a, 0x39, 0x2d, 0x77, 0x84,
	0xec, 0x09, 0x19, 0x8b, 0x33, 0x13, 0x6c, 0x94, 0x1d, 0xff, 0x3d, 0x4d, 0xf7, 0x31, 0x06, 0x5b,
	0x57, 0x0a, 0xbf, 0x74, 0x02, 0xd6, 0x62, 0xfd, 0x21, 0x93, 0x11, 0xdc, 0x00, 0x16, 0x77, 0x4b,
	0x66, 0xd5, 0x3c, 0xfa, 0xd7, 0xb2, 0xb8, 0x8b, 0xcf, 0xc0, 0x76, 0x8a, 0x91, 0xa1, 0x08, 0x24,
	0x83, 0x14, 0x14, 0xba, 0x4e, 0xc0, 0x34, 0xf6, 0xbf, 0xb1, 0x47, 0x32, 0x9b, 0x10, 0x05, 0x37,
	0x0b, 0xb3, 0xb7, 0x7d, 0xa3, 0xa5, 0x41, 0xbc, 0x93, 0xb2, 0xc8, 0x64, 0x14, 0x3e, 0x07, 0x30,
	0x7d, 0x98, 0xb8, 0xeb, 0xa0, 0xa8, 0xae, 0xc8, 0x92, 0x59, 0xfd, 0xf3, 0x9b, 0x3c, 0x26, 0x1b,
	0x53, 0x0b, 0x14, 0xb5, 0x09, 0xde, 0x81, 0x82, 0x6a, 0xc3, 0x83, 0x1f, 0xb7, 0xb2, 0x8b, 0xda,
	0x38, 0x0f, 0x89, 0xb3, 0xe0, 0xda, 0xc3, 0xe7, 0xf3, 0x89, 0x39, 0x7d, 0xf9, 0x78, 0xb2, 0x30,
	0xac, 0xc6, 0xcf, 0x57, 0xcb, 0x3e, 0xaa, 0xca, 0x40, 0x27, 0xdc, 0xbd, 0x87, 0x13, 0x50, 0xd4,
	0xbb, 0xc0, 0x1c, 0xf7, 0x6a, 0x7b, 0xfb, 0x30, 0x97, 0x49, 0x02, 0x1c, 0xaf, 0x03, 0x20, 0x58,
	0xc9, 0x09, 0x20, 0x9b, 0x17, 0xb3, 0x05, 0x32, 0xe7, 0x0b, 0x64, 0xbe, 0x2f, 0x90, 0xf9, 0xb8,
	0x44, 0xc6, 0x7c, 0x89, 0x8c, 0xd7, 0x25, 0x32, 0xae, 0x89, 0xc7, 0xa3, 0x9b, 0x61, 0x9b, 0x74,
	0x44, 0x8f, 0x4a, 0x9f, 0x87, 0xb5, 0x1e, 0x1b, 0xa5, 0x54, 0xb7, 0x6b, 0x59, 0x34, 0x0e, 0x99,
	0x6c, 0xff, 0xd5, 0xbf, 0xe3, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x49, 0x58, 0x65, 0x30, 0xbe,
	0x02, 0x00, 0x00,
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
	// Lane queries the a lane by its id.
	Lane(ctx context.Context, in *QueryLaneRequest, opts ...grpc.CallOption) (*QueryLaneResponse, error)
	// Lane queries all lanes in the x/blocksdk module
	Lanes(ctx context.Context, in *QueryLanesRequest, opts ...grpc.CallOption) (*QueryLanesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Lane(ctx context.Context, in *QueryLaneRequest, opts ...grpc.CallOption) (*QueryLaneResponse, error) {
	out := new(QueryLaneResponse)
	err := c.cc.Invoke(ctx, "/sdk.blocksdk.v1.Query/Lane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Lanes(ctx context.Context, in *QueryLanesRequest, opts ...grpc.CallOption) (*QueryLanesResponse, error) {
	out := new(QueryLanesResponse)
	err := c.cc.Invoke(ctx, "/sdk.blocksdk.v1.Query/Lanes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Lane queries the a lane by its id.
	Lane(context.Context, *QueryLaneRequest) (*QueryLaneResponse, error)
	// Lane queries all lanes in the x/blocksdk module
	Lanes(context.Context, *QueryLanesRequest) (*QueryLanesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Lane(ctx context.Context, req *QueryLaneRequest) (*QueryLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lane not implemented")
}
func (*UnimplementedQueryServer) Lanes(ctx context.Context, req *QueryLanesRequest) (*QueryLanesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lanes not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Lane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Lane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.blocksdk.v1.Query/Lane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Lane(ctx, req.(*QueryLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Lanes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLanesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Lanes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.blocksdk.v1.Query/Lanes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Lanes(ctx, req.(*QueryLanesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.blocksdk.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lane",
			Handler:    _Query_Lane_Handler,
		},
		{
			MethodName: "Lanes",
			Handler:    _Query_Lanes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdk/blocksdk/v1/query.proto",
}

func (m *QueryLaneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLaneRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLaneRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLaneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLaneResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLaneResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Lane.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLanesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLanesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLanesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLanesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLanesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLanesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for iNdEx := len(m.Lanes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Lanes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryLaneRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLaneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Lane.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLanesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLanesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for _, e := range m.Lanes {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryLaneRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLaneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLaneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *QueryLaneResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLaneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLaneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lane", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Lane.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *QueryLanesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLanesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLanesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryLanesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLanesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLanesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lanes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lanes = append(m.Lanes, Lane{})
			if err := m.Lanes[len(m.Lanes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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