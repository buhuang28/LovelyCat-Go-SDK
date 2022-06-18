// 代码由简易GoVCL IDE自动生成。
// 不要更改此文件名
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/win"
)

//::private::
type TForm1Fields struct {
	subItemHit win.TLVHitTestInfo
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.SetCaption(f.TForm.Caption() + "带图片版")
}

func (f *TForm1) OnGetTaskButtonClick(sender vcl.IObject) {

}

func (f *TForm1) OnApplyButtonClick(sender vcl.IObject) {
}

func (f *TForm1) OnTaskListViewClick(sender vcl.IObject) {

}

func (f *TForm1) OnComboBox1Exit(sender vcl.IObject) {

}

func (f *TForm1) OnAddLoginApiButtonClick(sender vcl.IObject) {
	item := f.LoginApiListView.Items().Add()
	item.SetCaption("")
	item.SubItems().Add("http://127.0.0.1:8001/login")
}

func (f *TForm1) OnLoginApiListViewClick(sender vcl.IObject) {

}

func (f *TForm1) OnLoginApiEditExit(sender vcl.IObject) {

}
