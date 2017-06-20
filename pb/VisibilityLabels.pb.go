// Code generated by protoc-gen-go.
// source: VisibilityLabels.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VisibilityLabelsRequest struct {
	VisLabel         []*VisibilityLabel `protobuf:"bytes,1,rep,name=visLabel" json:"visLabel,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *VisibilityLabelsRequest) Reset()                    { *m = VisibilityLabelsRequest{} }
func (m *VisibilityLabelsRequest) String() string            { return proto.CompactTextString(m) }
func (*VisibilityLabelsRequest) ProtoMessage()               {}
func (*VisibilityLabelsRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *VisibilityLabelsRequest) GetVisLabel() []*VisibilityLabel {
	if m != nil {
		return m.VisLabel
	}
	return nil
}

type VisibilityLabel struct {
	Label            []byte  `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Ordinal          *uint32 `protobuf:"varint,2,opt,name=ordinal" json:"ordinal,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VisibilityLabel) Reset()                    { *m = VisibilityLabel{} }
func (m *VisibilityLabel) String() string            { return proto.CompactTextString(m) }
func (*VisibilityLabel) ProtoMessage()               {}
func (*VisibilityLabel) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *VisibilityLabel) GetLabel() []byte {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *VisibilityLabel) GetOrdinal() uint32 {
	if m != nil && m.Ordinal != nil {
		return *m.Ordinal
	}
	return 0
}

type VisibilityLabelsResponse struct {
	Result           []*RegionActionResult `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *VisibilityLabelsResponse) Reset()                    { *m = VisibilityLabelsResponse{} }
