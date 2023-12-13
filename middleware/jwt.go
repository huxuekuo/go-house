package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-house/consts"
	"github.com/go-house/library"
	"github.com/golang-jwt/jwt/v5"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		// Set example variable
		t, _ := jwt.ParseWithClaims(tokenStr, &library.PassInfo{}, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return []byte(consts.AUTH_SECRETKEY), nil
		})
		if claims, ok := t.Claims.(*library.PassInfo); ok && t.Valid {
			c.Set("passInfo", claims)
		}
		// before request
		c.Next()
	}
}
