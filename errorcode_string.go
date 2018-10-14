// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package model

import "strconv"

const _ErrorCode_name = "ErrorCodeSuccessErrorCodeInvalidSpecErrorCodeInvalidQueueErrorCodeInvalidErrorCodeTimeoutLimitErrorCodeStorageLimitErrorCodeGPUCountLimitErrorCodeLimitErrorCodePermissionErrorCodeNoServerErrorCodeUnknown"

var _ErrorCode_index = [...]uint8{0, 16, 36, 57, 73, 94, 115, 137, 151, 170, 187, 203}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
