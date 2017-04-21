// Code generated by protoc-gen-go.
// source: niuniu_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of niuniu_client_poker from niuniu_base.proto

// Ignoring public import of niuniu_user_bill from niuniu_base.proto

// Ignoring public import of niuniu_desk_option from niuniu_base.proto

// Ignoring public import of niuniu_enum_protoid from niuniu_base.proto

// Ignoring public import of niuniu_enum_PokerType from niuniu_base.proto

// Ignoring public import of niuniu_enum_desk_state from niuniu_base.proto

// Ignoring public import of niuniu_enum_banker_rule from niuniu_base.proto

// desk 的信息
type NiuniuClientDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	DeskNumber       *string              `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	GameNumber       *int32               `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status           *NiuniuEnumDeskState `protobuf:"varint,5,opt,name=status,enum=ddproto.NiuniuEnumDeskState" json:"status,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	CircleNo         *int32               `protobuf:"varint,8,opt,name=circleNo" json:"circleNo,omitempty"`
	Owner            *uint32              `protobuf:"varint,10,opt,name=owner" json:"owner,omitempty"`
	CurrBanker       *uint32              `protobuf:"varint,11,opt,name=currBanker" json:"currBanker,omitempty"`
	DeskOption       *NiuniuDeskOption    `protobuf:"bytes,12,opt,name=deskOption" json:"deskOption,omitempty"`
	Users            []*NiuniuClientUser  `protobuf:"bytes,13,rep,name=users" json:"users,omitempty"`
	IsStart          *bool                `protobuf:"varint,14,opt,name=isStart" json:"isStart,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuClientDesk) Reset()                    { *m = NiuniuClientDesk{} }
func (m *NiuniuClientDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientDesk) ProtoMessage()               {}
func (*NiuniuClientDesk) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

func (m *NiuniuClientDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuClientDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuClientDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuClientDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuClientDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuClientDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuClientDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuClientDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuClientDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuClientDesk) GetDeskOption() *NiuniuDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *NiuniuClientDesk) GetUsers() []*NiuniuClientUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *NiuniuClientDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

// 用户属性
type NiuniuClientUser struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *NiuniuUserBill    `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool              `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32             `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuClientPoker `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	BankerScore      *int32             `protobuf:"varint,14,opt,name=bankerScore" json:"bankerScore,omitempty"`
	DoubleScore      *int32             `protobuf:"varint,15,opt,name=doubleScore" json:"doubleScore,omitempty"`
	IsReady          *bool              `protobuf:"varint,16,opt,name=isReady" json:"isReady,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *NiuniuClientUser) Reset()                    { *m = NiuniuClientUser{} }
func (m *NiuniuClientUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientUser) ProtoMessage()               {}
func (*NiuniuClientUser) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{1} }

func (m *NiuniuClientUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuClientUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuClientUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuClientUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuClientUser) GetPokers() *NiuniuClientPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *NiuniuClientUser) GetBankerScore() int32 {
	if m != nil && m.BankerScore != nil {
		return *m.BankerScore
	}
	return 0
}

func (m *NiuniuClientUser) GetDoubleScore() int32 {
	if m != nil && m.DoubleScore != nil {
		return *m.DoubleScore
	}
	return 0
}

func (m *NiuniuClientUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

// ==============================创建房间===============================
type NiuCreateDeskReq struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Option           *NiuniuDeskOption `protobuf:"bytes,2,opt,name=option" json:"option,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *NiuCreateDeskReq) Reset()                    { *m = NiuCreateDeskReq{} }
func (m *NiuCreateDeskReq) String() string            { return proto.CompactTextString(m) }
func (*NiuCreateDeskReq) ProtoMessage()               {}
func (*NiuCreateDeskReq) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{2} }

func (m *NiuCreateDeskReq) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuCreateDeskReq) GetOption() *NiuniuDeskOption {
	if m != nil {
		return m.Option
	}
	return nil
}

