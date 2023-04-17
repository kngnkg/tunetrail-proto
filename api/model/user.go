package model

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	UserName  string    `json:"userName" db:"user_name"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	IconUrl   string    `json:"iconUrl" db:"icon_url"`
	Bio       string    `json:"bio" db:"bio"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Users []User
