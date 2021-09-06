package server

import "backend_ft_tcs/log"

var logger = log.GetLogger("server")

func (serverInstance *defaultServer) routers() {
	var mdw = UserAuthHandler(serverInstance)
	// ping模块, 对网关开放
	pingNone := serverInstance.engine.Group("/ping")
	pingNone.GET("", serverInstance.Pong)

	//*************************************************************************公开路由******************************************************************************//

	commonNone := serverInstance.engine.Group("/api/common")
	userComNone := commonNone.Group("/user")
	userComNone.POST("/register", serverInstance.UserCreate)
	userComNone.POST("/login", serverInstance.UserLogin)

	//**************************************************************************验证路由***************************************************************************//
	internalNone := serverInstance.engine.Group("/api/internal")

	userNone := internalNone.Group("/user")
	userNone.Use(mdw)
	userNone.POST("/search", serverInstance.UserSearch)
	userNone.GET("/detail", serverInstance.UserDetail)

}