// ==============================进入房间===============================
type NiuEnterDeskReq struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskNumber       *string      `protobuf:"bytes,2,opt,name=deskNumber" json:"deskNumber,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuEnterDeskReq) Reset()                    { *m = NiuEnterDeskReq{} }
func (m *NiuEnterDeskReq) String() string            { return proto.CompactTextString(m) }
func (*NiuEnterDeskReq) ProtoMessage()               {}
func (*NiuEnterDeskReq) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{3} }

func (m *NiuEnterDeskReq) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuEnterDeskReq) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

type NiuEnterDeskAck struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskState        *NiuniuClientDesk `protobuf:"bytes,2,opt,name=deskState" json:"deskState,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *NiuEnterDeskAck) Reset()                    { *m = NiuEnterDeskAck{} }
func (m *NiuEnterDeskAck) String() string            { return proto.CompactTextString(m) }
func (*NiuEnterDeskAck) ProtoMessage()               {}
func (*NiuEnterDeskAck) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{4} }

func (m *NiuEnterDeskAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuEnterDeskAck) GetDeskState() *NiuniuClientDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

type NiuEnterDeskBc struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *NiuniuClientUser `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *NiuEnterDeskBc) Reset()                    { *m = NiuEnterDeskBc{} }
func (m *NiuEnterDeskBc) String() string            { return proto.CompactTextString(m) }
func (*NiuEnterDeskBc) ProtoMessage()               {}
func (*NiuEnterDeskBc) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{5} }

func (m *NiuEnterDeskBc) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuEnterDeskBc) GetUser() *NiuniuClientUser {
	if m != nil {
		return m.User
	}
	return nil
}

// =============================准备===================================
type NiuSwitchReadyReq struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsReady          *bool        `protobuf:"varint,2,opt,name=isReady" json:"isReady,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuSwitchReadyReq) Reset()                    { *m = NiuSwitchReadyReq{} }
func (m *NiuSwitchReadyReq) String() string            { return proto.CompactTextString(m) }
func (*NiuSwitchReadyReq) ProtoMessage()               {}
func (*NiuSwitchReadyReq) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{6} }

func (m *NiuSwitchReadyReq) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuSwitchReadyReq) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

type NiuSwitchReadyAck struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuSwitchReadyAck) Reset()                    { *m = NiuSwitchReadyAck{} }
func (m *NiuSwitchReadyAck) String() string            { return proto.CompactTextString(m) }
func (*NiuSwitchReadyAck) ProtoMessage()               {}
func (*NiuSwitchReadyAck) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{7} }

func (m *NiuSwitchReadyAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type NiuSwitchReadyBc struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *uint32      `protobuf:"varint,2,opt,name=user" json:"user,omitempty"`
	IsReady          *bool        `protobuf:"varint,3,opt,name=isReady" json:"isReady,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuSwitchReadyBc) Reset()                    { *m = NiuSwitchReadyBc{} }
func (m *NiuSwitchReadyBc) String() string            { return proto.CompactTextString(m) }
func (*NiuSwitchReadyBc) ProtoMessage()               {}
func (*NiuSwitchReadyBc) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{8} }

func (m *NiuSwitchReadyBc) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuSwitchReadyBc) GetUser() uint32 {
	if m != nil && m.User != nil {
		return *m.User
	}
	return 0
}

func (m *NiuSwitchReadyBc) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

// ==============================房主开局===================================
type NiuStartGameOt struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuStartGameOt) Reset()                    { *m = NiuStartGameOt{} }
func (m *NiuStartGameOt) String() string            { return proto.CompactTextString(m) }
func (*NiuStartGameOt) ProtoMessage()               {}
func (*NiuStartGameOt) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{9} }

func (m *NiuStartGameOt) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// =============================抢庄(发牌)===================================
type NiuQiangzhuangOt struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Pokers           *NiuniuClientPoker `protobuf:"bytes,2,opt,name=pokers" json:"pokers,omitempty"`
	CurrCircle       *int32             `protobuf:"varint,3,opt,name=currCircle" json:"currCircle,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *NiuQiangzhuangOt) Reset()                    { *m = NiuQiangzhuangOt{} }
func (m *NiuQiangzhuangOt) String() string            { return proto.CompactTextString(m) }
func (*NiuQiangzhuangOt) ProtoMessage()               {}
func (*NiuQiangzhuangOt) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{10} }

