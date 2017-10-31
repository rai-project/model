package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty" gorm:"primary_key" toml:"id,omitempty"`
	CreatedAt time.Time     `json:"created_at" toml:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" toml:"updated_at,omitempty"`
	DeletedAt *time.Time    `json:"deleted_at" sql:"index" toml:"deleted_at,omitempty"`
}
