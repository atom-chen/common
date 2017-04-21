// Code generated by protoc-gen-go.
// source: niuniu_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type NiuniuEnumProtoid int32

const (
	// //////////////////////////////////
	NiuniuEnumProtoid_NIU_PID_HEARTBEAT       NiuniuEnumProtoid = 0
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN      NiuniuEnumProtoid = 1
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN_ACK  NiuniuEnumProtoid = 2
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN      NiuniuEnumProtoid = 3
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN_ACK  NiuniuEnumProtoid = 4
	NiuniuEnumProtoid_NIU_PID_CREATE_DESK_REQ NiuniuEnumProtoid = 5
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_REQ  NiuniuEnumProtoid = 6
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_ACK  NiuniuEnumProtoid = 7
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_BC   NiuniuEnumProtoid = 8
	NiuniuEnumProtoid_NIU_PID_READY_REQ       NiuniuEnumProtoid = 9
	NiuniuEnumProtoid_NIU_PID_READY_ACK       NiuniuEnumProtoid = 10
	NiuniuEnumProtoid_NIU_PID_READY_BC        NiuniuEnumProtoid = 11
	NiuniuEnumProtoid_NIU_PID_START_GAME_OT   NiuniuEnumProtoid = 12
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_OT  NiuniuEnumProtoid = 13
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_REQ NiuniuEnumProtoid = 14
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_ACK NiuniuEnumProtoid = 15
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_BC  NiuniuEnumProtoid = 16
	NiuniuEnumProtoid_NIU_PID_JIABEI_OT       NiuniuEnumProtoid = 17
	NiuniuEnumProtoid_NIU_PID_JIABEI_REQ      NiuniuEnumProtoid = 18
	NiuniuEnumProtoid_NIU_PID_JIABEI_ACK      NiuniuEnumProtoid = 19
	NiuniuEnumProtoid_NIU_PID_JIABEI_BC       NiuniuEnumProtoid = 20
	NiuniuEnumProtoid_NIU_PID_BIPAI_RESULT_BC NiuniuEnumProtoid = 21
	NiuniuEnumProtoid_NIU_PID_GAME_END_BC     NiuniuEnumProtoid = 22
)

var NiuniuEnumProtoid_name = map[int32]string{
	0:  "NIU_PID_HEARTBEAT",
	1:  "NIU_PID_QUICK_CONN",
	2:  "NIU_PID_QUICK_CONN_ACK",
	3:  "NIU_PID_GAME_LOGIN",
	4:  "NIU_PID_GAME_LOGIN_ACK",
	5:  "NIU_PID_CREATE_DESK_REQ",
	6:  "NIU_PID_ENTER_DESK_REQ",
	7:  "NIU_PID_ENTER_DESK_ACK",
	8:  "NIU_PID_ENTER_DESK_BC",
	9:  "NIU_PID_READY_REQ",
	10: "NIU_PID_READY_ACK",
	11: "NIU_PID_READY_BC",
	12: "NIU_PID_START_GAME_OT",
	13: "NIU_PID_QIANGZHUANG_OT",
	14: "NIU_PID_QIANGZHUANG_REQ",
	15: "NIU_PID_QIANGZHUANG_ACK",
	16: "NIU_PID_QIANGZHUANG_BC",
	17: "NIU_PID_JIABEI_OT",
	18: "NIU_PID_JIABEI_REQ",
	19: "NIU_PID_JIABEI_ACK",
	20: "NIU_PID_JIABEI_BC",
	21: "NIU_PID_BIPAI_RESULT_BC",
	22: "NIU_PID_GAME_END_BC",
}
var NiuniuEnumProtoid_value = map[string]int32{
	"NIU_PID_HEARTBEAT":       0,
	"NIU_PID_QUICK_CONN":      1,
	"NIU_PID_QUICK_CONN_ACK":  2,
	"NIU_PID_GAME_LOGIN":      3,
	"NIU_PID_GAME_LOGIN_ACK":  4,
	"NIU_PID_CREATE_DESK_REQ": 5,
	"NIU_PID_ENTER_DESK_REQ":  6,
	"NIU_PID_ENTER_DESK_ACK":  7,
	"NIU_PID_ENTER_DESK_BC":   8,
	"NIU_PID_READY_REQ":       9,
	"NIU_PID_READY_ACK":       10,
	"NIU_PID_READY_BC":        11,
	"NIU_PID_START_GAME_OT":   12,
	"NIU_PID_QIANGZHUANG_OT":  13,
	"NIU_PID_QIANGZHUANG_REQ": 14,
	"NIU_PID_QIANGZHUANG_ACK": 15,
	"NIU_PID_QIANGZHUANG_BC":  16,
	"NIU_PID_JIABEI_OT":       17,
	"NIU_PID_JIABEI_REQ":      18,
	"NIU_PID_JIABEI_ACK":      19,
	"NIU_PID_JIABEI_BC":       20,
	"NIU_PID_BIPAI_RESULT_BC": 21,
	"NIU_PID_GAME_END_BC":     22,
}

