// Code generated by protoc-gen-go.
// source: TestProcedure.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestTableDDLStateData struct {
	TableName        *string `protobuf:"bytes,1,req,name=table_name,json=tableName" json:"table_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestTableDDLStateData) Reset()                    { *m = TestTableDDLStateData{} }
func (m *TestTableDDLStateData) String() string            { return proto.CompactTextString(m) }
func (*TestTableDDLStateData) ProtoMessage()               {}
func (*TestTableDDLStateData) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *TestTableDDLStateData) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

func init() {
	proto.RegisterType((*TestTableDDLStateData)(nil), "pb.TestTableDDLStateData")
}

var fileDescriptor22 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0e, 0x49, 0x2d, 0x2e,
	0x09, 0x28, 0xca, 0x4f, 0x4e, 0x4d, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x52, 0x32, 0xe3, 0x12, 0x05, 0x49, 0x85, 0x24, 0x26, 0xe5, 0xa4, 0xba, 0xb8,
	0xf8, 0x04, 0x97, 0x24, 0x96, 0xa4, 0xba, 0x24, 0x96, 0x24, 0x0a, 0xc9, 0x72, 0x71, 0x95, 0x80,
	0x04, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0x38, 0xc1, 0x22,
	0x7e, 0x40, 0x01, 0x27, 0x0f, 0x2e, 0xbd, 0xfc, 0xa2, 0x74, 0xbd, 0xc4, 0x82, 0xc4, 0xe4, 0x8c,
	0x54, 0xbd, 0x8c, 0xc4, 0x94, 0xfc, 0xfc, 0x02, 0xbd, 0x8c, 0xa4, 0xc4, 0xe2, 0x54, 0xbd, 0xcc,
	0x82, 0x64, 0x88, 0x05, 0x49, 0xa5, 0x69, 0x7a, 0xe9, 0xa9, 0x79, 0xa9, 0x45, 0x40, 0x33, 0x53,
	0x9c, 0x50, 0x9d, 0x10, 0x00, 0x52, 0x50, 0xdc, 0xc1, 0xc8, 0x08, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xd6, 0x01, 0x1a, 0xe3, 0x9b, 0x00, 0x00, 0x00,
}