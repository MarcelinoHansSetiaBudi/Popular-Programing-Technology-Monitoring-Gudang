package Models

type DetailOutcomeProduct struct {
	ID           int        `gorm:"column:id_income_product;primaryKey;autoIncrement" json:"id_income_product"`
	ProductID    string     `gorm:"column:product_id" json:"product_id"`
	ShiftStaffID string     `gorm:"column:shift_staff_id" json:"shift_staff_id"`
	Stock        int        `gorm:"column:stock" json:"stock"`
	Product      Product    `gorm:"foreignKey:ProductID"`
	ShiftStaff   ShiftStaff `gorm:"foreignKey:ShiftStaffID"`
}

type DetailOutcomeProductInput struct {
	ProductID    string `json:"product_id"`
	ShiftStaffID string `json:"shift_staff_id"`
	Stock        int    `json:"stock"`
}

type DetailOutcomeProductResponse struct {
	ID           int    `json:"id_income_product"`
	ProductID    string `json:"product_id"`
	ShiftStaffID string `json:"shift_staff_id"`
	Stock        int    `json:"stock"`
}

type DetailOutcomeProductListResponse struct {
	DetailOutcomeProduct []DetailOutcomeProductResponse `json:"DetailOutcomeProduct"`
}
