// Code generated by protoc-gen-go. DO NOT EDIT.
// source: niuniu_server.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of niuniu_client_poker from niuniu_base.proto

// Ignoring public import of niuniu_user_bill from niuniu_base.proto

// Ignoring public import of niuniu_desk_option from niuniu_base.proto

// Ignoring public import of niuniu_enum_protoid from niuniu_base.proto

// Ignoring public import of niuniu_enum_PokerType from niuniu_base.proto

// Ignoring public import of niuniu_enum_desk_state from niuniu_base.proto

// Ignoring public import of niuniu_enum_banker_rule from niuniu_base.proto

// 打出去的牌
type NiuniuSrvPoker struct {
	Pais             []*CommonSrvPokerPai  `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuSrvPoker) Reset()                    { *m = NiuniuSrvPoker{} }
func (m *NiuniuSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvPoker) ProtoMessage()               {}
func (*NiuniuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

func (m *NiuniuSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuSrvPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// desk 的信息
type NiuniuSrvDesk struct {
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
	IsStart          *bool                `protobuf:"varint,13,opt,name=isStart" json:"isStart,omitempty"`
	IsOnDissolve     *bool                `protobuf:"varint,14,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime     *int64               `protobuf:"varint,15,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	OneStartTime     *int64               `protobuf:"varint,16,opt,name=oneStartTime" json:"oneStartTime,omitempty"`
	AllStartTime     *int64               `protobuf:"varint,17,opt,name=allStartTime" json:"allStartTime,omitempty"`
	DissolveUser     *uint32              `protobuf:"varint,18,opt,name=dissolveUser" json:"dissolveUser,omitempty"`
	IsDaikai         *bool                `protobuf:"varint,19,opt,name=isDaikai" json:"isDaikai,omitempty"`
	DaikaiUser       *uint32              `protobuf:"varint,20,opt,name=daikaiUser" json:"daikaiUser,omitempty"`
	IsCoinRoom       *bool                `protobuf:"varint,21,opt,name=isCoinRoom" json:"isCoinRoom,omitempty"`
	IsOnGamming      *bool                `protobuf:"varint,22,opt,name=isOnGamming" json:"isOnGamming,omitempty"`
	SurplusTime      *int32               `protobuf:"varint,23,opt,name=surplusTime" json:"surplusTime,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvDesk) Reset()                    { *m = NiuniuSrvDesk{} }
func (m *NiuniuSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDesk) ProtoMessage()               {}
func (*NiuniuSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{1} }

func (m *NiuniuSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskNumber() string {
	if m != nil && m.DeskNumber != nil {
		return *m.DeskNumber
	}
	return ""
}

func (m *NiuniuSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *NiuniuSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvDesk) GetStatus() NiuniuEnumDeskState {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER
}

func (m *NiuniuSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *NiuniuSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *NiuniuSrvDesk) GetCurrBanker() uint32 {
	if m != nil && m.CurrBanker != nil {
		return *m.CurrBanker
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDeskOption() *NiuniuDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *NiuniuSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *NiuniuSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *NiuniuSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

func (m *NiuniuSrvDesk) GetOneStartTime() int64 {
	if m != nil && m.OneStartTime != nil {
		return *m.OneStartTime
	}
	return 0
}

func (m *NiuniuSrvDesk) GetAllStartTime() int64 {
	if m != nil && m.AllStartTime != nil {
		return *m.AllStartTime
	}
	return 0
}

func (m *NiuniuSrvDesk) GetDissolveUser() uint32 {
	if m != nil && m.DissolveUser != nil {
		return *m.DissolveUser
	}
	return 0
}

func (m *NiuniuSrvDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *NiuniuSrvDesk) GetDaikaiUser() uint32 {
	if m != nil && m.DaikaiUser != nil {
		return *m.DaikaiUser
	}
	return 0
}

func (m *NiuniuSrvDesk) GetIsCoinRoom() bool {
	if m != nil && m.IsCoinRoom != nil {
		return *m.IsCoinRoom
	}
	return false
}

func (m *NiuniuSrvDesk) GetIsOnGamming() bool {
	if m != nil && m.IsOnGamming != nil {
		return *m.IsOnGamming
	}
	return false
}

func (m *NiuniuSrvDesk) GetSurplusTime() int32 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

// 用户属性
type NiuniuSrvUser struct {
	UserId        *uint32         `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName      *string         `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Bill          *NiuniuUserBill `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	IsOnline      *bool           `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index         *int32          `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers        *NiuniuSrvPoker `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	BankerScore   *int32          `protobuf:"varint,14,opt,name=bankerScore" json:"bankerScore,omitempty"`
	DoubleScore   *int32          `protobuf:"varint,15,opt,name=doubleScore" json:"doubleScore,omitempty"`
	IsReady       *bool           `protobuf:"varint,16,opt,name=isReady" json:"isReady,omitempty"`
	LastScore     *int32          `protobuf:"varint,17,opt,name=lastScore" json:"lastScore,omitempty"`
	DissolveState *int32          `protobuf:"varint,18,opt,name=dissolveState" json:"dissolveState,omitempty"`
	IsOnWhiteList *bool           `protobuf:"varint,19,opt,name=isOnWhiteList" json:"isOnWhiteList,omitempty"`
	WhiteWinRate  *int32          `protobuf:"varint,20,opt,name=whiteWinRate" json:"whiteWinRate,omitempty"`
	// 金币场
	IsOnGamming      *bool  `protobuf:"varint,21,opt,name=isOnGamming" json:"isOnGamming,omitempty"`
	IsLeave          *bool  `protobuf:"varint,22,opt,name=isLeave" json:"isLeave,omitempty"`
	IsXuanpai        *bool  `protobuf:"varint,23,opt,name=isXuanpai" json:"isXuanpai,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NiuniuSrvUser) Reset()                    { *m = NiuniuSrvUser{} }
func (m *NiuniuSrvUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvUser) ProtoMessage()               {}
func (*NiuniuSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{2} }

func (m *NiuniuSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *NiuniuSrvUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *NiuniuSrvUser) GetBill() *NiuniuUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *NiuniuSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *NiuniuSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *NiuniuSrvUser) GetPokers() *NiuniuSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *NiuniuSrvUser) GetBankerScore() int32 {
	if m != nil && m.BankerScore != nil {
		return *m.BankerScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDoubleScore() int32 {
	if m != nil && m.DoubleScore != nil {
		return *m.DoubleScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *NiuniuSrvUser) GetLastScore() int32 {
	if m != nil && m.LastScore != nil {
		return *m.LastScore
	}
	return 0
}

func (m *NiuniuSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

func (m *NiuniuSrvUser) GetIsOnWhiteList() bool {
	if m != nil && m.IsOnWhiteList != nil {
		return *m.IsOnWhiteList
	}
	return false
}

func (m *NiuniuSrvUser) GetWhiteWinRate() int32 {
	if m != nil && m.WhiteWinRate != nil {
		return *m.WhiteWinRate
	}
	return 0
}

func (m *NiuniuSrvUser) GetIsOnGamming() bool {
	if m != nil && m.IsOnGamming != nil {
		return *m.IsOnGamming
	}
	return false
}

func (m *NiuniuSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *NiuniuSrvUser) GetIsXuanpai() bool {
	if m != nil && m.IsXuanpai != nil {
		return *m.IsXuanpai
	}
	return false
}

// room 的信息
type NiuniuSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuSrvRoom) Reset()                    { *m = NiuniuSrvRoom{} }
func (m *NiuniuSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvRoom) ProtoMessage()               {}
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{3} }

func (m *NiuniuSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *NiuniuSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

// desk快照索引列表
type NiuniuSrvDeskSnapshotIdIndex struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuSrvDeskSnapshotIdIndex) Reset()                    { *m = NiuniuSrvDeskSnapshotIdIndex{} }
func (m *NiuniuSrvDeskSnapshotIdIndex) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDeskSnapshotIdIndex) ProtoMessage()               {}
func (*NiuniuSrvDeskSnapshotIdIndex) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{4} }

