package model

import "time"

// Ranking holds info used to track project team rankings
type Ranking struct {
	OpFullRuntime time.Duration
	S3Key         string
	ProjectURL    string
}
