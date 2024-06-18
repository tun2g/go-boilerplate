package model

type Post struct {
	BaseModel
	Title       string `gorm:"type:varchar(200)" json:"title"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	UserId      string `gorm:"type:varchar(36)" json:"userId"`
}

type Comment struct {
	BaseModel
	Text   string `gorm:"type:varchar(200)" json:"text"`
	PostId string `gorm:"type:varchar(200)" json:"postId"`
	UserId string `gorm:"type:varchar(36)" json:"userId"`
}
