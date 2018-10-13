//go:generate go get -v github.com/ugorji/go/codec/codecgen
//go:generate go get -v github.com/valyala/fastjson
//go:generate codecgen -nx -rt=json -u=true -o codec.generated.go acl.go base.go credentials.go roles.go job_request.go job_response.go team.go user.go response.go jobresponse_ece408.go

package model
