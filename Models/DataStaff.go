package Models

// ORM Object Relational Mapping

// Buat DB
type DataStaff struct {
	ID      int    `gorm:"column:id_data_staff;int;primaryKey;autoIncrement"	json:"id_data_staff"`
	Name    string `gorm:"column:name;type:varchar(255);unique"					json:"name"`
	Address string `gorm:"column:address;type:varchar(255)"						json:"address"`
	Phone   string `gorm:"column:phone;type:varchar(255)"						json:"phone"`
	Email   string `gorm:"column:email;type:varchar(255);unique"				json:"email"`
}

type DataStaffInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type DataStaffResponse struct {
	ID      int    `json:"id_data_staff"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type DataStaffListResponse struct {
	DataStaff []DataStaffResponse `json:"Data_Staff"`
}