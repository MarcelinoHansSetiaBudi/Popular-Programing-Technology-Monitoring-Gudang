package Models

type DetailIncomeProduct struct {
	ID            int         `gorm:"column:id_income_product;primaryKey;autoIncrement" json:"id_income_product"`
	ProductID     string      `gorm:"column:product_id" json:"product_id"`
	ShiftStaffID  string      `gorm:"column:shift_staff_id" json:"shift_staff_id"`
	DistributorID string      `gorm:"column:id_distributor" json:"id_distributor"`
	Stock         int         `gorm:"column:stock" json:"stock"`
	Distributor   Distributor `gorm:"foreignKey:DistributorID"`
	ShiftStaff    ShiftStaff  `gorm:"foreignKey:ShiftStaffID"`
}

type DetailIncomeProductInput struct {
	ProductID    string `json:"product_id"`
	ShiftStaffID string `json:"shift_staff_id"`
	Stock        int    `json:"stock"`
}

type DetailIncomeProductResponse struct {
	ID           int    `json:"id_income_product"`
	ProductID    string `json:"product_id"`
	ShiftStaffID string `json:"shift_staff_id"`
	Stock        int    `json:"stock"`
}

type DetailIncomeProductListResponse struct {
	DetailIncomeProduct []DetailIncomeProductResponse `json:"DetailIncomeProduct"`
}
