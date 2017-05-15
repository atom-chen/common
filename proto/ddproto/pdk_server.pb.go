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

// Ignoring public import of pdk_base_playerInfo from pdk_base.proto

// Ignoring public import of pdk_base_playerRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_commonRateInfo from pdk_base.proto

// Ignoring public import of pdk_base_deskInfo from pdk_base.proto

// Ignoring public import of pdk_enum_protoId from pdk_base.proto

// Ignoring public import of pdk_enum_errorCode from pdk_base.proto

// Ignoring public import of pdk_enum_actType from pdk_base.proto

// Ignoring public import of pdk_enum_gameStatus from pdk_base.proto

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
func (*PdkSrvOutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

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
func (*PdkSrvDeskTongJi) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{1} }

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
func (*PdkSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{2} }

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
func (*PdkSrvGameData) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{3} }

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
	// 跑得快的账单
	Coin             *int64              `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	LoseUser         *uint32             `protobuf:"varint,2,opt,name=loseUser" json:"loseUser,omitempty"`
	WinUser          *uint32             `protobuf:"varint,3,opt,name=winUser" json:"winUser,omitempty"`
	Desc             *string             `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	OutPokerPais     *PdkSrvOutPokerPais `protobuf:"bytes,5,opt,name=OutPokerPais" json:"OutPokerPais,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *PdkSrvBillBean) Reset()                    { *m = PdkSrvBillBean{} }
func (m *PdkSrvBillBean) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvBillBean) ProtoMessage()               {}
func (*PdkSrvBillBean) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{4} }

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

func (m *PdkSrvBillBean) GetOutPokerPais() *PdkSrvOutPokerPais {
	if m != nil {
		return m.OutPokerPais
	}
	return nil
}

type PdkSrvBill struct {
	// 跑得快的账单
	WinCoin          *int64            `protobuf:"varint,1,opt,name=winCoin" json:"winCoin,omitempty"`
	BillBean         []*PdkSrvBillBean `protobuf:"bytes,2,rep,name=billBean" json:"billBean,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *PdkSrvBill) Reset()                    { *m = PdkSrvBill{} }
func (m *PdkSrvBill) String() string            { return proto.CompactTextString(m) }
func (*PdkSrvBill) ProtoMessage()               {}
func (*PdkSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{5} }

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
func (*PdkSrvUserStatisticsRound) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{6} }

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
func (*PdkSrvUserStatistics) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{7} }

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
func (*PdkSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{8} }

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
func (*PdkSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{9} }

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
func (*PdkSrvBak) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{10} }

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