func (m *VisibilityLabelsResponse) String() string            { return proto.CompactTextString(m) }
func (*VisibilityLabelsResponse) ProtoMessage()               {}
func (*VisibilityLabelsResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

func (m *VisibilityLabelsResponse) GetResult() []*RegionActionResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type SetAuthsRequest struct {
	User             []byte   `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth             [][]byte `protobuf:"bytes,2,rep,name=auth" json:"auth,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SetAuthsRequest) Reset()                    { *m = SetAuthsRequest{} }
func (m *SetAuthsRequest) String() string            { return proto.CompactTextString(m) }
func (*SetAuthsRequest) ProtoMessage()               {}
func (*SetAuthsRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

func (m *SetAuthsRequest) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SetAuthsRequest) GetAuth() [][]byte {
	if m != nil {
		return m.Auth
	}
	return nil
}

type UserAuthorizations struct {
	User             []byte   `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth             []uint32 `protobuf:"varint,2,rep,name=auth" json:"auth,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UserAuthorizations) Reset()                    { *m = UserAuthorizations{} }
func (m *UserAuthorizations) String() string            { return proto.CompactTextString(m) }
func (*UserAuthorizations) ProtoMessage()               {}
func (*UserAuthorizations) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{4} }

func (m *UserAuthorizations) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserAuthorizations) GetAuth() []uint32 {
	if m != nil {
		return m.Auth
	}
	return nil
}

type MultiUserAuthorizations struct {
	UserAuths        []*UserAuthorizations `protobuf:"bytes,1,rep,name=userAuths" json:"userAuths,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *MultiUserAuthorizations) Reset()                    { *m = MultiUserAuthorizations{} }
func (m *MultiUserAuthorizations) String() string            { return proto.CompactTextString(m) }
func (*MultiUserAuthorizations) ProtoMessage()               {}
func (*MultiUserAuthorizations) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{5} }

func (m *MultiUserAuthorizations) GetUserAuths() []*UserAuthorizations {
	if m != nil {
		return m.UserAuths
	}
	return nil
}

type GetAuthsRequest struct {
	User             []byte `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAuthsRequest) Reset()                    { *m = GetAuthsRequest{} }
func (m *GetAuthsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAuthsRequest) ProtoMessage()               {}
func (*GetAuthsRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{6} }

func (m *GetAuthsRequest) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

type GetAuthsResponse struct {
	User             []byte   `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Auth             [][]byte `protobuf:"bytes,2,rep,name=auth" json:"auth,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetAuthsResponse) Reset()                    { *m = GetAuthsResponse{} }
func (m *GetAuthsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAuthsResponse) ProtoMessage()               {}
func (*GetAuthsResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{7} }

func (m *GetAuthsResponse) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetAuthsResponse) GetAuth() [][]byte {
	if m != nil {
		return m.Auth
	}
	return nil
}

type ListLabelsRequest struct {
	Regex            *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListLabelsRequest) Reset()                    { *m = ListLabelsRequest{} }
func (m *ListLabelsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLabelsRequest) ProtoMessage()               {}
func (*ListLabelsRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{8} }

func (m *ListLabelsRequest) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

type ListLabelsResponse struct {
	Label            [][]byte `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ListLabelsResponse) Reset()                    { *m = ListLabelsResponse{} }
func (m *ListLabelsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLabelsResponse) ProtoMessage()               {}
func (*ListLabelsResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{9} }

func (m *ListLabelsResponse) GetLabel() [][]byte {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*VisibilityLabelsRequest)(nil), "pb.VisibilityLabelsRequest")
	proto.RegisterType((*VisibilityLabel)(nil), "pb.VisibilityLabel")
	proto.RegisterType((*VisibilityLabelsResponse)(nil), "pb.VisibilityLabelsResponse")
	proto.RegisterType((*SetAuthsRequest)(nil), "pb.SetAuthsRequest")
	proto.RegisterType((*UserAuthorizations)(nil), "pb.UserAuthorizations")
	proto.RegisterType((*MultiUserAuthorizations)(nil), "pb.MultiUserAuthorizations")
	proto.RegisterType((*GetAuthsRequest)(nil), "pb.GetAuthsRequest")
	proto.RegisterType((*GetAuthsResponse)(nil), "pb.GetAuthsResponse")
	proto.RegisterType((*ListLabelsRequest)(nil), "pb.ListLabelsRequest")
	proto.RegisterType((*ListLabelsResponse)(nil), "pb.ListLabelsResponse")
}

var fileDescriptor24 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x26, 0xad, 0x3f, 0xed, 0x71, 0x97, 0x6a, 0x5c, 0xbb, 0xa5, 0x7a, 0xb1, 0x0c, 0x08, 0x75,
	0x2f, 0x22, 0x2c, 0x7a, 0xe1, 0x1f, 0xd2, 0x15, 0x54, 0xa4, 0xa2, 0x64, 0xd1, 0xfb, 0x4c, 0xe7,
	0x38, 0x0d, 0x84, 0xc9, 0x98, 0x64, 0x16, 0xf5, 0x09, 0x7c, 0x0c, 0x5f, 0xca, 0xf7, 0x31, 0xe9,
	0x64, 0xbb, 0xed, 0x4c, 0xa1, 0xb2, 0x37, 0x25, 0xf9, 0xce, 0xf7, 0x9d, 0xf3, 0x9d, 0x9c, 0xd3,
	0x81, 0xe1, 0x57, 0x69, 0x65, 0x2a, 0x95, 0x74, 0x3f, 0x67, 0x22, 0x45, 0x65, 0x59, 0x69, 0xb4,
	0xd3, 0xb4, 0x53, 0xa6, 0xe3, 0xbd, 0x37, 0x4a, 0x62, 0xe1, 0x6a, 0x24, 0xf9, 0x00, 0x87, 0x4d,
	0x2e, 0xc7, 0xef, 0x15, 0x5a, 0x47, 0x1f, 0x43, 0xef, 0x5c, 0xda, 0x25, 0x36, 0x22, 0x47, 0xdd,
	0xc9, 0xad, 0x93, 0xbb, 0xac, 0x4c, 0x59, 0x83, 0xce, 0x57, 0xa4, 0x64, 0x0a, 0x83, 0x46, 0x90,
	0x1e, 0xc0, 0x75, 0x15, 0x13, 0x74, 0x26, 0x7b, 0xbc, 0xbe, 0xd0, 0x11, 0xdc, 0xd4, 0x26, 0x93,
	0x85, 0x50, 0xa3, 0xce, 0x11, 0x99, 0xec, 0xf3, 0x8b, 0xab, 0xb7, 0x33, 0x6a, 0xdb, 0xb1, 0xa5,
	0x2e, 0x2c, 0x52, 0x06, 0x37, 0x0c, 0xda, 0x4a, 0xb9, 0xe8, 0x66, 0x18, 0xdc, 0x70, 0xcc, 0xa5,
	0x2e, 0xa6, 0x73, 0xe7, 0x7f, 0xf9, 0x32, 0xca, 0x23, 0x2b, 0x79, 0x06, 0x83, 0x33, 0x74, 0xd3,
	0xca, 0x2d, 0x56, 0x2d, 0x51, 0xb8, 0x56, 0x59, 0x34, 0xd1, 0xcd, 0xf2, 0x1c, 0x30, 0xe1, 0x39,
	0xde, 0x49, 0x37, 0x60, 0xe1, 0x9c, 0xbc, 0x04, 0xfa, 0xc5, 0xc7, 0x82, 0x56, 0x1b, 0xf9, 0x4b,
	0x84, 0xec, 0x76, 0xa7, 0x7a, 0x3f, 0xaa, 0x3f, 0xc1, 0xe1, 0x47, 0x6f, 0x40, 0x6e, 0x49, 0xf1,
	0x04, 0xfa, 0x55, 0x44, 0xed, 0x7a, 0x1b, 0x6d, 0x2a, 0xbf, 0x24, 0x26, 0x0f, 0x61, 0xf0, 0x6e,
	0x77, 0x27, 0xc9, 0x73, 0xb8, 0x7d, 0x49, 0x8b, 0x8f, 0xf6, 0xbf, 0x1d, 0x3f, 0x82, 0x3b, 0x33,
	0x69, 0xdd, 0xe6, 0x06, 0xf8, 0xe9, 0x19, 0xcc, 0xf1, 0x87, 0x57, 0x93, 0x49, 0x9f, 0xd7, 0x97,
	0xe4, 0x18, 0xe8, 0x3a, 0x35, 0x16, 0x5a, 0x9b, 0x74, 0x77, 0x35, 0xe9, 0x93, 0xbf, 0x9d, 0xf6,
	0x7e, 0x9d, 0xa1, 0x39, 0x97, 0x73, 0xa4, 0x6f, 0xa1, 0x2f, 0xb2, 0xac, 0xc6, 0xe8, 0xfd, 0x2d,
	0xab, 0x75, 0xe1, 0x63, 0xfc, 0x60, 0x7b, 0x30, 0x56, 0x7e, 0x05, 0x3d, 0x1b, 0xdb, 0xa6, 0xcb,
	0x0d, 0x6d, 0x4c, 0x7d, 0x87, 0xfc, 0x35, 0xc0, 0x5c, 0xa1, 0x30, 0x57, 0x4e, 0xf0, 0x14, 0x7a,
	0xf9, 0x46, 0xfd, 0xc6, 0xac, 0xc6, 0x07, 0x9b, 0x60, 0x94, 0xbd, 0x00, 0x50, 0xab, 0x67, 0xa4,
	0xf7, 0x02, 0xa7, 0x35, 0x81, 0xf1, 0xb0, 0x09, 0xd7, 0xe2, 0xd3, 0x19, 0x1c, 0x6b, 0x93, 0x33,
	0x51, 0x8a, 0xf9, 0x02, 0xd9, 0x42, 0x64, 0x5a, 0x97, 0x6c, 0x91, 0x0a, 0x8b, 0xf5, 0xff, 0x3a,
	0xad, 0xbe, 0xb1, 0x1c, 0x0b, 0x34, 0xc2, 0x61, 0x76, 0xda, 0xfa, 0x1c, 0x7c, 0x0e, 0x1c, 0xfb,
	0x9e, 0xfc, 0x26, 0xe4, 0x0f, 0x21, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x69, 0x05, 0x79, 0xa5,
	0x2f, 0x04, 0x00, 0x00,
}