package server

import (
	"backend_ft_tcs/middleware"
	"github.com/gin-gonic/gin"
)

func UserAuthHandler(server *defaultServer) gin.HandlerFunc {
	//return func(context *gin.Context) {
	//	userId := context.GetHeader(XUSERID)
	//	if userId == "" {
	//		logrus.Error("middleware -> user:auth forbidden")
	//		server.ResponseError(context, FORBIDDEN)
	//		context.Abort()
	//		return
	//	}
	//
	//	id, err := strconv.Atoi(userId)
	//	if err != nil {
	//		server.InternalServiceError(context, err.Error())
	//		context.Abort()
	//		return
	//	}
	//
	//	service := user.NewUserService(server.db)
	//	users, err := service.FirstOrCreateUser(uint(id))
	//	if err != nil {
	//		logrus.Error("middleware ->userInit:", err.Error())
	//		server.InternalServiceError(context, err.Error())
	//		context.Abort()
	//		return
	//	}
	//	context.Set(UserCookie, &UserClaims{
	//		UserId: users.TritiumID,
	//		UUId:   users.UUID,
	//	})
	//}
	return middleware.GetSessionHandler()
}
