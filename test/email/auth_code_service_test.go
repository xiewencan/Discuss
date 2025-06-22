package email

import (
	"discuss/internal/service/email"
	"testing"
)

func TestSendAuthCode(t *testing.T) {

	// 测试用例
	tests := []struct {
		name        string
		email       string
		authCode    string
		expectError bool
	}{
		{
			name:        "正常发送验证码",
			email:       "2022902706@chd.edu.cn",
			authCode:    "123456",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := email.SendAuthCode(tt.email, tt.authCode)
			if tt.expectError && err == nil {
				t.Errorf("期望出现错误，但是没有错误")
			}
			if !tt.expectError && err != nil {
				t.Errorf("期望成功，但是出现错误：%v", err)
			}
		})
	}
}
