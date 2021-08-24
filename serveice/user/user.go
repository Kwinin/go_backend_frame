package user

import (
	"backend_ft_tcs/utils"
	_ "fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserModel struct {
	utils.BaseModel
	TrueName  string    `gorm:"index;comment:'真实姓名'" json:"true_name"` // 真实姓名
	Mobile    string    `gorm:"index;comment:'手机号码'" json:"mobile"`    // 手机号码
	Password  string    `gorm:"index;comment:'登录密码'" json:"password"`  // 登录密码
	IdCard    string    `gorm:"index;comment:'身份证'" json:"id_card"`    // 身份证
	Type      string    `gorm:"comment:'用户类型'" json:"type"`            // 用户类型
	Status    int8      `gorm:"comment:'1:可用 2：已删除'" json:"status"`    // 1:可用 2：已删除
	LastLogin time.Time `gorm:"comment:'上次登录时间'" json:"last_login"`    // 上次登录时间
	ThisLogin time.Time `gorm:"comment:'本次登录时间'" json:"this_login"`    // 本次登录时间
}

func newUserModel() *UserModel {
	um := new(UserModel)
	return um
}

func NewUserModel() *UserModel {
	return newUserModel()
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) BeforeCreate() (err error) {
	u.CreatedAt = time.Now()
	return nil
}

func (u *UserModel) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now()
	return nil
}

func InitUser(db *gorm.DB) error {
	var err error

	if db.HasTable(&UserModel{}) {
		err = db.AutoMigrate(&UserModel{}).Error
	} else {
		err = db.CreateTable(&UserModel{}).Error
	}

	return err
}

func NewUserModelWithMobile(db *gorm.DB, mobile string) *UserModel {
	r := newUserModel()
	r.Mobile = mobile
	notFound := db.Table(r.TableName()).
		Select("*").
		Where("mobile = ?", mobile).
		Where("status = 1").
		First(r).
		RecordNotFound()
	if notFound {
		return nil
	}
	return r
}

func (u *UserModel) Create(db *gorm.DB) error {
	return db.Model(&UserModel{}).Create(u).Error
}

func (u *UserModel) UpdateMap(db *gorm.DB, update map[string]interface{}) error {
	return db.Model(&UserModel{}).Where("id=?", u.Id).Update(update).Error
}

func (u *UserModel) Update(db *gorm.DB, update UserModel) error {
	return db.Model(&UserModel{}).Where("id=?", u.Id).Update(update).Error
}

func NewUserModelWithId(db *gorm.DB, id int) *UserModel {
	r := newUserModel()
	r.Id = id
	notFound := db.Table(r.TableName()).
		Select("*").
		Where("id = ?", id).
		Where("status = 1").
		First(r).
		RecordNotFound()
	if notFound {
		return nil
	}
	return r
}

type UserModelList []UserModel

func ListUserModelWithUserIds(db *gorm.DB, userIds []int, search string, limit, offset int) (*UserModelList, int, error) {
	userModel := newUserModel()
	userModelList := make(UserModelList, 0, limit)
	sql := db.Table(userModel.TableName()).
		Select("*").Where("status = ?", 1)
	if userIds != nil && len(userIds) > 0 {
		sql = sql.Where("id in (?)", userIds)
	}
	if search != "" {
		sql = sql.Where("true_name LIKE ? OR mobile LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	var total int
	err := sql.Count(&total).
		Limit(limit).
		Offset(offset).
		Scan(&userModelList).Error
	return &userModelList, total, err
}

func (uList *UserModelList) UserIds() []int {
	userIds := make([]int, 0, len(*uList))
	for _, v := range *uList {
		userIds = append(userIds, v.Id)
	}
	return userIds
}
