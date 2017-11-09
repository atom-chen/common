// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_hall.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of hall_req_event from hall.proto

// Ignoring public import of hall_ack_event from hall.proto

// Ignoring public import of hall_lottery_item from hall.proto

// Ignoring public import of hall_sign_lottery_info from hall.proto

// Ignoring public import of hall_draw_lottery_info from hall.proto

// Ignoring public import of hall_req_mail_list from hall.proto

// Ignoring public import of hall_ack_mail_list from hall.proto

// Ignoring public import of hall_req_task from hall.proto

// Ignoring public import of hall_ack_task from hall.proto

// Ignoring public import of hall_req_checkTask from hall.proto

// Ignoring public import of hall_ack_checkTask from hall.proto

// Ignoring public import of hall_req_task_sum from hall.proto

// Ignoring public import of hall_ack_task_sum from hall.proto

// Ignoring public import of hall_req_checkBonus from hall.proto

// Ignoring public import of hall_ack_checkBonus from hall.proto

// Ignoring public import of hall_req_bag_items from hall.proto

// Ignoring public import of hall_ack_bag_items from hall.proto

// Ignoring public import of hall_req_userData from hall.proto

// Ignoring public import of hall_ack_userData from hall.proto

// Ignoring public import of hall_req_updateRealData from hall.proto

// Ignoring public import of hall_ack_updateRealData from hall.proto

// Ignoring public import of hall_req_goods_list from hall.proto

// Ignoring public import of hall_ack_goods_list from hall.proto

// Ignoring public import of hall_req_goods_buy from hall.proto

// Ignoring public import of hall_ack_goods_buy from hall.proto

// Ignoring public import of hall_goods_item_msg from hall.proto

// Ignoring public import of hall_req_rank from hall.proto

// Ignoring public import of hall_ack_rank from hall.proto

// Ignoring public import of hall_req_draw_lottery from hall.proto

// Ignoring public import of hall_ack_draw_lottery from hall.proto

// Ignoring public import of hall_req_ds_lottery_info from hall.proto

// Ignoring public import of hall_ack_ds_lottery_info from hall.proto

// Ignoring public import of hall_req_friends_list from hall.proto

// Ignoring public import of hall_ack_friends_list from hall.proto

// Ignoring public import of hall_req_recommend_user_list from hall.proto

// Ignoring public import of hall_ack_recommend_user_list from hall.proto

// Ignoring public import of hall_req_friends_search from hall.proto

// Ignoring public import of hall_ack_friends_search from hall.proto

// Ignoring public import of hall_req_add_friend from hall.proto

// Ignoring public import of hall_ack_add_friend from hall.proto

// Ignoring public import of hall_req_del_friend from hall.proto

// Ignoring public import of hall_ack_del_friend from hall.proto

// Ignoring public import of hall_friend_state from hall.proto

// Ignoring public import of hall_user_info from hall.proto

// Ignoring public import of hall_ack_strongbox_info from hall.proto

// Ignoring public import of hall_req_strongbox_info from hall.proto

// Ignoring public import of hall_req_strongbox_access from hall.proto

// Ignoring public import of hall_ack_strongbox_access from hall.proto

// Ignoring public import of game_GameRecord from hall.proto

// Ignoring public import of BeanUserRecord from hall.proto

// Ignoring public import of BeanGameRecord from hall.proto

// Ignoring public import of game_AckGameRecord from hall.proto

// Ignoring public import of hall_bean_bisai from hall.proto

// Ignoring public import of hall_req_bisai from hall.proto

// Ignoring public import of hall_req_friend_lottery_list from hall.proto

// Ignoring public import of hall_ack_friend_lottery_list from hall.proto

// Ignoring public import of hall_req_friend_lottery_draw from hall.proto

// Ignoring public import of hall_ack_friend_lottery_draw from hall.proto

// Ignoring public import of hall_req_distance_matched from hall.proto

// Ignoring public import of hall_ack_distance_matched from hall.proto

// Ignoring public import of hall_req_agent_room_gaming_list from hall.proto

// Ignoring public import of hall_ack_agent_room_gaming_list from hall.proto

// Ignoring public import of hall_req_agent_room_history_list from hall.proto

// Ignoring public import of hall_ack_agent_room_history_list from hall.proto

// Ignoring public import of hall_req_bind_invcode from hall.proto

// Ignoring public import of hall_ack_bind_invcode from hall.proto

// Ignoring public import of hall_enum_goods_type from hall.proto

// Ignoring public import of hall_strongbox_access_type from hall.proto

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

