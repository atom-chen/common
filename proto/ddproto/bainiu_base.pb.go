// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bainiu_base.proto

/*
Package ddproto is a generated protocol buffer package.

It is generated from these files:
	bainiu_base.proto
	bainiu_play.proto
	bainiu_server.proto
	common_client.proto
	common_client_award.proto
	common_client_pay.proto
	common_enum.proto
	common_mj.proto
	common_pdk.proto
	common_server.proto
	common_server_award.proto
	common_server_config.proto
	common_server_model.proto
	common_server_msg.proto
	common_server_poker.proto
	common_server_redis.proto
	common_server_task.proto
	ddz_base.proto
	ddz_desk.proto
	ddz_hall.proto
	ddz_play.proto
	ddz_server.proto
	erddz_base.proto
	erddz_desk.proto
	erddz_hall.proto
	erddz_play.proto
	group_base.proto
	group_req.proto
	hall.proto
	hall_data.proto
	hall_playback.proto
	hot_update.proto
	mj_baishan_base.proto
	niuniu_base.proto
	niuniu_play.proto
	niuniu_server.proto
	pdk_base.proto
	pdk_desk.proto
	pdk_hall.proto
	pdk_play.proto
	pdk_server.proto
	pez_base.proto
	pez_desk.proto
	pez_hall.proto
	pez_play.proto
	phz_base.proto
	phz_desk.proto
	phz_play.proto
	phz_server.proto
	rpc_greeter.proto
	rpc_hall.proto
	rpc_niuniu.proto
	rpc_pdk.proto
	zjh_base.proto
	zjh_desk.proto
	zjh_hall.proto
	zjh_play.proto
	zjh_serever.proto

It has these top-level messages:
	BainiuClientPoker
	BainiuAreaInfo
	BainiuClientDesk
	BainiuClientUser
	BainiuEnterDeskReq
	BainiuEnterDeskAck
	BainiuYazhuOt
	BainiuYazhuReq
	BainiuYazhuAck
	BainiuYazhuBc
	BainiuFapaiBc
	BainiuWinUserItem
	BainiuAreaBillItem
	BainiuBipaiBc
	BainiuWincountReq
	BainiuWincountItem
	BainiuWincountAck
	BainiuSrvPoker
	ProtoHeader
	Heartbeat
	ServerInfo
	QuickConn
	AckQuickConn
	WeixinInfo
	CommonReqReg
	CommonReqRegViaInput
	CommonAckReg
	CommonReqGameLogin
	CommonReqGameLoginViaInput
	CommonAckGameLogin
	CommonReqQrLogin
	CommonAckQrLogin
	CommonReqQrWxInfo
	CommonAckQrWxInfo
	CommonAckReconnect
	CommonReqReconnect
	CommonReqGameState
	CommonAckGameState
	CommonReqLogout
	CommonAckLogout
	CommonReqFeedback
	ClientBasePoker
	CommonReqMessage
	CommonBcMessage
	CommonReqNotice
	CommonAckNotice
	CommonReqEnterAgentMode
	CommonAckEnterAgentMode
	CommonReqQuitAgentMode
	CommonAckQuitAgentMode
	CommonReqLeaveDesk
	CommonAckLeaveDesk
	CommonReqKickout
	CommonBcKickout
	CommonReqAllowance
	CommonAckAllowance
	CommonReqApplyDissolve
	CommonBcApplyDissolve
	CommonReqApplyDissolveBack
	CommonAckApplyDissolveBack
	CommonAckTimeout
	CommonBcUserBreak
	CommonReqClickStatistic
	CommonReqOffline
	CommonReqUploadLocation
	CommonBcLeaveTimeout
	CommonDeskByAgent
	CommonReqListCoinDesk
	CommonAckListCoinDesk
	CommonCoinDeskInfo
	AwardReqOnline
	WardAckOnline
	AwardReqOnlineInfo
	AwardAckOnlineInfo
	AwardReqGetNewUser
	AwardAckGetNewUser
	PayBaseProduct
	PayBasePaymodel
	PayBaseDetails
	WxpayReqUnifiedorder
	WxpayAckUnifiedorder
	WxpayReqSyncChecker
	WxpayAckSyncChecker
	ApplepayReqRechargecb
	ApplepayAcksRechargecb
	CardInfo
	RoomTypeInfo
	PlayOptions
	ChangShaPlayOptions
	BaiShanPlayOptions
	ZhuanZhuanPlayOptions
	HaiNanPlayOptions
	ComposeCard
	PlayerCard
	GangBean
	PdkBaseRoomTypeInfo
	PdkBaseTimerInfo
	GameSession
	CommonSrvGameUser
	CommonSrvGameDesk
	RedisDeskByAgent
	RedisUserGameBill
	UserGameBill
	HzhRobotGameBill
	Taward
	AwardMOnline
	TConfigSys
	TConfigDrawLottery
	User
	TNotice
	TGameCounts
	TUserTask
	Push
	CommonSrvMsgGameCount
	CommonSrvPokerPai
	RedisGameServerStatus
	TTask
	DdzBaseRoomTypeInfo
	DdzBasePlayerInfo
	DdzBasePlayerRateInfo
	DdzBaseCommonRateInfo
	DdzBaseTimerInfo
	DdzBaseDeskInfo
	DdzReqDissolveDesk
	DdzAckDissolveDesk
	DdzReqReady
	DdzAckReady
	DdzBaseWinCoinInfo
	DdzBaCurrentResult
	DdzBaseEndLotteryInfo
	DdzBcEndLottery
	DdzReqCreateDesk
	DdzAckCreateDesk
	DdzReqGameRecord
	DdzBaseUserRecord
	DdzBaseGameRecord
	DdzAckGameRecord
	DdzReqEnterDesk
	DdzAckEnterDesk
	DdzBcOpening
	DdzBcDealCards
	DdzReqShowHandPokers
	DdzAckShowHandPokers
	DdzReqJiaoDiZhu
	DdzAckJiaoDiZhu
	DdzReqRobDiZhu
	DdzAckRobDiZhu
	DdzReqDouble
	DdzAckDouble
	DdzBcStartPlay
	DdzReqMenuZhua
	DdzAckMenuZhua
	DdzReqSeeCards
	DdzAckSeeCards
	DdzReqPull
	DdzAckPull
	DdzReqOutCards
	DdzAckOutCards
	DdzReqActGuo
	DdzAckGuoAck
	DdzBcOverTurn
	DdzBcGameInfo
	DdzSrvOutPokerPais
	DdzSrvDeskTongJi
	DdzSrvDesk
	DdzSrvGameData
	DdzSrvBillBean
	DdzSrvBill
	DdzSrvUserStatisticsRound
	DdzSrvUserStatistics
	DdzSrvUser
	DdzSrvRoom
	DdzSrvBak
	ErddzBaseRoomTypeInfo
	ErddzBasePlayerInfo
	ErddzBasePlayerRateInfo
	ErddzBaseCommonRateInfo
	ErddzBaseTimerInfo
	ErddzBaseDeskInfo
	ErddzReqDissolveDesk
	ErddzAckDissolveDesk
	ErddzReqReady
	ErddzAckReady
	ErddzBaseWinCoinInfo
	ErddzBcCurrentResult
	ErddzBaseEndLotteryInfo
	ErddzBcEndLottery
	ErddzReqCreateDesk
	ErddzAckCreateDesk
	ErddzReqEnterDesk
	ErddzAckEnterDesk
	ErddzBcOpening
	ErddzBcDealCards
	ErddzReqJiaoDiZhu
	ErddzAckJiaoDiZhu
	ErddzReqRobDiZhu
	ErddzAckRobDiZhu
	ErddzReqRangcards
	ErddzAckRangcards
	ErddzReqDouble
	ErddzAckDouble
	ErddzBcStartPlay
	ErddzReqOutCards
	ErddzAckOutCards
	ErddzReqActGuo
	ErddzAckActGuo
	ErddzBcOverTurn
	ErddzBcGameInfo
	GroupInfo
	GroupMemberInfo
	GameOption
	GroupMsgItem
	HallGroupListReq
	HallGroupListAck
	HallGroupAddReq
	HallGroupAddAck
	HallCheckGroupUpdateReq
	HallCheckGroupUpdateAck
	HallGroupSendMsgReq
	HallGroupSendMsgBc
	HallGroupDeskListReq
	HallGroupDeskListAck
	HallReqEvent
	HallAckEvent
	HallLotteryItem
	HallSignLotteryInfo
	HallDrawLotteryInfo
	HallReqMailList
	HallAckMailList
	HallReqTask
	HallAckTask
	HallReqCheckTask
	HallAckCheckTask
	HallReqTaskSum
	HallAckTaskSum
	HallReqCheckBonus
	HallAckCheckBonus
	HallReqBagItems
	HallAckBagItems
	HallReqUserData
	HallAckUserData
	HallReqUpdateRealData
	HallAckUpdateRealData
	HallReqGoodsList
	HallAckGoodsList
	HallReqGoodsBuy
	HallAckGoodsBuy
	HallGoodsItemMsg
	HallReqRank
	HallAckRank
	HallReqDrawLottery
	HallAckDrawLottery
	HallReqDsLotteryInfo
	HallAckDsLotteryInfo
	HallReqFriendsList
	HallAckFriendsList
	HallReqRecommendUserList
	HallAckRecommendUserList
	HallReqFriendsSearch
	HallAckFriendsSearch
	HallReqAddFriend
	HallAckAddFriend
	HallReqDelFriend
	HallAckDelFriend
	HallFriendState
	HallUserInfo
	HallAckStrongboxInfo
	HallReqStrongboxInfo
	HallReqStrongboxAccess
	HallAckStrongboxAccess
	Game_GameRecord
	BeanUserRecord
	BeanGameRecord
	Game_AckGameRecord
	HallBeanBisai
	HallReqBisai
	HallReqFriendLotteryList
	HallAckFriendLotteryList
	HallReqFriendLotteryDraw
	HallAckFriendLotteryDraw
	HallReqDistanceMatched
	HallAckDistanceMatched
	HallReqAgentRoomGamingList
	HallAckAgentRoomGamingList
	HallReqAgentRoomHistoryList
	HallAckAgentRoomHistoryList
	HallItemEvent
	HallMailItem
	HallBagItem
	HallItemTask
	HallUserData
	HallRankItem
	CoinZone
	DiamondZone
	ExchangeZone
	BuyZone
	GoodsItem
	HallStrongboxInfo
	HallReqNewRoomList
	HallAckNewRoomList
	HallNewRoom_List
	PlaybackReqPage
	PlaybackAckPage
	PlayerInfo
	PlaybackSnapshot
	PlaybackDeskInfo
	PlaybackItem
	PlaybackReqPageByGameid
	VersionInfo
	AssetInfo
	HotupdateReqVersionInfo
	HotupdateAckVersionInfo
	HotupdateReqAssetsInfo
	HotupdateAckAssetsInfo
	HotupdateReqGameAssetsInfo
	HotupdateAckGameAssetsInfo
	NiuniuClientPoker
	NiuniuUserBill
	NiuniuDeskOption
	NiuniuSrvRoom
	NiuniuClientDesk
	NiuniuClientUser
	NiuCreateDeskReq
	NiuEnterDeskReq
	NiuEnterDeskAck
	NiuEnterDeskBc
	NiuSwitchReadyReq
	NiuSwitchReadyAck
	NiuSwitchReadyBc
	NiuStartGameOt
	NiuQiangzhuangOt
	NiuQiangzhuangReq
	NiuQiangzhuangAck
	NiuQiangzhuangResItem
	NiuQiangzhuangResBc
	NiuJiabeiOt
	NiuJiabeiReq
	NiuJiabeiAck
	NiuJiabeiBc
	NiuXuanpaiOt
	NiuXuanpaiReq
	NiuXuanpaiAck
	NiuXuanpaiBc
	NiuBipaiResultItem
	NiuBipaiResultBc
	NiuGameEnd
	NiuDeskDissolveDoneBc
	NiuOwnerDissolveReq
	NiuOwnerDissolveAck
	NiuOfflineBc
	NiuCoinRoomListReq
	NiuCoinRoomListAck
	NiuniuSrvPoker
	NiuniuSrvDesk
	NiuniuSrvUser
	NiuniuSrvDeskSnapshotIdIndex
	NiuniuSrvDeskSnapshot
	PdkBasePlayerInfo
	PdkBasePlayerRateInfo
	PdkBaseCommonRateInfo
	PdkBaseDeskInfo
	PdkReqDissolveDesk
	PdkAckDissolveDesk
	PdkReqReady
	PdkAckReady
	PdkBaseWinCoinInfo
	PdkBaCurrentResult
	PdkBaseEndLotteryInfo
	PdkBcEndLottery
	PdkReqCreateDesk
	PdkAckCreateDesk
	PdkReqGameRecord
	PdkBaseUserRecord
	PdkBaseGameRecord
	PdkAckGameRecord
	PdkReqEnterDesk
	PdkAckEnterDesk
	PdkBcOpening
	PdkBcDealCards
	PdkReqShowHandPokers
	PdkAckShowHandPokers
	PdkReqJiaoDiZhu
	PdkAckJiaoDiZhu
	PdkReqRobDiZhu
	PdkAckRobDiZhu
	PdkReqDouble
	PdkAckDouble
	PdkBcStartPlay
	PdkReqMenuZhua
	PdkAckMenuZhua
	PdkReqSeeCards
	PdkAckSeeCards
	PdkReqPull
	PdkAckPull
	PdkReqOutCards
	PdkAckOutCards
	PdkReqActGuo
	PdkAckGuoAck
	PdkOverturnPlayerInfo
	PdkBcOverTurn
	PdkBcGameInfo
	PdkSrvOutPokerPais
	PdkSrvDeskTongJi
	PdkSrvDesk
	PdkSrvGameData
	PdkSrvBillBean
	PdkSrvBill
	PdkSrvUserStatisticsRound
	PdkSrvUserStatistics
	PdkSrvUser
	PdkSrvRoom
	PdkSrvBak
	PlaybackPdkAckPage
	PdkPlayerInfo
	PdkPlaybackSnapshot
	PdkPlaybackDeskInfo
	PdkDataRecover
	PdkDataRecoverDeskIdList
	PezBase_PaiInfo
	PezBase_PlayConf
	PezBase_RoomTypeInfo
	PezBaseTimerInfo
	PezBase_PaiValue
	PezBase_PlayerCard
	PezLastScore
	PezBase_PlayerInfo
	PezBase_DeskGameInfo
	PezTuozi
	Pez_DissolveDesk
	Pez_AckDissolveDesk
	Pez_Ready
	Pez_AckReady
	Pez_EndLotteryInfo
	Pez_SendCurrentResult
	EndLottery
	Pez_SendEndLottery
	Pez_UserBean
	Pez_Bill
	PezReq_CreateRoom
	PezReq_AckCreateRoom
	PezReq_EnterRoom
	PezReq_AckEnterRoom
	Pez_Notice
	Pez_AckNotice
	Pez_GameRecord
	PezBeanUserRecord
	PezBeanGameRecord
	Pez_AckGameRecord
	Pez_Feedback
	PezUserCoinBean
	Pez_Opening
	Pez_DealCards
	Pez_Bet
	Pez_AckBet
	Pez_BCOpenPai
	Pez_SendGameInfo
	PhzPlayPaiIds
	PhzBaseCreateOption
	PhzBaseRoomInfo
	PhzBaseDeskInfo
	PhzBasePlayerInfo
	PhzReqCreateDesk
	PhzReqEnterDesk
	PhzAckDesk
	PhzReqGameInfo
	PhzDeskGameInfo
	PhzReqDissolveDesk
	PhzAckDissolveDesk
	PhzReq_Ready
	PhzAck_Ready
	PhzPlay_ReadyStatus
	PhzPlay_Opening
	PhzPlay_SendCards
	PhzReq_OutCards
	PhzAck_OutCards
	PhzPlay_MoPai
	PhzPlay_Overturn
	PhzPlay_CanPeng
	PhzReq_Peng
	PhzAck_Peng
	PhzPlay_ChiPai
	PhzPlayCanChi
	PhzReqChiPai
	PhzReqBiPai
	PhzAckBiPai
	PhzAckChiPai
	PhzPlayTiPai
	PhzPlayCanTiPai
	PhzPlayWeiPai
	PhzPlayPaoPai
	PhzPlayCanHuPai
	PhzReqHuPai
	PhzAckHuPai
	PhzReqPass
	PhzAckPass
	PhzPlayHandPokers
	PhzPlayRoundBillBean
	PhzPlayCurrentResult
	PhzPlayUserLotteryInfo
	PhzPlayUserEndResult
	PhzPlayEndLottery
	PlaybackPhzAckPage
	PhzPlaybackPlayerInfo
	PhzPlaybackSnapshot
	PhzPlaybackDeskInfo
	HelloRequest
	HelloReply
	RpcHallUpdateConfig
	HallRpcDeskEventMsg
	HallRpcGroupRefresh
	ZjhReqGetRoomList
	ZjhBaseRoomInfo
	ZjhAckRoomList
	ZjhReqEnterDesk
	ZjhBaseUserInfo
	ZjhDeskStateAck
	ZjhBcGameInfo
	ZjhBaseDeskInfo
	ZjhBcNewPlayerEnter
	ZjhReqLeave
	ZjhBcLeave
	ZjhBasePlayerInfo
	ZjhBcOpening
	ZjhReqReady
	ZjhBcReady
	ZjhBcOverTurn
	ZjhReqCheckCards
	ZjhBcCheckCards
	ZjhReqFold
	ZjhBcFold
	ZjhReqCall
	ZjhBcCall
	ZjhReqBloodFight
	ZjhBcBloodFight
	ZjhBcBloodEnd
	ZjhReqRaiseFight
	ZjhBcRaiseAck
	ZjhReqCompare
	ZjhBcCompare
	ZjhBcGameEnd
	ZjhBcDeskState
	ZjhReqDaShang
	ZjhBcDaShang
	ZjhSrvPoker
	ZjhSrvBill
	ZjhSrv_GameData
	ZjhSrvUser
	ZjhSrvDesk
	ZjhSrvRoom
*/
package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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

