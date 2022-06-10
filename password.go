package goutils

import "golang.org/x/crypto/bcrypt"

// HashPassword 密码加密
// 注意：相同密码每次调用生成的密码是不一样的，需要使用CheckHashPassword进行密码校验
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckHashPassword 密码校验
func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
