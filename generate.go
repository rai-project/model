//go:generate go get -v github.com/mailru/easyjson/...
//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate easyjson -snake_case -disallow_unknown_fields -pkg .
//go:generate stringer -type=ErrorCode

package model
