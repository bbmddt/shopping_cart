package middleware

import (
	"net/http"
	jwtHelper "shopping_cart/utils/jwt"

	"github.com/gin-gonic/gin"
)

// 管理員授權
func AuthAdminMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil && decodedClaims.IsAdmin {
				c.Next()
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "你沒有權限訪問!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你沒有授權!"})
		}
		c.Abort()

	}
}

// 使用者授權
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "你沒有權限訪問!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "你沒授權!"})
		}
		c.Abort()

	}
}
