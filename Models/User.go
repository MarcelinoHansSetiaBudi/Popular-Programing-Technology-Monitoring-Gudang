package Models

type User struct {
	ID          int    `gorm:"column:user_id;type:int;primaryKey;autoIncrement" 	json:"user_id"`
	Username    string `gorm:"column:username;type:varchar(255);not null;unique" 	json:"username"`
	Password    string `gorm:"column:password;type:varchar(255);not null;" 		json:"password"`
	Name        string `gorm:"column:name;type:varchar(255);"						json:"name"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255);"				json:"phone_number"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}