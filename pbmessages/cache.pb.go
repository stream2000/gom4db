// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cache.proto

package pbmessages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type REQUEST_MSG int32

const (
	REQUEST_MSG_Get_Request REQUEST_MSG = 0
	REQUEST_MSG_Set_Request REQUEST_MSG = 1
	REQUEST_MSG_Del_Request REQUEST_MSG = 2
)

var REQUEST_MSG_name = map[int32]string{
	0: "Get_Request",
	1: "Set_Request",
	2: "Del_Request",
}

var REQUEST_MSG_value = map[string]int32{
	"Get_Request": 0,
	"Set_Request": 1,
	"Del_Request": 2,
}

func (x REQUEST_MSG) String() string {
	return proto.EnumName(REQUEST_MSG_name, int32(x))
}

func (REQUEST_MSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}

type RESPONSE_MSG int32

const (
	RESPONSE_MSG_Get_Response     RESPONSE_MSG = 0
	RESPONSE_MSG_Set_Response     RESPONSE_MSG = 1
	RESPONSE_MSG_Del_Response     RESPONSE_MSG = 2
	RESPONSE_MSG_Unknown_Response RESPONSE_MSG = 3
)

var RESPONSE_MSG_name = map[int32]string{
	0: "Get_Response",
	1: "Set_Response",
	2: "Del_Response",
	3: "Unknown_Response",
}

var RESPONSE_MSG_value = map[string]int32{
	"Get_Response":     0,
	"Set_Response":     1,
	"Del_Response":     2,
	"Unknown_Response": 3,
}

func (x RESPONSE_MSG) String() string {
	return proto.EnumName(RESPONSE_MSG_name, int32(x))
}

