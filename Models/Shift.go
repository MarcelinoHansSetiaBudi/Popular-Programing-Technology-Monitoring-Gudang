package Models

// Buat DB
type Shift struct {
	ID          int    `gorm:"column:id_shift;primaryKey;autoIncrement" json:"id_shift"`
	ShiftStart  string `gorm:"column:shift_start;type:varchar(255)" json:"shift_start"`
	ShiftEnd    string `gorm:"column:shift_end;type:varchar(255)" json:"shift_end"`
}

type ShiftInput struct {
	ShiftStart string `json:"shift_start"`
	ShiftEnd   string `json:"shift_end"`
}

type ShiftResponse struct {
	ID          int    `json:"id_shift"`
	ShiftStart  string `json:"shift_start"`
	ShiftEnd    string `json:"shift_end"`
}

type ShiftListResponse struct {
	Shift []ShiftResponse `json:"Shift"`
}