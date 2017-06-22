// Code generated by protoc-gen-go.
// source: RowProcessor.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcessRequest struct {
	RowProcessorClassName              *string `protobuf:"bytes,1,req,name=row_processor_class_name,json=rowProcessorClassName" json:"row_processor_class_name,omitempty"`
	RowProcessorInitializerMessageName *string `protobuf:"bytes,2,opt,name=row_processor_initializer_message_name,json=rowProcessorInitializerMessageName" json:"row_processor_initializer_message_name,omitempty"`
	RowProcessorInitializerMessage     []byte  `protobuf:"bytes,3,opt,name=row_processor_initializer_message,json=rowProcessorInitializerMessage" json:"row_processor_initializer_message,omitempty"`
	NonceGroup                         *uint64 `protobuf:"varint,4,opt,name=nonce_group,json=nonceGroup" json:"nonce_group,omitempty"`
	Nonce                              *uint64 `protobuf:"varint,5,opt,name=nonce" json:"nonce,omitempty"`
	XXX_unrecognized                   []byte  `json:"-"`
}

func (m *ProcessRequest) Reset()                    { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()               {}
func (*ProcessRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *ProcessRequest) GetRowProcessorClassName() string {
	if m != nil && m.RowProcessorClassName != nil {
		return *m.RowProcessorClassName
	}
	return ""
}

func (m *ProcessRequest) GetRowProcessorInitializerMessageName() string {
	if m != nil && m.RowProcessorInitializerMessageName != nil {
		return *m.RowProcessorInitializerMessageName
	}
	return ""
}

func (m *ProcessRequest) GetRowProcessorInitializerMessage() []byte {
	if m != nil {
		return m.RowProcessorInitializerMessage
	}
	return nil
}

func (m *ProcessRequest) GetNonceGroup() uint64 {
	if m != nil && m.NonceGroup != nil {
		return *m.NonceGroup
	}
	return 0
}

func (m *ProcessRequest) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

type ProcessResponse struct {
	RowProcessorResult []byte `protobuf:"bytes,1,req,name=row_processor_result,json=rowProcessorResult" json:"row_processor_result,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *ProcessResponse) Reset()                    { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string            { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()               {}
func (*ProcessResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *ProcessResponse) GetRowProcessorResult() []byte {
	if m != nil {
		return m.RowProcessorResult
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcessRequest)(nil), "pb.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "pb.ProcessResponse")
}

var fileDescriptor24 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xe5, 0xfc, 0xad, 0x7e, 0x31, 0x54, 0x20, 0xb9, 0x45, 0x8a, 0x58, 0x40, 0xe9, 0x02,
	0x21, 0x16, 0x16, 0xea, 0x86, 0x7d, 0xbb, 0xa0, 0x5d, 0x80, 0x2a, 0x73, 0x80, 0xc8, 0x49, 0x87,
	0x34, 0x52, 0xea, 0x31, 0x76, 0x42, 0x25, 0x4e, 0xc0, 0x2d, 0xe0, 0xa8, 0x24, 0x4e, 0x81, 0x64,
	0x03, 0x4b, 0xbf, 0xf7, 0xe6, 0xf3, 0x68, 0x1e, 0x70, 0x49, 0xbb, 0x95, 0xa5, 0x04, 0x9d, 0x23,
	0x2b, 0x8c, 0xa5, 0x82, 0x78, 0x60, 0xe2, 0xc9, 0x7b, 0x00, 0x47, 0x7b, 0x5d, 0xe2, 0x73, 0x89,
	0xae, 0xe0, 0xb7, 0x10, 0x5a, 0xda, 0x45, 0xe6, 0x2b, 0x1d, 0x25, 0xb9, 0x72, 0x2e, 0xd2, 0x6a,
	0x8b, 0x21, 0x1b, 0x07, 0x57, 0x07, 0xf2, 0xc4, 0xb6, 0x60, 0xf3, 0xda, 0x7d, 0xa8, 0x4c, 0x2e,
	0xe1, 0xb2, 0x3b, 0x98, 0xe9, 0xac, 0xc8, 0x54, 0x9e, 0xbd, 0xa2, 0x8d, 0xb6, 0x95, 0xa4, 0x52,
	0x6c, 0x30, 0xc1, 0x98, 0x55, 0x98, 0x49, 0x1b, 0xb3, 0xfc, 0xc9, 0xde, 0x37, 0x51, 0xcf, 0x5c,
	0xc2, 0xc5, 0x9f, 0xcc, 0xf0, 0x5f, 0x85, 0x1b, 0xc8, 0xb3, 0xdf, 0x71, 0xfc, 0x1c, 0x0e, 0x35,
	0xe9, 0x04, 0xa3, 0xd4, 0x52, 0x69, 0xc2, 0x5e, 0x35, 0xd4, 0x93, 0xe0, 0xa5, 0xbb, 0x5a, 0xe1,
	0x23, 0xe8, 0xfb, 0x57, 0xd8, 0xf7, 0x56, 0xf3, 0x98, 0xcc, 0xe1, 0xf8, 0xfb, 0x40, 0xce, 0x90,
	0x76, 0xc8, 0x6f, 0x60, 0xd4, 0x5d, 0xca, 0xa2, 0x2b, 0xf3, 0xc2, 0x5f, 0x67, 0x20, 0x79, 0x7b,
	0x0f, 0xe9, 0x9d, 0xe9, 0x12, 0x86, 0xed, 0x02, 0x1e, 0xd1, 0xbe, 0x64, 0x09, 0xf2, 0x29, 0xfc,
	0xdf, 0x6b, 0x9c, 0x0b, 0x13, 0x8b, 0x6e, 0x13, 0xa7, 0xc3, 0x8e, 0xd6, 0x7c, 0x3e, 0x5b, 0xc0,
	0x35, 0xd9, 0x54, 0x28, 0xa3, 0x92, 0x0d, 0x8a, 0x8d, 0x5a, 0x13, 0x19, 0xb1, 0x89, 0x95, 0xc3,
	0xa6, 0xd6, 0xb8, 0x7c, 0x12, 0x29, 0x6a, 0xb4, 0xaa, 0xc0, 0xf5, 0xac, 0xd3, 0xfb, 0xaa, 0xf6,
	0xdd, 0x82, 0xbd, 0x31, 0xf6, 0xc1, 0xd8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x59, 0xe2,
	0xfb, 0x14, 0x02, 0x00, 0x00,
}
