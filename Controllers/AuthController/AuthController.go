package AuthController

import (
	"net/http"
	"strings"
	"time"

	database "FinPro/Database"
	"FinPro/Models"
	"FinPro/Utils/Token"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
    var form Models.LoginForm

    if err := c.ShouldBindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Permintaan tidak valid"})
        return
    }

    var user Models.User
    if err := database.DB.Where("username = ?", form.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Data pengguna tidak ditemukan"})
        return
    }

    if user.Password != form.Password {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Username atau Kata Sandi tidak valid"})
        return
    }

    existingToken := c.GetHeader("Authorization")
    if existingToken != "" {
        existingToken = strings.Replace(existingToken, "Bearer ", "", 1)
        Token.BlacklistToken(existingToken)
    }

    expirationTime := time.Now().Add(60 * time.Minute)
    tokenString, err := Token.GenerateTokenString(form, "SECRETEJWTKEYS123321RAHASIAhehe", expirationTime)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghasilkan token"})
        return
    }

    expTimeFormatted := expirationTime.Format("2006-01-02 15:04:05")
    c.JSON(http.StatusOK, gin.H{"token": "Bearer "+tokenString, "exp_token": expTimeFormatted})
}

func LogoutHandler(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }

    tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

    Token.BlacklistToken(tokenString)

    c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}