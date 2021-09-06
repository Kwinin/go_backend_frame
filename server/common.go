package server

import (
	"backend_ft_tcs/route"
	"backend_ft_tcs/serveice/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/url"
)

func GetDBConnection(conf *DBConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s"
	dsn := fmt.Sprintf(format, conf.User, conf.Password, conf.Address, conf.Port, conf.DbName, conf.Charset, url.QueryEscape(conf.Loc))
	logrus.Infof("dsn=%s", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	InitModel(db)
	db.LogMode(conf.LogMode)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetMaxOpenConns(conf.MaxOpen)
	return db, nil
}

type JwtCustomClaims struct {
	UserId      uint
	LoginVerify int
	jwt.StandardClaims
}

type UserClaims struct {
	UserId uint   `json:"user_id"`
	UUId   string `json:"uu_id"`
}

func InitModel(db *gorm.DB) error {
	var err error
	err = user.InitUser(db)
	if err != nil {
		logrus.Fatal("Init db user failed, ", err)
		return err
	}

	err = route.InitRoute(db)
	if err != nil {
		logrus.Fatal("Init db route failed, ", err)
		return err
	}
	return err

}