func (m *NiuQiangzhuangOt) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuQiangzhuangOt) GetPokers() *NiuniuClientPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *NiuQiangzhuangOt) GetCurrCircle() int32 {
	if m != nil && m.CurrCircle != nil {
		return *m.CurrCircle
	}
	return 0
}

type NiuQiangzhuangReq struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Score            *int32       `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuQiangzhuangReq) Reset()                    { *m = NiuQiangzhuangReq{} }
func (m *NiuQiangzhuangReq) String() string            { return proto.CompactTextString(m) }
func (*NiuQiangzhuangReq) ProtoMessage()               {}
func (*NiuQiangzhuangReq) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{11} }

func (m *NiuQiangzhuangReq) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuQiangzhuangReq) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

type NiuQiangzhuangAck struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuQiangzhuangAck) Reset()                    { *m = NiuQiangzhuangAck{} }
func (m *NiuQiangzhuangAck) String() string            { return proto.CompactTextString(m) }
func (*NiuQiangzhuangAck) ProtoMessage()               {}
func (*NiuQiangzhuangAck) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{12} }

func (m *NiuQiangzhuangAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type NiuQiangzhuangResItem struct {
	User             *uint32 `protobuf:"varint,1,opt,name=user" json:"user,omitempty"`
	Score            *int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	IsBanker         *bool   `protobuf:"varint,3,opt,name=isBanker" json:"isBanker,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuQiangzhuangResItem) Reset()                    { *m = NiuQiangzhuangResItem{} }
func (m *NiuQiangzhuangResItem) String() string            { return proto.CompactTextString(m) }
func (*NiuQiangzhuangResItem) ProtoMessage()               {}
func (*NiuQiangzhuangResItem) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{13} }

func (m *NiuQiangzhuangResItem) GetUser() uint32 {
	if m != nil && m.User != nil {
		return *m.User
	}
	return 0
}

func (m *NiuQiangzhuangResItem) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuQiangzhuangResItem) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

// 抢庄结果广播
type NiuQiangzhuangResBc struct {
	Header           *ProtoHeader             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Result           []*NiuQiangzhuangResItem `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *NiuQiangzhuangResBc) Reset()                    { *m = NiuQiangzhuangResBc{} }
func (m *NiuQiangzhuangResBc) String() string            { return proto.CompactTextString(m) }
func (*NiuQiangzhuangResBc) ProtoMessage()               {}
func (*NiuQiangzhuangResBc) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{14} }

func (m *NiuQiangzhuangResBc) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuQiangzhuangResBc) GetResult() []*NiuQiangzhuangResItem {
	if m != nil {
		return m.Result
	}
	return nil
}

// =============================加倍（发牌）===================================
type NiuJiabeiOt struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CuurBanker       *int32             `protobuf:"varint,2,opt,name=cuurBanker" json:"cuurBanker,omitempty"`
	Pokers           *NiuniuClientPoker `protobuf:"bytes,3,opt,name=pokers" json:"pokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *NiuJiabeiOt) Reset()                    { *m = NiuJiabeiOt{} }
func (m *NiuJiabeiOt) String() string            { return proto.CompactTextString(m) }
func (*NiuJiabeiOt) ProtoMessage()               {}
func (*NiuJiabeiOt) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{15} }

func (m *NiuJiabeiOt) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuJiabeiOt) GetCuurBanker() int32 {
	if m != nil && m.CuurBanker != nil {
		return *m.CuurBanker
	}
	return 0
}

func (m *NiuJiabeiOt) GetPokers() *NiuniuClientPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

type NiuJiabeiReq struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Score            *int32       `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	CurrCircle       *int32       `protobuf:"varint,3,opt,name=currCircle" json:"currCircle,omitempty"`
	CuurBanker       *int32       `protobuf:"varint,4,opt,name=cuurBanker" json:"cuurBanker,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuJiabeiReq) Reset()                    { *m = NiuJiabeiReq{} }
