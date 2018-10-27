package model

import (
	"strings"

	"github.com/rai-project/config"
	"gopkg.in/mgo.v2/bson"
)

// easyjson:json
type JobRequest struct {
	ID                 bson.ObjectId `json:"_id" bson:"_id" gorm:"primary_key" toml:"id,omitempty" validate:"required"`
	Base               `json:",inline" bson:",inline"`
	ClientVersion      config.VersionInfo `json:"client_version" validate:"required"`
	UploadKey          string             `json:"upload_key"`
	User               User               `json:"user" validate:"required,dive,required" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification" validate:"required,dive,required"`
}

type ResponseKind string

const (
	StderrResponse ResponseKind = "Stderr"
	StdoutResponse ResponseKind = "Stdout"
	StdinResponse  ResponseKind = "Stdin"
	EndResponse    ResponseKind = "End"
)

func (b *BuildImageSpecification) PushQ() bool {
	if b == nil {
		return false
	}
	push := b.Push
	if push == nil {
		return false
	}
	if !push.Push {
		return false
	}
	if strings.TrimSpace(push.Credentials.Username) == "" {
		return false
	}
	if strings.TrimSpace(push.Credentials.Password) == "" {
		return false
	}
	return true
}

func (j *JobRequest) PushQ() bool {
	buildImage := j.BuildSpecification.Commands.BuildImage
	return buildImage != nil && buildImage.PushQ()
}
