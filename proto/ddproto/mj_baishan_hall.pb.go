// Code generated by protoc-gen-go.
// source: mj_baishan_hall.proto
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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of CardInfo from mj_common.proto

// Ignoring public import of ErrorCode from mj_common.proto

// Ignoring public import of mj_enum_color from mj_common.proto

// Ignoring public import of mj_enum_gangType from mj_common.proto

// Ignoring public import of mj_enum_huType from mj_common.proto

// Ignoring public import of mj_enum_composeCardType from mj_common.proto

// Ignoring public import of mj_enum_paiType from mj_common.proto

// Ignoring public import of mj_enum_userGameStatus from mj_common.proto

// Ignoring public import of mj_enum_deskGameStatus from mj_common.proto

// Ignoring public import of MJRoomType from mj_common.proto

// Ignoring public import of PlayOptions from mj_baishan_base.proto

// Ignoring public import of ChangShaPlayOptions from mj_baishan_base.proto

// Ignoring public import of RoomTypeInfo from mj_baishan_base.proto

// Ignoring public import of ComposeCard from mj_baishan_base.proto

// Ignoring public import of PlayerCard from mj_baishan_base.proto

// Ignoring public import of PlayerInfo from mj_baishan_base.proto

// Ignoring public import of DeskGameInfo from mj_baishan_base.proto

// Ignoring public import of MJ_BAISHAN_PID from mj_baishan_base.proto

// Ignoring public import of MJOption from mj_baishan_base.proto

type Game_Notice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	ChannelId        *string      `protobuf:"bytes,3,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Notice) Reset()                    { *m = Game_Notice{} }
func (m *Game_Notice) String() string            { return proto.CompactTextString(m) }
func (*Game_Notice) ProtoMessage()               {}
func (*Game_Notice) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *Game_Notice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Notice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Game_Notice) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 公告的内容
type Game_AckNotice struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NoticeType       *int32       `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string      `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string      `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string      `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Id               *int32       `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	Fileds           []string     `protobuf:"bytes,7,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckNotice) Reset()                    { *m = Game_AckNotice{} }
func (m *Game_AckNotice) String() string            { return proto.CompactTextString(m) }
func (*Game_AckNotice) ProtoMessage()               {}
func (*Game_AckNotice) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *Game_AckNotice) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *Game_AckNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *Game_AckNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *Game_AckNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *Game_AckNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Game_AckNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 游戏战绩
type Game_GameRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	UserId           *uint32      `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_GameRecord) Reset()                    { *m = Game_GameRecord{} }
func (m *Game_GameRecord) String() string            { return proto.CompactTextString(m) }
func (*Game_GameRecord) ProtoMessage()               {}
func (*Game_GameRecord) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

func (m *Game_GameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_GameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Game_GameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Game_GameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

//
type BeanUserRecord struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	WinAmount        *int64       `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BeanUserRecord) Reset()                    { *m = BeanUserRecord{} }
func (m *BeanUserRecord) String() string            { return proto.CompactTextString(m) }
func (*BeanUserRecord) ProtoMessage()               {}
func (*BeanUserRecord) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{3} }

func (m *BeanUserRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BeanUserRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BeanUserRecord) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BeanUserRecord) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

//
type BeanGameRecord struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Id               *int32            `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DeskId           *int32            `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	BeginTime        *string           `protobuf:"bytes,4,opt,name=beginTime" json:"beginTime,omitempty"`
	Users            []*BeanUserRecord `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BeanGameRecord) Reset()                    { *m = BeanGameRecord{} }
func (m *BeanGameRecord) String() string            { return proto.CompactTextString(m) }
func (*BeanGameRecord) ProtoMessage()               {}
func (*BeanGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{4} }

func (m *BeanGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BeanGameRecord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *BeanGameRecord) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *BeanGameRecord) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *BeanGameRecord) GetUsers() []*BeanUserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