type BainiuEnumProtoid int32

const (
	// //////////////////////////////////
	BainiuEnumProtoid_BAINIU_PID_HEARTBEAT      BainiuEnumProtoid = 0
	BainiuEnumProtoid_BAINIU_PID_QUICK_CONN     BainiuEnumProtoid = 1
	BainiuEnumProtoid_BAINIU_PID_QUICK_CONN_ACK BainiuEnumProtoid = 2
	BainiuEnumProtoid_BAINIU_ENTER_DESK_REQ     BainiuEnumProtoid = 3
	BainiuEnumProtoid_BAINIU_ENTER_DESK_ACK     BainiuEnumProtoid = 4
	BainiuEnumProtoid_BAINIU_YAZHU_OT           BainiuEnumProtoid = 5
	BainiuEnumProtoid_BAINIU_YAZHU_REQ          BainiuEnumProtoid = 6
	BainiuEnumProtoid_BAINIU_YAZHU_ACK          BainiuEnumProtoid = 7
	BainiuEnumProtoid_BAINIU_YAZHU_BC           BainiuEnumProtoid = 8
	BainiuEnumProtoid_BAINIU_FAPAI_BC           BainiuEnumProtoid = 9
	BainiuEnumProtoid_BAINIU_BIPAI_BC           BainiuEnumProtoid = 10
	BainiuEnumProtoid_BAINIU_PID_WINCOUNT_REQ   BainiuEnumProtoid = 11
	BainiuEnumProtoid_BAINIU_PID_WINCOUNT_ACK   BainiuEnumProtoid = 12
)

