package models

import (
	"time"
)

type (
	User struct {
		Id        int       `json: "id"`
		NickName  string    `json: "nick_name"`
		Email     string    `json: "email"`
		CreatedAt time.Time `json: "created_at"`
		UpdatedAt time.Time `json: "updated_at"`
	}

	Recipe struct {
		Id          int       `json: "id"`
		UserId      string    `json: "user_id"`
		Title       string    `json: "title"`
		Description string    `json: "description"`
		CreatedAt   time.Time `json: "created_at"`
		UpdatedAt   time.Time `json: "updated_at"`
	}
)
