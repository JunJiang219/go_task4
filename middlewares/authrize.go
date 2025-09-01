package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 不需要鉴权的路径
var skipAuthPaths = []string{
	"/",
	"/user/register",
	"/user/login",
}

var jWTAuthSecretKey = []byte("task4_secret")

type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware(c *gin.Context) {
	for _, path := range skipAuthPaths {
		if path == c.FullPath() {
			c.Next()
			return
		}
	}

	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Request without authorization"})
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jWTAuthSecretKey, nil // 在这里你应该使用加密的密钥
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token parse error"})
		return
	}

	var claims *MyCustomClaims = nil
	var ok bool = false
	if claims, ok = token.Claims.(*MyCustomClaims); !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
		return
	}

	if claims.ExpiresAt.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	c.Set("user_id", claims.UserID)
	c.Next()
}

func NewJWTAuth(userID uint) (tokenStr string, err error) {
	claims := MyCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err = token.SignedString(jWTAuthSecretKey)
	return
}
