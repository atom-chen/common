// Code generated by protoc-gen-go.
// source: common_server_model.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户的资
type User struct {
	RoomCard            *int64  `protobuf:"varint,1,opt,name=RoomCard" json:"RoomCard,omitempty"`
	Pwd                 *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin                *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                  *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName            *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	Scores              *int32  `protobuf:"varint,6,opt,name=scores" json:"scores,omitempty"`
	LastDrawLotteryTime *string `protobuf:"bytes,7,opt,name=lastDrawLotteryTime" json:"lastDrawLotteryTime,omitempty"`
	LastSignTime        *string `protobuf:"bytes,8,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignTotalDays       *int32  `protobuf:"varint,9,opt,name=signTotalDays" json:"signTotalDays,omitempty"`
	SignContinuousDays  *int32  `protobuf:"varint,10,opt,name=signContinuousDays" json:"signContinuousDays,omitempty"`
	Diamond             *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	Diamond2            *int64  `protobuf:"varint,12,opt,name=Diamond2" json:"Diamond2,omitempty"`
	OpenId              *string `protobuf:"bytes,13,opt,name=openId" json:"openId,omitempty"`
	UnionId             *string `protobuf:"bytes,14,opt,name=UnionId" json:"UnionId,omitempty"`
	HeadUrl             *string `protobuf:"bytes,15,opt,name=headUrl" json:"headUrl,omitempty"`
	City                *string `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Sex                 *int32  `protobuf:"varint,17,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *User) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastDrawLotteryTime() string {
	if m != nil && m.LastDrawLotteryTime != nil {
		return *m.LastDrawLotteryTime
	}
	return ""
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignTotalDays() int32 {
	if m != nil && m.SignTotalDays != nil {
		return *m.SignTotalDays
	}
	return 0
}

func (m *User) GetSignContinuousDays() int32 {
	if m != nil && m.SignContinuousDays != nil {
		return *m.SignContinuousDays
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetDiamond2() int64 {
	if m != nil && m.Diamond2 != nil {
		return *m.Diamond2
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

// 通知公告
type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Fileds           []string `protobuf:"bytes,6,rep,name=fileds" json:"fileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetFileds() []string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

// 玩家游戏场次统计
type TGameCounts struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DdzCount         *int32  `protobuf:"varint,3,opt,name=ddzCount" json:"ddzCount,omitempty"`
	DdzWinCount      *int32  `protobuf:"varint,4,opt,name=ddzWinCount" json:"ddzWinCount,omitempty"`
	MjCount          *int32  `protobuf:"varint,5,opt,name=mjCount" json:"mjCount,omitempty"`
	MjWinCount       *int32  `protobuf:"varint,6,opt,name=mjWinCount" json:"mjWinCount,omitempty"`
	ThCount          *int32  `protobuf:"varint,7,opt,name=thCount" json:"thCount,omitempty"`
	ThWinWinCount    *int32  `protobuf:"varint,8,opt,name=thWinWinCount" json:"thWinWinCount,omitempty"`
	ZjhCount         *int32  `protobuf:"varint,9,opt,name=zjhCount" json:"zjhCount,omitempty"`
	ZjhWinCount      *int32  `protobuf:"varint,10,opt,name=zjhWinCount" json:"zjhWinCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TGameCounts) Reset()                    { *m = TGameCounts{} }
func (m *TGameCounts) String() string            { return proto.CompactTextString(m) }
func (*TGameCounts) ProtoMessage()               {}
func (*TGameCounts) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *TGameCounts) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TGameCounts) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TGameCounts) GetDdzCount() int32 {
	if m != nil && m.DdzCount != nil {
		return *m.DdzCount
	}
	return 0
}

func (m *TGameCounts) GetDdzWinCount() int32 {
	if m != nil && m.DdzWinCount != nil {
		return *m.DdzWinCount
	}
	return 0
}

func (m *TGameCounts) GetMjCount() int32 {
	if m != nil && m.MjCount != nil {
		return *m.MjCount
	}
	return 0
}

func (m *TGameCounts) GetMjWinCount() int32 {
	if m != nil && m.MjWinCount != nil {
		return *m.MjWinCount
	}
	return 0
}

func (m *TGameCounts) GetThCount() int32 {
	if m != nil && m.ThCount != nil {
		return *m.ThCount
	}
	return 0
}

func (m *TGameCounts) GetThWinWinCount() int32 {
	if m != nil && m.ThWinWinCount != nil {
		return *m.ThWinWinCount
	}
	return 0
}

func (m *TGameCounts) GetZjhCount() int32 {
	if m != nil && m.ZjhCount != nil {
		return *m.ZjhCount
	}
	return 0
}

func (m *TGameCounts) GetZjhWinCount() int32 {
	if m != nil && m.ZjhWinCount != nil {
		return *m.ZjhWinCount
	}
	return 0
}

// 玩家的任务
type TUserTask struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	TaskId           *int32  `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	AwardId          *int32  `protobuf:"varint,4,opt,name=awardId" json:"awardId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TUserTask) Reset()                    { *m = TUserTask{} }
func (m *TUserTask) String() string            { return proto.CompactTextString(m) }
func (*TUserTask) ProtoMessage()               {}
func (*TUserTask) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *TUserTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TUserTask) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TUserTask) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *TUserTask) GetAwardId() int32 {
	if m != nil && m.AwardId != nil {
		return *m.AwardId
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
	proto.RegisterType((*TGameCounts)(nil), "ddproto.TGameCounts")
	proto.RegisterType((*TUserTask)(nil), "ddproto.TUserTask")
}

var fileDescriptor6 = []byte{
<<<<<<< HEAD
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd5, 0xa6, 0x69, 0x9a, 0x69, 0xbb, 0xbb, 0x64, 0x41, 0x32, 0x70, 0x41, 0x3d, 0xed,
	0x89, 0x03, 0xaf, 0xd0, 0x4a, 0x50, 0x09, 0xf6, 0x00, 0xa9, 0x38, 0xae, 0xac, 0xd8, 0x50, 0x77,
	0x13, 0x4f, 0x65, 0xbb, 0x94, 0xec, 0x8b, 0xf1, 0x02, 0x3c, 0x18, 0xe3, 0x49, 0xd3, 0x8a, 0xcb,
	0xde, 0x32, 0x9f, 0x7f, 0x67, 0xe6, 0xff, 0xc7, 0xf0, 0xba, 0xc2, 0xa6, 0x41, 0xfb, 0xe0, 0xb5,
	0xfb, 0xa5, 0xdd, 0x43, 0x83, 0x4a, 0xd7, 0xef, 0xf7, 0x0e, 0x03, 0x16, 0x99, 0x52, 0xfc, 0xb1,
	0xf8, 0x33, 0x84, 0xd1, 0x86, 0x04, 0xc5, 0x0d, 0x4c, 0xbe, 0x22, 0x36, 0x4b, 0xe9, 0x94, 0x18,
	0xbc, 0x1b, 0xdc, 0x25, 0xc5, 0x14, 0x92, 0xfd, 0x51, 0x89, 0x21, 0x15, 0x79, 0x31, 0x83, 0x51,
	0x85, 0xc6, 0x8a, 0x84, 0x8f, 0x00, 0x86, 0x46, 0x89, 0x11, 0x7d, 0xcf, 0xe3, 0x45, 0x6b, 0xaa,
	0xc7, 0x7b, 0xd9, 0x68, 0x91, 0xb2, 0xf6, 0x0a, 0xc6, 0xbe, 0x42, 0xa7, 0xbd, 0x18, 0x53, 0x9d,
	0x16, 0x6f, 0xe1, 0xb6, 0x96, 0x3e, 0xac, 0x9c, 0x3c, 0x7e, 0xc6, 0x10, 0xb4, 0x6b, 0x4b, 0x43,
	0xe2, 0x8c, 0xc5, 0x2f, 0x61, 0x16, 0x0f, 0xbf, 0x99, 0x9f, 0x96, 0xe9, 0x84, 0xe9, 0x2b, 0x98,
	0xfb, 0x48, 0x30, 0xc8, 0x7a, 0x25, 0x5b, 0x2f, 0x72, 0xfe, 0xd3, 0x1b, 0x28, 0x22, 0x5e, 0xa2,
	0x0d, 0xc6, 0x1e, 0xf0, 0xe0, 0xf9, 0x0c, 0xf8, 0xec, 0x1a, 0xb2, 0x95, 0x91, 0xe4, 0x57, 0x89,
	0x29, 0x0f, 0x49, 0x83, 0x9d, 0xc0, 0x07, 0x31, 0x63, 0x42, 0x83, 0xe1, 0x5e, 0xdb, 0xb5, 0x12,
	0x73, 0xee, 0x42, 0x57, 0x36, 0xd6, 0x60, 0x04, 0x57, 0x3d, 0xd8, 0x6a, 0xa9, 0x36, 0xae, 0x16,
	0xd7, 0x67, 0xdb, 0x26, 0xb4, 0xe2, 0x86, 0x2b, 0x4a, 0xc4, 0xeb, 0xdf, 0xe2, 0x45, 0xec, 0xb7,
	0x68, 0x21, 0x2b, 0xef, 0x31, 0x98, 0x4a, 0x9f, 0xe2, 0x18, 0xf0, 0x18, 0x54, 0x58, 0xa6, 0x65,
	0xbb, 0xd7, 0x1c, 0x5e, 0x5a, 0xdc, 0xc2, 0xf4, 0xc4, 0x4c, 0xa8, 0x35, 0x67, 0xc8, 0x16, 0x3b,
	0x18, 0xdd, 0x68, 0x1b, 0x38, 0xce, 0xfc, 0x72, 0xff, 0x8b, 0x6e, 0xf0, 0x12, 0xe8, 0x0f, 0x53,
	0x6b, 0x15, 0x03, 0x4d, 0xee, 0xf2, 0xc5, 0xdf, 0x01, 0x4c, 0xcb, 0x8f, 0x14, 0xf8, 0x12, 0x0f,
	0x36, 0xf8, 0xff, 0xfa, 0x93, 0xf6, 0x40, 0xfb, 0x5c, 0x77, 0x8b, 0xe3, 0xf5, 0x28, 0xf5, 0xc4,
	0x42, 0x6e, 0xcc, 0xd3, 0x10, 0xf9, 0x6e, 0x6c, 0x07, 0x47, 0x7d, 0x7a, 0xcd, 0xae, 0x03, 0x69,
	0xef, 0xa3, 0xd9, 0x9d, 0x45, 0xe3, 0x5e, 0x14, 0xb6, 0x1d, 0xc8, 0x18, 0x90, 0x87, 0xb0, 0x25,
	0xd1, 0x59, 0x37, 0x61, 0x4c, 0x3d, 0x9f, 0x76, 0x27, 0x61, 0xde, 0xf7, 0x24, 0x72, 0x96, 0xf1,
	0xc6, 0x16, 0x9f, 0x20, 0x2f, 0xe3, 0xdb, 0x2b, 0xa5, 0x7f, 0x7c, 0xd6, 0x03, 0xd5, 0x81, 0x34,
	0x54, 0x27, 0xfd, 0x1c, 0xf2, 0x48, 0x0f, 0x75, 0xdd, 0xbd, 0xc1, 0xf4, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x85, 0xec, 0xde, 0x19, 0xea, 0x02, 0x00, 0x00,
=======
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0x1b, 0x31,
	0x10, 0x65, 0xbd, 0xfe, 0xc8, 0x4e, 0xe2, 0x34, 0xa8, 0x60, 0xd4, 0x52, 0x8a, 0x31, 0x3d, 0xf8,
	0x54, 0x4a, 0xff, 0x82, 0x03, 0xc5, 0xd0, 0xe6, 0xb0, 0x5d, 0x93, 0x63, 0x58, 0x2c, 0x35, 0x96,
	0x6d, 0x49, 0x46, 0x52, 0xea, 0xda, 0xbf, 0xa8, 0xd0, 0x5f, 0xd0, 0x7f, 0x57, 0x66, 0x24, 0x6f,
	0xd6, 0xb4, 0x87, 0xde, 0xe6, 0xbd, 0xa7, 0x59, 0xcd, 0xbc, 0xa7, 0x85, 0x57, 0x4b, 0xab, 0xb5,
	0x35, 0x0f, 0x5e, 0xba, 0xef, 0xd2, 0x3d, 0x68, 0x2b, 0xe4, 0xf6, 0xfd, 0xce, 0xd9, 0x60, 0xd9,
	0x40, 0x08, 0x2a, 0x26, 0x3f, 0x73, 0xe8, 0x2e, 0xbc, 0x74, 0xec, 0x35, 0x5c, 0x94, 0xd6, 0xea,
	0x59, 0xed, 0x04, 0xcf, 0xc6, 0xd9, 0x34, 0x2f, 0x1b, 0xcc, 0x6e, 0x20, 0xdf, 0xed, 0x05, 0xef,
	0x8c, 0xb3, 0x69, 0x51, 0x62, 0xc9, 0x18, 0x74, 0x97, 0x56, 0x19, 0x9e, 0xd3, 0x49, 0xaa, 0xd9,
	0x35, 0x74, 0x94, 0xe0, 0xdd, 0x71, 0x36, 0x1d, 0x96, 0x1d, 0x25, 0xf0, 0x8b, 0x46, 0x2d, 0x37,
	0x77, 0xb5, 0x96, 0xbc, 0x47, 0xad, 0x0d, 0x66, 0x23, 0xe8, 0xfb, 0xa5, 0x75, 0xd2, 0xf3, 0xfe,
	0x38, 0x9b, 0xf6, 0xca, 0x84, 0xd8, 0x07, 0x78, 0xb9, 0xad, 0x7d, 0xb8, 0x75, 0xf5, 0xfe, 0xb3,
	0x0d, 0x41, 0xba, 0x43, 0xa5, 0xb4, 0xe4, 0x03, 0x6a, 0xff, 0x97, 0xc4, 0x26, 0x70, 0x85, 0xf4,
	0x57, 0xf5, 0x68, 0xe8, 0xe8, 0x05, 0x1d, 0x3d, 0xe3, 0xd8, 0x1b, 0x28, 0xbc, 0x7a, 0x34, 0x33,
	0xfb, 0x64, 0x02, 0x2f, 0xe8, 0xc2, 0x67, 0x82, 0x71, 0x18, 0xdc, 0xaa, 0x5a, 0x5b, 0x23, 0x38,
	0xd0, 0x3a, 0x27, 0x88, 0x1b, 0xa4, 0xf2, 0x23, 0xbf, 0x8c, 0x9e, 0x9c, 0x30, 0x6e, 0x60, 0x77,
	0xd2, 0xcc, 0x05, 0xbf, 0xa2, 0x1b, 0x13, 0xc2, 0xaf, 0x2d, 0x8c, 0xb2, 0x28, 0x0c, 0x49, 0x38,
	0x41, 0x54, 0x56, 0xb2, 0x16, 0x0b, 0xb7, 0xe5, 0xd7, 0x51, 0x49, 0x90, 0xdc, 0x54, 0xe1, 0xc0,
	0x5f, 0x10, 0x4d, 0x35, 0x7a, 0xee, 0xe5, 0x0f, 0x7e, 0x43, 0xd3, 0x62, 0x39, 0xf9, 0x9d, 0xc1,
	0xa0, 0xba, 0xb3, 0x41, 0x2d, 0x65, 0xf2, 0x3a, 0x23, 0x11, 0xbd, 0x7e, 0x0b, 0x60, 0x48, 0xa9,
	0x0e, 0x3b, 0x49, 0x41, 0xf5, 0xca, 0x16, 0xc3, 0xc6, 0x70, 0x99, 0x90, 0x0a, 0x5b, 0x49, 0xb1,
	0x15, 0x65, 0x9b, 0x62, 0xef, 0x60, 0x18, 0xe1, 0xcc, 0x9a, 0x20, 0x4d, 0xa0, 0x20, 0x8b, 0xf2,
	0x9c, 0x7c, 0xbe, 0xe7, 0x8b, 0xd4, 0x36, 0xa5, 0xda, 0x62, 0xd0, 0x95, 0x6f, 0x6a, 0x2b, 0x05,
	0xe6, 0x9a, 0xa3, 0x2b, 0x11, 0x4d, 0x7e, 0x75, 0xe0, 0xb2, 0xfa, 0x54, 0x6b, 0x49, 0x96, 0xfb,
	0xbf, 0xe6, 0x1f, 0x41, 0xff, 0xc9, 0x4b, 0x37, 0x8f, 0x8f, 0x6c, 0x58, 0x26, 0x84, 0x09, 0x08,
	0x71, 0x8c, 0xc1, 0xe5, 0x74, 0xba, 0xc1, 0xb8, 0x93, 0x10, 0xc7, 0x7b, 0x95, 0x72, 0xed, 0x92,
	0xdc, 0xa6, 0xd0, 0x71, 0xbd, 0x8e, 0x6a, 0x8f, 0xd4, 0x13, 0xc4, 0x3d, 0xf4, 0xba, 0x69, 0x8d,
	0x6f, 0xb0, 0xc5, 0x60, 0x67, 0x58, 0x45, 0x71, 0x10, 0x3b, 0x13, 0x44, 0x9f, 0xc2, 0xea, 0x5e,
	0x99, 0xa6, 0xf9, 0x82, 0xf4, 0x73, 0x12, 0xe7, 0x3e, 0xae, 0x57, 0xed, 0x07, 0xd7, 0x60, 0x9c,
	0xfb, 0xb8, 0x5e, 0x35, 0xfd, 0x10, 0xe7, 0x6e, 0x51, 0x13, 0x09, 0x45, 0x85, 0x3f, 0x65, 0x55,
	0xfb, 0xcd, 0x7f, 0x5b, 0x35, 0x82, 0x7e, 0xa8, 0xfd, 0x66, 0x2e, 0x92, 0x51, 0x09, 0xe1, 0x2a,
	0xf5, 0xbe, 0x76, 0x62, 0x2e, 0x92, 0x45, 0x27, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x21, 0x65,
	0x67, 0x18, 0x20, 0x04, 0x00, 0x00,
>>>>>>> 88fdf6aa358fc21d03593299d7e65f1939caba30
}
