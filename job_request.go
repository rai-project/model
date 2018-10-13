package model

import (
	"strings"

	"github.com/rai-project/config"
)

// easyjson:json
type JobRequest struct {
	Base
	ClientVersion      config.VersionInfo `json:"client_version" validate:"required"`
	UploadKey          string             `json:"upload_key"`
	User               User               `json:"user" validate:"required,dive,required" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification" validate:"required,dive,required"`
}

// easyjson:json
type BuildCommands []string

// easyjson:json
type Push struct {
	Push        bool        `json:"push" yaml:"push"`
	ImageName   string      `json:"image_name" yaml:"image_name"`
	Registry    string      `json:"registry" yaml:"registry"`
	Credentials Credentials `json:"credentials" yaml:"credentials"`
}

// easyjson:json
type BuildImageSpecification struct {
	ImageName  string `json:"image_name" yaml:"image_name"`
	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
	Push       *Push  `json:"push" yaml:"push"`
	NoCache    bool   `json:"no_cache" yaml:"no_cache"`
}

// easyjson:json
type RAIBuildSpecification struct {
	Version        string `json:"version" yaml:"version"`
	ContainerImage string `json:"image" yaml:"image"`
}

// easyjson:json
type CommandsBuildSpecification struct {
	BuildImage *BuildImageSpecification `json:"build_image,omitempty" yaml:"build_image,omitempty"`
	Build      BuildCommands            `json:"build" yaml:"build"`
}

// easyjson:json
type BuildSpecification struct {
	RAI       RAIBuildSpecification      `json:"rai" yaml:"rai" validate:"required"`
	Resources Resources                  `json:"resources" yaml:"resources" validate:"required"`
	Commands  CommandsBuildSpecification `json:"commands" yaml:"commands" validate:"required"`
}

// easyjson:json
type Resources struct {
	CPU     CPUResources  `json:"cpu" yaml:"cpu" validate:"required,dive,required"`
	GPU     *GPUResources `json:"gpu" yaml:"gpu"`
	Network bool          `json:"network" yaml:"network"`
}

// easyjson:json
type CPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture" validate:"required"`
}

// easyjson:json
type GPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture" validate:"required"`
	Count        int    `json:"count" yaml:"count" validate:"required"`
	Memory       int64  `json:"memory" yaml:"memory" validate:"required"`
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
