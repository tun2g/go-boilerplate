package model

import (
	
	"fist-app/src/shared/model"

	"github.com/jinzhu/gorm"
	"github.com/google/uuid"

)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string `gorm:"type:varchar(200);UNIQUE"`
	Password string `gorm:"type:varchar(200);"`
	FullName string `gorm:"type:varchar(200);"`
	model.Auditable
}
