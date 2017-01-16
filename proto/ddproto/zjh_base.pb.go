// Code generated by protoc-gen-go.
// source: zjh_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	// //////////////////////////////////
	EProtoId_ZJH_PID_HEARTBEAT           EProtoId = 0
	EProtoId_ZJH_PID_QUICKCOnn           EProtoId = 1
	EProtoId_ZJH_PID_QUICKCONNA_ACK      EProtoId = 2
	EProtoId_ZJH_PID_GAME_LOGIN          EProtoId = 3
	EProtoId_ZJH_PID_GAME_LOGIN_ACK      EProtoId = 4
	EProtoId_ZJH_PID_GET_ROOM_LIST       EProtoId = 5
	EProtoId_ZJH_PID_ENTER_ROOM_LIST_ACK EProtoId = 6
	EProtoId_ZJH_PID_AUTO_ENTER_DESK     EProtoId = 7
	EProtoId_ZJH_PID_AUTO_ENTER_DESK_ACK EProtoId = 8
	EProtoId_ZJH_PID_SEND_GAMEINFO       EProtoId = 9
	EProtoId_ZJH_PID_BC_NEWENTER         EProtoId = 10
	EProtoId_ZJH_PID_BC_OPENING          EProtoId = 11
	EProtoId_ZJH_PID_REQ_READY           EProtoId = 12
	EProtoId_ZJH_PID_BC_READY            EProtoId = 13
	EProtoId_ZJH_PID_REQ_LEAVEDESK       EProtoId = 14
	EProtoId_ZJH_PID_BC_LEAVEDESK        EProtoId = 15
	EProtoId_ZJH_PID_BC_NEXT             EProtoId = 16
	EProtoId_ZJH_PID_REQ_KANPAI          EProtoId = 17
	EProtoId_ZJH_PID_BC_KANPAI           EProtoId = 18
	EProtoId_ZJH_PID_REQ_QIPAI           EProtoId = 19
	EProtoId_ZJH_PID_BC_QIPAI            EProtoId = 20
	EProtoId_ZJH_PID_REQ_GENZHU          EProtoId = 21
	EProtoId_ZJH_PID_BC_GENZHU           EProtoId = 22
	EProtoId_ZJH_PID_REQ_FAQI_XUEPIN     EProtoId = 23
	EProtoId_ZJH_PID_BC_FAQI_XUEPIN      EProtoId = 24
	EProtoId_ZJH_PID_BC_XUEPIN_END       EProtoId = 25
	EProtoId_ZJH_PID_REQ_JIAZHU          EProtoId = 26
	EProtoId_ZJH_PID_BC_JIAZHU           EProtoId = 27
	EProtoId_ZJH_PID_REQ_BIPAI           EProtoId = 28
	EProtoId_ZJH_PID_BC_BIPAI            EProtoId = 29
	EProtoId_ZJH_PID_BC_GAME_END         EProtoId = 30
	EProtoId_ZJH_PID_BC_DESK_STATE       EProtoId = 31
)

