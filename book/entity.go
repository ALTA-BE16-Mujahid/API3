package book

import (
	"github.com/mujahxd/api3-jwt/user"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
	UserID    int    `json:"user_id"`
	User      user.User
}
