// Code generated by protoc-gen-go.
// source: pdk_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of pdk_base_roomTypeInfo from pdk_base.proto

// Ignoring public import of pdk_base_playerInfo from pdk_base.proto

// Ignoring public import of pdk_base_playerRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_commonRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_timerInfo from pdk_base.proto

// Ignoring public import of pdk_base_deskInfo from pdk_base.proto

// Ignoring public import of pdk_enum_protoId from pdk_base.proto

// Ignoring public import of pdk_enum_errorCode from pdk_base.proto

// Ignoring public import of pdk_enum_paiType from pdk_base.proto

// Ignoring public import of pdk_enum_actType from pdk_base.proto

// Ignoring public import of pdk_enum_gameStatus from pdk_base.proto

// Ignoring public import of pdk_enum_playerStatus from pdk_base.proto

// Ignoring public import of pdk_enum_roomType from pdk_base.proto

// Ignoring public import of pdk_enum_enterType from pdk_base.proto

// Ignoring public import of pdk_enum_coinRoomLevel from pdk_base.proto

// Ignoring public import of pdk_enum_deskGameStatus from pdk_base.proto

// 打出去的牌
type PdkSrvOutPokerPais struct {
	KeyValue         *int32               `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32               `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool                `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32               `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32               `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32               `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32               `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PdkSrvOutPokerPais) Reset()                    { *m = PdkSrvOutPokerPais{} }
func (m *PdkSrvOutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvOutPokerPais) ProtoMessage()               {}
func (*PdkSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *PdkSrvOutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetPokerPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *PdkSrvOutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *PdkSrvOutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *PdkSrvOutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 对 desk的统计
type PdkSrvDeskTongJi struct {
	Bombs            []*PdkSrvOutPokerPais `protobuf:"bytes,1,rep,name=bombs" json:"bombs,omitempty"`
	CountQiangDiZhu  *int32                `protobuf:"varint,2,opt,name=countQiangDiZhu" json:"countQiangDiZhu,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkSrvDeskTongJi) Reset()                    { *m = PdkSrvDeskTongJi{} }
func (m *PdkSrvDeskTongJi) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvDeskTongJi) ProtoMessage()               {}
func (*PdkSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

func (m *PdkSrvDeskTongJi) GetBombs() []*PdkSrvOutPokerPais {
	if m != nil {
		return m.Bombs
	}
	return nil
}

func (m *PdkSrvDeskTongJi) GetCountQiangDiZhu() int32 {
	if m != nil && m.CountQiangDiZhu != nil {
		return *m.CountQiangDiZhu
	}
	return 0
}

// desk
type PdkSrvDesk struct {
	AllPokerPai      []*CommonSrvPokerPai   `protobuf:"bytes,1,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*CommonSrvPokerPai   `protobuf:"bytes,2,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *PdkSrvOutPokerPais    `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	DizhuPaiUser     *uint32                `protobuf:"varint,4,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	DiZhuUserId      *uint32                `protobuf:"varint,5,opt,name=diZhuUserId" json:"diZhuUserId,omitempty"`
	ActiveUserId     *uint32                `protobuf:"varint,6,opt,name=activeUserId" json:"activeUserId,omitempty"`
	Tongji           *PdkSrvDeskTongJi      `protobuf:"bytes,7,opt,name=tongji" json:"tongji,omitempty"`
	PdkType          *int32                 `protobuf:"varint,8,opt,name=pdkType" json:"pdkType,omitempty"`
	RoomType         *int32                 `protobuf:"varint,9,opt,name=RoomType" json:"RoomType,omitempty"`
	BoardsCount      *int32                 `protobuf:"varint,10,opt,name=BoardsCount" json:"BoardsCount,omitempty"`
	CapMax           *int64                 `protobuf:"varint,11,opt,name=CapMax" json:"CapMax,omitempty"`
	IsJiaoFen        *bool                  `protobuf:"varint,12,opt,name=IsJiaoFen" json:"IsJiaoFen,omitempty"`
	CommonRateInfo   *PdkBaseCommonRateInfo `protobuf:"bytes,13,opt,name=commonRateInfo" json:"commonRateInfo,omitempty"`
	PlayRate         *int32                 `protobuf:"varint,14,opt,name=playRate" json:"playRate,omitempty"`
	GameStatus       *int32                 `protobuf:"varint,15,opt,name=GameStatus" json:"GameStatus,omitempty"`
	InitRoomCoin     *int64                 `protobuf:"varint,16,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32                 `protobuf:"varint,17,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32                 `protobuf:"varint,18,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	PlayerNum        *int32                 `protobuf:"varint,19,opt,name=playerNum" json:"playerNum,omitempty"`
	TimeOutSecond    *int32                 `protobuf:"varint,20,opt,name=timeOutSecond" json:"timeOutSecond,omitempty"`
	UserMinCoin      *int64                 `protobuf:"varint,21,opt,name=userMinCoin" json:"userMinCoin,omitempty"`
	UserMaxCoin      *int64                 `protobuf:"varint,22,opt,name=userMaxCoin" json:"userMaxCoin,omitempty"`
	CoinTicket       *int64                 `protobuf:"varint,23,opt,name=coinTicket" json:"coinTicket,omitempty"`
	CoinRoomLv       *PdkEnumCoinRoomLevel  `protobuf:"varint,24,opt,name=coinRoomLv,enum=ddproto.PdkEnumCoinRoomLevel" json:"coinRoomLv,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *PdkSrvDesk) Reset()                    { *m = PdkSrvDesk{} }
func (m *PdkSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvDesk) ProtoMessage()               {}
func (*PdkSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

func (m *PdkSrvDesk) GetAllPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *PdkSrvDesk) GetDiPokerPai() []*CommonSrvPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *PdkSrvDesk) GetOutPai() *PdkSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *PdkSrvDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *PdkSrvDesk) GetDiZhuUserId() uint32 {
	if m != nil && m.DiZhuUserId != nil {
		return *m.DiZhuUserId
	}
	return 0
}

func (m *PdkSrvDesk) GetActiveUserId() uint32 {
	if m != nil && m.ActiveUserId != nil {
		return *m.ActiveUserId
	}
	return 0
}

func (m *PdkSrvDesk) GetTongji() *PdkSrvDeskTongJi {
	if m != nil {
		return m.Tongji
	}
	return nil
}

func (m *PdkSrvDesk) GetPdkType() int32 {
	if m != nil && m.PdkType != nil {
		return *m.PdkType
	}
	return 0
}

func (m *PdkSrvDesk) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *PdkSrvDesk) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *PdkSrvDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *PdkSrvDesk) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *PdkSrvDesk) GetCommonRateInfo() *PdkBaseCommonRateInfo {
	if m != nil {
		return m.CommonRateInfo
	}
	return nil
}

func (m *PdkSrvDesk) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *PdkSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PdkSrvDesk) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *PdkSrvDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *PdkSrvDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *PdkSrvDesk) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *PdkSrvDesk) GetTimeOutSecond() int32 {
	if m != nil && m.TimeOutSecond != nil {
		return *m.TimeOutSecond
	}
	return 0
}

func (m *PdkSrvDesk) GetUserMinCoin() int64 {
	if m != nil && m.UserMinCoin != nil {
		return *m.UserMinCoin
	}
	return 0
}

func (m *PdkSrvDesk) GetUserMaxCoin() int64 {
	if m != nil && m.UserMaxCoin != nil {
		return *m.UserMaxCoin
	}
	return 0
}

func (m *PdkSrvDesk) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *PdkSrvDesk) GetCoinRoomLv() PdkEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLv != nil {
		return *m.CoinRoomLv
	}
	return PdkEnumCoinRoomLevel_PDK_LV1
}

// 游戏数据
type PdkSrvGameData struct {
	HandPokers       []*CommonSrvPokerPai  `protobuf:"bytes,1,rep,name=handPokers" json:"handPokers,omitempty"`
	OutPaiList       []*PdkSrvOutPokerPais `protobuf:"bytes,2,rep,name=outPaiList" json:"outPaiList,omitempty"`
	OutPai           *PdkSrvOutPokerPais   `protobuf:"bytes,3,opt,name=outPai" json:"outPai,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkSrvGameData) Reset()                    { *m = PdkSrvGameData{} }
func (m *PdkSrvGameData) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvGameData) ProtoMessage()               {}
func (*PdkSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

func (m *PdkSrvGameData) GetHandPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PdkSrvGameData) GetOutPaiList() []*PdkSrvOutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *PdkSrvGameData) GetOutPai() *PdkSrvOutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

