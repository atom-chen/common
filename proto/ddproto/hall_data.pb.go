// Code generated by protoc-gen-go.
// source: hall_data.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HallEnumProtoId int32

const (
	HallEnumProtoId_HALL_PID_HEARTBEAT                   HallEnumProtoId = 0
	HallEnumProtoId_HALL_PID_QUICK_CONN                  HallEnumProtoId = 1
	HallEnumProtoId_HALL_PID_QUICK_CONN_ACK              HallEnumProtoId = 2
	HallEnumProtoId_HALL_PID_GAME_LOGIN                  HallEnumProtoId = 3
	HallEnumProtoId_HALL_PID_GAME_LOGIN_ACK              HallEnumProtoId = 4
	HallEnumProtoId_HALL_PID_WXPAYUNIFIEDORDER_REQ       HallEnumProtoId = 5
	HallEnumProtoId_HALL_PID_WXPAYUNIFIEDORDER_ACK       HallEnumProtoId = 6
	HallEnumProtoId_HALL_PID_WXPAYSYNCCHECKER_REQ        HallEnumProtoId = 7
	HallEnumProtoId_HALL_PID_WXPAYSYNCCHECKER_ACK        HallEnumProtoId = 8
	HallEnumProtoId_HALL_PID_USER_DATA                   HallEnumProtoId = 9
	HallEnumProtoId_HALL_PID_USER_DATA_ACK               HallEnumProtoId = 10
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY                HallEnumProtoId = 11
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY_ACK            HallEnumProtoId = 12
	HallEnumProtoId_HALL_PID_DS_LOTTERY_INFO_ACK         HallEnumProtoId = 13
	HallEnumProtoId_HALL_PID_ONLINEWARD_REQ              HallEnumProtoId = 14
	HallEnumProtoId_HALL_PID_ONLINEWARD_ACK              HallEnumProtoId = 15
	HallEnumProtoId_HALL_PID_EVENT_REQ                   HallEnumProtoId = 16
	HallEnumProtoId_HALL_PID_EVENT_ACK                   HallEnumProtoId = 17
	HallEnumProtoId_HALL_PID_GOODS_LIST_REQ              HallEnumProtoId = 18
	HallEnumProtoId_HALL_PID_GOODS_LIST_ACK              HallEnumProtoId = 19
	HallEnumProtoId_HALL_PID_HOTUPDATEREQVERSIONINFO     HallEnumProtoId = 20
	HallEnumProtoId_HALL_PID_HOTUPDATEACKVERSIONINFO     HallEnumProtoId = 21
	HallEnumProtoId_HALL_PID_BAG_ITEMS_REQ               HallEnumProtoId = 22
	HallEnumProtoId_HALL_PID_BAG_ITEMS_ACK               HallEnumProtoId = 23
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ HallEnumProtoId = 24
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK HallEnumProtoId = 25
	HallEnumProtoId_HALL_PID_APPLE_PAY_CB_REQ            HallEnumProtoId = 26
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_REQ     HallEnumProtoId = 27
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_ACK     HallEnumProtoId = 28
	HallEnumProtoId_HALL_PID_MAIL_LIST_REQ               HallEnumProtoId = 29
	HallEnumProtoId_HALL_PID_MAIL_LIST_ACK               HallEnumProtoId = 30
	HallEnumProtoId_HALL_PID_HALL_DSLOTTERYINFO_REQ      HallEnumProtoId = 31
	HallEnumProtoId_HALL_PID_FRIENDS_LIST_REQ            HallEnumProtoId = 32
	HallEnumProtoId_HALL_PID_FRIENDS_LIST_ACK            HallEnumProtoId = 33
	HallEnumProtoId_HALL_PID_FRIEND_ADD_REQ              HallEnumProtoId = 34
	HallEnumProtoId_HALL_PID_FRIEND_ADD_ACK              HallEnumProtoId = 35
	HallEnumProtoId_HALL_PID_FRIEND_DEL_REQ              HallEnumProtoId = 36
	HallEnumProtoId_HALL_PID_FRIEND_DEL_ACK              HallEnumProtoId = 37
)

