package Models

type ShiftStaff struct {
	ID        int       `gorm:"column:id_shift_staff;int;primaryKey;autoIncrement" 	json:"id_shift_staff"`
	StaffID   int       `gorm:"column:staff_id;int" 								json:"staff_id"`
	ShiftID   int       `gorm:"column:shift_id;int" 								json:"shift_id"`
	Shift     Shift     `gorm:"foreignKey:ShiftID"`
	DataStaff DataStaff `gorm:"foreignKey:StaffID"`
}

type ShiftStaffInput struct {
	StaffID int `json:"staff_id"`
	ShiftID int `json:"shift_id"`
}

type ShiftStaffResponse struct {
	ID      int `json:"id_shift_staff"`
	StaffID int `json:"staff_id"`
	ShiftID int `json:"shift_id"`
}

type ShiftStaffListResponse struct {
	ShiftStaff []ShiftStaffResponse `json:"ShiftStaff"`
}