package main

import "C"

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

//export _EventGroupMsg
func _EventGroupMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, final_from_wxid *C.char, final_from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
	return 0
}

//export _EventFriendMsg
func _EventFriendMsg(robot_wxid *C.char, _type C.int, from_wxid *C.char, from_name *C.char, to_wxid *C.char, Msg *C.char) int32 {
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
