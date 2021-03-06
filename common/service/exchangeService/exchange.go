package exchangeService

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//审核状态
type ExchangeState int32

const (
	PROCESS_ING    ExchangeState = 1 //"审核中"
	PROCESS_FALSE  ExchangeState = 2 //"申请被拒"
	PROCESS_TRUE   ExchangeState = 3 //"申请通过"
	PROCESS_SENDED ExchangeState = 4 //"已发放"
)

func (s ExchangeState) Name() string {
	switch int32(s) {
	case 1:
		return "待审核"
	case 2:
		return "已被拒"
	case 3:
		return "已通过"
	case 4:
		return "已发放"
	}
	return "未知状态"
}

//兑换记录
type ExchangeRecord struct {
	Id          bson.ObjectId             `bson:"_id"` //编号
	UserId      uint32                    //用户id
	Type        ddproto.HallEnumTradeType //兑换商品的类型
	Money       float64                   //兑换数量
	Name        string                    //姓名
	Phone       string                    //电话
	WxNumber    string                    //微信号码
	Address     string                    //收货地址
	Status      ExchangeState             //审核状态
	RequestTime time.Time                 //申请时间
	ProcessTime time.Time                 //审核时间
}

//新兑换记录
func NewRecord(userId uint32, goods_type ddproto.HallEnumTradeType, amount float64) (*ExchangeRecord, error) {
	user := userService.GetUserById(userId)
	if user.GetRealName() == "" || user.GetPhoneNumber() == "" || user.GetWxNumber() == "" {
		return nil, errors.New("信息填写不完整，兑换申请失败！")
	}
	if goods_type > 300 && goods_type < 400 {
		//实物兑换
		if user.GetRealAddress() == "" {
			return nil, errors.New("请填写收货地址！")
		}
		_, err_dec := userService.DECUserTicket(userId, int32(amount), "商城实物兑换扣奖券")
		if err_dec != nil {
			return nil, errors.New("您的奖券余额不足，兑换失败！")
		}

	} else if goods_type == ddproto.HallEnumTradeType_TRADE_BONUS {
		//扣除红包
		_, err_dec := userService.DECUserBonus(userId, amount, "商城实物兑换扣红包")
		if err_dec != nil {
			return nil, errors.New("您的红包余额不足，兑换失败！")
		}
	}

	new_record := &ExchangeRecord{
		Type:     goods_type,
		Money:    amount,
		UserId:   userId,
		Name:     user.GetRealName(),
		Phone:    user.GetPhoneNumber(),
		WxNumber: user.GetWxNumber(),
		Address:  user.GetRealAddress(),
	}
	return new_record.Insert()
}

//插入新行
func (r *ExchangeRecord) Insert() (*ExchangeRecord, error) {
	r.Id = bson.NewObjectId()
	r.Status = PROCESS_ING
	r.RequestTime = time.Now()
	err := db.C(tableName.DBT_ADMIN_EXCHANGE_RECORD).Insert(r)
	return r, err
}

//更新记录
func (r *ExchangeRecord) Save() error {
	return db.C(tableName.DBT_ADMIN_EXCHANGE_RECORD).Update(bson.M{
		"_id": r.Id,
	}, r)
}

//通过id获取一条记录
func GetRecordById(id string) *ExchangeRecord {
	row := new(ExchangeRecord)
	db.C(tableName.DBT_ADMIN_EXCHANGE_RECORD).Find(bson.M{
		"_id": bson.ObjectIdHex(id),
	}, row)
	return row
}