var HallEnumProtoId_name = map[int32]string{
	0:  "HALL_PID_HEARTBEAT",
	1:  "HALL_PID_QUICK_CONN",
	2:  "HALL_PID_QUICK_CONN_ACK",
	3:  "HALL_PID_GAME_LOGIN",
	4:  "HALL_PID_GAME_LOGIN_ACK",
	5:  "HALL_PID_WXPAYUNIFIEDORDER_REQ",
	6:  "HALL_PID_WXPAYUNIFIEDORDER_ACK",
	7:  "HALL_PID_WXPAYSYNCCHECKER_REQ",
	8:  "HALL_PID_WXPAYSYNCCHECKER_ACK",
	9:  "HALL_PID_USER_DATA",
	10: "HALL_PID_USER_DATA_ACK",
	11: "HALL_PID_DRAW_LOTTERY",
	12: "HALL_PID_DRAW_LOTTERY_ACK",
	13: "HALL_PID_DS_LOTTERY_INFO_ACK",
	14: "HALL_PID_ONLINEWARD_REQ",
	15: "HALL_PID_ONLINEWARD_ACK",
	16: "HALL_PID_EVENT_REQ",
	17: "HALL_PID_EVENT_ACK",
	18: "HALL_PID_GOODS_LIST_REQ",
	19: "HALL_PID_GOODS_LIST_ACK",
	20: "HALL_PID_HOTUPDATEREQVERSIONINFO",
	21: "HALL_PID_HOTUPDATEACKVERSIONINFO",
	22: "HALL_PID_BAG_ITEMS_REQ",
	23: "HALL_PID_BAG_ITEMS_ACK",
	24: "HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ",
	25: "HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK",
	26: "HALL_PID_APPLE_PAY_CB_REQ",
	27: "HALL_PID_HOTUPDATEASSETSINFO_REQ",
	28: "HALL_PID_HOTUPDATEASSETSINFO_ACK",
	29: "HALL_PID_MAIL_LIST_REQ",
	30: "HALL_PID_MAIL_LIST_ACK",
	31: "HALL_PID_HALL_DSLOTTERYINFO_REQ",
	32: "HALL_PID_FRIENDS_LIST_REQ",
	33: "HALL_PID_FRIENDS_LIST_ACK",
	34: "HALL_PID_FRIEND_ADD_REQ",
	35: "HALL_PID_FRIEND_ADD_ACK",
	36: "HALL_PID_FRIEND_DEL_REQ",
	37: "HALL_PID_FRIEND_DEL_ACK",
}
var HallEnumProtoId_value = map[string]int32{
	"HALL_PID_HEARTBEAT":                   0,
	"HALL_PID_QUICK_CONN":                  1,
	"HALL_PID_QUICK_CONN_ACK":              2,
	"HALL_PID_GAME_LOGIN":                  3,
	"HALL_PID_GAME_LOGIN_ACK":              4,
	"HALL_PID_WXPAYUNIFIEDORDER_REQ":       5,
	"HALL_PID_WXPAYUNIFIEDORDER_ACK":       6,
	"HALL_PID_WXPAYSYNCCHECKER_REQ":        7,
	"HALL_PID_WXPAYSYNCCHECKER_ACK":        8,
	"HALL_PID_USER_DATA":                   9,
	"HALL_PID_USER_DATA_ACK":               10,
	"HALL_PID_DRAW_LOTTERY":                11,
	"HALL_PID_DRAW_LOTTERY_ACK":            12,
	"HALL_PID_DS_LOTTERY_INFO_ACK":         13,
	"HALL_PID_ONLINEWARD_REQ":              14,
	"HALL_PID_ONLINEWARD_ACK":              15,
	"HALL_PID_EVENT_REQ":                   16,
	"HALL_PID_EVENT_ACK":                   17,
	"HALL_PID_GOODS_LIST_REQ":              18,
	"HALL_PID_GOODS_LIST_ACK":              19,
	"HALL_PID_HOTUPDATEREQVERSIONINFO":     20,
	"HALL_PID_HOTUPDATEACKVERSIONINFO":     21,
	"HALL_PID_BAG_ITEMS_REQ":               22,
	"HALL_PID_BAG_ITEMS_ACK":               23,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ": 24,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK": 25,
	"HALL_PID_APPLE_PAY_CB_REQ":            26,
	"HALL_PID_HOTUPDATEASSETSINFO_REQ":     27,
	"HALL_PID_HOTUPDATEASSETSINFO_ACK":     28,
	"HALL_PID_MAIL_LIST_REQ":               29,
	"HALL_PID_MAIL_LIST_ACK":               30,
	"HALL_PID_HALL_DSLOTTERYINFO_REQ":      31,
	"HALL_PID_FRIENDS_LIST_REQ":            32,
	"HALL_PID_FRIENDS_LIST_ACK":            33,
	"HALL_PID_FRIEND_ADD_REQ":              34,
	"HALL_PID_FRIEND_ADD_ACK":              35,
	"HALL_PID_FRIEND_DEL_REQ":              36,
	"HALL_PID_FRIEND_DEL_ACK":              37,
}

