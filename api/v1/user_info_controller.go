package v1

import (
	"discuss/internal/dto/request"
	"discuss/internal/service/gorm"
	"discuss/pkg/constants"
	"discuss/pkg/zlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register注册
// func Register(c *gin.Context) {
// 	var req request.RegisterRequest
// 	if err := c.BindJSON(&req); err != nil {
// 		zlog.Error(err.Error())
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    500,
// 			"message": constants.SYSTEM_ERROR,
// 		})
// 		return
// 	}
// 	fmt.Println(req)
// 	message, userInfo, ret := gorm.UserInfoService.Register(req)
// 	JsonBack(c, message, ret, userInfo)
// }

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	fmt.Println(req)
	message, userInfo, ret := gorm.UserInfoService.Login(req)
	JsonBack(c, message, ret, userInfo)
}
