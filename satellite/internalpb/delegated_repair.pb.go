// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delegated_repair.proto

package internalpb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"

	pb "storj.io/common/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RepairJobRequest struct {
	// When not the first request, this will include the result of the last job
	LastJobResult        *RepairJobResult `protobuf:"bytes,1,opt,name=last_job_result,json=lastJobResult,proto3" json:"last_job_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RepairJobRequest) Reset()         { *m = RepairJobRequest{} }
func (m *RepairJobRequest) String() string { return proto.CompactTextString(m) }
func (*RepairJobRequest) ProtoMessage()    {}
func (*RepairJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{0}
}
func (m *RepairJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobRequest.Unmarshal(m, b)
}
func (m *RepairJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobRequest.Marshal(b, m, deterministic)
}
func (m *RepairJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobRequest.Merge(m, src)
}
func (m *RepairJobRequest) XXX_Size() int {
	return xxx_messageInfo_RepairJobRequest.Size(m)
}
func (m *RepairJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobRequest proto.InternalMessageInfo

func (m *RepairJobRequest) GetLastJobResult() *RepairJobResult {
	if m != nil {
		return m.LastJobResult
	}
	return nil
}

type RepairJobResponse struct {
	// When a job is available, this will be filled in
	NewJob *RepairJobDefinition `protobuf:"bytes,1,opt,name=new_job,json=newJob,proto3" json:"new_job,omitempty"`
	// Otherwise, client should wait this many milliseconds and then try again
	ComeBackInMillis     int32    `protobuf:"varint,2,opt,name=come_back_in_millis,json=comeBackInMillis,proto3" json:"come_back_in_millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairJobResponse) Reset()         { *m = RepairJobResponse{} }
func (m *RepairJobResponse) String() string { return proto.CompactTextString(m) }
func (*RepairJobResponse) ProtoMessage()    {}
func (*RepairJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{1}
}
func (m *RepairJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobResponse.Unmarshal(m, b)
}
func (m *RepairJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobResponse.Marshal(b, m, deterministic)
}
func (m *RepairJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobResponse.Merge(m, src)
}
func (m *RepairJobResponse) XXX_Size() int {
	return xxx_messageInfo_RepairJobResponse.Size(m)
}
func (m *RepairJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobResponse proto.InternalMessageInfo

func (m *RepairJobResponse) GetNewJob() *RepairJobDefinition {
	if m != nil {
		return m.NewJob
	}
	return nil
}

func (m *RepairJobResponse) GetComeBackInMillis() int32 {
	if m != nil {
		return m.ComeBackInMillis
	}
	return 0
}

type RepairJobDefinition struct {
	// Identifier for this job
	JobId []byte `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Signed GET orders for all believed-healthy pieces to be downloaded
	GetOrders []*pb.AddressedOrderLimit `protobuf:"bytes,2,rep,name=get_orders,json=getOrders,proto3" json:"get_orders,omitempty"`
	// Private piece key to use for fetching
	PrivateKeyForGet []byte `protobuf:"bytes,3,opt,name=private_key_for_get,json=privateKeyForGet,proto3" json:"private_key_for_get,omitempty"`
	// Signed PUT orders for all possible pieces to be uploaded (not including
	// piece numbers in get_orders)
	PutOrders []*pb.AddressedOrderLimit `protobuf:"bytes,4,rep,name=put_orders,json=putOrders,proto3" json:"put_orders,omitempty"`
	// Private piece key to use for storing
	PrivateKeyForPut []byte `protobuf:"bytes,5,opt,name=private_key_for_put,json=privateKeyForPut,proto3" json:"private_key_for_put,omitempty"`
	// Redundancy scheme used by the segment to be repaired
	Redundancy *pb.RedundancyScheme `protobuf:"bytes,6,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	// Size of the segment to be repaired
	SegmentSize int64 `protobuf:"varint,7,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	// Target piece count (worker should try to upload enough pieces so that
	// this count is achieved)
	DesiredPieceCount int32 `protobuf:"varint,8,opt,name=desired_piece_count,json=desiredPieceCount,proto3" json:"desired_piece_count,omitempty"`
	// Job expiration time
	ExpirationTime       time.Time `protobuf:"bytes,9,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RepairJobDefinition) Reset()         { *m = RepairJobDefinition{} }
func (m *RepairJobDefinition) String() string { return proto.CompactTextString(m) }
func (*RepairJobDefinition) ProtoMessage()    {}
func (*RepairJobDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{2}
}
func (m *RepairJobDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobDefinition.Unmarshal(m, b)
}
func (m *RepairJobDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobDefinition.Marshal(b, m, deterministic)
}
func (m *RepairJobDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobDefinition.Merge(m, src)
}
func (m *RepairJobDefinition) XXX_Size() int {
	return xxx_messageInfo_RepairJobDefinition.Size(m)
}
func (m *RepairJobDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobDefinition proto.InternalMessageInfo

func (m *RepairJobDefinition) GetJobId() []byte {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *RepairJobDefinition) GetGetOrders() []*pb.AddressedOrderLimit {
	if m != nil {
		return m.GetOrders
	}
	return nil
}

func (m *RepairJobDefinition) GetPrivateKeyForGet() []byte {
	if m != nil {
		return m.PrivateKeyForGet
	}
	return nil
}

func (m *RepairJobDefinition) GetPutOrders() []*pb.AddressedOrderLimit {
	if m != nil {
		return m.PutOrders
	}
	return nil
}

func (m *RepairJobDefinition) GetPrivateKeyForPut() []byte {
	if m != nil {
		return m.PrivateKeyForPut
	}
	return nil
}

func (m *RepairJobDefinition) GetRedundancy() *pb.RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RepairJobDefinition) GetSegmentSize() int64 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *RepairJobDefinition) GetDesiredPieceCount() int32 {
	if m != nil {
		return m.DesiredPieceCount
	}
	return 0
}

func (m *RepairJobDefinition) GetExpirationTime() time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return time.Time{}
}