func (x HallEnumProtoId) Enum() *HallEnumProtoId {
	p := new(HallEnumProtoId)
	*p = x
	return p
}
func (x HallEnumProtoId) String() string {
	return proto.EnumName(HallEnumProtoId_name, int32(x))
}
func (x *HallEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumProtoId_value, data, "HallEnumProtoId")
	if err != nil {
		return err
	}
	*x = HallEnumProtoId(value)
	return nil
}
func (HallEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

// 活动类型
type HallEnumEvent int32

const (
	HallEnumEvent_TYPE_TIME HallEnumEvent = 1
	HallEnumEvent_TYPE_NEW  HallEnumEvent = 2
	HallEnumEvent_TYPE_NULL HallEnumEvent = 3
)

var HallEnumEvent_name = map[int32]string{
	1: "TYPE_TIME",
	2: "TYPE_NEW",
	3: "TYPE_NULL",
}
var HallEnumEvent_value = map[string]int32{
	"TYPE_TIME": 1,
	"TYPE_NEW":  2,
	"TYPE_NULL": 3,
}

func (x HallEnumEvent) Enum() *HallEnumEvent {
	p := new(HallEnumEvent)
	*p = x
	return p
}
func (x HallEnumEvent) String() string {
	return proto.EnumName(HallEnumEvent_name, int32(x))
}
func (x *HallEnumEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumEvent_value, data, "HallEnumEvent")
	if err != nil {
		return err
	}
	*x = HallEnumEvent(value)
	return nil
}
func (HallEnumEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

// 活动奖品
type HallEnum_Reward int32

const (
	HallEnum_Reward_RE_EXP  HallEnum_Reward = 1
	HallEnum_Reward_RE_GIFT HallEnum_Reward = 2
)

var HallEnum_Reward_name = map[int32]string{
	1: "RE_EXP",
	2: "RE_GIFT",
}
var HallEnum_Reward_value = map[string]int32{
	"RE_EXP":  1,
	"RE_GIFT": 2,
}

func (x HallEnum_Reward) Enum() *HallEnum_Reward {
	p := new(HallEnum_Reward)
	*p = x
	return p
}
func (x HallEnum_Reward) String() string {
	return proto.EnumName(HallEnum_Reward_name, int32(x))
}
func (x *HallEnum_Reward) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnum_Reward_value, data, "HallEnum_Reward")
	if err != nil {
		return err
	}
	*x = HallEnum_Reward(value)
	return nil
}
func (HallEnum_Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

// 邮件类型
type HallEnumMailType int32

const (
	HallEnumMailType_SYSTEM     HallEnumMailType = 1
	HallEnumMailType_FRIEND_ADD HallEnumMailType = 2
)

var HallEnumMailType_name = map[int32]string{
	1: "SYSTEM",
	2: "FRIEND_ADD",
}
var HallEnumMailType_value = map[string]int32{
	"SYSTEM":     1,
	"FRIEND_ADD": 2,
}

func (x HallEnumMailType) Enum() *HallEnumMailType {
	p := new(HallEnumMailType)
	*p = x
	return p
}
func (x HallEnumMailType) String() string {
	return proto.EnumName(HallEnumMailType_name, int32(x))
}
func (x *HallEnumMailType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumMailType_value, data, "HallEnumMailType")
	if err != nil {
		return err
	}
	*x = HallEnumMailType(value)
	return nil
}
func (HallEnumMailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

// 道具类型
type HallEnumPropsType int32

const (
	HallEnumPropsType_TYPE_LABA     HallEnumPropsType = 1
	HallEnumPropsType_TYPE_JIPAIQI  HallEnumPropsType = 2
	HallEnumPropsType_TYPE_FANGKA   HallEnumPropsType = 3
	HallEnumPropsType_TYPE_EXP_3000 HallEnumPropsType = 4
)

var HallEnumPropsType_name = map[int32]string{
	1: "TYPE_LABA",
	2: "TYPE_JIPAIQI",
	3: "TYPE_FANGKA",
	4: "TYPE_EXP_3000",
}
var HallEnumPropsType_value = map[string]int32{
	"TYPE_LABA":     1,
	"TYPE_JIPAIQI":  2,
	"TYPE_FANGKA":   3,
	"TYPE_EXP_3000": 4,
}

func (x HallEnumPropsType) Enum() *HallEnumPropsType {
	p := new(HallEnumPropsType)
	*p = x
	return p
}
func (x HallEnumPropsType) String() string {
	return proto.EnumName(HallEnumPropsType_name, int32(x))
}
func (x *HallEnumPropsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumPropsType_value, data, "HallEnumPropsType")
	if err != nil {
		return err
	}
	*x = HallEnumPropsType(value)
	return nil
}
func (HallEnumPropsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

// 任务类型
type HallEnumTaskType int32

const (
	HallEnumTaskType_TYPE_MJ  HallEnumTaskType = 1
	HallEnumTaskType_TYPE_DDZ HallEnumTaskType = 2
	HallEnumTaskType_TYPE_ZJH HallEnumTaskType = 3
)

var HallEnumTaskType_name = map[int32]string{
	1: "TYPE_MJ",
	2: "TYPE_DDZ",
	3: "TYPE_ZJH",
}
var HallEnumTaskType_value = map[string]int32{
	"TYPE_MJ":  1,
	"TYPE_DDZ": 2,
	"TYPE_ZJH": 3,
}

func (x HallEnumTaskType) Enum() *HallEnumTaskType {
	p := new(HallEnumTaskType)
	*p = x
	return p
}
func (x HallEnumTaskType) String() string {
	return proto.EnumName(HallEnumTaskType_name, int32(x))
}
func (x *HallEnumTaskType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTaskType_value, data, "HallEnumTaskType")
	if err != nil {
		return err
	}
	*x = HallEnumTaskType(value)
	return nil
}
func (HallEnumTaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

// vip等级
type HallUser_VIP int32

const (
	HallUser_VIP_LV_1 HallUser_VIP = 1
	HallUser_VIP_LV_2 HallUser_VIP = 2
	HallUser_VIP_LV_3 HallUser_VIP = 3
	HallUser_VIP_LV_4 HallUser_VIP = 4
	HallUser_VIP_LV_5 HallUser_VIP = 5
	HallUser_VIP_LV_6 HallUser_VIP = 6
)

var HallUser_VIP_name = map[int32]string{
	1: "LV_1",
	2: "LV_2",
	3: "LV_3",
	4: "LV_4",
	5: "LV_5",
	6: "LV_6",
}
var HallUser_VIP_value = map[string]int32{
	"LV_1": 1,
	"LV_2": 2,
	"LV_3": 3,
	"LV_4": 4,
	"LV_5": 5,
	"LV_6": 6,
}

func (x HallUser_VIP) Enum() *HallUser_VIP {
	p := new(HallUser_VIP)
	*p = x
	return p
}
func (x HallUser_VIP) String() string {
	return proto.EnumName(HallUser_VIP_name, int32(x))
}
func (x *HallUser_VIP) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallUser_VIP_value, data, "HallUser_VIP")
	if err != nil {
		return err
	}
	*x = HallUser_VIP(value)
	return nil
}
func (HallUser_VIP) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

