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
	CommonEnumGame_GID_SRC                    CommonEnumGame = 0
	CommonEnumGame_GID_TEXAS                  CommonEnumGame = 1
	CommonEnumGame_GID_MAHJONG                CommonEnumGame = 2
	CommonEnumGame_GID_DDZ                    CommonEnumGame = 3
	CommonEnumGame_GID_ZJH                    CommonEnumGame = 4
	CommonEnumGame_GID_HALL                   CommonEnumGame = 5
	CommonEnumGame_GID_PDK                    CommonEnumGame = 7
	CommonEnumGame_GID_ZXZ                    CommonEnumGame = 8
	CommonEnumGame_GID_MJBAISHAN              CommonEnumGame = 9
	CommonEnumGame_GID_MJCHANGCHUN            CommonEnumGame = 10
	CommonEnumGame_GID_NIUNIUJINGDIAN         CommonEnumGame = 11
	CommonEnumGame_GID_BANTUOZI               CommonEnumGame = 12
	CommonEnumGame_GID_DOUDIZHUERREN          CommonEnumGame = 13
	CommonEnumGame_GIDDOUDIZHUJINGDIANDONGBEI CommonEnumGame = 14
	CommonEnumGame_GID_TIANDAKENG             CommonEnumGame = 15
	CommonEnumGame_GID_MJ_SONGJIANGHE         CommonEnumGame = 16
	CommonEnumGame_GID_BAIRENNIUNIU           CommonEnumGame = 17
	CommonEnumGame_GID_ZHIZUNWUZHANG          CommonEnumGame = 18
	CommonEnumGame_GID_MJ_YIBIN               CommonEnumGame = 19
	CommonEnumGame_GID_MJ_PENGZHOU            CommonEnumGame = 20
)

