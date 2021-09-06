package middleware

import (
	"backend_ft_tcs/constant"
	"backend_ft_tcs/log"
	"backend_ft_tcs/middleware/sessions"
	"backend_ft_tcs/response"
	user2 "backend_ft_tcs/serveice/user"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

var logger = log.GetLogger("session")

var (
	AuthHeaderName = "Authorization"
	MaxAge         = 3600 * 24 * 7 // 1hour
)

//暂不细分权限
func GetSessionHandler() gin.HandlerFunc {
	return EntSession
}

func EntSession(c *gin.Context) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	store, _ := sessions.NewMySQLStoreFromConnection(db.DB(), "session", "/", MaxAge, []byte("something-very-secret"))
	defer store.Close()
	session, err := store.Get(c.Request, AuthHeaderName)
	if err != nil || session.Values["ID"] == nil {
		logger.Errorf("session id is %+v ", session.Values["ID"])
		c.JSON(http.StatusUnauthorized, response.FailBy(http.StatusUnauthorized))
		c.Abort()
		return
	}
	requestUrl := c.Request.URL.Path
	if session.Values["Id"].(string) == constant.JwtIdFront {
		//user status
		//GetStatus(c, session.Values["ID"].(uint64))
		//user不能访问后台管理接口
		if strings.HasPrefix(requestUrl, "/cms/admin") {
			c.JSON(http.StatusMethodNotAllowed, response.FailBy(http.StatusMethodNotAllowed))
			c.Abort()
			return
		}
	}

	claims := CustomClaims{
		ID:   session.Values["ID"].(int),
		Name: session.Values["Name"].(string),
	}
	claims.Id = session.Values["Id"].(string)
	c.Set(constant.ContextClaims, &claims)
	c.Next()
}

func GetStatus(c *gin.Context, userId uint64) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var ua user2.UserModel
	if err = tx.Where("id=?", userId).First(&ua).Error; err != nil {
		c.JSON(http.StatusOK, response.FailBy(constant.StatusUserLoginError))
		return
	}

	if ua.Status != 1 {
		c.JSON(http.StatusMethodNotAllowed, response.FailBy(constant.UserAccountForbid))
		c.Abort()
		return
	}
}

func getUser(c *gin.Context, userId uint64) (*user2.UserModel, error) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	var err error
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var ua user2.UserModel
	if err = tx.Where("id=?", userId).First(&ua).Error; err != nil {
		return nil, err
	}

	return &ua, err
}

func GetSessionInfo(c *gin.Context) (*CustomClaims, error) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	store, _ := sessions.NewMySQLStoreFromConnection(db.DB(), "session", "/", MaxAge, []byte("something-very-secret"))
	defer store.Close()
	session, err := store.Get(c.Request, AuthHeaderName)
	if err != nil {
		return nil, err
	}
	if session.Values["ID"] == nil {
		return nil, errors.New("not found user info")
	}
	claims := CustomClaims{
		ID:   session.Values["ID"].(int),
		Name: session.Values["Name"].(string),
	}
	claims.Id = session.Values["Id"].(string)
	return &claims, nil
}

func SaveSession(c *gin.Context, claims *CustomClaims) (token string, err error) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	store, _ := sessions.NewMySQLStoreFromConnection(db.DB(), "session", "/", MaxAge, []byte("something-very-secret"))
	defer store.Close()
	session, err := store.Get(c.Request, AuthHeaderName)
	if err != nil {
		logger.Errorf("%+v", err.Error())
		return "", err
	}
	session.Values["ID"] = claims.ID
	session.Values["Name"] = claims.Name
	session.Values["Id"] = claims.Id
	session.Values["identifier"] = fmt.Sprintf("%s_%d", claims.Id, claims.ID)
	err = session.Save(c.Request, c.Writer)
	return session.ID, err
}

func DelSession(c *gin.Context, claims *CustomClaims) (err error) {
	db := c.MustGet(constant.ContextDb).(*gorm.DB)
	store, _ := sessions.NewMySQLStoreFromConnection(db.DB(), "session", "/", MaxAge, []byte("something-very-secret"))
	defer store.Close()
	session, err := store.Get(c.Request, AuthHeaderName)
	if err != nil {
		return err
	}
	if claims.ID == session.Values["ID"] {
		storeErr := store.Delete(c.Request, c.Writer, session)
		if storeErr != nil {
			return storeErr
		}
	}
	return nil
}
