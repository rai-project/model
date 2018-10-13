package model

import "time"

type JobResponse struct {
	ID        string       `json:"id"`
	Kind      ResponseKind `json:"kind"`
	Body      []byte       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
}
