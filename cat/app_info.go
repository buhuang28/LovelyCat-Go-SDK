package cat

import "encoding/json"

type AppInfo struct {
	Name string `json:"name"` //插件名,这里插件名字要和生成的插件名一致

	Desc string `json:"desc"` //描述简介

	Author string `json:"author"` //作者

	Version string `json:"version"` //插件版本

	ApiVersion string `json:"api_version"` //插件SDK版本号，默认不可修改

	MenuTitle string `json:"menu_title"` //打开插件按钮标题

	CoverBase64 string `json:"cover_base_64"` //插件logo,图片经过转码后的base64数据

	DeveloperKey string `json:"developer_key"` // 818198829 开发者key，默认不可修改
}

func (a AppInfo) ToJson() string {
	marshal, _ := json.Marshal(a)
	return Bytes2Str(marshal)
}
