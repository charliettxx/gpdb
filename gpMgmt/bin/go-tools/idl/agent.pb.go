// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package idl

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type StartSegmentRequest struct {
	DataDir              string   `protobuf:"bytes,1,opt,name=dataDir,proto3" json:"dataDir,omitempty"`
	Wait                 bool     `protobuf:"varint,2,opt,name=wait,proto3" json:"wait,omitempty"`
	Timeout              int32    `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Options              string   `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartSegmentRequest) Reset()         { *m = StartSegmentRequest{} }
func (m *StartSegmentRequest) String() string { return proto.CompactTextString(m) }
func (*StartSegmentRequest) ProtoMessage()    {}
func (*StartSegmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *StartSegmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartSegmentRequest.Unmarshal(m, b)
}
func (m *StartSegmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartSegmentRequest.Marshal(b, m, deterministic)
}
func (m *StartSegmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartSegmentRequest.Merge(m, src)
}
func (m *StartSegmentRequest) XXX_Size() int {
	return xxx_messageInfo_StartSegmentRequest.Size(m)
}
func (m *StartSegmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartSegmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartSegmentRequest proto.InternalMessageInfo

func (m *StartSegmentRequest) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

func (m *StartSegmentRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *StartSegmentRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *StartSegmentRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

type StartSegmentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartSegmentReply) Reset()         { *m = StartSegmentReply{} }
func (m *StartSegmentReply) String() string { return proto.CompactTextString(m) }
func (*StartSegmentReply) ProtoMessage()    {}
func (*StartSegmentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *StartSegmentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartSegmentReply.Unmarshal(m, b)
}
func (m *StartSegmentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartSegmentReply.Marshal(b, m, deterministic)
}
func (m *StartSegmentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartSegmentReply.Merge(m, src)
}
func (m *StartSegmentReply) XXX_Size() int {
	return xxx_messageInfo_StartSegmentReply.Size(m)
}
func (m *StartSegmentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StartSegmentReply.DiscardUnknown(m)
}

var xxx_messageInfo_StartSegmentReply proto.InternalMessageInfo

type StopAgentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentRequest) Reset()         { *m = StopAgentRequest{} }
func (m *StopAgentRequest) String() string { return proto.CompactTextString(m) }
func (*StopAgentRequest) ProtoMessage()    {}
func (*StopAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *StopAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentRequest.Unmarshal(m, b)
}
func (m *StopAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentRequest.Marshal(b, m, deterministic)
}
func (m *StopAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentRequest.Merge(m, src)
}
func (m *StopAgentRequest) XXX_Size() int {
	return xxx_messageInfo_StopAgentRequest.Size(m)
}
func (m *StopAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentRequest proto.InternalMessageInfo

type StopAgentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentReply) Reset()         { *m = StopAgentReply{} }
func (m *StopAgentReply) String() string { return proto.CompactTextString(m) }
func (*StopAgentReply) ProtoMessage()    {}
func (*StopAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *StopAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentReply.Unmarshal(m, b)
}
func (m *StopAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentReply.Marshal(b, m, deterministic)
}
func (m *StopAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentReply.Merge(m, src)
}
func (m *StopAgentReply) XXX_Size() int {
	return xxx_messageInfo_StopAgentReply.Size(m)
}
func (m *StopAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentReply proto.InternalMessageInfo

type StatusAgentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusAgentRequest) Reset()         { *m = StatusAgentRequest{} }
func (m *StatusAgentRequest) String() string { return proto.CompactTextString(m) }
func (*StatusAgentRequest) ProtoMessage()    {}
func (*StatusAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{4}
}

func (m *StatusAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusAgentRequest.Unmarshal(m, b)
}
func (m *StatusAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusAgentRequest.Marshal(b, m, deterministic)
}
func (m *StatusAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusAgentRequest.Merge(m, src)
}
func (m *StatusAgentRequest) XXX_Size() int {
	return xxx_messageInfo_StatusAgentRequest.Size(m)
}
func (m *StatusAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusAgentRequest proto.InternalMessageInfo

type StatusAgentReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Uptime               string   `protobuf:"bytes,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Pid                  uint32   `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusAgentReply) Reset()         { *m = StatusAgentReply{} }
func (m *StatusAgentReply) String() string { return proto.CompactTextString(m) }
func (*StatusAgentReply) ProtoMessage()    {}
func (*StatusAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{5}
}

