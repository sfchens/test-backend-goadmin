package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func Md5(s string, rawOutput ...bool) string {
	h := md5.New()
	h.Write([]byte(s))

	if len(rawOutput) > 0 && rawOutput[0] == true {
		return string(h.Sum(nil))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 计算字符串的 sha1 散列值
// 如果可选的 raw_output 参数被设置为 true， 那么 sha1 摘要将以 20 字符长度的原始二进制格式返回， 否则返回值为 40 字符长度的十六进制数
func Sha1(s string, rawOutput ...bool) string {
	o := sha1.New()
	o.Write([]byte(s))

	if len(rawOutput) > 0 && rawOutput[0] == true {
		return string(o.Sum(nil))
	}
	return hex.EncodeToString(o.Sum(nil))
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
