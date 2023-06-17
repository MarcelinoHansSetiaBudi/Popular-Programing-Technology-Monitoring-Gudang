package Models

type Distributor struct {
	ID          int    `gorm:"column:id_distributor;int;primaryKey;autoIncrement"	json:"id_distributor"`
	Name        string `gorm:"column:name;type:varchar(255)" 						json:"name"`
	Address     string `gorm:"column:address;type:varchar(255)" 					json:"address"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255)" 				json:"phone_number"`
}

type DistributorInput struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type DistributorResponse struct {
	ID          int    `json:"id_distributor"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type DistributorListResponse struct {
	Distributor []DistributorResponse `json:"Distributor"`
}