var BainiuEnumProtoid_name = map[int32]string{
	0:  "BAINIU_PID_HEARTBEAT",
	1:  "BAINIU_PID_QUICK_CONN",
	2:  "BAINIU_PID_QUICK_CONN_ACK",
	3:  "BAINIU_ENTER_DESK_REQ",
	4:  "BAINIU_ENTER_DESK_ACK",
	5:  "BAINIU_YAZHU_OT",
	6:  "BAINIU_YAZHU_REQ",
	7:  "BAINIU_YAZHU_ACK",
	8:  "BAINIU_YAZHU_BC",
	9:  "BAINIU_FAPAI_BC",
	10: "BAINIU_BIPAI_BC",
	11: "BAINIU_PID_WINCOUNT_REQ",
	12: "BAINIU_PID_WINCOUNT_ACK",
}
var BainiuEnumProtoid_value = map[string]int32{
	"BAINIU_PID_HEARTBEAT":      0,
	"BAINIU_PID_QUICK_CONN":     1,
	"BAINIU_PID_QUICK_CONN_ACK": 2,
	"BAINIU_ENTER_DESK_REQ":     3,
	"BAINIU_ENTER_DESK_ACK":     4,
	"BAINIU_YAZHU_OT":           5,
	"BAINIU_YAZHU_REQ":          6,
	"BAINIU_YAZHU_ACK":          7,
	"BAINIU_YAZHU_BC":           8,
	"BAINIU_FAPAI_BC":           9,
	"BAINIU_BIPAI_BC":           10,
	"BAINIU_PID_WINCOUNT_REQ":   11,
	"BAINIU_PID_WINCOUNT_ACK":   12,
}

