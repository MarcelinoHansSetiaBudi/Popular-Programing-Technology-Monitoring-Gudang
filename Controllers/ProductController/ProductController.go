package ProductController

import (
	database "FinPro/Database"
	"FinPro/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var product []Models.Product
	database.DB.Find(&product)

	var productResponses []Models.ProductResponse
	for _, product := range product {
		productResponse := Models.ProductResponse{
			ID:          	product.ID,
			Name:		  	product.Name,
			Stock: 		product.Stock,
			Brand:    product.Brand,
		}
		productResponses = append(productResponses, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": productResponses})
}

func Create(c *gin.Context) {
	var ProductInput Models.ProductInput
	if err := c.ShouldBindJSON(&ProductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := Models.Product{
		Name: ProductInput.Name,
		Stock: ProductInput.Stock,
		Brand: ProductInput.Brand,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	productResponse := Models.ProductResponse{
		ID: 	product.ID,
		Name:   product.Name,
		Stock:	product.Stock,
		Brand: 	product.Brand,
	}

	c.JSON(http.StatusOK, gin.H{"data": productResponse})
}

func Read(c *gin.Context) {
	var product Models.Product
	if err := database.DB.Where("id_product = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	productResponse := Models.ProductResponse{
		ID:		product.ID, 
		Name:   product.Name,
		Stock:  product.Stock,
		Brand: 	product.Brand,
	}

	c.JSON(http.StatusOK, gin.H{"data": productResponse})
}

func Update(c *gin.Context) {
	var product Models.Product
	if err := database.DB.Where("id_product = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	var ProductInput Models.Product
	if err := c.ShouldBindJSON(&ProductInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name =	  ProductInput.Name
	product.Stock = ProductInput.Stock
	product.Brand = ProductInput.Brand

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	productResponse := Models.ProductResponse{
		ID: 			product.ID,		
		Name:         	product.Name,
		Stock:		  	product.Stock,
		Brand: 			product.Brand,
	}

	c.JSON(http.StatusOK, gin.H{"data": productResponse})
}

func Destroy(c *gin.Context) {
	var product Models.Product
	if err := database.DB.Where("id_product = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO DATA!"})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}