package cat

import "C"

var (
	AuthCode     int32
	Name         string
	Desc         string
	Author       string
	Version      string
	ApiVersion   string
	MenuuTitle   string
	CoverBase64  string
	DeveloperKey string
)

//export NewAppInfo
func NewAppInfo(session C.int) int32 {
	var app = AppInfo{
		Name:         Name,
		Desc:         Desc,
		Author:       Author,
		Version:      Version,
		ApiVersion:   ApiVersion,
		MenuTitle:    MenuuTitle,
		CoverBase64:  CoverBase64,
		DeveloperKey: DeveloperKey,
	}
	json := app.ToJson()
	authCode := Initialize(int64(session), json)
	AuthCode = authCode
	return authCode
}

//export _EventInit
func _EventInit() int32 {
	return EventInit()
}

//export _EventEnable
func _EventEnable() int32 {
	return EventEnable()
}

//export _EventStop
func _EventStop() int32 {
	return EventStop()
}

//export _EventLogin
func _EventLogin(rob_wxid *C.char, rob_wxname *C.char, _type C.int) int32 {
	return EventLogin(goString(rob_wxid), goString(rob_wxname), int32(_type))
}

/*
   robot_wxid  机器人帐号ID
   type        消息类型，1文本，3图片，34语音，42名片，43视频，47动态                      表情，48地理位置，49分享链接，2001红包，2002小程序，2003群邀请
   from_wxid   来源群ID
   from_name   来源群昵称
   final_from_wxid   具体发群消息的成员ID
   final_from_name   具体发群消息的成员昵称
   to_wxid     接收信息的人的ID，（一般是机器人收到消息，所以是机器人                          的id,如果是机器人主动发消息给别人，那就是别人的id）
   Msg         群消息内容
*/

//export _EventGroupMsg
func _EventGroupMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, final_from_wxid *C.char, final_from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
	return EventGroupMsg(goString(robot_wxid), int32(_type), goString(from_wxid), goString(from_name), goString(final_from_wxid), goString(final_from_name), goString(to_wxid), goString(Msg))
}

//export _EventFriendMsg
func _EventFriendMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
	return EventFriendMsg(goString(robot_wxid), int32(_type), goString(from_wxid), goString(from_name), goString(to_wxid), goString(Msg))
}

//export _EventReceivedTransfer
func _EventReceivedTransfer(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, to_wxid *C.char, money *C.char, Msg_json *C.char) int32 {
	return EventReceivedTransfer(goString(robot_wxid), goString(from_wxid), goString(from_name), goString(to_wxid), goString(money), goString(Msg_json))

}

//export _EventScanCashMoney
func _EventScanCashMoney(robot_wxid *C.char, pay_wxid *C.char, pay_name *C.char, money *C.char, Msg_json *C.char) int32 {
	return EventScanCashMoney(goString(robot_wxid), goString(pay_wxid), goString(pay_name), goString(money), goString(Msg_json))
}

//export _EventFriendVerify
func _EventFriendVerify(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, to_wxid *C.char, Msg_json *C.char) int32 {
	return EventFriendVerify(goString(robot_wxid), goString(from_wxid), goString(from_name), goString(to_wxid), goString(Msg_json))
}

//export _EventGroupMemberAdd
func _EventGroupMemberAdd(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, Msg_json *C.char) int32 {
	return EventGroupMemberAdd(goString(robot_wxid), goString(from_wxid), goString(from_name), goString(Msg_json))
}

//export _EventGroupMemberDecrease
func _EventGroupMemberDecrease(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, Msg_json *C.char) int32 {
	return EventGroupMemberDecrease(goString(robot_wxid), goString(from_wxid), goString(from_name), goString(Msg_json))
}

//export _EventSysMsg
func _EventSysMsg(robot_wxid *C.char, _type C.int, Msg_json *C.char) int32 {
	return EventSysMsg(goString(robot_wxid), int32(_type), goString(Msg_json))
}

//export _Menu
func _Menu() int32 {
	return Menu()
}
