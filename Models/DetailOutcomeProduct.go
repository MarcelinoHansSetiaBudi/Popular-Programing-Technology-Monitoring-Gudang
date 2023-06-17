package Models

type DetailOutcomeProduct struct {
	ID           int        `gorm:"column:id_income_product;int;primaryKey;autoIncrement"	json:"id_income_product"`
	ProductID    int        `gorm:"column:product_id;int" 									json:"product_id"`
	ShiftStaffID int        `gorm:"column:shift_staff_id;int"								json:"shift_staff_id"`
	Stock        int        `gorm:"column:stock"											json:"stock"`
	Product      Product    `gorm:"foreignKey:ProductID"`
	ShiftStaff   ShiftStaff `gorm:"foreignKey:ShiftStaffID"`
}

type DetailOutcomeProductInput struct {
	ProductID    int `json:"product_id"`
	ShiftStaffID int `json:"shift_staff_id"`
	Stock        int `json:"stock"`
}

type DetailOutcomeProductResponse struct {
	ID           int `json:"id_income_product"`
	ProductID    int `json:"product_id"`
	ShiftStaffID int `json:"shift_staff_id"`
	Stock        int `json:"stock"`
}

type DetailOutcomeProductListResponse struct {
	DetailOutcomeProduct []DetailOutcomeProductResponse `json:"DetailOutcomeProduct"`
}
