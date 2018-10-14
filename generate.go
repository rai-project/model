//go:generate go get -v github.com/mailru/easyjson/...
//go:generate go get github.com/alvaroloes/enumer
//go:generate easyjson -snake_case -disallow_unknown_fields -pkg .
//go:generate enumer -type=ErrorCode -json

package model
