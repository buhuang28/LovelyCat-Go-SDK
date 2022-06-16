package cat

import (
	"syscall"
)

var (
	api                     = syscall.NewLazyDLL("coupler.dll")
	ApiAcceptTransfer       = getApiProc("Api_AcceptTransfer")
	ApiAgreeFriendVerify    = getApiProc("Api_AgreeFriendVerify")
	ApiAgreeGroupInvite     = getApiProc("Api_AgreeGroupInvite")
	ApiAppendLogs           = getApiProc("Api_AppendLogs")
	ApiBuildingGroup        = getApiProc("Api_BuildingGroup")
	ApiDeleteFriend         = getApiProc("Api_DeleteFriend")
	ApiGetAppDirectory      = getApiProc("Api_GetAppDirectory")
	ApiGetFriendList        = getApiProc("Api_GetFriendList")
	ApiGetGroupList         = getApiProc("Api_GetGroupList")
	ApiGetGroupMemberList   = getApiProc("Api_GetGroupMemberList")
	ApiGetLoggedAccountList = getApiProc("Api_GetLoggedAccountList")
	ApiGetRobotHeadimgurl   = getApiProc("Api_GetRobotHeadimgurl")
	ApiGetRobotName         = getApiProc("Api_GetRobotName")
	ApiInitialize           = getApiProc("Api_Initialize")
	ApiInviteInGroup        = getApiProc("Api_InviteInGroup")
	ApiModifyFriendNote     = getApiProc("Api_ModifyFriendNote")
	ApiModifyGroupName      = getApiProc("Api_ModifyGroupName")
	ApiModifyGroupNotice    = getApiProc("Api_ModifyGroupNotice")
	ApiQuitGroup            = getApiProc("Api_QuitGroup")
	ApiReloadPlug           = getApiProc("Api_ReloadPlug")
	ApiSendEmojiMsg         = getApiProc("Api_SendEmojiMsg")
	ApiSendFileMsg          = getApiProc("Api_SendFileMsg")
	ApiSendGroupMsgAndAt    = getApiProc("Api_SendGroupMsgAndAt")
	ApiSendImageMsg         = getApiProc("Api_SendImageMsg")
	ApiSendLinkMsg          = getApiProc("Api_SendLinkMsg")
	ApiSendMusicMsg         = getApiProc("Api_SendMusicMsg")
	ApiSendTextMsg          = getApiProc("Api_SendTextMsg")
	ApiSendVideoMsg         = getApiProc("Api_SendVideoMsg")
	ApiSetFatal             = getApiProc("Api_SetFatal")
)

func getApiProc(name string) *syscall.LazyProc {
	return api.NewProc(name)
}

func Initialize(AuthCode int64, jsonInfo string) int32 {
	call, _, _ := ApiInitialize.Call(int2ptr(AuthCode), str2ptr(jsonInfo))
	return int32(call)
}

func AcceptTransfer(robot_wxid string, from_wxid string, json_msg string) int64 {
	call, _, _ := ApiAcceptTransfer.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(from_wxid), str2ptr(json_msg))
	return int64(call)
}
func AgreeFriendVerify(robot_wxid string, json_msg string) int64 {
	call, _, _ := ApiAgreeFriendVerify.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(json_msg))
	return int64(call)
}

func AgreeGroupInvite(robot_wxid string, json_msg string) int64 {
	call, _, _ := ApiAgreeGroupInvite.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(json_msg))
	return int64(call)
}

func AppendLogs(msg1 string) int64 {
	call, _, _ := ApiAppendLogs.Call(int2ptr(int64(AuthCode)), str2ptr(msg1))
	return int64(call)
}

func BuildingGroup(robot_wxid string, friends string) int64 {
	call, _, _ := ApiBuildingGroup.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(friends))
	return int64(call)
}

func GetRobotName(robot_wxid string) string {
	call, _, _ := ApiGetRobotName.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid))
	return ptr2str(call)
}

func DeleteFriend(robot_wxid string, friend_wxid string) int64 {
	call, _, _ := ApiDeleteFriend.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(friend_wxid))
	return int64(call)
}

