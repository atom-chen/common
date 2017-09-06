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
func (*NiuniuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{0} }

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
	XXX_unrecognized []byte               `json:"-"`
}

func (m *NiuniuSrvDesk) Reset()                    { *m = NiuniuSrvDesk{} }
func (m *NiuniuSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvDesk) ProtoMessage()               {}
func (*NiuniuSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{1} }

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

// 用户属性
type NiuniuSrvUser struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string         `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Bill             *NiuniuUserBill `protobuf:"bytes,6,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool           `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32          `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	Pokers           *NiuniuSrvPoker `protobuf:"bytes,13,opt,name=pokers" json:"pokers,omitempty"`
	BankerScore      *int32          `protobuf:"varint,14,opt,name=bankerScore" json:"bankerScore,omitempty"`
	DoubleScore      *int32          `protobuf:"varint,15,opt,name=doubleScore" json:"doubleScore,omitempty"`
	IsReady          *bool           `protobuf:"varint,16,opt,name=isReady" json:"isReady,omitempty"`
	LastScore        *int32          `protobuf:"varint,17,opt,name=lastScore" json:"lastScore,omitempty"`
	DissolveState    *int32          `protobuf:"varint,18,opt,name=dissolveState" json:"dissolveState,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NiuniuSrvUser) Reset()                    { *m = NiuniuSrvUser{} }
func (m *NiuniuSrvUser) String() string            { return proto.CompactTextString(m) }
func (*NiuniuSrvUser) ProtoMessage()               {}
func (*NiuniuSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{2} }

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
func (*NiuniuSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{3} }

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
func (*NiuniuSrvDeskSnapshotIdIndex) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{4} }

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
func (*NiuniuSrvDeskSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{5} }

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

func init() { proto.RegisterFile("niuniu_server.proto", fileDescriptor33) }

var fileDescriptor33 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xc9, 0xda, 0x74, 0xad, 0xbb, 0xee, 0x87, 0x37, 0x21, 0xaf, 0x4c, 0x23, 0x54, 0x1c,
	0x72, 0xa1, 0x82, 0x1e, 0xe0, 0xb0, 0x1b, 0xda, 0x65, 0x12, 0xda, 0x2a, 0x6f, 0x88, 0x63, 0x94,
	0x36, 0x16, 0x58, 0x4d, 0xec, 0xca, 0x4e, 0xca, 0x76, 0x84, 0xff, 0x98, 0x33, 0x17, 0xf4, 0x9e,
	0x93, 0x26, 0x99, 0xca, 0x29, 0x7e, 0x5f, 0x7f, 0xde, 0xf3, 0x8b, 0xdf, 0x7b, 0x26, 0xa7, 0x4a,
	0x16, 0x4a, 0x16, 0x91, 0x15, 0x66, 0x23, 0xcc, 0x74, 0x6d, 0x74, 0xae, 0xe9, 0x7e, 0x92, 0xe0,
	0x62, 0x7c, 0xbe, 0xd4, 0x59, 0xa6, 0x55, 0xb9, 0x1b, 0xad, 0xf5, 0xaa, 0x62, 0xc6, 0x27, 0xa5,
	0xe3, 0x22, 0xb6, 0xc2, 0x49, 0x93, 0x47, 0x72, 0x5c, 0x45, 0x33, 0x1b, 0x07, 0xd3, 0xf7, 0xa4,
	0xbb, 0x8e, 0xa5, 0x65, 0x7b, 0x41, 0x27, 0x1c, 0xce, 0x2e, 0xa6, 0x65, 0xe4, 0x69, 0x15, 0xb8,
	0x02, 0xe7, 0xb1, 0xe4, 0x48, 0xd2, 0x19, 0xe9, 0xe6, 0x4f, 0x6b, 0xc1, 0x3a, 0x81, 0x17, 0x1e,
	0xce, 0x2e, 0xb7, 0x1e, 0x65, 0x68, 0xa1, 0x8a, 0x2c, 0x9a, 0x83, 0xcb, 0xc3, 0xd3, 0x5a, 0x70,
	0x64, 0x27, 0x7f, 0xba, 0xe4, 0xa8, 0x71, 0x74, 0x22, 0xec, 0x8a, 0xbe, 0x24, 0x3d, 0xf8, 0xde,
	0x24, 0xcc, 0x0b, 0xbc, 0xd0, 0xe7, 0xa5, 0x45, 0x2f, 0x09, 0x81, 0xd5, 0x6d, 0x91, 0x2d, 0x84,
	0x61, 0x7b, 0x81, 0x17, 0x0e, 0x78, 0x43, 0x81, 0xfd, 0xef, 0x71, 0x26, 0xca, 0xfd, 0x0e, 0xfa,
	0x36, 0x14, 0x88, 0x6b, 0xb4, 0xce, 0x6e, 0x12, 0xd6, 0x75, 0x71, 0x9d, 0x45, 0x3f, 0x91, 0x9e,
	0xcd, 0xe3, 0xbc, 0xb0, 0xcc, 0xc7, 0xcc, 0x5f, 0xef, 0xcc, 0x1c, 0x0e, 0x8a, 0x80, 0x13, 0xbc,
	0xc4, 0xe9, 0x05, 0x19, 0xa4, 0xb1, 0xcd, 0xbf, 0x49, 0x25, 0x0c, 0xeb, 0x05, 0x5e, 0x38, 0xe2,
	0xb5, 0x40, 0xc7, 0xa4, 0xbf, 0x94, 0x66, 0x99, 0x8a, 0x5b, 0xcd, 0xfa, 0x78, 0xe0, 0xd6, 0xa6,
	0x67, 0xc4, 0xd7, 0x3f, 0xc1, 0x8b, 0xa0, 0x97, 0x33, 0xe0, 0x07, 0x96, 0x85, 0x31, 0x9f, 0x63,
	0xb5, 0x12, 0x86, 0x0d, 0x71, 0xab, 0xa1, 0xd0, 0x2b, 0x77, 0x01, 0x77, 0xeb, 0x5c, 0x6a, 0xc5,
	0x0e, 0x02, 0x2f, 0x1c, 0xce, 0x5e, 0x3d, 0x4f, 0x16, 0xf3, 0xd4, 0x88, 0xf0, 0x06, 0x4e, 0x19,
	0xd9, 0x97, 0xf6, 0x3e, 0x8f, 0x4d, 0xce, 0x46, 0x81, 0x17, 0xf6, 0x79, 0x65, 0xd2, 0x09, 0x39,
	0x90, 0xf6, 0x4e, 0x5d, 0x4b, 0x6b, 0x75, 0xba, 0x11, 0xec, 0x10, 0xb7, 0x5b, 0x1a, 0x30, 0x49,
	0xb9, 0x7e, 0x90, 0x99, 0x60, 0x47, 0x81, 0x17, 0x76, 0x78, 0x4b, 0x03, 0x46, 0x2b, 0x81, 0x31,
	0x91, 0x39, 0x76, 0x4c, 0x53, 0x03, 0x26, 0x4e, 0xd3, 0x9a, 0x39, 0x71, 0x4c, 0x53, 0x6b, 0x9e,
	0xf5, 0xd5, 0x0a, 0xc3, 0x28, 0x5e, 0x44, 0x4b, 0x83, 0xcb, 0x95, 0xf6, 0x3a, 0x96, 0xab, 0x58,
	0xb2, 0x53, 0xcc, 0x77, 0x6b, 0x63, 0x9f, 0xe0, 0x0a, 0xbd, 0xcf, 0xdc, 0x35, 0xd6, 0xca, 0xe4,
	0xef, 0x5e, 0xab, 0xe7, 0x0a, 0xeb, 0x7a, 0x03, 0xbe, 0x65, 0xcf, 0x8d, 0x78, 0x69, 0xc1, 0x39,
	0x4a, 0x2e, 0x57, 0xb7, 0x71, 0x26, 0xca, 0x8e, 0xdb, 0xda, 0xf4, 0x1d, 0xe9, 0x2e, 0x64, 0x9a,
	0x62, 0xe5, 0x87, 0xb3, 0xf3, 0xe7, 0x85, 0x80, 0x08, 0x11, 0x00, 0x1c, 0x31, 0x97, 0xf2, 0x9d,
	0x4a, 0xa5, 0x12, 0x58, 0x76, 0x4c, 0xd9, 0xd9, 0xd0, 0x0f, 0x52, 0x25, 0xe2, 0x11, 0x8b, 0xea,
	0x73, 0x67, 0xd0, 0x0f, 0xa4, 0x87, 0x23, 0x66, 0xb1, 0x62, 0x3b, 0x8e, 0xd8, 0x0e, 0x21, 0x2f,
	0x41, 0x1a, 0x90, 0xe1, 0x02, 0x9b, 0xe5, 0x7e, 0xa9, 0x8d, 0x2b, 0xa5, 0xcf, 0x9b, 0x12, 0x10,
	0x89, 0x2e, 0x16, 0xa9, 0x70, 0xc4, 0x91, 0x23, 0x1a, 0x92, 0xeb, 0x14, 0x2e, 0xe2, 0xe4, 0x09,
	0x4b, 0x88, 0x9d, 0x82, 0x66, 0xd5, 0xf0, 0xce, 0xf3, 0x04, 0x3d, 0x6b, 0x81, 0xbe, 0x25, 0xa3,
	0xaa, 0x46, 0xf7, 0x30, 0x27, 0x58, 0x38, 0x9f, 0xb7, 0xc5, 0xc9, 0x2f, 0xaf, 0x75, 0xfb, 0x30,
	0x83, 0x8d, 0xc9, 0xf4, 0x5a, 0x93, 0x39, 0x26, 0x7d, 0x58, 0xc1, 0x7b, 0x81, 0xb7, 0xef, 0xf3,
	0xad, 0x0d, 0xb9, 0xc0, 0xfa, 0x8b, 0xd8, 0x88, 0xb4, 0x1c, 0xf6, 0x5a, 0xa8, 0x76, 0x1f, 0x64,
	0x9e, 0x0a, 0x1c, 0xf7, 0x01, 0xaf, 0x85, 0xc9, 0x15, 0x79, 0xf3, 0xec, 0xd1, 0x89, 0xac, 0x8a,
	0xd7, 0xf6, 0x87, 0xce, 0x23, 0x99, 0x44, 0xee, 0xf6, 0x9b, 0xcf, 0x50, 0xa7, 0x7e, 0x86, 0x26,
	0xbf, 0x3d, 0xc2, 0xfe, 0xe7, 0x4d, 0x3f, 0x92, 0x01, 0x08, 0xee, 0xff, 0x3d, 0xac, 0x1a, 0xdb,
	0x55, 0x35, 0x80, 0x78, 0x8d, 0xd2, 0x29, 0xf1, 0xa1, 0x5f, 0xaa, 0xe7, 0x76, 0xa7, 0x0f, 0x00,
	0xdc, 0x61, 0xf3, 0x17, 0x73, 0xef, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xef, 0xa6, 0x3c,
	0x02, 0x06, 0x00, 0x00,
}
