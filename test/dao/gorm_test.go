package dao

import (
	"database/sql"
	"discuss/internal/dao"
	"discuss/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestInsertUserInfo 测试插入用户信息
func TestInsertUserInfo(t *testing.T) {

	// 准备插入的用户信息
	userInfo := &model.UserInfo{
		Uuid:          "U2025061412347",
		Nickname:      "小红",
		Telephone:     "13700000002",
		Email:         "xiaohong@example.com",
		Avatar:        "/static/avatars/xh.png",
		Gender:        0,
		Signature:     "喜欢旅行，热爱生活",
		Password:      "e10adc3949ba59ab",
		Birthday:      "19980308",
		CreatedAt:     time.Date(2025, 6, 14, 14, 22, 1, 0, time.Local),
		IsAdmin:       0,
		Status:        0,
		LastOnlineAt:  sql.NullTime{Valid: false},
		LastOfflineAt: sql.NullTime{Valid: false},
	}

	// 插入用户信息
	result := dao.GormDb.Create(userInfo)

	assert.NoError(t, result.Error, "插入用户信息时出错")

	// 验证插入是否成功
	var insertedUserInfo model.UserInfo
	result = dao.GormDb.First(&insertedUserInfo, "uuid = ?", userInfo.Uuid)
	assert.NoError(t, result.Error, "查询用户信息时出错")
	assert.Equal(t, userInfo.Uuid, insertedUserInfo.Uuid, "插入的用户信息与查询结果不一致")
}