// 用于更新hall的配置文件
// 更新配置文件
type RpcHallUpdateConfig struct {
	ConfigId         *int32 `protobuf:"varint,1,opt,name=configId" json:"configId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpcHallUpdateConfig) Reset()                    { *m = RpcHallUpdateConfig{} }
func (m *RpcHallUpdateConfig) String() string            { return proto.CompactTextString(m) }
func (*RpcHallUpdateConfig) ProtoMessage()               {}
func (*RpcHallUpdateConfig) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{0} }

func (m *RpcHallUpdateConfig) GetConfigId() int32 {
	if m != nil && m.ConfigId != nil {
		return *m.ConfigId
	}
	return 0
}

// 代开房间状态更新消息
type HallRpcDeskEventMsg struct {
	Msg              *string            `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Desk             *CommonDeskByAgent `protobuf:"bytes,2,opt,name=desk" json:"desk,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallRpcDeskEventMsg) Reset()                    { *m = HallRpcDeskEventMsg{} }
func (m *HallRpcDeskEventMsg) String() string            { return proto.CompactTextString(m) }
func (*HallRpcDeskEventMsg) ProtoMessage()               {}
func (*HallRpcDeskEventMsg) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{1} }

func (m *HallRpcDeskEventMsg) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *HallRpcDeskEventMsg) GetDesk() *CommonDeskByAgent {
	if m != nil {
		return m.Desk
	}
	return nil
}

// 俱乐部状态更新
type HallRpcGroupRefresh struct {
	GroupId          *int32 `protobuf:"varint,1,opt,name=groupId" json:"groupId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallRpcGroupRefresh) Reset()                    { *m = HallRpcGroupRefresh{} }
func (m *HallRpcGroupRefresh) String() string            { return proto.CompactTextString(m) }
func (*HallRpcGroupRefresh) ProtoMessage()               {}
func (*HallRpcGroupRefresh) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{2} }

func (m *HallRpcGroupRefresh) GetGroupId() int32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func init() {
	proto.RegisterType((*RpcHallUpdateConfig)(nil), "ddproto.rpc_hall_update_config")
	proto.RegisterType((*HallRpcDeskEventMsg)(nil), "ddproto.hall_rpc_desk_event_msg")
	proto.RegisterType((*HallRpcGroupRefresh)(nil), "ddproto.hall_rpc_group_refresh")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HallRpc service

type HallRpcClient interface {
	// 创建房间
	GetAgentRoomList(ctx context.Context, in *HallReqAgentRoomGamingList, opts ...grpc.CallOption) (*HallAckAgentRoomGamingList, error)
	SendDeskEventMsg(ctx context.Context, in *HallRpcDeskEventMsg, opts ...grpc.CallOption) (*ProtoHeader, error)
	RefreshGroupInfo(ctx context.Context, in *HallRpcGroupRefresh, opts ...grpc.CallOption) (*ProtoHeader, error)
}

type hallRpcClient struct {
	cc *grpc.ClientConn
}

func NewHallRpcClient(cc *grpc.ClientConn) HallRpcClient {
	return &hallRpcClient{cc}
}

func (c *hallRpcClient) GetAgentRoomList(ctx context.Context, in *HallReqAgentRoomGamingList, opts ...grpc.CallOption) (*HallAckAgentRoomGamingList, error) {
	out := new(HallAckAgentRoomGamingList)
	err := grpc.Invoke(ctx, "/ddproto.HallRpc/GetAgentRoomList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallRpcClient) SendDeskEventMsg(ctx context.Context, in *HallRpcDeskEventMsg, opts ...grpc.CallOption) (*ProtoHeader, error) {
	out := new(ProtoHeader)
	err := grpc.Invoke(ctx, "/ddproto.HallRpc/SendDeskEventMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hallRpcClient) RefreshGroupInfo(ctx context.Context, in *HallRpcGroupRefresh, opts ...grpc.CallOption) (*ProtoHeader, error) {
	out := new(ProtoHeader)
	err := grpc.Invoke(ctx, "/ddproto.HallRpc/RefreshGroupInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HallRpc service

type HallRpcServer interface {
	// 创建房间
	GetAgentRoomList(context.Context, *HallReqAgentRoomGamingList) (*HallAckAgentRoomGamingList, error)
	SendDeskEventMsg(context.Context, *HallRpcDeskEventMsg) (*ProtoHeader, error)
	RefreshGroupInfo(context.Context, *HallRpcGroupRefresh) (*ProtoHeader, error)
}

func RegisterHallRpcServer(s *grpc.Server, srv HallRpcServer) {
	s.RegisterService(&_HallRpc_serviceDesc, srv)
}

func _HallRpc_GetAgentRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HallReqAgentRoomGamingList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallRpcServer).GetAgentRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.HallRpc/GetAgentRoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallRpcServer).GetAgentRoomList(ctx, req.(*HallReqAgentRoomGamingList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HallRpc_SendDeskEventMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HallRpcDeskEventMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallRpcServer).SendDeskEventMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.HallRpc/SendDeskEventMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallRpcServer).SendDeskEventMsg(ctx, req.(*HallRpcDeskEventMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _HallRpc_RefreshGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HallRpcGroupRefresh)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HallRpcServer).RefreshGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.HallRpc/RefreshGroupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HallRpcServer).RefreshGroupInfo(ctx, req.(*HallRpcGroupRefresh))
	}
	return interceptor(ctx, in, info, handler)
}