func (x BainiuEnumProtoid) Enum() *BainiuEnumProtoid {
	p := new(BainiuEnumProtoid)
	*p = x
	return p
}
func (x BainiuEnumProtoid) String() string {
	return proto.EnumName(BainiuEnumProtoid_name, int32(x))
}
func (x *BainiuEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuEnumProtoid_value, data, "BainiuEnumProtoid")
	if err != nil {
		return err
	}
	*x = BainiuEnumProtoid(value)
	return nil
}
func (BainiuEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 押注区域
type BainiuAreaName int32

const (
	BainiuAreaName_BAINIU_AREA_BANKER BainiuAreaName = 0
	BainiuAreaName_BAINIU_AREA_TIAN   BainiuAreaName = 1
	BainiuAreaName_BAINIU_AREA_DI     BainiuAreaName = 2
	BainiuAreaName_BAINIU_AREA_XUAN   BainiuAreaName = 3
	BainiuAreaName_BAINIU_AREA_HUANG  BainiuAreaName = 4
)

var BainiuAreaName_name = map[int32]string{
	0: "BAINIU_AREA_BANKER",
	1: "BAINIU_AREA_TIAN",
	2: "BAINIU_AREA_DI",
	3: "BAINIU_AREA_XUAN",
	4: "BAINIU_AREA_HUANG",
}
var BainiuAreaName_value = map[string]int32{
	"BAINIU_AREA_BANKER": 0,
	"BAINIU_AREA_TIAN":   1,
	"BAINIU_AREA_DI":     2,
	"BAINIU_AREA_XUAN":   3,
	"BAINIU_AREA_HUANG":  4,
}

func (x BainiuAreaName) Enum() *BainiuAreaName {
	p := new(BainiuAreaName)
	*p = x
	return p
}
func (x BainiuAreaName) String() string {
	return proto.EnumName(BainiuAreaName_name, int32(x))
}
func (x *BainiuAreaName) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuAreaName_value, data, "BainiuAreaName")
	if err != nil {
		return err
	}
	*x = BainiuAreaName(value)
	return nil
}
func (BainiuAreaName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// =================================百人牛牛牌型================================
// 牛牛牌的类型
type BainiuEnum_PokerType int32

const (
	BainiuEnum_PokerType_BAINIU_NO_NIU           BainiuEnum_PokerType = 1
	BainiuEnum_PokerType_BAINIU_NIU_1            BainiuEnum_PokerType = 2
	BainiuEnum_PokerType_BAINIU_NIU_2            BainiuEnum_PokerType = 3
	BainiuEnum_PokerType_BAINIU_NIU_3            BainiuEnum_PokerType = 4
	BainiuEnum_PokerType_BAINIU_NIU_4            BainiuEnum_PokerType = 5
	BainiuEnum_PokerType_BAINIU_NIU_5            BainiuEnum_PokerType = 6
	BainiuEnum_PokerType_BAINIU_NIU_6            BainiuEnum_PokerType = 7
	BainiuEnum_PokerType_BAINIU_NIU_7            BainiuEnum_PokerType = 8
	BainiuEnum_PokerType_BAINIU_NIU_8            BainiuEnum_PokerType = 9
	BainiuEnum_PokerType_BAINIU_NIU_9            BainiuEnum_PokerType = 10
	BainiuEnum_PokerType_BAINIU_NIU_NIU          BainiuEnum_PokerType = 11
	BainiuEnum_PokerType_BAINIU_YIN_NIU          BainiuEnum_PokerType = 12
	BainiuEnum_PokerType_BAINIU_JIN_NIU          BainiuEnum_PokerType = 13
	BainiuEnum_PokerType_BAINIU_WU_XIAO_NIU      BainiuEnum_PokerType = 14
	BainiuEnum_PokerType_BAINIU_NIU_ZHA_DAN      BainiuEnum_PokerType = 15
	BainiuEnum_PokerType_BAINIU_NIU_YI_TIAO_LONG BainiuEnum_PokerType = 16
)

var BainiuEnum_PokerType_name = map[int32]string{
	1:  "BAINIU_NO_NIU",
	2:  "BAINIU_NIU_1",
	3:  "BAINIU_NIU_2",
	4:  "BAINIU_NIU_3",
	5:  "BAINIU_NIU_4",
	6:  "BAINIU_NIU_5",
	7:  "BAINIU_NIU_6",
	8:  "BAINIU_NIU_7",
	9:  "BAINIU_NIU_8",
	10: "BAINIU_NIU_9",
	11: "BAINIU_NIU_NIU",
	12: "BAINIU_YIN_NIU",
	13: "BAINIU_JIN_NIU",
	14: "BAINIU_WU_XIAO_NIU",
	15: "BAINIU_NIU_ZHA_DAN",
	16: "BAINIU_NIU_YI_TIAO_LONG",
}
var BainiuEnum_PokerType_value = map[string]int32{
	"BAINIU_NO_NIU":           1,
	"BAINIU_NIU_1":            2,
	"BAINIU_NIU_2":            3,
	"BAINIU_NIU_3":            4,
	"BAINIU_NIU_4":            5,
	"BAINIU_NIU_5":            6,
	"BAINIU_NIU_6":            7,
	"BAINIU_NIU_7":            8,
	"BAINIU_NIU_8":            9,
	"BAINIU_NIU_9":            10,
	"BAINIU_NIU_NIU":          11,
	"BAINIU_YIN_NIU":          12,
	"BAINIU_JIN_NIU":          13,
	"BAINIU_WU_XIAO_NIU":      14,
	"BAINIU_NIU_ZHA_DAN":      15,
	"BAINIU_NIU_YI_TIAO_LONG": 16,
}

func (x BainiuEnum_PokerType) Enum() *BainiuEnum_PokerType {
	p := new(BainiuEnum_PokerType)
	*p = x
	return p
}
func (x BainiuEnum_PokerType) String() string {
	return proto.EnumName(BainiuEnum_PokerType_name, int32(x))
}
func (x *BainiuEnum_PokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuEnum_PokerType_value, data, "BainiuEnum_PokerType")
	if err != nil {
		return err
	}
	*x = BainiuEnum_PokerType(value)
	return nil
}
func (BainiuEnum_PokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 打出去的牌
type BainiuClientPoker struct {
	Pais             []*ClientBasePoker    `protobuf:"bytes,1,rep,name=pais" json:"pais,omitempty"`
	Type             *BainiuEnum_PokerType `protobuf:"varint,2,opt,name=type,enum=ddproto.BainiuEnum_PokerType" json:"type,omitempty"`
	Area             *BainiuAreaName       `protobuf:"varint,3,opt,name=area,enum=ddproto.BainiuAreaName" json:"area,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *BainiuClientPoker) Reset()                    { *m = BainiuClientPoker{} }
func (m *BainiuClientPoker) String() string            { return proto.CompactTextString(m) }
func (*BainiuClientPoker) ProtoMessage()               {}
func (*BainiuClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BainiuClientPoker) GetPais() []*ClientBasePoker {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *BainiuClientPoker) GetType() BainiuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return BainiuEnum_PokerType_BAINIU_NO_NIU
}

func (m *BainiuClientPoker) GetArea() BainiuAreaName {
	if m != nil && m.Area != nil {
		return *m.Area
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func init() {
	proto.RegisterType((*BainiuClientPoker)(nil), "ddproto.bainiu_client_poker")
	proto.RegisterEnum("ddproto.BainiuEnumProtoid", BainiuEnumProtoid_name, BainiuEnumProtoid_value)
	proto.RegisterEnum("ddproto.BainiuAreaName", BainiuAreaName_name, BainiuAreaName_value)
	proto.RegisterEnum("ddproto.BainiuEnum_PokerType", BainiuEnum_PokerType_name, BainiuEnum_PokerType_value)
}

func init() { proto.RegisterFile("bainiu_base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x6f, 0xda, 0x30,
	0x18, 0xc7, 0x49, 0x48, 0x4b, 0xfb, 0x40, 0xe9, 0x83, 0x29, 0x1b, 0x74, 0xda, 0x54, 0xed, 0x54,
	0x21, 0x0d, 0x69, 0xec, 0xfd, 0x68, 0x20, 0x2b, 0x1e, 0x93, 0x43, 0xb3, 0x44, 0x7d, 0xb9, 0x58,
	0x69, 0xc9, 0x01, 0x6d, 0x90, 0x88, 0xd2, 0x43, 0x8f, 0xfb, 0x36, 0xfb, 0x42, 0x3b, 0xec, 0xdb,
	0x4c, 0x76, 0xbc, 0xca, 0x73, 0xe9, 0xc1, 0x52, 0xf4, 0xfb, 0xbf, 0xd8, 0x7e, 0x1c, 0x68, 0x5c,
	0x25, 0xf3, 0xe5, 0xfc, 0x56, 0x5c, 0x25, 0x37, 0x69, 0x2f, 0x5f, 0x65, 0xeb, 0x8c, 0x54, 0x66,
	0x33, 0xf5, 0x71, 0xd8, 0xbc, 0xce, 0x16, 0x8b, 0x6c, 0x29, 0xae, 0x7f, 0xcc, 0xd3, 0xe5, 0xba,
	0x50, 0x5f, 0xfe, 0x72, 0xa0, 0xa9, 0x33, 0x05, 0x17, 0x79, 0xf6, 0x3d, 0x5d, 0x91, 0x1e, 0x78,
	0x79, 0x32, 0xbf, 0x69, 0x3b, 0x47, 0xe5, 0xe3, 0x6a, 0xff, 0xb0, 0xa7, 0x4b, 0x7a, 0xda, 0x24,
	0xfb, 0x0b, 0x67, 0xa8, 0x7c, 0xa4, 0x0f, 0xde, 0xfa, 0x2e, 0x4f, 0xdb, 0xee, 0x91, 0x73, 0x5c,
	0xef, 0xbf, 0xb8, 0xf7, 0xeb, 0xee, 0x74, 0x79, 0xbb, 0x10, 0x53, 0xe9, 0x8f, 0xee, 0xf2, 0x34,
	0x54, 0x5e, 0xf2, 0x0a, 0xbc, 0x64, 0x95, 0x26, 0xed, 0xb2, 0xca, 0x74, 0xec, 0x8c, 0xd4, 0xc4,
	0x32, 0x59, 0xa4, 0xa1, 0xb2, 0x75, 0xff, 0xb8, 0xf7, 0x47, 0x55, 0x75, 0xca, 0x3c, 0x9f, 0x91,
	0x36, 0x1c, 0x0c, 0x28, 0xe3, 0x2c, 0x16, 0x53, 0x36, 0x12, 0x63, 0x9f, 0x86, 0xd1, 0xc0, 0xa7,
	0x11, 0x96, 0x48, 0x07, 0x5a, 0x86, 0x72, 0x1a, 0xb3, 0xe1, 0x44, 0x0c, 0x03, 0xce, 0xd1, 0x21,
	0xcf, 0xa1, 0xb3, 0x51, 0x12, 0x74, 0x38, 0x41, 0xd7, 0x48, 0xfa, 0x3c, 0xf2, 0x43, 0x31, 0xf2,
	0xbf, 0x4d, 0x44, 0xe8, 0x9f, 0x62, 0x79, 0xb3, 0x24, 0x53, 0x1e, 0x69, 0xc2, 0xbe, 0x96, 0x2e,
	0xe8, 0xe5, 0x38, 0x16, 0x41, 0x84, 0x5b, 0xe4, 0x00, 0xf0, 0x3f, 0x28, 0x5b, 0xb6, 0x1f, 0x50,
	0x59, 0x50, 0x79, 0x50, 0x30, 0x18, 0xe2, 0x8e, 0x01, 0x3f, 0xd3, 0x29, 0x65, 0x12, 0xee, 0x1a,
	0x70, 0xc0, 0x34, 0x04, 0xf2, 0x0c, 0x9e, 0x1a, 0x97, 0x3a, 0x63, 0x7c, 0x18, 0xc4, 0x3c, 0x52,
	0x3b, 0x56, 0x1f, 0x13, 0xe5, 0xc6, 0xb5, 0xee, 0x4f, 0x07, 0xd0, 0x1e, 0x3b, 0x79, 0x02, 0x44,
	0x27, 0x68, 0xe8, 0x53, 0x31, 0xa0, 0x7c, 0xe2, 0x87, 0x58, 0x32, 0xce, 0xae, 0x78, 0xc4, 0xa8,
	0x9c, 0x28, 0x81, 0xba, 0x49, 0x47, 0x0c, 0x5d, 0xdb, 0x79, 0x1e, 0x53, 0x8e, 0x65, 0xd2, 0x82,
	0x86, 0x49, 0xc7, 0x31, 0xe5, 0x27, 0xe8, 0x75, 0x7f, 0xbb, 0xd0, 0xda, 0xf8, 0xbb, 0x90, 0x06,
	0xec, 0xe9, 0x00, 0x0f, 0x04, 0x67, 0x31, 0x3a, 0x04, 0xa1, 0xf6, 0x0f, 0xb1, 0x58, 0xbc, 0x46,
	0xd7, 0x22, 0x7d, 0x2c, 0x5b, 0xe4, 0x0d, 0x7a, 0x16, 0x79, 0x8b, 0x5b, 0x16, 0x79, 0x87, 0xdb,
	0x16, 0x79, 0x8f, 0x15, 0x8b, 0x7c, 0xc0, 0x1d, 0x8b, 0x7c, 0xc4, 0x5d, 0x8b, 0x7c, 0x42, 0x30,
	0xe6, 0xa1, 0x17, 0x56, 0x0d, 0x76, 0xc1, 0xb8, 0x62, 0x35, 0x83, 0x7d, 0xd1, 0x6c, 0xcf, 0x98,
	0xfc, 0x59, 0x2c, 0xce, 0x19, 0x2d, 0x6e, 0x5d, 0x37, 0xb8, 0x5c, 0x97, 0x63, 0x2a, 0x46, 0x94,
	0xe3, 0xbe, 0xf1, 0xb6, 0x45, 0xb7, 0x7c, 0x93, 0x40, 0x7c, 0x0d, 0xf8, 0x09, 0xe2, 0xb4, 0xf4,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xac, 0x44, 0x51, 0x16, 0x04, 0x00, 0x00,
}
