// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/response/article.proto

package response // import "JupiterBlog/micro/protobuf/response"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/protobuf/types/known/timestamppb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ArticleInfo struct {
	Id                   uint32   `protobuf:"fixed32,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CateId               uint32   `protobuf:"fixed32,4,opt,name=cate_id,json=cateId,proto3" json:"cate_id,omitempty"`
	ReadAmount           uint64   `protobuf:"fixed64,5,opt,name=read_amount,json=readAmount,proto3" json:"read_amount,omitempty"`
	Sort                 int32    `protobuf:"fixed32,6,opt,name=sort,proto3" json:"sort,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Status               uint32   `protobuf:"fixed32,10,opt,name=status,proto3" json:"status,omitempty"`
	Image                string   `protobuf:"bytes,11,opt,name=image,proto3" json:"image,omitempty"`
	Summary              string   `protobuf:"bytes,12,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleInfo) Reset()         { *m = ArticleInfo{} }
func (m *ArticleInfo) String() string { return proto.CompactTextString(m) }
func (*ArticleInfo) ProtoMessage()    {}
func (*ArticleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_article_a1b947462f49a19d, []int{0}
}
func (m *ArticleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleInfo.Unmarshal(m, b)
}
func (m *ArticleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleInfo.Marshal(b, m, deterministic)
}
func (dst *ArticleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleInfo.Merge(dst, src)
}
func (m *ArticleInfo) XXX_Size() int {
	return xxx_messageInfo_ArticleInfo.Size(m)
}
func (m *ArticleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleInfo proto.InternalMessageInfo

func (m *ArticleInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ArticleInfo) GetCateId() uint32 {
	if m != nil {
		return m.CateId
	}
	return 0
}

func (m *ArticleInfo) GetReadAmount() uint64 {
	if m != nil {
		return m.ReadAmount
	}
	return 0
}

func (m *ArticleInfo) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

func (m *ArticleInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ArticleInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ArticleInfo) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *ArticleInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ArticleInfo) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ArticleInfo) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func init() {
	proto.RegisterType((*ArticleInfo)(nil), "response.ArticleInfo")
}

func init() {
	proto.RegisterFile("protobuf/response/article.proto", fileDescriptor_article_a1b947462f49a19d)
}

var fileDescriptor_article_a1b947462f49a19d = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xec, 0x30,
	0x14, 0x86, 0x69, 0xef, 0x4c, 0x3b, 0x73, 0xe6, 0xa2, 0x10, 0x44, 0x83, 0x20, 0x53, 0x14, 0xa1,
	0x2b, 0xbb, 0xf0, 0x09, 0x3a, 0xbb, 0x71, 0xd9, 0xa5, 0x9b, 0x92, 0x69, 0xce, 0x94, 0x40, 0xd3,
	0x94, 0xe4, 0x64, 0xe1, 0xdb, 0xf9, 0x68, 0xd2, 0xa4, 0x75, 0xe3, 0x2e, 0xff, 0xf7, 0xe5, 0x24,
	0xf9, 0x09, 0x1c, 0x27, 0x6b, 0xc8, 0x5c, 0xfc, 0xb5, 0xb2, 0xe8, 0x26, 0x33, 0x3a, 0xac, 0x84,
	0x25, 0xd5, 0x0d, 0xf8, 0x16, 0x0c, 0xdb, 0xad, 0xfc, 0xf1, 0xd8, 0x1b, 0xd3, 0x0f, 0x58, 0xfd,
	0x4e, 0x90, 0xd2, 0xe8, 0x48, 0xe8, 0x29, 0x6e, 0x7d, 0xfe, 0x4e, 0xe1, 0x50, 0xc7, 0xe1, 0xf3,
	0x78, 0x35, 0xec, 0x06, 0x52, 0x25, 0x79, 0x52, 0x24, 0x65, 0xde, 0xa4, 0x4a, 0xb2, 0x3b, 0xd8,
	0x92, 0xa2, 0x01, 0x79, 0x5a, 0x24, 0xe5, 0xbe, 0x89, 0x81, 0x71, 0xc8, 0x3b, 0x33, 0x12, 0x8e,
	0xc4, 0xff, 0x05, 0xbe, 0x46, 0xf6, 0x00, 0x79, 0x27, 0x08, 0x5b, 0x25, 0xf9, 0x26, 0x1c, 0x92,
	0xcd, 0xf1, 0x2c, 0xd9, 0x11, 0x0e, 0x16, 0x85, 0x6c, 0x85, 0x36, 0x7e, 0x24, 0xbe, 0x2d, 0x92,
	0x32, 0x6b, 0x60, 0x46, 0x75, 0x20, 0x8c, 0xc1, 0xc6, 0x19, 0x4b, 0x3c, 0x2b, 0x92, 0xf2, 0xb6,
	0x09, 0x6b, 0xf6, 0x04, 0xd0, 0x59, 0x14, 0x84, 0xb2, 0x15, 0xc4, 0xf3, 0x70, 0xd5, 0x7e, 0x21,
	0x75, 0xd0, 0x7e, 0x92, 0xab, 0xde, 0x45, 0xbd, 0x90, 0xa8, 0x25, 0x0e, 0xb8, 0xe8, 0x7d, 0xd4,
	0x0b, 0xa9, 0x89, 0xdd, 0x43, 0xe6, 0x48, 0x90, 0x77, 0x1c, 0xe2, 0x4b, 0x63, 0x9a, 0x2b, 0x2b,
	0x2d, 0x7a, 0xe4, 0x87, 0x58, 0x39, 0x84, 0xb9, 0xb2, 0xf3, 0x5a, 0x0b, 0xfb, 0xc5, 0xff, 0xc7,
	0xca, 0x4b, 0x3c, 0xbd, 0x7e, 0xbe, 0x7c, 0xf8, 0x49, 0x11, 0xda, 0xd3, 0x60, 0xfa, 0x4a, 0xab,
	0xce, 0x9a, 0xea, 0xcf, 0x17, 0x5d, 0xb2, 0x80, 0xde, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0x53, 0x7c, 0x29, 0xbe, 0x01, 0x00, 0x00,
}