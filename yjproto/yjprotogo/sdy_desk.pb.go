// Code generated by protoc-gen-go.
// source: sdy_desk.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// 房主解散房间(未开局)
type SdyReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqDissolveDesk) Reset()                    { *m = SdyReqDissolveDesk{} }
func (m *SdyReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyReqDissolveDesk) ProtoMessage()               {}
func (*SdyReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SdyReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type SdyAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckDissolveDesk) Reset()                    { *m = SdyAckDissolveDesk{} }
func (m *SdyAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyAckDissolveDesk) ProtoMessage()               {}
func (*SdyAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SdyAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *SdyAckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type SdyReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqReady) Reset()                    { *m = SdyReqReady{} }
func (m *SdyReqReady) String() string            { return proto.CompactTextString(m) }
func (*SdyReqReady) ProtoMessage()               {}
func (*SdyReqReady) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *SdyReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type SdyAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckReady) Reset()                    { *m = SdyAckReady{} }
func (m *SdyAckReady) String() string            { return proto.CompactTextString(m) }
func (*SdyAckReady) ProtoMessage()               {}
func (*SdyAckReady) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *SdyAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *SdyAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type SdyBaseWinCoinInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	NickName         *string      `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	BaseValue        *int32       `protobuf:"varint,4,opt,name=baseValue" json:"baseValue,omitempty"`
	WinCoin          *int64       `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64       `protobuf:"varint,6,opt,name=coin" json:"coin,omitempty"`
	IsBanker         *bool        `protobuf:"varint,7,opt,name=isBanker" json:"isBanker,omitempty"`
	Rate             *int32       `protobuf:"varint,8,opt,name=rate" json:"rate,omitempty"`
	Description      *string      `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	HandPokers       []int32      `protobuf:"varint,10,rep,name=handPokers" json:"handPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBaseWinCoinInfo) Reset()                    { *m = SdyBaseWinCoinInfo{} }
