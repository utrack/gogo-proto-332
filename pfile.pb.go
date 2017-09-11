// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pfile.proto

/*
Package listingpb is a generated protocol buffer package.

It is generated from these files:
	pfile.proto

It has these top-level messages:
	TestMsg
	Foo
*/
package listingpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestMsg struct {
	// repeated Foo foos = 1 [(gogoproto.nullable) = false];
	// Foo bar = 2 [(gogoproto.nullable) = false];
	FooMap map[string]Foo `protobuf:"bytes,3,rep,name=foo_map,json=fooMap" json:"foo_map" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TestMsg) Reset()                    { *m = TestMsg{} }
func (m *TestMsg) String() string            { return proto.CompactTextString(m) }
func (*TestMsg) ProtoMessage()               {}
func (*TestMsg) Descriptor() ([]byte, []int) { return fileDescriptorPfile, []int{0} }

func (m *TestMsg) GetFooMap() map[string]Foo {
	if m != nil {
		return m.FooMap
	}
	return nil
}

type Foo struct {
	Bar string `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptorPfile, []int{1} }

func (m *Foo) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func init() {
	proto.RegisterType((*TestMsg)(nil), "listingpb.TestMsg")
	golang_proto.RegisterType((*TestMsg)(nil), "listingpb.TestMsg")
	proto.RegisterType((*Foo)(nil), "listingpb.Foo")
	golang_proto.RegisterType((*Foo)(nil), "listingpb.Foo")
}
func (m *TestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FooMap) > 0 {
		for k, _ := range m.FooMap {
			dAtA[i] = 0x1a
			i++
			v := m.FooMap[k]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovPfile(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovPfile(uint64(len(k))) + msgSize
			i = encodeVarintPfile(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPfile(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintPfile(dAtA, i, uint64((&v).Size()))
			n1, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	return i, nil
}

func (m *Foo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Bar) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPfile(dAtA, i, uint64(len(m.Bar)))
		i += copy(dAtA[i:], m.Bar)
	}
	return i, nil
}

func encodeFixed64Pfile(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Pfile(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPfile(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TestMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.FooMap) > 0 {
		for k, v := range m.FooMap {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovPfile(uint64(len(k))) + 1 + l + sovPfile(uint64(l))
			n += mapEntrySize + 1 + sovPfile(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Foo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Bar)
	if l > 0 {
		n += 1 + l + sovPfile(uint64(l))
	}
	return n
}

func sovPfile(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPfile(x uint64) (n int) {
	return sovPfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPfile
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FooMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPfile
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthPfile
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.FooMap == nil {
				m.FooMap = make(map[string]Foo)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPfile
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPfile
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthPfile
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthPfile
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &Foo{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.FooMap[mapkey] = *mapvalue
			} else {
				var mapvalue Foo
				m.FooMap[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPfile
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
func (m *Foo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPfile
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Foo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Foo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPfile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPfile
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
func skipPfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPfile
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
					return 0, ErrIntOverflowPfile
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
					return 0, ErrIntOverflowPfile
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPfile
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPfile
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
				next, err := skipPfile(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthPfile = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPfile   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pfile.proto", fileDescriptorPfile) }
func init() { golang_proto.RegisterFile("pfile.proto", fileDescriptorPfile) }

var fileDescriptorPfile = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x48, 0xcb, 0xcc,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xc9, 0x2c, 0x2e, 0xc9, 0xcc, 0x4b,
	0x2f, 0x48, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x48, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2,
	0x53, 0x69, 0x32, 0x23, 0x17, 0x7b, 0x48, 0x6a, 0x71, 0x89, 0x6f, 0x71, 0xba, 0x90, 0x2d, 0x17,
	0x7b, 0x5a, 0x7e, 0x7e, 0x7c, 0x6e, 0x62, 0x81, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x9c,
	0x1e, 0xdc, 0x5c, 0x3d, 0xa8, 0x22, 0x3d, 0xb7, 0xfc, 0x7c, 0xdf, 0xc4, 0x02, 0xd7, 0xbc, 0x92,
	0xa2, 0x4a, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xd8, 0xd2, 0xc0, 0x42, 0x52, 0x9e, 0x5c,
	0xdc, 0x48, 0x92, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0xa6, 0x90, 0x0a, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0xb7, 0x11, 0x1f, 0x92, 0xe9, 0x6e, 0xf9, 0xf9, 0x41, 0x10, 0x49, 0x2b, 0x26, 0x0b, 0x46,
	0x25, 0x71, 0x2e, 0x66, 0xb7, 0xfc, 0x7c, 0x90, 0x11, 0x49, 0x89, 0x45, 0x30, 0x23, 0x92, 0x12,
	0x8b, 0x9c, 0x44, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x0e, 0x3c, 0x96, 0x63, 0x4c, 0x62, 0x03, 0xfb, 0xc5, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0x70, 0xbc, 0x6b, 0x14, 0x01, 0x00, 0x00,
}