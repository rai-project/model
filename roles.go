package model

import "strings"

type Role string

var Roles = []Role{
	"power",
	"student",
	"ece408_student",
	"guest",
	"admin",
}

func (r Role) ToACL() ACL {
	return ACL{}
}

func (r0 Role) Validate() bool {
	r := strings.ToLower(string(r0))
	for _, e := range Roles {
		if string(e) == r {
			return true
		}
	}
	return false
}
