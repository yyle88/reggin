package reghan

type ResponseType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func MakeResponse[resType any](res *resType, erx error) *ResponseType {
	if erx != nil {
		return &ResponseType{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ResponseType{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}
