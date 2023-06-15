// package main

// import (
// 	database "FinPro/Database"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"

// 	"github.com/gorilla/mux"

// 	controllers "FinPro/Controllers/DistributorController"
// 	"FinPro/Models"
// )

// func main(){
//  	database.ConnectDB()
//  	suhu := mux.NewRouter()
// // 	suhu.HandleFunc("/api", Controllers.).Methods("POST")
// // 	suhu.HandleFunc().Methods("GET")
// // }
// // Inisialisasi database
// dsn := "user:password@tcp(localhost:3306)/database"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// if err != nil {
// 	log.Fatal(err)
// }

// // Migrasi model
// if err := db.AutoMigrate(&Models.Distributor{}); err != nil {
// 	log.Fatal(err)
// }

// // Inisialisasi controller
// distributorController := &controllers.DistributorController{
// 	DB: db,
// }

// // Inisialisasi Echo
// e := echo.New()

// // Mendaftarkan rute-rute
// distributorController.RegisterRoutes(e)

// // Menjalankan server
// e.Start(":8080")
// }

package main

func main() {
	
}
