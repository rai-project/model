package model

import "time"

// easyjson:json
type ACL struct {
	MaxRuntimeLimit time.Duration
	MaxStorageLimit uint64
	NetworkAccess   bool
	MountAccess     bool
}
