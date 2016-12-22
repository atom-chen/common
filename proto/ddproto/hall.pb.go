// Code generated by protoc-gen-go.
// source: hall.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of hall_lottery_item from hall_data.proto

// Ignoring public import of hall_item_event from hall_data.proto

// Ignoring public import of hall_item_email from hall_data.proto

// Ignoring public import of hall_item_bag from hall_data.proto

// Ignoring public import of hall_item_task from hall_data.proto

// Ignoring public import of hall_userData from hall_data.proto

// Ignoring public import of hall_item_rank from hall_data.proto

// Ignoring public import of hall_shop from hall_data.proto

// Ignoring public import of CoinZone from hall_data.proto

// Ignoring public import of DiamondZone from hall_data.proto

// Ignoring public import of ExchangeZone from hall_data.proto

// Ignoring public import of BuyZone from hall_data.proto

// Ignoring public import of GoodsItem from hall_data.proto

// Ignoring public import of hall_enum_protoId from hall_data.proto

// Ignoring public import of hall_enum_event from hall_data.proto

// Ignoring public import of hall_enum_Reward from hall_data.proto

// Ignoring public import of hall_enum_emailType from hall_data.proto

// Ignoring public import of hall_enum_bagItemType from hall_data.proto

// Ignoring public import of hall_enum_bagItemRank from hall_data.proto

// Ignoring public import of hall_enum_taskType from hall_data.proto

// Ignoring public import of hall_enum_shopType from hall_data.proto

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 活动列表
type HallReqEvent struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqEvent) Reset()                    { *m = HallReqEvent{} }
func (m *HallReqEvent) String() string            { return proto.CompactTextString(m) }
func (*HallReqEvent) ProtoMessage()               {}
func (*HallReqEvent) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *HallReqEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckEvent struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	EventList        []*HallItemEvent `protobuf:"bytes,2,rep,name=eventList" json:"eventList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallAckEvent) Reset()                    { *m = HallAckEvent{} }
func (m *HallAckEvent) String() string            { return proto.CompactTextString(m) }
func (*HallAckEvent) ProtoMessage()               {}
func (*HallAckEvent) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *HallAckEvent) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckEvent) GetEventList() []*HallItemEvent {
	if m != nil {
		return m.EventList
	}
	return nil
}

// 邮件
type HallReqMail struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqMail) Reset()                    { *m = HallReqMail{} }
func (m *HallReqMail) String() string            { return proto.CompactTextString(m) }
func (*HallReqMail) ProtoMessage()               {}
func (*HallReqMail) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *HallReqMail) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckMail struct {
	Header           *ProtoHeader     `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	EmailList        []*HallItemEmail `protobuf:"bytes,2,rep,name=emailList" json:"emailList,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallAckMail) Reset()                    { *m = HallAckMail{} }
func (m *HallAckMail) String() string            { return proto.CompactTextString(m) }
func (*HallAckMail) ProtoMessage()               {}
func (*HallAckMail) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *HallAckMail) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckMail) GetEmailList() []*HallItemEmail {
	if m != nil {
		return m.EmailList
	}
	return nil
}

// 任务
type HallReqTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqTask) Reset()                    { *m = HallReqTask{} }
func (m *HallReqTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqTask) ProtoMessage()               {}
func (*HallReqTask) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *HallReqTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckTask struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TaskList         []*HallItemTask `protobuf:"bytes,2,rep,name=taskList" json:"taskList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckTask) Reset()                    { *m = HallAckTask{} }
func (m *HallAckTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckTask) ProtoMessage()               {}
func (*HallAckTask) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *HallAckTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckTask) GetTaskList() []*HallItemTask {
	if m != nil {
		return m.TaskList
	}
	return nil
}

