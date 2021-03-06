package task

import (
	"testing"
	"casino_common/common/sys"
	db_init "casino_common/common/db"
	"casino_common/proto/ddproto"
	"log"
	"casino_common/common/service/taskService"
	"casino_common/common/service/taskService/taskType"
	"casino_common/common/service/countService/gameCounter"
	"casino_common/common/service/countService/countType"
	"github.com/golang/protobuf/proto"
	"casino_common/common/service/countService"
	"time"
	"casino_common/common/service/countService/countModel"
)

func init() {
	db_init.InitMongoDb("192.168.199.200", 27017, "test", "id",[]string{})
	sys.InitRedis("192.168.199.200:6379","test")

}

func InitBonusTask() {
	taskService.RegistTask(taskType.Task{
		TaskInfo: taskService.GetTaskInfo(103),
		ValidateFun: func(user_task *taskType.UserTask) {
			counter := gameCounter.GetUserCounterByType(user_task.UserId, user_task.TaskType)
			user_task.SumNo = counter.DayCount - user_task.LastCheckSumNo
			if user_task.SumNo >= user_task.TaskSum {
				user_task.SumNo = user_task.TaskSum
				user_task.IsDone = true
			}
			user_task.SetUserState(user_task.UserId, user_task.TaskState)
		},
		GetRewardFun: func(user_task *taskType.UserTask) []*ddproto.HallBagItem {
			return []*ddproto.HallBagItem{
				&ddproto.HallBagItem{
					Type: ddproto.HallEnumTradeType_TRADE_COIN.Enum(),
					Amount: proto.Float64(12),
				},
			}
		},
		AfterRewardFun: func(user_task *taskType.UserTask) {

		},
	})
}

func ShowUserTask(user_id uint32, task_id int32) {
	user_task := taskService.GetUserTask(user_id, task_id)
	if user_task != nil {
		//user_task.ValidateFun(user_task)
		log.Println("isDone", user_task.IsDone,"isCheck", user_task.IsCheck,"sumNo",user_task.SumNo,"lastCheckSumNo",user_task.LastCheckSumNo,"repeatNo",user_task.RepeatNo, "data", user_task.Data)
	}else {
		log.Println("用户", user_id, "的任务", task_id, "不存在。")
	}
}

//测试：红包任务
func TestBonusTask(t *testing.T)  {
	var user_id uint32 = 11
	var task_id int32 = 103
	var count_type countType.CountType = countType.ALL_GAME_WIN
	InitBonusTask()

	for i:=0;i<120;i++ {
		log.Println(i)
		ShowUserTask(user_id, task_id)
		gameCounter.Add(user_id, count_type, 1)
		taskService.OnTask(count_type, user_id)
		ShowUserTask(user_id, task_id)
		err,name,_ := taskService.CheckTaskReward(user_id, task_id)
		if err != nil {
			log.Println(err.Error())
		}else {
			log.Println(name)
		}
		ShowUserTask(user_id, task_id)
	}
	log.Println("end")
}

//测试：regist task
func TestInitTask(t *testing.T) {
	InitTask()
	row := countService.T_game_log{
		UserId: 11,
		GameId:ddproto.CommonEnumGame_GID_ZJH,
		GameNumber: 323541,
		RoomType: ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN,
		RoomLevel:1,
		Bill: 640,
		IsWine: true,
		StartTime:time.Now().Unix() - int64(10*60*time.Second),
		EndTime: time.Now().Unix(),
	}
	t.Log(gameCounter.GetDayCounter(11))
	ShowUserTask(11, 221)
	row.DoCountAndTask()
	//尝试领取
	err,name,_ := taskService.CheckTaskReward(11, 221)
	t.Log(err, name)
	t.Log(gameCounter.GetDayCounter(11))
	ShowUserTask(11, 221)
}


func TestNearTask(t *testing.T) {
	InitTask()
	task := taskService.GetUserNearBonusTask(112, ddproto.CommonEnumGame_GID_ZJH)
	t.Log(task)
}

//测试：新版红包任务
func TestNewBonusTask(t *testing.T) {
	InitTask()
	for i:= 0; i< 20; i++ {
		log.Println(i)
		ShowUserTask(11, 203)
		row := countService.T_game_log{
			UserId: 11,
			GameId:ddproto.CommonEnumGame_GID_DDZ,
			GameNumber: 323541,
			RoomType: ddproto.COMMON_ENUM_ROOMTYPE_DESK_COIN,
			RoomLevel:1,
			Bill: 640,
			CoinFee: 2000,
			IsWine: true,
			StartTime:time.Now().Unix() - int64(10 * 60 * time.Second),
			EndTime: time.Now().Unix(),
		}
		row.DoCountAndTask()
		ShowUserTask(11, 203)
		err, name,_ := taskService.CheckTaskReward(11, 203)
		log.Println(err, name)
		ShowUserTask(11, 203)
	}
}

//测试：红包任务奖励计算
func TestBonusTaskCompute(t *testing.T) {
	user_coin_fee := float64(countModel.GetUserDayCoinFeeByGameId(11, ddproto.CommonEnumGame_GID_ZJH).DayCount)
	conf := GetBonusTaskComputeConfigByGameId(ddproto.CommonEnumGame_GID_ZJH)
	reward := (user_coin_fee/10000)*conf.CoinPercent*conf.BonusPercent

	t.Log(reward)
	if reward < conf.Min {
		reward = conf.Min
	}else if reward > conf.Max {
		reward = conf.Max
	}

	t.Log(user_coin_fee)
	t.Log(conf)
	t.Log(reward)
}

//测试：再赢多少局可以领红包
func TestNearBonusNum(t *testing.T) {
	//fill_game := ddproto.CommonEnumGame_GID_ZJH
	InitTask()
	var user_id uint32 = 11
	user_task := taskService.GetUserTask(user_id, 203)
	t.Log(user_task)
	ShowUserTask(user_id, 201)
}
