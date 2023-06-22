package Middleware

import (
	database "FinPro/Database"
	"FinPro/Models"
	Token "FinPro/Utils/Token"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	tokenString := authHeaderParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRETEJWTKEYS123321RAHASIAhehe"), nil
	})

	if err != nil || !token.Valid || Token.TokenIsBlacklisted(tokenString) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime := int64(claims["exp"].(float64))
		currentTime := time.Now().Unix()
		if currentTime > expirationTime {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, Token Expired"})
			return
		}

		var user Models.User
		database.DB.Where("username = ?", claims["sub"]).First(&user)
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("user", user)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}