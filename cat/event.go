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
	return 0
}

//export _EventEnable
func _EventEnable() int32 {
	return 0
}

//export _EventStop
func _EventStop() int32 {
	return 0
}

//export _EventLogin
func _EventLogin(rob_wxid *C.char, rob_wxname *C.char, _type C.int) int32 {
	return 0
}

/*
   robot_wxid  机器人帐号ID
   type        消息类型，1文本，3图片，34语音，42名片，43视频，47动态表情，48地理位置，49分享链接，2001红包，2002小程序，2003群邀请
   from_wxid   来源群ID
   from_name   来源群昵称
   final_from_wxid   具体发群消息的成员ID
   final_from_name   具体发群消息的成员昵称
   to_wxid     接收信息的人的ID，（一般是机器人收到消息，所以是机器人的id,如果是机器人主动发消息给别人，那就是别人的id）
   Msg         群消息内容
*/
//export _EventGroupMsg
func _EventGroupMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, final_from_wxid *C.char, final_from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
	botId := goString(robot_wxid)
	msgType := int32(_type)
	fromGroupId := goString(from_wxid)
	//fromGroupName := goString(from_name)
	//fromVId := goString(final_from_wxid)
	//fromVName := goString(final_from_name)
	//ToVId := goString(to_wxid)
	message := goString(Msg)
	switch msgType {
	case TEXT:
		SendTextMsg(botId, fromGroupId, "我是复读机:\r"+message)
	case IMAGE:
		SendImageMsg(botId, fromGroupId, "图片路径")
	}
	return 0
}

//export _EventFriendMsg
func _EventFriendMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
	SendTextMsg(goString(robot_wxid), goString(from_wxid), "我是复读机:\r"+goString(Msg))
	return 0
}

//export _EventReceivedTransfer
func _EventReceivedTransfer(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, to_wxid *C.char, money *C.char, Msg_json *C.char) int32 {
	return 0
}

//export _EventScanCashMoney
func _EventScanCashMoney(robot_wxid *C.char, pay_wxid *C.char, pay_name *C.char, money *C.char, Msg_json *C.char) int32 {
	return 0
}

//export _EventFriendVerify
func _EventFriendVerify(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, to_wxid *C.char, Msg_json *C.char) int32 {
	return 0
}

//export _EventGroupMemberAdd
func _EventGroupMemberAdd(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, Msg_json *C.char) int32 {
	return 0
}

//export _EventGroupMemberDecrease
func _EventGroupMemberDecrease(robot_wxid *C.char, from_wxid *C.char, from_name *C.char, Msg_json *C.char) int32 {
	return 0
}

//export _EventSysMsg
func _EventSysMsg(robot_wxid *C.char, _type C.int, Msg_json *C.char) int32 {
	return 0
}

//export _Menu
func _Menu() int32 {
	return 0
}