//
type Game_AckGameRecord struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32           `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	GameId           *int32            `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	Records          []*BeanGameRecord `protobuf:"bytes,4,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Game_AckGameRecord) Reset()                    { *m = Game_AckGameRecord{} }
func (m *Game_AckGameRecord) String() string            { return proto.CompactTextString(m) }
func (*Game_AckGameRecord) ProtoMessage()               {}
func (*Game_AckGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{5} }

func (m *Game_AckGameRecord) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckGameRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_AckGameRecord) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *Game_AckGameRecord) GetRecords() []*BeanGameRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

// 反馈信息的协议
type Game_Feedback struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Message          *string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Feedback) Reset()                    { *m = Game_Feedback{} }
func (m *Game_Feedback) String() string            { return proto.CompactTextString(m) }
func (*Game_Feedback) ProtoMessage()               {}
func (*Game_Feedback) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{6} }

func (m *Game_Feedback) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Feedback) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// 创建房间
type Game_CreateRoom struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Game_CreateRoom) Reset()                    { *m = Game_CreateRoom{} }
func (m *Game_CreateRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_CreateRoom) ProtoMessage()               {}
func (*Game_CreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{7} }

func (m *Game_CreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_CreateRoom) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type Game_AckCreateRoom struct {
	Header           *ProtoHeader  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32        `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string       `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	UserBalance      *int64        `protobuf:"varint,4,opt,name=userBalance" json:"userBalance,omitempty"`
	CreateFee        *int64        `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Game_AckCreateRoom) Reset()                    { *m = Game_AckCreateRoom{} }
func (m *Game_AckCreateRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_AckCreateRoom) ProtoMessage()               {}
func (*Game_AckCreateRoom) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{8} }

func (m *Game_AckCreateRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckCreateRoom) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Game_AckCreateRoom) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Game_AckCreateRoom) GetUserBalance() int64 {
	if m != nil && m.UserBalance != nil {
		return *m.UserBalance
	}
	return 0
}

func (m *Game_AckCreateRoom) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *Game_AckCreateRoom) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 客户端请求进入 room, 服务器返回game_SendGameInfo
type Game_EnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	MatchId          *int32       `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	TableId          *int32       `protobuf:"varint,3,opt,name=tableId" json:"tableId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32      `protobuf:"varint,5,opt,name=userId" json:"userId,omitempty"`
	RoomType         *int32       `protobuf:"varint,6,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32       `protobuf:"varint,7,opt,name=roomLevel" json:"roomLevel,omitempty"`
	EnterType        *int32       `protobuf:"varint,8,opt,name=enterType" json:"enterType,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_EnterRoom) Reset()                    { *m = Game_EnterRoom{} }
func (m *Game_EnterRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_EnterRoom) ProtoMessage()               {}
func (*Game_EnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{9} }

func (m *Game_EnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_EnterRoom) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *Game_EnterRoom) GetTableId() int32 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *Game_EnterRoom) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *Game_EnterRoom) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_EnterRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *Game_EnterRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *Game_EnterRoom) GetEnterType() int32 {
	if m != nil && m.EnterType != nil {
		return *m.EnterType
	}
	return 0
}

type Game_AckEnterRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckEnterRoom) Reset()                    { *m = Game_AckEnterRoom{} }
func (m *Game_AckEnterRoom) String() string            { return proto.CompactTextString(m) }
func (*Game_AckEnterRoom) ProtoMessage()               {}
func (*Game_AckEnterRoom) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{10} }

func (m *Game_AckEnterRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*Game_Notice)(nil), "ddproto.game_Notice")
	proto.RegisterType((*Game_AckNotice)(nil), "ddproto.game_AckNotice")
	proto.RegisterType((*Game_GameRecord)(nil), "ddproto.game_GameRecord")
	proto.RegisterType((*BeanUserRecord)(nil), "ddproto.BeanUserRecord")
	proto.RegisterType((*BeanGameRecord)(nil), "ddproto.BeanGameRecord")
	proto.RegisterType((*Game_AckGameRecord)(nil), "ddproto.game_AckGameRecord")
	proto.RegisterType((*Game_Feedback)(nil), "ddproto.game_Feedback")
	proto.RegisterType((*Game_CreateRoom)(nil), "ddproto.game_CreateRoom")
	proto.RegisterType((*Game_AckCreateRoom)(nil), "ddproto.game_AckCreateRoom")
	proto.RegisterType((*Game_EnterRoom)(nil), "ddproto.game_EnterRoom")
	proto.RegisterType((*Game_AckEnterRoom)(nil), "ddproto.game_AckEnterRoom")
}

