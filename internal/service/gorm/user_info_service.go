package gorm

import (
	"discuss/internal/dao"
	"discuss/internal/dto/request"
	"discuss/internal/dto/respond"
	"discuss/internal/model"
	"discuss/internal/utils/random"

	emailService "discuss/internal/service/email"
	myredis "discuss/internal/service/redis"
	"discuss/pkg/constants"
	"discuss/pkg/zlog"

	// redis "github.com/go-redis/redis/v8"
	//"encoding/json"
	"errors"
	"fmt"
	"strconv"
	//"github.com/ugorji/go/codec"
	"gorm.io/gorm"
	"time"
)

type userInfoService struct{}

var UserInfoService = new(userInfoService)

// checkEmailExist 检查邮箱是否存在
func (u *userInfoService) checkEmailExist(email string) (model.UserInfo, int) {
	var user model.UserInfo
	res := dao.GormDb.First(&user, "email = ?", email)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return model.UserInfo{}, 0 // 邮箱不存在
		}
		zlog.Error(res.Error.Error())
		return model.UserInfo{}, -1 // 系统错误
	}
	return user, 1 // 邮箱存在
}

// Login 用户登录
func (u *userInfoService) Login(req request.LoginRequest) (string, *respond.LoginRespond, int) {
	password := req.Password
	email := req.Email
	var user model.UserInfo
	user, exist := u.checkEmailExist(email)
	switch exist {
	case -1:
		message := constants.SYSTEM_ERROR
		zlog.Error(message)
		return message, nil, -1
	case 0:
		message := "用户不存在,请先注册"
		zlog.Error(message)
		return message, nil, -2
	case 1:
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
	default:
		message := "未知错误,请稍后再试"
		zlog.Error(message)
		return message, nil, -1
	}

}

// EmailLogin 邮箱验证码登录
func (u *userInfoService) EmailLogin(req request.EmailLoginRequest) (string, *respond.LoginRespond, int) {
	var user model.UserInfo
	email := req.Email

	user, exist := u.checkEmailExist(email)
	switch exist {
	case -1:
		message := constants.SYSTEM_ERROR
		zlog.Error(message)
		return message, nil, -1
	case 0:
		message := "用户不存在,请先注册"
		zlog.Error(message)
		return message, nil, -2
	case 1:
		key := fmt.Sprintf("email_code_%s", email)
		code, err := myredis.GetKey(key)
		emailcode := req.EmailCode
		if err != nil {
			message := constants.SYSTEM_ERROR
			zlog.Error(message)
			return message, nil, -1
		}
		if emailcode == "" {
			message := "验证码不能为空"
			zlog.Info(message)
			return message, nil, -2
		}
		if code != emailcode {
			message := "验证码错误,请重新输入"
			zlog.Info(message)
			return message, nil, -2
		} else {
			if err := myredis.DelKeyIfExists(key); err != nil {
				zlog.Error(err.Error())
				return constants.SYSTEM_ERROR, nil, -1
			}
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
		loginRsp.CreatedAt = fmt.Sprintf("%d.%d.%d", year, month, day)
		return "登陆成功", loginRsp, 0

	default:
		message := "未知错误,请稍后再试"
		zlog.Error(message)
		return message, nil, -1
	}
}

func (u *userInfoService) SendEmailCode(email string) (string, int) {
	code := strconv.Itoa(random.GetRandomInt(6))
	key := fmt.Sprintf("email_code_%s", email)
	if err := myredis.SetKeyEx(key, code, time.Minute*3); err != nil {
		zlog.Error(err.Error())
		return constants.SYSTEM_ERROR, -1
	}
	if err := emailService.SendAuthCode(email, code); err != nil {
		zlog.Error(err.Error())
		return constants.SYSTEM_ERROR, -1
	}
	return "验证码发送成功", 0
}

// func (u *userInfoService) checkEmailExist(email string) bool {
// 	var user model.UserInfo
// 	res := dao.GormDb.First(&user, "email = ?", email)
// 	if res.Error != nil {
// 		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
// 			return false
// 		}
// 		return false
// 	}
// 	return true
// }

// func (u *userInfoService) checkEmailVaild(email string)       {}
// func (u *userInfoService) checkNicknameValid(nickname string) {}

// func (u *userInfoService) Register(req request.RegisterRequest) (string, *respond.LoginRespond, int) {

// }
