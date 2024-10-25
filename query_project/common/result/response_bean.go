package result

type SuccessBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
type NullJson struct{}

func Success(data interface{}) *SuccessBean {
	return &SuccessBean{0, "OK", data}
}

type ErrorBean struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode int, errMsg string) *ErrorBean {
	return &ErrorBean{errCode, errMsg}
}
