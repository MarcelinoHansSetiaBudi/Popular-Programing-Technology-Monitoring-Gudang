package main

import (
	datastaffController "FinPro/Controllers/DataStaffController"
	distributorController "FinPro/Controllers/DistributorController"
	productController "FinPro/Controllers/ProductController"
	shiftController "FinPro/Controllers/ShiftController"
	shiftstaffController "FinPro/Controllers/ShiftStaffController"
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

	// Shift routes
	router.GET("api/gudang/shift", shiftController.GetAllShift)
	router.POST("api/gudang/shift", shiftController.Create)
	router.GET("api/gudang/shift/:id", shiftController.Read)
	router.PUT("api/gudang/shift/:id", shiftController.Update)
	router.DELETE("api/gudang/shift/:id", shiftController.Destroy)

	// Data Staff routes
	router.GET("api/gudang/data-staff", datastaffController.GetAllDataStaff)
	router.POST("api/gudang/data-staff", datastaffController.Create)
	router.GET("api/gudang/data-staff/:id", datastaffController.Read)
	router.PUT("api/gudang/data-staff/:id", datastaffController.Update)
	router.DELETE("api/gudang/data-staff/:id", datastaffController.Destroy)

	// Shift Staff routes
	router.GET("api/gudang/shift-staff", shiftstaffController.SelectAllShiftStaff)
	router.POST("api/gudang/shift-staff", shiftstaffController.Create)
	router.GET("api/gudang/shift-staff/:id", shiftstaffController.Read)
	router.PUT("api/gudang/shift-staff/:id", shiftstaffController.Update)
	router.DELETE("api/gudang/shift-staff/:id", shiftstaffController.Destroy)
	// Route Prefix Address
	router.Run("localhost:8080")
}
