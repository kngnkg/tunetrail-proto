package model

type LikeInfo struct {
	LikesCount int  `json:"likesCount" db:"likes_count" binding:"required"`
	Liked      bool `json:"liked" db:"liked" binding:"required"`
}
