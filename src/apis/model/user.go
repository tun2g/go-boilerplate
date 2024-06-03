package model

type User struct {
	BaseModel
	Email    string `gorm:"type:varchar(200)" json:"email"`
	Password string `gorm:"type:varchar(200)" json:"password"`
	FullName string `gorm:"type:varchar(200)" json:"fullName"`
	Role     string `gorm:"type:varchar(200)" json:"role"`
}
