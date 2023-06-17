package database

import (
	"FinPro/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/PPT-Monitoring_gudang"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Models.DataStaff{},
		&Models.Shift{},
		&Models.Distributor{},
		&Models.Product{},
		&Models.ShiftStaff{},
		&Models.DetailIncomeProduct{},
		&Models.DetailOutcomeProduct{},
	)
	DB = db
}