var EProtoId_name = map[int32]string{
	0:  "ZJH_PID_HEARTBEAT",
	1:  "ZJH_PID_QUICKCOnn",
	2:  "ZJH_PID_QUICKCONNA_ACK",
	3:  "ZJH_PID_GAME_LOGIN",
	4:  "ZJH_PID_GAME_LOGIN_ACK",
	5:  "ZJH_PID_GET_ROOM_LIST",
	6:  "ZJH_PID_ENTER_ROOM_LIST_ACK",
	7:  "ZJH_PID_AUTO_ENTER_DESK",
	8:  "ZJH_PID_AUTO_ENTER_DESK_ACK",
	9:  "ZJH_PID_SEND_GAMEINFO",
	10: "ZJH_PID_BC_NEWENTER",
	11: "ZJH_PID_BC_OPENING",
	12: "ZJH_PID_REQ_READY",
	13: "ZJH_PID_BC_READY",
	14: "ZJH_PID_REQ_LEAVEDESK",
	15: "ZJH_PID_BC_LEAVEDESK",
	16: "ZJH_PID_BC_NEXT",
	17: "ZJH_PID_REQ_KANPAI",
	18: "ZJH_PID_BC_KANPAI",
	19: "ZJH_PID_REQ_QIPAI",
	20: "ZJH_PID_BC_QIPAI",
	21: "ZJH_PID_REQ_GENZHU",
	22: "ZJH_PID_BC_GENZHU",
	23: "ZJH_PID_REQ_FAQI_XUEPIN",
	24: "ZJH_PID_BC_FAQI_XUEPIN",
	25: "ZJH_PID_BC_XUEPIN_END",
	26: "ZJH_PID_REQ_JIAZHU",
	27: "ZJH_PID_BC_JIAZHU",
	28: "ZJH_PID_REQ_BIPAI",
	29: "ZJH_PID_BC_BIPAI",
	30: "ZJH_PID_BC_GAME_END",
	31: "ZJH_PID_BC_DESK_STATE",
}
var EProtoId_value = map[string]int32{
	"ZJH_PID_HEARTBEAT":           0,
	"ZJH_PID_QUICKCOnn":           1,
	"ZJH_PID_QUICKCONNA_ACK":      2,
	"ZJH_PID_GAME_LOGIN":          3,
	"ZJH_PID_GAME_LOGIN_ACK":      4,
	"ZJH_PID_GET_ROOM_LIST":       5,
	"ZJH_PID_ENTER_ROOM_LIST_ACK": 6,
	"ZJH_PID_AUTO_ENTER_DESK":     7,
	"ZJH_PID_AUTO_ENTER_DESK_ACK": 8,
	"ZJH_PID_SEND_GAMEINFO":       9,
	"ZJH_PID_BC_NEWENTER":         10,
	"ZJH_PID_BC_OPENING":          11,
	"ZJH_PID_REQ_READY":           12,
	"ZJH_PID_BC_READY":            13,
	"ZJH_PID_REQ_LEAVEDESK":       14,
	"ZJH_PID_BC_LEAVEDESK":        15,
	"ZJH_PID_BC_NEXT":             16,
	"ZJH_PID_REQ_KANPAI":          17,
	"ZJH_PID_BC_KANPAI":           18,
	"ZJH_PID_REQ_QIPAI":           19,
	"ZJH_PID_BC_QIPAI":            20,
	"ZJH_PID_REQ_GENZHU":          21,
	"ZJH_PID_BC_GENZHU":           22,
	"ZJH_PID_REQ_FAQI_XUEPIN":     23,
	"ZJH_PID_BC_FAQI_XUEPIN":      24,
	"ZJH_PID_BC_XUEPIN_END":       25,
	"ZJH_PID_REQ_JIAZHU":          26,
	"ZJH_PID_BC_JIAZHU":           27,
	"ZJH_PID_REQ_BIPAI":           28,
	"ZJH_PID_BC_BIPAI":            29,
	"ZJH_PID_BC_GAME_END":         30,
	"ZJH_PID_BC_DESK_STATE":       31,
}

func (x EProtoId) Enum() *EProtoId {
	p := new(EProtoId)
	*p = x
	return p
}
func (x EProtoId) String() string {
	return proto.EnumName(EProtoId_name, int32(x))
}
func (x *EProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoId_value, data, "EProtoId")
	if err != nil {
		return err
	}
	*x = EProtoId(value)
	return nil
}
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

// 玩家当前状态
type ZjhEnumPlayerGameStatus int32

const (
	ZjhEnumPlayerGameStatus_zjh_SS_GAMING   ZjhEnumPlayerGameStatus = 1
	ZjhEnumPlayerGameStatus_zjh_SS_STANDING ZjhEnumPlayerGameStatus = 2
	ZjhEnumPlayerGameStatus_zjh_SS_SITED    ZjhEnumPlayerGameStatus = 3
	ZjhEnumPlayerGameStatus_zjh_SS_READY    ZjhEnumPlayerGameStatus = 4
	ZjhEnumPlayerGameStatus_zjh_SS_OFFLINE  ZjhEnumPlayerGameStatus = 5
)

