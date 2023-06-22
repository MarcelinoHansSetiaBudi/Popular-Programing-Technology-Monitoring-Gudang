package main

import (
	authController "FinPro/Controllers/AuthController"
	datastaffController "FinPro/Controllers/DataStaffController"
	detailincomeProductController "FinPro/Controllers/DetailIncomeProductController"
	detailoutcomeProductController "FinPro/Controllers/DetailOutcomeProductController"
	distributorController "FinPro/Controllers/DistributorController"
	productController "FinPro/Controllers/ProductController"
	shiftController "FinPro/Controllers/ShiftController"
	shiftstaffController "FinPro/Controllers/ShiftStaffController"
	userController "FinPro/Controllers/UserController"
	database "FinPro/Database"
	middleware "FinPro/Middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDB()

	// Users Routes
	router.POST("/api/users/login", authController.LoginHandler)
	router.POST("/api/users", userController.UserCreate)

	authMiddleware := middleware.RequireAuth

	auth := router.Group("/api")
	auth.Use(authMiddleware)
	{
		// User 
		auth.POST("users/logout", authController.LogoutHandler)

		// Distributor routes
		auth.GET("gudang/distributor", distributorController.GetAllDistributor)
		auth.POST("gudang/distributor", distributorController.Create)
		auth.GET("gudang/distributor/:id", distributorController.Read)
		auth.PUT("gudang/distributor/:id", distributorController.Update)
		auth.DELETE("gudang/distributor/:id", distributorController.Destroy)

		// Product routes
		auth.GET("gudang/product", productController.GetAllProduct)
		auth.POST("gudang/product", productController.Create)
		auth.GET("gudang/product/:id", productController.Read)
		auth.PUT("gudang/product/:id", productController.Update)
		auth.DELETE("gudang/product/:id", productController.Destroy)

		// Shift routes
		auth.GET("gudang/shift", shiftController.GetAllShift)
		auth.POST("gudang/shift", shiftController.Create)
		auth.GET("gudang/shift/:id", shiftController.Read)
		auth.PUT("gudang/shift/:id", shiftController.Update)
		auth.DELETE("gudang/shift/:id", shiftController.Destroy)

		// Data Staff routes
		auth.GET("gudang/data-staff", datastaffController.GetAllDataStaff)
		auth.POST("gudang/data-staff", datastaffController.Create)
		auth.GET("gudang/data-staff/:id", datastaffController.Read)
		auth.PUT("gudang/data-staff/:id", datastaffController.Update)
		auth.DELETE("gudang/data-staff/:id", datastaffController.Destroy)

		// Shift Staff routes
		auth.GET("gudang/shift-staff", shiftstaffController.GetAllShiftStaff)
		auth.POST("gudang/shift-staff", shiftstaffController.Create)
		auth.GET("gudang/shift-staff/:id", shiftstaffController.Read)
		auth.PUT("gudang/shift-staff/:id", shiftstaffController.Update)
		auth.DELETE("gudang/shift-staff/:id", shiftstaffController.Destroy)
		
		// Detail Income Product routes
		auth.GET("gudang/detail-income", detailincomeProductController.GetAllDetailIncome)
		auth.POST("gudang/detail-income", detailincomeProductController.Create)
		auth.GET("gudang/detail-income/:id", detailincomeProductController.Read)
		auth.PUT("gudang/detail-income/:id", detailincomeProductController.Update)
		auth.DELETE("gudang/detail-income/:id", detailincomeProductController.Destroy)

		// Detail Outcome Product routes
		auth.GET("gudang/detail-outcome", detailoutcomeProductController.GetAllDetailOutcome)
		auth.POST("gudang/detail-outcome", detailoutcomeProductController.Create)
		auth.GET("gudang/detail-outcome/:id", detailoutcomeProductController.Read)
		auth.PUT("gudang/detail-outcome/:id", detailoutcomeProductController.Update)
		auth.DELETE("gudang/detail-outcome/:id", detailoutcomeProductController.Destroy)
	}
	
	
	// Route Prefix Address
	router.Run("localhost:8080")
}