var fileDescriptor21 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x31, 0x49, 0x9a, 0x09, 0x49, 0xa8, 0x4b, 0x84, 0xd5, 0x53, 0x65, 0x21, 0x54, 0x09,
	0x29, 0x87, 0xde, 0x38, 0x36, 0x15, 0xa5, 0x95, 0xa0, 0x8a, 0xa2, 0x56, 0x95, 0xb8, 0x44, 0x1b,
	0xef, 0x34, 0x59, 0x62, 0xef, 0x46, 0x5e, 0x17, 0xc4, 0x9d, 0x2b, 0x7f, 0x83, 0x3b, 0xfc, 0x42,
	0x66, 0x3f, 0x9c, 0x04, 0x2a, 0xa1, 0x44, 0xbd, 0x58, 0xf6, 0xdb, 0xf1, 0x7b, 0x6f, 0xde, 0xce,
	0x40, 0x3f, 0xff, 0x3c, 0x99, 0x32, 0xa1, 0xe7, 0x4c, 0x4e, 0xe6, 0x2c, 0xcb, 0x06, 0xcb, 0x42,
	0x95, 0x2a, 0x6a, 0x72, 0x6e, 0x5f, 0x0e, 0x0f, 0x52, 0x95, 0xe7, 0x4a, 0x4e, 0xd2, 0x4c, 0xa0,
	0x2c, 0xdd, 0xe9, 0x61, 0x8f, 0x7e, 0x72, 0xb8, 0x07, 0x36, 0x59, 0xa6, 0x4c, 0xa3, 0x83, 0x93,
	0x4f, 0xd0, 0x9e, 0xb1, 0x1c, 0x27, 0x57, 0xaa, 0x14, 0x29, 0x46, 0xaf, 0xa0, 0x31, 0x47, 0xc6,
	0xb1, 0x88, 0x83, 0xa3, 0xe0, 0xb8, 0x7d, 0xf2, 0x62, 0xe0, 0x55, 0x06, 0x23, 0xf3, 0xbc, 0xb0,
	0x67, 0x51, 0x04, 0x20, 0x6d, 0xfd, 0xf5, 0xb7, 0x25, 0xc6, 0x35, 0xaa, 0xac, 0x47, 0xfb, 0xd0,
	0x4a, 0x89, 0x5b, 0x62, 0x76, 0xc9, 0xe3, 0x90, 0xa0, 0x56, 0xf2, 0x33, 0x80, 0xae, 0x25, 0x3f,
	0x4d, 0x17, 0x8f, 0xe6, 0x3f, 0x80, 0xb6, 0xc7, 0x44, 0x99, 0xa1, 0x53, 0x88, 0xfa, 0xd0, 0x71,
	0xe0, 0x99, 0x92, 0x25, 0x35, 0x1f, 0x3f, 0xb5, 0xf0, 0xea, 0xff, 0x8f, 0x98, 0xab, 0xb8, 0x6e,
	0x31, 0x80, 0x9a, 0xe0, 0x71, 0xc3, 0x72, 0x75, 0xa1, 0x71, 0x27, 0x32, 0xe4, 0x3a, 0x6e, 0x1e,
	0x85, 0x64, 0x34, 0x85, 0x9e, 0xf5, 0xf9, 0x9e, 0x1e, 0x63, 0x4c, 0x55, 0xc1, 0xb7, 0x34, 0xea,
	0x48, 0x6b, 0x15, 0xa9, 0x21, 0xf1, 0xdd, 0xdb, 0xef, 0x7b, 0x8d, 0x05, 0x7d, 0x1b, 0x53, 0x9d,
	0x64, 0x01, 0xdd, 0x21, 0x32, 0x79, 0x43, 0xd8, 0x4e, 0x1a, 0xc4, 0x73, 0xe3, 0x78, 0x8c, 0x4e,
	0x27, 0x7a, 0x0e, 0x7b, 0x57, 0x82, 0x02, 0x25, 0x2d, 0x9f, 0x02, 0x45, 0x7f, 0x2b, 0xe4, 0x69,
	0xae, 0xee, 0x7d, 0x02, 0x61, 0xf2, 0x23, 0x70, 0x6a, 0x8f, 0xed, 0x88, 0xa3, 0x5e, 0xac, 0x3a,
	0x22, 0x9d, 0x29, 0xce, 0x84, 0xbc, 0x16, 0x24, 0xed, 0x92, 0x7e, 0x0d, 0x75, 0xd3, 0xa4, 0xa6,
	0x90, 0x43, 0xe2, 0x7c, 0xb9, 0xe2, 0xfc, 0xbb, 0xd5, 0xe4, 0x7b, 0x00, 0x51, 0x35, 0x0a, 0x3b,
	0x7b, 0x5a, 0x27, 0xe9, 0x12, 0xf8, 0x37, 0xe9, 0x63, 0x68, 0x16, 0x96, 0x4f, 0x93, 0xab, 0x87,
	0x36, 0xd6, 0x7a, 0xc9, 0x39, 0x74, 0xac, 0x8b, 0x73, 0x44, 0x3e, 0x65, 0xe9, 0x62, 0x4b, 0x03,
	0x3d, 0x68, 0xe6, 0xa8, 0x35, 0x9b, 0xb9, 0x61, 0x6c, 0x25, 0xdc, 0x0f, 0xcc, 0x59, 0x81, 0xac,
	0xc4, 0xb1, 0x52, 0xf9, 0x96, 0x4c, 0x6f, 0xe0, 0x59, 0x41, 0xd5, 0x66, 0xae, 0x2f, 0xe5, 0x9d,
	0xb2, 0x74, 0xed, 0x93, 0xfe, 0xaa, 0x76, 0xbc, 0x71, 0x98, 0xfc, 0xda, 0x08, 0x6d, 0x67, 0xa5,
	0xf5, 0xe5, 0xb9, 0xcb, 0xa4, 0xb1, 0x59, 0x32, 0xad, 0xbf, 0x52, 0x0c, 0x7e, 0x6c, 0x68, 0xa3,
	0x4c, 0xac, 0x43, 0x96, 0x31, 0x99, 0xba, 0x0b, 0x0d, 0xed, 0x1a, 0x5b, 0x29, 0x8a, 0xc8, 0x6e,
	0x4e, 0xf8, 0xc0, 0x73, 0xe3, 0x7f, 0x9e, 0x7f, 0x57, 0x3b, 0xff, 0x8e, 0x16, 0xb2, 0xd8, 0xc1,
	0xaf, 0xc9, 0x98, 0x95, 0xe9, 0x7c, 0x65, 0x98, 0x80, 0x92, 0x4d, 0xb3, 0xf5, 0x35, 0x53, 0x07,
	0x23, 0xea, 0xe0, 0xd6, 0x74, 0xe0, 0xa6, 0x6f, 0x3d, 0x18, 0xf5, 0x6a, 0x35, 0x2a, 0xa7, 0x7e,
	0xd3, 0xa9, 0x1d, 0x83, 0x7c, 0xc0, 0x2f, 0x98, 0xd1, 0xb2, 0x7b, 0x08, 0x8d, 0x37, 0x5b, 0xb5,
	0x67, 0xa0, 0xe4, 0x2d, 0xec, 0x57, 0x39, 0xef, 0x68, 0x7b, 0x58, 0xbb, 0x08, 0x47, 0x4f, 0x46,
	0xc1, 0xa8, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x1b, 0xd3, 0xe2, 0xa7, 0x05, 0x00, 0x00,
}