// 单个活动
type HallItemEvent struct {
	Id               *int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumEvent   `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumEvent" json:"type,omitempty"`
	Reward           *HallEnum_Reward `protobuf:"varint,3,opt,name=reward,enum=ddproto.HallEnum_Reward" json:"reward,omitempty"`
	RichText         []string         `protobuf:"bytes,5,rep,name=richText" json:"richText,omitempty"`
	Title            *string          `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallItemEvent) Reset()                    { *m = HallItemEvent{} }
func (m *HallItemEvent) String() string            { return proto.CompactTextString(m) }
func (*HallItemEvent) ProtoMessage()               {}
func (*HallItemEvent) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *HallItemEvent) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemEvent) GetType() HallEnumEvent {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumEvent_TYPE_TIME
}

func (m *HallItemEvent) GetReward() HallEnum_Reward {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return HallEnum_Reward_RE_EXP
}

func (m *HallItemEvent) GetRichText() []string {
	if m != nil {
		return m.RichText
	}
	return nil
}

func (m *HallItemEvent) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

// 单个邮件
type HallMailItem struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumMailType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumMailType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content          *string           `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	IsWatch          *bool             `protobuf:"varint,5,opt,name=isWatch" json:"isWatch,omitempty"`
	IsCheck          *bool             `protobuf:"varint,6,opt,name=isCheck" json:"isCheck,omitempty"`
	Attach           []*HallBagItem    `protobuf:"bytes,7,rep,name=attach" json:"attach,omitempty"`
	Date             *int64            `protobuf:"varint,8,opt,name=date" json:"date,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallMailItem) Reset()                    { *m = HallMailItem{} }
func (m *HallMailItem) String() string            { return proto.CompactTextString(m) }
func (*HallMailItem) ProtoMessage()               {}
func (*HallMailItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *HallMailItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallMailItem) GetType() HallEnumMailType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumMailType_SYSTEM
}

func (m *HallMailItem) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallMailItem) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *HallMailItem) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *HallMailItem) GetIsCheck() bool {
	if m != nil && m.IsCheck != nil {
		return *m.IsCheck
	}
	return false
}

func (m *HallMailItem) GetAttach() []*HallBagItem {
	if m != nil {
		return m.Attach
	}
	return nil
}

func (m *HallMailItem) GetDate() int64 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

// 背包单个道具
type HallBagItem struct {
	Type             *HallEnumPropsType `protobuf:"varint,1,opt,name=type,enum=ddproto.HallEnumPropsType" json:"type,omitempty"`
	Amount           *int32             `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallBagItem) Reset()                    { *m = HallBagItem{} }
func (m *HallBagItem) String() string            { return proto.CompactTextString(m) }
func (*HallBagItem) ProtoMessage()               {}
func (*HallBagItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *HallBagItem) GetType() HallEnumPropsType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumPropsType_TYPE_LABA
}

func (m *HallBagItem) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// 单个任务
type HallItemTask struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumTaskType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumTaskType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Reward           *string           `protobuf:"bytes,4,opt,name=reward" json:"reward,omitempty"`
	Num              *int32            `protobuf:"varint,5,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallItemTask) Reset()                    { *m = HallItemTask{} }
func (m *HallItemTask) String() string            { return proto.CompactTextString(m) }
func (*HallItemTask) ProtoMessage()               {}
func (*HallItemTask) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *HallItemTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemTask) GetType() HallEnumTaskType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumTaskType_TYPE_MJ
}

func (m *HallItemTask) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallItemTask) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

func (m *HallItemTask) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 个人信息
type HallUserData struct {
	UserName          *string       `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	UserId            *int32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NiceValue         *int64        `protobuf:"varint,3,opt,name=niceValue" json:"niceValue,omitempty"`
	EvilValue         *int64        `protobuf:"varint,4,opt,name=evilValue" json:"evilValue,omitempty"`
	UserLevel         *int32        `protobuf:"varint,5,opt,name=userLevel" json:"userLevel,omitempty"`
	UserVIP           *bool         `protobuf:"varint,6,opt,name=userVIP" json:"userVIP,omitempty"`
	UserVIPLv         *HallUser_VIP `protobuf:"varint,7,opt,name=userVIPLv,enum=ddproto.HallUser_VIP" json:"userVIPLv,omitempty"`
	UserMoney         *int64        `protobuf:"varint,8,opt,name=userMoney" json:"userMoney,omitempty"`
	UserDiamond       *int64        `protobuf:"varint,9,opt,name=userDiamond" json:"userDiamond,omitempty"`
	UserRedBag        *string       `protobuf:"bytes,10,opt,name=userRedBag" json:"userRedBag,omitempty"`
	UserLotteryTicket *int64        `protobuf:"varint,11,opt,name=userLotteryTicket" json:"userLotteryTicket,omitempty"`
	Sex               *bool         `protobuf:"varint,12,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized  []byte        `json:"-"`
}

func (m *HallUserData) Reset()                    { *m = HallUserData{} }
func (m *HallUserData) String() string            { return proto.CompactTextString(m) }
func (*HallUserData) ProtoMessage()               {}
func (*HallUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *HallUserData) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *HallUserData) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallUserData) GetNiceValue() int64 {
	if m != nil && m.NiceValue != nil {
		return *m.NiceValue
	}
	return 0
}

func (m *HallUserData) GetEvilValue() int64 {
	if m != nil && m.EvilValue != nil {
		return *m.EvilValue
	}
	return 0
}

func (m *HallUserData) GetUserLevel() int32 {
	if m != nil && m.UserLevel != nil {
		return *m.UserLevel
	}
	return 0
}

func (m *HallUserData) GetUserVIP() bool {
	if m != nil && m.UserVIP != nil {
		return *m.UserVIP
	}
	return false
}

func (m *HallUserData) GetUserVIPLv() HallUser_VIP {
	if m != nil && m.UserVIPLv != nil {
		return *m.UserVIPLv
	}
	return HallUser_VIP_LV_1
}

func (m *HallUserData) GetUserMoney() int64 {
	if m != nil && m.UserMoney != nil {
		return *m.UserMoney
	}
	return 0
}

func (m *HallUserData) GetUserDiamond() int64 {
	if m != nil && m.UserDiamond != nil {
		return *m.UserDiamond
	}
	return 0
}

func (m *HallUserData) GetUserRedBag() string {
	if m != nil && m.UserRedBag != nil {
		return *m.UserRedBag
	}
	return ""
}

func (m *HallUserData) GetUserLotteryTicket() int64 {
	if m != nil && m.UserLotteryTicket != nil {
		return *m.UserLotteryTicket
	}
	return 0
}

func (m *HallUserData) GetSex() bool {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return false
}

// 排行信息
type HallItemRank struct {
	Placing          *int32  `protobuf:"varint,1,opt,name=placing" json:"placing,omitempty"`
	UserId           *int32  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RankInfo         *string `protobuf:"bytes,3,opt,name=rankInfo" json:"rankInfo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallItemRank) Reset()                    { *m = HallItemRank{} }
