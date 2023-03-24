package kit

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 密码加密
func Encrypt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// Decrypt 密码验证
func Decrypt(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
