package response

import "backend_ft_tcs/constant"

type Response struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg,omitempty" example:"success"`
	Data interface{} `json:"data,omitempty" example:""`
}

type Page struct {
	Records  interface{} `json:"records"`
	Total    uint32      `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}

func Success() *Response {
	return SuccessWith(nil)
}

func SuccessWith(data interface{}) *Response {
	return &Response{
		Code: constant.Success,
		Msg:  constant.TranslateErrCode(constant.Success),
		Data: data,
	}
}

func PageSuccess(Records interface{}, Total uint32, PageNum int, PageSize int) *Response {
	return SuccessWith(&Page{Records, Total, PageNum, PageSize})
}

func ParamError() *Response {
	return &Response{
		Code: constant.ParamError,
		Msg:  constant.TranslateErrCode(constant.ParamError),
	}
}

func FailBy(code int) *Response {
	return &Response{
		Code: code,
		Msg:  constant.TranslateErrCode(code),
	}
}
