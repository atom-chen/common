// Code generated by protoc-gen-go. DO NOT EDIT.
// source: erddz_play.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of erddz_base_roomTypeInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_commonRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_timerInfo from erddz_base.proto

// Ignoring public import of erddz_base_deskInfo from erddz_base.proto

// Ignoring public import of erddz_enum_protoId from erddz_base.proto

// Ignoring public import of erddz_enum_errorCode from erddz_base.proto

// Ignoring public import of erddz_enum_paiType from erddz_base.proto

// Ignoring public import of erddz_enum_actType from erddz_base.proto

// Ignoring public import of erddz_enum_deskStatus from erddz_base.proto

// Ignoring public import of erddz_enum_playerStatus from erddz_base.proto

// Ignoring public import of erddz_enum_roomType from erddz_base.proto

// Ignoring public import of erddz_enum_coinRoomLevel from erddz_base.proto

// 开局（接收服务端消息）
type ErddzBcOpening struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32                `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Ticket           *int64                 `protobuf:"varint,3,opt,name=ticket" json:"ticket,omitempty"`
	UserCoin         *int64                 `protobuf:"varint,4,opt,name=userCoin" json:"userCoin,omitempty"`
	UserRoomCard     *int64                 `protobuf:"varint,5,opt,name=userRoomCard" json:"userRoomCard,omitempty"`
	PlayerInfos      []*ErddzBasePlayerInfo `protobuf:"bytes,6,rep,name=playerInfos" json:"playerInfos,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzBcOpening) Reset()                    { *m = ErddzBcOpening{} }
func (m *ErddzBcOpening) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcOpening) ProtoMessage()               {}
func (*ErddzBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *ErddzBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcOpening) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzBcOpening) GetTicket() int64 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *ErddzBcOpening) GetUserCoin() int64 {
	if m != nil && m.UserCoin != nil {
		return *m.UserCoin
	}
	return 0
}

func (m *ErddzBcOpening) GetUserRoomCard() int64 {
	if m != nil && m.UserRoomCard != nil {
		return *m.UserRoomCard
	}
	return 0
}

func (m *ErddzBcOpening) GetPlayerInfos() []*ErddzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// 发牌Ï
type ErddzBcDealCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []*ClientBasePoker `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	ShowPokerUserId  *uint32            `protobuf:"varint,3,opt,name=showPokerUserId" json:"showPokerUserId,omitempty"`
	ShowPokerId      *int32             `protobuf:"varint,4,opt,name=showPokerId" json:"showPokerId,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ErddzBcDealCards) Reset()                    { *m = ErddzBcDealCards{} }
func (m *ErddzBcDealCards) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcDealCards) ProtoMessage()               {}
func (*ErddzBcDealCards) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

func (m *ErddzBcDealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcDealCards) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *ErddzBcDealCards) GetShowPokerUserId() uint32 {
	if m != nil && m.ShowPokerUserId != nil {
		return *m.ShowPokerUserId
	}
	return 0
}

func (m *ErddzBcDealCards) GetShowPokerId() int32 {
	if m != nil && m.ShowPokerId != nil {
		return *m.ShowPokerId
	}
	return 0
}

// 叫地主
type ErddzReqJiaoDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool        `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqJiaoDiZhu) Reset()                    { *m = ErddzReqJiaoDiZhu{} }
func (m *ErddzReqJiaoDiZhu) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqJiaoDiZhu) ProtoMessage()               {}
func (*ErddzReqJiaoDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

func (m *ErddzReqJiaoDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqJiaoDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqJiaoDiZhu) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

// 叫地主回复
type ErddzAckJiaoDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Jiao             *bool        `protobuf:"varint,3,opt,name=jiao" json:"jiao,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckJiaoDiZhu) Reset()                    { *m = ErddzAckJiaoDiZhu{} }
func (m *ErddzAckJiaoDiZhu) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckJiaoDiZhu) ProtoMessage()               {}
func (*ErddzAckJiaoDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

func (m *ErddzAckJiaoDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckJiaoDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckJiaoDiZhu) GetJiao() bool {
	if m != nil && m.Jiao != nil {
		return *m.Jiao
	}
	return false
}

// 抢地主
type ErddzReqRobDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rob              *bool        `protobuf:"varint,3,opt,name=rob" json:"rob,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqRobDiZhu) Reset()                    { *m = ErddzReqRobDiZhu{} }
func (m *ErddzReqRobDiZhu) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqRobDiZhu) ProtoMessage()               {}
func (*ErddzReqRobDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

func (m *ErddzReqRobDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqRobDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqRobDiZhu) GetRob() bool {
	if m != nil && m.Rob != nil {
		return *m.Rob
	}
	return false
}

type ErddzAckRobDiZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Rob              *bool        `protobuf:"varint,3,opt,name=rob" json:"rob,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckRobDiZhu) Reset()                    { *m = ErddzAckRobDiZhu{} }
func (m *ErddzAckRobDiZhu) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckRobDiZhu) ProtoMessage()               {}
func (*ErddzAckRobDiZhu) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

