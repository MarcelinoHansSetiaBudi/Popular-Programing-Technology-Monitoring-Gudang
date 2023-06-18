package shiftstaffcontroller

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllShiftStaff(c *gin.Context) {
	var shiftstaff []Models.ShiftStaff
	database.DB.Find(&shiftstaff)

	var shiftstaffRepsonses []Models.ShiftStaffResponse
	for _, shiftstaff := range shiftstaff {
		shiftstaffResponse := Models.ShiftStaffResponse{
			ID:         shiftstaff.ID,
			StaffID:	shiftstaff.StaffID,
			ShiftID: 	shiftstaff.ShiftID,
		}
		shiftstaffRepsonses = append(shiftstaffRepsonses, shiftstaffResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftstaffRepsonses})
}

func Create(c *gin.Context) {
	var shiftstaffInput Models.ShiftStaffInput
	if err := c.ShouldBindJSON(&shiftstaffInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shiftstaff := Models.ShiftStaff{
		StaffID:	shiftstaffInput.StaffID,
		ShiftID:	shiftstaffInput.ShiftID,
	}

	if err := database.DB.Create(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shift staff"})
		return
	}

	shiftstaffResponse := Models.ShiftStaffResponse{
		ID:			shiftstaff.ID,
		StaffID:	shiftstaff.StaffID,
		ShiftID: 	shiftstaff.ShiftID,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftstaffResponse})
}

func Read(c *gin.Context) {
	var shiftstaff Models.ShiftStaff
	if err := database.DB.Where("id_shift_staff = ?", c.Param("id")).First(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	shiftstaffResponse := Models.ShiftStaffResponse{
		ID: 		shiftstaff.ID,
		StaffID: 	shiftstaff.StaffID,
		ShiftID: 	shiftstaff.ShiftID,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftstaffResponse})
}

func Update(c *gin.Context) {
	var shiftstaff Models.ShiftStaff
	if err := database.DB.Where("id_shift_staff = ?", c.Param("id")).First(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var shiftstaffInput Models.ShiftStaffInput
	if err := c.ShouldBindJSON(&shiftstaffInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shiftstaff.StaffID =	  shiftstaffInput.StaffID
	shiftstaff.ShiftID =	  shiftstaffInput.ShiftID

	if err := database.DB.Save(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shift staff"})
		return
	}

	shiftstaffResponse := Models.ShiftStaffResponse{
		ID: 		shiftstaff.ID,
		StaffID: 	shiftstaff.StaffID,
		ShiftID: 	shiftstaff.ShiftID,
	}

	c.JSON(http.StatusOK, gin.H{"data": shiftstaffResponse})
}

func Destroy(c *gin.Context) {
	var shiftstaff Models.ShiftStaff
	if err := database.DB.Where("id_shift_staff = ?", c.Param("id")).First(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&shiftstaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shift staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