func (m *NiuniuSrvDeskSnapshotIdIndex) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

// 牌桌快照
type NiuniuSrvDeskSnapshot struct {
	DeskState        *NiuniuSrvDesk   `protobuf:"bytes,1,opt,name=deskState" json:"deskState,omitempty"`
	Users            []*NiuniuSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *NiuniuSrvDeskSnapshot) Reset()                    { *m = NiuniuSrvDeskSnapshot{} }
func (m *NiuniuSrvDeskSnapshot) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDeskSnapshot) ProtoMessage()               {}
func (*NiuniuSrvDeskSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{5} }

func (m *NiuniuSrvDeskSnapshot) GetDeskState() *NiuniuSrvDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *NiuniuSrvDeskSnapshot) GetUsers() []*NiuniuSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*NiuniuSrvPoker)(nil), "ddproto.niuniu_srv_poker")
	proto.RegisterType((*NiuniuSrvDesk)(nil), "ddproto.niuniu_srv_desk")
	proto.RegisterType((*NiuniuSrvUser)(nil), "ddproto.niuniu_srv_user")
	proto.RegisterType((*NiuniuSrvRoom)(nil), "ddproto.niuniu_srv_room")
	proto.RegisterType((*NiuniuSrvDeskSnapshotIdIndex)(nil), "ddproto.niuniu_srv_desk_snapshot_id_index")
	proto.RegisterType((*NiuniuSrvDeskSnapshot)(nil), "ddproto.niuniu_srv_desk_snapshot")
}

