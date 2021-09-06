package server

import (
	"backend_ft_tcs/constant"
	"backend_ft_tcs/middleware"
	"backend_ft_tcs/serveice/user"
	"backend_ft_tcs/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 获取用户信息
type UserCreateArgs struct {
	TrueName string `json:"true_name" form:"true_name" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	IdCard   string `json:"id_card" form:"id_card" `
}

// @Description 创建用户接口
// @Summary 创建用户接口
// @Author kwinin
// @Tags 用户模块
// @Accept  json
// @Produce json
// @Router /common/user/register [post]
// @Param true_name formData string false "姓名"
// @Param mobile formData string true "电话"
// @Param password formData string true "密码"
// @Param id_card formData string false "身份证"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) UserCreate(ctx *gin.Context) {
	args := new(UserCreateArgs)
	if err := ctx.ShouldBind(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}
	if err := ValidatorInstance().Struct(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}
	//if ok := utils.CheckPassword(args.Password); !ok {
	//	s.ResponseError(ctx, 1, "密码由8~20数字和字母组成，必须包含大写字母")
	//	return
	//}
	//if ok := utils.CheckPhone(args.Mobile); !ok {
	//	s.ResponseError(ctx, 1, "请输入正确的手机号码")
	//	return
	//}

	userService := user.NewUserService(s.db)
	_, err := userService.UserCreate(args.TrueName, args.Mobile, args.Password, args.IdCard)
	if err != nil {
		s.ResponseError(ctx, constant.Failure, err.Error())
		return
	}
	s.ResponseSuccess(ctx, nil)
	return
}

type UserLoginArgs struct {
	Mobile   string `json:"mobile" form:"mobile" example:"12345678"`
	Password string `json:"password" form:"password" example:"123456"`
}
type UserLoginVO struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

// @Description 登陆接口
// @Summary 登陆接口
// @Author kwinin
// @Date 08/23/20
// @Tags 用户模块
// @Accept  json
// @Produce json
// @Router /common/user/login [post]
// @Param mobile body UserLoginArgs true "账号"
// @Success 200 {object} UserLoginVO "登录"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) UserLogin(ctx *gin.Context) {
	args := new(UserLoginArgs)
	if err := ctx.ShouldBind(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}
	var token string
	if err := ValidatorInstance().Struct(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}

	vo := UserLoginVO{}
	userService := user.NewUserService(s.db)
	ua, err := userService.UserLogin(args.Mobile, args.Password, s.conf.Secret)
	if err != nil {
		s.ResponseError(ctx, constant.UserLoginError, "账号密码不正确")
		return
	}

	if token, err = middleware.UserSessionID(ctx, ua); err != nil {
		logger.Errorf("%+v", err.Error())
		s.ResponseError(ctx, constant.DataErr)
		return
	}
	if err != nil {
		s.ResponseError(ctx, constant.UserLoginError, "账号密码不正确")
		return
	}

	vo.UserId = ua.Id
	vo.Token = token

	s.ResponseSuccess(ctx, vo)
	return
}

type UserSearchArgs struct {
	Search   string `json:"search" form:"search"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required"`
	PageNum  int    `json:"page_num" form:"page_num" binding:"required"`
}

type UserSearchItemVO struct {
	UserID    int    `json:"user_id"`
	TrueName  string `json:"true_name"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"created_at"`
}
type UserSearchListVO []UserSearchItemVO

type UserSearchVO struct {
	List  UserSearchListVO `json:"list"`
	Total int              `json:"total"`
}

// @Description 查询用户列表接口
// @Summary 查询用户列表接口
// @Author kwinin
// @Tags 用户模块
// @Accept  json
// @Produce json
// @Router /internal/user/search [post]
// @Param search formData string false "根据手机号和姓名进行模糊查询"
// @Param page_size formData int true "页数"
// @Param page_num formData int true "页码"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) UserSearch(ctx *gin.Context) {
	args := new(UserSearchArgs)
	if err := ctx.ShouldBind(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}
	if err := ValidatorInstance().Struct(args); err != nil {
		logger.Errorf("%+v", err.Error())
		s.InvalidParametersError(ctx)
		return
	}
	if args.PageSize > 100 {
		s.InvalidParametersError(ctx)
		return
	}

	userService := user.NewUserService(s.db)
	list, total, err := userService.UserSearch(nil, args.Search, args.PageSize, args.PageSize*(args.PageNum-1))
	if err != nil {
		s.ResponseError(ctx, constant.DBError, "数据库操作:"+err.Error())
		return
	}
	listVO := make(UserSearchListVO, 0, args.PageNum)
	for _, v := range *list {

		vo := UserSearchItemVO{
			UserID:    v.Id,
			TrueName:  v.TrueName,
			Mobile:    v.Mobile,
			CreatedAt: utils.FormatTime(v.CreatedAt),
		}
		listVO = append(listVO, vo)
	}
	userSearchVO := UserSearchVO{
		List:  listVO,
		Total: total,
	}
	s.ResponseSuccess(ctx, userSearchVO)
	return
}

type UserDetailVO struct {
	Mobile   string `json:"mobile" form:"mobile"` // 手机号码
	TrueName string `json:"true_name" form:"true_name"`
}

// @Description 获取用户信息接口
// @Summary 获取用户信息接口
// @Author kwinin
// @Tags 用户模块
// @Accept  x-www-form-urlencoded
// @Produce json
// @Router /internal/user/detail [get]
// @Success 200 {object} UserDetailVO "detail"
// @See https://github.com/swaggo/swag#api-operation
func (s *defaultServer) UserDetail(ctx *gin.Context) {
	claims := ctx.MustGet(constant.ContextClaims).(*middleware.CustomClaims)

	userService := user.NewUserService(s.db)
	userModel := userService.GetUserDetail(claims.ID)
	if userModel == nil {
		s.ResponseError(ctx, constant.CanNotFindAuth, "未找到用户")
		return
	}

	vo := UserDetailVO{
		Mobile:   userModel.Mobile,
		TrueName: userModel.TrueName,
	}
	s.ResponseSuccess(ctx, vo)
	return
}
