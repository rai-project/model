package model

import "time"

// ffjson: nodecoder
type ACL struct {
	MaxRuntimeLimit time.Duration
	MaxStorageLimit uint64
	NetworkAccess   bool
	MountAccess     bool
}
