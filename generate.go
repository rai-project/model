//go:generate go get -v github.com/pquerna/ffjson
//go:generate go get -v github.com/valyala/fastjson
//go:generate ffjson -nodecoder acl.go
//go:generate ffjson -nodecoder base.go
//go:generate ffjson -nodecoder credentials.go
//go:generate ffjson -nodecoder ece408_jobresponse.go
//go:generate ffjson -nodecoder job_request.go
//go:generate ffjson -nodecoder job_response.go
//go:generate ffjson -nodecoder response.go
//go:generate ffjson -nodecoder team.go
//go:generate ffjson -nodecoder user.go

package model
