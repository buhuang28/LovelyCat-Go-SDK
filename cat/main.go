package cat

var EventInit func() int32
var EventEnable func() int32
var EventStop func() int32
var EventLogin func(rob_wxid string, rob_wxname string, _type int32) int32
var EventGroupMsg func(robot_wxid string, _type int32, from_wxid string, from_name string, final_from_wxid string, final_from_name string, to_wxid string, Msg string) int32
var EventFriendMsg func(robot_wxid string, _type int32, from_wxid string, from_name string, to_wxid string, Msg string) int32
var EventReceivedTransfer func(robot_wxid string, from_wxid string, from_name string, to_wxid string, money string, Msg_json string) int32
var EventScanCashMoney func(robot_wxid string, pay_wxid string, pay_name string, money string, Msg_json string) int32
var EventFriendVerify func(robot_wxid string, from_wxid string, from_name string, to_wxid string, Msg_json string) int32
var EventGroupMemberAdd func(robot_wxid string, from_wxid string, from_name string, Msg_json string) int32
var EventGroupMemberDecrease func(robot_wxid string, from_wxid string, from_name string, Msg_json string) int32
var EventSysMsg func(robot_wxid string, _type int32, Msg_json string) int32
var Menu func() int32
