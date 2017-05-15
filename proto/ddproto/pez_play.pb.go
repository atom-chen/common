// Code generated by protoc-gen-go.
// source: pez_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_timerInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_lastScore from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_tuozi from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_PEZTYPE from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_Bet from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of pez_enum_TuoziType from pez_base.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of pez_DissolveDesk from pez_desk.proto

// Ignoring public import of pez_AckDissolveDesk from pez_desk.proto

// Ignoring public import of pez_Ready from pez_desk.proto

// Ignoring public import of pez_AckReady from pez_desk.proto

// Ignoring public import of pez_EndLotteryInfo from pez_desk.proto

// Ignoring public import of pez_SendCurrentResult from pez_desk.proto

// Ignoring public import of EndLottery from pez_desk.proto

// Ignoring public import of pez_SendEndLottery from pez_desk.proto

// Ignoring public import of pez_UserBean from pez_desk.proto

// Ignoring public import of pez_Bill from pez_desk.proto

// 链接类型
type PEZ_RECONNECT_TYPE int32

const (
	PEZ_RECONNECT_TYPE_PEZ_NORMAL    PEZ_RECONNECT_TYPE = 1
	PEZ_RECONNECT_TYPE_PEZ_RECONNECT PEZ_RECONNECT_TYPE = 2
)

var PEZ_RECONNECT_TYPE_name = map[int32]string{
	1: "PEZ_NORMAL",
	2: "PEZ_RECONNECT",
}
var PEZ_RECONNECT_TYPE_value = map[string]int32{
	"PEZ_NORMAL":    1,
	"PEZ_RECONNECT": 2,
}

func (x PEZ_RECONNECT_TYPE) Enum() *PEZ_RECONNECT_TYPE {
	p := new(PEZ_RECONNECT_TYPE)
	*p = x
	return p
}
func (x PEZ_RECONNECT_TYPE) String() string {
	return proto.EnumName(PEZ_RECONNECT_TYPE_name, int32(x))
}
func (x *PEZ_RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PEZ_RECONNECT_TYPE_value, data, "PEZ_RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = PEZ_RECONNECT_TYPE(value)
	return nil
}
func (PEZ_RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor41, []int{0} }

// 积分
type PezUserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezUserCoinBean) Reset()                    { *m = PezUserCoinBean{} }
func (m *PezUserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*PezUserCoinBean) ProtoMessage()               {}
func (*PezUserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{0} }

func (m *PezUserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezUserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 开局（接收服务端消息）
type Pez_Opening struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrPlayCount    *int32             `protobuf:"varint,2,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,3,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_Opening) Reset()                    { *m = Pez_Opening{} }
func (m *Pez_Opening) String() string            { return proto.CompactTextString(m) }
func (*Pez_Opening) ProtoMessage()               {}
func (*Pez_Opening) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{1} }

func (m *Pez_Opening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Opening) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *Pez_Opening) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发牌
type Pez_DealCards struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Banker           *uint32           `protobuf:"varint,2,opt,name=banker" json:"banker,omitempty"`
	Dice1            *int32            `protobuf:"varint,3,opt,name=dice1" json:"dice1,omitempty"`
	Dice2            *int32            `protobuf:"varint,4,opt,name=dice2" json:"dice2,omitempty"`
	LastScore        []*PezLastScore   `protobuf:"bytes,5,rep,name=lastScore" json:"lastScore,omitempty"`
	TimerInfo        *PezBaseTimerInfo `protobuf:"bytes,6,opt,name=timerInfo" json:"timerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Pez_DealCards) Reset()                    { *m = Pez_DealCards{} }
func (m *Pez_DealCards) String() string            { return proto.CompactTextString(m) }
func (*Pez_DealCards) ProtoMessage()               {}
func (*Pez_DealCards) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{2} }

func (m *Pez_DealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DealCards) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *Pez_DealCards) GetDice1() int32 {
	if m != nil && m.Dice1 != nil {
		return *m.Dice1
	}
	return 0
}

