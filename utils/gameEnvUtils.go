package utils

import (
	"casino_common/proto/ddproto"
	//"casino_common/common/log"
	"casino_common/common/Error"
	"casino_common/common/log"
	"fmt"
)

/*

	gid : 游戏id
	roomType : 金币场 还是朋友桌
*/
func GetGameName(gid int32, roomType int32) string {
	log.T("查询游戏的名字 gid:%v roomType %v", gid, roomType)
	if gid == int32(ddproto.CommonEnumGame_GID_TEXAS) {
		return "德州扑克"
	} else if gid == int32(ddproto.CommonEnumGame_GID_MAHJONG) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "麻将金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "麻将朋友桌"
		} else {
			return "麻将"
		}
	} else if gid == int32(ddproto.CommonEnumGame_GID_DDZ) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "斗地主金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "斗地主朋友桌"
		} else {
			return "斗地主"
		}
	} else if gid == int32(ddproto.CommonEnumGame_GID_ZJH) {
		return "炸金花"
	} else if gid == int32(ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN) {
		return "牛牛朋友桌"
	} else if gid == int32(ddproto.CommonEnumGame_GID_PDK) {
		if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN) {
			return "跑得快金币场"
		} else if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
			return "跑得快朋友桌"
		} else {
			return "跑得快"
		}
	}
	return "其他"
}

func GetGameAndTypeName(gid, roomType int32) string {
	gameName := ""
	switch gid {
	case int32(ddproto.CommonEnumGame_GID_TEXAS):
		gameName = "德州扑克"
	case int32(ddproto.CommonEnumGame_GID_MJ_HAINAN):
		gameName = "海南麻将"
	case int32(ddproto.CommonEnumGame_GID_ZJH):
		gameName = "炸金花"
	case int32(ddproto.CommonEnumGame_GID_BAIRENNIUNIU):
		gameName = "百人牛牛"
	case int32(ddproto.CommonEnumGame_GID_PDK):
		gameName = "跑得快"
	case int32(ddproto.CommonEnumGame_GID_MAHJONG):
		gameName = "四川麻将"
	case int32(ddproto.CommonEnumGame_GID_ZXZ):
		gameName = "捉虾子"
	case int32(ddproto.CommonEnumGame_GID_MJBAISHAN):
		gameName = "白山麻将"
	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		gameName = "红中麻将"
	case int32(ddproto.CommonEnumGame_GID_DDZ):
		gameName = "斗地主"
	case int32(ddproto.CommonEnumGame_GID_MJ_SONGJIANGHE):
		gameName = "松江河"
	case int32(ddproto.CommonEnumGame_GID_PAOHUZI):
		gameName = "跑胡子"
	case int32(ddproto.CommonEnumGame_GID_PAOYAO):
		gameName = "刨幺"
	case int32(ddproto.CommonEnumGame_GID_CHANGBAI):
		gameName = "长白麻将"
	case int32(ddproto.CommonEnumGame_GID_ZHADAN):
		gameName = "炸弹"
	case int32(ddproto.CommonEnumGame_GID_MJ_SHENQI):
		gameName = "神奇麻将"
	case int32(ddproto.CommonEnumGame_GID_PHZ_SHAOYANGBOPI):
		gameName = "邵阳剥皮"
	case int32(ddproto.CommonEnumGame_GID_MJ_CHANGSHA):
		gameName = "长沙麻将"
	case int32(ddproto.CommonEnumGame_GID_MJ_ZIGONG):
		gameName = "自贡麻将"

		//未用到的gameId
	case int32(ddproto.CommonEnumGame_GID_MJCHANGCHUN):
		gameName = "长春麻将"
	case int32(ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN):
		gameName = "经典牛牛"
	case int32(ddproto.CommonEnumGame_GID_BANTUOZI):
		gameName = "搬坨子"
	case int32(ddproto.CommonEnumGame_GID_DOUDIZHUERREN):
		gameName = "二人斗地主"
	case int32(ddproto.CommonEnumGame_GIDDOUDIZHUJINGDIANDONGBEI):
		gameName = "东北经典斗地主"
	case int32(ddproto.CommonEnumGame_GID_TIANDAKENG):
		gameName = "填大坑"
	case int32(ddproto.CommonEnumGame_GID_ZHIZUNWUZHANG):
		gameName = "自助五张"
	case int32(ddproto.CommonEnumGame_GID_MJ_YIBIN):
		gameName = "宜宾麻将"
	case int32(ddproto.CommonEnumGame_GID_MJ_PENGZHOU):
		gameName = "彭州麻将"
	case int32(ddproto.CommonEnumGame_GID_SANDAYI):
		gameName = "三打一"

	default:
		gameName = "其他"
	}

	roomTypeName := ""
	switch roomType {
	case int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND):
		roomTypeName = "朋友桌"
	case int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN):
		roomTypeName = "金币场"
	default:

	}

	return gameName + roomTypeName
}

func GetEnterError(gid, roomType int32, password string) error {
	msg := fmt.Sprintf("【%v】游戏还没有结束，请完成后再来", GetGameAndTypeName(gid, roomType))
	retErr := Error.NewError(int32(ddproto.COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM_ERROR_EC_OTHER_LV_DESK_GAMING), msg)
	if roomType == int32(ddproto.COMMON_ENUM_ROOMTYPE_DESK_FRIEND) {
		//朋友桌 需要返回桌号
		retErr = Error.NewError(Error.GetErrorCode(retErr), fmt.Sprintf("%s 桌号【%s】", Error.GetErrorMsg(retErr), password))
	}
	return retErr
}