func (m *NiuJiabeiReq) String() string            { return proto.CompactTextString(m) }
func (*NiuJiabeiReq) ProtoMessage()               {}
func (*NiuJiabeiReq) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{16} }

func (m *NiuJiabeiReq) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuJiabeiReq) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuJiabeiReq) GetCurrCircle() int32 {
	if m != nil && m.CurrCircle != nil {
		return *m.CurrCircle
	}
	return 0
}

func (m *NiuJiabeiReq) GetCuurBanker() int32 {
	if m != nil && m.CuurBanker != nil {
		return *m.CuurBanker
	}
	return 0
}

type NiuJiabeiAck struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuJiabeiAck) Reset()                    { *m = NiuJiabeiAck{} }
func (m *NiuJiabeiAck) String() string            { return proto.CompactTextString(m) }
func (*NiuJiabeiAck) ProtoMessage()               {}
func (*NiuJiabeiAck) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{17} }

func (m *NiuJiabeiAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type NiuJiabeiBc struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Score            *int32       `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NiuJiabeiBc) Reset()                    { *m = NiuJiabeiBc{} }
func (m *NiuJiabeiBc) String() string            { return proto.CompactTextString(m) }
func (*NiuJiabeiBc) ProtoMessage()               {}
func (*NiuJiabeiBc) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{18} }

func (m *NiuJiabeiBc) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuJiabeiBc) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// ==============================跟庄家比牌结果============================
type NiuBipaiResultItem struct {
	Poker            *NiuniuClientPoker `protobuf:"bytes,1,opt,name=poker" json:"poker,omitempty"`
	Score            *int32             `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *NiuBipaiResultItem) Reset()                    { *m = NiuBipaiResultItem{} }
func (m *NiuBipaiResultItem) String() string            { return proto.CompactTextString(m) }
func (*NiuBipaiResultItem) ProtoMessage()               {}
func (*NiuBipaiResultItem) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{19} }

func (m *NiuBipaiResultItem) GetPoker() *NiuniuClientPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *NiuBipaiResultItem) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

type NiuBipaiResultBc struct {
	UserState        []*NiuBipaiResultItem `protobuf:"bytes,1,rep,name=userState" json:"userState,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuBipaiResultBc) Reset()                    { *m = NiuBipaiResultBc{} }
func (m *NiuBipaiResultBc) String() string            { return proto.CompactTextString(m) }
func (*NiuBipaiResultBc) ProtoMessage()               {}
func (*NiuBipaiResultBc) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{20} }

func (m *NiuBipaiResultBc) GetUserState() []*NiuBipaiResultItem {
	if m != nil {
		return m.UserState
	}
	return nil
}

type NiuGameEnd struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data             []*NiuniuUserBill `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *NiuGameEnd) Reset()                    { *m = NiuGameEnd{} }
func (m *NiuGameEnd) String() string            { return proto.CompactTextString(m) }
func (*NiuGameEnd) ProtoMessage()               {}
func (*NiuGameEnd) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{21} }