var fileDescriptor38 = []byte{
	// 1330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x5d, 0x6f, 0xdb, 0x36,
	0x17, 0x7e, 0x1d, 0xc7, 0x5f, 0x74, 0x92, 0xe6, 0x65, 0xd3, 0x56, 0x4d, 0xb3, 0xce, 0x10, 0x8a,
	0xc2, 0xdb, 0x80, 0x62, 0x48, 0x51, 0x0c, 0x18, 0xf6, 0x81, 0x26, 0xc6, 0xd6, 0x14, 0xe9, 0xea,
	0xd1, 0xed, 0x8a, 0xf5, 0xc6, 0xa0, 0x2d, 0xce, 0x66, 0x2d, 0x89, 0x86, 0x28, 0xb9, 0x49, 0x7f,
	0xc3, 0xae, 0xf7, 0x47, 0xf6, 0x07, 0x76, 0xb5, 0xfb, 0x5d, 0xef, 0xcf, 0x0c, 0xe7, 0x90, 0x92,
	0x28, 0xc7, 0x58, 0x8b, 0x5d, 0xc5, 0xe7, 0xe1, 0x43, 0xf1, 0x7c, 0xf0, 0x39, 0x87, 0x21, 0xfb,
	0xcb, 0x60, 0x31, 0xd6, 0x22, 0x59, 0x89, 0xe4, 0xc1, 0x32, 0x51, 0xa9, 0xa2, 0xad, 0x20, 0xc0,
	0x1f, 0x87, 0xb7, 0xa7, 0x2a, 0x8a, 0x54, 0x6c, 0x57, 0xc7, 0x4b, 0xb5, 0xc8, 0x39, 0x87, 0x7b,
	0xb0, 0x6b, 0xc2, 0xb5, 0x30, 0xb6, 0xff, 0xc7, 0x16, 0x39, 0xc0, 0x0f, 0x25, 0xab, 0xb1, 0xca,
	0xd2, 0x21, 0x50, 0x87, 0x5c, 0x6a, 0x7a, 0x48, 0xda, 0x0b, 0x71, 0xf9, 0x13, 0x0f, 0x33, 0xe1,
	0xd5, 0x7a, 0xb5, 0x7e, 0x83, 0x15, 0x36, 0xfd, 0x92, 0x74, 0x96, 0x39, 0xd1, 0xdb, 0xea, 0xd5,
	0xfb, 0xdd, 0xe3, 0xa3, 0x07, 0xf6, 0xf0, 0x07, 0xf9, 0xd9, 0xc9, 0x6a, 0x9c, 0x93, 0x58, 0x49,
	0xa7, 0x94, 0x6c, 0xa7, 0x97, 0x4b, 0xe1, 0xd5, 0xf1, 0x9b, 0xf8, 0x9b, 0xde, 0x24, 0x4d, 0xa9,
	0x4f, 0x54, 0x34, 0xf1, 0xb6, 0x7b, 0xb5, 0x7e, 0x9b, 0x59, 0x8b, 0xde, 0x25, 0x64, 0xaa, 0xb2,
	0x38, 0x1d, 0x64, 0xf2, 0x9d, 0xf4, 0x1a, 0xb8, 0xc3, 0x41, 0xe8, 0x3d, 0xb2, 0x8b, 0xd6, 0x88,
	0xc7, 0xef, 0xe6, 0x3c, 0x9e, 0x79, 0x4d, 0xa4, 0x54, 0x41, 0xea, 0x93, 0x1d, 0x03, 0x48, 0x43,
	0x6a, 0x21, 0xa9, 0x82, 0x15, 0x9c, 0x9f, 0x2d, 0xa7, 0xed, 0x70, 0x2c, 0x06, 0x5e, 0x66, 0x5a,
	0x24, 0x67, 0x81, 0xd7, 0xe9, 0xd5, 0xfa, 0xbb, 0xcc, 0x5a, 0xbe, 0x26, 0x34, 0xcf, 0x60, 0x20,
	0xf4, 0xe2, 0x85, 0x8a, 0x67, 0x4f, 0x25, 0x7d, 0x48, 0x1a, 0x13, 0x15, 0x4d, 0xb4, 0x57, 0xc3,
	0xfc, 0x7c, 0x54, 0xe4, 0x67, 0x53, 0xb6, 0x99, 0xe1, 0xd2, 0x3e, 0xb9, 0x86, 0x47, 0xfe, 0x28,
	0x79, 0x3c, 0x1b, 0xc8, 0xd7, 0xf3, 0xcc, 0xdb, 0x42, 0x4f, 0xd6, 0x61, 0xff, 0xaf, 0x16, 0xd9,
	0x71, 0x4f, 0xa5, 0xdf, 0x90, 0x2e, 0x0f, 0xc3, 0xfc, 0x8b, 0xf6, 0xd4, 0x7f, 0xaf, 0x8a, 0xbb,
	0x81, 0x7e, 0x45, 0x48, 0x20, 0x8b, 0xed, 0x1f, 0x52, 0x54, 0x87, 0x4f, 0x1f, 0x91, 0x26, 0xc4,
	0xc3, 0x25, 0xd6, 0xf5, 0xbd, 0xe1, 0x5a, 0x32, 0xa4, 0x3d, 0x90, 0xef, 0xe6, 0xd9, 0x90, 0xcb,
	0x97, 0x5a, 0x24, 0x58, 0xfe, 0x5d, 0x56, 0xc1, 0x68, 0x8f, 0x74, 0x03, 0x08, 0xf9, 0xa5, 0xc9,
	0x7d, 0x03, 0x29, 0x2e, 0x04, 0x5f, 0xe1, 0xd3, 0x54, 0xae, 0x84, 0xa5, 0x34, 0xcd, 0x57, 0x5c,
	0x8c, 0x3e, 0x24, 0xcd, 0x54, 0xc5, 0xb3, 0x37, 0x12, 0xcb, 0xdf, 0x3d, 0xbe, 0x73, 0xc5, 0xc1,
	0xb2, 0x76, 0xcc, 0x52, 0xa9, 0x47, 0x5a, 0xcb, 0x60, 0xf1, 0x02, 0xae, 0xab, 0xb9, 0x10, 0xb9,
	0x09, 0xea, 0x60, 0x4a, 0x45, 0xb8, 0xd4, 0x31, 0xea, 0xc8, 0x6d, 0x70, 0xf8, 0x44, 0xf1, 0x24,
	0xd0, 0xa7, 0x50, 0x33, 0x8f, 0xe0, 0xb2, 0x0b, 0xc1, 0x4d, 0x3a, 0xe5, 0xcb, 0x67, 0xfc, 0xc2,
	0xeb, 0xf6, 0x6a, 0xfd, 0x3a, 0xb3, 0x16, 0x3d, 0x22, 0x9d, 0x33, 0xfd, 0x54, 0x72, 0xf5, 0x9d,
	0x88, 0xbd, 0x1d, 0x94, 0x42, 0x09, 0xd0, 0x27, 0x64, 0xcf, 0x94, 0x81, 0xf1, 0x54, 0x9c, 0xc5,
	0xbf, 0x28, 0x6f, 0x17, 0x43, 0xe9, 0x55, 0x42, 0x01, 0x6d, 0x8f, 0xab, 0x3c, 0xb6, 0xb6, 0x0f,
	0xbc, 0x5f, 0x86, 0xfc, 0x12, 0x6c, 0x6f, 0xcf, 0x78, 0x9f, 0xdb, 0xa0, 0xb9, 0xef, 0x79, 0x24,
	0x46, 0x29, 0x4f, 0x33, 0xed, 0x5d, 0x33, 0x9a, 0x2b, 0x11, 0x48, 0xb6, 0x8c, 0x65, 0x0a, 0xd1,
	0x9e, 0x2a, 0x19, 0x7b, 0xfb, 0x18, 0x41, 0x05, 0x43, 0x5d, 0x66, 0x49, 0x32, 0x0c, 0xf9, 0xa5,
	0xc9, 0xc1, 0xff, 0xad, 0x2e, 0x5d, 0x90, 0xde, 0x27, 0x7b, 0xa9, 0x4a, 0x79, 0x58, 0xd2, 0x28,
	0xd2, 0xd6, 0x50, 0xc8, 0x0a, 0x78, 0x27, 0x92, 0x1f, 0xb2, 0xc8, 0xbb, 0x8e, 0x94, 0x12, 0x80,
	0xb3, 0x52, 0x19, 0x89, 0xe7, 0x59, 0x3a, 0x12, 0x53, 0x15, 0x07, 0xde, 0x81, 0x39, 0xab, 0x02,
	0x42, 0x4d, 0x40, 0xad, 0xcf, 0x64, 0x8c, 0x4e, 0xdf, 0x40, 0xa7, 0x5d, 0xa8, 0x60, 0xf0, 0x0b,
	0x64, 0xdc, 0x74, 0x18, 0x06, 0x32, 0xdd, 0x48, 0xc6, 0x2f, 0xe4, 0x74, 0x21, 0x52, 0xef, 0x16,
	0x12, 0x1c, 0x84, 0x7e, 0x6b, 0xd6, 0x21, 0x0b, 0xe7, 0x2b, 0xcf, 0xeb, 0xd5, 0xfa, 0x7b, 0xc7,
	0x1f, 0x57, 0x6a, 0x23, 0xe2, 0x2c, 0x1a, 0x17, 0x1c, 0xb1, 0x12, 0x21, 0x73, 0xb6, 0xf8, 0x7f,
	0xd6, 0x6c, 0x53, 0x4f, 0x56, 0xe3, 0x19, 0x8f, 0xc4, 0x80, 0xa7, 0x1c, 0x74, 0x39, 0xe7, 0x71,
	0x80, 0xda, 0xd1, 0x1f, 0x24, 0x6b, 0x87, 0x4f, 0xbf, 0x26, 0xc4, 0x48, 0xed, 0x5c, 0xea, 0xd4,
	0xaa, 0xfa, 0x3d, 0xda, 0x74, 0x36, 0xfc, 0x47, 0x59, 0xfb, 0xbf, 0x3b, 0x81, 0x4c, 0x64, 0x18,
	0x9e, 0x08, 0x1e, 0x43, 0xe3, 0x87, 0x58, 0x71, 0x98, 0xd4, 0x19, 0xfe, 0x86, 0x8b, 0x18, 0x2a,
	0x8d, 0x1a, 0xc5, 0x46, 0xb7, 0xcb, 0x0a, 0x1b, 0xc4, 0xf7, 0x56, 0xc6, 0xb8, 0x54, 0xc7, 0xa5,
	0xdc, 0x84, 0x2f, 0x05, 0x42, 0x4f, 0xb1, 0x5b, 0x74, 0x18, 0xfe, 0xa6, 0x8f, 0xc9, 0xce, 0x73,
	0xc7, 0x15, 0x6c, 0x13, 0xef, 0xf5, 0xb7, 0xb2, 0xc5, 0x1f, 0x97, 0x1d, 0x15, 0x9c, 0xb6, 0x0e,
	0x9c, 0x96, 0x3e, 0xe7, 0x26, 0x7d, 0x44, 0xda, 0x79, 0x58, 0x36, 0xa7, 0xb7, 0xaf, 0x1c, 0x94,
	0x13, 0x58, 0x41, 0xf5, 0x7f, 0xdd, 0x22, 0x77, 0xf2, 0x65, 0xb8, 0x58, 0xa0, 0x28, 0xa9, 0x53,
	0x39, 0xd5, 0x4c, 0x65, 0x71, 0x40, 0x0f, 0x48, 0x23, 0x81, 0x1f, 0x76, 0xde, 0x1a, 0xc3, 0x75,
	0x63, 0xab, 0xea, 0xc6, 0x11, 0xe9, 0xe0, 0x58, 0xc0, 0xc9, 0x69, 0xe6, 0x69, 0x09, 0xd0, 0x2f,
	0x48, 0x6b, 0xa2, 0x22, 0xf4, 0x71, 0xfb, 0x43, 0xea, 0x9e, 0xb3, 0x2b, 0xdd, 0xa1, 0xb1, 0xd6,
	0x1d, 0x0e, 0x49, 0x5b, 0xea, 0xd1, 0x32, 0x91, 0x76, 0xd8, 0xb6, 0x59, 0x61, 0x83, 0xa3, 0x52,
	0x0f, 0xa0, 0x75, 0x63, 0x8f, 0x6d, 0xb3, 0xdc, 0x84, 0xc0, 0xa4, 0x7e, 0x25, 0x63, 0xec, 0xa2,
	0x6d, 0x66, 0x0c, 0xff, 0xb7, 0x2d, 0x72, 0x73, 0x73, 0x3a, 0xe8, 0x09, 0xe9, 0x60, 0xf0, 0xe8,
	0xbd, 0xb9, 0xf3, 0xf7, 0xae, 0x78, 0xbf, 0x21, 0x85, 0xac, 0xdc, 0x06, 0xae, 0x62, 0x32, 0x5e,
	0xd9, 0xc4, 0x35, 0x58, 0x61, 0x17, 0x99, 0x3b, 0x57, 0x5a, 0x54, 0x32, 0x07, 0x00, 0xb4, 0x02,
	0xf3, 0x38, 0x30, 0x71, 0x6e, 0x9b, 0x06, 0xee, 0x40, 0xe5, 0xc3, 0x04, 0xa3, 0xad, 0x3c, 0x4c,
	0x30, 0xe0, 0x4a, 0x65, 0x9a, 0xeb, 0x95, 0xb9, 0x4b, 0x48, 0xc4, 0x2f, 0x5e, 0xd9, 0xa2, 0xb6,
	0x4c, 0x23, 0x29, 0x11, 0xff, 0xef, 0x66, 0x79, 0x13, 0x21, 0x48, 0xe7, 0xe5, 0x51, 0x73, 0x5f,
	0x1e, 0x70, 0x0f, 0xf3, 0x3e, 0x81, 0x21, 0x6e, 0xba, 0x87, 0x39, 0x81, 0x15, 0x54, 0x68, 0x99,
	0x1a, 0x9b, 0xf9, 0x93, 0x73, 0x7c, 0x51, 0xd8, 0x0c, 0x54, 0x41, 0x68, 0xf4, 0x39, 0x00, 0x13,
	0xc8, 0xa6, 0xa1, 0x82, 0x41, 0x0b, 0x77, 0xec, 0x13, 0x91, 0x3f, 0xd2, 0xd6, 0x50, 0x78, 0xd7,
	0x18, 0x64, 0x74, 0xfa, 0x4c, 0xc4, 0xaf, 0xe7, 0x19, 0xb7, 0x59, 0x59, 0x87, 0x5d, 0x26, 0xd8,
	0xd0, 0x7a, 0x5a, 0x55, 0xa6, 0x85, 0xcb, 0x28, 0x46, 0xa7, 0xe7, 0x7c, 0xc0, 0x95, 0x1d, 0xd1,
	0x55, 0xb0, 0x8c, 0xe2, 0xe9, 0x00, 0xa3, 0xe8, 0xb8, 0x51, 0x18, 0x8c, 0x7e, 0x4a, 0xf6, 0xed,
	0xa6, 0xb9, 0x7a, 0x6b, 0x1b, 0xad, 0x99, 0xda, 0x57, 0x70, 0x1c, 0x7f, 0x2e, 0xaf, 0x8b, 0x37,
	0xba, 0x82, 0xd1, 0x4f, 0xc8, 0x36, 0x68, 0x1e, 0x27, 0x78, 0xf7, 0xf8, 0xc6, 0xc6, 0xd6, 0xc0,
	0x90, 0x52, 0x34, 0xc5, 0x5d, 0xa7, 0x29, 0x1e, 0x90, 0x86, 0x9e, 0xaa, 0xc4, 0x8c, 0xe6, 0x3a,
	0x33, 0x06, 0x4c, 0x17, 0x5d, 0x5c, 0x76, 0x9c, 0xcb, 0xdd, 0xb5, 0xe9, 0xb2, 0x41, 0x13, 0xce,
	0x96, 0x8a, 0xac, 0xf7, 0xd7, 0x64, 0x7d, 0x44, 0x3a, 0x6f, 0x24, 0x57, 0x23, 0x3c, 0xd6, 0x0c,
	0xeb, 0x12, 0x80, 0x98, 0xed, 0x34, 0x75, 0xc7, 0x74, 0x05, 0x33, 0xe2, 0x7f, 0x3c, 0x13, 0x71,
	0x8a, 0x23, 0x1a, 0xc5, 0x8f, 0x26, 0xdd, 0x27, 0x75, 0x2d, 0x2e, 0xec, 0x58, 0x86, 0x9f, 0xe0,
	0x49, 0x02, 0x4f, 0x05, 0x9e, 0x04, 0x38, 0x89, 0x1b, 0xac, 0xb0, 0xe9, 0xe7, 0xe4, 0xba, 0xc9,
	0xf9, 0xe3, 0xe5, 0x32, 0xbc, 0x1c, 0x48, 0xad, 0x55, 0xb8, 0x12, 0x38, 0x8e, 0x1b, 0x6c, 0xd3,
	0x12, 0x88, 0x43, 0x4f, 0x31, 0xaa, 0x5b, 0x48, 0xb2, 0x16, 0xe0, 0x6f, 0x26, 0x88, 0x7b, 0x06,
	0x37, 0x96, 0x7f, 0xbf, 0x14, 0x17, 0x9c, 0x0a, 0x3c, 0xf8, 0x7b, 0x96, 0xb7, 0x5d, 0x6b, 0xf9,
	0x82, 0x74, 0x8b, 0x82, 0xf1, 0x05, 0x14, 0x15, 0x5e, 0x88, 0x48, 0xda, 0x54, 0x54, 0x58, 0xc4,
	0x59, 0xb4, 0xa0, 0x9f, 0x91, 0x06, 0xd4, 0x21, 0xff, 0xd7, 0xe8, 0xc6, 0xc6, 0x2a, 0x31, 0xc3,
	0x19, 0xfe, 0x6f, 0x58, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x06, 0x7c, 0xf2, 0x35, 0xcb, 0x0d,
	0x00, 0x00,
}
