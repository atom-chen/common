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
	CommonEnumGame_GID_SANDAYI                CommonEnumGame = 21
	CommonEnumGame_GID_ZHUANZHUAN             CommonEnumGame = 22
	CommonEnumGame_GID_PAOHUZI                CommonEnumGame = 23
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
	21: "GID_SANDAYI",
	22: "GID_ZHUANZHUAN",
	23: "GID_PAOHUZI",
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
	"GID_SANDAYI":                21,
	"GID_ZHUANZHUAN":             22,
	"GID_PAOHUZI":                23,
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
	COMMON_ENUM_KICKOUT_K_OWNER                      COMMON_ENUM_KICKOUT = 7
)

var COMMON_ENUM_KICKOUT_name = map[int32]string{
	1: "K_COINUNMATCHED",
	2: "K_SYSTEM",
	3: "K_OFFLINE",
	4: "K_GAME_BLOCKED",
	5: "K_TIMEOUT_NOTREADY_LOTTERY",
	6: "K_TIMEOUT_NOTREADY_ENTERDESK",
	7: "K_OWNER",
}
var COMMON_ENUM_KICKOUT_value = map[string]int32{
	"K_COINUNMATCHED":              1,
	"K_SYSTEM":                     2,
	"K_OFFLINE":                    3,
	"K_GAME_BLOCKED":               4,
	"K_TIMEOUT_NOTREADY_LOTTERY":   5,
	"K_TIMEOUT_NOTREADY_ENTERDESK": 6,
	"K_OWNER":                      7,
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
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x8e, 0xe4, 0x44,
	0x0c, 0xc7, 0xe9, 0xee, 0xf9, 0xf4, 0x7c, 0x79, 0x6a, 0x3e, 0x96, 0x81, 0x5d, 0x04, 0x12, 0xa7,
	0x3e, 0x20, 0x4e, 0x70, 0xae, 0xa4, 0xdc, 0x49, 0x75, 0x12, 0x57, 0xab, 0x52, 0xd5, 0xb3, 0xc9,
	0xa5, 0x84, 0xd8, 0x15, 0xa7, 0xd9, 0x46, 0x08, 0x5e, 0x8a, 0xa7, 0xe0, 0x35, 0x78, 0x1b, 0xe4,
	0x64, 0x33, 0xf4, 0x48, 0x73, 0x89, 0x62, 0xff, 0x5d, 0xf6, 0xcf, 0x1f, 0x70, 0xfd, 0xeb, 0xee,
	0xe9, 0x69, 0xf7, 0x29, 0x7d, 0xfc, 0xf4, 0xd7, 0xd3, 0x0f, 0xbf, 0xff, 0xb1, 0xfb, 0x73, 0xa7,
	0x8e, 0x3f, 0x7c, 0x18, 0x7e, 0x96, 0xff, 0x2c, 0x00, 0xf7, 0xe4, 0xf4, 0xdb, 0x2f, 0x4f, 0x1f,
	0xd5, 0x19, 0x1c, 0x17, 0xd6, 0xa4, 0xd6, 0xe7, 0xf8, 0x85, 0xba, 0x80, 0x53, 0x31, 0x02, 0xbd,
	0xd7, 0x2d, 0xce, 0xd4, 0x15, 0x9c, 0x89, 0xd9, 0xe8, 0x72, 0xed, 0xb8, 0xc0, 0xf9, 0x14, 0x6c,
	0x4c, 0x8f, 0x8b, 0xc9, 0xe8, 0xd7, 0x25, 0x1e, 0xa8, 0x73, 0x38, 0x11, 0xa3, 0xd4, 0x75, 0x8d,
	0x87, 0x93, 0xb4, 0x31, 0x15, 0x1e, 0x3f, 0xc7, 0xbd, 0xef, 0xf1, 0x44, 0x5d, 0xc3, 0xc5, 0x90,
	0x72, 0x9d, 0x69, 0xdb, 0x96, 0x9a, 0xf1, 0x54, 0xdd, 0xc0, 0xd5, 0xe8, 0xca, 0x4b, 0xcd, 0x45,
	0x5e, 0x46, 0x46, 0x50, 0xf7, 0xa0, 0xc4, 0xc9, 0x36, 0xb2, 0x8d, 0x6b, 0xcb, 0x85, 0xb1, 0x9a,
	0xf1, 0x4c, 0x21, 0x9c, 0x8b, 0x3f, 0xd3, 0x1c, 0xa2, 0xeb, 0x2d, 0x9e, 0xab, 0x3b, 0xb8, 0x1e,
	0x98, 0x5c, 0x34, 0xb6, 0x2f, 0x23, 0x79, 0x4f, 0x8c, 0x17, 0xea, 0x1b, 0xf8, 0xaa, 0xb0, 0x66,
	0xf2, 0x4e, 0x19, 0x8c, 0xe3, 0x22, 0x23, 0x8b, 0x97, 0x4a, 0xc1, 0xe5, 0xd0, 0xaa, 0x38, 0x75,
	0x45, 0x5c, 0xe0, 0xd5, 0x54, 0xb4, 0x59, 0xa7, 0xd6, 0x71, 0xb1, 0xb6, 0x9a, 0x8b, 0x92, 0x10,
	0xd5, 0x2d, 0xe0, 0x58, 0xd4, 0x7a, 0xe2, 0x11, 0x09, 0xaf, 0xa7, 0xc2, 0x7d, 0x69, 0xfb, 0xc8,
	0x8f, 0xb1, 0x17, 0x7a, 0x54, 0x13, 0x61, 0xb3, 0x4e, 0x9d, 0xcd, 0x2c, 0xe3, 0xcd, 0xff, 0x0d,
	0xa6, 0x0d, 0x71, 0xd1, 0x97, 0x2e, 0xe2, 0xed, 0x34, 0xdb, 0x56, 0xca, 0x77, 0x16, 0xef, 0x26,
	0xa0, 0xbe, 0x8c, 0x9a, 0x87, 0x0f, 0xde, 0x4f, 0x41, 0x1b, 0xed, 0xca, 0xd8, 0x5b, 0x7c, 0xb3,
	0xfc, 0x09, 0x6e, 0x73, 0xd7, 0x34, 0x8e, 0x13, 0x71, 0x6c, 0x92, 0x77, 0xae, 0x09, 0xdd, 0x86,
	0x24, 0xd0, 0x50, 0x5b, 0xa5, 0x95, 0xb7, 0xc4, 0x06, 0x67, 0xb2, 0xc9, 0xc1, 0x91, 0x3b, 0xcb,
	0x38, 0x5f, 0xfe, 0x08, 0xf7, 0xfb, 0xef, 0x0a, 0xdd, 0x50, 0x1b, 0x74, 0x88, 0xad, 0x02, 0x38,
	0x62, 0x27, 0x36, 0xce, 0xe4, 0xbf, 0xd0, 0x8d, 0x95, 0x55, 0x2f, 0x7f, 0x7e, 0xf9, 0xc2, 0x53,
	0x4d, 0xba, 0xa5, 0xa0, 0x0b, 0x75, 0x0a, 0x87, 0x3e, 0x6d, 0xbc, 0xc3, 0x99, 0x30, 0xfb, 0xa4,
	0x37, 0x9b, 0x9a, 0xd2, 0x96, 0xbc, 0x5d, 0x75, 0x38, 0x5f, 0xfe, 0x3d, 0x83, 0x9b, 0xfd, 0x97,
	0x95, 0xcd, 0x2b, 0x17, 0x83, 0x4c, 0x61, 0xc4, 0x89, 0xdc, 0xe8, 0x90, 0x97, 0x24, 0x98, 0xe7,
	0x70, 0x52, 0xa5, 0xb6, 0x6b, 0x03, 0x35, 0x38, 0x17, 0xe8, 0x2a, 0xb9, 0xd5, 0xaa, 0xb6, 0x4c,
	0xb8, 0x90, 0xec, 0xd5, 0x80, 0x9a, 0xb2, 0xda, 0xe5, 0x15, 0x19, 0x3c, 0x90, 0xb5, 0x56, 0x29,
	0xd8, 0x86, 0x5c, 0x0c, 0x89, 0x5d, 0xf0, 0xa4, 0x4d, 0x97, 0x6a, 0x17, 0x02, 0xf9, 0x0e, 0x0f,
	0xd5, 0xb7, 0xf0, 0xf6, 0x15, 0x9d, 0x38, 0x90, 0x97, 0x79, 0xe0, 0x91, 0x9c, 0x63, 0x95, 0xdc,
	0x23, 0x93, 0xc7, 0xe3, 0xe5, 0x0a, 0x1e, 0xf6, 0x59, 0xa5, 0x95, 0xce, 0xd8, 0xb6, 0x75, 0xf5,
	0x96, 0x04, 0x4e, 0x9b, 0xa4, 0x0b, 0x4f, 0x34, 0x4e, 0x54, 0x9b, 0xe4, 0x69, 0x15, 0x5b, 0xc2,
	0xf9, 0x67, 0x91, 0x9d, 0xce, 0x03, 0x2e, 0x96, 0x01, 0x4e, 0xb2, 0xc0, 0x69, 0xd8, 0xc5, 0x03,
	0xdc, 0x65, 0x81, 0xe5, 0x37, 0x31, 0x3d, 0xc6, 0x96, 0x7c, 0x16, 0x43, 0x70, 0x8c, 0x33, 0xf5,
	0x0e, 0x1e, 0x26, 0xa9, 0x2d, 0xb5, 0xa7, 0xda, 0x72, 0x25, 0xb0, 0x43, 0xc3, 0x73, 0x75, 0x09,
	0xe0, 0xc9, 0x64, 0xba, 0x10, 0x11, 0x17, 0xcb, 0x7f, 0x67, 0xf0, 0xfd, 0x3e, 0x1e, 0x79, 0xef,
	0xfc, 0x50, 0x64, 0xec, 0x48, 0xa6, 0x29, 0x27, 0xa0, 0xbe, 0x86, 0x37, 0xa3, 0x46, 0xf9, 0x30,
	0xe2, 0x14, 0xd9, 0x90, 0x4f, 0x8d, 0x95, 0xa2, 0x6f, 0xe1, 0xcb, 0x97, 0xa2, 0xdb, 0x92, 0x4f,
	0xb5, 0x6d, 0x6c, 0xc0, 0xb9, 0x20, 0xbd, 0xf6, 0x74, 0x94, 0x17, 0xea, 0x3b, 0x78, 0xf7, 0x2c,
	0xbb, 0x50, 0x8a, 0xb2, 0x4d, 0xc3, 0x65, 0x7d, 0xbe, 0x94, 0x03, 0x19, 0xf9, 0x73, 0x08, 0x3b,
	0xa6, 0x51, 0xd6, 0x5b, 0x6d, 0x6b, 0x9d, 0xd5, 0x84, 0x87, 0xb2, 0xc8, 0x97, 0x49, 0xf0, 0xe8,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0xd1, 0xa0, 0x84, 0xa9, 0x04, 0x00, 0x00,
}