// 领取任务
type HallReqGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqGetTask) Reset()                    { *m = HallReqGetTask{} }
func (m *HallReqGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallReqGetTask) ProtoMessage()               {}
func (*HallReqGetTask) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *HallReqGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckGetTask struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FinishedNum      *int32       `protobuf:"varint,2,opt,name=finishedNum" json:"finishedNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckGetTask) Reset()                    { *m = HallAckGetTask{} }
func (m *HallAckGetTask) String() string            { return proto.CompactTextString(m) }
func (*HallAckGetTask) ProtoMessage()               {}
func (*HallAckGetTask) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

func (m *HallAckGetTask) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckGetTask) GetFinishedNum() int32 {
	if m != nil && m.FinishedNum != nil {
		return *m.FinishedNum
	}
	return 0
}

// 背包
type HallReqBag struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqBag) Reset()                    { *m = HallReqBag{} }
func (m *HallReqBag) String() string            { return proto.CompactTextString(m) }
func (*HallReqBag) ProtoMessage()               {}
func (*HallReqBag) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

func (m *HallReqBag) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckBag struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	BagItemList      []*HallItemBag `protobuf:"bytes,2,rep,name=bagItemList" json:"bagItemList,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *HallAckBag) Reset()                    { *m = HallAckBag{} }
func (m *HallAckBag) String() string            { return proto.CompactTextString(m) }
func (*HallAckBag) ProtoMessage()               {}
func (*HallAckBag) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{9} }

func (m *HallAckBag) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckBag) GetBagItemList() []*HallItemBag {
	if m != nil {
		return m.BagItemList
	}
	return nil
}

// 个人信息
type HallReqUserData struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqUserData) Reset()                    { *m = HallReqUserData{} }
func (m *HallReqUserData) String() string            { return proto.CompactTextString(m) }
func (*HallReqUserData) ProtoMessage()               {}
func (*HallReqUserData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func (m *HallReqUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckUserData struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *HallUserData `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *HallAckUserData) Reset()                    { *m = HallAckUserData{} }
func (m *HallAckUserData) String() string            { return proto.CompactTextString(m) }
func (*HallAckUserData) ProtoMessage()               {}
func (*HallAckUserData) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{11} }

func (m *HallAckUserData) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckUserData) GetUser() *HallUserData {
	if m != nil {
		return m.User
	}
	return nil
}

// 商城
type HallReqShop struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqShop) Reset()                    { *m = HallReqShop{} }
func (m *HallReqShop) String() string            { return proto.CompactTextString(m) }
func (*HallReqShop) ProtoMessage()               {}
func (*HallReqShop) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{12} }

func (m *HallReqShop) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAck_Shop struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ShopInfo         *HallShop    `protobuf:"bytes,2,opt,name=shopInfo" json:"shopInfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAck_Shop) Reset()                    { *m = HallAck_Shop{} }
func (m *HallAck_Shop) String() string            { return proto.CompactTextString(m) }
func (*HallAck_Shop) ProtoMessage()               {}
func (*HallAck_Shop) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{13} }

func (m *HallAck_Shop) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAck_Shop) GetShopInfo() *HallShop {
	if m != nil {
		return m.ShopInfo
	}
	return nil
}

// 排行
type HallReqRank struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqRank) Reset()                    { *m = HallReqRank{} }
func (m *HallReqRank) String() string            { return proto.CompactTextString(m) }
func (*HallReqRank) ProtoMessage()               {}
func (*HallReqRank) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{14} }

func (m *HallReqRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckRank struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RankList         []*HallItemRank `protobuf:"bytes,2,rep,name=rankList" json:"rankList,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallAckRank) Reset()                    { *m = HallAckRank{} }
func (m *HallAckRank) String() string            { return proto.CompactTextString(m) }
func (*HallAckRank) ProtoMessage()               {}
func (*HallAckRank) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{15} }

func (m *HallAckRank) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckRank) GetRankList() []*HallItemRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

// 转盘签到奖励物品信息
type HallAckDsLotteryItems struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DialItems        []*HallLotteryItem `protobuf:"bytes,2,rep,name=dialItems" json:"dialItems,omitempty"`
	SignItems        []*HallLotteryItem `protobuf:"bytes,3,rep,name=SignItems" json:"SignItems,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallAckDsLotteryItems) Reset()                    { *m = HallAckDsLotteryItems{} }
func (m *HallAckDsLotteryItems) String() string            { return proto.CompactTextString(m) }
func (*HallAckDsLotteryItems) ProtoMessage()               {}
func (*HallAckDsLotteryItems) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{16} }

func (m *HallAckDsLotteryItems) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDsLotteryItems) GetDialItems() []*HallLotteryItem {
	if m != nil {
		return m.DialItems
	}
	return nil
}

func (m *HallAckDsLotteryItems) GetSignItems() []*HallLotteryItem {
	if m != nil {
		return m.SignItems
	}
	return nil
}

// 抽奖动作
type HallReqDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallReqDrawLottery) Reset()                    { *m = HallReqDrawLottery{} }
func (m *HallReqDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallReqDrawLottery) ProtoMessage()               {}
func (*HallReqDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{17} }