var ZjhEnumPlayerGameStatus_name = map[int32]string{
	1: "zjh_SS_GAMING",
	2: "zjh_SS_STANDING",
	3: "zjh_SS_SITED",
	4: "zjh_SS_READY",
	5: "zjh_SS_OFFLINE",
}
var ZjhEnumPlayerGameStatus_value = map[string]int32{
	"zjh_SS_GAMING":   1,
	"zjh_SS_STANDING": 2,
	"zjh_SS_SITED":    3,
	"zjh_SS_READY":    4,
	"zjh_SS_OFFLINE":  5,
}

func (x ZjhEnumPlayerGameStatus) Enum() *ZjhEnumPlayerGameStatus {
	p := new(ZjhEnumPlayerGameStatus)
	*p = x
	return p
}
func (x ZjhEnumPlayerGameStatus) String() string {
	return proto.EnumName(ZjhEnumPlayerGameStatus_name, int32(x))
}
func (x *ZjhEnumPlayerGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumPlayerGameStatus_value, data, "ZjhEnumPlayerGameStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumPlayerGameStatus(value)
	return nil
}
func (ZjhEnumPlayerGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

// 桌面状态
type ZjhEnumDeskState int32

const (
	ZjhEnumDeskState_DESK_IS_GAMING ZjhEnumDeskState = 1
	ZjhEnumDeskState_DESK_IS_WAIT   ZjhEnumDeskState = 2
)

var ZjhEnumDeskState_name = map[int32]string{
	1: "DESK_IS_GAMING",
	2: "DESK_IS_WAIT",
}
var ZjhEnumDeskState_value = map[string]int32{
	"DESK_IS_GAMING": 1,
	"DESK_IS_WAIT":   2,
}

func (x ZjhEnumDeskState) Enum() *ZjhEnumDeskState {
	p := new(ZjhEnumDeskState)
	*p = x
	return p
}
func (x ZjhEnumDeskState) String() string {
	return proto.EnumName(ZjhEnumDeskState_name, int32(x))
}
func (x *ZjhEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumDeskState_value, data, "ZjhEnumDeskState")
	if err != nil {
		return err
	}
	*x = ZjhEnumDeskState(value)
	return nil
}
func (ZjhEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

// 房间玩家状态
type ZjhEnumUserState int32

const (
	ZjhEnumUserState_USER_IS_GAMING ZjhEnumUserState = 1
	ZjhEnumUserState_USER_IS_STAND  ZjhEnumUserState = 2
	ZjhEnumUserState_USER_IS_SITED  ZjhEnumUserState = 3
)

var ZjhEnumUserState_name = map[int32]string{
	1: "USER_IS_GAMING",
	2: "USER_IS_STAND",
	3: "USER_IS_SITED",
}
var ZjhEnumUserState_value = map[string]int32{
	"USER_IS_GAMING": 1,
	"USER_IS_STAND":  2,
	"USER_IS_SITED":  3,
}

func (x ZjhEnumUserState) Enum() *ZjhEnumUserState {
	p := new(ZjhEnumUserState)
	*p = x
	return p
}
func (x ZjhEnumUserState) String() string {
	return proto.EnumName(ZjhEnumUserState_name, int32(x))
}
func (x *ZjhEnumUserState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserState_value, data, "ZjhEnumUserState")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserState(value)
	return nil
}
func (ZjhEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{3} }

// 房间类型
type ZjhEnumRoomType int32

const (
	ZjhEnumRoomType_ROOM_TYPE_FRIEND   ZjhEnumRoomType = 1
	ZjhEnumRoomType_ROOM_TYPE_NORMAL   ZjhEnumRoomType = 2
	ZjhEnumRoomType_ROOM_TYPE_REDBLACK ZjhEnumRoomType = 3
)

var ZjhEnumRoomType_name = map[int32]string{
	1: "ROOM_TYPE_FRIEND",
	2: "ROOM_TYPE_NORMAL",
	3: "ROOM_TYPE_REDBLACK",
}
var ZjhEnumRoomType_value = map[string]int32{
	"ROOM_TYPE_FRIEND":   1,
	"ROOM_TYPE_NORMAL":   2,
	"ROOM_TYPE_REDBLACK": 3,
}

