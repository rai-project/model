package model

import (
	"time"
)

// easyjson:json
type Base struct {
	CreatedAt time.Time  `json:"created_at" toml:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at" toml:"updated_at,omitempty" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index" toml:"deleted_at,omitempty"`
}
