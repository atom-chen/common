// Code generated by protoc-gen-go.
// source: hall_playback_pdk.proto
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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pdk_base_roomTypeInfo from common_pdk.proto

// Ignoring public import of pdk_base_timerInfo from common_pdk.proto

// Ignoring public import of pdk_enum_paiType from common_pdk.proto

// Ignoring public import of pdk_enum_playerStatus from common_pdk.proto

// Ignoring public import of pdk_enum_roomType from common_pdk.proto

// Ignoring public import of pdk_enum_enterType from common_pdk.proto

// Ignoring public import of pdk_enum_coinRoomLevel from common_pdk.proto

// Ignoring public import of pdk_enum_deskGameStatus from common_pdk.proto

// 跑得快的操作类型
type PlaybackPdkActType int32

const (
	PlaybackPdkActType_PPAT_GUO   PlaybackPdkActType = 1
	PlaybackPdkActType_PPAT_DAPAI PlaybackPdkActType = 2
)

var PlaybackPdkActType_name = map[int32]string{
	1: "PPAT_GUO",
	2: "PPAT_DAPAI",
}
var PlaybackPdkActType_value = map[string]int32{
	"PPAT_GUO":   1,
	"PPAT_DAPAI": 2,
}

func (x PlaybackPdkActType) Enum() *PlaybackPdkActType {
	p := new(PlaybackPdkActType)
	*p = x
	return p
}
func (x PlaybackPdkActType) String() string {
	return proto.EnumName(PlaybackPdkActType_name, int32(x))
}
func (x *PlaybackPdkActType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlaybackPdkActType_value, data, "PlaybackPdkActType")
	if err != nil {
		return err
	}
	*x = PlaybackPdkActType(value)
	return nil
}
func (PlaybackPdkActType) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

type PlaybackPdkAckPage struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 分页信息
	Length *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Page   *int32 `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Total  *int32 `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	// 快照信息
	PlaybackSnapshots []*PdkPlaybackSnapshot `protobuf:"bytes,5,rep,name=playbackSnapshots" json:"playbackSnapshots,omitempty"`
	XXX_unrecognized  []byte                 `json:"-"`
}

func (m *PlaybackPdkAckPage) Reset()                    { *m = PlaybackPdkAckPage{} }
func (m *PlaybackPdkAckPage) String() string            { return proto.CompactTextString(m) }
func (*PlaybackPdkAckPage) ProtoMessage()               {}
func (*PlaybackPdkAckPage) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *PlaybackPdkAckPage) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PlaybackPdkAckPage) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPage() int32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetTotal() int32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *PlaybackPdkAckPage) GetPlaybackSnapshots() []*PdkPlaybackSnapshot {
	if m != nil {
		return m.PlaybackSnapshots
	}
	return nil
}

