package DistributorController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDistributor(c *gin.Context) {
	var distributor []Models.Distributor
	database.DB.Find(&distributor)

	var distributorResponses []Models.DistributorResponse
	for _, distributor := range distributor {
		distributorResponse := Models.DistributorResponse{
			ID:          	distributor.ID,
			Name:		  	distributor.Name,
			Address: 		distributor.Address,
			PhoneNumber:    distributor.PhoneNumber,
		}
		distributorResponses = append(distributorResponses, distributorResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": distributorResponses})
}

func Create(c *gin.Context) {
	var distributorInput Models.DistributorInput
	if err := c.ShouldBindJSON(&distributorInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	distributor := Models.Distributor{
		Name: distributorInput.Name,
		Address: distributorInput.Address,
		PhoneNumber: distributorInput.PhoneNumber,
	}

	if err := database.DB.Create(&distributor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create distributor"})
		return
	}

	distributorResponse := Models.DistributorResponse{
		Name:          distributor.Name,
		Address:		  distributor.Address,
		PhoneNumber: distributor.PhoneNumber,
	}

	c.JSON(http.StatusOK, gin.H{"data": distributorResponse})
}

func Read(c *gin.Context) {
	var distributor Models.Distributor
	if err := database.DB.Where("id_distributor = ?", c.Param("id")).First(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	distributorResponse := Models.DistributorResponse{
		ID:          distributor.ID,
		Name:		  distributor.Name,
		Address: distributor.Address,
		PhoneNumber:            distributor.PhoneNumber,
	}

	c.JSON(http.StatusOK, gin.H{"data": distributorResponse})
}

func Update(c *gin.Context) {
	var distributor Models.Distributor
	if err := database.DB.Where("id_distributor = ?", c.Param("id")).First(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var distributorInput Models.Distributor
	if err := c.ShouldBindJSON(&distributorInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	distributor.Name =	  distributorInput.Name
	distributor.Address = distributorInput.Address
	distributor.PhoneNumber = distributorInput.PhoneNumber

	if err := database.DB.Save(&distributor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update distributor"})
		return
	}

	distributorResponse := Models.DistributorResponse{
		Name:          distributor.Name,
		Address:		  distributor.Address,
		PhoneNumber: distributor.PhoneNumber,
	}

	c.JSON(http.StatusOK, gin.H{"data": distributorResponse})
}

func Destroy(c *gin.Context) {
	var distributor Models.Distributor
	if err := database.DB.Where("distributor_id = ?", c.Param("id")).First(&distributor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&distributor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete distributor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}