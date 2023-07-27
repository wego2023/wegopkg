package base

// Response
// @Description: 响应体
type Response struct {
	Code int         `json:"code"` // 0:成功，其他：失败
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 数据
}

// NewResponse
// @Description: 创建响应体
// @return *Response
func NewResponse() *Response {
	return &Response{}
}

// Fail
// @Description: 响应错误
// @receiver: base
// @param: code
// @param: err
func (response *Response) Fail(code int, err string) *Response {
	response.Code = code
	response.Msg = err
	return response
}

// Fail
// @Description: 响应错误
// @param: code
// @param: err
// @return *Response
func Fail(code int, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
	}
}

// FailData
// @Description: 响应错误
// @receiver: response
// @param: code
// @param: data
// @param: err
func (response *Response) FailData(code int, data interface{}, err string) *Response {
	response.Code = code
	response.Msg = err
	response.Data = data
	return response
}

// FailData
// @Description: 响应错误
// @param: code
// @param: data
// @param: err
// @return *Response
func FailData(code int, data interface{}, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
		Data: data,
	}
}

// Success
// @Description: 响应成功
// @receiver: response
// @param: data
func (response *Response) Success() *Response {
	response.Code = 0
	response.Data = ""
	response.Msg = "success"
	return response
}

// Success
// @Description: 响应成功
// @return *Response
func Success() *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: "",
	}
}

// OK
// @Description: 响应成功
// @receiver: response
// @param: data
func (response *Response) OK(data interface{}) *Response {
	response.Code = 0
	response.Data = data
	response.Msg = "success"
	return response
}

// OK
// @Description: 响应成功
// @param: data
// @return *Response
func OK(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

// OkMsg
// @Description: 响应成功
// @receiver: response
// @param: data
// @param: err
func (response *Response) OkMsg(data interface{}, err string) *Response {
	response.Code = 0
	response.Data = data
	response.Msg = err
	return response
}

// OkMsg
// @Description: 响应成功
// @param: data
// @param: err
// @return *Response
func OkMsg(data interface{}, err string) *Response {
	return &Response{
		Code: 0,
		Msg:  err,
		Data: data,
	}
}
