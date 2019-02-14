// Code generated by protoc-gen-go. DO NOT EDIT.
// source: explorer.proto

package explorer_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type ChainMeta struct {
	Height               string   `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	TotalCandidates      int64    `protobuf:"varint,2,opt,name=totalCandidates,proto3" json:"totalCandidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainMeta) Reset()         { *m = ChainMeta{} }
func (m *ChainMeta) String() string { return proto.CompactTextString(m) }
func (*ChainMeta) ProtoMessage()    {}
func (*ChainMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{0}
}

func (m *ChainMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainMeta.Unmarshal(m, b)
}
func (m *ChainMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainMeta.Marshal(b, m, deterministic)
}
func (m *ChainMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainMeta.Merge(m, src)
}
func (m *ChainMeta) XXX_Size() int {
	return xxx_messageInfo_ChainMeta.Size(m)
}
func (m *ChainMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ChainMeta proto.InternalMessageInfo

func (m *ChainMeta) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *ChainMeta) GetTotalCandidates() int64 {
	if m != nil {
		return m.TotalCandidates
	}
	return 0
}

type Bucket struct {
	// hex string
	Voter         string `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Votes         string `protobuf:"bytes,2,opt,name=votes,proto3" json:"votes,omitempty"`
	WeightedVotes string `protobuf:"bytes,3,opt,name=weightedVotes,proto3" json:"weightedVotes,omitempty"`
	// seconds
	RemainingDuration    *duration.Duration `protobuf:"bytes,4,opt,name=remainingDuration,proto3" json:"remainingDuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

func (m *Bucket) GetVotes() string {
	if m != nil {
		return m.Votes
	}
	return ""
}

func (m *Bucket) GetWeightedVotes() string {
	if m != nil {
		return m.WeightedVotes
	}
	return ""
}

func (m *Bucket) GetRemainingDuration() *duration.Duration {
	if m != nil {
		return m.RemainingDuration
	}
	return nil
}

type Candidate struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// hex string
	PubKey               string   `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	TotalWeightedVotes   string   `protobuf:"bytes,3,opt,name=totalWeightedVotes,proto3" json:"totalWeightedVotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candidate) Reset()         { *m = Candidate{} }
func (m *Candidate) String() string { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()    {}
func (*Candidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{2}
}

func (m *Candidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidate.Unmarshal(m, b)
}
func (m *Candidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidate.Marshal(b, m, deterministic)
}
func (m *Candidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidate.Merge(m, src)
}
func (m *Candidate) XXX_Size() int {
	return xxx_messageInfo_Candidate.Size(m)
}
func (m *Candidate) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidate.DiscardUnknown(m)
}

var xxx_messageInfo_Candidate proto.InternalMessageInfo

func (m *Candidate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Candidate) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Candidate) GetTotalWeightedVotes() string {
	if m != nil {
		return m.TotalWeightedVotes
	}
	return ""
}

type GetCandidatesRequest struct {
	Height               string   `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCandidatesRequest) Reset()         { *m = GetCandidatesRequest{} }
func (m *GetCandidatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCandidatesRequest) ProtoMessage()    {}
func (*GetCandidatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{3}
}

func (m *GetCandidatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCandidatesRequest.Unmarshal(m, b)
}
func (m *GetCandidatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCandidatesRequest.Marshal(b, m, deterministic)
}
func (m *GetCandidatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCandidatesRequest.Merge(m, src)
}
func (m *GetCandidatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCandidatesRequest.Size(m)
}
func (m *GetCandidatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCandidatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCandidatesRequest proto.InternalMessageInfo

func (m *GetCandidatesRequest) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *GetCandidatesRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetCandidatesRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBucketsByCandidateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height               string   `protobuf:"bytes,2,opt,name=height,proto3" json:"height,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBucketsByCandidateRequest) Reset()         { *m = GetBucketsByCandidateRequest{} }
func (m *GetBucketsByCandidateRequest) String() string { return proto.CompactTextString(m) }
func (*GetBucketsByCandidateRequest) ProtoMessage()    {}
func (*GetBucketsByCandidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{4}
}

func (m *GetBucketsByCandidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBucketsByCandidateRequest.Unmarshal(m, b)
}
func (m *GetBucketsByCandidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBucketsByCandidateRequest.Marshal(b, m, deterministic)
}
func (m *GetBucketsByCandidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBucketsByCandidateRequest.Merge(m, src)
}
func (m *GetBucketsByCandidateRequest) XXX_Size() int {
	return xxx_messageInfo_GetBucketsByCandidateRequest.Size(m)
}
func (m *GetBucketsByCandidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBucketsByCandidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBucketsByCandidateRequest proto.InternalMessageInfo

func (m *GetBucketsByCandidateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetBucketsByCandidateRequest) GetHeight() string {
	if m != nil {
		return m.Height
	}
	return ""
}

func (m *GetBucketsByCandidateRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetBucketsByCandidateRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type CandidateResponse struct {
	Candidates           []*Candidate `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CandidateResponse) Reset()         { *m = CandidateResponse{} }
func (m *CandidateResponse) String() string { return proto.CompactTextString(m) }
func (*CandidateResponse) ProtoMessage()    {}
func (*CandidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{5}
}

func (m *CandidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CandidateResponse.Unmarshal(m, b)
}
func (m *CandidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CandidateResponse.Marshal(b, m, deterministic)
}
func (m *CandidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandidateResponse.Merge(m, src)
}
func (m *CandidateResponse) XXX_Size() int {
	return xxx_messageInfo_CandidateResponse.Size(m)
}
func (m *CandidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CandidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CandidateResponse proto.InternalMessageInfo

func (m *CandidateResponse) GetCandidates() []*Candidate {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type BucketResponse struct {
	Buckets              []*Bucket `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BucketResponse) Reset()         { *m = BucketResponse{} }
func (m *BucketResponse) String() string { return proto.CompactTextString(m) }
func (*BucketResponse) ProtoMessage()    {}
func (*BucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92013ac430d0de85, []int{6}
}

func (m *BucketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketResponse.Unmarshal(m, b)
}
func (m *BucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketResponse.Marshal(b, m, deterministic)
}
func (m *BucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketResponse.Merge(m, src)
}
func (m *BucketResponse) XXX_Size() int {
	return xxx_messageInfo_BucketResponse.Size(m)
}
func (m *BucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BucketResponse proto.InternalMessageInfo

func (m *BucketResponse) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainMeta)(nil), "explorer_pb.ChainMeta")
	proto.RegisterType((*Bucket)(nil), "explorer_pb.Bucket")
	proto.RegisterType((*Candidate)(nil), "explorer_pb.Candidate")
	proto.RegisterType((*GetCandidatesRequest)(nil), "explorer_pb.GetCandidatesRequest")
	proto.RegisterType((*GetBucketsByCandidateRequest)(nil), "explorer_pb.GetBucketsByCandidateRequest")
	proto.RegisterType((*CandidateResponse)(nil), "explorer_pb.CandidateResponse")
	proto.RegisterType((*BucketResponse)(nil), "explorer_pb.BucketResponse")
}

func init() { proto.RegisterFile("explorer.proto", fileDescriptor_92013ac430d0de85) }

var fileDescriptor_92013ac430d0de85 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x90, 0x92, 0x89, 0x5a, 0xd4, 0xa1, 0x44, 0xc1, 0x45, 0x55, 0xb0, 0x38, 0x98,
	0x03, 0xae, 0x14, 0x24, 0x2e, 0x1c, 0x90, 0x5a, 0xaa, 0x1c, 0xaa, 0x5e, 0x2c, 0x04, 0x17, 0x24,
	0xb4, 0x6e, 0x26, 0x8e, 0x45, 0xe2, 0x35, 0xf6, 0x18, 0x9a, 0x1b, 0xff, 0xc2, 0x8f, 0xa2, 0xec,
	0xae, 0x13, 0xc7, 0xb1, 0x6f, 0x3b, 0xf3, 0x66, 0xe7, 0xcd, 0x7b, 0x33, 0x70, 0x4a, 0x8f, 0xe9,
	0x52, 0x66, 0x94, 0xf9, 0x69, 0x26, 0x59, 0xe2, 0xa0, 0x8c, 0x7f, 0xa4, 0xa1, 0x73, 0x19, 0x49,
	0x19, 0x2d, 0xe9, 0x4a, 0x41, 0x61, 0x31, 0xbf, 0x9a, 0x15, 0x99, 0xe0, 0x58, 0x26, 0xba, 0xd8,
	0xb9, 0xa8, 0xe3, 0xb4, 0x4a, 0x79, 0xad, 0x41, 0xf7, 0x1e, 0xfa, 0x37, 0x0b, 0x11, 0x27, 0xf7,
	0xc4, 0x02, 0x87, 0xd0, 0x5b, 0x50, 0x1c, 0x2d, 0x78, 0x64, 0x8d, 0x2d, 0xaf, 0x1f, 0x98, 0x08,
	0x3d, 0x78, 0xc6, 0x92, 0xc5, 0xf2, 0x46, 0x24, 0xb3, 0x78, 0x26, 0x98, 0xf2, 0x51, 0x67, 0x6c,
	0x79, 0x76, 0x50, 0x4f, 0xbb, 0xff, 0x2c, 0xe8, 0x5d, 0x17, 0x0f, 0x3f, 0x89, 0xf1, 0x1c, 0x9e,
	0xfc, 0x96, 0x4c, 0x99, 0xe9, 0xa5, 0x83, 0x32, 0xab, 0x1b, 0x98, 0x6c, 0x8e, 0x6f, 0xe0, 0xe4,
	0x8f, 0xa2, 0xa2, 0xd9, 0x57, 0x85, 0xda, 0x0a, 0xdd, 0x4f, 0xe2, 0x14, 0xce, 0x32, 0x5a, 0x89,
	0x38, 0x89, 0x93, 0xe8, 0xb3, 0xd1, 0x38, 0xea, 0x8e, 0x2d, 0x6f, 0x30, 0x79, 0xe9, 0x6b, 0x91,
	0x7e, 0x29, 0xd2, 0x2f, 0x0b, 0x82, 0xc3, 0x3f, 0x6e, 0x04, 0xfd, 0xed, 0xcc, 0x88, 0xd0, 0x4d,
	0xc4, 0x8a, 0xcc, 0x98, 0xea, 0xbd, 0x31, 0x22, 0x2d, 0xc2, 0x3b, 0x5a, 0x9b, 0x31, 0x4d, 0x84,
	0x3e, 0xa0, 0x52, 0xfc, 0xad, 0x61, 0xd8, 0x06, 0xc4, 0xfd, 0x0e, 0xe7, 0x53, 0xe2, 0x9d, 0x3f,
	0x01, 0xfd, 0x2a, 0x28, 0xe7, 0x56, 0xa3, 0x87, 0xd0, 0x93, 0xf3, 0x79, 0x4e, 0x6c, 0xfc, 0x35,
	0xd1, 0xc6, 0xb5, 0x65, 0xbc, 0x8a, 0x59, 0x51, 0xd9, 0x81, 0x0e, 0xdc, 0x47, 0x78, 0x35, 0x25,
	0xd6, 0x76, 0xe7, 0xd7, 0xeb, 0x2d, 0x4d, 0xc9, 0xd2, 0xa2, 0xcc, 0x30, 0x77, 0x5a, 0x98, 0xed,
	0x66, 0xe6, 0x6e, 0x95, 0xf9, 0x0e, 0xce, 0x2a, 0x6c, 0x79, 0x2a, 0x93, 0x9c, 0xf0, 0x03, 0xc0,
	0xc3, 0xee, 0x40, 0xac, 0xb1, 0xed, 0x0d, 0x26, 0x43, 0xbf, 0x72, 0xa9, 0xfe, 0xee, 0x4f, 0xa5,
	0xd2, 0xfd, 0x04, 0xa7, 0x5a, 0xc3, 0xb6, 0xd3, 0x3b, 0x38, 0x0e, 0xb5, 0x2a, 0xd3, 0xe6, 0xf9,
	0x5e, 0x1b, 0x53, 0x5d, 0xd6, 0x4c, 0xfe, 0x76, 0xe0, 0xe9, 0xad, 0xc1, 0xf1, 0x23, 0x1c, 0x47,
	0xc4, 0xfa, 0x9c, 0x0f, 0x8e, 0xe2, 0x76, 0x73, 0xf9, 0x4e, 0x6d, 0xa8, 0xf2, 0xfc, 0xdd, 0x23,
	0xfc, 0x02, 0x27, 0x51, 0x75, 0x5f, 0xf8, 0x7a, 0xaf, 0xb4, 0x69, 0x97, 0xce, 0x65, 0x8b, 0x44,
	0x23, 0xc6, 0x3d, 0x42, 0x01, 0x2f, 0xa2, 0xa6, 0x3d, 0xe1, 0xdb, 0x7a, 0xf7, 0xd6, 0x5d, 0x3a,
	0x17, 0x4d, 0x0e, 0x6c, 0x29, 0xc2, 0x9e, 0x92, 0xf8, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x26, 0x58, 0x6d, 0xf6, 0x29, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExplorerClient is the client API for Explorer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExplorerClient interface {
	// get the blockchain meta data
	GetMeta(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ChainMeta, error)
	// get candidates
	GetCandidates(ctx context.Context, in *GetCandidatesRequest, opts ...grpc.CallOption) (*CandidateResponse, error)
	// get buckets by candidate
	GetBucketsByCandidate(ctx context.Context, in *GetBucketsByCandidateRequest, opts ...grpc.CallOption) (*BucketResponse, error)
}

type explorerClient struct {
	cc *grpc.ClientConn
}

func NewExplorerClient(cc *grpc.ClientConn) ExplorerClient {
	return &explorerClient{cc}
}

func (c *explorerClient) GetMeta(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ChainMeta, error) {
	out := new(ChainMeta)
	err := c.cc.Invoke(ctx, "/explorer_pb.Explorer/getMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerClient) GetCandidates(ctx context.Context, in *GetCandidatesRequest, opts ...grpc.CallOption) (*CandidateResponse, error) {
	out := new(CandidateResponse)
	err := c.cc.Invoke(ctx, "/explorer_pb.Explorer/getCandidates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerClient) GetBucketsByCandidate(ctx context.Context, in *GetBucketsByCandidateRequest, opts ...grpc.CallOption) (*BucketResponse, error) {
	out := new(BucketResponse)
	err := c.cc.Invoke(ctx, "/explorer_pb.Explorer/getBucketsByCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExplorerServer is the server API for Explorer service.
type ExplorerServer interface {
	// get the blockchain meta data
	GetMeta(context.Context, *empty.Empty) (*ChainMeta, error)
	// get candidates
	GetCandidates(context.Context, *GetCandidatesRequest) (*CandidateResponse, error)
	// get buckets by candidate
	GetBucketsByCandidate(context.Context, *GetBucketsByCandidateRequest) (*BucketResponse, error)
}

func RegisterExplorerServer(s *grpc.Server, srv ExplorerServer) {
	s.RegisterService(&_Explorer_serviceDesc, srv)
}

func _Explorer_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/explorer_pb.Explorer/GetMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServer).GetMeta(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Explorer_GetCandidates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandidatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServer).GetCandidates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/explorer_pb.Explorer/GetCandidates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServer).GetCandidates(ctx, req.(*GetCandidatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Explorer_GetBucketsByCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketsByCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServer).GetBucketsByCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/explorer_pb.Explorer/GetBucketsByCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServer).GetBucketsByCandidate(ctx, req.(*GetBucketsByCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Explorer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "explorer_pb.Explorer",
	HandlerType: (*ExplorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMeta",
			Handler:    _Explorer_GetMeta_Handler,
		},
		{
			MethodName: "getCandidates",
			Handler:    _Explorer_GetCandidates_Handler,
		},
		{
			MethodName: "getBucketsByCandidate",
			Handler:    _Explorer_GetBucketsByCandidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "explorer.proto",
}
