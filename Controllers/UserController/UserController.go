package UserController

import (
	"net/http"

	database "FinPro/Database"
	"FinPro/Models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var user Models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat data pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
