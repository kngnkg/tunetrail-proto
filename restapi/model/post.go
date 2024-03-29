package model

import "time"

type Post struct {
	Id       string `json:"id" db:"id" binding:"required"`
	ParentId string `json:"parentId" db:"parent_id" binding:"required"`
	Body     string `json:"body" db:"body" binding:"required"`
	User     *User  `json:"user" db:"user"`
	LikeInfo
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at" binding:"required"`
}

type Pagination struct {
	NextCursor     string `json:"nextCursor"`     // 次のページのカーソル
	PreviousCursor string `json:"previousCursor"` // 前のページのカーソル
	Limit          int    `json:"limit"`          // 1ページあたりの件数
}
