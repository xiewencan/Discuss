package valid

import (
	"regexp"
)

func IsNicknameValid(nickname string) bool {
	if len(nickname) < 2 || len(nickname) > 10 {
		return false
	}
	// 使用正则表达式匹配名字
	pattern := `^[\p{Han}a-zA-Z_]+$`
	matched, _ := regexp.MatchString(pattern, nickname)
	return matched
}

func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func IsPasswordValid(password string) bool {
	return true
}
