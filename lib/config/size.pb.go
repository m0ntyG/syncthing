// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/size.proto

package config

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/syncthing/syncthing/proto/ext"
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

type Size struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value" xml:",chardata"`
	Unit  string  `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit" xml:"unit,attr"`
}

func (m *Size) Reset()      { *m = Size{} }
func (*Size) ProtoMessage() {}
func (*Size) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d75cb8f619bd299, []int{0}
}
func (m *Size) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Size) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Size.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Size) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Size.Merge(m, src)
}
func (m *Size) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Size) XXX_DiscardUnknown() {
	xxx_messageInfo_Size.DiscardUnknown(m)
}

var xxx_messageInfo_Size proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Size)(nil), "config.Size")
}

func init() { proto.RegisterFile("lib/config/size.proto", fileDescriptor_4d75cb8f619bd299) }

var fileDescriptor_4d75cb8f619bd299 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xc9, 0x4c, 0xd2,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0xce, 0xac, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x83, 0x08, 0x49, 0x29, 0x17, 0xa5, 0x16, 0xe4, 0x17, 0xeb, 0x83, 0x05, 0x93,
	0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30, 0x0b, 0xa2, 0x58, 0x8a, 0x33, 0xb5,
	0xa2, 0x04, 0xc2, 0x54, 0xea, 0x66, 0xe4, 0x62, 0x09, 0xce, 0xac, 0x4a, 0x15, 0xb2, 0xe7, 0x62,
	0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x74, 0xd2, 0x7c, 0x75, 0x4f,
	0x1e, 0x22, 0xf0, 0xe9, 0x9e, 0x3c, 0x7f, 0x45, 0x6e, 0x8e, 0x95, 0x92, 0x4e, 0x72, 0x46, 0x62,
	0x51, 0x4a, 0x62, 0x49, 0xa2, 0xd2, 0xab, 0xf3, 0x2a, 0x9c, 0x70, 0x5e, 0x10, 0x44, 0x99, 0x90,
	0x0d, 0x17, 0x4b, 0x69, 0x5e, 0x66, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc6, 0xab,
	0x7b, 0xf2, 0x60, 0x3e, 0x5c, 0x3b, 0x88, 0xa3, 0x93, 0x58, 0x52, 0x52, 0x04, 0xd6, 0x0e, 0xe7,
	0x05, 0x81, 0x55, 0x59, 0xb1, 0xcc, 0x58, 0x20, 0xcf, 0xe0, 0xe4, 0x7d, 0xe2, 0xa1, 0x1c, 0xc3,
	0x85, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x82, 0xc7, 0x72, 0x8c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x99,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x5c, 0x99, 0x97, 0x5c, 0x92,
	0x91, 0x99, 0x97, 0x8e, 0xc4, 0x42, 0x84, 0x4e, 0x12, 0x1b, 0xd8, 0x87, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x65, 0x1e, 0xa3, 0x25, 0x32, 0x01, 0x00, 0x00,
}

func (m *Size) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Size) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Size) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Unit) > 0 {
		i -= len(m.Unit)
		copy(dAtA[i:], m.Unit)
		i = encodeVarintSize(dAtA, i, uint64(len(m.Unit)))
		i--
		dAtA[i] = 0x12
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintSize(dAtA []byte, offset int, v uint64) int {
	offset -= sovSize(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Size) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovSize(uint64(l))
	}
	return n
}

func sovSize(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSize(x uint64) (n int) {
	return sovSize(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Size) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSize
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
			return fmt.Errorf("proto: Size: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Size: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSize
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
				return ErrInvalidLengthSize
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSize
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSize(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSize
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
func skipSize(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSize
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
					return 0, ErrIntOverflowSize
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
					return 0, ErrIntOverflowSize
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
				return 0, ErrInvalidLengthSize
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSize
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSize
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSize        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSize          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSize = fmt.Errorf("proto: unexpected end of group")
)
