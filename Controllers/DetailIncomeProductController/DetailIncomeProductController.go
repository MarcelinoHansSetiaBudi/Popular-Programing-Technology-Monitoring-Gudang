package DetailIncomeProductController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDetailIncome(c *gin.Context) {
	var detailincomeproducts []Models.DetailIncomeProduct
	database.DB.Find(&detailincomeproducts)

	var detailincomeproductsResponses []Models.DetailIncomeProductResponse
	for _, detailincomeproducts := range detailincomeproducts {
		detailincomeproductsResponse := Models.DetailIncomeProductResponse{
			ID:          	detailincomeproducts.ID,
			ProductID:		detailincomeproducts.ProductID,
			ShiftStaffID: 	detailincomeproducts.ShiftStaffID,
			Stock:    		detailincomeproducts.Stock,
		}
		detailincomeproductsResponses = append(detailincomeproductsResponses, detailincomeproductsResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductsResponses})
}

func Create(c *gin.Context) {
	var detailincomeproductsInput Models.DetailIncomeProductInput
	if err := c.ShouldBindJSON(&detailincomeproductsInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detailincomeproducts := Models.DetailIncomeProduct{
		ProductID: 		detailincomeproductsInput.ProductID,
		ShiftStaffID: 	detailincomeproductsInput.ShiftStaffID,
		DistributorID: 	detailincomeproductsInput.DistributorID,
		Stock: 			detailincomeproductsInput.Stock,
	}

	if err := database.DB.Create(&detailincomeproducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create detail income product"})
		return
	}

	detailincomeproductsResponse := Models.DetailIncomeProductResponse{
		ID: 			detailincomeproducts.ID,	
		ProductID:      detailincomeproducts.ProductID,
		ShiftStaffID:	detailincomeproducts.ShiftStaffID,
		Stock: 			detailincomeproducts.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductsResponse})
}

func Read(c *gin.Context) {
	var detailincomeproducts Models.DetailIncomeProduct
	if err := database.DB.Where("id_income_product = ?", c.Param("id")).First(&detailincomeproducts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	detailincomeproductsResponse := Models.DetailIncomeProductResponse{
		ID:          	detailincomeproducts.ID,
		ProductID:		detailincomeproducts.ProductID,
		ShiftStaffID: 	detailincomeproducts.ShiftStaffID,
		Stock:          detailincomeproducts.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductsResponse})
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

	detailincomeproduct.ProductID = detailincomeproductInput.ProductID
	detailincomeproduct.ShiftStaffID = detailincomeproductInput.ShiftStaffID
	detailincomeproduct.Stock = detailincomeproductInput.Stock

	if err := database.DB.Save(&detailincomeproduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update detail income product"})
		return
	}

	detailincomeproductResponse := Models.DetailIncomeProductResponse{
		ID:           detailincomeproduct.ID,
		ProductID:    detailincomeproduct.ProductID,
		ShiftStaffID: detailincomeproduct.ShiftStaffID,
		Stock:        detailincomeproduct.Stock,
	}

	c.JSON(http.StatusOK, gin.H{"data": detailincomeproductResponse})
}

func Destroy(c *gin.Context) {
	var detailincomeproducts Models.DetailIncomeProduct
	if err := database.DB.Where("id_income_product = ?", c.Param("id")).First(&detailincomeproducts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&detailincomeproducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete detail income product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}