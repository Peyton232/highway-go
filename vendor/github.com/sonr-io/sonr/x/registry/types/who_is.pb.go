// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/who_is.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// WhoIs is the entry pointing a registered name to a user account address, Did Url string, and a DIDDocument.
type WhoIs struct {
	// Name is the registered name of the User
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// DID is the DID of the account
	Did string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	// Document is the DID Document of the registered name and account encoded as JSON
	Document []byte `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
	// Creator is the DID of the creator of the DID Document
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *WhoIs) Reset()         { *m = WhoIs{} }
func (m *WhoIs) String() string { return proto.CompactTextString(m) }
func (*WhoIs) ProtoMessage()    {}
func (*WhoIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d9bdfc8d37d9424, []int{0}
}
func (m *WhoIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhoIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhoIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhoIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoIs.Merge(m, src)
}
func (m *WhoIs) XXX_Size() int {
	return m.Size()
}
func (m *WhoIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhoIs proto.InternalMessageInfo

func (m *WhoIs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WhoIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhoIs) GetDocument() []byte {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *WhoIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*WhoIs)(nil), "sonrio.sonr.registry.WhoIs")
}

func init() { proto.RegisterFile("registry/who_is.proto", fileDescriptor_9d9bdfc8d37d9424) }

var fileDescriptor_9d9bdfc8d37d9424 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0xcf, 0xc8, 0x8f, 0xcf, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x29, 0xce, 0xcf, 0x2b, 0xca, 0xcc, 0xd7, 0x03, 0x51, 0x7a, 0x30, 0x25, 0x4a,
	0xc9, 0x5c, 0xac, 0xe1, 0x19, 0xf9, 0x9e, 0xc5, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x4a, 0x66, 0x8a,
	0x04, 0x13, 0x58, 0x08, 0xc4, 0x14, 0x92, 0xe2, 0xe2, 0x48, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd,
	0x2b, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf3, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b,
	0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x58, 0xc0, 0x3a, 0x60, 0x5c, 0x27, 0xa7, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x07, 0x39, 0x4c, 0x37, 0x33, 0x1f, 0x4c, 0xeb, 0x57, 0xe8, 0xc3, 0x7d, 0x51,
	0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x85, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x8f, 0x37, 0xf5, 0xde, 0x00, 0x00, 0x00,
}

func (m *WhoIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhoIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhoIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Document) > 0 {
		i -= len(m.Document)
		copy(dAtA[i:], m.Document)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Document)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhoIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhoIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhoIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	l = len(m.Document)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	return n
}

func sovWhoIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhoIs(x uint64) (n int) {
	return sovWhoIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhoIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhoIs
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
			return fmt.Errorf("proto: WhoIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhoIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Document = append(m.Document[:0], dAtA[iNdEx:postIndex]...)
			if m.Document == nil {
				m.Document = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhoIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhoIs
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
func skipWhoIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhoIs
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
					return 0, ErrIntOverflowWhoIs
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
					return 0, ErrIntOverflowWhoIs
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
				return 0, ErrInvalidLengthWhoIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhoIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhoIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhoIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhoIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhoIs = fmt.Errorf("proto: unexpected end of group")
)