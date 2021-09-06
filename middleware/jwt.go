package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"backend_ft_tcs/constant"
	"backend_ft_tcs/response"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		now := time.Now().Unix()
		if err == nil && claims.VerifyExpiresAt(now, false) && claims.VerifyIssuedAt(now, false) {
			key := fmt.Sprintf("%s_%d", claims.Id, claims.ID)
			// 前段必须每次更新token且做成同步访问
			if v, k := TokenMap[key]; k {
				if v == token {
					if token, err = j.RefreshToken(token); err == nil {
						c.Header("Authorization", token)
					}
				} else {
					c.JSON(http.StatusUnauthorized, response.FailBy(constant.AnotherClientLoginIn))
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, response.FailBy(http.StatusUnauthorized))
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, response.FailBy(http.StatusUnauthorized))
			c.Abort()
			return
		}
		c.Set(constant.ContextClaims, claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "candy"
	TokenMap                = make(map[string]string)
)

type CustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}
func GetSignKey() string {
	return SignKey
}
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 服务器持久化
	tokenStr, _ := token.SignedString(j.SigningKey)
	key := fmt.Sprintf("%s_%d", claims.Id, claims.ID)
	TokenMap[key] = tokenStr //会顶掉前面登录的账号
	return token.SignedString(j.SigningKey)
}
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
