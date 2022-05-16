package model

import "time"

// xAPI 学习行为数据规范

// 行为数据内容 statement = actor+verb+object(+result+context+authority+timestamp)

// actor
// 	uuid 用户唯一标识符
// verb
// 	login 登录
// 	navigate 浏览
// 	practice 做题
// 	examine 考试
// 	logout 登出
// object
// 	module 功能模块
// 	question 习题
// 	wrongQuestion 错题
// result
// context
// 	ip ip地址
// 	device 设备信息
// authority
// timestamp 时间戳

type XAPIData struct {
	Actor     XActor    `json:"actor"`
	Verb      string    `json:"verb"` // login login-oauth-ourspark logout navigate study practice examine
	Object    XObject   `json:"object"`
	Result    struct{}  `json:"result"`
	Context   XContext  `json:"context"`
	Authority struct{}  `json:"authority"`
	Timestamp time.Time `json:"timestamp"`
}

type XActor struct {
	UUID string `json:"uuid"`
}

type XObject struct {
	Info string `json:"info"`
}

type XContext struct {
	IP     string `json:"ip"`
	Device string `json:"device"`
}
