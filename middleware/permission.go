package middleware

import (
	"backend_ft_tcs/constant"
	user2 "backend_ft_tcs/serveice/user"
	"github.com/gin-gonic/gin"
)

func UserSessionID(c *gin.Context, ua *user2.UserModel) (string, error) {
	claims := &CustomClaims{
		ID:   ua.Id,
		Name: ua.TrueName,
	}
	claims.Id = constant.JwtIdFront
	return SaveSession(c, claims)
}
