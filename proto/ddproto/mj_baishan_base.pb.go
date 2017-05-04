// Code generated by protoc-gen-go.
// source: mj_baishan_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MjBaishanPaitype int32

const (
	MjBaishanPaitype_PENGPENGHU     MjBaishanPaitype = 1
	MjBaishanPaitype_MENQING_FOUR   MjBaishanPaitype = 2
	MjBaishanPaitype_MENQING_THREE  MjBaishanPaitype = 3
	MjBaishanPaitype_MENQING        MjBaishanPaitype = 4
	MjBaishanPaitype_HUNYISE        MjBaishanPaitype = 5
	MjBaishanPaitype_SHOUBAYI       MjBaishanPaitype = 6
	MjBaishanPaitype_YIBANGAO       MjBaishanPaitype = 7
	MjBaishanPaitype_SIGUIYI        MjBaishanPaitype = 8
	MjBaishanPaitype_QINGYISE       MjBaishanPaitype = 9
	MjBaishanPaitype_SANHUATONGSHUN MjBaishanPaitype = 10
	MjBaishanPaitype_QIDUI          MjBaishanPaitype = 11
	MjBaishanPaitype_SHISANYAO      MjBaishanPaitype = 12
)

var MjBaishanPaitype_name = map[int32]string{
	1:  "PENGPENGHU",
	2:  "MENQING_FOUR",
	3:  "MENQING_THREE",
	4:  "MENQING",
	5:  "HUNYISE",
	6:  "SHOUBAYI",
	7:  "YIBANGAO",
	8:  "SIGUIYI",
	9:  "QINGYISE",
	10: "SANHUATONGSHUN",
	11: "QIDUI",
	12: "SHISANYAO",
}
var MjBaishanPaitype_value = map[string]int32{
	"PENGPENGHU":     1,
	"MENQING_FOUR":   2,
	"MENQING_THREE":  3,
	"MENQING":        4,
	"HUNYISE":        5,
	"SHOUBAYI":       6,
	"YIBANGAO":       7,
	"SIGUIYI":        8,
	"QINGYISE":       9,
	"SANHUATONGSHUN": 10,
	"QIDUI":          11,
	"SHISANYAO":      12,
}

func (x MjBaishanPaitype) Enum() *MjBaishanPaitype {
	p := new(MjBaishanPaitype)
	*p = x
	return p
}
func (x MjBaishanPaitype) String() string {
	return proto.EnumName(MjBaishanPaitype_name, int32(x))
}
func (x *MjBaishanPaitype) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjBaishanPaitype_value, data, "MjBaishanPaitype")
	if err != nil {
		return err
	}
	*x = MjBaishanPaitype(value)
	return nil
}
func (MjBaishanPaitype) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

type MjBaishanHutype int32

const (
	MjBaishanHutype_BAOTING            MjBaishanHutype = 1
	MjBaishanHutype_ZIMO               MjBaishanHutype = 2
	MjBaishanHutype_BANKER             MjBaishanHutype = 3
	MjBaishanHutype_GANGHOUKAI_HU      MjBaishanHutype = 4
	MjBaishanHutype_GANGHOUKAI_DIANPAO MjBaishanHutype = 5
)

var MjBaishanHutype_name = map[int32]string{
	1: "BAOTING",
	2: "ZIMO",
	3: "BANKER",
	4: "GANGHOUKAI_HU",
	5: "GANGHOUKAI_DIANPAO",
}
var MjBaishanHutype_value = map[string]int32{
	"BAOTING":            1,
	"ZIMO":               2,
	"BANKER":             3,
	"GANGHOUKAI_HU":      4,
	"GANGHOUKAI_DIANPAO": 5,
}

func (x MjBaishanHutype) Enum() *MjBaishanHutype {
	p := new(MjBaishanHutype)
	*p = x
	return p
}
func (x MjBaishanHutype) String() string {
	return proto.EnumName(MjBaishanHutype_name, int32(x))
}
func (x *MjBaishanHutype) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MjBaishanHutype_value, data, "MjBaishanHutype")
	if err != nil {
		return err
	}
	*x = MjBaishanHutype(value)
	return nil
}
func (MjBaishanHutype) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

