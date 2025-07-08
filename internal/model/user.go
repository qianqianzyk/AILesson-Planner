package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Uuid      int       `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
