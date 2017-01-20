// Code generated by protoc-gen-go.
// source: common_enum.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏的id
type CommonEnumGame int32

const (
	CommonEnumGame_GID_SRC     CommonEnumGame = 0
	CommonEnumGame_GID_TEXAS   CommonEnumGame = 1
	CommonEnumGame_GID_MAHJONG CommonEnumGame = 2
	CommonEnumGame_GID_DDZ     CommonEnumGame = 3
	CommonEnumGame_GID_ZJH     CommonEnumGame = 4
	CommonEnumGame_GID_HALL    CommonEnumGame = 5
)

var CommonEnumGame_name = map[int32]string{
	0: "GID_SRC",
	1: "GID_TEXAS",
	2: "GID_MAHJONG",
	3: "GID_DDZ",
	4: "GID_ZJH",
	5: "GID_HALL",
}
var CommonEnumGame_value = map[string]int32{
	"GID_SRC":     0,
	"GID_TEXAS":   1,
	"GID_MAHJONG": 2,
	"GID_DDZ":     3,
	"GID_ZJH":     4,
	"GID_HALL":    5,
}

func (x CommonEnumGame) Enum() *CommonEnumGame {
	p := new(CommonEnumGame)
	*p = x
	return p
}
func (x CommonEnumGame) String() string {
	return proto.EnumName(CommonEnumGame_name, int32(x))
}
func (x *CommonEnumGame) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommonEnumGame_value, data, "CommonEnumGame")
	if err != nil {
		return err
	}
	*x = CommonEnumGame(value)
	return nil
}
func (CommonEnumGame) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// 房间的类型
type COMMON_ENUM_ROOMTYPE int32

const (
	COMMON_ENUM_ROOMTYPE_DESK_FRIEND COMMON_ENUM_ROOMTYPE = 1
	COMMON_ENUM_ROOMTYPE_DESK_COIN   COMMON_ENUM_ROOMTYPE = 2
)

var COMMON_ENUM_ROOMTYPE_name = map[int32]string{
	1: "DESK_FRIEND",
	2: "DESK_COIN",
}
var COMMON_ENUM_ROOMTYPE_value = map[string]int32{
	"DESK_FRIEND": 1,
	"DESK_COIN":   2,
}

func (x COMMON_ENUM_ROOMTYPE) Enum() *COMMON_ENUM_ROOMTYPE {
	p := new(COMMON_ENUM_ROOMTYPE)
	*p = x
	return p
}
func (x COMMON_ENUM_ROOMTYPE) String() string {
	return proto.EnumName(COMMON_ENUM_ROOMTYPE_name, int32(x))
}
func (x *COMMON_ENUM_ROOMTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ROOMTYPE_value, data, "COMMON_ENUM_ROOMTYPE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ROOMTYPE(value)
	return nil
}
func (COMMON_ENUM_ROOMTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// 游戏的状态
type COMMON_ENUM_GAMESTATUS int32

const (
	COMMON_ENUM_GAMESTATUS_NOGAME COMMON_ENUM_GAMESTATUS = 1
	COMMON_ENUM_GAMESTATUS_GAMING COMMON_ENUM_GAMESTATUS = 2
)

var COMMON_ENUM_GAMESTATUS_name = map[int32]string{
	1: "NOGAME",
	2: "GAMING",
}
var COMMON_ENUM_GAMESTATUS_value = map[string]int32{
	"NOGAME": 1,
	"GAMING": 2,
}

func (x COMMON_ENUM_GAMESTATUS) Enum() *COMMON_ENUM_GAMESTATUS {
	p := new(COMMON_ENUM_GAMESTATUS)
	*p = x
	return p
}
func (x COMMON_ENUM_GAMESTATUS) String() string {
	return proto.EnumName(COMMON_ENUM_GAMESTATUS_name, int32(x))
}
func (x *COMMON_ENUM_GAMESTATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_GAMESTATUS_value, data, "COMMON_ENUM_GAMESTATUS")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_GAMESTATUS(value)
	return nil
}
func (COMMON_ENUM_GAMESTATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// 发布版本
type COMMON_ENUM_RELEASETAG int32

const (
	COMMON_ENUM_RELEASETAG_R_PRO          COMMON_ENUM_RELEASETAG = 1
	COMMON_ENUM_RELEASETAG_R_APPLE_VERIFY COMMON_ENUM_RELEASETAG = 2
)

var COMMON_ENUM_RELEASETAG_name = map[int32]string{
	1: "R_PRO",
	2: "R_APPLE_VERIFY",
}
var COMMON_ENUM_RELEASETAG_value = map[string]int32{
	"R_PRO":          1,
	"R_APPLE_VERIFY": 2,
}

func (x COMMON_ENUM_RELEASETAG) Enum() *COMMON_ENUM_RELEASETAG {
	p := new(COMMON_ENUM_RELEASETAG)
	*p = x
	return p
}
func (x COMMON_ENUM_RELEASETAG) String() string {
	return proto.EnumName(COMMON_ENUM_RELEASETAG_name, int32(x))
}
func (x *COMMON_ENUM_RELEASETAG) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_RELEASETAG_value, data, "COMMON_ENUM_RELEASETAG")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_RELEASETAG(value)
	return nil
}
func (COMMON_ENUM_RELEASETAG) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// 踢人的踢出类型
type COMMON_ENUM_KICKOUT int32

