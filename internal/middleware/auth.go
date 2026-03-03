package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   string   `json:"sub"`
	TenantID string   `json:"tenant_id"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" { c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth header"}); c.Abort(); return }
		token := strings.TrimPrefix(auth, "Bearer ")
		claims := &Claims{}
		t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
		if err != nil || !t.Valid { c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"}); c.Abort(); return }
		c.Set("user_id", claims.UserID); c.Set("tenant_id", claims.TenantID); c.Next()
	}
}