func init() {
	proto.RegisterEnum("ddproto.MjBaishanPaitype", MjBaishanPaitype_name, MjBaishanPaitype_value)
	proto.RegisterEnum("ddproto.MjBaishanHutype", MjBaishanHutype_name, MjBaishanHutype_value)
}

func init() { proto.RegisterFile("mj_baishan_base.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xdd, 0x6a, 0x3a, 0x41,
	0x0c, 0xc5, 0xf1, 0x73, 0x35, 0x7e, 0x10, 0x03, 0xff, 0xff, 0x43, 0x78, 0xd1, 0x77, 0xc8, 0xe2,
	0x74, 0x26, 0x88, 0x19, 0x75, 0x9c, 0x8b, 0xed, 0xcd, 0xb2, 0xa2, 0xa0, 0x85, 0xd6, 0xa5, 0xda,
	0x8b, 0xbe, 0x61, 0x1f, 0xab, 0xcc, 0xd2, 0xc2, 0x5e, 0x04, 0x72, 0x4e, 0xce, 0x81, 0xfc, 0xe0,
	0xdf, 0xdb, 0x6b, 0x79, 0xac, 0xae, 0xf7, 0x4b, 0xf5, 0x5e, 0x1e, 0xab, 0xfb, 0xf9, 0xa9, 0xfe,
	0xb8, 0x3d, 0x6e, 0x94, 0x9d, 0x4e, 0xcd, 0xb2, 0xfc, 0xee, 0x00, 0xb5, 0x22, 0x75, 0x75, 0x7d,
	0x7c, 0xd5, 0x67, 0x9a, 0x03, 0x6c, 0x8d, 0xda, 0x34, 0x2e, 0x62, 0x87, 0x10, 0xa6, 0x1b, 0xa3,
	0x3b, 0x51, 0x5b, 0x3e, 0xfb, 0xb8, 0xc7, 0x2e, 0x2d, 0x60, 0xf6, 0xe7, 0x1c, 0xdc, 0xde, 0x18,
	0xec, 0xd1, 0x04, 0xb2, 0x5f, 0x0b, 0xfb, 0x49, 0xb8, 0xa8, 0x85, 0x04, 0x83, 0x03, 0x9a, 0xc2,
	0x28, 0x38, 0x1f, 0x73, 0x2e, 0x04, 0x87, 0x49, 0x15, 0x92, 0xb3, 0x5a, 0xf6, 0x98, 0xa5, 0x60,
	0x10, 0x1b, 0xa5, 0x10, 0x1c, 0xa5, 0x53, 0xea, 0x37, 0xb5, 0x31, 0x11, 0xcc, 0x03, 0xab, 0x8b,
	0x7c, 0xf0, 0x6a, 0x83, 0x8b, 0x8a, 0x40, 0x63, 0x18, 0xec, 0x64, 0x15, 0x05, 0x27, 0x34, 0x83,
	0x71, 0x70, 0x12, 0x58, 0x0b, 0xf6, 0x38, 0x5d, 0x56, 0xb0, 0x68, 0x91, 0x5c, 0x3e, 0x1b, 0x90,
	0x09, 0x64, 0x39, 0xfb, 0x43, 0xfa, 0xa9, 0x43, 0x23, 0xe8, 0xbf, 0xc8, 0xc6, 0x63, 0x97, 0x00,
	0x86, 0x39, 0xeb, 0xda, 0xec, 0xb1, 0x97, 0x48, 0x2c, 0xab, 0x75, 0x3e, 0xae, 0x59, 0x4a, 0x17,
	0xb1, 0x4f, 0xff, 0x81, 0x5a, 0xd6, 0x4a, 0x58, 0xb7, 0xec, 0x71, 0xf0, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x13, 0xc1, 0x41, 0x4e, 0x01, 0x00, 0x00,
}
