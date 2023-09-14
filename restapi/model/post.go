package model

import "time"

type Post struct {
	Id        string    `json:"id" db:"id" binding:"required"`
	ParentId  string    `json:"parentId" db:"parent_id" binding:"required"`
	User      *User     `json:"user" db:"user"`
	Body      string    `json:"body" db:"body" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" binding:"required"`
}