func (m *ErddzAckRobDiZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckRobDiZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckRobDiZhu) GetRob() bool {
	if m != nil && m.Rob != nil {
		return *m.Rob
	}
	return false
}

// 让牌
type ErddzReqRangcards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RangNum          *uint32      `protobuf:"varint,3,opt,name=rangNum" json:"rangNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqRangcards) Reset()                    { *m = ErddzReqRangcards{} }
func (m *ErddzReqRangcards) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqRangcards) ProtoMessage()               {}
func (*ErddzReqRangcards) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

func (m *ErddzReqRangcards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqRangcards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqRangcards) GetRangNum() uint32 {
	if m != nil && m.RangNum != nil {
		return *m.RangNum
	}
	return 0
}

type ErddzAckRangcards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RangNum          *uint32      `protobuf:"varint,3,opt,name=rangNum" json:"rangNum,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckRangcards) Reset()                    { *m = ErddzAckRangcards{} }
func (m *ErddzAckRangcards) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckRangcards) ProtoMessage()               {}
func (*ErddzAckRangcards) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

func (m *ErddzAckRangcards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckRangcards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckRangcards) GetRangNum() uint32 {
	if m != nil && m.RangNum != nil {
		return *m.RangNum
	}
	return 0
}

// 加倍
type ErddzReqDouble struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Double           *bool        `protobuf:"varint,3,opt,name=double" json:"double,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqDouble) Reset()                    { *m = ErddzReqDouble{} }
func (m *ErddzReqDouble) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqDouble) ProtoMessage()               {}
func (*ErddzReqDouble) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

func (m *ErddzReqDouble) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqDouble) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqDouble) GetDouble() bool {
	if m != nil && m.Double != nil {
		return *m.Double
	}
	return false
}

type ErddzAckDouble struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Double           *bool        `protobuf:"varint,3,opt,name=double" json:"double,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckDouble) Reset()                    { *m = ErddzAckDouble{} }
func (m *ErddzAckDouble) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckDouble) ProtoMessage()               {}
func (*ErddzAckDouble) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

func (m *ErddzAckDouble) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckDouble) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckDouble) GetDouble() bool {
	if m != nil && m.Double != nil {
		return *m.Double
	}
	return false
}

// 叫地主结束，开始游戏 (广播)
type ErddzBcStartPlay struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	FootPokers       []*ClientBasePoker `protobuf:"bytes,2,rep,name=footPokers" json:"footPokers,omitempty"`
	FootRate         *int32             `protobuf:"varint,3,opt,name=footRate" json:"footRate,omitempty"`
	Dizhu            *uint32            `protobuf:"varint,4,opt,name=dizhu" json:"dizhu,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ErddzBcStartPlay) Reset()                    { *m = ErddzBcStartPlay{} }
func (m *ErddzBcStartPlay) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcStartPlay) ProtoMessage()               {}
func (*ErddzBcStartPlay) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

func (m *ErddzBcStartPlay) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcStartPlay) GetFootPokers() []*ClientBasePoker {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

func (m *ErddzBcStartPlay) GetFootRate() int32 {
	if m != nil && m.FootRate != nil {
		return *m.FootRate
	}
	return 0
}

func (m *ErddzBcStartPlay) GetDizhu() uint32 {
	if m != nil && m.Dizhu != nil {
		return *m.Dizhu
	}
	return 0
}

// 出牌
type ErddzReqOutCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	OutCards         []*ClientBasePoker `protobuf:"bytes,2,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ErddzReqOutCards) Reset()                    { *m = ErddzReqOutCards{} }
func (m *ErddzReqOutCards) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqOutCards) ProtoMessage()               {}
func (*ErddzReqOutCards) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{11} }