type PdkSrvBillBean struct {
	// 斗地主的账单
	Coin             *int64  `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32 `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32 `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkSrvBillBean) Reset()                    { *m = PdkSrvBillBean{} }
func (m *PdkSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvBillBean) ProtoMessage()               {}
func (*PdkSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

func (m *PdkSrvBillBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkSrvBillBean) GetLoseUser() uint32 {
	if m != nil && m.LoseUser != nil {
		return *m.LoseUser
	}
	return 0
}

func (m *PdkSrvBillBean) GetWinUser() uint32 {
	if m != nil && m.WinUser != nil {
		return *m.WinUser
	}
	return 0
}

func (m *PdkSrvBillBean) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

type PdkSrvBill struct {
	// 斗地主的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*PdkSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkSrvBill) Reset()                    { *m = PdkSrvBill{} }
func (m *PdkSrvBill) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvBill) ProtoMessage()               {}
func (*PdkSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

func (m *PdkSrvBill) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkSrvBill) GetBillBean() []*PdkSrvBillBean {
	if m != nil {
		return m.BillBean
	}
	return nil
}

type PdkSrvUserStatisticsRound struct {
	// 玩家每一局的统计信息
	Round            *int32                `protobuf:"varint,1,opt,name=round" json:"round,omitempty"`
	WinCoin          *int64                `protobuf:"varint,2,opt,name=winCoin" json:"winCoin,omitempty"`
	CountBomb        *int32                `protobuf:"varint,3,opt,name=countBomb" json:"countBomb,omitempty"`
	BomBean          []*PdkSrvOutPokerPais `protobuf:"bytes,4,rep,name=bomBean" json:"bomBean,omitempty"`
	PlayRate         *int32                `protobuf:"varint,5,opt,name=playRate" json:"playRate,omitempty"`
	IsSpring         *bool                 `protobuf:"varint,6,opt,name=isSpring" json:"isSpring,omitempty"`
	IsDizhu          *bool                 `protobuf:"varint,7,opt,name=isDizhu" json:"isDizhu,omitempty"`
	IsWin            *bool                 `protobuf:"varint,8,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PdkSrvUserStatisticsRound) Reset()                    { *m = PdkSrvUserStatisticsRound{} }
func (m *PdkSrvUserStatisticsRound) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvUserStatisticsRound) ProtoMessage()               {}
func (*PdkSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

func (m *PdkSrvUserStatisticsRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *PdkSrvUserStatisticsRound) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *PdkSrvUserStatisticsRound) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *PdkSrvUserStatisticsRound) GetBomBean() []*PdkSrvOutPokerPais {
	if m != nil {
		return m.BomBean
	}
	return nil
}

