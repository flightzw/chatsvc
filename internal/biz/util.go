package biz

import "github.com/flightzw/chatsvc/internal/utils/hash"

// 生成 hash 密码串
func generatePasswordStr(password string) string {
	salt := hash.GenerateSalt(16)
	hashPassword := hash.GenerateBcryptHashPassword(password, salt)
	return hashPassword + ":" + salt
}
