package library

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-house/consts"
	"github.com/golang-jwt/jwt/v5"
)

func GetPassinfo(ctx *gin.Context) *PassInfo {
	var passInfo PassInfo
	if info, ok := ctx.Get("passinfo"); ok {
		if passInfo, ok = info.(PassInfo); ok {
			return &passInfo
		}
	}
	return nil
}

func GenerateJWT(uid uint64) (string, error) {
	claims := PassInfo{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(consts.AUTH_SECRETKEY))

	return s, err
}
