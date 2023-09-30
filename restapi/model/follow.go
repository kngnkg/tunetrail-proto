package model

type FollowInfo struct {
	IsFollowing bool `json:"isFollowing" db:"is_following"`
	IsFollowed  bool `json:"isFollowed" db:"is_followed"`
}
