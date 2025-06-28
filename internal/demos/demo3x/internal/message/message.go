package message

// Response defile you response struct type
// 自定义的返回结构
type Response struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}