func (x ZjhEnumRoomType) Enum() *ZjhEnumRoomType {
	p := new(ZjhEnumRoomType)
	*p = x
	return p
}
func (x ZjhEnumRoomType) String() string {
	return proto.EnumName(ZjhEnumRoomType_name, int32(x))
}
func (x *ZjhEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumRoomType_value, data, "ZjhEnumRoomType")
	if err != nil {
		return err
	}
	*x = ZjhEnumRoomType(value)
	return nil
}
func (ZjhEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("ddproto.ZjhEnumPlayerGameStatus", ZjhEnumPlayerGameStatus_name, ZjhEnumPlayerGameStatus_value)
	proto.RegisterEnum("ddproto.ZjhEnumDeskState", ZjhEnumDeskState_name, ZjhEnumDeskState_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserState", ZjhEnumUserState_name, ZjhEnumUserState_value)
	proto.RegisterEnum("ddproto.ZjhEnumRoomType", ZjhEnumRoomType_name, ZjhEnumRoomType_value)
}

var fileDescriptor20 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x94, 0xcb, 0x72, 0xda, 0x30,
	0x14, 0x86, 0x1b, 0x72, 0xad, 0x9a, 0xcb, 0x41, 0xdc, 0x42, 0x68, 0x9b, 0x3d, 0x8b, 0x3e, 0x40,
	0x77, 0xb2, 0x7d, 0x00, 0x05, 0x23, 0x1b, 0x5b, 0x94, 0x24, 0x1b, 0x0d, 0x1d, 0x3c, 0xd3, 0x69,
	0xcb, 0x65, 0xb8, 0x2c, 0xd2, 0xbe, 0x56, 0x1f, 0xb0, 0x23, 0x19, 0x83, 0xd0, 0xce, 0xf3, 0xfd,
	0xe7, 0xd7, 0xf9, 0x25, 0x1d, 0x99, 0xdc, 0xfe, 0xf9, 0xf9, 0x43, 0x7d, 0x9f, 0xac, 0xb3, 0x2f,
	0xcb, 0xd5, 0x62, 0xb3, 0xa0, 0x97, 0xd3, 0xa9, 0xf9, 0x68, 0xff, 0xbb, 0x20, 0x57, 0x18, 0xeb,
	0x4f, 0x3e, 0xa5, 0x35, 0x52, 0x7e, 0x7d, 0xea, 0xa9, 0x98, 0x07, 0xaa, 0x87, 0x2c, 0x91, 0x1e,
	0x32, 0x09, 0xef, 0x6c, 0x3c, 0x1c, 0x71, 0xbf, 0xef, 0x47, 0xf3, 0x39, 0x9c, 0xd0, 0x07, 0x52,
	0x77, 0xb0, 0x10, 0x4c, 0x31, 0xbf, 0x0f, 0x25, 0x5a, 0x27, 0xb4, 0xd0, 0xba, 0x6c, 0x80, 0x2a,
	0x8c, 0xba, 0x5c, 0xc0, 0xa9, 0xed, 0x39, 0x70, 0xe3, 0x39, 0xa3, 0x4d, 0x52, 0xdb, 0x6b, 0x28,
	0x55, 0x12, 0x45, 0x03, 0x15, 0xf2, 0x54, 0xc2, 0x39, 0x7d, 0x24, 0xad, 0x42, 0x42, 0x21, 0x31,
	0x39, 0x88, 0xc6, 0x7b, 0x41, 0x5b, 0xa4, 0x51, 0x14, 0xb0, 0x91, 0x8c, 0x76, 0x55, 0x01, 0xa6,
	0x7d, 0xb8, 0xb4, 0xdd, 0x8e, 0x68, 0xdc, 0x57, 0x76, 0xe7, 0x14, 0x45, 0x1e, 0x8d, 0x8b, 0x4e,
	0x04, 0xef, 0x69, 0x83, 0x54, 0x0a, 0xc9, 0xf3, 0x95, 0xc0, 0xb1, 0x31, 0x03, 0xb1, 0x77, 0xe8,
	0xf9, 0x2a, 0x8a, 0x51, 0x70, 0xd1, 0x85, 0x0f, 0xf6, 0x61, 0x25, 0x38, 0x54, 0x09, 0xb2, 0xe0,
	0x05, 0xae, 0x69, 0x95, 0x80, 0x55, 0x9e, 0xd3, 0x1b, 0xbb, 0xb1, 0x2e, 0x0e, 0x91, 0x7d, 0x43,
	0x13, 0xfa, 0x96, 0xde, 0x93, 0xaa, 0x65, 0x38, 0x28, 0x77, 0xb4, 0x42, 0xee, 0x8e, 0x22, 0x3d,
	0x4b, 0x00, 0x3b, 0x8e, 0x5e, 0xa9, 0xcf, 0x44, 0xcc, 0x38, 0x94, 0xed, 0x38, 0x9e, 0x5f, 0x60,
	0xea, 0xa6, 0x1c, 0x72, 0x8d, 0x2b, 0x4e, 0xca, 0x9c, 0x56, 0xdd, 0xb5, 0xbb, 0x28, 0x5e, 0x7b,
	0x23, 0xa8, 0x39, 0x6b, 0xef, 0x70, 0xdd, 0xbe, 0x0b, 0x5d, 0xde, 0x61, 0x43, 0xae, 0x9e, 0x47,
	0x18, 0x73, 0x01, 0x0d, 0x7b, 0x00, 0x3c, 0xff, 0x48, 0xbb, 0xb7, 0x4f, 0xc3, 0xf3, 0x77, 0x58,
	0xa1, 0x08, 0xa0, 0xe9, 0x46, 0x78, 0xe2, 0x4c, 0xf7, 0x7a, 0x70, 0x22, 0xec, 0x70, 0xcb, 0xdd,
	0x9e, 0x67, 0x36, 0xf2, 0xd1, 0xd9, 0x5e, 0x4e, 0x3f, 0x39, 0x57, 0x6c, 0xc6, 0x52, 0x37, 0xfd,
	0xec, 0xe4, 0x31, 0xf3, 0x92, 0x4a, 0x26, 0x11, 0x1e, 0xdb, 0x7f, 0x49, 0x53, 0xbf, 0xa8, 0x6c,
	0xbe, 0x9d, 0xa9, 0xe5, 0xef, 0xc9, 0x5b, 0xb6, 0xea, 0x4e, 0x66, 0x59, 0xba, 0x99, 0x6c, 0xb6,
	0x6b, 0x5a, 0x26, 0x37, 0x5a, 0x4c, 0x53, 0xbd, 0x98, 0x9e, 0x8a, 0x13, 0x7d, 0x67, 0x3b, 0x94,
	0x4a, 0x26, 0x02, 0x0d, 0x4b, 0x14, 0xc8, 0x75, 0x01, 0xb9, 0xc4, 0x00, 0x4e, 0x2d, 0x92, 0x4f,
	0xc8, 0x19, 0xa5, 0xf9, 0xd3, 0x4d, 0x53, 0x15, 0x75, 0x3a, 0x21, 0x17, 0x08, 0xe7, 0xed, 0xaf,
	0x84, 0xee, 0x9b, 0x4f, 0xb3, 0xf5, 0x2f, 0xdd, 0x36, 0xd3, 0x95, 0x26, 0x22, 0xb7, 0xda, 0x02,
	0xb9, 0x2e, 0xd8, 0x98, 0x71, 0x09, 0xa5, 0xb6, 0xb0, 0xbc, 0xdb, 0x75, 0xb6, 0xda, 0x7b, 0x47,
	0x29, 0x26, 0x47, 0xde, 0x32, 0xb9, 0x29, 0x98, 0xc9, 0x0c, 0xa5, 0x23, 0x94, 0x27, 0x6e, 0x8f,
	0x49, 0x79, 0xbf, 0xde, 0x6a, 0xb1, 0x98, 0xc9, 0xb7, 0x65, 0xa6, 0xcf, 0xd9, 0x3c, 0x50, 0xf9,
	0x12, 0xa3, 0xea, 0x24, 0x5c, 0x1f, 0xe7, 0xc9, 0x31, 0x15, 0x51, 0x32, 0x60, 0x61, 0xfe, 0xa7,
	0x38, 0xd0, 0x04, 0x03, 0x2f, 0xd4, 0x6f, 0xf2, 0xf4, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0x6f, 0xf3, 0xcb, 0xb2, 0x04, 0x00, 0x00,
}
