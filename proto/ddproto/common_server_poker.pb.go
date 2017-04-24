// Code generated by protoc-gen-go.
// source: common_server_poker.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 牌 ,单张扑克牌
type CommonSrvPokerPai struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Des              *string `protobuf:"bytes,2,opt,name=des" json:"des,omitempty"`
	Value            *int32  `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
	Flower           *int32  `protobuf:"varint,4,opt,name=flower" json:"flower,omitempty"`
	Name             *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvPokerPai) Reset()                    { *m = CommonSrvPokerPai{} }
func (m *CommonSrvPokerPai) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvPokerPai) ProtoMessage()               {}
func (*CommonSrvPokerPai) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *CommonSrvPokerPai) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *CommonSrvPokerPai) GetDes() string {
	if m != nil && m.Des != nil {
		return *m.Des
	}
	return ""
}

func (m *CommonSrvPokerPai) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *CommonSrvPokerPai) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

func (m *CommonSrvPokerPai) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonSrvPokerPai)(nil), "ddproto.common_srv_pokerPai")
}

var fileDescriptor10 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x2f, 0xc8, 0xcf, 0x4e, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0x94, 0xa2, 0xb9, 0x84, 0x61,
	0xaa, 0x8a, 0xca, 0x20, 0x4a, 0x02, 0x12, 0x33, 0x85, 0xb8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x85, 0xb8, 0xb9, 0x98, 0x53, 0x52, 0x8b, 0x25, 0x98, 0x80, 0x1c, 0x4e,
	0x21, 0x5e, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x66, 0xb0, 0x1c, 0x1f, 0x17, 0x5b,
	0x5a, 0x4e, 0x7e, 0x79, 0x6a, 0x91, 0x04, 0x0b, 0x98, 0xcf, 0xc3, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0x0a, 0x52, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x40, 0xcb, 0x80, 0xc5, 0x81, 0x00,
	0x00, 0x00,
}