// 玩家的信息
type PdkPlayerInfo struct {
	PlayerPokers     []*ClientBasePoker   `protobuf:"bytes,1,rep,name=playerPokers" json:"playerPokers,omitempty"`
	Coin             *int64               `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	NickName         *string              `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	Sex              *int32               `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	UserId           *uint32              `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool                `protobuf:"varint,6,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32               `protobuf:"varint,7,opt,name=bReady" json:"bReady,omitempty"`
	Status           *PdkEnumPlayerStatus `protobuf:"varint,8,opt,name=status,enum=ddproto.PdkEnumPlayerStatus" json:"status,omitempty"`
	WxInfo           *WeixinInfo          `protobuf:"bytes,9,opt,name=wxInfo" json:"wxInfo,omitempty"`
	OnlineStatus     *int32               `protobuf:"varint,10,opt,name=onlineStatus" json:"onlineStatus,omitempty"`
	RemainPaiCount   *int32               `protobuf:"varint,11,opt,name=remainPaiCount" json:"remainPaiCount,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,12,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,13,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	RoomCard         *int64               `protobuf:"varint,14,opt,name=roomCard" json:"roomCard,omitempty"`
	IsBanker         *bool                `protobuf:"varint,15,opt,name=isBanker" json:"isBanker,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlayerInfo) Reset()                    { *m = PdkPlayerInfo{} }
func (m *PdkPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlayerInfo) ProtoMessage()               {}
func (*PdkPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *PdkPlayerInfo) GetPlayerPokers() []*ClientBasePoker {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

func (m *PdkPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkPlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PdkPlayerInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PdkPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkPlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PdkPlayerInfo) GetStatus() PdkEnumPlayerStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PdkEnumPlayerStatus_PDK_NORMAL_DEFAULT
}

func (m *PdkPlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PdkPlayerInfo) GetOnlineStatus() int32 {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return 0
}

func (m *PdkPlayerInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

func (m *PdkPlayerInfo) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlayerInfo) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

func (m *PdkPlayerInfo) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *PdkPlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

// 回放快照信息（包含玩家信息、桌面信息、房间信息、右上角时间信息等），另外包含协议信息（PID 和 二进制数据）
type PdkPlaybackSnapshot struct {
	PlayerInfo       []*PdkPlayerInfo     `protobuf:"bytes,1,rep,name=PlayerInfo" json:"PlayerInfo,omitempty"`
	PlayBackDeskInfo *PdkPlaybackDeskInfo `protobuf:"bytes,2,opt,name=PlayBackDeskInfo" json:"PlayBackDeskInfo,omitempty"`
	RoomTypeInfo     *PdkBaseRoomTypeInfo `protobuf:"bytes,3,opt,name=RoomTypeInfo" json:"RoomTypeInfo,omitempty"`
	Time             *string              `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	ActType          *PlaybackPdkActType  `protobuf:"varint,5,opt,name=actType,enum=ddproto.PlaybackPdkActType" json:"actType,omitempty"`
	Userid           *uint32              `protobuf:"varint,6,opt,name=userid" json:"userid,omitempty"`
	OutPais          []*ClientBasePoker   `protobuf:"bytes,7,rep,name=outPais" json:"outPais,omitempty"`
	OutPaiType       *PdkEnumPaiType      `protobuf:"varint,8,opt,name=outPaiType,enum=ddproto.PdkEnumPaiType" json:"outPaiType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkPlaybackSnapshot) Reset()                    { *m = PdkPlaybackSnapshot{} }
func (m *PdkPlaybackSnapshot) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackSnapshot) ProtoMessage()               {}
func (*PdkPlaybackSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{2} }

func (m *PdkPlaybackSnapshot) GetPlayerInfo() []*PdkPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetPlayBackDeskInfo() *PdkPlaybackDeskInfo {
	if m != nil {
		return m.PlayBackDeskInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetRoomTypeInfo() *PdkBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *PdkPlaybackSnapshot) GetActType() PlaybackPdkActType {
	if m != nil && m.ActType != nil {
		return *m.ActType
	}
	return PlaybackPdkActType_PPAT_GUO
}

func (m *PdkPlaybackSnapshot) GetUserid() uint32 {
	if m != nil && m.Userid != nil {
		return *m.Userid
	}
	return 0
}

func (m *PdkPlaybackSnapshot) GetOutPais() []*ClientBasePoker {
	if m != nil {
		return m.OutPais
	}
	return nil
}

func (m *PdkPlaybackSnapshot) GetOutPaiType() PdkEnumPaiType {
	if m != nil && m.OutPaiType != nil {
		return *m.OutPaiType
	}
	return PdkEnumPaiType_PDK_ERRORCARD
}

// 回放桌面信息
type PdkPlaybackDeskInfo struct {
	GameStatus       *int32            `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	PlayerNum        *int32            `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveUserId     *uint32           `protobuf:"varint,4,opt,name=activeUserId" json:"activeUserId,omitempty"`
	CurrPlayCount    *int32            `protobuf:"varint,10,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32            `protobuf:"varint,11,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	RoomNumber       *string           `protobuf:"bytes,12,opt,name=roomNumber" json:"roomNumber,omitempty"`
	BankerId         *uint32           `protobuf:"varint,14,opt,name=bankerId" json:"bankerId,omitempty"`
	EnterType        *PdkEnumEnterType `protobuf:"varint,18,opt,name=enterType,enum=ddproto.PdkEnumEnterType" json:"enterType,omitempty"`
	CoinTicket       *int64            `protobuf:"varint,19,opt,name=coinTicket" json:"coinTicket,omitempty"`
	TimerInfo        *PdkBaseTimerInfo `protobuf:"bytes,20,opt,name=timerInfo" json:"timerInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkPlaybackDeskInfo) Reset()                    { *m = PdkPlaybackDeskInfo{} }
func (m *PdkPlaybackDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkPlaybackDeskInfo) ProtoMessage()               {}
func (*PdkPlaybackDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{3} }

func (m *PdkPlaybackDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *PdkPlaybackDeskInfo) GetBankerId() uint32 {
	if m != nil && m.BankerId != nil {
		return *m.BankerId
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetEnterType() PdkEnumEnterType {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return PdkEnumEnterType_PDK_FRIEND
}

func (m *PdkPlaybackDeskInfo) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *PdkPlaybackDeskInfo) GetTimerInfo() *PdkBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

type PdkDataRecoverPlayerInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool   `protobuf:"varint,2,opt,name=isOwner" json:"isOwner,omitempty"`
	IsReady          *bool   `protobuf:"varint,3,opt,name=isReady" json:"isReady,omitempty"`
	HandPokers       []int32 `protobuf:"varint,4,rep,name=HandPokers" json:"HandPokers,omitempty"`
	UserGameStatus   *int32  `protobuf:"varint,5,opt,name=UserGameStatus" json:"UserGameStatus,omitempty"`
	DissolveStatus   *int32  `protobuf:"varint,6,opt,name=DissolveStatus" json:"DissolveStatus,omitempty"`
	OutPaiList       []int32 `protobuf:"varint,7,rep,name=OutPaiList" json:"OutPaiList,omitempty"`
	RemainPaiCount   *int32  `protobuf:"varint,8,opt,name=RemainPaiCount" json:"RemainPaiCount,omitempty"`
	OnlineStatus     *bool   `protobuf:"varint,9,opt,name=OnlineStatus" json:"OnlineStatus,omitempty"`
	IsBigWinner      *bool   `protobuf:"varint,10,opt,name=isBigWinner" json:"isBigWinner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,11,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64  `protobuf:"varint,12,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkDataRecoverPlayerInfo) Reset()                    { *m = PdkDataRecoverPlayerInfo{} }
