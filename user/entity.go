package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone" gorm:"type:varchar(13);primaryKey"`
	Password string `json:"password"`
}