func GetAppDirectory() string {
	call, _, _ := ApiGetAppDirectory.Call(int2ptr(int64(AuthCode)))
	return ptr2str(call)
}

func GetFriendList(robot_wxid string, is_refresh bool) string {
	call, _, _ := ApiGetFriendList.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), int2ptr(int64(cBool(is_refresh))))
	return ptr2str(call)
}

func GetGroupList(robot_wxid string, is_refresh bool) string {

	call, _, _ := ApiGetGroupList.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), int2ptr(int64(cBool(is_refresh))))
	return ptr2str(call)
}

func GetGroupMemberList(robot_wxid string, group_wxid string, is_refresh bool) string {
	call, _, _ := ApiGetGroupMemberList.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid), int2ptr(int64(cBool(is_refresh))))
	return ptr2str(call)
}

func GetLoggedAccountList() string {
	call, _, _ := ApiGetLoggedAccountList.Call(int2ptr(int64(AuthCode)))
	return ptr2str(call)
}

func GetRobotHeadimgurl(robot_wxid string) string {
	call, _, _ := ApiGetRobotHeadimgurl.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid))
	return ptr2str(call)
}

func InviteInGroup(robot_wxid string, group_wxid string, friend_wxid string) int64 {
	call, _, _ := ApiInviteInGroup.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid), str2ptr(friend_wxid))
	return int64(call)
}

func ModifyFriendNote(robot_wxid string, friend_wxid string, note string) int64 {
	call, _, _ := ApiModifyFriendNote.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(friend_wxid), str2ptr(note))
	return int64(call)
}

func ModifyGroupName(robot_wxid string, group_wxid string, group_name string) int64 {
	call, _, _ := ApiModifyGroupName.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid), str2ptr(group_name))
	return int64(call)
}

func ModifyGroupNotice(robot_wxid string, group_wxid string, content string) int64 {
	call, _, _ := ApiModifyGroupNotice.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid), str2ptr(content))
	return int64(call)
}

func QuitGroup(robot_wxid string, group_wxid string) int64 {
	call, _, _ := ApiQuitGroup.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid))
	return int64(call)
}

func ReloadPlug() int64 {
	call, _, _ := ApiReloadPlug.Call(int2ptr(int64(AuthCode)))
	return int64(call)
}

func SendEmojiMsg(robot_wxid string, to_wxid string, path string) int64 {
	call, _, _ := ApiSendEmojiMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(path))
	return int64(call)
}

func SendFileMsg(robot_wxid string, to_wxid string, path string) int64 {
	call, _, _ := ApiSendFileMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(path))
	return int64(call)
}

func SendGroupMsgAndAt(robot_wxid string, group_wxid string, member_wxid string, member_name string, msg string) int64 {
	call, _, _ := ApiSendGroupMsgAndAt.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(group_wxid), str2ptr(member_wxid), str2ptr(member_name), str2ptr(msg))
	return int64(call)
}

func SendImageMsg(robot_wxid string, to_wxid string, path string) int64 {
	call, _, _ := ApiSendImageMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(path))
	return int64(call)
}

func SendLinkMsg(robot_wxid string, to_wxid string, title string, text string, target_url string, pic_url string, icon_url string) int64 {
	call, _, _ := ApiSendLinkMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(title), str2ptr(text), str2ptr(target_url), str2ptr(pic_url), str2ptr(icon_url))
	return int64(call)
}

func SendMusicMsg(robot_wxid string, to_wxid string, name string, _type string) int64 {
	call, _, _ := ApiSendMusicMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(name), str2ptr(_type))
	return int64(call)
}

func SendTextMsg(robot_wxid string, to_wxid string, msg string) int64 {
	call, _, _ := ApiSendTextMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(msg))
	return int64(call)
}

func SendVideoMsg(robot_wxid string, to_wxid string, path string) int64 {
	call, _, _ := ApiSendVideoMsg.Call(int2ptr(int64(AuthCode)), str2ptr(robot_wxid), str2ptr(to_wxid), str2ptr(path))
	return int64(call)
}

func SetFatal() int64 {
	call, _, _ := ApiSetFatal.Call(int2ptr(int64(AuthCode)))
	return int64(call)
}