func init() { proto.RegisterFile("niuniu_server.proto", fileDescriptor35) }

var fileDescriptor35 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x95, 0xcf, 0x6e, 0xe3, 0x36,
	0x10, 0xc6, 0xab, 0xb5, 0xe5, 0xb5, 0xe9, 0xf5, 0x66, 0xc3, 0xcd, 0x6e, 0x19, 0x37, 0x48, 0x55,
	0xa3, 0x07, 0x5d, 0x6a, 0xb4, 0x3e, 0xb4, 0x87, 0xdc, 0xda, 0x00, 0x45, 0x80, 0x20, 0x31, 0x98,
	0x14, 0xe9, 0x4d, 0xa0, 0x2d, 0x22, 0x21, 0x2c, 0x91, 0x82, 0x28, 0x39, 0xc9, 0xb1, 0x7d, 0x91,
	0x3e, 0x55, 0xdf, 0xa7, 0x98, 0xa1, 0x64, 0x49, 0xae, 0xf7, 0x64, 0xce, 0xc7, 0x6f, 0x86, 0x7f,
	0xf4, 0xe3, 0x98, 0x7c, 0xd4, 0xaa, 0xd4, 0xaa, 0x8c, 0xac, 0xcc, 0xb7, 0x32, 0x9f, 0x67, 0xb9,
	0x29, 0x0c, 0x7d, 0x1b, 0xc7, 0x38, 0x98, 0x9e, 0xae, 0x4d, 0x9a, 0x1a, 0x5d, 0xcd, 0x46, 0x99,
	0xd9, 0xd4, 0x9e, 0xe9, 0x71, 0x95, 0xb8, 0x12, 0x56, 0x3a, 0x69, 0xf6, 0x42, 0x3e, 0xd4, 0xd5,
	0xf2, 0xad, 0x33, 0xd3, 0x1f, 0x49, 0x3f, 0x13, 0xca, 0xb2, 0x37, 0x41, 0x2f, 0x1c, 0x2f, 0xce,
	0xe6, 0x55, 0xe5, 0x79, 0x5d, 0xb8, 0x36, 0x2e, 0x85, 0xe2, 0xe8, 0xa4, 0x0b, 0xd2, 0x2f, 0x5e,
	0x33, 0xc9, 0x7a, 0x81, 0x17, 0xbe, 0x5f, 0x9c, 0xef, 0x32, 0xaa, 0xd2, 0x52, 0x97, 0x69, 0xb4,
	0x84, 0x94, 0xfb, 0xd7, 0x4c, 0x72, 0xf4, 0xce, 0xfe, 0xf5, 0xc9, 0x51, 0x6b, 0xe9, 0x58, 0xda,
	0x0d, 0xfd, 0x4c, 0x06, 0xf0, 0x7b, 0x15, 0x33, 0x2f, 0xf0, 0x42, 0x9f, 0x57, 0x11, 0x3d, 0x27,
	0x04, 0x46, 0x37, 0x65, 0xba, 0x92, 0x39, 0x7b, 0x13, 0x78, 0xe1, 0x88, 0xb7, 0x14, 0x98, 0x7f,
	0x14, 0xa9, 0xac, 0xe6, 0x7b, 0x98, 0xdb, 0x52, 0xa0, 0x6e, 0x6e, 0x4c, 0x7a, 0x15, 0xb3, 0xbe,
	0xab, 0xeb, 0x22, 0xfa, 0x0b, 0x19, 0xd8, 0x42, 0x14, 0xa5, 0x65, 0x3e, 0xee, 0xfc, 0xdb, 0x83,
	0x3b, 0x87, 0x85, 0x22, 0xf0, 0x49, 0x5e, 0xd9, 0xe9, 0x19, 0x19, 0x25, 0xc2, 0x16, 0x0f, 0x4a,
	0xcb, 0x9c, 0x0d, 0x02, 0x2f, 0x9c, 0xf0, 0x46, 0xa0, 0x53, 0x32, 0x5c, 0xab, 0x7c, 0x9d, 0xc8,
	0x1b, 0xc3, 0x86, 0xb8, 0xe0, 0x2e, 0xa6, 0x27, 0xc4, 0x37, 0xcf, 0x90, 0x45, 0x30, 0xcb, 0x05,
	0x70, 0x80, 0x75, 0x99, 0xe7, 0xbf, 0x0a, 0xbd, 0x91, 0x39, 0x1b, 0xe3, 0x54, 0x4b, 0xa1, 0x17,
	0xee, 0x02, 0x6e, 0xb3, 0x42, 0x19, 0xcd, 0xde, 0x05, 0x5e, 0x38, 0x5e, 0x7c, 0xb3, 0xbf, 0x59,
	0xdc, 0xa7, 0x41, 0x0b, 0x6f, 0xd9, 0x29, 0x23, 0x6f, 0x95, 0xbd, 0x2b, 0x44, 0x5e, 0xb0, 0x49,
	0xe0, 0x85, 0x43, 0x5e, 0x87, 0x74, 0x46, 0xde, 0x29, 0x7b, 0xab, 0x2f, 0x95, 0xb5, 0x26, 0xd9,
	0x4a, 0xf6, 0x1e, 0xa7, 0x3b, 0x1a, 0x78, 0xe2, 0x6a, 0x7c, 0xaf, 0x52, 0xc9, 0x8e, 0x02, 0x2f,
	0xec, 0xf1, 0x8e, 0x06, 0x1e, 0xa3, 0x25, 0xd6, 0x44, 0xcf, 0x07, 0xe7, 0x69, 0x6b, 0xe0, 0x11,
	0x49, 0xd2, 0x78, 0x8e, 0x9d, 0xa7, 0xad, 0xb5, 0xd7, 0xfa, 0xc3, 0xca, 0x9c, 0x51, 0xbc, 0x88,
	0x8e, 0x06, 0x97, 0xab, 0xec, 0xa5, 0x50, 0x1b, 0xa1, 0xd8, 0x47, 0xdc, 0xef, 0x2e, 0x46, 0x4e,
	0x70, 0x84, 0xd9, 0x27, 0xee, 0x1a, 0x1b, 0x05, 0xe6, 0x95, 0xfd, 0xcd, 0x28, 0xcd, 0x8d, 0x49,
	0xd9, 0x27, 0xcc, 0x6e, 0x29, 0x34, 0x20, 0x63, 0x38, 0xfb, 0xef, 0x22, 0x4d, 0x95, 0x7e, 0x64,
	0x9f, 0xd1, 0xd0, 0x96, 0xc0, 0x61, 0xcb, 0x3c, 0x4b, 0x4a, 0x8b, 0x87, 0xf8, 0x1a, 0xbf, 0x6e,
	0x5b, 0x9a, 0xfd, 0xd3, 0xef, 0x70, 0x5d, 0x5a, 0xc7, 0x1f, 0xfc, 0x56, 0x5c, 0x4f, 0x78, 0x15,
	0xc1, 0x59, 0xb4, 0x5a, 0x6f, 0x6e, 0x44, 0x2a, 0x2b, 0xaa, 0x77, 0x31, 0xfd, 0x81, 0xf4, 0x57,
	0x2a, 0x49, 0x90, 0xae, 0xf1, 0xe2, 0x74, 0xff, 0x63, 0x43, 0x85, 0x08, 0x0c, 0x1c, 0x6d, 0xee,
	0x5a, 0x6e, 0x75, 0xa2, 0xb4, 0x44, 0xb4, 0xf0, 0x5a, 0x5c, 0x0c, 0xcc, 0x29, 0x1d, 0xcb, 0x17,
	0x04, 0xc7, 0xe7, 0x2e, 0xa0, 0x3f, 0x91, 0x01, 0x3e, 0x63, 0x8b, 0x54, 0x1c, 0x58, 0x62, 0xf7,
	0xd0, 0x79, 0x65, 0x84, 0xd3, 0xaf, 0x10, 0xc8, 0xbb, 0xb5, 0xc9, 0x1d, 0x2e, 0x3e, 0x6f, 0x4b,
	0xe0, 0x88, 0x4d, 0xb9, 0x4a, 0xa4, 0x73, 0x1c, 0x39, 0x47, 0x4b, 0x72, 0x34, 0x72, 0x29, 0xe2,
	0x57, 0xc4, 0x04, 0x69, 0xc4, 0xb0, 0x7e, 0x54, 0x2e, 0xf3, 0x18, 0x33, 0x1b, 0x81, 0x7e, 0x4f,
	0x26, 0x35, 0x07, 0x77, 0xf0, 0x16, 0x11, 0x0e, 0x9f, 0x77, 0x45, 0x70, 0xc1, 0xb1, 0x1f, 0x9e,
	0x54, 0x21, 0xaf, 0x95, 0x2d, 0x2a, 0x44, 0xba, 0x22, 0x70, 0xf6, 0x0c, 0xc1, 0x83, 0xd2, 0x1c,
	0x4a, 0x9d, 0x60, 0xa9, 0x8e, 0xb6, 0xcf, 0xc2, 0xa7, 0xff, 0xb3, 0x80, 0x27, 0xb9, 0x96, 0x62,
	0x2b, 0x2b, 0x52, 0xea, 0x10, 0x4e, 0xa2, 0xec, 0x9f, 0xa5, 0xd0, 0x99, 0x50, 0xc8, 0xc8, 0x90,
	0x37, 0xc2, 0xec, 0x2f, 0xaf, 0x43, 0x08, 0xf4, 0xa2, 0x56, 0x87, 0xf2, 0x3a, 0x1d, 0x6a, 0x4a,
	0x86, 0x30, 0x82, 0xbe, 0x89, 0x84, 0xf8, 0x7c, 0x17, 0xc3, 0x2a, 0x30, 0xbe, 0x96, 0x5b, 0x99,
	0x54, 0x4d, 0xaf, 0x11, 0xea, 0xd9, 0x7b, 0x55, 0x24, 0x12, 0xdb, 0xde, 0x88, 0x37, 0xc2, 0xec,
	0x82, 0x7c, 0xb7, 0xd7, 0x7c, 0x23, 0xab, 0x45, 0x66, 0x9f, 0x4c, 0x11, 0xa9, 0x38, 0x72, 0x84,
	0xb4, 0xdb, 0x71, 0xaf, 0x69, 0xc7, 0xb3, 0xbf, 0x3d, 0xc2, 0xbe, 0x94, 0x4d, 0x7f, 0x26, 0x23,
	0x10, 0xdc, 0x37, 0xf2, 0x90, 0x2c, 0x76, 0x88, 0x2c, 0x30, 0xf1, 0xc6, 0x4a, 0xe7, 0xc4, 0x07,
	0xa6, 0xeb, 0xbf, 0x9d, 0x83, 0x39, 0x60, 0xe0, 0xce, 0xb6, 0xfc, 0x6a, 0xe9, 0xfd, 0x17, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0x31, 0xcb, 0xa9, 0x0a, 0x07, 0x00, 0x00,
}
