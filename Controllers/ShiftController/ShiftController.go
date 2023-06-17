package ShiftController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllShift(c *gin.Context) {
	var shift []Models.Shift
	database.DB.Find(&shift)

	var shiftResponses []Models.ShiftResponse
	for _, shift := range shift {
		ShiftResponse := Models.ShiftResponse{
			ID:          	shift.ID,
			ShiftStart:		  	shift.ShiftStart,
			ShiftEnd: 		shift.ShiftEnd,
		}
		shiftResponses = append(shiftResponses, ShiftResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftResponses})
}

func Create(c *gin.Context) {
	var ShiftInput Models.ShiftInput
	if err := c.ShouldBindJSON(&ShiftInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shift := Models.Shift{
		ShiftStart: ShiftInput.ShiftStart,
		ShiftEnd: 	ShiftInput.ShiftEnd,
	}

	if err := database.DB.Create(&shift).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shift"})
		return
	}

	shiftResponse := Models.ShiftResponse{
		ID:			shift.ID,
		ShiftStart: shift.ShiftStart,
		ShiftEnd: 	shift.ShiftEnd,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftResponse})
}

func Read(c *gin.Context) {
	var shift Models.Shift
	if err := database.DB.Where("id_shift = ?", c.Param("id")).First(&shift).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	shiftResponse := Models.ShiftResponse{
		ID: 		shift.ID,
		ShiftStart:	shift.ShiftStart,
		ShiftEnd:  	shift.ShiftEnd,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftResponse})
}

func Update(c *gin.Context) {
	var shift Models.Shift
	if err := database.DB.Where("id_shift = ?", c.Param("id")).First(&shift).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var ShiftInput Models.Shift
	if err := c.ShouldBindJSON(&ShiftInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shift.ShiftStart = 	ShiftInput.ShiftStart
	shift.ShiftEnd = 	ShiftInput.ShiftEnd

	if err := database.DB.Save(&shift).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shift"})
		return
	}

	shiftResponse := Models.ShiftResponse{
		ID: 		shift.ID,
		ShiftStart:	shift.ShiftStart,
		ShiftEnd:	shift.ShiftEnd,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftResponse})
}

func Destroy(c *gin.Context) {
	var shift Models.Shift
	if err := database.DB.Where("id_shift = ?", c.Param("id")).First(&shift).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&shift).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shift"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}