package user

import (
	userModel "gobpframe/apps/demo/models/user"
	userService "gobpframe/apps/demo/services/user"
	"gobpframe/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"

	"gobpframe/errs"
	"gobpframe/resp"
)

type Controller struct{}

func NewController() *Controller {
	userService.AutoMigrate()
	return &Controller{}
}

// Create godoc
// @summary	新建用户
// @tags	用户
// @Security OAuth2Password
// @product json
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称"
// @Param password formData string true "密码"
// @Param gender formData string true "性别"
// @router	/user/create [PUT]
// @success 200 {object} string "返回状态"
func (ctrl *Controller) Create(ctx *gin.Context) {
	username := ctx.PostForm("username")
	nickname := ctx.PostForm("nickname")
	password := ctx.PostForm("password")
	genderStr := ctx.DefaultPostForm("gender", "1")

	if username == "" || password == "" {
		resp.Failed(ctx, errs.ErrReqParamMissing, []string{"username", "password"})
		return
	}

	gender, _ := strconv.Atoi(genderStr)
	user := &userModel.User{
		Username: username,
		Nickname: nickname,
		Password: password,
		Gender:   gender,
	}
	logger.Debug("create user:", user)
	user = userService.Create(user)
	resp.Success(ctx, user)
}

// Login godoc
// @summary	登录
// @tags	用户
// @Security OAuth2Password
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "返回状态"
// @Router /user/login [post]
func (ctrl *Controller) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "" || password == "" {
		resp.Failed(ctx, errs.ErrReqParamMissing, []string{"username", "password"})
		return
	}

	logged, usr := userService.Login(username, password)
	if !logged {
		resp.Failed(ctx, errs.ErrUserPassword, map[string]interface{}{username: username, password: password})
		return
	}
	resp.Success(ctx, usr)
}
