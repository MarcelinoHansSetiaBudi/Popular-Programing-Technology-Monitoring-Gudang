package Models

// Buat DB
type Product struct {
	ID    int    `gorm:"column:id_product;primaryKey;autoIncrement" json:"id_product"`
	Name  string `gorm:"column:name;type:varchar(255)" json:"name"`
	Stock string `gorm:"column:stock;type:varchar(255)" json:"stock"`
	Brand string `gorm:"column:brand;type:varchar(255)" json:"brand"`
}

type ProductInput struct {
	Name  string `json:"name"`
	Stock string `json:"stock"`
	Brand string `json:"brand"`
}

type ProductResponse struct {
	ID    int    `json:"id_product"`
	Name  string `json:"name"`
	Stock string `json:"stock"`
	Brand string `json:"brand"`
}

type ProductListResponse struct {
	Product []ProductResponse `json:"Product"`
}
