package server

import (
	"backend_ft_tcs/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Data    interface{} `json:"data"`    // 数据结果集
}

func ResponseOutput(c *gin.Context, code int, message []string, data interface{}) {
	var msg string
	if len(message) <= 0 {
		msg = constant.TranslateErrCode(code, message...)
	} else {
		msg = message[0]
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func responseError(c *gin.Context, code int, message []string) {
	ResponseOutput(c, code, message, nil)
}

func (this *defaultServer) ResponseSuccess(c *gin.Context, result interface{}, msg ...string) {
	ResponseOutput(c, constant.Success, msg, result)
}

func (this *defaultServer) ResponseError(c *gin.Context, code int, message ...string) {
	responseError(c, code, message)
}

func (this *defaultServer) InvalidParametersError(c *gin.Context) {
	responseError(c, constant.ParamError, nil)
}

func (this *defaultServer) InternalServiceError(c *gin.Context, message ...string) {
	responseError(c, constant.Failure, message)
}

func (this *defaultServer) NoPermissionError(c *gin.Context, message ...string) {
	responseError(c, constant.UserForbiddenError, message)
}
