package model

import (
	"casino_common/proto/ddproto"
	"casino_common/utils/numUtils"
	"casino_common/utils/timeUtils"
	"github.com/golang/protobuf/proto"
	"time"
)

//一把结束,战绩可以通过这个表来查询
type T_zhadan_desk_round struct {
	DeskId     int32  //房间Id
	Password   string //房间号
	GameNumber int32  //一局游戏编号
	UserIds    string
	BeginTime  time.Time
	EndTime    time.Time
	TotalRound int32
	Records    []NiuRecordBean
}

func (t T_zhadan_desk_round) TransRecord() *ddproto.BeanGameRecord {
	result := &ddproto.BeanGameRecord{
		BeginTime: proto.String(timeUtils.Format(t.BeginTime)),
		DeskId:    proto.Int32(t.DeskId),
		Password:  proto.String(t.Password),
		Id:        proto.Int32(t.GameNumber),
		RoundStr:  proto.String(numUtils.Int2String2(t.TotalRound)),
	}

	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

func (t T_zhadan_desk_round) TransRecordCoin() *ddproto.BeanGameRecordCoin {
	result := &ddproto.BeanGameRecordCoin{
		BeginTime: proto.String(timeUtils.Format(t.EndTime)),
		DeskId:    proto.Int32(t.DeskId),
		Password:  proto.String(t.Password),
		Id:        proto.Int32(t.GameNumber),
		Users:     []*ddproto.BeanUserRecord{},
	}

	for _, bean := range t.Records {
		b := bean.TransBeanUserRecord()
		result.Users = append(result.Users, b)
	}
	return result
}

type ZhadanRecordBean struct {
	UserId    uint32
	NickName  string
	WinAmount int64
}

func (b ZhadanRecordBean) TransBeanUserRecord() *ddproto.BeanUserRecord {
	result := &ddproto.BeanUserRecord{
		NickName:  proto.String(b.NickName),
		UserId:    proto.Uint32(b.UserId),
		WinAmount: proto.Int64(b.WinAmount),
	}
	return result
}
