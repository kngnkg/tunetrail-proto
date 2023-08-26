package model

import "time"

type Post struct {
	Id        string    `json:"id" db:"id" binding:"required"`
	UserId    UserID    `json:"userId" db:"user_id" binding:"required"`
	Body      string    `json:"body" db:"body" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" binding:"required"`
}
