package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"hospital-admin-api/config"
	"time"
)

type Token struct {
	UserId               int64 `json:"user_id"` // 用户id
	RoleId               int64 `json:"role_id"` // 角色id
	DeptId               int64 `json:"dept_id"` // 部门id
	PostId               int64 `json:"post_id"` // 岗位id
	jwt.RegisteredClaims       // v5版本新加的方法
}

// NewJWT GenerateJWT 生成JWT
func (t Token) NewJWT() (string, error) {
	ttl := time.Duration(config.C.Jwt.Ttl)

	t.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(ttl * time.Hour)) // 过期时间24小时
	t.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())                       // 签发时间
	t.RegisteredClaims.NotBefore = jwt.NewNumericDate(time.Now())                      // 生效时间

	// 使用HS256签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	s, err := token.SignedString([]byte(config.C.Jwt.SignKey))

	return s, err
}

// ParseJwt 解析JWT
func ParseJwt(authorization string) (*Token, error) {
	t, err := jwt.ParseWithClaims(authorization, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.Jwt.SignKey), nil
	})

	if claims, ok := t.Claims.(*Token); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