func (m *PdkSrvUserStatisticsRound) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *PdkSrvUserStatisticsRound) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *PdkSrvUserStatisticsRound) GetIsDizhu() bool {
	if m != nil && m.IsDizhu != nil {
		return *m.IsDizhu
	}
	return false
}

func (m *PdkSrvUserStatisticsRound) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type PdkSrvUserStatistics struct {
	// 玩家统计信息
	RoundBean        []*PdkSrvUserStatisticsRound `protobuf:"bytes,1,rep,name=roundBean" json:"roundBean,omitempty"`
	CountWin         *int32                       `protobuf:"varint,2,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32                       `protobuf:"varint,3,opt,name=countLose" json:"countLose,omitempty"`
	CountSpring      *int32                       `protobuf:"varint,4,opt,name=countSpring" json:"countSpring,omitempty"`
	CountDizhu       *int32                       `protobuf:"varint,5,opt,name=countDizhu" json:"countDizhu,omitempty"`
	CountBomb        *int32                       `protobuf:"varint,6,opt,name=countBomb" json:"countBomb,omitempty"`
	MaxWinCoin       *int64                       `protobuf:"varint,7,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *PdkSrvUserStatistics) Reset()                    { *m = PdkSrvUserStatistics{} }
func (m *PdkSrvUserStatistics) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvUserStatistics) ProtoMessage()               {}
func (*PdkSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

func (m *PdkSrvUserStatistics) GetRoundBean() []*PdkSrvUserStatisticsRound {
	if m != nil {
		return m.RoundBean
	}
	return nil
}

func (m *PdkSrvUserStatistics) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *PdkSrvUserStatistics) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

func (m *PdkSrvUserStatistics) GetCountSpring() int32 {
	if m != nil && m.CountSpring != nil {
		return *m.CountSpring
	}
	return 0
}

func (m *PdkSrvUserStatistics) GetCountDizhu() int32 {
	if m != nil && m.CountDizhu != nil {
		return *m.CountDizhu
	}
	return 0
}

func (m *PdkSrvUserStatistics) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *PdkSrvUserStatistics) GetMaxWinCoin() int64 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

// user
type PdkSrvUser struct {
	UserId              *uint32               `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData            *PdkSrvGameData       `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	StatusHLQiang       *int32                `protobuf:"varint,3,opt,name=statusHLQiang" json:"statusHLQiang,omitempty"`
	StatusHLJiao        *int32                `protobuf:"varint,4,opt,name=statusHLJiao" json:"statusHLJiao,omitempty"`
	StatusHLJiaBei      *int32                `protobuf:"varint,5,opt,name=statusHLJiaBei" json:"statusHLJiaBei,omitempty"`
	StatusSCMenZhua     *int32                `protobuf:"varint,6,opt,name=statusSCMenZhua" json:"statusSCMenZhua,omitempty"`
	StatusSCZhuaPai     *int32                `protobuf:"varint,7,opt,name=statusSCZhuaPai" json:"statusSCZhuaPai,omitempty"`
	StatusSCLaDao       *int32                `protobuf:"varint,8,opt,name=statusSCLaDao" json:"statusSCLaDao,omitempty"`
	StatusJDJiao        *int32                `protobuf:"varint,9,opt,name=statusJDJiao" json:"statusJDJiao,omitempty"`
	StatusShowPokers    *int32                `protobuf:"varint,10,opt,name=statusShowPokers" json:"statusShowPokers,omitempty"`
	IsShowPokers        *bool                 `protobuf:"varint,11,opt,name=isShowPokers" json:"isShowPokers,omitempty"`
	Bill                *PdkSrvBill           `protobuf:"bytes,12,opt,name=bill" json:"bill,omitempty"`
	Coin                *int64                `protobuf:"varint,13,opt,name=coin" json:"coin,omitempty"`
	Score               *int64                `protobuf:"varint,14,opt,name=score" json:"score,omitempty"`
	Statistics          *PdkSrvUserStatistics `protobuf:"bytes,15,opt,name=statistics" json:"statistics,omitempty"`
	PlayRate            *int32                `protobuf:"varint,16,opt,name=playRate" json:"playRate,omitempty"`
	JiaoScore           *int32                `protobuf:"varint,17,opt,name=jiaoScore" json:"jiaoScore,omitempty"`
	TimeOutCount        *int32                `protobuf:"varint,18,opt,name=timeOutCount" json:"timeOutCount,omitempty"`
	IsAgent             *bool                 `protobuf:"varint,19,opt,name=isAgent" json:"isAgent,omitempty"`
	Sex                 *int32                `protobuf:"varint,20,opt,name=sex" json:"sex,omitempty"`
	RoomCard            *int32                `protobuf:"varint,21,opt,name=roomCard" json:"roomCard,omitempty"`
	StatusApplyDissolve *int32                `protobuf:"varint,22,opt,name=statusApplyDissolve" json:"statusApplyDissolve,omitempty"`
	ScRate              *int32                `protobuf:"varint,23,opt,name=scRate" json:"scRate,omitempty"`
	JbRate              *int32                `protobuf:"varint,24,opt,name=jbRate" json:"jbRate,omitempty"`
	XXX_unrecognized    []byte                `json:"-"`
}

