package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// easyjson:json
type Base struct {
	ID        bson.ObjectId `json:"id" bson:"_id" gorm:"primary_key" toml:"id,omitempty" validate:"required"`
	CreatedAt time.Time     `json:"created_at" toml:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time     `json:"updated_at" toml:"updated_at,omitempty" validate:"required"`
	DeletedAt *time.Time    `json:"deleted_at" sql:"index" toml:"deleted_at,omitempty"`
}