func (m *NiuGameEnd) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *NiuGameEnd) GetData() []*NiuniuUserBill {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*NiuniuClientDesk)(nil), "ddproto.niuniu_client_desk")
	proto.RegisterType((*NiuniuClientUser)(nil), "ddproto.niuniu_client_user")
	proto.RegisterType((*NiuCreateDeskReq)(nil), "ddproto.niu_create_desk_req")
	proto.RegisterType((*NiuEnterDeskReq)(nil), "ddproto.niu_enter_desk_req")
	proto.RegisterType((*NiuEnterDeskAck)(nil), "ddproto.niu_enter_desk_ack")
	proto.RegisterType((*NiuEnterDeskBc)(nil), "ddproto.niu_enter_desk_bc")
	proto.RegisterType((*NiuSwitchReadyReq)(nil), "ddproto.niu_switch_ready_req")
	proto.RegisterType((*NiuSwitchReadyAck)(nil), "ddproto.niu_switch_ready_ack")
	proto.RegisterType((*NiuSwitchReadyBc)(nil), "ddproto.niu_switch_ready_bc")
	proto.RegisterType((*NiuStartGameOt)(nil), "ddproto.niu_start_game_ot")
	proto.RegisterType((*NiuQiangzhuangOt)(nil), "ddproto.niu_qiangzhuang_ot")
	proto.RegisterType((*NiuQiangzhuangReq)(nil), "ddproto.niu_qiangzhuang_req")
	proto.RegisterType((*NiuQiangzhuangAck)(nil), "ddproto.niu_qiangzhuang_ack")
	proto.RegisterType((*NiuQiangzhuangResItem)(nil), "ddproto.niu_qiangzhuang_res_item")
	proto.RegisterType((*NiuQiangzhuangResBc)(nil), "ddproto.niu_qiangzhuang_res_bc")
	proto.RegisterType((*NiuJiabeiOt)(nil), "ddproto.niu_jiabei_ot")
	proto.RegisterType((*NiuJiabeiReq)(nil), "ddproto.niu_jiabei_req")
	proto.RegisterType((*NiuJiabeiAck)(nil), "ddproto.niu_jiabei_ack")
	proto.RegisterType((*NiuJiabeiBc)(nil), "ddproto.niu_jiabei_bc")
	proto.RegisterType((*NiuBipaiResultItem)(nil), "ddproto.niu_bipai_result_item")
	proto.RegisterType((*NiuBipaiResultBc)(nil), "ddproto.niu_bipai_result_bc")
	proto.RegisterType((*NiuGameEnd)(nil), "ddproto.niu_game_end")
}

