package gorm

import (
	"discuss/internal/dao"
	"discuss/internal/dto/request"
	"discuss/internal/dto/respond"
	"discuss/internal/model"
	//"discuss/internal/service/email"
	"discuss/pkg/constants"
	"discuss/pkg/zlog"
	//"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	//"time"
)

type userInfoService struct{}

var UserInfoService = new(userInfoService)

// Login 用户登录
func (u *userInfoService) Login(req request.LoginRequest) (string, *respond.LoginRespond, int) {
	password := req.Password
	email := req.Email
	var user model.UserInfo
	res := dao.GormDb.First(&user, "email = ?", email)
	if res.Error != nil {
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				message := "用户不存在,请先注册"
				zlog.Error(message)
				return message, nil, -2
			}
		}
		zlog.Error(res.Error.Error())
		return constants.SYSTEM_ERROR, nil, -1
	}

	if user.Password != password {
		message := "密码错误,请重新输入"
		zlog.Error(message)
		return message, nil, -2
	}

	loginRsp := &respond.LoginRespond{
		Uuid:      user.Uuid,
		Telephone: user.Telephone,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Birthday:  user.Birthday,
		Signature: user.Signature,
		IsAdmin:   user.IsAdmin,
		Status:    user.Status,
	}
	year, month, day := user.CreatedAt.Date()
	loginRsp.CreatedAt = fmt.Sprintf("%d-%02d-%02d", year, month, day)

	return "登陆成功", loginRsp, 0

}
