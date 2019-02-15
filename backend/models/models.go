package models

import (
	"time"
)

type (
	Recipe struct {
		Id        int `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json: "id"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
