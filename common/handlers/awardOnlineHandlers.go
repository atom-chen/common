package handlers

import (
	"casino_common/common/cfg"
	"casino_common/common/consts"
	"casino_common/common/service/awardService"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
)

//在线奖励的处理
func HandlerAwardOnline(args []interface{}) {
	m := args[0].(*ddproto.AwardReqOnline)
	a := args[1].(gate.Agent)
	//处理奖励
	awardService.DoAwardOnline(m.GetHeader().GetUserId(), a)
}

//在线奖励的信息
func HandlerAwardOnlineInfo(args []interface{}) {
	m := args[0].(*ddproto.AwardReqOnlineInfo)
	a := args[1].(gate.Agent)
	//处理奖励
	awardService.GetOnlineInfo(m.GetHeader().GetUserId(), a)
}

//获取新手奖励
func HandlerNewUserAward(args []interface{}) {
	m := args[0].(*ddproto.AwardReqGetNewUser)
	a := args[1].(gate.Agent)
	userId := m.GetHeader().GetUserId()
	//返回信息
	ack := &ddproto.AwardAckGetNewUser{
		Header: &ddproto.ProtoHeader{
			UserId: proto.Uint32(userId),
			Error:  proto.String(fmt.Sprintf("恭喜你，成功领取 %.2f红包 分享奖励。", cfg.GetNewUserAward())),
			Code:   proto.Int32(consts.ACK_RESULT_SUCC),
		},
	}
	user := userService.GetUserById(userId)
	if !user.GetNewUserAward() {
		return
	}

	//增加红包信息处理新手奖励
	_, err := userService.INCRUserBonus(userId, cfg.GetNewUserAward(), "领取新手奖励")
	if err != nil {
		ack.Header.Code = proto.Int32(consts.ACK_RESULT_ERROR)
		ack.Header.Error = proto.String("领取分享奖励失败。")
		a.WriteMsg(ack)
		return
	}

	//更新redis和mgo
	user = userService.GetUserById(userId)
	user.NewUserAward = proto.Bool(false)
	userService.UpdateUser2Mgo(user)

	//领取成功,返回信息
	a.WriteMsg(ack)
}
