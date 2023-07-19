package my_jwt

import (
	"csf/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"

	"csf/library/viper"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// BaseClaims 这里可根据需求来修改
type BaseClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte // 签名的发行者
}

// NewJWT 创建 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(viper.NewViper.GetString("jwt.signingKey")),
	}
}

// CreateClaims 创建 Claims
func (j *JWT) CreateClaims(baseClaims BaseClaims) MyClaims {
	claims := MyClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                                            // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(viper.NewViper.GetInt("jwt.expiresTime")), // 过期时间
			Issuer:    viper.NewViper.GetString("jwt.issuer"),                              // 签名的发行者
		},
	}
	return claims
}

// GenerateToken 根据给定的 username 和 password 创建 token
func (j *JWT) GenerateToken(username, password string) (string, error) {
	claims := j.CreateClaims(BaseClaims{
		Username: utils.Md5(username, true),
		Password: utils.Md5(password, true),
	})
	return j.CreateToken(claims)
}

// CreateToken 创建 token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