func (x NiuniuEnumProtoid) Enum() *NiuniuEnumProtoid {
	p := new(NiuniuEnumProtoid)
	*p = x
	return p
}
func (x NiuniuEnumProtoid) String() string {
	return proto.EnumName(NiuniuEnumProtoid_name, int32(x))
}
func (x *NiuniuEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumProtoid_value, data, "NiuniuEnumProtoid")
	if err != nil {
		return err
	}
	*x = NiuniuEnumProtoid(value)
	return nil
}
func (NiuniuEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

// =================================公共================================
// 牛牛牌的类型
type NiuniuEnum_PokerType int32

const (
	NiuniuEnum_PokerType_NO_NIU           NiuniuEnum_PokerType = 1
	NiuniuEnum_PokerType_NIU_1            NiuniuEnum_PokerType = 2
	NiuniuEnum_PokerType_NIU_2            NiuniuEnum_PokerType = 3
	NiuniuEnum_PokerType_NIU_3            NiuniuEnum_PokerType = 4
	NiuniuEnum_PokerType_NIU_4            NiuniuEnum_PokerType = 5
	NiuniuEnum_PokerType_NIU_5            NiuniuEnum_PokerType = 6
	NiuniuEnum_PokerType_NIU_6            NiuniuEnum_PokerType = 7
	NiuniuEnum_PokerType_NIU_7            NiuniuEnum_PokerType = 8
	NiuniuEnum_PokerType_NIU_8            NiuniuEnum_PokerType = 9
	NiuniuEnum_PokerType_NIU_9            NiuniuEnum_PokerType = 10
	NiuniuEnum_PokerType_NIU_NIU          NiuniuEnum_PokerType = 11
	NiuniuEnum_PokerType_YIN_NIU          NiuniuEnum_PokerType = 12
	NiuniuEnum_PokerType_JIN_NIU          NiuniuEnum_PokerType = 13
	NiuniuEnum_PokerType_WU_XIAO_NIU      NiuniuEnum_PokerType = 14
	NiuniuEnum_PokerType_NIU_ZHA_DAN      NiuniuEnum_PokerType = 15
	NiuniuEnum_PokerType_NIU_YI_TIAO_LONG NiuniuEnum_PokerType = 16
)

var NiuniuEnum_PokerType_name = map[int32]string{
	1:  "NO_NIU",
	2:  "NIU_1",
	3:  "NIU_2",
	4:  "NIU_3",
	5:  "NIU_4",
	6:  "NIU_5",
	7:  "NIU_6",
	8:  "NIU_7",
	9:  "NIU_8",
	10: "NIU_9",
	11: "NIU_NIU",
	12: "YIN_NIU",
	13: "JIN_NIU",
	14: "WU_XIAO_NIU",
	15: "NIU_ZHA_DAN",
	16: "NIU_YI_TIAO_LONG",
}
var NiuniuEnum_PokerType_value = map[string]int32{
	"NO_NIU":           1,
	"NIU_1":            2,
	"NIU_2":            3,
	"NIU_3":            4,
	"NIU_4":            5,
	"NIU_5":            6,
	"NIU_6":            7,
	"NIU_7":            8,
	"NIU_8":            9,
	"NIU_9":            10,
	"NIU_NIU":          11,
	"YIN_NIU":          12,
	"JIN_NIU":          13,
	"WU_XIAO_NIU":      14,
	"NIU_ZHA_DAN":      15,
	"NIU_YI_TIAO_LONG": 16,
}

func (x NiuniuEnum_PokerType) Enum() *NiuniuEnum_PokerType {
	p := new(NiuniuEnum_PokerType)
	*p = x
	return p
}
func (x NiuniuEnum_PokerType) String() string {
	return proto.EnumName(NiuniuEnum_PokerType_name, int32(x))
}
func (x *NiuniuEnum_PokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnum_PokerType_value, data, "NiuniuEnum_PokerType")
	if err != nil {
		return err
	}
	*x = NiuniuEnum_PokerType(value)
	return nil
}
func (NiuniuEnum_PokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{1} }