func (m *Pez_DealCards) GetDice2() int32 {
	if m != nil && m.Dice2 != nil {
		return *m.Dice2
	}
	return 0
}

func (m *Pez_DealCards) GetLastScore() []*PezLastScore {
	if m != nil {
		return m.LastScore
	}
	return nil
}

func (m *Pez_DealCards) GetTimerInfo() *PezBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

// 押注
type Pez_Bet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetNum           *int64  `protobuf:"varint,2,opt,name=betNum" json:"betNum,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_Bet) Reset()                    { *m = Pez_Bet{} }
func (m *Pez_Bet) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bet) ProtoMessage()               {}
func (*Pez_Bet) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{3} }

func (m *Pez_Bet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_Bet) GetBetNum() int64 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *Pez_Bet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_AckBet struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BetCount         *int64       `protobuf:"varint,3,opt,name=betCount" json:"betCount,omitempty"`
	Time             *int32       `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckBet) Reset()                    { *m = Pez_AckBet{} }
func (m *Pez_AckBet) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckBet) ProtoMessage()               {}
func (*Pez_AckBet) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{4} }

func (m *Pez_AckBet) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckBet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckBet) GetBetCount() int64 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *Pez_AckBet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_BCOpenPai struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	HandPai          []*PezBase_PaiInfo `protobuf:"bytes,2,rep,name=handPai" json:"handPai,omitempty"`
	KeyValue         *int32             `protobuf:"varint,3,opt,name=keyValue" json:"keyValue,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_BCOpenPai) Reset()                    { *m = Pez_BCOpenPai{} }
func (m *Pez_BCOpenPai) String() string            { return proto.CompactTextString(m) }
func (*Pez_BCOpenPai) ProtoMessage()               {}
func (*Pez_BCOpenPai) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{5} }

func (m *Pez_BCOpenPai) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_BCOpenPai) GetHandPai() []*PezBase_PaiInfo {
	if m != nil {
		return m.HandPai
	}
	return nil
}

func (m *Pez_BCOpenPai) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

// 发送游戏信息(广播)
type Pez_SendGameInfo struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 1. 首先是牌桌的玩家数据（玩家数据包括其id昵称筹码头像等基本信息，其手牌数据，以及自己打开牌的数据，还有状态是否已经押注了，玩家在整局的总输赢）
	PlayerInfo []*PezBase_PlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	// 2. 桌面信息（包括：游戏是否结束，当前哪个玩家未押注，倒计时剩余时间）
	DeskGameInfo *PezBase_DeskGameInfo `protobuf:"bytes,3,opt,name=deskGameInfo" json:"deskGameInfo,omitempty"`
	//
	SenderUserId     *uint32 `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool   `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_SendGameInfo) Reset()                    { *m = Pez_SendGameInfo{} }
func (m *Pez_SendGameInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendGameInfo) ProtoMessage()               {}
func (*Pez_SendGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{6} }

func (m *Pez_SendGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_SendGameInfo) GetPlayerInfo() []*PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetDeskGameInfo() *PezBase_DeskGameInfo {
	if m != nil {
		return m.DeskGameInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *Pez_SendGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func init() {
	proto.RegisterType((*PezUserCoinBean)(nil), "ddproto.pez_user_coin_bean")
	proto.RegisterType((*Pez_Opening)(nil), "ddproto.pez_Opening")
	proto.RegisterType((*Pez_DealCards)(nil), "ddproto.pez_DealCards")
	proto.RegisterType((*Pez_Bet)(nil), "ddproto.pez_Bet")
	proto.RegisterType((*Pez_AckBet)(nil), "ddproto.pez_AckBet")
	proto.RegisterType((*Pez_BCOpenPai)(nil), "ddproto.pez_BCOpenPai")
	proto.RegisterType((*Pez_SendGameInfo)(nil), "ddproto.pez_SendGameInfo")
	proto.RegisterEnum("ddproto.PEZ_RECONNECT_TYPE", PEZ_RECONNECT_TYPE_name, PEZ_RECONNECT_TYPE_value)
}

var fileDescriptor41 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x49, 0xdb, 0xdd, 0xbd, 0xdd, 0x2c, 0xbb, 0xe3, 0x52, 0x62, 0x15, 0x29, 0xc1, 0x87,
	0x22, 0x52, 0xb0, 0x0a, 0x22, 0xf8, 0x60, 0x9b, 0x2d, 0xee, 0x82, 0xdb, 0x86, 0xd9, 0x55, 0xd0,
	0x97, 0x30, 0x4d, 0xae, 0x6e, 0x68, 0x32, 0x29, 0xf9, 0x00, 0xeb, 0xa3, 0x4f, 0xfe, 0x10, 0x7f,
	0xa1, 0xbf, 0x40, 0x66, 0x32, 0x4d, 0x1b, 0xea, 0x3e, 0xf4, 0xa5, 0xcc, 0x39, 0x73, 0xee, 0xed,
	0x39, 0xf7, 0x66, 0xe0, 0x64, 0x89, 0x3f, 0xbd, 0x65, 0xc4, 0x56, 0x83, 0x65, 0x9a, 0xe4, 0x09,
	0x39, 0x08, 0x02, 0x79, 0xe8, 0x3e, 0xf4, 0x93, 0x38, 0x4e, 0xb8, 0xe7, 0x47, 0x21, 0xf2, 0xbc,
	0xbc, 0xed, 0x4a, 0xf5, 0x9c, 0x65, 0xa8, 0xf0, 0x99, 0x12, 0x21, 0x2f, 0xe2, 0x6d, 0x49, 0x80,
	0xd9, 0xa2, 0xc4, 0xf6, 0x7b, 0x20, 0x82, 0x29, 0x32, 0x4c, 0x3d, 0x3f, 0x09, 0xb9, 0x37, 0x47,
	0xc6, 0x49, 0x07, 0x5a, 0x82, 0xb9, 0x0a, 0x2c, 0xad, 0xa7, 0xf5, 0x4d, 0xaa, 0x10, 0x21, 0xd0,
	0x10, 0x22, 0x4b, 0xef, 0x69, 0x7d, 0x83, 0xca, 0xb3, 0xfd, 0x47, 0x83, 0xb6, 0x68, 0x31, 0x5b,
	0x22, 0x0f, 0xf9, 0x77, 0xf2, 0x02, 0x5a, 0x77, 0xc8, 0x02, 0x4c, 0x65, 0x6d, 0x7b, 0x78, 0x3e,
	0x50, 0x9e, 0x07, 0xae, 0xf8, 0xbd, 0x94, 0x77, 0x54, 0x69, 0xc8, 0x33, 0x30, 0x9d, 0x22, 0x4d,
	0xdd, 0x88, 0xad, 0x9c, 0xa4, 0xe0, 0xb9, 0x6c, 0xdd, 0xa4, 0x75, 0x92, 0x8c, 0xc0, 0x14, 0x0e,
	0x9c, 0x24, 0xe4, 0x63, 0x64, 0x3c, 0xb3, 0x8c, 0x9e, 0xd1, 0x6f, 0x0f, 0x1f, 0x57, 0xad, 0x77,
	0x33, 0xd0, 0x7a, 0x85, 0xfd, 0x57, 0x03, 0x53, 0xa8, 0x2e, 0x90, 0x45, 0x0e, 0x4b, 0x83, 0x6c,
	0x4f, 0xa3, 0x1d, 0x68, 0xcd, 0x19, 0x5f, 0x60, 0x2a, 0x1d, 0x9a, 0x54, 0x21, 0x72, 0x0e, 0xcd,
	0x20, 0xf4, 0xf1, 0xa5, 0x65, 0x48, 0xe3, 0x25, 0x58, 0xb3, 0x43, 0xab, 0xb1, 0x61, 0x87, 0xe4,
	0x35, 0x1c, 0x45, 0x2c, 0xcb, 0x6f, 0xfc, 0x24, 0x45, 0xab, 0x29, 0x23, 0x74, 0x6a, 0x11, 0xaa,
	0x5b, 0xba, 0x11, 0x92, 0xb7, 0x70, 0x94, 0x87, 0x31, 0xa6, 0x57, 0xfc, 0x5b, 0x62, 0xb5, 0xa4,
	0xd5, 0x7a, 0x70, 0xb1, 0x71, 0xaf, 0x92, 0xd0, 0x8d, 0xda, 0xbe, 0x86, 0x03, 0x21, 0x18, 0x63,
	0x7e, 0xef, 0x4a, 0x45, 0x2e, 0xcc, 0xa7, 0x45, 0xac, 0x96, 0xaa, 0x90, 0x58, 0xb5, 0xe8, 0xa3,
	0x62, 0xc9, 0xb3, 0xfd, 0x4b, 0x03, 0x10, 0xfd, 0x46, 0xfe, 0x42, 0xb4, 0xdc, 0x7b, 0x80, 0xca,
	0x80, 0x5e, 0x33, 0xd0, 0x85, 0xc3, 0x39, 0xe6, 0xe5, 0xf2, 0x0d, 0x69, 0xa1, 0xc2, 0x95, 0x89,
	0xc6, 0x96, 0x89, 0x1f, 0xe5, 0x1e, 0xc7, 0x8e, 0xf8, 0xe0, 0x5c, 0x16, 0xde, 0x9b, 0xec, 0x15,
	0x1c, 0xdc, 0x31, 0x1e, 0xb8, 0x2c, 0xb4, 0x74, 0x39, 0xeb, 0x47, 0xbb, 0x53, 0x73, 0x59, 0x28,
	0x67, 0xb6, 0x56, 0x0a, 0x37, 0x0b, 0x5c, 0x7d, 0x66, 0x51, 0xb1, 0x8e, 0x5e, 0x61, 0xfb, 0xb7,
	0x0e, 0xa7, 0xa2, 0xf2, 0x06, 0x79, 0xf0, 0x81, 0xc5, 0x28, 0x2a, 0xf7, 0x1c, 0xc2, 0x3b, 0x00,
	0xf1, 0x9a, 0xd5, 0x32, 0x4b, 0x5b, 0x4f, 0xfe, 0x63, 0xab, 0xd2, 0xd0, 0x2d, 0x3d, 0x19, 0xc3,
	0xb1, 0x78, 0xba, 0xeb, 0xff, 0x96, 0x06, 0xdb, 0xc3, 0xa7, 0xbb, 0xf5, 0x17, 0x5b, 0x2a, 0x5a,
	0xab, 0x21, 0x36, 0x1c, 0x67, 0xc8, 0x03, 0x4c, 0x3f, 0x95, 0x33, 0x6b, 0xc8, 0x99, 0xd5, 0x38,
	0xd2, 0x83, 0x76, 0x98, 0x51, 0xf4, 0x13, 0xce, 0xd1, 0xcf, 0xad, 0x66, 0x4f, 0xeb, 0x1f, 0xd2,
	0x6d, 0xea, 0xf9, 0x1b, 0x20, 0xee, 0xe4, 0xab, 0x47, 0x27, 0xce, 0x6c, 0x3a, 0x9d, 0x38, 0xb7,
	0xde, 0xed, 0x17, 0x77, 0x42, 0x4e, 0x00, 0x04, 0x3b, 0x9d, 0xd1, 0xeb, 0xd1, 0xc7, 0x53, 0x8d,
	0x9c, 0x81, 0x59, 0x53, 0x9d, 0xea, 0x63, 0xfd, 0xd2, 0x70, 0x1f, 0xb8, 0x9a, 0xab, 0xbb, 0xc6,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xcf, 0xfa, 0x06, 0xdd, 0x04, 0x00, 0x00,
}
