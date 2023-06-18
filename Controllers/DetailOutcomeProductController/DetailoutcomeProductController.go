package DetailOutcomeProductController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDetailOutcome(c *gin.Context) {
	var detailoutcomeproduct []Models.DetailOutcomeProduct
	database.DB.Find(&detailoutcomeproduct)

	var detailoutcomeproductResponses []Models.DetailOutcomeProduct
	for _, detailoutcomeproduct := range detailoutcomeproduct {
		detailoutcomeproductResponse := Models.DetailOutcomeProduct{
			ID:          	detailoutcomeproduct.ID,
			ProductID:		detailoutcomeproduct.ProductID,
			ShiftStaffID:   detailoutcomeproduct.ShiftStaffID,
			Stock: 			detailoutcomeproduct.Stock,
		}
		detailoutcomeproductResponses = append(detailoutcomeproductResponses, detailoutcomeproductResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": detailoutcomeproductResponses})
}

func Create(c *gin.Context) {
	var detailoutcomeproductInput Models.DetailOutcomeProductInput
	if err := c.ShouldBindJSON(&detailoutcomeproductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailoutcomeproduct := Models.DetailOutcomeProduct{
		ProductID: 		detailoutcomeproductInput.ProductID,
		ShiftStaffID:	detailoutcomeproductInput.ShiftStaffID,
		Stock: 			detailoutcomeproductInput.Stock,
	}

	if err := database.DB.Create(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create outcome product"})
		return
	}

	detailoutcomeproductsResponse := Models.DetailOutcomeProductResponse{
		ID:          	detailoutcomeproduct.ID,
		ProductID:		detailoutcomeproduct.ProductID,
		ShiftStaffID:   detailoutcomeproduct.ShiftStaffID,
		Stock: 			detailoutcomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailoutcomeproductsResponse})
}

func Read(c *gin.Context) {
	var detailoutcomeproduct Models.DetailOutcomeProduct
	if err := database.DB.Where("id_outcome_product = ?", c.Param("id")).First(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	detailoutcomeproductResponse := Models.DetailOutcomeProductResponse{
		ID:          	detailoutcomeproduct.ID,
		ProductID:		detailoutcomeproduct.ProductID,
		ShiftStaffID:   detailoutcomeproduct.ShiftStaffID,
		Stock: 			detailoutcomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailoutcomeproductResponse})
}

func Update(c *gin.Context) {
	var detailoutcomeproduct Models.DetailOutcomeProduct
	if err := database.DB.Where("id_outcome_product = ?", c.Param("id")).First(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var detailoutcomeproductInput Models.DetailOutcomeProduct
	if err := c.ShouldBindJSON(&detailoutcomeproductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailoutcomeproduct.ProductID =	detailoutcomeproductInput.ProductID
	detailoutcomeproduct.ShiftStaffID = detailoutcomeproductInput.ShiftStaffID
	detailoutcomeproduct.Stock = 		detailoutcomeproductInput.Stock

	if err := database.DB.Save(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update outcome product"})
		return
	}

	detailoutcomeproductResponse := Models.DetailOutcomeProductResponse{
		ID:          	detailoutcomeproduct.ID,
		ProductID:		detailoutcomeproduct.ProductID,
		ShiftStaffID:   detailoutcomeproduct.ShiftStaffID,
		Stock: 			detailoutcomeproduct.Stock,
	}


	c.JSON(http.StatusOK, gin.H{"data": detailoutcomeproductResponse})
}

func Destroy(c *gin.Context) {
	var detailoutcomeproduct Models.DetailOutcomeProduct
	if err := database.DB.Where("id_outcome_product = ?", c.Param("id")).First(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&detailoutcomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete outcome product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}