package model

import "time"

type Base struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
