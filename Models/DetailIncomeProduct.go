package Models

type DetailIncomeProduct struct {
	ID            int         `gorm:"column:id_income_product;int;primaryKey;autoIncrement"	json:"id_income_product"`
	ProductID     int         `gorm:"column:product_id;int"									json:"product_id"`
	ShiftStaffID  int         `gorm:"column:shift_staff_id;int"								json:"shift_staff_id"`
	DistributorID int         `gorm:"column:distributor_id;int"								json:"distributor_id"`
	Stock         int         `gorm:"column:stock" json:"stock"`
	Distributor   Distributor `gorm:"foreignKey:DistributorID"`
	ShiftStaff    ShiftStaff  `gorm:"foreignKey:ShiftStaffID"`
	Product       Product     `gorm:foreignKey:ProductID`
}

type DetailIncomeProductInput struct {
	ProductID     int `json:"product_id"`
	ShiftStaffID  int `json:"shift_staff_id"`
	DistributorID int `json:"distributor_id"`
	Stock         int `json:"stock"`
}

type DetailIncomeProductResponse struct {
	ID           int `json:"id_income_product"`
	ProductID    int `json:"product_id"`
	ShiftStaffID int `json:"shift_staff_id"`
	Stock        int `json:"stock"`
}

type DetailIncomeProductListResponse struct {
	DetailIncomeProduct []DetailIncomeProductResponse `json:"DetailIncomeProduct"`
}
