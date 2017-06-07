package NiuDeskRoundDao

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//查询总战绩
func GetNiuDeskRoundByUserId(userId uint32) []model.T_niuniu_desk_round {
	var deskRecords []model.T_niuniu_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_NIU_DESK_ROUND_ALL).Find(bson.M{"userids": bson.RegEx{querKey, "."}}).Sort("-endtime").Limit(20).All(&deskRecords)
	})

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]牛牛相关的战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}

//查询牌桌内战绩
func GetNiuDeskRoundByDeskId(userId uint32, deskId int32) []model.T_niuniu_desk_round {
	var deskRecords []model.T_niuniu_desk_round
	querKey, _ := numUtils.Uint2String(userId)
	db.Query("", func(d *mgo.Database) {
		d.C(tableName.DBT_NIU_DESK_ROUND_ONE).Find(bson.M{
			"userids": bson.RegEx{querKey, "."},
			"deskid":  deskId,
		}).Sort("-gamenumber").Limit(20).All(&deskRecords)
	})

	if deskRecords == nil || len(deskRecords) <= 0 {
		log.T("没有找到玩家[%v]牛牛相关的牌桌内战绩...", userId)
		return nil
	} else {
		return deskRecords
	}
}