type RepairJobResult struct {
	// Identifier for this job, as given in RepairJobResponse
	JobId []byte `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Set nonzero only if the segment could not be reconstructed because of
	// too few pieces available.
	IrreparablePiecesRetrieved int32 `protobuf:"varint,2,opt,name=irreparable_pieces_retrieved,json=irreparablePiecesRetrieved,proto3" json:"irreparable_pieces_retrieved,omitempty"`
	// Set only if the segment could not be reconstructed.
	ReconstructError string `protobuf:"bytes,3,opt,name=reconstruct_error,json=reconstructError,proto3" json:"reconstruct_error,omitempty"`
	// Set only if new pieces could not be stored to any new nodes.
	StoreError string `protobuf:"bytes,4,opt,name=store_error,json=storeError,proto3" json:"store_error,omitempty"`
	// PieceHashes signed by storage nodes which were used to accomplish repair
	NewPiecesStored []*pb.PieceHash `protobuf:"bytes,5,rep,name=new_pieces_stored,json=newPiecesStored,proto3" json:"new_pieces_stored,omitempty"`
	// A copy of the put_orders list as provided in the corresponding
	// RepairJobDefinition
	PutOrders []*pb.AddressedOrderLimit `protobuf:"bytes,6,rep,name=put_orders,json=putOrders,proto3" json:"put_orders,omitempty"`
	// Pieces which should be _removed_ from the pointer. This will include
	// pieces for which the expected owning storage node returned a "not found"
	// error, as well as pieces which were downloaded but failed their
	// validation check.
	DeletePieceNums      []int32  `protobuf:"varint,7,rep,packed,name=delete_piece_nums,json=deletePieceNums,proto3" json:"delete_piece_nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairJobResult) Reset()         { *m = RepairJobResult{} }
func (m *RepairJobResult) String() string { return proto.CompactTextString(m) }
func (*RepairJobResult) ProtoMessage()    {}
func (*RepairJobResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{3}
}
func (m *RepairJobResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobResult.Unmarshal(m, b)
}
func (m *RepairJobResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobResult.Marshal(b, m, deterministic)
}
func (m *RepairJobResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobResult.Merge(m, src)
}
func (m *RepairJobResult) XXX_Size() int {
	return xxx_messageInfo_RepairJobResult.Size(m)
}
func (m *RepairJobResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobResult.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobResult proto.InternalMessageInfo

func (m *RepairJobResult) GetJobId() []byte {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *RepairJobResult) GetIrreparablePiecesRetrieved() int32 {
	if m != nil {
		return m.IrreparablePiecesRetrieved
	}
	return 0
}

func (m *RepairJobResult) GetReconstructError() string {
	if m != nil {
		return m.ReconstructError
	}
	return ""
}

func (m *RepairJobResult) GetStoreError() string {
	if m != nil {
		return m.StoreError
	}
	return ""
}

func (m *RepairJobResult) GetNewPiecesStored() []*pb.PieceHash {
	if m != nil {
		return m.NewPiecesStored
	}
	return nil
}

func (m *RepairJobResult) GetPutOrders() []*pb.AddressedOrderLimit {
	if m != nil {
		return m.PutOrders
	}
	return nil
}

func (m *RepairJobResult) GetDeletePieceNums() []int32 {
	if m != nil {
		return m.DeletePieceNums
	}
	return nil
}

func init() {
	proto.RegisterType((*RepairJobRequest)(nil), "satellite.delegated_repair.RepairJobRequest")
	proto.RegisterType((*RepairJobResponse)(nil), "satellite.delegated_repair.RepairJobResponse")
	proto.RegisterType((*RepairJobDefinition)(nil), "satellite.delegated_repair.RepairJobDefinition")
	proto.RegisterType((*RepairJobResult)(nil), "satellite.delegated_repair.RepairJobResult")
}

func init() { proto.RegisterFile("delegated_repair.proto", fileDescriptor_04d00d18c724d5a7) }

var fileDescriptor_04d00d18c724d5a7 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x5d, 0xd7, 0xb5, 0x5b, 0xdd, 0xb1, 0xb6, 0x99, 0x40, 0x51, 0x01, 0xb5, 0x14, 0x21, 0x55,
	0x8c, 0xa5, 0xd2, 0xb8, 0x04, 0x24, 0xd8, 0xf8, 0xd9, 0x06, 0x83, 0x29, 0xe5, 0x8a, 0x1b, 0xcb,
	0x89, 0xbf, 0x66, 0xee, 0x12, 0x3b, 0xd8, 0xce, 0xc6, 0x76, 0xc3, 0x0b, 0x70, 0xc1, 0x63, 0xf1,
	0x14, 0x20, 0x2e, 0x78, 0x0f, 0x64, 0x27, 0xed, 0xaa, 0x69, 0x43, 0xe5, 0x2e, 0xfe, 0xce, 0xb1,
	0xcf, 0xc9, 0xf7, 0x1d, 0x1b, 0xdd, 0xa2, 0x10, 0x43, 0x44, 0x34, 0x50, 0x2c, 0x21, 0x25, 0x4c,
	0x7a, 0xa9, 0x14, 0x5a, 0x38, 0x6d, 0x45, 0x34, 0xc4, 0x31, 0xd3, 0xe0, 0x5d, 0x66, 0xb4, 0x51,
	0x24, 0x22, 0x91, 0xf3, 0xda, 0x9d, 0x48, 0x88, 0x28, 0x86, 0x81, 0x5d, 0x05, 0xd9, 0x68, 0xa0,
	0x59, 0x02, 0x4a, 0x93, 0x24, 0x2d, 0x08, 0x6b, 0x09, 0x68, 0xc2, 0xf8, 0x68, 0xb2, 0x61, 0x55,
	0x48, 0x0a, 0x52, 0x15, 0xab, 0x46, 0x2a, 0x18, 0xd7, 0x20, 0x69, 0x90, 0x17, 0x7a, 0x11, 0x6a,
	0xfa, 0x56, 0x65, 0x5f, 0x04, 0x3e, 0x7c, 0xce, 0x40, 0x69, 0x67, 0x88, 0x1a, 0x31, 0x51, 0x1a,
	0x8f, 0x45, 0x80, 0x25, 0xa8, 0x2c, 0xd6, 0x6e, 0xa9, 0x5b, 0xea, 0xd7, 0xb7, 0x36, 0xbc, 0xeb,
	0x5d, 0x7a, 0x33, 0xc7, 0x98, 0x2d, 0xfe, 0x0d, 0x73, 0xc6, 0x74, 0xd9, 0xfb, 0x56, 0x42, 0xad,
	0x59, 0x4a, 0x2a, 0xb8, 0x02, 0x67, 0x17, 0x2d, 0x73, 0x38, 0x35, 0x4a, 0x85, 0xc4, 0x60, 0x2e,
	0x89, 0x97, 0x30, 0x62, 0x9c, 0x69, 0x26, 0xb8, 0x5f, 0xe5, 0x70, 0xba, 0x2f, 0x02, 0x67, 0x13,
	0xad, 0x87, 0x22, 0x01, 0x1c, 0x90, 0xf0, 0x18, 0x33, 0x8e, 0x13, 0x16, 0xc7, 0x4c, 0xb9, 0x8b,
	0xdd, 0x52, 0xbf, 0xe2, 0x37, 0x0d, 0xb4, 0x4d, 0xc2, 0xe3, 0x3d, 0x7e, 0x60, 0xeb, 0xbd, 0x3f,
	0x65, 0xb4, 0x7e, 0xc5, 0x71, 0xce, 0x4d, 0x54, 0x35, 0xbf, 0xcd, 0xa8, 0xf5, 0xb3, 0xea, 0x57,
	0xc6, 0x22, 0xd8, 0xa3, 0xce, 0x53, 0x84, 0x22, 0xd0, 0x38, 0xef, 0xa5, 0xbb, 0xd8, 0x2d, 0xf7,
	0xeb, 0x5b, 0x77, 0xbd, 0x69, 0xab, 0x5f, 0x50, 0x2a, 0x41, 0x29, 0xa0, 0x1f, 0x0c, 0xe1, 0x1d,
	0x4b, 0x98, 0xf6, 0x6b, 0x11, 0x68, 0xbb, 0x54, 0xc6, 0x5b, 0x2a, 0xd9, 0x09, 0xd1, 0x80, 0x8f,
	0xe1, 0x0c, 0x8f, 0x84, 0xc4, 0x11, 0x68, 0xb7, 0x6c, 0x15, 0x9a, 0x05, 0xf4, 0x16, 0xce, 0x5e,
	0x0b, 0xf9, 0x06, 0xb4, 0x11, 0x4b, 0xb3, 0xa9, 0xd8, 0xd2, 0x5c, 0x62, 0x69, 0xf6, 0x0f, 0xb1,
	0x34, 0xd3, 0x6e, 0xe5, 0x0a, 0xb1, 0xc3, 0x4c, 0x3b, 0x4f, 0x10, 0x92, 0x40, 0x33, 0x4e, 0x09,
	0x0f, 0xcf, 0xdc, 0xaa, 0x1d, 0xc2, 0x6d, 0xef, 0x22, 0x26, 0xfe, 0x14, 0x1c, 0x86, 0x47, 0x90,
	0x80, 0x3f, 0x43, 0x77, 0xee, 0xa1, 0x55, 0x05, 0x51, 0x02, 0x5c, 0x63, 0xc5, 0xce, 0xc1, 0x5d,
	0xee, 0x96, 0xfa, 0x65, 0xbf, 0x5e, 0xd4, 0x86, 0xec, 0x1c, 0x1c, 0x0f, 0xad, 0x53, 0x50, 0x4c,
	0x02, 0xc5, 0x29, 0x83, 0x10, 0x70, 0x28, 0x32, 0xae, 0xdd, 0x15, 0x3b, 0x97, 0x56, 0x01, 0x1d,
	0x1a, 0x64, 0xc7, 0x00, 0xce, 0x01, 0x6a, 0xc0, 0x97, 0x94, 0x49, 0x62, 0xc6, 0x81, 0x4d, 0xba,
	0xdd, 0x9a, 0x35, 0xd5, 0xf6, 0xf2, 0xe8, 0x7b, 0x93, 0xe8, 0x7b, 0x1f, 0x27, 0xd1, 0xdf, 0x5e,
	0xf9, 0xf1, 0xb3, 0xb3, 0xf0, 0xfd, 0x57, 0xa7, 0xe4, 0xaf, 0x5d, 0x6c, 0x36, 0x70, 0xef, 0xf7,
	0x22, 0x6a, 0x5c, 0x4a, 0xe6, 0x75, 0x33, 0x7e, 0x8e, 0xee, 0x30, 0x69, 0x92, 0x26, 0x49, 0x10,
	0x43, 0xee, 0x56, 0x61, 0x09, 0x5a, 0x32, 0x38, 0x01, 0x5a, 0x44, 0xa9, 0x3d, 0xc3, 0xb1, 0xb6,
	0x95, 0x3f, 0x61, 0x38, 0x1b, 0xa8, 0x25, 0x21, 0x14, 0x5c, 0x69, 0x99, 0x85, 0x1a, 0x83, 0x94,
	0x42, 0xda, 0x29, 0xd7, 0xfc, 0xe6, 0x0c, 0xf0, 0xca, 0xd4, 0x9d, 0x0e, 0xaa, 0x2b, 0x2d, 0x24,
	0x14, 0xb4, 0x25, 0x4b, 0x43, 0xb6, 0x94, 0x13, 0x9e, 0xa1, 0x96, 0xb9, 0x1b, 0x85, 0x0f, 0x0b,
	0x50, 0xb7, 0x62, 0xd3, 0xd0, 0xf2, 0x8a, 0x5b, 0x6d, 0x1d, 0xec, 0x12, 0x75, 0xe4, 0x37, 0x38,
	0x9c, 0xe6, 0x7e, 0x86, 0x96, 0x79, 0x29, 0x45, 0xd5, 0xff, 0x4c, 0xd1, 0x43, 0xd4, 0x32, 0xd7,
	0x4f, 0x17, 0x7d, 0xc0, 0x3c, 0x4b, 0x94, 0xbb, 0xdc, 0x2d, 0xf7, 0x2b, 0x7e, 0x23, 0x07, 0xac,
	0xd8, 0xfb, 0x2c, 0x51, 0x5b, 0x5f, 0x27, 0x37, 0x7b, 0x47, 0x08, 0x49, 0x19, 0x27, 0x5a, 0x48,
	0x67, 0x8c, 0x6a, 0xd3, 0xbe, 0x3b, 0x8f, 0xe6, 0x7c, 0x38, 0xec, 0xfb, 0xd3, 0xde, 0x9c, 0xf7,
	0x99, 0xb1, 0x6f, 0x48, 0x6f, 0x61, 0xfb, 0xc1, 0xa7, 0xfb, 0xa6, 0x3d, 0x63, 0x8f, 0x89, 0x81,
	0xfd, 0x18, 0x4c, 0x0f, 0x18, 0xd8, 0x1c, 0x73, 0x12, 0xa7, 0x41, 0x50, 0xb5, 0xc9, 0x79, 0xfc,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x7f, 0xbc, 0x37, 0x84, 0x05, 0x00, 0x00,
}