var fileDescriptor29 = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x6d, 0x4f, 0x13, 0x41,
	0x10, 0xb6, 0x2d, 0x57, 0xda, 0x29, 0x2d, 0xb0, 0xa0, 0x39, 0x5f, 0xa2, 0x78, 0x31, 0xf1, 0x05,
	0x53, 0x02, 0x1f, 0x4c, 0x8c, 0x7e, 0x52, 0x3f, 0x80, 0x89, 0x40, 0xec, 0x07, 0x13, 0x13, 0x73,
	0xd9, 0xbb, 0xdb, 0xd0, 0xa5, 0xd7, 0xbb, 0x72, 0xbb, 0x17, 0xc4, 0xf8, 0xef, 0xfc, 0x11, 0xfe,
	0x1d, 0x77, 0x67, 0x0f, 0x72, 0xbd, 0xb6, 0xb0, 0x98, 0x90, 0x96, 0x9b, 0x9b, 0x9d, 0x67, 0xe6,
	0x79, 0x66, 0x66, 0x0b, 0xeb, 0x09, 0xcf, 0xd5, 0x9f, 0x3f, 0x89, 0xe9, 0x45, 0x7f, 0x92, 0xa5,
	0x32, 0x25, 0xcb, 0x51, 0x84, 0xff, 0x3c, 0xb8, 0x7c, 0x17, 0x50, 0xc1, 0xcc, 0x3b, 0xef, 0x4f,
	0x1d, 0x48, 0x61, 0x0d, 0x63, 0xce, 0x12, 0xe9, 0x47, 0x4c, 0x8c, 0x48, 0x0f, 0x9a, 0xfa, 0xfb,
	0x20, 0x72, 0x6b, 0x5b, 0xb5, 0x17, 0x0e, 0x21, 0x00, 0xfa, 0xf9, 0x30, 0x1f, 0x07, 0x2c, 0x73,
	0xeb, 0xca, 0xd6, 0xd6, 0xb6, 0x13, 0x3a, 0x66, 0x85, 0xad, 0x81, 0x7e, 0xea, 0x5c, 0x96, 0xa6,
	0x63, 0x75, 0x6e, 0x09, 0x9f, 0x77, 0xa0, 0x29, 0x24, 0x95, 0xb9, 0x70, 0x1d, 0xf5, 0xdc, 0xdb,
	0x7b, 0xd2, 0x2f, 0x72, 0xe9, 0x17, 0xa0, 0x2c, 0xc9, 0xc7, 0x08, 0xe9, 0x6b, 0x3f, 0x46, 0xd6,
	0xa1, 0x1d, 0x53, 0x21, 0xbf, 0xf1, 0x44, 0xc5, 0x6c, 0xaa, 0x33, 0x5d, 0xb2, 0x06, 0xad, 0x90,
	0x67, 0x61, 0xcc, 0x0e, 0x53, 0xb7, 0x85, 0x51, 0xbb, 0xe0, 0xa4, 0xe7, 0xda, 0x01, 0xd0, 0x41,
	0x25, 0x12, 0xe6, 0x59, 0xf6, 0x81, 0x26, 0x23, 0x65, 0xeb, 0xa0, 0x6d, 0xc7, 0x24, 0x7c, 0x34,
	0x91, 0x3c, 0x4d, 0xdc, 0x15, 0x65, 0xeb, 0xec, 0x3d, 0xac, 0x82, 0x23, 0x6e, 0x8a, 0x2e, 0xe4,
	0x15, 0x38, 0xb9, 0x60, 0x99, 0x70, 0xbb, 0x5b, 0x8d, 0x79, 0xbe, 0x05, 0x3b, 0xda, 0x87, 0xac,
	0xc2, 0x32, 0x17, 0x03, 0x49, 0x33, 0xe9, 0xf6, 0x54, 0xe4, 0x96, 0xf7, 0xb7, 0x56, 0x65, 0x11,
	0xfd, 0x14, 0x1b, 0xfa, 0xbb, 0x60, 0xb1, 0x4b, 0x9e, 0xc3, 0x52, 0xc0, 0xe3, 0x18, 0xeb, 0xea,
	0xec, 0xdd, 0xaf, 0x42, 0x68, 0x5f, 0x5f, 0x3b, 0xe8, 0x92, 0xb9, 0x38, 0x4a, 0x62, 0xc5, 0x02,
	0xd6, 0xd8, 0xd2, 0x25, 0xf3, 0x24, 0x62, 0x3f, 0xb1, 0x14, 0x87, 0xbc, 0x86, 0xe6, 0x24, 0x1d,
	0x99, 0x74, 0x75, 0xac, 0x47, 0x0b, 0xd2, 0x45, 0x27, 0xb2, 0x01, 0x9d, 0x00, 0xc9, 0x19, 0x84,
	0x69, 0xc6, 0x30, 0x67, 0x47, 0x1b, 0xa3, 0x34, 0x0f, 0x62, 0x66, 0x8c, 0xab, 0x68, 0xc4, 0xca,
	0xbe, 0x32, 0x1a, 0x5d, 0xb8, 0x6b, 0x58, 0xd9, 0x10, 0x36, 0x30, 0x5c, 0xc6, 0x94, 0x3a, 0x86,
	0xb0, 0x8c, 0x9d, 0x91, 0x67, 0xd0, 0x1c, 0x2a, 0x2f, 0x45, 0x77, 0x0d, 0xf1, 0x37, 0xaf, 0xf0,
	0x8f, 0xf5, 0xe7, 0x3e, 0xbe, 0x23, 0xdb, 0xd0, 0x34, 0xec, 0x62, 0xc7, 0x5c, 0x2f, 0x80, 0x77,
	0x88, 0x14, 0xaa, 0x86, 0x90, 0x8a, 0x85, 0x5b, 0x02, 0xcd, 0x69, 0x4f, 0xef, 0x74, 0x26, 0x1e,
	0x0d, 0x47, 0x96, 0xf1, 0xfa, 0xd0, 0xd6, 0x27, 0x06, 0xba, 0x25, 0x17, 0xe5, 0x5e, 0x1a, 0x17,
	0x2f, 0xc2, 0xb1, 0x2b, 0x63, 0x05, 0xa1, 0x25, 0xd4, 0x4b, 0x58, 0xd2, 0xba, 0xdf, 0x80, 0xa2,
	0x5d, 0xbc, 0x2f, 0xb0, 0xa9, 0x4d, 0xe2, 0x9c, 0xcb, 0x70, 0xa8, 0xd8, 0x51, 0x2a, 0xdd, 0x82,
	0xa3, 0x92, 0xb4, 0x75, 0x94, 0xf6, 0xfd, 0x9c, 0x70, 0xd6, 0x14, 0x79, 0xdf, 0x4d, 0x63, 0x4c,
	0x9d, 0xb6, 0x2e, 0x7a, 0xa5, 0x54, 0x74, 0xb7, 0x9c, 0x59, 0x03, 0x33, 0x7b, 0x6b, 0xe8, 0x14,
	0x7a, 0xc2, 0x7c, 0xbd, 0x63, 0xfc, 0x54, 0x5a, 0xa6, 0xf5, 0xdb, 0xa8, 0x7e, 0xc6, 0x69, 0x72,
	0xf2, 0x6b, 0x98, 0xab, 0x4f, 0xeb, 0xb3, 0xa5, 0xa1, 0xaa, 0x5b, 0x0c, 0x55, 0xb1, 0x75, 0x3e,
	0xe2, 0x6a, 0x32, 0xeb, 0xcf, 0xfb, 0x6c, 0x48, 0x29, 0xa3, 0xdb, 0x0b, 0xa4, 0x46, 0x5c, 0xe0,
	0x28, 0xd6, 0x31, 0xd6, 0xbb, 0xd9, 0x58, 0xf6, 0xea, 0x1c, 0x80, 0x3b, 0x9b, 0x88, 0xf0, 0xb9,
	0x64, 0xe3, 0x2b, 0xf2, 0xcd, 0x4e, 0x9a, 0x46, 0x35, 0x9b, 0xa7, 0xd8, 0xa4, 0x46, 0x8c, 0x33,
	0xb8, 0x37, 0x2f, 0x94, 0xb5, 0xd6, 0xbb, 0xea, 0x4a, 0x60, 0x22, 0x8f, 0xa5, 0x42, 0xd0, 0x9b,
	0xf5, 0x69, 0x99, 0xd5, 0xb9, 0x19, 0x7a, 0xe7, 0xd0, 0xd5, 0xef, 0x4e, 0x39, 0x0d, 0x18, 0xb7,
	0xd7, 0x0f, 0x15, 0xc9, 0x2f, 0xef, 0x81, 0x7a, 0x65, 0x51, 0x36, 0x6e, 0xd6, 0xd4, 0x1b, 0x43,
	0xaf, 0x04, 0xfc, 0xbf, 0xd2, 0xcd, 0x6b, 0x8d, 0x4a, 0x72, 0x78, 0x3b, 0x7a, 0x6f, 0xa6, 0xe0,
	0xec, 0xd5, 0xfd, 0x34, 0xc5, 0x8f, 0xb5, 0x12, 0x95, 0x06, 0x1b, 0xc0, 0x5d, 0xfc, 0x31, 0xc0,
	0x27, 0x54, 0xd7, 0xaa, 0x25, 0x32, 0x0d, 0xb2, 0x0d, 0x0e, 0xd2, 0x51, 0x04, 0xbb, 0x7e, 0x0c,
	0x2a, 0x41, 0xf7, 0x4d, 0xd7, 0x4e, 0x05, 0x55, 0x09, 0xee, 0x42, 0x5b, 0xf7, 0x9c, 0x59, 0xa8,
	0x35, 0xec, 0x83, 0xc7, 0x53, 0x7d, 0x30, 0x93, 0x85, 0xf7, 0x03, 0x56, 0xf4, 0x0b, 0x1c, 0x7f,
	0x96, 0x44, 0x96, 0x35, 0xaa, 0x2b, 0x36, 0xa2, 0x92, 0x16, 0xbd, 0xb6, 0xf8, 0x8a, 0x3d, 0xbe,
	0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x6e, 0xa9, 0xb0, 0x2a, 0x09, 0x00, 0x00,
}