func (m *StatusAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusAgentReply.Unmarshal(m, b)
}
func (m *StatusAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusAgentReply.Marshal(b, m, deterministic)
}
func (m *StatusAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusAgentReply.Merge(m, src)
}
func (m *StatusAgentReply) XXX_Size() int {
	return xxx_messageInfo_StatusAgentReply.Size(m)
}
func (m *StatusAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusAgentReply proto.InternalMessageInfo

func (m *StatusAgentReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StatusAgentReply) GetUptime() string {
	if m != nil {
		return m.Uptime
	}
	return ""
}

func (m *StatusAgentReply) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type ValidateHostEnvRequest struct {
	HostAddressList      []string `protobuf:"bytes,1,rep,name=hostAddressList,proto3" json:"hostAddressList,omitempty"`
	DirectoryList        []string `protobuf:"bytes,2,rep,name=DirectoryList,proto3" json:"DirectoryList,omitempty"`
	PortList             []string `protobuf:"bytes,3,rep,name=portList,proto3" json:"portList,omitempty"`
	Locale               *Locale  `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	GpVersion            string   `protobuf:"bytes,5,opt,name=gpVersion,proto3" json:"gpVersion,omitempty"`
	Forced               bool     `protobuf:"varint,6,opt,name=forced,proto3" json:"forced,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateHostEnvRequest) Reset()         { *m = ValidateHostEnvRequest{} }
func (m *ValidateHostEnvRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateHostEnvRequest) ProtoMessage()    {}
func (*ValidateHostEnvRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{6}
}

func (m *ValidateHostEnvRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateHostEnvRequest.Unmarshal(m, b)
}
func (m *ValidateHostEnvRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateHostEnvRequest.Marshal(b, m, deterministic)
}
func (m *ValidateHostEnvRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateHostEnvRequest.Merge(m, src)
}
func (m *ValidateHostEnvRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateHostEnvRequest.Size(m)
}
func (m *ValidateHostEnvRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateHostEnvRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateHostEnvRequest proto.InternalMessageInfo

func (m *ValidateHostEnvRequest) GetHostAddressList() []string {
	if m != nil {
		return m.HostAddressList
	}
	return nil
}

func (m *ValidateHostEnvRequest) GetDirectoryList() []string {
	if m != nil {
		return m.DirectoryList
	}
	return nil
}

func (m *ValidateHostEnvRequest) GetPortList() []string {
	if m != nil {
		return m.PortList
	}
	return nil
}

func (m *ValidateHostEnvRequest) GetLocale() *Locale {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *ValidateHostEnvRequest) GetGpVersion() string {
	if m != nil {
		return m.GpVersion
	}
	return ""
}

func (m *ValidateHostEnvRequest) GetForced() bool {
	if m != nil {
		return m.Forced
	}
	return false
}

type ValidateHostEnvReply struct {
	Messages             []*LogMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ValidateHostEnvReply) Reset()         { *m = ValidateHostEnvReply{} }
func (m *ValidateHostEnvReply) String() string { return proto.CompactTextString(m) }
func (*ValidateHostEnvReply) ProtoMessage()    {}
func (*ValidateHostEnvReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{7}
}

func (m *ValidateHostEnvReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateHostEnvReply.Unmarshal(m, b)
}
func (m *ValidateHostEnvReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateHostEnvReply.Marshal(b, m, deterministic)
}
func (m *ValidateHostEnvReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateHostEnvReply.Merge(m, src)
}
func (m *ValidateHostEnvReply) XXX_Size() int {
	return xxx_messageInfo_ValidateHostEnvReply.Size(m)
}
func (m *ValidateHostEnvReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateHostEnvReply.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateHostEnvReply proto.InternalMessageInfo

func (m *ValidateHostEnvReply) GetMessages() []*LogMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type MakeSegmentRequest struct {
	Segment              *Segment          `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	Locale               *Locale           `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Encoding             string            `protobuf:"bytes,3,opt,name=Encoding,proto3" json:"Encoding,omitempty"`
	SegConfig            map[string]string `protobuf:"bytes,4,rep,name=SegConfig,proto3" json:"SegConfig,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CoordinatorAddrs     []string          `protobuf:"bytes,5,rep,name=coordinatorAddrs,proto3" json:"coordinatorAddrs,omitempty"`
	HbaHostNames         bool              `protobuf:"varint,6,opt,name=hbaHostNames,proto3" json:"hbaHostNames,omitempty"`
	DataChecksums        bool              `protobuf:"varint,7,opt,name=dataChecksums,proto3" json:"dataChecksums,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MakeSegmentRequest) Reset()         { *m = MakeSegmentRequest{} }
func (m *MakeSegmentRequest) String() string { return proto.CompactTextString(m) }
func (*MakeSegmentRequest) ProtoMessage()    {}
func (*MakeSegmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{8}
}

func (m *MakeSegmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeSegmentRequest.Unmarshal(m, b)
}
func (m *MakeSegmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeSegmentRequest.Marshal(b, m, deterministic)
}
func (m *MakeSegmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeSegmentRequest.Merge(m, src)
}
func (m *MakeSegmentRequest) XXX_Size() int {
	return xxx_messageInfo_MakeSegmentRequest.Size(m)
}
func (m *MakeSegmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeSegmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeSegmentRequest proto.InternalMessageInfo

func (m *MakeSegmentRequest) GetSegment() *Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MakeSegmentRequest) GetLocale() *Locale {
	if m != nil {
		return m.Locale
	}
	return nil
}

func (m *MakeSegmentRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *MakeSegmentRequest) GetSegConfig() map[string]string {
	if m != nil {
		return m.SegConfig
	}
	return nil
}

func (m *MakeSegmentRequest) GetCoordinatorAddrs() []string {
	if m != nil {
		return m.CoordinatorAddrs
	}
	return nil
}

func (m *MakeSegmentRequest) GetHbaHostNames() bool {
	if m != nil {
		return m.HbaHostNames
	}
	return false
}

func (m *MakeSegmentRequest) GetDataChecksums() bool {
	if m != nil {
		return m.DataChecksums
	}
	return false
}

type MakeSegmentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeSegmentReply) Reset()         { *m = MakeSegmentReply{} }
func (m *MakeSegmentReply) String() string { return proto.CompactTextString(m) }
func (*MakeSegmentReply) ProtoMessage()    {}
func (*MakeSegmentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{9}
}

func (m *MakeSegmentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeSegmentReply.Unmarshal(m, b)
}
func (m *MakeSegmentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeSegmentReply.Marshal(b, m, deterministic)
}
func (m *MakeSegmentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeSegmentReply.Merge(m, src)
}
func (m *MakeSegmentReply) XXX_Size() int {
	return xxx_messageInfo_MakeSegmentReply.Size(m)
}
func (m *MakeSegmentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeSegmentReply.DiscardUnknown(m)
}

var xxx_messageInfo_MakeSegmentReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartSegmentRequest)(nil), "idl.StartSegmentRequest")
	proto.RegisterType((*StartSegmentReply)(nil), "idl.StartSegmentReply")
	proto.RegisterType((*StopAgentRequest)(nil), "idl.StopAgentRequest")
	proto.RegisterType((*StopAgentReply)(nil), "idl.StopAgentReply")
	proto.RegisterType((*StatusAgentRequest)(nil), "idl.StatusAgentRequest")
	proto.RegisterType((*StatusAgentReply)(nil), "idl.StatusAgentReply")
	proto.RegisterType((*ValidateHostEnvRequest)(nil), "idl.ValidateHostEnvRequest")
	proto.RegisterType((*ValidateHostEnvReply)(nil), "idl.ValidateHostEnvReply")
	proto.RegisterType((*MakeSegmentRequest)(nil), "idl.MakeSegmentRequest")
	proto.RegisterMapType((map[string]string)(nil), "idl.MakeSegmentRequest.SegConfigEntry")
	proto.RegisterType((*MakeSegmentReply)(nil), "idl.MakeSegmentReply")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x7a, 0x5b, 0x73, 0xba, 0x4b, 0x39, 0xbb, 0x10, 0x02, 0x0f, 0x55, 0x40, 0x53, 0x05,
	0x52, 0x91, 0x0a, 0x0f, 0x68, 0x42, 0x42, 0xbb, 0x49, 0x48, 0x6c, 0x3c, 0xa4, 0x68, 0x0f, 0xbc,
	0x79, 0x8d, 0x97, 0x5a, 0x4b, 0xe3, 0x60, 0x3b, 0x9b, 0xfa, 0x1b, 0xf9, 0x11, 0xfc, 0x0c, 0x5e,
	0x91, 0x1d, 0xa7, 0x6b, 0xd2, 0xf1, 0xe6, 0xef, 0x3b, 0x9f, 0x8f, 0xcf, 0xd5, 0xd0, 0x23, 0x31,
	0x4d, 0xd5, 0x28, 0x13, 0x5c, 0x71, 0x6c, 0xb2, 0x28, 0xf1, 0xdd, 0x59, 0x7e, 0x53, 0xe0, 0xe0,
	0x01, 0xf6, 0x26, 0x8a, 0x08, 0x35, 0xa1, 0xf1, 0x9c, 0xa6, 0x2a, 0xa4, 0xbf, 0x72, 0x2a, 0x15,
	0x7a, 0xb0, 0x19, 0x11, 0x45, 0xce, 0x99, 0xf0, 0x9c, 0x81, 0x33, 0x74, 0xc3, 0x12, 0x22, 0x42,
	0xeb, 0x81, 0x30, 0xe5, 0x35, 0x06, 0xce, 0xb0, 0x1b, 0x9a, 0xb3, 0x56, 0x2b, 0x36, 0xa7, 0x3c,
	0x57, 0x5e, 0x6b, 0xe0, 0x0c, 0xdb, 0x61, 0x09, 0xb5, 0x85, 0x67, 0x8a, 0xf1, 0x54, 0x7a, 0xed,
	0xc2, 0x8f, 0x85, 0xc1, 0x1e, 0x3c, 0xab, 0x3e, 0x9c, 0x25, 0x8b, 0x00, 0xa1, 0x3f, 0x51, 0x3c,
	0x3b, 0x89, 0x1f, 0x43, 0x09, 0xfa, 0xb0, 0xb3, 0xc2, 0x69, 0xd5, 0x3e, 0xe0, 0x44, 0x11, 0x95,
	0xcb, 0x8a, 0xee, 0x87, 0xbe, 0xbb, 0xc2, 0x66, 0xc9, 0x02, 0x0f, 0xa1, 0x23, 0x0d, 0x67, 0xb3,
	0xb0, 0x48, 0xf3, 0x79, 0xa6, 0x63, 0x34, 0x69, 0xb8, 0xa1, 0x45, 0xd8, 0x87, 0x66, 0xc6, 0x22,
	0xaf, 0x39, 0x70, 0x86, 0xdb, 0xa1, 0x3e, 0x06, 0x7f, 0x1c, 0x38, 0xbc, 0x26, 0x09, 0x8b, 0x88,
	0xa2, 0x5f, 0xb9, 0x54, 0x17, 0xe9, 0x7d, 0x59, 0xa3, 0x21, 0xec, 0xce, 0xb8, 0x54, 0x27, 0x51,
	0x24, 0xa8, 0x94, 0x97, 0x4c, 0x2a, 0xcf, 0x19, 0x34, 0x87, 0x6e, 0x58, 0xa7, 0xf1, 0x0d, 0x6c,
	0x9f, 0x33, 0x41, 0xa7, 0x8a, 0x8b, 0x85, 0xd1, 0x35, 0x8c, 0xae, 0x4a, 0xa2, 0x0f, 0xdd, 0x8c,
	0x0b, 0x65, 0x04, 0x4d, 0x23, 0x58, 0x62, 0x7c, 0x0d, 0x9d, 0x84, 0x4f, 0x49, 0x42, 0x4d, 0x81,
	0x7b, 0xe3, 0xde, 0x88, 0x45, 0xc9, 0xe8, 0xd2, 0x50, 0xa1, 0x35, 0xe1, 0x2b, 0x70, 0xe3, 0xec,
	0x9a, 0x0a, 0xc9, 0x78, 0x6a, 0xcb, 0xfd, 0x48, 0xe8, 0x9c, 0x6f, 0xb9, 0x98, 0xd2, 0xc8, 0xeb,
	0x98, 0xd6, 0x59, 0x14, 0x9c, 0xc1, 0xfe, 0x5a, 0x82, 0xba, 0x76, 0xef, 0xa0, 0x3b, 0xa7, 0x52,
	0x92, 0x98, 0x4a, 0x93, 0x57, 0x6f, 0xbc, 0x6b, 0x1f, 0x8d, 0xaf, 0x0a, 0x3e, 0x5c, 0x0a, 0x82,
	0xbf, 0x0d, 0xc0, 0x2b, 0x72, 0x47, 0x6b, 0x63, 0x74, 0x04, 0x9b, 0xb2, 0x60, 0x4c, 0x03, 0x7a,
	0xe3, 0x2d, 0xe3, 0xa2, 0x54, 0x95, 0xc6, 0x95, 0xf4, 0x1a, 0xff, 0x4f, 0xcf, 0x87, 0xee, 0x45,
	0x3a, 0xe5, 0x11, 0x4b, 0x63, 0xd3, 0x21, 0x37, 0x5c, 0x62, 0x3c, 0x07, 0x77, 0x42, 0xe3, 0x33,
	0x9e, 0xde, 0xb2, 0xd8, 0x6b, 0x99, 0x68, 0x8f, 0x8c, 0x8f, 0xf5, 0xa0, 0x46, 0x4b, 0xe1, 0x45,
	0xaa, 0xc4, 0x22, 0x7c, 0xbc, 0x88, 0x6f, 0xa1, 0x3f, 0xe5, 0x5c, 0x44, 0x2c, 0x25, 0x8a, 0x0b,
	0xdd, 0x41, 0x3d, 0xb6, 0xba, 0x13, 0x6b, 0x3c, 0x06, 0xb0, 0x35, 0xbb, 0x21, 0xba, 0x62, 0xdf,
	0xc9, 0x9c, 0x4a, 0x5b, 0xd4, 0x0a, 0xa7, 0xfb, 0xae, 0xd7, 0xe6, 0x6c, 0x46, 0xa7, 0x77, 0x32,
	0x9f, 0x4b, 0x6f, 0xd3, 0x88, 0xaa, 0xa4, 0xff, 0x19, 0x76, 0xaa, 0x21, 0xe9, 0x31, 0xbc, 0xa3,
	0x0b, 0x3b, 0xb3, 0xfa, 0x88, 0xfb, 0xd0, 0xbe, 0x27, 0x49, 0x5e, 0xce, 0x6b, 0x01, 0x8e, 0x1b,
	0x9f, 0x1c, 0xbd, 0x32, 0x95, 0x1c, 0xb3, 0x64, 0x31, 0xfe, 0xdd, 0x80, 0xb6, 0xd9, 0x02, 0xfc,
	0x08, 0x2d, 0xbd, 0x3c, 0x78, 0x50, 0xd4, 0xbd, 0xb6, 0x5b, 0xfe, 0x5e, 0x9d, 0xd6, 0xeb, 0xb5,
	0x81, 0xc7, 0xd0, 0x29, 0x56, 0x09, 0x9f, 0x5b, 0x41, 0x7d, 0xdb, 0xfc, 0x83, 0x75, 0x43, 0x71,
	0xf7, 0x0b, 0xf4, 0x56, 0xe2, 0xb1, 0x0e, 0xd6, 0xbb, 0x60, 0x1d, 0xd4, 0x43, 0x0f, 0x36, 0xf0,
	0x14, 0xb6, 0x56, 0x3f, 0x06, 0xf4, 0xca, 0x97, 0xea, 0x9f, 0x94, 0x7f, 0xf8, 0x84, 0xa5, 0xf0,
	0xf1, 0x0d, 0x76, 0x6b, 0x33, 0x8d, 0x2f, 0x8d, 0xf8, 0xe9, 0x55, 0xf6, 0x5f, 0x3c, 0x6d, 0x34,
	0xce, 0x4e, 0xbb, 0x3f, 0x3b, 0xa3, 0xd1, 0x7b, 0x16, 0x25, 0x37, 0x1d, 0xf3, 0x67, 0x7e, 0xf8,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x72, 0xd2, 0xf0, 0x06, 0x52, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	Stop(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error)
	Status(ctx context.Context, in *StatusAgentRequest, opts ...grpc.CallOption) (*StatusAgentReply, error)
	MakeSegment(ctx context.Context, in *MakeSegmentRequest, opts ...grpc.CallOption) (*MakeSegmentReply, error)
	StartSegment(ctx context.Context, in *StartSegmentRequest, opts ...grpc.CallOption) (*StartSegmentReply, error)
	ValidateHostEnv(ctx context.Context, in *ValidateHostEnvRequest, opts ...grpc.CallOption) (*ValidateHostEnvReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Stop(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error) {
	out := new(StopAgentReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Status(ctx context.Context, in *StatusAgentRequest, opts ...grpc.CallOption) (*StatusAgentReply, error) {
	out := new(StatusAgentReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) MakeSegment(ctx context.Context, in *MakeSegmentRequest, opts ...grpc.CallOption) (*MakeSegmentReply, error) {
	out := new(MakeSegmentReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/MakeSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StartSegment(ctx context.Context, in *StartSegmentRequest, opts ...grpc.CallOption) (*StartSegmentReply, error) {
	out := new(StartSegmentReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/StartSegment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ValidateHostEnv(ctx context.Context, in *ValidateHostEnvRequest, opts ...grpc.CallOption) (*ValidateHostEnvReply, error) {
	out := new(ValidateHostEnvReply)
	err := c.cc.Invoke(ctx, "/idl.Agent/ValidateHostEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	Stop(context.Context, *StopAgentRequest) (*StopAgentReply, error)
	Status(context.Context, *StatusAgentRequest) (*StatusAgentReply, error)
	MakeSegment(context.Context, *MakeSegmentRequest) (*MakeSegmentReply, error)
	StartSegment(context.Context, *StartSegmentRequest) (*StartSegmentReply, error)
	ValidateHostEnv(context.Context, *ValidateHostEnvRequest) (*ValidateHostEnvReply, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) Stop(ctx context.Context, req *StopAgentRequest) (*StopAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedAgentServer) Status(ctx context.Context, req *StatusAgentRequest) (*StatusAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedAgentServer) MakeSegment(ctx context.Context, req *MakeSegmentRequest) (*MakeSegmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeSegment not implemented")
}
func (*UnimplementedAgentServer) StartSegment(ctx context.Context, req *StartSegmentRequest) (*StartSegmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSegment not implemented")
}
func (*UnimplementedAgentServer) ValidateHostEnv(ctx context.Context, req *ValidateHostEnvRequest) (*ValidateHostEnvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateHostEnv not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Stop(ctx, req.(*StopAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Status(ctx, req.(*StatusAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_MakeSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).MakeSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/MakeSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).MakeSegment(ctx, req.(*MakeSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StartSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StartSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/StartSegment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StartSegment(ctx, req.(*StartSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ValidateHostEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateHostEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ValidateHostEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/ValidateHostEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ValidateHostEnv(ctx, req.(*ValidateHostEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _Agent_Stop_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Agent_Status_Handler,
		},
		{
			MethodName: "MakeSegment",
			Handler:    _Agent_MakeSegment_Handler,
		},
		{
			MethodName: "StartSegment",
			Handler:    _Agent_StartSegment_Handler,
		},
		{
			MethodName: "ValidateHostEnv",
			Handler:    _Agent_ValidateHostEnv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}