func (m *PdkSrvUser) Reset()                    { *m = PdkSrvUser{} }
func (m *PdkSrvUser) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvUser) ProtoMessage()               {}
func (*PdkSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

func (m *PdkSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkSrvUser) GetGameData() *PdkSrvGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *PdkSrvUser) GetStatusHLQiang() int32 {
	if m != nil && m.StatusHLQiang != nil {
		return *m.StatusHLQiang
	}
	return 0
}

func (m *PdkSrvUser) GetStatusHLJiao() int32 {
	if m != nil && m.StatusHLJiao != nil {
		return *m.StatusHLJiao
	}
	return 0
}

func (m *PdkSrvUser) GetStatusHLJiaBei() int32 {
	if m != nil && m.StatusHLJiaBei != nil {
		return *m.StatusHLJiaBei
	}
	return 0
}

func (m *PdkSrvUser) GetStatusSCMenZhua() int32 {
	if m != nil && m.StatusSCMenZhua != nil {
		return *m.StatusSCMenZhua
	}
	return 0
}

func (m *PdkSrvUser) GetStatusSCZhuaPai() int32 {
	if m != nil && m.StatusSCZhuaPai != nil {
		return *m.StatusSCZhuaPai
	}
	return 0
}

func (m *PdkSrvUser) GetStatusSCLaDao() int32 {
	if m != nil && m.StatusSCLaDao != nil {
		return *m.StatusSCLaDao
	}
	return 0
}

func (m *PdkSrvUser) GetStatusJDJiao() int32 {
	if m != nil && m.StatusJDJiao != nil {
		return *m.StatusJDJiao
	}
	return 0
}

func (m *PdkSrvUser) GetStatusShowPokers() int32 {
	if m != nil && m.StatusShowPokers != nil {
		return *m.StatusShowPokers
	}
	return 0
}

func (m *PdkSrvUser) GetIsShowPokers() bool {
	if m != nil && m.IsShowPokers != nil {
		return *m.IsShowPokers
	}
	return false
}

func (m *PdkSrvUser) GetBill() *PdkSrvBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *PdkSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PdkSrvUser) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *PdkSrvUser) GetStatistics() *PdkSrvUserStatistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *PdkSrvUser) GetPlayRate() int32 {
	if m != nil && m.PlayRate != nil {
		return *m.PlayRate
	}
	return 0
}

func (m *PdkSrvUser) GetJiaoScore() int32 {
	if m != nil && m.JiaoScore != nil {
		return *m.JiaoScore
	}
	return 0
}

func (m *PdkSrvUser) GetTimeOutCount() int32 {
	if m != nil && m.TimeOutCount != nil {
		return *m.TimeOutCount
	}
	return 0
}

func (m *PdkSrvUser) GetIsAgent() bool {
	if m != nil && m.IsAgent != nil {
		return *m.IsAgent
	}
	return false
}

func (m *PdkSrvUser) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PdkSrvUser) GetRoomCard() int32 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *PdkSrvUser) GetStatusApplyDissolve() int32 {
	if m != nil && m.StatusApplyDissolve != nil {
		return *m.StatusApplyDissolve
	}
	return 0
}

func (m *PdkSrvUser) GetScRate() int32 {
	if m != nil && m.ScRate != nil {
		return *m.ScRate
	}
	return 0
}

func (m *PdkSrvUser) GetJbRate() int32 {
	if m != nil && m.JbRate != nil {
		return *m.JbRate
	}
	return 0
}

