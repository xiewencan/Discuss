package email

import (
	"discuss/internal/config"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendAuthCode(email string, authCode string) error {
	conf := config.GetConfig().EmailConfig
	fmt.Printf("%+v\n", conf)
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Discuss注册")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是: %s", authCode))

	d := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
