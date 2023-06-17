package DataStaffController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDataStaff(c *gin.Context) {
	var datastaff []Models.DataStaff
	database.DB.Find(&datastaff)

	var datastaffResponses []Models.DataStaffResponse
	for _, datastaff := range datastaff {
		datastaffResponse := Models.DataStaffResponse{
			ID:      	datastaff.ID,
			Name:		datastaff.Name,
			Address: 	datastaff.Address,
			Phone:		datastaff.Phone,
			Email: 		datastaff.Email,
		}
		datastaffResponses = append(datastaffResponses, datastaffResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": datastaffResponses})
}

func Create(c *gin.Context) {
	var datastaffInput Models.DataStaffInput
	if err := c.ShouldBindJSON(&datastaffInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datastaff := Models.DataStaff{
		Name: 		datastaffInput.Name,
		Address: 	datastaffInput.Address,
		Phone: 		datastaffInput.Phone,
		Email: 		datastaffInput.Email,
	}

	if err := database.DB.Create(&datastaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create data staff"})
		return
	}

	datastaffResponse := Models.DataStaffResponse{
		ID: 		datastaff.ID,
		Name:   	datastaff.Name,
		Address:	datastaff.Address,
		Phone: 		datastaff.Phone,
		Email: 		datastaff.Email,

	}

	c.JSON(http.StatusOK, gin.H{"data": datastaffResponse})
}

func Read(c *gin.Context) {
	var datastaff Models.DataStaff
	if err := database.DB.Where("id_data_staff = ?", c.Param("id")).First(&datastaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	datastaffResponse := Models.DataStaffResponse{
		ID:         datastaff.ID,
		Name:		datastaff.Name,
		Address: 	datastaff.Address,
		Phone:      datastaff.Phone,
		Email: 		datastaff.Email,
	}

	c.JSON(http.StatusOK, gin.H{"data": datastaffResponse})
}

func Update(c *gin.Context) {
	var datastaff Models.DataStaff
	if err := database.DB.Where("id_data_staff = ?", c.Param("id")).First(&datastaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var datastaffInput Models.DataStaff
	if err := c.ShouldBindJSON(&datastaffInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	datastaff.Name =	datastaffInput.Name
	datastaff.Address =	datastaffInput.Address
	datastaff.Phone =	datastaffInput.Phone
	datastaff.Email =	datastaffInput.Email

	if err := database.DB.Save(&datastaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data staff"})
		return
	}

	datastaffResponse := Models.DataStaffResponse{
		ID: 		datastaff.ID,	
		Name:       datastaff.Name,
		Address:	datastaff.Address,
		Phone: 		datastaff.Phone,
		Email:		datastaff.Email, 
	}

	c.JSON(http.StatusOK, gin.H{"data": datastaffResponse})
}

func Destroy(c *gin.Context) {
	var datastaff Models.DataStaff
	if err := database.DB.Where("id_data_staff = ?", c.Param("id")).First(&datastaff).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&datastaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}