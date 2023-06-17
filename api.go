package main

import (
	distributorController "FinPro/Controllers/DistributorController"
	productController "FinPro/Controllers/ProductController"
	database "FinPro/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	// Distributor routes
	router.GET("api/gudang/distributor", distributorController.GetAllDistributor)
	router.POST("api/gudang/distributor", distributorController.Create)
	router.GET("api/gudang/distributor/:id", distributorController.Read)
	router.PUT("api/gudang/distributor/:id", distributorController.Update)
	router.DELETE("api/gudang/ditributor/:id", distributorController.Destroy)

	// Product routes
	router.GET("api/gudang/product", productController.GetAllProduct)
	router.POST("api/gudang/product", productController.Create)
	router.GET("api/gudang/product/:id", productController.Read)
	router.PUT("api/gudang/product/:id", productController.Update)
	router.DELETE("api/gudang/product/:id", productController.Destroy)

	// Route Prefix Address
	router.Run("localhost:8080")
}
