package route

import (
	"backend_ft_tcs/utils"
	"github.com/jinzhu/gorm"
)

type Route struct {
	utils.BaseModel
	Name string `gorm:"size:64;NOT NULL;unique_index" json:"name"`
}

func (Route) TableName() string {
	return "route"
}

func InitRoute(db *gorm.DB) error {
	var err error

	if db.HasTable(&Route{}) {
		err = db.AutoMigrate(&Route{}).Error
	} else {
		err = db.CreateTable(&Route{}).Error
	}

	return err
}

func dropRoute(db *gorm.DB) {
	db.DropTableIfExists(&Route{})
}