func (m *HallItemRank) String() string            { return proto.CompactTextString(m) }
func (*HallItemRank) ProtoMessage()               {}
func (*HallItemRank) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *HallItemRank) GetPlacing() int32 {
	if m != nil && m.Placing != nil {
		return *m.Placing
	}
	return 0
}

func (m *HallItemRank) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallItemRank) GetRankInfo() string {
	if m != nil && m.RankInfo != nil {
		return *m.RankInfo
	}
	return ""
}

// 金币专区
type CoinZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CoinZone) Reset()                    { *m = CoinZone{} }
func (m *CoinZone) String() string            { return proto.CompactTextString(m) }
func (*CoinZone) ProtoMessage()               {}
func (*CoinZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *CoinZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 钻石专区
type DiamondZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DiamondZone) Reset()                    { *m = DiamondZone{} }
func (m *DiamondZone) String() string            { return proto.CompactTextString(m) }
func (*DiamondZone) ProtoMessage()               {}
func (*DiamondZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *DiamondZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 兑换专区
type ExchangeZone struct {
	Money            []*GoodsItem `protobuf:"bytes,1,rep,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ExchangeZone) Reset()                    { *m = ExchangeZone{} }
func (m *ExchangeZone) String() string            { return proto.CompactTextString(m) }
func (*ExchangeZone) ProtoMessage()               {}
func (*ExchangeZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *ExchangeZone) GetMoney() []*GoodsItem {
	if m != nil {
		return m.Money
	}
	return nil
}

// 购买专区
type BuyZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BuyZone) Reset()                    { *m = BuyZone{} }
func (m *BuyZone) String() string            { return proto.CompactTextString(m) }
func (*BuyZone) ProtoMessage()               {}
func (*BuyZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *BuyZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 商品类型
type GoodsItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int32  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Img              *string `protobuf:"bytes,3,opt,name=img" json:"img,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GoodsItem) Reset()                    { *m = GoodsItem{} }
func (m *GoodsItem) String() string            { return proto.CompactTextString(m) }
func (*GoodsItem) ProtoMessage()               {}
func (*GoodsItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *GoodsItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GoodsItem) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *GoodsItem) GetImg() string {
	if m != nil && m.Img != nil {
		return *m.Img
	}
	return ""
}

func init() {
	proto.RegisterType((*HallItemEvent)(nil), "ddproto.hall_item_event")
	proto.RegisterType((*HallMailItem)(nil), "ddproto.hall_mail_item")
	proto.RegisterType((*HallBagItem)(nil), "ddproto.hall_bag_item")
	proto.RegisterType((*HallItemTask)(nil), "ddproto.hall_item_task")
	proto.RegisterType((*HallUserData)(nil), "ddproto.hall_userData")
	proto.RegisterType((*HallItemRank)(nil), "ddproto.hall_item_rank")
	proto.RegisterType((*CoinZone)(nil), "ddproto.CoinZone")
	proto.RegisterType((*DiamondZone)(nil), "ddproto.DiamondZone")
	proto.RegisterType((*ExchangeZone)(nil), "ddproto.ExchangeZone")
	proto.RegisterType((*BuyZone)(nil), "ddproto.BuyZone")
	proto.RegisterType((*GoodsItem)(nil), "ddproto.GoodsItem")
	proto.RegisterEnum("ddproto.HallEnumProtoId", HallEnumProtoId_name, HallEnumProtoId_value)
	proto.RegisterEnum("ddproto.HallEnumEvent", HallEnumEvent_name, HallEnumEvent_value)
	proto.RegisterEnum("ddproto.HallEnum_Reward", HallEnum_Reward_name, HallEnum_Reward_value)
	proto.RegisterEnum("ddproto.HallEnumMailType", HallEnumMailType_name, HallEnumMailType_value)
	proto.RegisterEnum("ddproto.HallEnumPropsType", HallEnumPropsType_name, HallEnumPropsType_value)
	proto.RegisterEnum("ddproto.HallEnumTaskType", HallEnumTaskType_name, HallEnumTaskType_value)
	proto.RegisterEnum("ddproto.HallUser_VIP", HallUser_VIP_name, HallUser_VIP_value)
}

var fileDescriptor14 = []byte{
	// 1251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdf, 0x53, 0xdb, 0xc6,
	0x13, 0xff, 0xca, 0xf2, 0xcf, 0x35, 0x90, 0xe3, 0x42, 0x88, 0x48, 0x20, 0x71, 0x1c, 0xbe, 0x33,
	0x1e, 0x37, 0x43, 0x20, 0x49, 0x3b, 0x7d, 0xc9, 0x83, 0xb0, 0x0e, 0x50, 0x90, 0x65, 0x71, 0x16,
	0x26, 0xd0, 0x99, 0x6a, 0x54, 0x5b, 0x05, 0x0d, 0xb6, 0xc4, 0x18, 0x99, 0x86, 0x3f, 0xa2, 0x0f,
	0xfd, 0x3b, 0xfa, 0xda, 0xbf, 0xae, 0x4f, 0x9d, 0x3b, 0x9d, 0x65, 0x2b, 0xb6, 0x4b, 0xf3, 0xb6,
	0xbb, 0x9f, 0xfd, 0xec, 0xed, 0x7e, 0x6e, 0x6f, 0xe6, 0xe0, 0xd1, 0x95, 0xdb, 0xef, 0x3b, 0x3d,
	0x37, 0x72, 0x77, 0x6e, 0x86, 0x61, 0x14, 0xe2, 0x42, 0xaf, 0xc7, 0x8d, 0xea, 0x5f, 0x92, 0x00,
	0xfd, 0xc8, 0x1b, 0x38, 0xde, 0x9d, 0x17, 0x44, 0x78, 0x05, 0x32, 0x7e, 0x4f, 0x91, 0x2a, 0x52,
	0x2d, 0x47, 0x33, 0x7e, 0x0f, 0xbf, 0x81, 0x6c, 0x74, 0x7f, 0xe3, 0x29, 0x99, 0x8a, 0x54, 0x5b,
	0x79, 0xa7, 0xec, 0x08, 0xee, 0x0e, 0xe7, 0x79, 0xc1, 0x48, 0xf0, 0x28, 0xcf, 0xc2, 0x7b, 0x90,
	0x1f, 0x7a, 0xbf, 0xb9, 0xc3, 0x9e, 0x22, 0xf3, 0xfc, 0x8d, 0x39, 0xf9, 0x94, 0x27, 0x50, 0x91,
	0x88, 0x9f, 0x41, 0x71, 0xe8, 0x77, 0xaf, 0x6c, 0xef, 0x4b, 0xa4, 0xe4, 0x2a, 0x72, 0xad, 0x44,
	0x13, 0x1f, 0xaf, 0x41, 0x2e, 0xf2, 0xa3, 0xbe, 0xa7, 0xe4, 0x2b, 0x52, 0xad, 0x44, 0x63, 0xa7,
	0xfa, 0xb7, 0x04, 0x2b, 0xbc, 0xdc, 0xc0, 0xf5, 0xe3, 0xde, 0x67, 0xba, 0xde, 0x4d, 0x75, 0xbd,
	0x39, 0xa7, 0x0b, 0xce, 0x65, 0x39, 0xa2, 0xf3, 0xe4, 0x28, 0x79, 0xea, 0x28, 0xac, 0x40, 0xa1,
	0x1b, 0x06, 0x91, 0x17, 0x44, 0x4a, 0x96, 0xc7, 0xc7, 0x2e, 0x43, 0xfc, 0xdb, 0x33, 0x37, 0xea,
	0x5e, 0x29, 0xb9, 0x8a, 0x54, 0x2b, 0xd2, 0xb1, 0x1b, 0x23, 0x8d, 0x2b, 0xaf, 0x7b, 0xcd, 0xdb,
	0xe6, 0x08, 0x77, 0xf1, 0x0e, 0xe4, 0xdd, 0x28, 0x72, 0xbb, 0x57, 0x4a, 0xa1, 0x22, 0xd7, 0xca,
	0xef, 0xd6, 0xd3, 0x7d, 0xfd, 0xe2, 0x5e, 0xf2, 0x69, 0xa8, 0xc8, 0xc2, 0x18, 0xb2, 0x3d, 0x37,
	0xf2, 0x94, 0x62, 0x45, 0xaa, 0xc9, 0x94, 0xdb, 0xd5, 0x0b, 0x58, 0x4e, 0x25, 0xe3, 0x3d, 0x31,
	0xaa, 0xc4, 0x47, 0xdd, 0x9a, 0x33, 0xea, 0xcd, 0x30, 0xbc, 0xb9, 0x9d, 0x9e, 0x75, 0x1d, 0xf2,
	0xee, 0x20, 0x1c, 0x05, 0x11, 0xd7, 0x27, 0x47, 0x85, 0x57, 0xfd, 0x63, 0x2c, 0x2c, 0xdf, 0x87,
	0xc8, 0xbd, 0xbd, 0x9e, 0x11, 0xf6, 0x6d, 0x4a, 0xd8, 0xe7, 0x73, 0x4e, 0x63, 0x34, 0xfb, 0x21,
	0x5d, 0xd7, 0x93, 0x3d, 0x89, 0x65, 0x1d, 0x2f, 0x03, 0x02, 0xd9, 0x1c, 0x0d, 0xb8, 0xa2, 0x39,
	0xca, 0xcc, 0xea, 0xef, 0xb2, 0x18, 0x78, 0x74, 0xeb, 0x0d, 0x35, 0x37, 0x72, 0xd9, 0xc2, 0x30,
	0xdb, 0x74, 0x07, 0xf1, 0xd0, 0x25, 0x9a, 0xf8, 0xac, 0x2e, 0xb3, 0xf5, 0xde, 0x78, 0xb2, 0xd8,
	0xc3, 0x9b, 0x50, 0x0a, 0xfc, 0xae, 0xd7, 0x71, 0xfb, 0xa3, 0xb8, 0x13, 0x99, 0x4e, 0x02, 0x0c,
	0xf5, 0xee, 0xfc, 0x7e, 0x8c, 0x66, 0x63, 0x34, 0x09, 0x30, 0x94, 0x55, 0x31, 0xbc, 0x3b, 0xaf,
	0x2f, 0x3a, 0x9b, 0x04, 0xd8, 0x6d, 0x33, 0xa7, 0xa3, 0x5b, 0xe3, 0xdb, 0x16, 0x2e, 0xfe, 0x10,
	0xf3, 0x3a, 0xba, 0x65, 0xdc, 0x29, 0x05, 0xae, 0xd7, 0x57, 0x17, 0xce, 0x60, 0xa7, 0xa3, 0x5b,
	0x74, 0x92, 0x38, 0x3e, 0xad, 0x19, 0x06, 0xde, 0xbd, 0xb8, 0xf8, 0x49, 0x00, 0x57, 0xa0, 0xcc,
	0x75, 0xf0, 0xdd, 0x41, 0x18, 0xf4, 0x94, 0x12, 0xc7, 0xa7, 0x43, 0xf8, 0x05, 0x00, 0x73, 0xa9,
	0xd7, 0xdb, 0x77, 0x2f, 0x15, 0xe0, 0xfa, 0x4c, 0x45, 0xf0, 0x1b, 0x58, 0xe5, 0xcd, 0x87, 0x51,
	0xe4, 0x0d, 0xef, 0x6d, 0xbf, 0x7b, 0xed, 0x45, 0x4a, 0x99, 0xd7, 0x99, 0x05, 0xd8, 0x7d, 0xdc,
	0x7a, 0x5f, 0x94, 0x25, 0x3e, 0x19, 0x33, 0xab, 0x3f, 0x4f, 0xaf, 0xc8, 0xd0, 0x0d, 0xae, 0x99,
	0x02, 0x37, 0x7d, 0xb7, 0xeb, 0x07, 0x97, 0x62, 0x4f, 0xc6, 0xee, 0xc2, 0xdb, 0x60, 0x4f, 0xde,
	0x0d, 0xae, 0xf5, 0xe0, 0xd7, 0x50, 0xac, 0x45, 0xe2, 0x57, 0x77, 0xa1, 0xd8, 0x08, 0xfd, 0xe0,
	0x22, 0x0c, 0x3c, 0xbc, 0x0d, 0xf2, 0x8d, 0x7b, 0xaf, 0x48, 0xfc, 0xb1, 0xe0, 0x44, 0xbb, 0xc3,
	0x30, 0xec, 0xdd, 0xea, 0xec, 0xa1, 0x30, 0xb8, 0xfa, 0x1e, 0xca, 0x62, 0xf8, 0x6f, 0x20, 0xfd,
	0x08, 0x4b, 0xe4, 0x4b, 0xf7, 0xca, 0x0d, 0x2e, 0x3d, 0xce, 0xaa, 0x41, 0x6e, 0xc0, 0x25, 0x5f,
	0xcc, 0x8b, 0x13, 0xaa, 0x6f, 0xa1, 0xb0, 0x3f, 0xba, 0xff, 0x86, 0xa3, 0x1a, 0x50, 0x4a, 0x22,
	0x33, 0xef, 0x69, 0x6d, 0x7c, 0x6e, 0xac, 0x50, 0xec, 0x30, 0xd9, 0xfd, 0xc1, 0xa5, 0xd0, 0x86,
	0x99, 0xf5, 0x3f, 0x4b, 0xb0, 0x9a, 0x7a, 0xd1, 0x51, 0xa8, 0xf7, 0xf0, 0x3a, 0xe0, 0x23, 0xd5,
	0x30, 0x1c, 0x4b, 0xd7, 0x9c, 0x23, 0xa2, 0x52, 0x7b, 0x9f, 0xa8, 0x36, 0xfa, 0x1f, 0x7e, 0x0a,
	0x8f, 0x93, 0xf8, 0xc9, 0xa9, 0xde, 0x38, 0x76, 0x1a, 0x2d, 0xd3, 0x44, 0x12, 0x7e, 0x0e, 0x4f,
	0xe7, 0x00, 0x8e, 0xda, 0x38, 0x46, 0x99, 0x14, 0xeb, 0x50, 0x6d, 0x12, 0xc7, 0x68, 0x1d, 0xea,
	0x26, 0x92, 0x53, 0xac, 0x09, 0xc0, 0x59, 0x59, 0x5c, 0x85, 0x17, 0x09, 0x78, 0xf6, 0xd9, 0x52,
	0xcf, 0x4f, 0x4d, 0xfd, 0x40, 0x27, 0x5a, 0x8b, 0x6a, 0x84, 0x3a, 0x94, 0x9c, 0xa0, 0xdc, 0x03,
	0x39, 0xac, 0x4e, 0x1e, 0xbf, 0x82, 0xad, 0x74, 0x4e, 0xfb, 0xdc, 0x6c, 0x34, 0x8e, 0x48, 0xe3,
	0x58, 0x94, 0x29, 0xfc, 0x7b, 0x0a, 0xab, 0x52, 0x4c, 0x29, 0x72, 0xda, 0x26, 0xd4, 0xd1, 0x54,
	0x5b, 0x45, 0x25, 0xfc, 0x0c, 0xd6, 0x67, 0xe3, 0x9c, 0x03, 0x78, 0x03, 0x9e, 0x24, 0x98, 0x46,
	0xd5, 0x33, 0xc7, 0x68, 0xd9, 0x36, 0xa1, 0xe7, 0xa8, 0x8c, 0xb7, 0x60, 0x63, 0x2e, 0xc4, 0x99,
	0x4b, 0xb8, 0x02, 0x9b, 0x13, 0xb8, 0x9d, 0x80, 0xba, 0x79, 0xd0, 0xe2, 0x19, 0xcb, 0x29, 0xe9,
	0x5a, 0xa6, 0xa1, 0x9b, 0xe4, 0x4c, 0xa5, 0x1a, 0x9f, 0x67, 0x65, 0x11, 0xc8, 0x98, 0x8f, 0x52,
	0x93, 0x90, 0x0e, 0x31, 0x6d, 0x4e, 0x42, 0x73, 0xe2, 0x2c, 0x7f, 0x35, 0x7d, 0x49, 0xad, 0x16,
	0x6b, 0x47, 0x6f, 0xc7, 0x24, 0xbc, 0x08, 0x64, 0xcc, 0xc7, 0x78, 0x1b, 0x2a, 0x93, 0x2d, 0x6a,
	0xd9, 0xa7, 0x96, 0xa6, 0xda, 0x84, 0x92, 0x93, 0x0e, 0xa1, 0x6d, 0xbd, 0x65, 0xb2, 0x69, 0xd0,
	0xda, 0xfc, 0x2c, 0xb5, 0x71, 0x3c, 0x9d, 0xf5, 0x24, 0xa5, 0xf3, 0xbe, 0x7a, 0xe8, 0xe8, 0x36,
	0x69, 0xb6, 0x79, 0x13, 0xeb, 0x0b, 0x30, 0xd6, 0xc3, 0x53, 0x5c, 0x83, 0xed, 0xd9, 0xea, 0x6c,
	0xd7, 0xd4, 0x76, 0x9b, 0xd8, 0x6d, 0xae, 0x28, 0xab, 0xa2, 0xfc, 0xa7, 0x4c, 0x56, 0x73, 0x23,
	0x75, 0x79, 0xaa, 0x65, 0x19, 0xc4, 0xb1, 0xd4, 0x73, 0xa7, 0xb1, 0xcf, 0x0b, 0x3d, 0x5b, 0x30,
	0x50, 0xfa, 0xb8, 0xe7, 0x0f, 0x66, 0xb1, 0xa3, 0x36, 0x53, 0xa3, 0x35, 0x55, 0xdd, 0x98, 0x68,
	0xbf, 0xb5, 0x00, 0x63, 0xbc, 0x17, 0xf8, 0x35, 0xbc, 0x9c, 0x54, 0x67, 0x86, 0xd6, 0x16, 0x4b,
	0x94, 0xb4, 0xf0, 0x32, 0x35, 0xc7, 0x01, 0xd5, 0x89, 0x39, 0x7d, 0xb7, 0x95, 0xc5, 0x30, 0x3b,
	0xe2, 0x55, 0xea, 0xea, 0x63, 0xd8, 0x51, 0xb5, 0x78, 0x03, 0xab, 0x8b, 0x40, 0xc6, 0x7c, 0x3d,
	0x0f, 0xd4, 0x88, 0xc1, 0x99, 0xdb, 0x8b, 0x40, 0xc6, 0xfc, 0x7f, 0xfd, 0xa3, 0xf8, 0x57, 0x4e,
	0xfe, 0x87, 0x78, 0x19, 0x4a, 0xf6, 0xb9, 0x45, 0x1c, 0x5b, 0x6f, 0x12, 0x24, 0xe1, 0x25, 0x28,
	0x72, 0xd7, 0x24, 0x67, 0x28, 0x93, 0x80, 0xe6, 0xa9, 0x61, 0x20, 0xb9, 0xfe, 0x1d, 0xa0, 0xaf,
	0xbf, 0x8b, 0x18, 0x20, 0x4f, 0x89, 0x43, 0x3e, 0x5b, 0x48, 0xc2, 0x65, 0x28, 0x50, 0xe2, 0x1c,
	0xea, 0x07, 0x36, 0xca, 0xd4, 0xf7, 0xe0, 0xf1, 0x9c, 0x5f, 0x1d, 0xcb, 0x6f, 0x9f, 0xb7, 0x6d,
	0xd2, 0x44, 0x12, 0x5e, 0x01, 0x98, 0x0c, 0x87, 0x32, 0xf5, 0x9f, 0x60, 0x6d, 0xde, 0xef, 0x28,
	0x69, 0xc3, 0x50, 0xf7, 0x55, 0x24, 0x61, 0x04, 0x4b, 0xdc, 0xfd, 0xa4, 0x5b, 0xaa, 0x7e, 0xa2,
	0xa3, 0x0c, 0x7e, 0x04, 0x65, 0x1e, 0x39, 0x50, 0xcd, 0xc3, 0x63, 0x15, 0xc9, 0x78, 0x15, 0x96,
	0x79, 0x80, 0x7c, 0xb6, 0x9c, 0xf7, 0xbb, 0xbb, 0xbb, 0x28, 0x5b, 0xff, 0x08, 0x78, 0xf6, 0x33,
	0xc4, 0x5a, 0xe6, 0x89, 0xcd, 0x4f, 0x53, 0xc3, 0x6b, 0xda, 0x05, 0xca, 0x24, 0xde, 0xc5, 0xa7,
	0x23, 0x24, 0xd7, 0x8f, 0xa7, 0xbe, 0x3b, 0xec, 0x6f, 0x80, 0x8b, 0x90, 0x35, 0x3a, 0xce, 0x1e,
	0x92, 0x84, 0xf5, 0x0e, 0x65, 0x84, 0xf5, 0x1e, 0xc9, 0xc2, 0xfa, 0x80, 0xb2, 0xc2, 0xfa, 0x1e,
	0xe5, 0x84, 0xf5, 0x03, 0xca, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xfd, 0x81, 0x60, 0xfb,
	0x0b, 0x00, 0x00,
}