// 房间状态
type NiuniuEnumDeskState int32

const (
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER       NiuniuEnumDeskState = 1
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_READY       NiuniuEnumDeskState = 2
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_START       NiuniuEnumDeskState = 3
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_QIANGZHUANG NiuniuEnumDeskState = 4
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_JIABEI      NiuniuEnumDeskState = 5
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_BIPAI       NiuniuEnumDeskState = 6
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_RESULT      NiuniuEnumDeskState = 7
)

var NiuniuEnumDeskState_name = map[int32]string{
	1: "NIU_DESK_STATUS_WAIT_ENTER",
	2: "NIU_DESK_STATUS_WAIT_READY",
	3: "NIU_DESK_STATUS_WAIT_START",
	4: "NIU_DESK_STATUS_WAIT_QIANGZHUANG",
	5: "NIU_DESK_STATUS_WAIT_JIABEI",
	6: "NIU_DESK_STATUS_WAIT_BIPAI",
	7: "NIU_DESK_STATUS_WAIT_RESULT",
}
var NiuniuEnumDeskState_value = map[string]int32{
	"NIU_DESK_STATUS_WAIT_ENTER":       1,
	"NIU_DESK_STATUS_WAIT_READY":       2,
	"NIU_DESK_STATUS_WAIT_START":       3,
	"NIU_DESK_STATUS_WAIT_QIANGZHUANG": 4,
	"NIU_DESK_STATUS_WAIT_JIABEI":      5,
	"NIU_DESK_STATUS_WAIT_BIPAI":       6,
	"NIU_DESK_STATUS_WAIT_RESULT":      7,
}

func (x NiuniuEnumDeskState) Enum() *NiuniuEnumDeskState {
	p := new(NiuniuEnumDeskState)
	*p = x
	return p
}
func (x NiuniuEnumDeskState) String() string {
	return proto.EnumName(NiuniuEnumDeskState_name, int32(x))
}
func (x *NiuniuEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumDeskState_value, data, "NiuniuEnumDeskState")
	if err != nil {
		return err
	}
	*x = NiuniuEnumDeskState(value)
	return nil
}
func (NiuniuEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{2} }

// 坐庄规则
type NiuniuEnumBankerRule int32

const (
	NiuniuEnumBankerRule_DING_ZHUANG       NiuniuEnumBankerRule = 1
	NiuniuEnumBankerRule_SUI_JI_ZUO_ZHUANG NiuniuEnumBankerRule = 2
	NiuniuEnumBankerRule_QIANG_ZHUANG      NiuniuEnumBankerRule = 3
)

var NiuniuEnumBankerRule_name = map[int32]string{
	1: "DING_ZHUANG",
	2: "SUI_JI_ZUO_ZHUANG",
	3: "QIANG_ZHUANG",
}
var NiuniuEnumBankerRule_value = map[string]int32{
	"DING_ZHUANG":       1,
	"SUI_JI_ZUO_ZHUANG": 2,
	"QIANG_ZHUANG":      3,
}

