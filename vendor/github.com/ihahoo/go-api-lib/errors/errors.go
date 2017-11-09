package errors

// E 接口错误信息
type E struct {
	Status int
	Code   int
	Msg    string
}

// Err 调用接口返回的错误信息
type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
