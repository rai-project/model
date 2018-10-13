package model

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