// room
type PdkSrvRoom struct {
	RoomId           *int32 `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PdkSrvRoom) Reset()                    { *m = PdkSrvRoom{} }
func (m *PdkSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvRoom) ProtoMessage()               {}
func (*PdkSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

func (m *PdkSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

// 备份专用...
type PdkSrvBak struct {
	Desk             *PdkSrvDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*PdkSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PdkSrvBak) Reset()                    { *m = PdkSrvBak{} }
func (m *PdkSrvBak) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvBak) ProtoMessage()               {}
func (*PdkSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

func (m *PdkSrvBak) GetDesk() *PdkSrvDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *PdkSrvBak) GetUsers() []*PdkSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*PdkSrvOutPokerPais)(nil), "ddproto.pdk_srv_outPokerPais")
	proto.RegisterType((*PdkSrvDeskTongJi)(nil), "ddproto.pdk_srv_deskTongJi")
	proto.RegisterType((*PdkSrvDesk)(nil), "ddproto.pdk_srv_desk")
	proto.RegisterType((*PdkSrvGameData)(nil), "ddproto.pdk_srv_gameData")
	proto.RegisterType((*PdkSrvBillBean)(nil), "ddproto.pdk_srv_billBean")
	proto.RegisterType((*PdkSrvBill)(nil), "ddproto.pdk_srv_bill")
	proto.RegisterType((*PdkSrvUserStatisticsRound)(nil), "ddproto.pdk_srv_userStatisticsRound")
	proto.RegisterType((*PdkSrvUserStatistics)(nil), "ddproto.pdk_srv_userStatistics")
	proto.RegisterType((*PdkSrvUser)(nil), "ddproto.pdk_srv_user")
	proto.RegisterType((*PdkSrvRoom)(nil), "ddproto.pdk_srv_room")
	proto.RegisterType((*PdkSrvBak)(nil), "ddproto.pdk_srv_bak")
}

var fileDescriptor25 = []byte{
	// 1313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x5d, 0x6f, 0xdb, 0x36,
	0x17, 0x7e, 0x1d, 0xc7, 0x5f, 0x74, 0x93, 0xe6, 0x65, 0xd3, 0x56, 0x4d, 0xb3, 0xce, 0x10, 0x8a,
	0xc2, 0xdb, 0x80, 0x62, 0x48, 0x51, 0x0c, 0x18, 0xf6, 0x81, 0x26, 0xc6, 0xd6, 0x14, 0xe9, 0x96,
	0xd1, 0xed, 0x8a, 0xf5, 0xc6, 0xa0, 0x2d, 0xce, 0x66, 0x2d, 0x89, 0x82, 0x28, 0xb9, 0x49, 0x7f,
	0xc3, 0xae, 0xf7, 0x7b, 0x76, 0xb5, 0xfb, 0x5d, 0xef, 0xcf, 0x0c, 0xe7, 0x90, 0x92, 0x28, 0xc7,
	0x58, 0x8b, 0x5d, 0xc5, 0xe7, 0xe1, 0x43, 0xf1, 0x7c, 0xf0, 0x39, 0x87, 0x21, 0x7b, 0x49, 0xb0,
	0x9c, 0x68, 0x91, 0xae, 0x44, 0xfa, 0x30, 0x49, 0x55, 0xa6, 0x68, 0x27, 0x08, 0xf0, 0xc7, 0xc1,
	0x9d, 0x99, 0x8a, 0x22, 0x15, 0xdb, 0xd5, 0x49, 0xa2, 0x96, 0x05, 0xe7, 0x60, 0x17, 0x76, 0x4d,
	0xb9, 0x16, 0xc6, 0xf6, 0xff, 0xd8, 0x22, 0xfb, 0xf8, 0xa1, 0x74, 0x35, 0x51, 0x79, 0x76, 0x0e,
	0xd4, 0x73, 0x2e, 0x35, 0x3d, 0x20, 0xdd, 0xa5, 0xb8, 0xfc, 0x99, 0x87, 0xb9, 0xf0, 0x1a, 0x83,
	0xc6, 0xb0, 0xc5, 0x4a, 0x9b, 0x7e, 0x49, 0x7a, 0x49, 0x41, 0xf4, 0xb6, 0x06, 0xcd, 0x61, 0xff,
	0xe8, 0xf0, 0xa1, 0x3d, 0xfc, 0x61, 0x71, 0x76, 0xba, 0x9a, 0x14, 0x24, 0x56, 0xd1, 0x29, 0x25,
	0xdb, 0xd9, 0x65, 0x22, 0xbc, 0x26, 0x7e, 0x13, 0x7f, 0xd3, 0x5b, 0xa4, 0x2d, 0xf5, 0xb1, 0x8a,
	0xa6, 0xde, 0xf6, 0xa0, 0x31, 0xec, 0x32, 0x6b, 0xd1, 0x7b, 0x84, 0xcc, 0x54, 0x1e, 0x67, 0xa3,
	0x5c, 0xbe, 0x93, 0x5e, 0x0b, 0x77, 0x38, 0x08, 0xbd, 0x4f, 0x76, 0xd0, 0x1a, 0xf3, 0xf8, 0xdd,
	0x82, 0xc7, 0x73, 0xaf, 0x8d, 0x94, 0x3a, 0x48, 0x7d, 0x72, 0xcd, 0x00, 0xd2, 0x90, 0x3a, 0x48,
	0xaa, 0x61, 0x25, 0xe7, 0x17, 0xcb, 0xe9, 0x3a, 0x1c, 0x8b, 0x81, 0x97, 0xb9, 0x16, 0xe9, 0x69,
	0xe0, 0xf5, 0x06, 0x8d, 0xe1, 0x0e, 0xb3, 0x96, 0xaf, 0x09, 0x2d, 0x32, 0x18, 0x08, 0xbd, 0x7c,
	0xa1, 0xe2, 0xf9, 0x33, 0x49, 0x1f, 0x91, 0xd6, 0x54, 0x45, 0x53, 0xed, 0x35, 0x30, 0x3f, 0x1f,
	0x95, 0xf9, 0xd9, 0x94, 0x6d, 0x66, 0xb8, 0x74, 0x48, 0xae, 0xe3, 0x91, 0x3f, 0x49, 0x1e, 0xcf,
	0x47, 0xf2, 0xf5, 0x22, 0xf7, 0xb6, 0xd0, 0x93, 0x75, 0xd8, 0xff, 0xab, 0x43, 0xae, 0xb9, 0xa7,
	0xd2, 0x6f, 0x48, 0x9f, 0x87, 0x61, 0xf1, 0x45, 0x7b, 0xea, 0xbf, 0x57, 0xc5, 0xdd, 0x40, 0xbf,
	0x22, 0x24, 0x90, 0xe5, 0xf6, 0x0f, 0x29, 0xaa, 0xc3, 0xa7, 0x8f, 0x49, 0x1b, 0xe2, 0xe1, 0x12,
	0xeb, 0xfa, 0xde, 0x70, 0x2d, 0x19, 0xd2, 0x1e, 0xc8, 0x77, 0x8b, 0xfc, 0x9c, 0xcb, 0x97, 0x5a,
	0xa4, 0x58, 0xfe, 0x1d, 0x56, 0xc3, 0xe8, 0x80, 0xf4, 0x03, 0x08, 0xf9, 0xa5, 0xc9, 0x7d, 0x0b,
	0x29, 0x2e, 0x04, 0x5f, 0xe1, 0xb3, 0x4c, 0xae, 0x84, 0xa5, 0xb4, 0xcd, 0x57, 0x5c, 0x8c, 0x3e,
	0x22, 0xed, 0x4c, 0xc5, 0xf3, 0x37, 0x12, 0xcb, 0xdf, 0x3f, 0xba, 0x7b, 0xc5, 0xc1, 0xaa, 0x76,
	0xcc, 0x52, 0xa9, 0x47, 0x3a, 0x49, 0xb0, 0x7c, 0x01, 0xd7, 0xd5, 0x5c, 0x88, 0xc2, 0x04, 0x75,
	0x30, 0xa5, 0x22, 0x5c, 0xea, 0x19, 0x75, 0x14, 0x36, 0x38, 0x7c, 0xac, 0x78, 0x1a, 0xe8, 0x13,
	0xa8, 0x99, 0x47, 0x70, 0xd9, 0x85, 0xe0, 0x26, 0x9d, 0xf0, 0xe4, 0x39, 0xbf, 0xf0, 0xfa, 0x83,
	0xc6, 0xb0, 0xc9, 0xac, 0x45, 0x0f, 0x49, 0xef, 0x54, 0x3f, 0x93, 0x5c, 0x7d, 0x27, 0x62, 0xef,
	0x1a, 0x4a, 0xa1, 0x02, 0xe8, 0x53, 0xb2, 0x6b, 0xca, 0xc0, 0x78, 0x26, 0x4e, 0xe3, 0x5f, 0x95,
	0xb7, 0x83, 0xa1, 0x0c, 0x6a, 0xa1, 0x80, 0xb6, 0x27, 0x75, 0x1e, 0x5b, 0xdb, 0x07, 0xde, 0x27,
	0x21, 0xbf, 0x04, 0xdb, 0xdb, 0x35, 0xde, 0x17, 0x36, 0x68, 0xee, 0x7b, 0x1e, 0x89, 0x71, 0xc6,
	0xb3, 0x5c, 0x7b, 0xd7, 0x8d, 0xe6, 0x2a, 0x04, 0x92, 0x2d, 0x63, 0x99, 0x41, 0xb4, 0x27, 0x4a,
	0xc6, 0xde, 0x1e, 0x46, 0x50, 0xc3, 0x50, 0x97, 0x79, 0x9a, 0x9e, 0x87, 0xfc, 0xd2, 0xe4, 0xe0,
	0xff, 0x56, 0x97, 0x2e, 0x48, 0x1f, 0x90, 0xdd, 0x4c, 0x65, 0x3c, 0xac, 0x68, 0x14, 0x69, 0x6b,
	0x28, 0x64, 0x05, 0xbc, 0x13, 0xe9, 0x0f, 0x79, 0xe4, 0xdd, 0x40, 0x4a, 0x05, 0xc0, 0x59, 0x99,
	0x8c, 0xc4, 0x8f, 0x79, 0x36, 0x16, 0x33, 0x15, 0x07, 0xde, 0xbe, 0x39, 0xab, 0x06, 0x42, 0x4d,
	0x40, 0xad, 0xcf, 0x65, 0x8c, 0x4e, 0xdf, 0x44, 0xa7, 0x5d, 0xa8, 0x64, 0xf0, 0x0b, 0x64, 0xdc,
	0x72, 0x18, 0x06, 0x32, 0xdd, 0x48, 0xc6, 0x2f, 0xe4, 0x6c, 0x29, 0x32, 0xef, 0x36, 0x12, 0x1c,
	0x84, 0x7e, 0x6b, 0xd6, 0x21, 0x0b, 0x67, 0x2b, 0xcf, 0x1b, 0x34, 0x86, 0xbb, 0x47, 0x1f, 0xd7,
	0x6a, 0x23, 0xe2, 0x3c, 0x9a, 0x94, 0x1c, 0xb1, 0x12, 0x21, 0x73, 0xb6, 0xf8, 0x7f, 0x36, 0x6c,
	0x53, 0x4f, 0x57, 0x93, 0x39, 0x8f, 0xc4, 0x88, 0x67, 0x1c, 0x74, 0xb9, 0xe0, 0x71, 0x80, 0xda,
	0xd1, 0x1f, 0x24, 0x6b, 0x87, 0x4f, 0xbf, 0x26, 0xc4, 0x48, 0xed, 0x4c, 0xea, 0xcc, 0xaa, 0xfa,
	0x3d, 0xda, 0x74, 0x36, 0xfc, 0x47, 0x59, 0xfb, 0x49, 0x15, 0xc7, 0x54, 0x86, 0xe1, 0xb1, 0xe0,
	0x31, 0xf4, 0x7d, 0x08, 0x15, 0x67, 0x49, 0x93, 0xe1, 0x6f, 0xb8, 0x87, 0xa1, 0xd2, 0x28, 0x51,
	0xec, 0x73, 0x3b, 0xac, 0xb4, 0x41, 0x7b, 0x6f, 0x65, 0x8c, 0x4b, 0x4d, 0x5c, 0x2a, 0x4c, 0xf8,
	0x52, 0x20, 0xf4, 0x0c, 0x9b, 0x45, 0x8f, 0xe1, 0x6f, 0x7f, 0x52, 0x75, 0x43, 0x38, 0xd1, 0xee,
	0x3e, 0xa9, 0x0e, 0x2c, 0x4c, 0xfa, 0x98, 0x74, 0x0b, 0x9f, 0x6c, 0x3e, 0xee, 0x5c, 0x09, 0xaa,
	0x20, 0xb0, 0x92, 0xea, 0xff, 0xb6, 0x45, 0xee, 0x16, 0xcb, 0x70, 0x29, 0x40, 0x0d, 0x52, 0x67,
	0x72, 0xa6, 0x99, 0xca, 0xe3, 0x80, 0xee, 0x93, 0x56, 0x0a, 0x3f, 0xec, 0xac, 0x34, 0x86, 0xeb,
	0xc6, 0x56, 0xdd, 0x8d, 0x43, 0xd2, 0xc3, 0x96, 0x8e, 0x53, 0xcf, 0xcc, 0xc2, 0x0a, 0xa0, 0x5f,
	0x90, 0xce, 0x54, 0x45, 0xe8, 0xe3, 0xf6, 0x87, 0xd4, 0xac, 0x60, 0xd7, 0x94, 0xdd, 0x5a, 0x53,
	0xf6, 0x01, 0xe9, 0x4a, 0x3d, 0x4e, 0x52, 0x69, 0x07, 0x65, 0x97, 0x95, 0x36, 0x38, 0x2a, 0xf5,
	0x08, 0xda, 0x2e, 0xf6, 0xc7, 0x2e, 0x2b, 0x4c, 0x08, 0x4c, 0xea, 0x57, 0x32, 0xc6, 0x0e, 0xd8,
	0x65, 0xc6, 0xf0, 0x7f, 0xdf, 0x22, 0xb7, 0x36, 0xa7, 0x83, 0x1e, 0x93, 0x1e, 0x06, 0x8f, 0xde,
	0x9b, 0xfb, 0x7a, 0xff, 0x8a, 0xf7, 0x1b, 0x52, 0xc8, 0xaa, 0x6d, 0xe0, 0x2a, 0x26, 0xe3, 0x95,
	0x4d, 0x5c, 0x8b, 0x95, 0x76, 0x99, 0xb9, 0x33, 0xa5, 0x45, 0x2d, 0x73, 0x00, 0x80, 0x8c, 0xcd,
	0x60, 0x37, 0x71, 0x6e, 0x9b, 0xe6, 0xeb, 0x40, 0xd5, 0xa3, 0x02, 0xa3, 0xad, 0x3d, 0x2a, 0x30,
	0xe0, 0x5a, 0x65, 0xda, 0xeb, 0x95, 0xb9, 0x47, 0x48, 0xc4, 0x2f, 0x5e, 0xd9, 0xa2, 0x76, 0x4c,
	0x13, 0xa8, 0x10, 0xff, 0xef, 0x76, 0x75, 0x13, 0x21, 0x48, 0xe7, 0xd5, 0xd0, 0x70, 0x5f, 0x0d,
	0x70, 0x0f, 0x0b, 0x8d, 0x63, 0x88, 0x9b, 0xee, 0x61, 0x41, 0x60, 0x25, 0x15, 0xda, 0x9d, 0xc6,
	0x46, 0xfc, 0xf4, 0x0c, 0x5f, 0x03, 0x36, 0x03, 0x75, 0x10, 0x9a, 0x74, 0x01, 0xc0, 0xf4, 0xb0,
	0x69, 0xa8, 0x61, 0xd0, 0x7e, 0x1d, 0xfb, 0x58, 0x14, 0x0f, 0xac, 0x35, 0x14, 0xde, 0x24, 0x06,
	0x19, 0x9f, 0x3c, 0x17, 0xf1, 0xeb, 0x45, 0xce, 0x6d, 0x56, 0xd6, 0x61, 0x97, 0x09, 0x36, 0xb4,
	0x8d, 0x4e, 0x9d, 0x69, 0xe1, 0x2a, 0x8a, 0xf1, 0xc9, 0x19, 0x1f, 0x71, 0x65, 0xc7, 0x6b, 0x1d,
	0xac, 0xa2, 0x78, 0x36, 0xc2, 0x28, 0x7a, 0x6e, 0x14, 0x06, 0xa3, 0x9f, 0x92, 0x3d, 0xbb, 0x69,
	0xa1, 0xde, 0xda, 0x26, 0x69, 0x26, 0xee, 0x15, 0x1c, 0x47, 0x97, 0xcb, 0xeb, 0xe3, 0x8d, 0xae,
	0x61, 0xf4, 0x13, 0xb2, 0x0d, 0x9a, 0xc7, 0xe9, 0xdb, 0x3f, 0xba, 0xb9, 0xb1, 0x35, 0x30, 0xa4,
	0x94, 0x1d, 0x6d, 0xc7, 0xe9, 0x68, 0xfb, 0xa4, 0xa5, 0x67, 0x2a, 0x35, 0x63, 0xb5, 0xc9, 0x8c,
	0x01, 0x93, 0x41, 0x97, 0x97, 0x1d, 0x67, 0x6a, 0x7f, 0x6d, 0x32, 0x6c, 0xd0, 0x84, 0xb3, 0xa5,
	0x26, 0xeb, 0xbd, 0x35, 0x59, 0x1f, 0x92, 0xde, 0x1b, 0xc9, 0xd5, 0x18, 0x8f, 0x35, 0x83, 0xb6,
	0x02, 0x20, 0x66, 0x3b, 0x09, 0xdd, 0x11, 0x5b, 0xc3, 0x8c, 0xf8, 0x9f, 0xcc, 0x45, 0x9c, 0xe1,
	0x78, 0x45, 0xf1, 0xa3, 0x49, 0xf7, 0x48, 0x53, 0x8b, 0x0b, 0x3b, 0x52, 0xe1, 0x27, 0x78, 0x92,
	0xc2, 0x98, 0xe7, 0x69, 0x80, 0x53, 0xb4, 0xc5, 0x4a, 0x9b, 0x7e, 0x4e, 0x6e, 0x98, 0x9c, 0x3f,
	0x49, 0x92, 0xf0, 0x72, 0x24, 0xb5, 0x56, 0xe1, 0x4a, 0xe0, 0x28, 0x6d, 0xb1, 0x4d, 0x4b, 0x20,
	0x0e, 0x3d, 0xc3, 0xa8, 0x6e, 0x23, 0xc9, 0x5a, 0x80, 0xbf, 0x99, 0x22, 0xee, 0x19, 0xdc, 0x58,
	0xfe, 0x83, 0x4a, 0x5c, 0x70, 0x2a, 0xf0, 0xe0, 0xef, 0x69, 0xd1, 0x76, 0xad, 0xe5, 0x0b, 0xd2,
	0x2f, 0x0b, 0xc6, 0x97, 0x50, 0x54, 0x78, 0xdd, 0x21, 0x69, 0x53, 0x51, 0x61, 0x11, 0x07, 0xc9,
	0x92, 0x7e, 0x46, 0x5a, 0x50, 0x87, 0xe2, 0xdf, 0x9a, 0x9b, 0x1b, 0xab, 0xc4, 0x0c, 0xe7, 0xfc,
	0x7f, 0xe7, 0x8d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x96, 0xb1, 0x2e, 0xa6, 0x87, 0x0d, 0x00,
	0x00,
}