var _HallRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddproto.HallRpc",
	HandlerType: (*HallRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgentRoomList",
			Handler:    _HallRpc_GetAgentRoomList_Handler,
		},
		{
			MethodName: "SendDeskEventMsg",
			Handler:    _HallRpc_SendDeskEventMsg_Handler,
		},
		{
			MethodName: "RefreshGroupInfo",
			Handler:    _HallRpc_RefreshGroupInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_hall.proto",
}

func init() { proto.RegisterFile("rpc_hall.proto", fileDescriptor50) }

var fileDescriptor50 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x50, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0x5e, 0xa7, 0x32, 0x3d, 0x82, 0x94, 0x28, 0x73, 0x14, 0xc4, 0xd1, 0xab, 0x5d, 0x0d, 0x1c,
	0xbe, 0x80, 0xa0, 0x6c, 0xe2, 0x84, 0x51, 0xef, 0x0d, 0xb1, 0x39, 0xcb, 0x4a, 0xf3, 0x53, 0x93,
	0x4c, 0xf0, 0x25, 0x7c, 0x66, 0x69, 0xba, 0x55, 0xdc, 0xd8, 0x4d, 0x38, 0x39, 0x7c, 0x7f, 0xe7,
	0x83, 0x0b, 0x5b, 0xe5, 0x74, 0xc5, 0xa4, 0x1c, 0x57, 0xd6, 0x78, 0x43, 0x7a, 0x9c, 0x87, 0x21,
	0x81, 0xbf, 0x65, 0x72, 0x99, 0x1b, 0xa5, 0x8c, 0xa6, 0xb9, 0x2c, 0x50, 0xfb, 0x66, 0x99, 0xde,
	0x43, 0x7f, 0xcb, 0xa5, 0xeb, 0x8a, 0x33, 0x8f, 0x34, 0x37, 0x7a, 0x59, 0x08, 0x92, 0xc0, 0x69,
	0x33, 0x3d, 0xf3, 0x41, 0x34, 0x8c, 0x46, 0x27, 0x59, 0xfb, 0x4f, 0xdf, 0xe1, 0x3a, 0x30, 0x6a,
	0x2a, 0x47, 0x57, 0x52, 0xfc, 0x42, 0xed, 0xa9, 0x72, 0x82, 0xc4, 0x70, 0xa4, 0x9c, 0x08, 0x8c,
	0xb3, 0xac, 0x1e, 0xc9, 0x1d, 0x1c, 0xd7, 0x98, 0x41, 0x77, 0x18, 0x8d, 0xce, 0x27, 0x37, 0xe3,
	0x4d, 0xb6, 0xf1, 0x26, 0x4e, 0xe0, 0x7f, 0x7c, 0x53, 0x26, 0x50, 0xfb, 0x2c, 0x40, 0xd3, 0x09,
	0xf4, 0x5b, 0x7d, 0x61, 0xcd, 0xba, 0xa2, 0x16, 0x97, 0x16, 0xdd, 0x8a, 0x0c, 0xa0, 0x17, 0x16,
	0x6d, 0xa8, 0xed, 0x77, 0xf2, 0xd3, 0x85, 0xde, 0x8c, 0x49, 0x99, 0x55, 0x39, 0x59, 0x41, 0x3c,
	0x45, 0xff, 0x10, 0x14, 0x8d, 0x51, 0xf3, 0xc2, 0x79, 0x32, 0x6a, 0x8d, 0x1b, 0x69, 0xfc, 0x6c,
	0x2c, 0xa9, 0x35, 0x46, 0x51, 0xc1, 0x54, 0xa1, 0x05, 0x95, 0x85, 0xf3, 0xc9, 0x0e, 0x92, 0xe5,
	0xe5, 0x01, 0x64, 0xda, 0x21, 0x73, 0x88, 0xdf, 0x50, 0xf3, 0x47, 0x74, 0xe5, 0x53, 0xdd, 0xc1,
	0xab, 0x13, 0x64, 0xb8, 0xe3, 0xb4, 0x57, 0x52, 0x72, 0xd5, 0x22, 0x16, 0xf5, 0x3b, 0x43, 0xc6,
	0xd1, 0xa6, 0x1d, 0xf2, 0x02, 0x71, 0xd6, 0x1c, 0x3a, 0x0d, 0x57, 0xe9, 0xa5, 0x21, 0xb7, 0xfb,
	0x6a, 0xff, 0x2a, 0x39, 0x24, 0xb6, 0xe8, 0x2c, 0xa2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0x6e, 0xc4, 0x27, 0x19, 0x02, 0x00, 0x00,
}
