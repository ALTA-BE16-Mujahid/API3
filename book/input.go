package book

import "github.com/mujahxd/api3-jwt/user"

type CreateBookInput struct {
	Title     string `json:"title" binding:"required"`
	Year      int    `json:"year" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	User      user.User
}

type GetBookDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
