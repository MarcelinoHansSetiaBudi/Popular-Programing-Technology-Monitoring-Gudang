package DetailIncomeProductController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDetailIncome(c *gin.Context) {
	var detailincomeproduct []Models.DetailIncomeProduct
	database.DB.Find(&detailincomeproduct)

	var detailincomeproductResponses []Models.DetailIncomeProductResponse
	for _, detailincomeproduct := range detailincomeproduct {
		detailincomeproductResponse := Models.DetailIncomeProductResponse{
			ID:          	detailincomeproduct.ID,
			ProductID:		detailincomeproduct.ProductID,
			ShiftStaffID: 	detailincomeproduct.ShiftStaffID,
			Stock:    		detailincomeproduct.Stock,
		}
		detailincomeproductResponses = append(detailincomeproductResponses, detailincomeproductResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductResponses})
}

func Create(c *gin.Context) {
	var detailincomeproductInput Models.DetailIncomeProductInput
	if err := c.ShouldBindJSON(&detailincomeproductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailincomeproduct := Models.DetailIncomeProduct{
		ProductID: 		detailincomeproductInput.ProductID,
		ShiftStaffID: 	detailincomeproductInput.ShiftStaffID,
		Stock: 			detailincomeproductInput.Stock,
	}

	if err := database.DB.Create(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create detail income product"})
		return
	}

	detailincomeproductResponse := Models.DetailIncomeProductResponse{
		ID: 			detailincomeproduct.ID,	
		ProductID:          	detailincomeproduct.ProductID,
		ShiftStaffID:		detailincomeproduct.ShiftStaffID,
		Stock: 	detailincomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductResponse})
}

func Read(c *gin.Context) {
	var detailincomeproduct Models.DetailIncomeProduct
	if err := database.DB.Where("id_income_product = ?", c.Param("id")).First(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	detailincomeproductResponse := Models.DetailIncomeProductResponse{
		ID:          	detailincomeproduct.ID,
		ProductID:		detailincomeproduct.ProductID,
		ShiftStaffID: 	detailincomeproduct.ShiftStaffID,
		Stock:          detailincomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductResponse})
}

func Update(c *gin.Context) {
	var detailincomeproduct Models.DetailIncomeProduct
	if err := database.DB.Where("id_income_product = ?", c.Param("id")).First(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var detailincomeproductInput Models.DetailIncomeProduct
	if err := c.ShouldBindJSON(&detailincomeproductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailincomeproduct.ProductID =	  	detailincomeproduct.ProductID
	detailincomeproduct.ShiftStaff = 	detailincomeproduct.ShiftStaff
	detailincomeproduct.Stock = 		detailincomeproduct.Stock

	if err := database.DB.Save(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update detail income product"})
		return
	}

	detailincomeproductResponse := Models.DetailIncomeProductResponse{
		ID:          	detailincomeproduct.ID,
		ProductID:		detailincomeproduct.ProductID,
		ShiftStaffID: 	detailincomeproduct.ShiftStaffID,
		Stock:          detailincomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductResponse})
}

func Destroy(c *gin.Context) {
	var detailincomeproduct Models.DetailIncomeProduct
	if err := database.DB.Where("id_income_product = ?", c.Param("id")).First(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete detail income product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}