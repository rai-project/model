package model

import (
	"time"

	"github.com/rai-project/config"
)

// easyjson:json
type JobResponse struct {
	ID        string                 `json:"id" validate:"required"`
	Kind      ResponseKind           `json:"kind" validate:"required"`
	Body      []byte                 `json:"body" validate:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time              `json:"created_at" validate:"required"`
}

func (JobResponse) TableName() string {
	return config.App.Name + "_jobs"
}