func (RESPONSE_MSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type SetRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}
func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{2}
}
func (m *DelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRequest.Merge(m, src)
}
func (m *DelRequest) XXX_Size() int {
	return m.Size()
}
func (m *DelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelRequest proto.InternalMessageInfo

type Request struct {
	Type                 REQUEST_MSG `protobuf:"varint,1,opt,name=type,proto3,enum=pbmessages.REQUEST_MSG" json:"type,omitempty"`
	GetRequest           *GetRequest `protobuf:"bytes,2,opt,name=getRequest,proto3" json:"getRequest,omitempty"`
	SetRequest           *SetRequest `protobuf:"bytes,3,opt,name=setRequest,proto3" json:"setRequest,omitempty"`
	DelRequest           *DelRequest `protobuf:"bytes,4,opt,name=delRequest,proto3" json:"delRequest,omitempty"`
	Key                  string      `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{3}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() REQUEST_MSG {
	if m != nil {
		return m.Type
	}
	return REQUEST_MSG_Get_Request
}

func (m *Request) GetGetRequest() *GetRequest {
	if m != nil {
		return m.GetRequest
	}
	return nil
}

func (m *Request) GetSetRequest() *SetRequest {
	if m != nil {
		return m.SetRequest
	}
	return nil
}

func (m *Request) GetDelRequest() *DelRequest {
	if m != nil {
		return m.DelRequest
	}
	return nil
}

func (m *Request) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type UnifiedResponse struct {
	Type                 RESPONSE_MSG `protobuf:"varint,1,opt,name=type,proto3,enum=pbmessages.RESPONSE_MSG" json:"type,omitempty"`
	Error                bool         `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMsg             string       `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	Value                string       `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	RedirectDir          string       `protobuf:"bytes,5,opt,name=redirectDir,proto3" json:"redirectDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnifiedResponse) Reset()         { *m = UnifiedResponse{} }
func (m *UnifiedResponse) String() string { return proto.CompactTextString(m) }
func (*UnifiedResponse) ProtoMessage()    {}
func (*UnifiedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{4}
}
func (m *UnifiedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnifiedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnifiedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnifiedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnifiedResponse.Merge(m, src)
}
func (m *UnifiedResponse) XXX_Size() int {
	return m.Size()
}
func (m *UnifiedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnifiedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnifiedResponse proto.InternalMessageInfo

func (m *UnifiedResponse) GetType() RESPONSE_MSG {
	if m != nil {
		return m.Type
	}
	return RESPONSE_MSG_Get_Response
}

func (m *UnifiedResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *UnifiedResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *UnifiedResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *UnifiedResponse) GetRedirectDir() string {
	if m != nil {
		return m.RedirectDir
	}
	return ""
}

func init() {
	proto.RegisterEnum("pbmessages.REQUEST_MSG", REQUEST_MSG_name, REQUEST_MSG_value)
	proto.RegisterEnum("pbmessages.RESPONSE_MSG", RESPONSE_MSG_name, RESPONSE_MSG_value)
	proto.RegisterType((*GetRequest)(nil), "pbmessages.GetRequest")
	proto.RegisterType((*SetRequest)(nil), "pbmessages.SetRequest")
	proto.RegisterType((*DelRequest)(nil), "pbmessages.DelRequest")
	proto.RegisterType((*Request)(nil), "pbmessages.Request")
	proto.RegisterType((*UnifiedResponse)(nil), "pbmessages.UnifiedResponse")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x4e, 0xc2, 0x40,
	0x14, 0xc7, 0x19, 0x3e, 0x14, 0x5e, 0x89, 0x34, 0x13, 0xa2, 0x8d, 0x8b, 0x86, 0x74, 0x45, 0xd0,
	0xb0, 0xc0, 0xc4, 0xb5, 0x31, 0x10, 0x56, 0xf8, 0x31, 0x23, 0x1b, 0x37, 0x84, 0x8f, 0x27, 0x36,
	0x60, 0x5b, 0x67, 0x8a, 0x86, 0x9b, 0x78, 0x05, 0x6f, 0xe2, 0xd2, 0x23, 0x18, 0xbc, 0x80, 0x47,
	0x30, 0xd3, 0x42, 0x3b, 0x0d, 0xbb, 0x79, 0xbf, 0xbc, 0xff, 0xcc, 0xfb, 0xcd, 0x0c, 0x18, 0xd3,
	0xf1, 0xf4, 0x19, 0xdb, 0x81, 0xf0, 0x43, 0x9f, 0x42, 0x30, 0x79, 0x41, 0x29, 0xc7, 0x73, 0x94,
	0x4e, 0x15, 0xa0, 0x8f, 0x21, 0xc3, 0xd7, 0x15, 0xca, 0xd0, 0x71, 0x00, 0x78, 0x52, 0xd1, 0x3a,
	0x94, 0xde, 0xc6, 0xcb, 0x15, 0x5a, 0xa4, 0x41, 0x9a, 0x15, 0x16, 0x17, 0x2a, 0xd1, 0xc5, 0xe5,
	0x2e, 0xf1, 0x47, 0xe0, 0x70, 0xd7, 0x7f, 0x06, 0xc5, 0x70, 0x1d, 0xc4, 0xed, 0x47, 0x9d, 0x93,
	0x76, 0x7a, 0x4c, 0x9b, 0xf5, 0xee, 0x87, 0x3d, 0xfe, 0x30, 0x1a, 0xf0, 0x3e, 0x8b, 0x9a, 0xe8,
	0x25, 0xc0, 0x3c, 0x39, 0xca, 0xca, 0x37, 0x48, 0xd3, 0xe8, 0x1c, 0xeb, 0x91, 0x74, 0x2c, 0xa6,
	0x75, 0xaa, 0x9c, 0x4c, 0x73, 0x85, 0xfd, 0x1c, 0xd7, 0x72, 0x32, 0x93, 0x9b, 0x25, 0x63, 0x5b,
	0xc5, 0xfd, 0x5c, 0x2a, 0xc5, 0xb4, 0x4e, 0x6a, 0x42, 0x61, 0x81, 0x6b, 0xab, 0x14, 0x5d, 0x81,
	0x5a, 0x3a, 0x9f, 0x04, 0x6a, 0x43, 0xcf, 0x7d, 0x72, 0x71, 0xc6, 0x50, 0x06, 0xbe, 0x27, 0x91,
	0x9e, 0x67, 0xd4, 0xad, 0xac, 0x3a, 0xbf, 0xbb, 0xbd, 0xe1, 0x3d, 0xcd, 0xbd, 0x0e, 0x25, 0x14,
	0xc2, 0x17, 0x91, 0x76, 0x99, 0xc5, 0x05, 0x3d, 0x85, 0x72, 0xb4, 0x18, 0xc8, 0x79, 0xe4, 0x55,
	0x61, 0x49, 0x9d, 0x3e, 0x45, 0x51, 0x7b, 0x0a, 0xda, 0x00, 0x43, 0xe0, 0xcc, 0x15, 0x38, 0x0d,
	0xbb, 0xae, 0xd8, 0xce, 0xa8, 0xa3, 0xd6, 0x15, 0x18, 0xda, 0xd5, 0xd3, 0x1a, 0x18, 0x7d, 0x0c,
	0x47, 0x5b, 0x37, 0x33, 0xa7, 0x00, 0xd7, 0x00, 0x51, 0xa0, 0x8b, 0xcb, 0x04, 0xe4, 0x5b, 0x8f,
	0x50, 0xd5, 0x0d, 0xa8, 0x09, 0xd5, 0x78, 0x8b, 0xd8, 0xdc, 0xcc, 0x29, 0xc2, 0x75, 0x42, 0x14,
	0x89, 0x37, 0xd9, 0x92, 0x3c, 0xad, 0x83, 0x39, 0xf4, 0x16, 0x9e, 0xff, 0xee, 0xa5, 0xb4, 0x70,
	0x6d, 0x7e, 0x6d, 0x6c, 0xf2, 0xbd, 0xb1, 0xc9, 0xcf, 0xc6, 0x26, 0x1f, 0xbf, 0x76, 0x6e, 0x72,
	0x10, 0xfd, 0xd0, 0x8b, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x0c, 0x1c, 0xb5, 0xb0, 0x02,
	0x00, 0x00,
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *SetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCache(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCache(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x2a
	}
	if m.DelRequest != nil {
		{
			size, err := m.DelRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.SetRequest != nil {
		{
			size, err := m.SetRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.GetRequest != nil {
		{
			size, err := m.GetRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCache(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintCache(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UnifiedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnifiedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnifiedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RedirectDir) > 0 {
		i -= len(m.RedirectDir)
		copy(dAtA[i:], m.RedirectDir)
		i = encodeVarintCache(dAtA, i, uint64(len(m.RedirectDir)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCache(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ErrorMsg) > 0 {
		i -= len(m.ErrorMsg)
		copy(dAtA[i:], m.ErrorMsg)
		i = encodeVarintCache(dAtA, i, uint64(len(m.ErrorMsg)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Error {
		i--
		if m.Error {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintCache(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCache(dAtA []byte, offset int, v uint64) int {
	offset -= sovCache(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCache(uint64(m.Type))
	}
	if m.GetRequest != nil {
		l = m.GetRequest.Size()
		n += 1 + l + sovCache(uint64(l))
	}
	if m.SetRequest != nil {
		l = m.SetRequest.Size()
		n += 1 + l + sovCache(uint64(l))
	}
	if m.DelRequest != nil {
		l = m.DelRequest.Size()
		n += 1 + l + sovCache(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UnifiedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCache(uint64(m.Type))
	}
	if m.Error {
		n += 2
	}
	l = len(m.ErrorMsg)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	l = len(m.RedirectDir)
	if l > 0 {
		n += 1 + l + sovCache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCache(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCache(x uint64) (n int) {
	return sovCache(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
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
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
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
			return fmt.Errorf("proto: SetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
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
			return fmt.Errorf("proto: DelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= REQUEST_MSG(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GetRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GetRequest == nil {
				m.GetRequest = &GetRequest{}
			}
			if err := m.GetRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SetRequest == nil {
				m.SetRequest = &SetRequest{}
			}
			if err := m.SetRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DelRequest == nil {
				m.DelRequest = &DelRequest{}
			}
			if err := m.DelRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnifiedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
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
			return fmt.Errorf("proto: UnifiedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnifiedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= RESPONSE_MSG(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Error = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedirectDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
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
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCache
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedirectDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCache(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCache
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
					return 0, ErrIntOverflowCache
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
					return 0, ErrIntOverflowCache
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
				return 0, ErrInvalidLengthCache
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCache
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCache
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCache        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCache          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCache = fmt.Errorf("proto: unexpected end of group")
)
