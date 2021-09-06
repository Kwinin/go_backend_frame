package user

import (
	"backend_ft_tcs/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserService interface {
	// 获取用户信息接口
	//
	// 子平台获取用户信息接口
	GetUserDetail(userId int) *UserModel
	// 用户模块
	//
	// 用户登录
	UserLogin(mobile, password, SecretKey string) (*UserModel, error)
	// 创建用户
	UserCreate(trueName, mobile, password, idCard string) (*UserModel, error)
	// 查找用户
	UserSearch(userIds []int, search string, limit, offset int) (*UserModelList, int, error)
	// 删除用户
	UserDelete(userId int) error
	// 编辑用户
	UserEdit(userId int, data UserModel) error
	// 账户中心
	//
	// 获取账户信息
	AccountDetail(userId int) (*UserModel, error)
	// 编辑账户中心信息
	AccountEdit(userId int, mobile, trueName string) error
	// 修改账户密码
	AccountPassword(userId int, oldPassword, newPassword, secret string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	u := new(userService)
	u.db = db
	return u
}

func (us *userService) GetUserDetail(userId int) *UserModel {
	return NewUserModelWithId(us.db, userId)
}

func (us *userService) UserLogin(mobile, password, SecretKey string) (*UserModel, error) {
	user := NewUserModelWithMobile(us.db, mobile)
	if user == nil {
		return nil, fmt.Errorf("%s", "未找到该用户")
	}

	if utils.CheckPassword(user.Password, password) {
		// 修改本次登录时间和上次登录时间
		var err error
		updateData := map[string]interface{}{
			"last_login": user.ThisLogin,
			"this_login": time.Now(),
		}
		err = user.UpdateMap(us.db, updateData)
		if err != nil {
			return nil, fmt.Errorf("%s", "数据库操作:更新登录时间失败")
		}
		return user, nil
	}

	return nil, fmt.Errorf("%s", "密码不正确")

}

func (us *userService) UserCreate(trueName, mobile, password, idCard string) (*UserModel, error) {
	user := NewUserModelWithMobile(us.db, mobile)
	if user != nil {
		return nil, fmt.Errorf("%s", "账户已存在")
	}
	newUser := NewUserModel()

	encodePwd, err := utils.CryptoPassword(password)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	newUser.TrueName = trueName
	newUser.Mobile = mobile
	newUser.Password = encodePwd
	newUser.Status = 1
	newUser.IdCard = idCard
	newUser.LastLogin = time.Now()
	newUser.ThisLogin = time.Now()
	err = newUser.Create(us.db)
	if err != nil {
		return nil, fmt.Errorf("数据库错误: %s", err.Error())
	}
	return newUser, nil

}

func (us *userService) UserSearch(userIds []int, search string, limit, offset int) (*UserModelList, int, error) {
	list, total, err := ListUserModelWithUserIds(us.db, userIds, search, limit, offset)
	return list, total, err

}

func (us *userService) UserDelete(userId int) error {
	user := NewUserModelWithId(us.db, userId)
	if user == nil {
		return fmt.Errorf("%s", "账户不存在")
	}
	updateData := map[string]interface{}{}
	updateData["status"] = 0
	return user.UpdateMap(us.db, updateData)
}

func (us *userService) UserEdit(userId int, data UserModel) error {
	user := NewUserModelWithId(us.db, userId)
	if user == nil {
		return fmt.Errorf("%s", "账户不存在")
	}
	updateData := UserModel{
		TrueName: data.TrueName,
		Mobile:   data.Mobile,
		IdCard:   data.IdCard,
		Type:     data.Type,
		Status:   data.Status,
	}
	return user.Update(us.db, updateData)
}

func (us *userService) AccountDetail(userId int) (*UserModel, error) {
	user := NewUserModelWithId(us.db, userId)
	if user == nil {
		return nil, fmt.Errorf("%s", "账户不存在或账号已被删除")
	}
	return user, nil
}

func (us *userService) AccountEdit(userId int, mobile, trueName string) error {
	user := NewUserModelWithId(us.db, userId)
	if user == nil {
		return fmt.Errorf("%s", "账户不存在")
	}
	oldUser := NewUserModelWithMobile(us.db, mobile)
	if oldUser != nil {
		return fmt.Errorf("%s", "此账号已存在")
	}
	updateData := map[string]interface{}{}
	updateData["true_name"] = trueName
	updateData["mobile"] = mobile
	return user.UpdateMap(us.db, updateData)
}

func (us *userService) AccountPassword(userId int, oldPassword, newPassword, secret string) error {
	user := NewUserModelWithId(us.db, userId)
	if user == nil {
		return fmt.Errorf("%s", "账户不存在")
	}
	if user.Password != utils.Password(oldPassword, secret) {
		return fmt.Errorf("%s", "密码不正确")
	}
	updateData := map[string]interface{}{}
	updateData["password"] = utils.Password(newPassword, secret)
	return user.UpdateMap(us.db, updateData)
}
