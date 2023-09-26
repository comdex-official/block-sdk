// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/proposals/v1/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// ProposalInfo contains the metadata about a given proposal that was built by
// the block-sdk. This is used to verify and consilidate proposal data across
// the network.
type ProposalInfo struct {
	// Lanes contains information about how each partial proposal
	// was constructed by the block-sdk lanes.
	Lanes map[string]*LaneInfo `protobuf:"bytes,1,rep,name=lanes,proto3" json:"lanes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ProposalInfo) Reset()         { *m = ProposalInfo{} }
func (m *ProposalInfo) String() string { return proto.CompactTextString(m) }
func (*ProposalInfo) ProtoMessage()    {}
func (*ProposalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6b8540ee6bc1e, []int{0}
}
func (m *ProposalInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalInfo.Merge(m, src)
}
func (m *ProposalInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProposalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalInfo proto.InternalMessageInfo

func (m *ProposalInfo) GetLanes() map[string]*LaneInfo {
	if m != nil {
		return m.Lanes
	}
	return nil
}

// LaneInfo contains the metadata about a given lane that was included in a
// proposal.
type LaneInfo struct {
	// NumTxs is the number of transactions in the proposal that were included
	// from this lane.
	NumTxs uint64 `protobuf:"varint,1,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
}

func (m *LaneInfo) Reset()         { *m = LaneInfo{} }
func (m *LaneInfo) String() string { return proto.CompactTextString(m) }
func (*LaneInfo) ProtoMessage()    {}
func (*LaneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d6b8540ee6bc1e, []int{1}
}
func (m *LaneInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LaneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LaneInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LaneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaneInfo.Merge(m, src)
}
func (m *LaneInfo) XXX_Size() int {
	return m.Size()
}
func (m *LaneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LaneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LaneInfo proto.InternalMessageInfo

func (m *LaneInfo) GetNumTxs() uint64 {
	if m != nil {
		return m.NumTxs
	}
	return 0
}

func init() {
	proto.RegisterType((*ProposalInfo)(nil), "sdk.proposals.v1.ProposalInfo")
	proto.RegisterMapType((map[string]*LaneInfo)(nil), "sdk.proposals.v1.ProposalInfo.LanesEntry")
	proto.RegisterType((*LaneInfo)(nil), "sdk.proposals.v1.LaneInfo")
}

func init() { proto.RegisterFile("sdk/proposals/v1/types.proto", fileDescriptor_b5d6b8540ee6bc1e) }

var fileDescriptor_b5d6b8540ee6bc1e = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4e, 0xc9, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0x29, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x4e, 0xc9, 0xd6, 0x83, 0xcb,
	0xea, 0x95, 0x19, 0x2a, 0x2d, 0x65, 0xe4, 0xe2, 0x09, 0x80, 0x0a, 0x78, 0xe6, 0xa5, 0xe5, 0x0b,
	0xd9, 0x73, 0xb1, 0xe6, 0x24, 0xe6, 0xa5, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x69,
	0xea, 0xa1, 0x6b, 0xd1, 0x43, 0x56, 0xae, 0xe7, 0x03, 0x52, 0xeb, 0x9a, 0x57, 0x52, 0x54, 0x19,
	0x04, 0xd1, 0x27, 0x15, 0xc2, 0xc5, 0x85, 0x10, 0x14, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x0c, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0x30, 0x2d, 0x00, 0x69, 0x07, 0x19, 0x1e,
	0x04, 0x51, 0x68, 0xc5, 0x64, 0xc1, 0xa8, 0xa4, 0xcc, 0xc5, 0x01, 0x13, 0x16, 0x12, 0xe7, 0x62,
	0xcf, 0x2b, 0xcd, 0x8d, 0x2f, 0xa9, 0x28, 0x06, 0x9b, 0xcb, 0x12, 0xc4, 0x96, 0x57, 0x9a, 0x1b,
	0x52, 0x51, 0xec, 0xe4, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc6,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5, 0xd9, 0x99, 0x05, 0xba,
	0xb9, 0xa9, 0x65, 0xfa, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xba, 0xa0, 0x00, 0x03, 0xb3, 0x90, 0x82,
	0x0d, 0x1c, 0x66, 0x49, 0x6c, 0xe0, 0x40, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x85,
	0x41, 0xd1, 0x54, 0x01, 0x00, 0x00,
}

func (m *ProposalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for k := range m.Lanes {
			v := m.Lanes[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTypes(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTypes(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LaneInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LaneInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LaneInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumTxs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NumTxs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for k, v := range m.Lanes {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTypes(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTypes(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *LaneInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumTxs != 0 {
		n += 1 + sovTypes(uint64(m.NumTxs))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ProposalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lanes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lanes == nil {
				m.Lanes = make(map[string]*LaneInfo)
			}
			var mapkey string
			var mapvalue *LaneInfo
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTypes
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTypes
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &LaneInfo{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Lanes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LaneInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LaneInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LaneInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTxs", wireType)
			}
			m.NumTxs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTxs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
