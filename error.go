package model

import (
	json "encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Error struct {
	code  int      `json:"code,omitempty"`
	name  string   `json:"name,omitempty"`
	error error    `json:"error,omitempty"`
	stack []string `json:"stack,omitempty"`
}

func (e *Error) MarshalJSON() ([]byte, error) {
	code := fmt.Sprintf("\"code\": \"%v\"", e.code)
	name := fmt.Sprintf("\"name\": \"%v\"", e.name)
	message := fmt.Sprintf("\"message\": \"%v\"", e.error)

	var stack string
	stackData := strings.Split(fmt.Sprintf("%+v", errors.WithStack(e.error)), "\n")
	bts, err := json.Marshal(stackData)
	if err != nil {
		stack = fmt.Sprintf("\"stack\": []")
	} else {
		stack = fmt.Sprintf("\"stack\": %s", string(bts))
	}

	res := fmt.Sprintf("{%s, %s, %s, %s}", name, message, code, stack)
	return []byte(res), nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%v", e.name, e.error.Error())
}

func (e *Error) WithCode(code int) *Error {
	e.code = code
	return e
}

func (e *Error) WithName(name string) *Error {
	e.name = name
	return e
}

func (e *Error) Wrap(err error) *Error {
	e.error = err
	return e
}

func (e *Error) Code() int32 {
	return int32(e.code)
}
