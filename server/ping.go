package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// 验证单例
func ValidatorInstance() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

type PingVO struct {
	RealIP        string `json:"real_ip"` // 远程客户端ip地址
	UserId        string `json:"userId"`  // 用户id
	Authorization string `json:"authorization"`
}

func (s *defaultServer) Pong(ctx *gin.Context) {
	vo := PingVO{
		RealIP: "127.0.0.1",
		UserId: "10001",
	}
	s.ResponseSuccess(ctx, vo)
	return
}
