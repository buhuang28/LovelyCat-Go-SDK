package main

import (
	"LovelyCat-sdk/cat"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
	"runtime"
	"time"
)

func init() {
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		// 在go的dll中应该使用专用的初始函数
		// 此方法也只适合第三方非govcl程序，不适合go+govcl+godll+govcl方式
		rtl.InitGoDll(0) //0则自动获取当前线程Id
		vcl.Application.Icon().SetHandle(win.LoadIcon2(win.GetSelfModuleHandle(), "DEFAULT_ICON"))
		// 这。。。。。。有点无语啊
		go func() {
			time.Sleep(time.Second * 3)
			vcl.ThreadSync(func() {
				// 新建一个窗口
				if Form1 == nil {
					Form1 = NewForm1(nil)
					Form1.SetShowInTaskBar(types.StAlways)
					Form1.SetFormStyle(types.FsStayOnTop)
				}
			})
		}()
		vcl.Application.Run()
	}()

	cat.Name = "demo"
	cat.Desc = "不慌的SDK"
	cat.Author = "不慌"
	cat.Version = "1.0.2"
	cat.ApiVersion = "4.1"
	cat.MenuuTitle = "设置"
	cat.CoverBase64 = "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"
	cat.DeveloperKey = "818198829"

	cat.EventInit = EventInit
	cat.EventEnable = EventEnable
	cat.EventStop = EventStop
	cat.EventLogin = EventLogin
	cat.EventGroupMsg = EventGroupMsg
	cat.EventFriendMsg = EventFriendMsg
	cat.EventReceivedTransfer = EventReceivedTransfer
	cat.EventScanCashMoney = EventScanCashMoney
	cat.EventFriendVerify = EventFriendVerify
	cat.EventGroupMemberAdd = EventGroupMemberAdd
	cat.EventGroupMemberDecrease = EventGroupMemberDecrease
	cat.EventSysMsg = EventSysMsg
	cat.Menu = Menu
}

func EventInit() int32 {
	return 0
}

func EventEnable() int32 {
	return 0
}

func EventStop() int32 {
	return 0
}

func EventLogin(rob_wxid string, rob_wxname string, _type int32) int32 {
	return 0
}

func EventGroupMsg(robot_wxid string, _type int32, from_wxid string, from_name string, final_from_wxid string, final_from_name string, to_wxid string, Msg string) int32 {
	return 0
}

func EventFriendMsg(robot_wxid string, _type int32, from_wxid string, from_name string, to_wxid string, Msg string) int32 {
	return 0
}

func EventReceivedTransfer(robot_wxid string, from_wxid string, from_name string, to_wxid string, money string, Msg_json string) int32 {
	return 0
}

func EventScanCashMoney(robot_wxid string, pay_wxid string, pay_name string, money string, Msg_json string) int32 {
	return 0
}

func EventFriendVerify(robot_wxid string, from_wxid string, from_name string, to_wxid string, Msg_json string) int32 {
	return 0
}

func EventGroupMemberAdd(robot_wxid string, from_wxid string, from_name string, Msg_json string) int32 {
	return 0
}

func EventGroupMemberDecrease(robot_wxid string, from_wxid string, from_name string, Msg_json string) int32 {
	return 0
}

func EventSysMsg(robot_wxid string, _type int32, Msg_json string) int32 {
	return 0
}

func Menu() int32 {
	vcl.ThreadSync(func() {
		if Form1 != nil {
			Form1.Show()
		}
	})
	cat.AppendLogs("滑稽啦啦啦")
	return 0
}