func (m *ErddzReqOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqOutCards) GetOutCards() []*ClientBasePoker {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type ErddzAckOutCards struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	CardType         *ErddzEnumPaiType  `protobuf:"varint,3,opt,name=cardType,enum=ddproto.ErddzEnumPaiType" json:"cardType,omitempty"`
	RemainPokers     *int32             `protobuf:"varint,4,opt,name=remainPokers" json:"remainPokers,omitempty"`
	OutCards         []*ClientBasePoker `protobuf:"bytes,5,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ErddzAckOutCards) Reset()                    { *m = ErddzAckOutCards{} }
func (m *ErddzAckOutCards) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckOutCards) ProtoMessage()               {}
func (*ErddzAckOutCards) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{12} }

func (m *ErddzAckOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckOutCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckOutCards) GetCardType() ErddzEnumPaiType {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return ErddzEnumPaiType_ERDDZ_ERRORCARD
}

func (m *ErddzAckOutCards) GetRemainPokers() int32 {
	if m != nil && m.RemainPokers != nil {
		return *m.RemainPokers
	}
	return 0
}

func (m *ErddzAckOutCards) GetOutCards() []*ClientBasePoker {
	if m != nil {
		return m.OutCards
	}
	return nil
}

// 过牌
type ErddzReqActGuo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqActGuo) Reset()                    { *m = ErddzReqActGuo{} }
func (m *ErddzReqActGuo) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqActGuo) ProtoMessage()               {}
func (*ErddzReqActGuo) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{13} }

func (m *ErddzReqActGuo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqActGuo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type ErddzAckActGuo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckActGuo) Reset()                    { *m = ErddzAckActGuo{} }
func (m *ErddzAckActGuo) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckActGuo) ProtoMessage()               {}
func (*ErddzAckActGuo) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{14} }

func (m *ErddzAckActGuo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckActGuo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 轮到谁操作
type ErddzBcOverTurn struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32                `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Time             *int32                 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	ActType          *ErddzEnumActType      `protobuf:"varint,4,opt,name=actType,enum=ddproto.ErddzEnumActType" json:"actType,omitempty"`
	CanOutCards      *bool                  `protobuf:"varint,5,opt,name=canOutCards" json:"canOutCards,omitempty"`
	PlayerInfos      []*ErddzBasePlayerInfo `protobuf:"bytes,6,rep,name=playerInfos" json:"playerInfos,omitempty"`
	IsCanPass        *bool                  `protobuf:"varint,7,opt,name=isCanPass" json:"isCanPass,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzBcOverTurn) Reset()                    { *m = ErddzBcOverTurn{} }
func (m *ErddzBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcOverTurn) ProtoMessage()               {}
func (*ErddzBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{15} }

func (m *ErddzBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcOverTurn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzBcOverTurn) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *ErddzBcOverTurn) GetActType() ErddzEnumActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return ErddzEnumActType_ERDDZ_T_NORMAL_ACT
}

func (m *ErddzBcOverTurn) GetCanOutCards() bool {
	if m != nil && m.CanOutCards != nil {
		return *m.CanOutCards
	}
	return false
}

func (m *ErddzBcOverTurn) GetPlayerInfos() []*ErddzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

func (m *ErddzBcOverTurn) GetIsCanPass() bool {
	if m != nil && m.IsCanPass != nil {
		return *m.IsCanPass
	}
	return false
}

// 游戏信息(广播)
type ErddzBcGameInfo struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerInfo       []*ErddzBasePlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	ErddzDeskInfo    *ErddzBaseDeskInfo     `protobuf:"bytes,3,opt,name=erddzDeskInfo" json:"erddzDeskInfo,omitempty"`
	SenderUserId     *uint32                `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *int32                 `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzBcGameInfo) Reset()                    { *m = ErddzBcGameInfo{} }
func (m *ErddzBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcGameInfo) ProtoMessage()               {}
func (*ErddzBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{16} }

func (m *ErddzBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcGameInfo) GetPlayerInfo() []*ErddzBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *ErddzBcGameInfo) GetErddzDeskInfo() *ErddzBaseDeskInfo {
	if m != nil {
		return m.ErddzDeskInfo
	}
	return nil
}

func (m *ErddzBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *ErddzBcGameInfo) GetIsReconnect() int32 {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return 0
}

func init() {
	proto.RegisterType((*ErddzBcOpening)(nil), "ddproto.erddz_bc_opening")
	proto.RegisterType((*ErddzBcDealCards)(nil), "ddproto.erddz_bc_dealCards")
	proto.RegisterType((*ErddzReqJiaoDiZhu)(nil), "ddproto.erddz_req_jiaoDiZhu")
	proto.RegisterType((*ErddzAckJiaoDiZhu)(nil), "ddproto.erddz_ack_jiaoDiZhu")
	proto.RegisterType((*ErddzReqRobDiZhu)(nil), "ddproto.erddz_req_robDiZhu")
	proto.RegisterType((*ErddzAckRobDiZhu)(nil), "ddproto.erddz_ack_robDiZhu")
	proto.RegisterType((*ErddzReqRangcards)(nil), "ddproto.erddz_req_rangcards")
	proto.RegisterType((*ErddzAckRangcards)(nil), "ddproto.erddz_ack_rangcards")
	proto.RegisterType((*ErddzReqDouble)(nil), "ddproto.erddz_req_double")
	proto.RegisterType((*ErddzAckDouble)(nil), "ddproto.erddz_ack_double")
	proto.RegisterType((*ErddzBcStartPlay)(nil), "ddproto.erddz_bc_startPlay")
	proto.RegisterType((*ErddzReqOutCards)(nil), "ddproto.erddz_req_outCards")
	proto.RegisterType((*ErddzAckOutCards)(nil), "ddproto.erddz_ack_outCards")
	proto.RegisterType((*ErddzReqActGuo)(nil), "ddproto.erddz_req_actGuo")
	proto.RegisterType((*ErddzAckActGuo)(nil), "ddproto.erddz_ack_actGuo")
	proto.RegisterType((*ErddzBcOverTurn)(nil), "ddproto.erddz_bc_overTurn")
	proto.RegisterType((*ErddzBcGameInfo)(nil), "ddproto.erddz_bc_gameInfo")
}

func init() { proto.RegisterFile("erddz_play.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0xb3, 0x4d, 0x13, 0x9c, 0x16, 0x8a, 0x5b, 0xa1, 0x55, 0xa8, 0x50, 0xb4, 0xa7, 0x1c,
	0x50, 0x0f, 0x95, 0x00, 0x89, 0x03, 0x42, 0x6d, 0x25, 0xda, 0x0b, 0x44, 0x56, 0x91, 0x10, 0x97,
	0xc8, 0x59, 0xbb, 0xcd, 0x92, 0xac, 0x1d, 0x6c, 0x2f, 0xa8, 0xfd, 0x05, 0xfe, 0x83, 0x33, 0x5f,
	0xc2, 0xaf, 0x20, 0xf1, 0x05, 0x68, 0xbc, 0xde, 0x5d, 0xa7, 0xa2, 0xa2, 0x29, 0x21, 0x97, 0x95,
	0x67, 0xfc, 0x3c, 0xcf, 0x6f, 0xc6, 0x9e, 0x35, 0xda, 0xe2, 0x8a, 0xb1, 0xcb, 0xe1, 0x6c, 0x4a,
	0x2f, 0xf6, 0x66, 0x4a, 0x1a, 0x89, 0x5b, 0x8c, 0xd9, 0x41, 0x77, 0x3b, 0x91, 0x59, 0x26, 0xc5,
	0x30, 0x99, 0xa6, 0x5c, 0x98, 0x62, 0xb6, 0xeb, 0xf0, 0x23, 0xaa, 0x79, 0xe1, 0x89, 0x7f, 0x05,
	0x65, 0x90, 0x51, 0x32, 0x94, 0x33, 0x2e, 0x52, 0x71, 0x8e, 0x9f, 0xa0, 0xf5, 0x31, 0xa7, 0x8c,
	0xab, 0x28, 0xe8, 0x05, 0xfd, 0xce, 0xfe, 0xce, 0x9e, 0x8b, 0xba, 0x37, 0x80, 0xef, 0xb1, 0x9d,
	0x23, 0x0e, 0x83, 0x1f, 0xa2, 0xf5, 0x5c, 0x73, 0x75, 0xc2, 0xa2, 0x46, 0x2f, 0xe8, 0x6f, 0x12,
	0x67, 0x81, 0xdf, 0xa4, 0xc9, 0x84, 0x9b, 0x28, 0xec, 0x05, 0xfd, 0x90, 0x38, 0x0b, 0x77, 0x51,
	0x1b, 0x10, 0x87, 0x32, 0x15, 0xd1, 0x9a, 0x9d, 0xa9, 0x6c, 0x1c, 0xa3, 0x0d, 0x18, 0x13, 0x29,
	0xb3, 0x43, 0xaa, 0x58, 0xd4, 0xb4, 0xf3, 0x73, 0x3e, 0xfc, 0x0a, 0x75, 0x40, 0x30, 0x57, 0x27,
	0xe2, 0x4c, 0xea, 0x68, 0xbd, 0x17, 0xf6, 0x3b, 0xfb, 0x8f, 0xab, 0x2d, 0xd6, 0x12, 0x87, 0x35,
	0x8c, 0xf8, 0x4b, 0xe2, 0x1f, 0x01, 0xc2, 0x95, 0x68, 0xc6, 0xe9, 0x14, 0xe2, 0xea, 0x05, 0x65,
	0xbf, 0x44, 0x1b, 0x45, 0xcc, 0x81, 0x9c, 0x70, 0xa5, 0xa3, 0x86, 0xdd, 0x47, 0xb7, 0x5a, 0x53,
	0x24, 0xde, 0x6d, 0x04, 0x20, 0x64, 0x0e, 0x8f, 0xfb, 0xe8, 0xbe, 0x1e, 0xcb, 0x2f, 0xd6, 0x7a,
	0x57, 0xe4, 0x2f, 0xb4, 0xf9, 0xbb, 0xea, 0xc6, 0x3d, 0xd4, 0xa9, 0x5c, 0x27, 0xcc, 0xe6, 0xac,
	0x49, 0x7c, 0x57, 0x2c, 0xd1, 0x76, 0xa1, 0x47, 0xf1, 0x4f, 0xc3, 0x8f, 0x29, 0x95, 0x47, 0xe9,
	0x87, 0x71, 0xbe, 0xa4, 0x3a, 0x62, 0xb4, 0x06, 0x21, 0xed, 0xee, 0xda, 0xc4, 0x8e, 0x6b, 0x42,
	0x9a, 0x4c, 0x56, 0x42, 0x38, 0x2d, 0x2b, 0x06, 0x0a, 0x95, 0x1c, 0x2d, 0x93, 0x6f, 0x0b, 0x85,
	0x4a, 0x8e, 0x1c, 0x1d, 0x0c, 0x6b, 0x36, 0x90, 0xf7, 0xdf, 0xd9, 0x72, 0xbf, 0x7a, 0x8a, 0x8a,
	0xf3, 0xe4, 0x16, 0xc7, 0xf1, 0x3a, 0xba, 0x08, 0xb5, 0x20, 0xe4, 0x9b, 0x3c, 0x73, 0xc7, 0xab,
	0x34, 0x6b, 0x5a, 0x2b, 0x72, 0x65, 0xb4, 0xb3, 0xb2, 0xe1, 0x80, 0x5a, 0x26, 0xf3, 0xd1, 0x94,
	0x2f, 0xaf, 0xe1, 0x14, 0xf1, 0x5c, 0x72, 0x9d, 0x55, 0x33, 0x82, 0xd0, 0x95, 0x30, 0x7e, 0xf7,
	0x1b, 0x8c, 0x36, 0x54, 0x99, 0xc1, 0x94, 0x5e, 0x2c, 0x48, 0xfa, 0x02, 0xa1, 0x33, 0x29, 0xcd,
	0x8d, 0xdb, 0x8b, 0x87, 0x86, 0x1e, 0x0b, 0x16, 0xa1, 0xa6, 0xd8, 0x5a, 0x93, 0x54, 0x36, 0xde,
	0x41, 0x4d, 0x96, 0x5e, 0x8e, 0x73, 0xdb, 0x48, 0x36, 0x49, 0x61, 0xc4, 0x97, 0xfe, 0x05, 0x93,
	0xb9, 0xb9, 0x4d, 0x4b, 0x7c, 0x86, 0xda, 0xe5, 0xca, 0x1b, 0xec, 0xb7, 0xc2, 0xc6, 0x3f, 0x03,
	0xff, 0xbe, 0xdd, 0x92, 0xfc, 0xba, 0x1a, 0x3d, 0x47, 0x6d, 0x38, 0xd8, 0xa7, 0x17, 0xb3, 0x22,
	0x15, 0xf7, 0xf6, 0x1f, 0x5d, 0xf9, 0x57, 0x70, 0x91, 0x67, 0xc3, 0x19, 0x4d, 0x01, 0x42, 0x2a,
	0x30, 0xfc, 0x8b, 0x14, 0xcf, 0x68, 0x2a, 0x5c, 0x05, 0x8a, 0xbe, 0x3b, 0xe7, 0x9b, 0x53, 0xdc,
	0x5c, 0x40, 0xf1, 0x7b, 0xff, 0x12, 0xd0, 0xc4, 0xbc, 0xce, 0xe5, 0x72, 0xe4, 0xd6, 0x91, 0x21,
	0x95, 0x4b, 0x8d, 0xfc, 0xad, 0x81, 0x1e, 0xd4, 0x4f, 0x85, 0xcf, 0x5c, 0x9d, 0xe6, 0x4a, 0x2c,
	0xaf, 0xe5, 0x9b, 0x34, 0x2b, 0xcf, 0xaa, 0x1d, 0xe3, 0xa7, 0xa8, 0x45, 0x13, 0x63, 0xeb, 0xb6,
	0x76, 0x7d, 0xdd, 0x1c, 0x84, 0x94, 0x58, 0xf8, 0x5b, 0x26, 0x54, 0xbc, 0xad, 0xab, 0x02, 0x17,
	0xd3, 0x77, 0xfd, 0xfb, 0x03, 0x02, 0xef, 0xa2, 0xbb, 0xa9, 0x3e, 0xa4, 0x62, 0x40, 0xb5, 0x8e,
	0x5a, 0x96, 0xa1, 0x76, 0xc4, 0x5f, 0xfd, 0x44, 0x9d, 0xd3, 0x8c, 0xc3, 0xa2, 0x85, 0x5f, 0x17,
	0xa8, 0x26, 0x74, 0x97, 0xe9, 0x6f, 0x5b, 0xf4, 0x56, 0xe0, 0x03, 0xb4, 0x69, 0x41, 0x47, 0x5c,
	0x4f, 0x6c, 0x88, 0xd0, 0x92, 0xee, 0xfe, 0x29, 0x04, 0x73, 0x18, 0x32, 0xbf, 0x04, 0x2e, 0x80,
	0xe6, 0x82, 0x55, 0xcf, 0x93, 0xa2, 0x5f, 0xcc, 0xf9, 0x20, 0xdb, 0xa9, 0x26, 0x3c, 0x91, 0x42,
	0xf0, 0xc4, 0xd8, 0x6c, 0x37, 0x89, 0xef, 0x3a, 0x68, 0x1c, 0x87, 0x83, 0x3b, 0x83, 0xe0, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x1c, 0x91, 0xea, 0xac, 0x0a, 0x00, 0x00,
}
