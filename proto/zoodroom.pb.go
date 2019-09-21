// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zoodroom.proto

package zoodroom

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AddFeedbackRequest struct {
	FullName             string   `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackRequest) Reset()         { *m = AddFeedbackRequest{} }
func (m *AddFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackRequest) ProtoMessage()    {}
func (*AddFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_282c5c41d3f9130f, []int{0}
}

func (m *AddFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackRequest.Unmarshal(m, b)
}
func (m *AddFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackRequest.Merge(m, src)
}
func (m *AddFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackRequest.Size(m)
}
func (m *AddFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackRequest proto.InternalMessageInfo

func (m *AddFeedbackRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *AddFeedbackRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddFeedbackRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AddFeedbackRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type AddFeedbackResponse struct {
	IsCreated            bool     `protobuf:"varint,1,opt,name=is_created,json=isCreated,proto3" json:"is_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFeedbackResponse) Reset()         { *m = AddFeedbackResponse{} }
func (m *AddFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddFeedbackResponse) ProtoMessage()    {}
func (*AddFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_282c5c41d3f9130f, []int{1}
}

func (m *AddFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFeedbackResponse.Unmarshal(m, b)
}
func (m *AddFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFeedbackResponse.Merge(m, src)
}
func (m *AddFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddFeedbackResponse.Size(m)
}
func (m *AddFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFeedbackResponse proto.InternalMessageInfo

func (m *AddFeedbackResponse) GetIsCreated() bool {
	if m != nil {
		return m.IsCreated
	}
	return false
}

type GetFeedbackRequest struct {
	FeebackId            string   `protobuf:"bytes,1,opt,name=feeback_id,json=feebackId,proto3" json:"feeback_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackRequest) Reset()         { *m = GetFeedbackRequest{} }
func (m *GetFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackRequest) ProtoMessage()    {}
func (*GetFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_282c5c41d3f9130f, []int{2}
}

func (m *GetFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackRequest.Unmarshal(m, b)
}
func (m *GetFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackRequest.Merge(m, src)
}
func (m *GetFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackRequest.Size(m)
}
func (m *GetFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackRequest proto.InternalMessageInfo

func (m *GetFeedbackRequest) GetFeebackId() string {
	if m != nil {
		return m.FeebackId
	}
	return ""
}

type GetFeedbackResponse struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedbackResponse) Reset()         { *m = GetFeedbackResponse{} }
func (m *GetFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeedbackResponse) ProtoMessage()    {}
func (*GetFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_282c5c41d3f9130f, []int{3}
}

func (m *GetFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedbackResponse.Unmarshal(m, b)
}
func (m *GetFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *GetFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedbackResponse.Merge(m, src)
}
func (m *GetFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeedbackResponse.Size(m)
}
func (m *GetFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedbackResponse proto.InternalMessageInfo

func (m *GetFeedbackResponse) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *GetFeedbackResponse) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *GetFeedbackResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetFeedbackResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *GetFeedbackResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*AddFeedbackRequest)(nil), "zoodroom.AddFeedbackRequest")
	proto.RegisterType((*AddFeedbackResponse)(nil), "zoodroom.AddFeedbackResponse")
	proto.RegisterType((*GetFeedbackRequest)(nil), "zoodroom.GetFeedbackRequest")
	proto.RegisterType((*GetFeedbackResponse)(nil), "zoodroom.GetFeedbackResponse")
}

func init() { proto.RegisterFile("zoodroom.proto", fileDescriptor_282c5c41d3f9130f) }

var fileDescriptor_282c5c41d3f9130f = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x0b, 0xb4, 0x0a, 0x63, 0xd4, 0x64, 0xeb, 0x81, 0xa8, 0x4d, 0xcc, 0x9e, 0x3c, 0xf5,
	0x60, 0x7d, 0x01, 0x63, 0xa2, 0x69, 0x62, 0x3c, 0xa0, 0x77, 0x42, 0xd9, 0x69, 0xdc, 0x08, 0x3b,
	0xc8, 0x6e, 0x35, 0xf1, 0x0d, 0x7c, 0x16, 0x5f, 0xd2, 0xb2, 0x50, 0x05, 0x5b, 0x8e, 0xdf, 0x37,
	0x3b, 0x7f, 0x7e, 0x33, 0x0b, 0x47, 0x9f, 0x44, 0xa2, 0x24, 0xca, 0xa7, 0x45, 0x49, 0x86, 0x98,
	0xbf, 0xd1, 0xfc, 0x03, 0xd8, 0x8d, 0x10, 0x77, 0x88, 0x62, 0x91, 0xa4, 0xaf, 0x11, 0xbe, 0xad,
	0x50, 0x1b, 0x76, 0x06, 0xc1, 0x72, 0x95, 0x65, 0xb1, 0x4a, 0x72, 0x0c, 0x9d, 0x0b, 0xe7, 0x32,
	0x88, 0xfc, 0xca, 0x78, 0x5c, 0x6b, 0x76, 0x02, 0x23, 0xcc, 0x13, 0x99, 0x85, 0xae, 0x0d, 0xd4,
	0xa2, 0x72, 0x8b, 0x17, 0x52, 0x18, 0x7a, 0xb5, 0x6b, 0x05, 0x0b, 0x61, 0x3f, 0x25, 0x65, 0x50,
	0x99, 0x70, 0x68, 0xfd, 0x8d, 0xe4, 0xd7, 0x30, 0xee, 0x34, 0xd6, 0x05, 0x29, 0x8d, 0x6c, 0x02,
	0x20, 0x75, 0x9c, 0x96, 0x98, 0x18, 0x14, 0xb6, 0xb5, 0x1f, 0x05, 0x52, 0xdf, 0xd6, 0x06, 0x9f,
	0x01, 0xbb, 0x47, 0xf3, 0x7f, 0xdc, 0x75, 0xd2, 0x12, 0xb1, 0x72, 0x62, 0x29, 0x9a, 0x79, 0x83,
	0xc6, 0x99, 0x0b, 0xfe, 0xe5, 0xc0, 0xb8, 0x93, 0xd5, 0xf4, 0x3a, 0x06, 0xef, 0xef, 0xbd, 0x3b,
	0x17, 0x5d, 0x6c, 0xb7, 0x0f, 0xdb, 0xdb, 0x89, 0x3d, 0xec, 0xc1, 0x1e, 0x75, 0xb0, 0xaf, 0xbe,
	0x1d, 0x38, 0x7c, 0xa6, 0x3c, 0x31, 0xf4, 0x84, 0xe5, 0xbb, 0x4c, 0x91, 0x3d, 0xc0, 0x41, 0x6b,
	0x11, 0xec, 0x7c, 0xfa, 0x7b, 0xab, 0xed, 0xc3, 0x9c, 0x4e, 0x7a, 0xa2, 0x35, 0x11, 0x1f, 0x54,
	0xd5, 0x5a, 0xa8, 0xed, 0x6a, 0xdb, 0x7b, 0x6b, 0x57, 0xdb, 0xb1, 0x1f, 0x3e, 0x58, 0xec, 0xd9,
	0xef, 0x32, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x97, 0xef, 0x89, 0x4a, 0x40, 0x02, 0x00, 0x00,
}