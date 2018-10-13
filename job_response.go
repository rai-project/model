package model

import (
	"time"

	"github.com/rai-project/config"
)

// ffjson: nodecoder
type JobResponse struct {
	ID        string       `json:"id"`
	Kind      ResponseKind `json:"kind"`
	Body      []byte       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
}

func (JobResponse) TableName() string {
	return config.App.Name + "_jobs"
}
