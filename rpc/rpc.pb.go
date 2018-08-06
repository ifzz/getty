// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: getty.proto

/*
	Package prc is a generated protocol buffer package.

	It is generated from these files:
		getty.proto

	It has these top-level messages:
		GettyRPCRequestHeader
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CallType int32

const (
	CT_UNKOWN        CallType = 0
	CT_OneWay        CallType = 1
	CT_TwoWay        CallType = 2
	CT_TwoWayNoReply CallType = 3
)

var CallType_name = map[int32]string{
	0: "CT_UNKOWN",
	1: "CT_OneWay",
	2: "CT_TwoWay",
	3: "CT_TwoWayNoReply",
}
var CallType_value = map[string]int32{
	"CT_UNKOWN":        0,
	"CT_OneWay":        1,
	"CT_TwoWay":        2,
	"CT_TwoWayNoReply": 3,
}

func (x CallType) Enum() *CallType {
	p := new(CallType)
	*p = x
	return p
}
func (x CallType) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(CallType_name, int32(x))
}
func (x *CallType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CallType_value, data, "CallType")
	if err != nil {
		return err
	}
	*x = CallType(value)
	return nil
}
func (CallType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGetty, []int{0} }

type GettyRPCRequestHeader struct {
	Service  string   `protobuf:"bytes,1,opt,name=Service" json:"Service"`
	Method   string   `protobuf:"bytes,2,opt,name=Method" json:"Method"`
	CallType CallType `protobuf:"varint,3,opt,name=CallType,enum=prc.CallType" json:"CallType"`
}

func (m *GettyRPCRequestHeader) Reset()                    { *m = GettyRPCRequestHeader{} }
func (*GettyRPCRequestHeader) ProtoMessage()               {}
func (*GettyRPCRequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorGetty, []int{0} }

func init() {
	proto.RegisterType((*GettyRPCRequestHeader)(nil), "prc.GettyRPCRequestHeader")
	proto.RegisterEnum("prc.CallType", CallType_name, CallType_value)
}
func (x CallType) String() string {
	s, ok := CallType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *GettyRPCRequestHeader) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GettyRPCRequestHeader)
	if !ok {
		that2, ok := that.(GettyRPCRequestHeader)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GettyRPCRequestHeader")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GettyRPCRequestHeader but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GettyRPCRequestHeader but is not nil && this == nil")
	}
	if this.Service != that1.Service {
		return fmt.Errorf("Service this(%v) Not Equal that(%v)", this.Service, that1.Service)
	}
	if this.Method != that1.Method {
		return fmt.Errorf("Method this(%v) Not Equal that(%v)", this.Method, that1.Method)
	}
	if this.CallType != that1.CallType {
		return fmt.Errorf("CallType this(%v) Not Equal that(%v)", this.CallType, that1.CallType)
	}
	return nil
}
func (this *GettyRPCRequestHeader) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GettyRPCRequestHeader)
	if !ok {
		that2, ok := that.(GettyRPCRequestHeader)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if this.CallType != that1.CallType {
		return false
	}
	return true
}
func (this *GettyRPCRequestHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prc.GettyRPCRequestHeader{")
	s = append(s, "Service: "+fmt.Sprintf("%#v", this.Service)+",\n")
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "CallType: "+fmt.Sprintf("%#v", this.CallType)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGetty(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GettyRPCRequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GettyRPCRequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGetty(dAtA, i, uint64(len(m.Service)))
	i += copy(dAtA[i:], m.Service)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGetty(dAtA, i, uint64(len(m.Method)))
	i += copy(dAtA[i:], m.Method)
	dAtA[i] = 0x18
	i++
	i = encodeVarintGetty(dAtA, i, uint64(m.CallType))
	return i, nil
}

func encodeFixed64Getty(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Getty(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGetty(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GettyRPCRequestHeader) Size() (n int) {
	var l int
	_ = l
	l = len(m.Service)
	n += 1 + l + sovGetty(uint64(l))
	l = len(m.Method)
	n += 1 + l + sovGetty(uint64(l))
	n += 1 + sovGetty(uint64(m.CallType))
	return n
}

func sovGetty(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGetty(x uint64) (n int) {
	return sovGetty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GettyRPCRequestHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GettyRPCRequestHeader{`,
		`Service:` + fmt.Sprintf("%v", this.Service) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`CallType:` + fmt.Sprintf("%v", this.CallType) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGetty(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GettyRPCRequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGetty
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
			return fmt.Errorf("proto: GettyRPCRequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GettyRPCRequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetty
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
				return ErrInvalidLengthGetty
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetty
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
				return ErrInvalidLengthGetty
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallType", wireType)
			}
			m.CallType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetty
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallType |= (CallType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGetty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGetty
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
func skipGetty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGetty
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
					return 0, ErrIntOverflowGetty
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
					return 0, ErrIntOverflowGetty
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
				return 0, ErrInvalidLengthGetty
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGetty
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
				next, err := skipGetty(dAtA[start:])
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
	ErrInvalidLengthGetty = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGetty   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("getty.proto", fileDescriptorGetty) }

var fileDescriptorGetty = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4f, 0x2d, 0x29,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x28, 0x4a, 0x96, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb,
	0x25, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa3, 0xd4, 0xc6, 0xc8, 0x25, 0xea,
	0x0e, 0x32, 0x23, 0x28, 0xc0, 0x39, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0x23, 0x35, 0x31,
	0x25, 0xb5, 0x48, 0x48, 0x8e, 0x8b, 0x3d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xd3, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x98, 0xa0, 0x90, 0x0c, 0x17,
	0x9b, 0x6f, 0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x04, 0x13, 0x92, 0x34, 0x54, 0x4c, 0x48, 0x9f, 0x8b,
	0xc3, 0x39, 0x31, 0x27, 0x27, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0xcf, 0x88,
	0x57, 0xaf, 0xa0, 0x28, 0x59, 0x0f, 0x26, 0x08, 0x55, 0x0e, 0x57, 0xa4, 0xe5, 0x8b, 0xd0, 0x20,
	0xc4, 0xcb, 0xc5, 0xe9, 0x1c, 0x12, 0x1f, 0xea, 0xe7, 0xed, 0x1f, 0xee, 0x27, 0xc0, 0x00, 0xe5,
	0xfa, 0xe7, 0xa5, 0x86, 0x27, 0x56, 0x0a, 0x30, 0x42, 0xb9, 0x21, 0xe5, 0xf9, 0x20, 0x2e, 0x93,
	0x90, 0x08, 0x97, 0x00, 0x9c, 0xeb, 0x97, 0x1f, 0x94, 0x5a, 0x90, 0x53, 0x29, 0xc0, 0xec, 0x64,
	0x72, 0xe2, 0xa1, 0x1c, 0xc3, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x3c, 0x78, 0x28,
	0xc7, 0xf8, 0xe1, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0x37, 0xfb, 0x13,
	0x0a, 0x4f, 0x01, 0x00, 0x00,
}