func (x NiuniuEnumBankerRule) Enum() *NiuniuEnumBankerRule {
	p := new(NiuniuEnumBankerRule)
	*p = x
	return p
}
func (x NiuniuEnumBankerRule) String() string {
	return proto.EnumName(NiuniuEnumBankerRule_name, int32(x))
}
func (x *NiuniuEnumBankerRule) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumBankerRule_value, data, "NiuniuEnumBankerRule")
	if err != nil {
		return err
	}
	*x = NiuniuEnumBankerRule(value)
	return nil
}
func (NiuniuEnumBankerRule) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{3} }

// 打出去的牌
type NiuniuClientPoker struct {
	Pais             []*ClientBasePoker    `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuClientPoker) Reset()                    { *m = NiuniuClientPoker{} }
func (m *NiuniuClientPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientPoker) ProtoMessage()               {}
func (*NiuniuClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *NiuniuClientPoker) GetPais() []*ClientBasePoker {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuClientPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// 对局账单信息
type NiuniuUserBill struct {
	Score            *int32 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	CountHasNiu      *int32 `protobuf:"varint,2,opt,name=count_has_niu" json:"count_has_niu,omitempty"`
	CountNoNiu       *int32 `protobuf:"varint,3,opt,name=count_no_niu" json:"count_no_niu,omitempty"`
	CountWin         *int32 `protobuf:"varint,4,opt,name=count_win" json:"count_win,omitempty"`
	CountLost        *int32 `protobuf:"varint,5,opt,name=count_lost" json:"count_lost,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NiuniuUserBill) Reset()                    { *m = NiuniuUserBill{} }
func (m *NiuniuUserBill) String() string            { return proto.CompactTextString(m) }
func (*NiuniuUserBill) ProtoMessage()               {}
func (*NiuniuUserBill) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{1} }