func (m *PdkDataRecoverPlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverPlayerInfo) ProtoMessage()               {}
func (*PdkDataRecoverPlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{4} }

func (m *PdkDataRecoverPlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetHandPokers() []int32 {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PdkDataRecoverPlayerInfo) GetUserGameStatus() int32 {
	if m != nil && m.UserGameStatus != nil {
		return *m.UserGameStatus
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetDissolveStatus() int32 {
	if m != nil && m.DissolveStatus != nil {
		return *m.DissolveStatus
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetOutPaiList() []int32 {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *PdkDataRecoverPlayerInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetOnlineStatus() bool {
	if m != nil && m.OnlineStatus != nil {
		return *m.OnlineStatus
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetIsBigWinner() bool {
	if m != nil && m.IsBigWinner != nil {
		return *m.IsBigWinner
	}
	return false
}

func (m *PdkDataRecoverPlayerInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkDataRecoverPlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

type PdkDataRecoverDeskInfo struct {
	DeskGameStatus   *int32            `protobuf:"varint,1,opt,name=DeskGameStatus" json:"DeskGameStatus,omitempty"`
	PlayerNum        *int32            `protobuf:"varint,2,opt,name=PlayerNum" json:"PlayerNum,omitempty"`
	ActiveUserId     *uint32           `protobuf:"varint,3,opt,name=ActiveUserId" json:"ActiveUserId,omitempty"`
	IsFollowPai      *bool             `protobuf:"varint,4,opt,name=isFollowPai" json:"isFollowPai,omitempty"`
	CurrentRound     *int32            `protobuf:"varint,5,opt,name=CurrentRound" json:"CurrentRound,omitempty"`
	TotalRound       *int32            `protobuf:"varint,6,opt,name=TotalRound" json:"TotalRound,omitempty"`
	TimerInfo        *PdkBaseTimerInfo `protobuf:"bytes,7,opt,name=timerInfo" json:"timerInfo,omitempty"`
	LatestOutPokers  []int32           `protobuf:"varint,8,rep,name=LatestOutPokers" json:"LatestOutPokers,omitempty"`
	RoomNumber       *string           `protobuf:"bytes,9,opt,name=RoomNumber" json:"RoomNumber,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkDataRecoverDeskInfo) Reset()                    { *m = PdkDataRecoverDeskInfo{} }
func (m *PdkDataRecoverDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverDeskInfo) ProtoMessage()               {}
func (*PdkDataRecoverDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{5} }

func (m *PdkDataRecoverDeskInfo) GetDeskGameStatus() int32 {
	if m != nil && m.DeskGameStatus != nil {
		return *m.DeskGameStatus
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetIsFollowPai() bool {
	if m != nil && m.IsFollowPai != nil {
		return *m.IsFollowPai
	}
	return false
}

func (m *PdkDataRecoverDeskInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetTotalRound() int32 {
	if m != nil && m.TotalRound != nil {
		return *m.TotalRound
	}
	return 0
}

func (m *PdkDataRecoverDeskInfo) GetTimerInfo() *PdkBaseTimerInfo {
	if m != nil {
		return m.TimerInfo
	}
	return nil
}

func (m *PdkDataRecoverDeskInfo) GetLatestOutPokers() []int32 {
	if m != nil {
		return m.LatestOutPokers
	}
	return nil
}

func (m *PdkDataRecoverDeskInfo) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

type PdkDataRecover struct {
	PlayerInfo       []*PdkDataRecoverPlayerInfo `protobuf:"bytes,1,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *PdkDataRecoverDeskInfo     `protobuf:"bytes,2,opt,name=deskInfo" json:"deskInfo,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *PdkDataRecover) Reset()                    { *m = PdkDataRecover{} }
func (m *PdkDataRecover) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecover) ProtoMessage()               {}
func (*PdkDataRecover) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{6} }

func (m *PdkDataRecover) GetPlayerInfo() []*PdkDataRecoverPlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *PdkDataRecover) GetDeskInfo() *PdkDataRecoverDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

type PdkDataRecoverDeskIdList struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=DeskId" json:"DeskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkDataRecoverDeskIdList) Reset()                    { *m = PdkDataRecoverDeskIdList{} }
func (m *PdkDataRecoverDeskIdList) String() string            { return proto.CompactTextString(m) }
func (*PdkDataRecoverDeskIdList) ProtoMessage()               {}
func (*PdkDataRecoverDeskIdList) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{7} }

func (m *PdkDataRecoverDeskIdList) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

func init() {
	proto.RegisterType((*PlaybackPdkAckPage)(nil), "ddproto.playback_pdk_ack_page")
	proto.RegisterType((*PdkPlayerInfo)(nil), "ddproto.PdkPlayerInfo")
	proto.RegisterType((*PdkPlaybackSnapshot)(nil), "ddproto.PdkPlaybackSnapshot")
	proto.RegisterType((*PdkPlaybackDeskInfo)(nil), "ddproto.PdkPlaybackDeskInfo")
	proto.RegisterType((*PdkDataRecoverPlayerInfo)(nil), "ddproto.PdkDataRecoverPlayerInfo")
	proto.RegisterType((*PdkDataRecoverDeskInfo)(nil), "ddproto.PdkDataRecoverDeskInfo")
	proto.RegisterType((*PdkDataRecover)(nil), "ddproto.PdkDataRecover")
	proto.RegisterType((*PdkDataRecoverDeskIdList)(nil), "ddproto.PdkDataRecoverDeskIdList")
	proto.RegisterEnum("ddproto.PlaybackPdkActType", PlaybackPdkActType_name, PlaybackPdkActType_value)
}

var fileDescriptor30 = []byte{
	// 1077 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0x1c, 0x45,
	0x10, 0x65, 0xbc, 0xde, 0x5b, 0x79, 0xed, 0x98, 0x76, 0x30, 0x8d, 0x13, 0x85, 0x65, 0x85, 0xd0,
	0x0a, 0x50, 0x1e, 0x2c, 0x88, 0x88, 0x90, 0x90, 0x7c, 0x11, 0xb1, 0x51, 0x64, 0x8f, 0x3a, 0x8e,
	0xf2, 0xb8, 0x6a, 0xcf, 0x34, 0x76, 0x6b, 0x76, 0xba, 0x57, 0x33, 0xbd, 0xbe, 0xfc, 0x05, 0xe2,
	0x85, 0xff, 0x40, 0xe2, 0x4b, 0xf8, 0x0e, 0xfe, 0x01, 0x55, 0xf5, 0xec, 0x5c, 0x76, 0x1d, 0xa1,
	0xf0, 0x62, 0x77, 0x9d, 0x3e, 0x5d, 0xd3, 0x55, 0x75, 0xaa, 0x7a, 0xe1, 0xd3, 0x6b, 0x39, 0x9d,
	0x4e, 0x66, 0x53, 0x79, 0x7f, 0x29, 0xa3, 0x64, 0x32, 0x8b, 0x93, 0xe7, 0xb3, 0xcc, 0x3a, 0xcb,
	0xba, 0x71, 0x4c, 0x8b, 0xbd, 0x9d, 0xc8, 0xa6, 0xa9, 0x35, 0x93, 0x68, 0xaa, 0x95, 0x71, 0x7e,
	0x77, 0x6f, 0xbb, 0x00, 0x4b, 0xfe, 0xe8, 0xef, 0x00, 0x3e, 0xa9, 0xbb, 0x99, 0xd0, 0x7f, 0x79,
	0xa5, 0xd8, 0xb7, 0xd0, 0xb9, 0x56, 0x32, 0x56, 0x19, 0x0f, 0x86, 0xc1, 0x78, 0x63, 0xff, 0xf1,
	0xf3, 0xc2, 0xf5, 0xf3, 0x10, 0xff, 0x9e, 0xd0, 0x9e, 0x28, 0x38, 0x6c, 0x17, 0x3a, 0x53, 0x65,
	0xae, 0xdc, 0x35, 0x5f, 0x1b, 0x06, 0xe3, 0xb6, 0x28, 0x2c, 0xc6, 0x60, 0x1d, 0xbd, 0xf1, 0x16,
	0xa1, 0xb4, 0x66, 0x8f, 0xa1, 0xed, 0xac, 0x93, 0x53, 0xbe, 0x4e, 0xa0, 0x37, 0xd8, 0x2f, 0xf0,
	0xf1, 0xe2, 0x22, 0x6f, 0x8c, 0x9c, 0xe5, 0xd7, 0xd6, 0xe5, 0xbc, 0x3d, 0x6c, 0x8d, 0x37, 0xf6,
	0x9f, 0x56, 0x9f, 0x8e, 0x93, 0x70, 0x89, 0x24, 0x56, 0x8f, 0x8d, 0xfe, 0x5a, 0x87, 0xcd, 0x82,
	0xaa, 0xb2, 0x53, 0xf3, 0xab, 0x65, 0x3f, 0xc1, 0x60, 0x46, 0x56, 0x68, 0x13, 0x95, 0xe5, 0x3c,
	0x20, 0xc7, 0x7b, 0xa5, 0x63, 0x9f, 0xa6, 0xc9, 0xa5, 0xcc, 0xd5, 0x64, 0x86, 0x14, 0xd1, 0xe0,
	0x63, 0x1c, 0x91, 0xd5, 0x86, 0xa2, 0x6b, 0x09, 0x5a, 0xb3, 0x3d, 0xe8, 0x19, 0x1d, 0x25, 0x67,
	0x32, 0xf5, 0xf1, 0xf5, 0x45, 0x69, 0xb3, 0x6d, 0x68, 0xe5, 0xea, 0xae, 0x88, 0x10, 0x97, 0x98,
	0xa1, 0x79, 0xae, 0xb2, 0xd3, 0x98, 0xb7, 0x87, 0xc1, 0x78, 0x53, 0x14, 0x16, 0xe3, 0xd0, 0xd5,
	0xf9, 0xf9, 0xad, 0x51, 0x19, 0xef, 0x0c, 0x83, 0x71, 0x4f, 0x2c, 0x4c, 0x3c, 0x71, 0x29, 0x94,
	0x8c, 0xef, 0x79, 0xd7, 0xe7, 0xd4, 0x5b, 0xec, 0x05, 0x74, 0x72, 0x27, 0xdd, 0x3c, 0xe7, 0xbd,
	0x61, 0x30, 0xde, 0xda, 0x7f, 0x56, 0x46, 0x81, 0x05, 0x54, 0x66, 0x9e, 0x4e, 0xfc, 0xdd, 0xdf,
	0x10, 0x4b, 0x14, 0x6c, 0xf6, 0x0d, 0x74, 0x6e, 0xef, 0x30, 0x1b, 0xbc, 0x4f, 0x15, 0xdd, 0x29,
	0xcf, 0xbd, 0x53, 0xfa, 0x4e, 0x1b, 0xdc, 0x12, 0x05, 0x85, 0x8d, 0x60, 0x60, 0xcd, 0x54, 0x1b,
	0xe5, 0x9d, 0x70, 0xa0, 0x2b, 0x34, 0x30, 0xf6, 0x15, 0x6c, 0x65, 0x2a, 0x95, 0xda, 0x84, 0x52,
	0x1f, 0xd9, 0xb9, 0x71, 0x7c, 0x83, 0x58, 0x4b, 0x28, 0xfb, 0x0e, 0xba, 0x76, 0xee, 0x42, 0xa9,
	0x73, 0x3e, 0xf8, 0xcf, 0xbc, 0x2f, 0xa8, 0xec, 0x25, 0x80, 0x5f, 0x5e, 0xdc, 0xcf, 0x14, 0xdf,
	0xa4, 0x50, 0x3f, 0x7b, 0x20, 0x54, 0x4f, 0x10, 0x35, 0x32, 0x56, 0x26, 0xb3, 0x36, 0x3d, 0x92,
	0x59, 0xcc, 0xb7, 0xa8, 0x62, 0xa5, 0x8d, 0x7b, 0x3a, 0x3f, 0x94, 0x26, 0x51, 0x19, 0x7f, 0x44,
	0x09, 0x2f, 0xed, 0xd1, 0x9f, 0x2d, 0xd8, 0x79, 0x40, 0x62, 0xec, 0x05, 0x40, 0xa5, 0xa5, 0x42,
	0x3b, 0xbb, 0xcb, 0xa2, 0xf4, 0xbb, 0xa2, 0xc6, 0x64, 0x27, 0xb0, 0x8d, 0xd6, 0xa1, 0x8c, 0x92,
	0x63, 0x95, 0x27, 0x74, 0x7a, 0x8d, 0x72, 0xff, 0xa0, 0xa4, 0x17, 0x1c, 0xb1, 0x72, 0x8a, 0x1d,
	0xc2, 0x40, 0x58, 0x9b, 0x62, 0x74, 0xe4, 0xa5, 0x45, 0x5e, 0x9a, 0x95, 0xa7, 0x24, 0x66, 0x35,
	0x96, 0x68, 0x9c, 0x41, 0x0d, 0x3b, 0x9d, 0x2a, 0x12, 0x65, 0x5f, 0xd0, 0x9a, 0xfd, 0x00, 0x5d,
	0x19, 0x39, 0xca, 0x70, 0x7b, 0x59, 0x4c, 0xcd, 0xb1, 0xe0, 0x26, 0x0e, 0xd3, 0xbc, 0xa0, 0x2f,
	0xf4, 0xac, 0x63, 0x92, 0x6d, 0xa1, 0x67, 0x1d, 0xd7, 0x8b, 0xdd, 0xfd, 0xbf, 0xc5, 0xee, 0x7d,
	0x40, 0xb1, 0x47, 0x7f, 0x34, 0x8b, 0x56, 0xa6, 0xec, 0x19, 0xc0, 0x2b, 0x99, 0x2e, 0xf4, 0x1b,
	0x90, 0x32, 0x6b, 0x08, 0x7b, 0x0a, 0x7d, 0xdf, 0x26, 0x67, 0xf3, 0xb4, 0x98, 0x4f, 0x15, 0x80,
	0xfa, 0x97, 0x91, 0xd3, 0x37, 0xea, 0xad, 0x6f, 0xda, 0x75, 0x0a, 0xb2, 0x81, 0xb1, 0x2f, 0x61,
	0x33, 0x9a, 0x67, 0x19, 0x7e, 0xd9, 0xcb, 0xdf, 0x37, 0x49, 0x13, 0xc4, 0x2e, 0xa1, 0x09, 0x57,
	0xd1, 0x8a, 0x2e, 0x69, 0xa2, 0x78, 0x5f, 0x2c, 0xde, 0xd9, 0x3c, 0xbd, 0x54, 0x19, 0x1f, 0x50,
	0x91, 0x6a, 0x08, 0x0a, 0xf7, 0x92, 0x64, 0x7a, 0xea, 0x45, 0xbd, 0x29, 0x4a, 0x9b, 0xbd, 0x84,
	0xbe, 0x32, 0x4e, 0x65, 0x94, 0x3d, 0x46, 0xd9, 0x7b, 0xb2, 0x9a, 0xbd, 0x92, 0x22, 0x2a, 0x36,
	0x7e, 0x16, 0xa7, 0xd9, 0x85, 0x8e, 0x12, 0xe5, 0xf8, 0x0e, 0x75, 0x4b, 0x0d, 0x41, 0xd7, 0xa8,
	0x14, 0x2f, 0xfd, 0xc7, 0x24, 0xbb, 0x27, 0xab, 0xb2, 0x2b, 0x29, 0xa2, 0x62, 0x8f, 0x7e, 0x6f,
	0x01, 0x0f, 0xe3, 0xe4, 0x58, 0x3a, 0x29, 0x54, 0x64, 0x6f, 0x54, 0x56, 0xeb, 0x8d, 0x6a, 0x1e,
	0x06, 0xef, 0x9b, 0x87, 0x6b, 0xcd, 0x79, 0x48, 0x3b, 0x7e, 0x20, 0xb6, 0x16, 0x3b, 0x7e, 0x22,
	0x3e, 0x03, 0x38, 0x91, 0x26, 0x2e, 0x66, 0xfb, 0xfa, 0xb0, 0x85, 0xa5, 0xae, 0x10, 0x2c, 0x01,
	0x96, 0xac, 0x26, 0x87, 0xb6, 0x2f, 0x41, 0x13, 0x45, 0xde, 0xb1, 0xce, 0x73, 0x3b, 0xbd, 0x59,
	0xf0, 0x3a, 0x9e, 0xd7, 0x44, 0xf1, 0x7b, 0xe7, 0x24, 0xc0, 0xd7, 0x3a, 0x77, 0x24, 0xf3, 0xb6,
	0xa8, 0x21, 0xe8, 0x47, 0x34, 0x07, 0x63, 0xcf, 0xfb, 0x69, 0xa2, 0x28, 0xb2, 0xf3, 0xfa, 0x90,
	0xed, 0x53, 0x58, 0x0d, 0x8c, 0x0d, 0x61, 0x43, 0xe7, 0x87, 0xfa, 0xea, 0x9d, 0x36, 0x98, 0x13,
	0x20, 0x4a, 0x1d, 0xc2, 0xbc, 0xdc, 0x6a, 0x73, 0x84, 0xcf, 0xd3, 0x06, 0x95, 0x6f, 0x61, 0x96,
	0xaf, 0xd6, 0xa0, 0x7a, 0xb5, 0x46, 0xff, 0xac, 0xc1, 0x6e, 0xb3, 0x28, 0x65, 0xc7, 0x60, 0xf8,
	0x2a, 0x4f, 0x56, 0xba, 0x66, 0x09, 0xc5, 0xce, 0x09, 0xcb, 0xce, 0xf1, 0xef, 0x7d, 0x05, 0x60,
	0x50, 0x07, 0xf5, 0xce, 0x69, 0xf9, 0xce, 0xa9, 0x63, 0x3e, 0xa8, 0x9f, 0xed, 0x74, 0x6a, 0x6f,
	0x43, 0xa9, 0xa9, 0xb9, 0x28, 0xa8, 0x12, 0x42, 0x2f, 0x47, 0xf3, 0x2c, 0x53, 0xc6, 0x09, 0x3b,
	0x37, 0x71, 0x51, 0xb0, 0x06, 0x86, 0x65, 0xb8, 0xc0, 0x1e, 0xf2, 0x0c, 0x5f, 0xaa, 0x1a, 0xd2,
	0x94, 0x6e, 0xf7, 0x43, 0xa4, 0xcb, 0xc6, 0xf0, 0xe8, 0xb5, 0x74, 0x2a, 0x77, 0x58, 0x55, 0x2f,
	0xab, 0x1e, 0x95, 0x79, 0x19, 0xc6, 0x4b, 0x88, 0xaa, 0x6d, 0xfb, 0xbe, 0x6d, 0x2b, 0x64, 0xf4,
	0x5b, 0x00, 0x5b, 0xcd, 0x7c, 0xb3, 0x03, 0x80, 0xd9, 0xf2, 0x73, 0xf2, 0x45, 0xfd, 0x41, 0x78,
	0xb0, 0x63, 0x44, 0xed, 0x10, 0xfb, 0x11, 0x7a, 0x71, 0xf3, 0x45, 0xf9, 0xfc, 0x3d, 0x0e, 0xca,
	0x47, 0xa5, 0x3c, 0x30, 0xda, 0x5f, 0x6e, 0x4b, 0xe2, 0xc4, 0x24, 0xdd, 0x5d, 0xe8, 0x78, 0x8b,
	0xee, 0xd5, 0x16, 0x85, 0xf5, 0xf5, 0xf7, 0x2b, 0xbf, 0x13, 0xfd, 0x83, 0xc0, 0x06, 0xd0, 0x0b,
	0xc3, 0x83, 0x8b, 0xc9, 0xab, 0xb7, 0xe7, 0xdb, 0x01, 0xdb, 0x02, 0x20, 0xeb, 0xf8, 0x20, 0x3c,
	0x38, 0xdd, 0x5e, 0x0b, 0x3f, 0x0a, 0x83, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x41, 0xa6,
	0x19, 0xad, 0x0a, 0x00, 0x00,
}