func (m *HallReqDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type HallAckDrawLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	LotteryId        *int32       `protobuf:"varint,2,opt,name=lotteryId" json:"lotteryId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HallAckDrawLottery) Reset()                    { *m = HallAckDrawLottery{} }
func (m *HallAckDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*HallAckDrawLottery) ProtoMessage()               {}
func (*HallAckDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{18} }

func (m *HallAckDrawLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HallAckDrawLottery) GetLotteryId() int32 {
	if m != nil && m.LotteryId != nil {
		return *m.LotteryId
	}
	return 0
}

func init() {
	proto.RegisterType((*HallReqEvent)(nil), "ddproto.hall_req_event")
	proto.RegisterType((*HallAckEvent)(nil), "ddproto.hall_ack_event")
	proto.RegisterType((*HallReqMail)(nil), "ddproto.hall_req_mail")
	proto.RegisterType((*HallAckMail)(nil), "ddproto.hall_ack_mail")
	proto.RegisterType((*HallReqTask)(nil), "ddproto.hall_req_task")
	proto.RegisterType((*HallAckTask)(nil), "ddproto.hall_ack_task")
	proto.RegisterType((*HallReqGetTask)(nil), "ddproto.hall_req_getTask")
	proto.RegisterType((*HallAckGetTask)(nil), "ddproto.hall_ack_getTask")
	proto.RegisterType((*HallReqBag)(nil), "ddproto.hall_req_bag")
	proto.RegisterType((*HallAckBag)(nil), "ddproto.hall_ack_bag")
	proto.RegisterType((*HallReqUserData)(nil), "ddproto.hall_req_userData")
	proto.RegisterType((*HallAckUserData)(nil), "ddproto.hall_ack_userData")
	proto.RegisterType((*HallReqShop)(nil), "ddproto.hall_req_shop")
	proto.RegisterType((*HallAck_Shop)(nil), "ddproto.hall_ack_Shop")
	proto.RegisterType((*HallReqRank)(nil), "ddproto.hall_req_rank")
	proto.RegisterType((*HallAckRank)(nil), "ddproto.hall_ack_rank")
	proto.RegisterType((*HallAckDsLotteryItems)(nil), "ddproto.hall_ack_ds_lottery_items")
	proto.RegisterType((*HallReqDrawLottery)(nil), "ddproto.hall_req_draw_lottery")
	proto.RegisterType((*HallAckDrawLottery)(nil), "ddproto.hall_ack_draw_lottery")
}

var fileDescriptor12 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0x5b, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x9d, 0xd6, 0x4b, 0xf7, 0x6c, 0xbd, 0xc5, 0xdb, 0xba, 0xf8, 0xb0, 0xe4, 0x69, 0x11,
	0xd9, 0x87, 0x0a, 0xe2, 0x83, 0x8a, 0x82, 0x82, 0x0b, 0x22, 0x4b, 0xea, 0xfb, 0x92, 0x99, 0xa4,
	0x33, 0xa1, 0x73, 0xa9, 0x49, 0xb6, 0xe2, 0x87, 0xf3, 0xbb, 0xc9, 0xc9, 0xa4, 0x99, 0x19, 0x2c,
	0x2c, 0x19, 0x5f, 0x4a, 0x49, 0xfe, 0xbf, 0xff, 0xfc, 0x38, 0x9c, 0x2c, 0x40, 0xc1, 0xcb, 0x72,
	0x75, 0xa1, 0x1b, 0xdb, 0x90, 0x3b, 0x42, 0xb8, 0x7f, 0xe6, 0xf7, 0xf1, 0x70, 0x2b, 0xb8, 0xe5,
	0xed, 0xcd, 0xfc, 0x51, 0xd6, 0x54, 0x55, 0x53, 0x6f, 0xb3, 0x52, 0xc9, 0xda, 0xb6, 0x87, 0xf4,
	0x03, 0xdc, 0x73, 0x39, 0x2d, 0x7f, 0x6e, 0xe5, 0xa5, 0xac, 0x2d, 0x79, 0x05, 0xb7, 0x0b, 0xc9,
	0x85, 0xd4, 0xb3, 0x64, 0x91, 0x2c, 0xa7, 0x27, 0x8f, 0x57, 0xbe, 0x71, 0xb5, 0xc1, 0xbf, 0x5f,
	0xdd, 0x1d, 0xf3, 0x19, 0x7a, 0xe9, 0x79, 0x9e, 0x9d, 0x8f, 0xe1, 0xc9, 0x1b, 0x98, 0x38, 0xec,
	0x9b, 0x32, 0x76, 0x76, 0xb0, 0x38, 0x5c, 0x4e, 0x4f, 0x66, 0x01, 0x70, 0xcd, 0xca, 0xca, 0xaa,
	0xad, 0x66, 0x5d, 0x94, 0xbe, 0x87, 0xbb, 0xc1, 0xbb, 0xe2, 0xaa, 0x8c, 0xd4, 0xde, 0x79, 0x1c,
	0xb5, 0xe3, 0x71, 0x67, 0x8d, 0xd8, 0x3e, 0x6b, 0xcc, 0xb0, 0x2e, 0x3a, 0xb0, 0xb6, 0xdc, 0x9c,
	0x47, 0x5a, 0xeb, 0x9e, 0x75, 0x3c, 0x4e, 0x5e, 0xc3, 0x11, 0x52, 0x3d, 0xe9, 0x67, 0xd7, 0x48,
	0x63, 0x84, 0x85, 0x20, 0xfd, 0x08, 0x0f, 0x82, 0x72, 0x2e, 0xed, 0x8f, 0x78, 0xeb, 0xd4, 0x37,
	0xa0, 0xf5, 0xa8, 0x06, 0xb2, 0x80, 0xe9, 0x99, 0xaa, 0x95, 0x29, 0xa4, 0xf8, 0xbe, 0xab, 0x66,
	0x07, 0x8b, 0x64, 0x79, 0x8b, 0xf5, 0x8f, 0xe8, 0x3b, 0x38, 0x0e, 0x96, 0x29, 0xcf, 0xa3, 0x97,
	0xf8, 0x38, 0x18, 0x46, 0xd3, 0xe4, 0x2d, 0x4c, 0x53, 0x9e, 0xaf, 0xad, 0xac, 0x7a, 0x93, 0x7d,
	0x7a, 0xcd, 0x64, 0x53, 0x9e, 0xb3, 0x7e, 0x94, 0x7e, 0x82, 0x87, 0xc1, 0x7a, 0x67, 0xa4, 0xfe,
	0xcc, 0x2d, 0x8f, 0x54, 0xaf, 0x7c, 0x05, 0xaa, 0x8f, 0xab, 0x20, 0x2f, 0xe1, 0x26, 0x92, 0x6e,
	0xac, 0xff, 0x88, 0x5f, 0x75, 0x32, 0x97, 0x19, 0x2c, 0xb0, 0x29, 0x9a, 0x8b, 0x68, 0xdb, 0x6e,
	0x81, 0x4f, 0xa3, 0x71, 0xb2, 0x82, 0x23, 0xfc, 0xe8, 0xba, 0x3e, 0x6b, 0xbc, 0x2d, 0x19, 0xda,
	0xe2, 0x2d, 0x0b, 0x99, 0x81, 0xad, 0xe6, 0xf5, 0xff, 0x3c, 0xb7, 0x78, 0x1c, 0x9f, 0x1b, 0x52,
	0x7b, 0x9e, 0x1b, 0x46, 0x58, 0x08, 0xd2, 0x3f, 0x09, 0x3c, 0x0f, 0x1f, 0x15, 0x66, 0x5b, 0x36,
	0xd6, 0x4a, 0xfd, 0xdb, 0x65, 0x4d, 0xf4, 0x62, 0x4e, 0x84, 0xe2, 0x25, 0xae, 0x9b, 0xf1, 0x06,
	0xf3, 0xa1, 0x41, 0xbf, 0x9d, 0x75, 0x61, 0x24, 0x4f, 0x55, 0x5e, 0xb7, 0xe4, 0xe1, 0x7e, 0x32,
	0x84, 0xe9, 0x17, 0x78, 0x12, 0x46, 0x2e, 0x34, 0xff, 0x75, 0x15, 0x8c, 0x1c, 0x7d, 0xe6, 0x6b,
	0xdc, 0x14, 0x46, 0xd7, 0x90, 0x17, 0x30, 0xf1, 0xe0, 0x5a, 0xf8, 0x9f, 0x8d, 0xee, 0x60, 0x73,
	0x63, 0x93, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xd8, 0x25, 0x8d, 0x3b, 0x07, 0x00, 0x00,
}