func (m *SdyBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBaseWinCoinInfo) ProtoMessage()               {}
func (*SdyBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *SdyBaseWinCoinInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBaseWinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *SdyBaseWinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *SdyBaseWinCoinInfo) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *SdyBaseWinCoinInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *SdyBaseWinCoinInfo) GetHandPokers() []int32 {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

// 本局结果(广播)
type SdyBcCurrentResult struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*SdyBaseWinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	IsWangKou        *bool                 `protobuf:"varint,3,opt,name=isWangKou" json:"isWangKou,omitempty"`
	IsPoPai          *bool                 `protobuf:"varint,4,opt,name=isPoPai" json:"isPoPai,omitempty"`
	IsKouDi          *bool                 `protobuf:"varint,5,opt,name=isKouDi" json:"isKouDi,omitempty"`
	IsGuang          *bool                 `protobuf:"varint,6,opt,name=isGuang" json:"isGuang,omitempty"`
	IsJiao_ShangChe  *bool                 `protobuf:"varint,7,opt,name=isJiao_ShangChe,json=isJiaoShangChe" json:"isJiao_ShangChe,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SdyBcCurrentResult) Reset()                    { *m = SdyBcCurrentResult{} }
func (m *SdyBcCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcCurrentResult) ProtoMessage()               {}
func (*SdyBcCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *SdyBcCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcCurrentResult) GetWinCoinInfo() []*SdyBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *SdyBcCurrentResult) GetIsWangKou() bool {
	if m != nil && m.IsWangKou != nil {
		return *m.IsWangKou
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsPoPai() bool {
	if m != nil && m.IsPoPai != nil {
		return *m.IsPoPai
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsKouDi() bool {
	if m != nil && m.IsKouDi != nil {
		return *m.IsKouDi
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsGuang() bool {
	if m != nil && m.IsGuang != nil {
		return *m.IsGuang
	}
	return false
}

func (m *SdyBcCurrentResult) GetIsJiao_ShangChe() bool {
	if m != nil && m.IsJiao_ShangChe != nil {
		return *m.IsJiao_ShangChe
	}
	return false
}

type SdyBcEndLotteryInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NickName         *string      `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool        `protobuf:"varint,4,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool        `protobuf:"varint,5,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64       `protobuf:"varint,6,opt,name=winCoin" json:"winCoin,omitempty"`
	MaxWinCoin       *int32       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	CountBomb        *int32       `protobuf:"varint,8,opt,name=countBomb" json:"countBomb,omitempty"`
	CountWin         *int32       `protobuf:"varint,9,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32       `protobuf:"varint,10,opt,name=countLose" json:"countLose,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcEndLotteryInfo) Reset()                    { *m = SdyBcEndLotteryInfo{} }
func (m *SdyBcEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcEndLotteryInfo) ProtoMessage()               {}
func (*SdyBcEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *SdyBcEndLotteryInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcEndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *SdyBcEndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *SdyBcEndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *SdyBcEndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetMaxWinCoin() int32 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *SdyBcEndLotteryInfo) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

func init() {
	proto.RegisterType((*SdyReqDissolveDesk)(nil), "yjprotogo.sdy_req_dissolveDesk")
	proto.RegisterType((*SdyAckDissolveDesk)(nil), "yjprotogo.sdy_ack_dissolveDesk")
	proto.RegisterType((*SdyReqReady)(nil), "yjprotogo.sdy_req_ready")
	proto.RegisterType((*SdyAckReady)(nil), "yjprotogo.sdy_ack_ready")
	proto.RegisterType((*SdyBaseWinCoinInfo)(nil), "yjprotogo.sdy_base_winCoinInfo")
	proto.RegisterType((*SdyBcCurrentResult)(nil), "yjprotogo.sdy_bc_currentResult")
	proto.RegisterType((*SdyBcEndLotteryInfo)(nil), "yjprotogo.sdy_bc_endLotteryInfo")
}

var fileDescriptor5 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xfe, 0xd3, 0xac, 0x5d, 0xe2, 0x6a, 0xfb, 0x51, 0x80, 0xc9, 0xaa, 0xd0, 0x88, 0x7a, 0x43,
	0xae, 0x7a, 0xb1, 0x37, 0xa0, 0x9b, 0xc4, 0xca, 0x26, 0xa8, 0x8c, 0x44, 0xef, 0x88, 0xdc, 0xf8,
	0x90, 0x9a, 0x34, 0x76, 0xb1, 0x13, 0x46, 0x5f, 0x83, 0x07, 0xe0, 0x39, 0xb8, 0xe3, 0xd5, 0x90,
	0x1d, 0x37, 0xcd, 0x04, 0xbb, 0xd8, 0x04, 0x57, 0xf5, 0xf7, 0x7d, 0x3e, 0xfe, 0xce, 0xf9, 0x7a,
	0x82, 0x8e, 0x35, 0xdb, 0xa6, 0x0c, 0x74, 0x31, 0xd9, 0x28, 0x59, 0xc9, 0x28, 0xdc, 0x7e, 0xb2,
	0x87, 0x5c, 0x8e, 0x1e, 0x67, 0xb2, 0x2c, 0xa5, 0x48, 0xb3, 0x35, 0x07, 0x51, 0x35, 0xfa, 0xf8,
	0x03, 0x7a, 0x62, 0x2a, 0x14, 0x7c, 0x4e, 0x19, 0xd7, 0x5a, 0xae, 0xbf, 0xc0, 0x05, 0xe8, 0x22,
	0x9a, 0xa0, 0xc1, 0x0a, 0x28, 0x03, 0x85, 0xbd, 0xd8, 0x4b, 0x86, 0x67, 0x27, 0x93, 0xf6, 0xa1,
	0xc9, 0xdc, 0xfc, 0x5e, 0x5a, 0x95, 0xb8, 0x5b, 0xd1, 0x09, 0x1a, 0xd4, 0x1a, 0xd4, 0x8c, 0xe1,
	0x5e, 0xec, 0x25, 0x47, 0xc4, 0xa1, 0xf1, 0x37, 0xaf, 0x31, 0xa0, 0x59, 0xf1, 0x4f, 0x0c, 0x0c,
	0x6f, 0xc6, 0x9d, 0x31, 0xec, 0xc7, 0x5e, 0xd2, 0x27, 0x0e, 0x45, 0x23, 0x14, 0x6c, 0xa8, 0xd6,
	0x0b, 0xa9, 0x18, 0x3e, 0x88, 0xbd, 0x24, 0x24, 0x2d, 0x1e, 0x2f, 0xd0, 0xd1, 0x6e, 0x68, 0x05,
	0x94, 0x6d, 0xff, 0xda, 0xb4, 0xbc, 0x79, 0xd8, 0x0c, 0xfb, 0xb0, 0x87, 0x1f, 0x21, 0xbf, 0xd4,
	0xb9, 0x7d, 0x35, 0x24, 0xe6, 0xd8, 0xb1, 0xf2, 0x6f, 0x59, 0xfd, 0xe8, 0x35, 0xc1, 0x2e, 0xa9,
	0x86, 0xf4, 0x86, 0x8b, 0x73, 0xc9, 0xc5, 0x4c, 0x7c, 0x94, 0xf7, 0xb6, 0x1c, 0xa1, 0x40, 0xf0,
	0xac, 0x78, 0x43, 0x4b, 0x70, 0xbe, 0x2d, 0xbe, 0xcb, 0x3c, 0x7a, 0x86, 0x42, 0xe3, 0xfb, 0x9e,
	0xae, 0x6b, 0xb0, 0xe9, 0xf6, 0xc9, 0x9e, 0x88, 0x30, 0x3a, 0x74, 0x0d, 0xe1, 0x7e, 0xec, 0x25,
	0x3e, 0xd9, 0xc1, 0x28, 0x42, 0x07, 0x99, 0xa1, 0x07, 0x96, 0xb6, 0x67, 0xe3, 0xcf, 0xf5, 0x94,
	0x8a, 0x02, 0x14, 0x3e, 0x8c, 0xbd, 0x24, 0x20, 0x2d, 0x36, 0xf7, 0x15, 0xad, 0x00, 0x07, 0xd6,
	0xc2, 0x9e, 0xa3, 0x18, 0x0d, 0x19, 0xe8, 0x4c, 0xf1, 0x4d, 0xc5, 0xa5, 0xc0, 0xa1, 0x6d, 0xb9,
	0x4b, 0x45, 0xa7, 0x08, 0xad, 0xa8, 0x60, 0x73, 0x59, 0x80, 0xd2, 0x18, 0xc5, 0x7e, 0xd2, 0x27,
	0x1d, 0x66, 0xfc, 0x7d, 0x17, 0x5d, 0x96, 0x66, 0xb5, 0x52, 0x20, 0x2a, 0x02, 0xba, 0x5e, 0x57,
	0xf7, 0x8e, 0xee, 0x25, 0x1a, 0x76, 0x92, 0xc7, 0xbd, 0xd8, 0x4f, 0x86, 0x67, 0xcf, 0x3b, 0x45,
	0x7f, 0xfa, 0x83, 0x48, 0xb7, 0xc6, 0x24, 0xc9, 0xf5, 0x82, 0x8a, 0xfc, 0x4a, 0xd6, 0x36, 0xe4,
	0x80, 0xec, 0x09, 0x93, 0x24, 0xd7, 0x73, 0x39, 0xa7, 0xdc, 0xa6, 0x1c, 0x90, 0x1d, 0x6c, 0x94,
	0x2b, 0x59, 0x5f, 0x70, 0x9b, 0xb1, 0x55, 0x2c, 0x6c, 0x94, 0x57, 0x35, 0x15, 0xb9, 0x8d, 0xd9,
	0x2a, 0x16, 0x46, 0x2f, 0xd0, 0xff, 0x5c, 0xbf, 0xe6, 0x54, 0xa6, 0xef, 0x56, 0x54, 0xe4, 0xe7,
	0x2b, 0x70, 0x81, 0x1f, 0x37, 0xf4, 0x8e, 0x1d, 0xff, 0xec, 0xa1, 0xa7, 0x2e, 0x20, 0x10, 0xec,
	0x5a, 0x56, 0x15, 0xa8, 0xed, 0x83, 0x96, 0xeb, 0xae, 0xaf, 0xb6, 0xbb, 0x74, 0xfe, 0xef, 0x4b,
	0xb7, 0xe4, 0xf9, 0x82, 0x0b, 0x37, 0xb3, 0x43, 0xcd, 0x60, 0x6f, 0x6f, 0x04, 0xa8, 0xfd, 0xc8,
	0x16, 0x76, 0x17, 0x6e, 0x70, 0x7b, 0xe1, 0x4e, 0x11, 0x2a, 0xe9, 0xd7, 0x85, 0x13, 0x0f, 0xed,
	0x1a, 0x75, 0x18, 0x13, 0x7f, 0x26, 0x6b, 0x51, 0x4d, 0x65, 0xb9, 0x74, 0x5b, 0xb6, 0x27, 0x4c,
	0x97, 0x16, 0x98, 0x5e, 0x42, 0x2b, 0xb6, 0xb8, 0xad, 0xbc, 0x96, 0x1a, 0x30, 0xea, 0x54, 0x1a,
	0x62, 0xda, 0xbb, 0xf4, 0xe7, 0xff, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x59, 0xa4, 0xb7,
	0x8d, 0x05, 0x00, 0x00,
}
