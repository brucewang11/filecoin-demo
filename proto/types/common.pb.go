// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/common.proto

package types

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

type AddPieceReq struct {
	Ip                   []byte   `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPieceReq) Reset()         { *m = AddPieceReq{} }
func (m *AddPieceReq) String() string { return proto.CompactTextString(m) }
func (*AddPieceReq) ProtoMessage()    {}
func (*AddPieceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{0}
}

func (m *AddPieceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPieceReq.Unmarshal(m, b)
}
func (m *AddPieceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPieceReq.Marshal(b, m, deterministic)
}
func (m *AddPieceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPieceReq.Merge(m, src)
}
func (m *AddPieceReq) XXX_Size() int {
	return xxx_messageInfo_AddPieceReq.Size(m)
}
func (m *AddPieceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPieceReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPieceReq proto.InternalMessageInfo

func (m *AddPieceReq) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

type AddPieceRes struct {
	SectorId             uint64   `protobuf:"varint,2,opt,name=SectorId,proto3" json:"SectorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPieceRes) Reset()         { *m = AddPieceRes{} }
func (m *AddPieceRes) String() string { return proto.CompactTextString(m) }
func (*AddPieceRes) ProtoMessage()    {}
func (*AddPieceRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{1}
}

func (m *AddPieceRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPieceRes.Unmarshal(m, b)
}
func (m *AddPieceRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPieceRes.Marshal(b, m, deterministic)
}
func (m *AddPieceRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPieceRes.Merge(m, src)
}
func (m *AddPieceRes) XXX_Size() int {
	return xxx_messageInfo_AddPieceRes.Size(m)
}
func (m *AddPieceRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPieceRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddPieceRes proto.InternalMessageInfo

func (m *AddPieceRes) GetSectorId() uint64 {
	if m != nil {
		return m.SectorId
	}
	return 0
}

type GenPoStReq struct {
	Ip                   []byte   `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Seed                 []byte   `protobuf:"bytes,2,opt,name=Seed,proto3" json:"Seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenPoStReq) Reset()         { *m = GenPoStReq{} }
func (m *GenPoStReq) String() string { return proto.CompactTextString(m) }
func (*GenPoStReq) ProtoMessage()    {}
func (*GenPoStReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{2}
}

func (m *GenPoStReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenPoStReq.Unmarshal(m, b)
}
func (m *GenPoStReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenPoStReq.Marshal(b, m, deterministic)
}
func (m *GenPoStReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenPoStReq.Merge(m, src)
}
func (m *GenPoStReq) XXX_Size() int {
	return xxx_messageInfo_GenPoStReq.Size(m)
}
func (m *GenPoStReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GenPoStReq.DiscardUnknown(m)
}

var xxx_messageInfo_GenPoStReq proto.InternalMessageInfo

func (m *GenPoStReq) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *GenPoStReq) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

type GenPoStRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenPoStRes) Reset()         { *m = GenPoStRes{} }
func (m *GenPoStRes) String() string { return proto.CompactTextString(m) }
func (*GenPoStRes) ProtoMessage()    {}
func (*GenPoStRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{3}
}

func (m *GenPoStRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenPoStRes.Unmarshal(m, b)
}
func (m *GenPoStRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenPoStRes.Marshal(b, m, deterministic)
}
func (m *GenPoStRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenPoStRes.Merge(m, src)
}
func (m *GenPoStRes) XXX_Size() int {
	return xxx_messageInfo_GenPoStRes.Size(m)
}
func (m *GenPoStRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GenPoStRes.DiscardUnknown(m)
}

var xxx_messageInfo_GenPoStRes proto.InternalMessageInfo

type SealResultReq struct {
	SectorID             uint64       `protobuf:"varint,1,opt,name=SectorID,proto3" json:"SectorID,omitempty"`
	SealStatusCode       uint32       `protobuf:"varint,2,opt,name=SealStatusCode,proto3" json:"SealStatusCode,omitempty"`
	SealErrorMsg         string       `protobuf:"bytes,3,opt,name=SealErrorMsg,proto3" json:"SealErrorMsg,omitempty"`
	CommD                []byte       `protobuf:"bytes,4,opt,name=CommD,proto3" json:"CommD,omitempty"`
	CommR                []byte       `protobuf:"bytes,5,opt,name=CommR,proto3" json:"CommR,omitempty"`
	CommRStar            []byte       `protobuf:"bytes,6,opt,name=CommRStar,proto3" json:"CommRStar,omitempty"`
	Proof                []byte       `protobuf:"bytes,7,opt,name=Proof,proto3" json:"Proof,omitempty"`
	Pieces               []*PieceMeta `protobuf:"bytes,8,rep,name=Pieces,proto3" json:"Pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SealResultReq) Reset()         { *m = SealResultReq{} }
func (m *SealResultReq) String() string { return proto.CompactTextString(m) }
func (*SealResultReq) ProtoMessage()    {}
func (*SealResultReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{4}
}

func (m *SealResultReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealResultReq.Unmarshal(m, b)
}
func (m *SealResultReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealResultReq.Marshal(b, m, deterministic)
}
func (m *SealResultReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealResultReq.Merge(m, src)
}
func (m *SealResultReq) XXX_Size() int {
	return xxx_messageInfo_SealResultReq.Size(m)
}
func (m *SealResultReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SealResultReq.DiscardUnknown(m)
}

var xxx_messageInfo_SealResultReq proto.InternalMessageInfo

func (m *SealResultReq) GetSectorID() uint64 {
	if m != nil {
		return m.SectorID
	}
	return 0
}

func (m *SealResultReq) GetSealStatusCode() uint32 {
	if m != nil {
		return m.SealStatusCode
	}
	return 0
}

func (m *SealResultReq) GetSealErrorMsg() string {
	if m != nil {
		return m.SealErrorMsg
	}
	return ""
}

func (m *SealResultReq) GetCommD() []byte {
	if m != nil {
		return m.CommD
	}
	return nil
}

func (m *SealResultReq) GetCommR() []byte {
	if m != nil {
		return m.CommR
	}
	return nil
}

func (m *SealResultReq) GetCommRStar() []byte {
	if m != nil {
		return m.CommRStar
	}
	return nil
}

func (m *SealResultReq) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *SealResultReq) GetPieces() []*PieceMeta {
	if m != nil {
		return m.Pieces
	}
	return nil
}

type PieceMeta struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Size                 uint64   `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	InclusionProof       []byte   `protobuf:"bytes,3,opt,name=InclusionProof,proto3" json:"InclusionProof,omitempty"`
	CommP                []byte   `protobuf:"bytes,4,opt,name=CommP,proto3" json:"CommP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceMeta) Reset()         { *m = PieceMeta{} }
func (m *PieceMeta) String() string { return proto.CompactTextString(m) }
func (*PieceMeta) ProtoMessage()    {}
func (*PieceMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{5}
}

func (m *PieceMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceMeta.Unmarshal(m, b)
}
func (m *PieceMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceMeta.Marshal(b, m, deterministic)
}
func (m *PieceMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceMeta.Merge(m, src)
}
func (m *PieceMeta) XXX_Size() int {
	return xxx_messageInfo_PieceMeta.Size(m)
}
func (m *PieceMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceMeta.DiscardUnknown(m)
}

var xxx_messageInfo_PieceMeta proto.InternalMessageInfo

func (m *PieceMeta) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PieceMeta) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceMeta) GetInclusionProof() []byte {
	if m != nil {
		return m.InclusionProof
	}
	return nil
}

func (m *PieceMeta) GetCommP() []byte {
	if m != nil {
		return m.CommP
	}
	return nil
}

type SealResultRes struct {
	Got                  bool     `protobuf:"varint,1,opt,name=Got,proto3" json:"Got,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SealResultRes) Reset()         { *m = SealResultRes{} }
func (m *SealResultRes) String() string { return proto.CompactTextString(m) }
func (*SealResultRes) ProtoMessage()    {}
func (*SealResultRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{6}
}

func (m *SealResultRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealResultRes.Unmarshal(m, b)
}
func (m *SealResultRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealResultRes.Marshal(b, m, deterministic)
}
func (m *SealResultRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealResultRes.Merge(m, src)
}
func (m *SealResultRes) XXX_Size() int {
	return xxx_messageInfo_SealResultRes.Size(m)
}
func (m *SealResultRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SealResultRes.DiscardUnknown(m)
}

var xxx_messageInfo_SealResultRes proto.InternalMessageInfo

func (m *SealResultRes) GetGot() bool {
	if m != nil {
		return m.Got
	}
	return false
}

type PoStResultReq struct {
	Proof                []byte   `protobuf:"bytes,1,opt,name=Proof,proto3" json:"Proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoStResultReq) Reset()         { *m = PoStResultReq{} }
func (m *PoStResultReq) String() string { return proto.CompactTextString(m) }
func (*PoStResultReq) ProtoMessage()    {}
func (*PoStResultReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{7}
}

func (m *PoStResultReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoStResultReq.Unmarshal(m, b)
}
func (m *PoStResultReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoStResultReq.Marshal(b, m, deterministic)
}
func (m *PoStResultReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoStResultReq.Merge(m, src)
}
func (m *PoStResultReq) XXX_Size() int {
	return xxx_messageInfo_PoStResultReq.Size(m)
}
func (m *PoStResultReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PoStResultReq.DiscardUnknown(m)
}

var xxx_messageInfo_PoStResultReq proto.InternalMessageInfo

func (m *PoStResultReq) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type PoStResultRes struct {
	Got                  bool     `protobuf:"varint,1,opt,name=Got,proto3" json:"Got,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoStResultRes) Reset()         { *m = PoStResultRes{} }
func (m *PoStResultRes) String() string { return proto.CompactTextString(m) }
func (*PoStResultRes) ProtoMessage()    {}
func (*PoStResultRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c604b3c47fe2c93, []int{8}
}

func (m *PoStResultRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoStResultRes.Unmarshal(m, b)
}
func (m *PoStResultRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoStResultRes.Marshal(b, m, deterministic)
}
func (m *PoStResultRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoStResultRes.Merge(m, src)
}
func (m *PoStResultRes) XXX_Size() int {
	return xxx_messageInfo_PoStResultRes.Size(m)
}
func (m *PoStResultRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PoStResultRes.DiscardUnknown(m)
}

var xxx_messageInfo_PoStResultRes proto.InternalMessageInfo

func (m *PoStResultRes) GetGot() bool {
	if m != nil {
		return m.Got
	}
	return false
}

func init() {
	proto.RegisterType((*AddPieceReq)(nil), "types.AddPieceReq")
	proto.RegisterType((*AddPieceRes)(nil), "types.AddPieceRes")
	proto.RegisterType((*GenPoStReq)(nil), "types.GenPoStReq")
	proto.RegisterType((*GenPoStRes)(nil), "types.GenPoStRes")
	proto.RegisterType((*SealResultReq)(nil), "types.SealResultReq")
	proto.RegisterType((*PieceMeta)(nil), "types.PieceMeta")
	proto.RegisterType((*SealResultRes)(nil), "types.SealResultRes")
	proto.RegisterType((*PoStResultReq)(nil), "types.PoStResultReq")
	proto.RegisterType((*PoStResultRes)(nil), "types.PoStResultRes")
}

func init() { proto.RegisterFile("proto/types/common.proto", fileDescriptor_1c604b3c47fe2c93) }

var fileDescriptor_1c604b3c47fe2c93 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x86, 0x69, 0xab, 0xce, 0x1e, 0xab, 0x48, 0xd8, 0x45, 0x18, 0x1b, 0x74, 0x81, 0x8d, 0xee,
	0x46, 0xc7, 0xf6, 0x04, 0x43, 0x87, 0x94, 0x21, 0x94, 0xf4, 0x09, 0xba, 0x36, 0x1b, 0x82, 0x6d,
	0xba, 0x26, 0x5e, 0xb8, 0xbb, 0xbd, 0xf9, 0xc8, 0x69, 0xb5, 0x55, 0xbc, 0x3b, 0xf9, 0x4e, 0xce,
	0xff, 0xf7, 0xfc, 0x29, 0xd0, 0xb2, 0x92, 0x5a, 0xce, 0xf5, 0xbe, 0x14, 0x6a, 0x9e, 0xca, 0x3c,
	0x97, 0xc5, 0x0c, 0x11, 0xe9, 0x23, 0x63, 0x77, 0x30, 0x7a, 0xcb, 0xb2, 0x68, 0x23, 0x52, 0xc1,
	0xc5, 0x0f, 0x99, 0x80, 0x1d, 0x96, 0xd4, 0xf2, 0xad, 0xc0, 0xe3, 0x76, 0x58, 0xb2, 0xa7, 0x6e,
	0x5b, 0x91, 0x1b, 0x18, 0xc6, 0x22, 0xd5, 0xb2, 0x0a, 0x33, 0x6a, 0xfb, 0x56, 0xd0, 0xe3, 0xc7,
	0x33, 0x7b, 0x06, 0x58, 0x89, 0x22, 0x92, 0xb1, 0xbe, 0x20, 0x44, 0x08, 0xf4, 0x62, 0x21, 0xea,
	0x29, 0x8f, 0x63, 0xcd, 0xbc, 0xce, 0x84, 0x62, 0x7f, 0x36, 0x8c, 0x63, 0x91, 0x6c, 0xb9, 0x50,
	0xbb, 0x2d, 0x6a, 0xb4, 0x6e, 0x4b, 0x54, 0x6a, 0xdd, 0x96, 0xe4, 0x11, 0x26, 0xe6, 0x72, 0xac,
	0x13, 0xbd, 0x53, 0x0b, 0x99, 0x09, 0x54, 0x1e, 0xf3, 0x33, 0x4a, 0x18, 0x78, 0x86, 0xbc, 0x57,
	0x95, 0xac, 0xd6, 0xea, 0x9b, 0x3a, 0xbe, 0x15, 0xb8, 0xfc, 0x84, 0x91, 0x6b, 0xe8, 0x2f, 0x64,
	0x9e, 0x2f, 0x69, 0x0f, 0x3f, 0xae, 0x3e, 0x1c, 0x28, 0xa7, 0xfd, 0x96, 0x72, 0x72, 0x0b, 0x2e,
	0x16, 0xb1, 0x4e, 0x2a, 0x3a, 0xc0, 0x4e, 0x0b, 0xcc, 0x4c, 0x54, 0x49, 0xf9, 0x45, 0xaf, 0xea,
	0x19, 0x3c, 0x90, 0x00, 0x06, 0x98, 0xa0, 0xa2, 0x43, 0xdf, 0x09, 0x46, 0x2f, 0xd3, 0x19, 0x66,
	0x3f, 0x43, 0xb8, 0x16, 0x3a, 0xe1, 0x4d, 0x9f, 0x49, 0x70, 0x8f, 0x90, 0x4c, 0xc1, 0xf9, 0x10,
	0x7b, 0xdc, 0xdc, 0xe5, 0xa6, 0xc4, 0x10, 0x37, 0xbf, 0xa2, 0x89, 0x1e, 0x6b, 0x13, 0x44, 0x58,
	0xa4, 0xdb, 0x9d, 0xda, 0xc8, 0xa2, 0xf6, 0x76, 0xd0, 0xfb, 0x8c, 0x1e, 0xd6, 0x89, 0xba, 0x4b,
	0x46, 0xec, 0xfe, 0x34, 0x73, 0x65, 0x4c, 0x57, 0x52, 0xa3, 0xe9, 0x90, 0x9b, 0x92, 0x3d, 0xc0,
	0xb8, 0x79, 0xa2, 0xe6, 0x59, 0x8e, 0x4b, 0x5a, 0x9d, 0x25, 0x8d, 0x52, 0xf7, 0xda, 0x05, 0xa5,
	0xcf, 0x01, 0xfe, 0x79, 0xaf, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x94, 0x07, 0xe8, 0x95,
	0x02, 0x00, 0x00,
}
