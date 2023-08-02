package model

type Tokens struct {
	Id      string `json:"id" binding:"required"`
	Access  string `json:"access" binding:"required"`
	Refresh string `json:"refresh" binding:"required"`
}