const (
	COMMON_ENUM_KICKOUT_K_TIMEOUT       COMMON_ENUM_KICKOUT = 1
	COMMON_ENUM_KICKOUT_K_COINUNMATCHED COMMON_ENUM_KICKOUT = 2
)

var COMMON_ENUM_KICKOUT_name = map[int32]string{
	1: "K_TIMEOUT",
	2: "K_COINUNMATCHED",
}
var COMMON_ENUM_KICKOUT_value = map[string]int32{
	"K_TIMEOUT":       1,
	"K_COINUNMATCHED": 2,
}

func (x COMMON_ENUM_KICKOUT) Enum() *COMMON_ENUM_KICKOUT {
	p := new(COMMON_ENUM_KICKOUT)
	*p = x
	return p
}
func (x COMMON_ENUM_KICKOUT) String() string {
	return proto.EnumName(COMMON_ENUM_KICKOUT_name, int32(x))
}
func (x *COMMON_ENUM_KICKOUT) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_KICKOUT_value, data, "COMMON_ENUM_KICKOUT")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_KICKOUT(value)
	return nil
}
func (COMMON_ENUM_KICKOUT) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.CommonEnumGame", CommonEnumGame_name, CommonEnumGame_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ROOMTYPE", COMMON_ENUM_ROOMTYPE_name, COMMON_ENUM_ROOMTYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_GAMESTATUS", COMMON_ENUM_GAMESTATUS_name, COMMON_ENUM_GAMESTATUS_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_RELEASETAG", COMMON_ENUM_RELEASETAG_name, COMMON_ENUM_RELEASETAG_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_KICKOUT", COMMON_ENUM_KICKOUT_name, COMMON_ENUM_KICKOUT_value)
}

var fileDescriptor3 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8e, 0x41, 0x4f, 0xc3, 0x20,
	0x18, 0x86, 0xb5, 0x3a, 0xe7, 0xbe, 0xa9, 0x43, 0x66, 0xfc, 0x11, 0x3d, 0x18, 0x4f, 0x1a, 0x8f,
	0x04, 0xbe, 0xb5, 0xac, 0x05, 0x1a, 0xa0, 0xc6, 0xed, 0x42, 0x8c, 0x5b, 0x76, 0xea, 0x6a, 0x8c,
	0xfe, 0x7f, 0x61, 0x66, 0x49, 0xbd, 0xbd, 0x0f, 0xf0, 0xf0, 0xbe, 0x70, 0xfb, 0xd1, 0x77, 0x5d,
	0xbf, 0x0f, 0xdb, 0xfd, 0x4f, 0xf7, 0xf0, 0xf9, 0xd5, 0x7f, 0xf7, 0x74, 0xbc, 0xd9, 0x1c, 0x42,
	0xbe, 0x03, 0x32, 0xb8, 0x0d, 0xbb, 0xf7, 0x6e, 0x4b, 0xa7, 0x30, 0x2e, 0xa4, 0x08, 0xce, 0x72,
	0x72, 0x42, 0xaf, 0x61, 0x92, 0xc0, 0xe3, 0x1b, 0x73, 0xe4, 0x94, 0xce, 0x60, 0x9a, 0x50, 0xb1,
	0x72, 0x69, 0x74, 0x41, 0xb2, 0xe3, 0x63, 0x21, 0xd6, 0xe4, 0xec, 0x08, 0xeb, 0x65, 0x49, 0xce,
	0xe9, 0x15, 0x5c, 0x26, 0x28, 0x59, 0x5d, 0x93, 0x51, 0xfe, 0x04, 0x77, 0xdc, 0x28, 0x65, 0x74,
	0x40, 0xdd, 0xaa, 0x60, 0x8d, 0x51, 0x7e, 0xd5, 0x60, 0xfa, 0x50, 0xa0, 0xab, 0xc2, 0xc2, 0x4a,
	0xd4, 0x22, 0x36, 0xc4, 0xc2, 0xc3, 0x01, 0x37, 0x52, 0x93, 0x2c, 0x7f, 0x84, 0xfb, 0xa1, 0x57,
	0x30, 0x85, 0xce, 0x33, 0xdf, 0x3a, 0x0a, 0x70, 0xa1, 0x4d, 0xe2, 0x28, 0xc5, 0x1c, 0x93, 0x4c,
	0x8b, 0xf2, 0xe7, 0xff, 0x86, 0xc5, 0x1a, 0x99, 0x43, 0xcf, 0x0a, 0x3a, 0x81, 0x91, 0x0d, 0x8d,
	0x35, 0x51, 0xa0, 0x70, 0x63, 0x03, 0x6b, 0x9a, 0x1a, 0xc3, 0x2b, 0x5a, 0xb9, 0x58, 0x45, 0xf1,
	0x05, 0xe6, 0x43, 0xb1, 0x92, 0xbc, 0x32, 0xad, 0x4f, 0x83, 0xaa, 0xe0, 0xa5, 0xc2, 0x08, 0xd1,
	0x9c, 0xc3, 0xec, 0x6f, 0x5c, 0xab, 0x15, 0xf3, 0xbc, 0x44, 0x41, 0xb2, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0xd2, 0xb5, 0x13, 0x63, 0x01, 0x00, 0x00,
}