func (m *NiuniuUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuniuUserBill) GetCountHasNiu() int32 {
	if m != nil && m.CountHasNiu != nil {
		return *m.CountHasNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountNoNiu() int32 {
	if m != nil && m.CountNoNiu != nil {
		return *m.CountNoNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *NiuniuUserBill) GetCountLost() int32 {
	if m != nil && m.CountLost != nil {
		return *m.CountLost
	}
	return 0
}

// desk 配置选项
type NiuniuDeskOption struct {
	MinUser          *int32                `protobuf:"varint,1,opt,name=minUser" json:"minUser,omitempty"`
	MaxUser          *int32                `protobuf:"varint,2,opt,name=maxUser" json:"maxUser,omitempty"`
	MaxCircle        *int32                `protobuf:"varint,3,opt,name=maxCircle" json:"maxCircle,omitempty"`
	HasFlower        *bool                 `protobuf:"varint,4,opt,name=hasFlower" json:"hasFlower,omitempty"`
	BankRule         *NiuniuEnumBankerRule `protobuf:"varint,5,opt,name=bankRule,enum=ddproto.NiuniuEnumBankerRule" json:"bankRule,omitempty"`
	IsFlowerPlay     *bool                 `protobuf:"varint,6,opt,name=isFlowerPlay" json:"isFlowerPlay,omitempty"`
	IsJiaoFenJiaBei  *bool                 `protobuf:"varint,7,opt,name=isJiaoFenJiaBei" json:"isJiaoFenJiaBei,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuDeskOption) Reset()                    { *m = NiuniuDeskOption{} }
func (m *NiuniuDeskOption) String() string            { return proto.CompactTextString(m) }
func (*NiuniuDeskOption) ProtoMessage()               {}
func (*NiuniuDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{2} }

func (m *NiuniuDeskOption) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxUser() int32 {
	if m != nil && m.MaxUser != nil {
		return *m.MaxUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxCircle() int32 {
	if m != nil && m.MaxCircle != nil {
		return *m.MaxCircle
	}
	return 0
}

func (m *NiuniuDeskOption) GetHasFlower() bool {
	if m != nil && m.HasFlower != nil {
		return *m.HasFlower
	}
	return false
}

func (m *NiuniuDeskOption) GetBankRule() NiuniuEnumBankerRule {
	if m != nil && m.BankRule != nil {
		return *m.BankRule
	}
	return NiuniuEnumBankerRule_DING_ZHUANG
}

func (m *NiuniuDeskOption) GetIsFlowerPlay() bool {
	if m != nil && m.IsFlowerPlay != nil {
		return *m.IsFlowerPlay
	}
	return false
}

func (m *NiuniuDeskOption) GetIsJiaoFenJiaBei() bool {
	if m != nil && m.IsJiaoFenJiaBei != nil {
		return *m.IsJiaoFenJiaBei
	}
	return false
}

func init() {
	proto.RegisterType((*NiuniuClientPoker)(nil), "ddproto.niuniu_client_poker")
	proto.RegisterType((*NiuniuUserBill)(nil), "ddproto.niuniu_user_bill")
	proto.RegisterType((*NiuniuDeskOption)(nil), "ddproto.niuniu_desk_option")
	proto.RegisterEnum("ddproto.NiuniuEnumProtoid", NiuniuEnumProtoid_name, NiuniuEnumProtoid_value)
	proto.RegisterEnum("ddproto.NiuniuEnum_PokerType", NiuniuEnum_PokerType_name, NiuniuEnum_PokerType_value)
	proto.RegisterEnum("ddproto.NiuniuEnumDeskState", NiuniuEnumDeskState_name, NiuniuEnumDeskState_value)
	proto.RegisterEnum("ddproto.NiuniuEnumBankerRule", NiuniuEnumBankerRule_name, NiuniuEnumBankerRule_value)
}

var fileDescriptor28 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x4f, 0xe2, 0x5c,
	0x10, 0x96, 0x2f, 0xd1, 0x01, 0xb4, 0x1c, 0x04, 0x79, 0x31, 0xf1, 0x25, 0x66, 0x2f, 0x0c, 0xd9,
	0x98, 0x2c, 0xfb, 0x7d, 0x59, 0xa0, 0x62, 0xd5, 0x2d, 0x08, 0x6d, 0x5c, 0xbd, 0x39, 0xa9, 0xd0,
	0x64, 0x1b, 0xa1, 0x25, 0x14, 0xe2, 0x7a, 0xb7, 0xff, 0x6c, 0x7f, 0xcd, 0x5e, 0xec, 0xbf, 0xd8,
	0x99, 0xd3, 0x83, 0xd6, 0x48, 0x13, 0x23, 0xd3, 0x67, 0x66, 0x9e, 0xf9, 0x7a, 0x5a, 0x28, 0x7a,
	0xee, 0x12, 0xff, 0xf8, 0x9d, 0x1d, 0x38, 0x27, 0xb3, 0xb9, 0xbf, 0xf0, 0x59, 0x76, 0x3c, 0x16,
	0x46, 0xad, 0x34, 0xf2, 0xa7, 0x53, 0xdf, 0xe3, 0xa3, 0x89, 0xeb, 0x78, 0x8b, 0xd0, 0x7b, 0x34,
	0x85, 0x92, 0x4c, 0x09, 0x61, 0x3e, 0xf3, 0xef, 0x9d, 0x39, 0x3b, 0x86, 0xf4, 0xcc, 0x76, 0x83,
	0x6a, 0xb2, 0x9e, 0x3a, 0xce, 0x35, 0x6b, 0x27, 0x92, 0xe3, 0x44, 0x06, 0x11, 0xbd, 0x8c, 0x7c,
	0x0b, 0xe9, 0xc5, 0xe3, 0xcc, 0xa9, 0xa6, 0xea, 0x89, 0xe3, 0x9d, 0xe6, 0xe1, 0x53, 0xa4, 0x64,
	0x75, 0xbc, 0xe5, 0x94, 0xf7, 0x29, 0xd2, 0xc4, 0xa8, 0xa3, 0x25, 0x28, 0xd2, 0xb1, 0x0c, 0x9c,
	0x39, 0xbf, 0x73, 0x27, 0x13, 0x56, 0x80, 0x4c, 0x30, 0xf2, 0xe7, 0x4e, 0x35, 0x81, 0x14, 0x19,
	0x56, 0x86, 0xc2, 0xc8, 0x5f, 0x62, 0x91, 0x1f, 0x76, 0xc0, 0x31, 0x12, 0x7b, 0x20, 0x78, 0x0f,
	0xf2, 0x21, 0xec, 0xf9, 0x02, 0x4d, 0x09, 0xb4, 0x08, 0xdb, 0x21, 0xfa, 0xe0, 0x7a, 0xd5, 0xb4,
	0x80, 0x18, 0x40, 0x08, 0x4d, 0xfc, 0x60, 0x51, 0xcd, 0x10, 0x76, 0xf4, 0x3b, 0x01, 0x4c, 0xd6,
	0x1d, 0x3b, 0xc1, 0x3d, 0xf7, 0x67, 0x0b, 0xd7, 0xf7, 0xd8, 0x2e, 0x64, 0xa7, 0xae, 0x67, 0x61,
	0x27, 0xb2, 0x36, 0x01, 0xf6, 0x4f, 0x01, 0x24, 0x57, 0xfc, 0x08, 0xb4, 0xdd, 0xf9, 0x68, 0xe2,
	0x3c, 0x97, 0xc4, 0xce, 0x4e, 0x27, 0xfe, 0x03, 0x46, 0x51, 0xc9, 0x2d, 0xd6, 0x84, 0xad, 0x3b,
	0xdb, 0xbb, 0x1f, 0x2c, 0x31, 0x28, 0x23, 0xf6, 0x50, 0x5f, 0xbb, 0x07, 0x0a, 0xc2, 0xa9, 0xe7,
	0x18, 0x47, 0xf3, 0xb8, 0x92, 0xa5, 0x3f, 0xb1, 0x1f, 0xab, 0x9b, 0x82, 0x69, 0x1f, 0x76, 0xdd,
	0xe0, 0xdc, 0xb5, 0xfd, 0x53, 0xc7, 0xc3, 0x9f, 0x96, 0xe3, 0x56, 0xb3, 0xe4, 0x68, 0xfc, 0x4d,
	0x3f, 0x1d, 0x4a, 0x50, 0x09, 0x72, 0x77, 0x8c, 0xdb, 0x2a, 0x1a, 0xba, 0xc5, 0xfb, 0x7a, 0x87,
	0x9f, 0x69, 0xea, 0xc0, 0x6c, 0x69, 0xaa, 0xa9, 0x6c, 0xb0, 0x0a, 0xb0, 0x15, 0x7c, 0x65, 0xe9,
	0xed, 0x0b, 0xde, 0xee, 0x19, 0x86, 0x92, 0x60, 0x35, 0xa8, 0xbc, 0xc6, 0xb9, 0xda, 0xbe, 0x50,
	0x92, 0xd1, 0x9c, 0xae, 0xfa, 0x4d, 0xe3, 0x97, 0xbd, 0xae, 0x6e, 0x28, 0xa9, 0x68, 0xce, 0x33,
	0x2e, 0x72, 0xd2, 0xec, 0x00, 0xf6, 0x57, 0xbe, 0xf6, 0x00, 0x4b, 0x6b, 0xbc, 0xa3, 0x0d, 0x2f,
	0xf8, 0x40, 0xbb, 0x52, 0x32, 0xd1, 0x44, 0xcd, 0x30, 0xb5, 0xc1, 0xb3, 0x6f, 0x33, 0xc6, 0x47,
	0xa4, 0x59, 0xf6, 0x1f, 0x94, 0xd7, 0xf8, 0x5a, 0x6d, 0x65, 0x2b, 0x3a, 0x2e, 0x96, 0xeb, 0xdc,
	0x08, 0xb6, 0xed, 0xd7, 0x30, 0x11, 0x01, 0xee, 0x58, 0x79, 0x09, 0x23, 0x47, 0x2e, 0x4a, 0x3f,
	0x34, 0x71, 0x65, 0xe1, 0x54, 0x3d, 0x53, 0xc9, 0xbf, 0x58, 0x8f, 0xae, 0x1a, 0xdd, 0xdb, 0x33,
	0x0b, 0xff, 0x93, 0xaf, 0x10, 0x1d, 0x35, 0xea, 0xa3, 0x06, 0x76, 0xe2, 0x9c, 0xd4, 0xc6, 0x6e,
	0x1c, 0x2b, 0x36, 0xa3, 0x44, 0x3b, 0x3f, 0xd7, 0xd5, 0x96, 0xa6, 0x53, 0xb1, 0x62, 0xf4, 0x16,
	0x12, 0xa6, 0x3a, 0x6c, 0x0d, 0x4e, 0x25, 0x4a, 0x6b, 0x68, 0x90, 0x7d, 0x2f, 0xda, 0x56, 0x4b,
	0xef, 0xab, 0xc4, 0x32, 0xb4, 0x2e, 0x4d, 0x72, 0x96, 0x51, 0x6b, 0xa5, 0x17, 0x77, 0xd5, 0x8c,
	0x0e, 0x39, 0x2a, 0x8d, 0x3f, 0x09, 0x28, 0xaf, 0x7d, 0x7d, 0xf1, 0xd5, 0xda, 0x34, 0x7a, 0x1c,
	0xb3, 0x50, 0x4a, 0xdb, 0x90, 0xa1, 0xf4, 0x77, 0xa8, 0x1c, 0x69, 0x36, 0x51, 0x2c, 0xd2, 0x7c,
	0x8f, 0xda, 0x90, 0xe6, 0x07, 0x54, 0x82, 0x34, 0x3f, 0xe2, 0xe1, 0xa5, 0xf9, 0x09, 0xef, 0x2c,
	0xcd, 0xcf, 0x78, 0x57, 0x69, 0x7e, 0xc1, 0x5b, 0x4a, 0xf3, 0x2b, 0xde, 0x2f, 0x07, 0x59, 0x32,
	0xa9, 0x5e, 0x8e, 0x1e, 0x6e, 0x50, 0x77, 0xf4, 0x90, 0xa7, 0x87, 0x73, 0xf9, 0x50, 0xc0, 0xb7,
	0x36, 0x77, 0x6d, 0xf1, 0xef, 0xba, 0x1a, 0xb6, 0xb6, 0x43, 0x00, 0xe5, 0xdd, 0x9e, 0xa9, 0xbc,
	0xa3, 0x1a, 0x78, 0x01, 0x29, 0x84, 0x1b, 0x9d, 0x9b, 0x14, 0x75, 0xd9, 0x33, 0xba, 0x8a, 0xd2,
	0xf8, 0x95, 0x84, 0x4a, 0x74, 0x4e, 0xf1, 0x69, 0x08, 0x16, 0xf6, 0xc2, 0x61, 0x87, 0x50, 0xa3,
	0x04, 0x21, 0x3c, 0x14, 0x89, 0x69, 0x0d, 0xf9, 0xb5, 0xaa, 0x9b, 0xa1, 0x1e, 0x71, 0xf8, 0x38,
	0xbf, 0x90, 0x19, 0x6e, 0x24, 0xce, 0x2f, 0x04, 0x87, 0x6b, 0x7a, 0x03, 0xf5, 0xb5, 0xfe, 0x88,
	0x3e, 0x70, 0x83, 0xff, 0xc3, 0xc1, 0xda, 0xa8, 0xf0, 0xc4, 0xb8, 0xd7, 0xb8, 0x32, 0xe2, 0xd8,
	0xb8, 0xec, 0x38, 0x82, 0x50, 0x06, 0x4a, 0xb6, 0x31, 0x84, 0xfd, 0xb8, 0x0f, 0x14, 0x2e, 0xb1,
	0xa3, 0xa3, 0x4c, 0x65, 0x37, 0x09, 0xd2, 0xd8, 0xd0, 0xd2, 0xb1, 0x38, 0xbf, 0xb5, 0x7a, 0x2b,
	0x38, 0xc9, 0x14, 0xc8, 0x8b, 0xae, 0x57, 0x48, 0xaa, 0xbf, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x77, 0x02, 0x08, 0x87, 0x06, 0x00, 0x00,
}
