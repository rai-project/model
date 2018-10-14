package model

type ErrorCode int

const (
	ErrorCodeSuccess ErrorCode = iota
	ErrorCodeInvalidSpec
	ErrorCodeInvalidQueue
	ErrorCodeInvalid
	ErrorCodeTimeoutLimit
	ErrorCodeStorageLimit
	ErrorCodeGPUCountLimit
	ErrorCodeLimit
	ErrorCodePermission
	ErrorCodeNoServer
	ErrorCodeUnknown
)

const (
	Success = ErrorCodeSuccess
)