var CommonEnumGame_name = map[int32]string{
	0:  "GID_SRC",
	1:  "GID_TEXAS",
	2:  "GID_MAHJONG",
	3:  "GID_DDZ",
	4:  "GID_ZJH",
	5:  "GID_HALL",
	7:  "GID_PDK",
	8:  "GID_ZXZ",
	9:  "GID_MJBAISHAN",
	10: "GID_MJCHANGCHUN",
	11: "GID_NIUNIUJINGDIAN",
	12: "GID_BANTUOZI",
	13: "GID_DOUDIZHUERREN",
	14: "GIDDOUDIZHUJINGDIANDONGBEI",
	15: "GID_TIANDAKENG",
	16: "GID_MJ_SONGJIANGHE",
	17: "GID_BAIRENNIUNIU",
	18: "GID_ZHIZUNWUZHANG",
	19: "GID_MJ_YIBIN",
	20: "GID_MJ_PENGZHOU",
}
var CommonEnumGame_value = map[string]int32{
	"GID_SRC":                    0,
	"GID_TEXAS":                  1,
	"GID_MAHJONG":                2,
	"GID_DDZ":                    3,
	"GID_ZJH":                    4,
	"GID_HALL":                   5,
	"GID_PDK":                    7,
	"GID_ZXZ":                    8,
	"GID_MJBAISHAN":              9,
	"GID_MJCHANGCHUN":            10,
	"GID_NIUNIUJINGDIAN":         11,
	"GID_BANTUOZI":               12,
	"GID_DOUDIZHUERREN":          13,
	"GIDDOUDIZHUJINGDIANDONGBEI": 14,
	"GID_TIANDAKENG":             15,
	"GID_MJ_SONGJIANGHE":         16,
	"GID_BAIRENNIUNIU":           17,
	"GID_ZHIZUNWUZHANG":          18,
	"GID_MJ_YIBIN":               19,
	"GID_MJ_PENGZHOU":            20,
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
func (CommonEnumGame) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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
func (COMMON_ENUM_ROOMTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

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
func (COMMON_ENUM_GAMESTATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

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
func (COMMON_ENUM_RELEASETAG) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// 踢人的踢出类型
type COMMON_ENUM_KICKOUT int32

const (
	COMMON_ENUM_KICKOUT_K_COINUNMATCHED              COMMON_ENUM_KICKOUT = 1
	COMMON_ENUM_KICKOUT_K_SYSTEM                     COMMON_ENUM_KICKOUT = 2
	COMMON_ENUM_KICKOUT_K_OFFLINE                    COMMON_ENUM_KICKOUT = 3
	COMMON_ENUM_KICKOUT_K_GAME_BLOCKED               COMMON_ENUM_KICKOUT = 4
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_LOTTERY   COMMON_ENUM_KICKOUT = 5
	COMMON_ENUM_KICKOUT_K_TIMEOUT_NOTREADY_ENTERDESK COMMON_ENUM_KICKOUT = 6
)

var COMMON_ENUM_KICKOUT_name = map[int32]string{
	1: "K_COINUNMATCHED",
	2: "K_SYSTEM",
	3: "K_OFFLINE",
	4: "K_GAME_BLOCKED",
	5: "K_TIMEOUT_NOTREADY_LOTTERY",
	6: "K_TIMEOUT_NOTREADY_ENTERDESK",
}
var COMMON_ENUM_KICKOUT_value = map[string]int32{
	"K_COINUNMATCHED":              1,
	"K_SYSTEM":                     2,
	"K_OFFLINE":                    3,
	"K_GAME_BLOCKED":               4,
	"K_TIMEOUT_NOTREADY_LOTTERY":   5,
	"K_TIMEOUT_NOTREADY_ENTERDESK": 6,
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
func (COMMON_ENUM_KICKOUT) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

// 玩家申请解散房间的操作类型
type COMMON_ENUM_APPLYDISSOLVE int32

const (
	COMMON_ENUM_APPLYDISSOLVE_AD_AGREE  COMMON_ENUM_APPLYDISSOLVE = 1
	COMMON_ENUM_APPLYDISSOLVE_AD_REFUSE COMMON_ENUM_APPLYDISSOLVE = 2
	COMMON_ENUM_APPLYDISSOLVE_AD_NOACT  COMMON_ENUM_APPLYDISSOLVE = 3
)

var COMMON_ENUM_APPLYDISSOLVE_name = map[int32]string{
	1: "AD_AGREE",
	2: "AD_REFUSE",
	3: "AD_NOACT",
}
var COMMON_ENUM_APPLYDISSOLVE_value = map[string]int32{
	"AD_AGREE":  1,
	"AD_REFUSE": 2,
	"AD_NOACT":  3,
}

func (x COMMON_ENUM_APPLYDISSOLVE) Enum() *COMMON_ENUM_APPLYDISSOLVE {
	p := new(COMMON_ENUM_APPLYDISSOLVE)
	*p = x
	return p
}
func (x COMMON_ENUM_APPLYDISSOLVE) String() string {
	return proto.EnumName(COMMON_ENUM_APPLYDISSOLVE_name, int32(x))
}
func (x *COMMON_ENUM_APPLYDISSOLVE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_APPLYDISSOLVE_value, data, "COMMON_ENUM_APPLYDISSOLVE")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_APPLYDISSOLVE(value)
	return nil
}
func (COMMON_ENUM_APPLYDISSOLVE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

// 按钮的类型
type BTN_TYPE int32

const (
	BTN_TYPE_BTNTYPE_NEWUSERBUTTON     BTN_TYPE = 1
	BTN_TYPE_BTNTYPE_SHARELINKTIMELINE BTN_TYPE = 2
	BTN_TYPE_REDBAGSHAR                BTN_TYPE = 3
)

var BTN_TYPE_name = map[int32]string{
	1: "BTNTYPE_NEWUSERBUTTON",
	2: "BTNTYPE_SHARELINKTIMELINE",
	3: "REDBAGSHAR",
}
var BTN_TYPE_value = map[string]int32{
	"BTNTYPE_NEWUSERBUTTON":     1,
	"BTNTYPE_SHARELINKTIMELINE": 2,
	"REDBAGSHAR":                3,
}

func (x BTN_TYPE) Enum() *BTN_TYPE {
	p := new(BTN_TYPE)
	*p = x
	return p
}
func (x BTN_TYPE) String() string {
	return proto.EnumName(BTN_TYPE_name, int32(x))
}
func (x *BTN_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BTN_TYPE_value, data, "BTN_TYPE")
	if err != nil {
		return err
	}
	*x = BTN_TYPE(value)
	return nil
}
func (BTN_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

// 进入金币场房间的错误类型
type COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM int32

const (
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_MIN       COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 1
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_OVER_LIMIT      COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 2
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_COIN_UNDER_LIMIT     COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 3
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER_LV_DESK_GAMING COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 4
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_NONE_DESK_AVAILABLE  COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 5
	COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER                COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM = 6
)

var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name = map[int32]string{
	1: "ERROR_EC_COIN_UNDER_MIN",
	2: "ERROR_EC_COIN_OVER_LIMIT",
	3: "ERROR_EC_COIN_UNDER_LIMIT",
	4: "ERROR_EC_OTHER_LV_DESK_GAMING",
	5: "ERROR_EC_NONE_DESK_AVAILABLE",
	6: "ERROR_EC_OTHER",
}
var COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value = map[string]int32{
	"ERROR_EC_COIN_UNDER_MIN":       1,
	"ERROR_EC_COIN_OVER_LIMIT":      2,
	"ERROR_EC_COIN_UNDER_LIMIT":     3,
	"ERROR_EC_OTHER_LV_DESK_GAMING": 4,
	"ERROR_EC_NONE_DESK_AVAILABLE":  5,
	"ERROR_EC_OTHER":                6,
}

func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) Enum() *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM {
	p := new(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM)
	*p = x
	return p
}
func (x COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) String() string {
	return proto.EnumName(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, int32(x))
}
func (x *COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value, data, "COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM")
	if err != nil {
		return err
	}
	*x = COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM(value)
	return nil
}
func (COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{7}
}

func init() {
	proto.RegisterEnum("ddproto.CommonEnumGame", CommonEnumGame_name, CommonEnumGame_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ROOMTYPE", COMMON_ENUM_ROOMTYPE_name, COMMON_ENUM_ROOMTYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_GAMESTATUS", COMMON_ENUM_GAMESTATUS_name, COMMON_ENUM_GAMESTATUS_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_RELEASETAG", COMMON_ENUM_RELEASETAG_name, COMMON_ENUM_RELEASETAG_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_KICKOUT", COMMON_ENUM_KICKOUT_name, COMMON_ENUM_KICKOUT_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_APPLYDISSOLVE", COMMON_ENUM_APPLYDISSOLVE_name, COMMON_ENUM_APPLYDISSOLVE_value)
	proto.RegisterEnum("ddproto.BTN_TYPE", BTN_TYPE_name, BTN_TYPE_value)
	proto.RegisterEnum("ddproto.COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM", COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_name, COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_value)
}

var fileDescriptor6 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0xbd, 0x89, 0xf3, 0x9c, 0xbc, 0x68, 0xe6, 0x71, 0x6f, 0x6e, 0x93, 0xa2, 0x05, 0xba, 0xf2,
	0xa2, 0xe8, 0xaa, 0x5d, 0x53, 0xe2, 0x58, 0xa2, 0x25, 0x91, 0x06, 0x45, 0x3a, 0x91, 0x36, 0x44,
	0xd1, 0x04, 0x5d, 0x39, 0x2e, 0x8a, 0xf6, 0x27, 0xfa, 0x0f, 0xfd, 0xb0, 0xfe, 0x4d, 0x87, 0x52,
	0x95, 0x3a, 0x40, 0x00, 0x2e, 0x38, 0x73, 0xe6, 0x71, 0xe6, 0xcc, 0xc0, 0xf8, 0xd3, 0x6a, 0xb9,
	0x5c, 0x3d, 0x84, 0xfb, 0x87, 0xef, 0xcb, 0xb7, 0x5f, 0xbe, 0xae, 0xbe, 0xad, 0xf8, 0xee, 0xdd,
	0x5d, 0xf7, 0x99, 0xfc, 0x18, 0x01, 0x5b, 0x83, 0xc3, 0xe7, 0x8f, 0xcb, 0x7b, 0x7e, 0x00, 0xbb,
	0x99, 0x92, 0xa1, 0xb6, 0x29, 0xfb, 0x87, 0x1f, 0xc1, 0x7e, 0x34, 0x1c, 0xde, 0x8a, 0x9a, 0x6d,
	0xf0, 0x13, 0x38, 0x88, 0x66, 0x25, 0xf2, 0x99, 0xd1, 0x19, 0xdb, 0x1c, 0x82, 0xa5, 0x6c, 0xd9,
	0x68, 0x30, 0xda, 0x59, 0xce, 0xb6, 0xf8, 0x21, 0xec, 0x45, 0x23, 0x17, 0x65, 0xc9, 0xb6, 0x07,
	0x68, 0x2e, 0x0b, 0xb6, 0xfb, 0x18, 0x77, 0xdb, 0xb2, 0x3d, 0x3e, 0x86, 0xa3, 0xae, 0xe4, 0x2c,
	0x11, 0xaa, 0xce, 0x85, 0x66, 0xfb, 0xfc, 0x14, 0x4e, 0x7a, 0x57, 0x4a, 0x76, 0x96, 0xe6, 0x5e,
	0x33, 0xe0, 0x17, 0xc0, 0xa3, 0x53, 0x2b, 0x4f, 0x6f, 0xa6, 0x74, 0x26, 0x15, 0x05, 0x1f, 0x70,
	0x06, 0x87, 0xd1, 0x9f, 0x08, 0xed, 0xbc, 0x69, 0x15, 0x3b, 0xe4, 0xe7, 0x30, 0xee, 0x38, 0x19,
	0x2f, 0x55, 0x9b, 0x7b, 0xb4, 0x16, 0x35, 0x3b, 0xe2, 0x2f, 0xe1, 0x7f, 0x72, 0x0f, 0xde, 0xa1,
	0x82, 0xa4, 0x39, 0x12, 0x54, 0xec, 0x98, 0x73, 0x38, 0xee, 0x46, 0x8d, 0x4e, 0x51, 0x20, 0x8d,
	0x77, 0x32, 0x34, 0xad, 0x66, 0xa1, 0xa6, 0xb8, 0x19, 0x41, 0x59, 0x8e, 0x8c, 0xf1, 0x33, 0x60,
	0x7d, 0x53, 0x45, 0xb5, 0x7b, 0x4a, 0x6c, 0x3c, 0x34, 0x6e, 0x73, 0xd5, 0x7a, 0x7d, 0xe3, 0xdb,
	0xc8, 0x9e, 0xf1, 0x81, 0x21, 0x15, 0x69, 0x54, 0xa2, 0x34, 0x3b, 0xfd, 0x3b, 0x60, 0x98, 0x53,
	0x9f, 0x36, 0x37, 0x9e, 0x9d, 0x4d, 0xde, 0xc3, 0x59, 0x6a, 0xaa, 0xca, 0xe8, 0x80, 0xda, 0x57,
	0xc1, 0x1a, 0x53, 0xb9, 0x66, 0x8e, 0x51, 0x73, 0x89, 0x75, 0x11, 0xa6, 0x56, 0xa1, 0x96, 0xb4,
	0x04, 0xda, 0x49, 0xe7, 0x48, 0x0d, 0x15, 0xdb, 0x9c, 0xbc, 0x83, 0x8b, 0xf5, 0xbc, 0x4c, 0x54,
	0x58, 0x3b, 0xe1, 0x7c, 0xcd, 0x01, 0x76, 0xb4, 0x89, 0x36, 0x25, 0xd1, 0x9f, 0x7e, 0x2a, 0x2e,
	0x6d, 0xf2, 0xe1, 0x69, 0x86, 0xc5, 0x12, 0x45, 0x8d, 0x4e, 0x64, 0x7c, 0x1f, 0xb6, 0x6d, 0x98,
	0x5b, 0x43, 0x09, 0x24, 0x87, 0x0d, 0x62, 0x3e, 0x2f, 0x31, 0x2c, 0xd0, 0xaa, 0x69, 0x43, 0x89,
	0x3f, 0x37, 0xe0, 0x74, 0x3d, 0xb3, 0x50, 0x69, 0x61, 0xbc, 0x8b, 0xf3, 0xf4, 0x74, 0xbc, 0xae,
	0x84, 0x4b, 0x73, 0x8c, 0x34, 0xe9, 0x00, 0x8a, 0x50, 0x37, 0xb5, 0xc3, 0x8a, 0x0e, 0x85, 0x48,
	0x17, 0xc1, 0x4c, 0xa7, 0xa5, 0xd2, 0x48, 0xa7, 0x42, 0xd5, 0x8b, 0x8e, 0x6a, 0x48, 0x4a, 0x93,
	0x16, 0x94, 0xb0, 0x15, 0x17, 0x54, 0x90, 0xfc, 0x15, 0x52, 0xc9, 0xa0, 0x8d, 0xb3, 0x28, 0x64,
	0x13, 0x4a, 0xe3, 0x1c, 0xda, 0x86, 0x6e, 0xe8, 0x15, 0x5c, 0x3d, 0x83, 0xa3, 0x26, 0x38, 0xea,
	0xc1, 0x76, 0x26, 0x53, 0xb8, 0x5c, 0xa7, 0x17, 0xd9, 0x37, 0x52, 0xd5, 0xb5, 0x29, 0x17, 0x18,
	0xf9, 0x08, 0x19, 0x44, 0x66, 0x11, 0x7b, 0x11, 0xc9, 0xb2, 0x38, 0xf5, 0x35, 0x12, 0xbd, 0x1e,
	0xd4, 0x46, 0xa4, 0x8e, 0x8d, 0x26, 0x0e, 0xf6, 0x12, 0xa7, 0x43, 0x27, 0xff, 0x25, 0x9c, 0xd3,
	0x3f, 0x7e, 0x83, 0xc6, 0x1b, 0x8a, 0xb6, 0x89, 0x77, 0xce, 0x68, 0xaa, 0x71, 0x0d, 0x97, 0x03,
	0x44, 0x97, 0x4b, 0x32, 0x2a, 0x5d, 0x44, 0x7e, 0xdd, 0x8c, 0x9b, 0xfc, 0x18, 0xc0, 0xa2, 0x4c,
	0x44, 0x16, 0x41, 0xaa, 0xfa, 0x6b, 0x03, 0xde, 0xac, 0xd3, 0xa3, 0xc3, 0x34, 0xb6, 0x6b, 0xd2,
	0x0f, 0x11, 0x05, 0x8c, 0x5b, 0xe7, 0x2f, 0xe0, 0xdf, 0x1e, 0xc3, 0xb4, 0x53, 0x35, 0x78, 0x2d,
	0xd1, 0x06, 0xda, 0x1e, 0x35, 0xbd, 0x82, 0xff, 0x9e, 0x82, 0x86, 0xd6, 0x13, 0x4a, 0x55, 0x29,
	0x47, 0x3d, 0x89, 0xd2, 0x73, 0xa9, 0x3d, 0x3c, 0xe2, 0xaf, 0xe1, 0xfa, 0x11, 0x36, 0x2e, 0x8f,
	0xc8, 0x22, 0x74, 0xc7, 0xf4, 0xe7, 0x38, 0xb6, 0xa2, 0xca, 0x8f, 0x21, 0xda, 0x68, 0xec, 0x61,
	0xb1, 0x10, 0xaa, 0x14, 0x49, 0x89, 0xb4, 0x07, 0xda, 0xdd, 0xd3, 0x22, 0x6c, 0xe7, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x68, 0x0a, 0xa8, 0x66, 0x04, 0x00, 0x00,
}
