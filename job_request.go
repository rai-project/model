package model

import (
	"strings"

	"github.com/rai-project/config"
)

// ffjson
type JobRequest struct {
	Base
	ClientVersion      config.VersionInfo `json:"client_version"`
	UploadKey          string             `json:"upload_key"`
	User               User               `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification"`
}

// ffjson
type BuildCommands []string

// ffjson
type Push struct {
	Push        bool        `json:"push" yaml:"push"`
	ImageName   string      `json:"image_name" yaml:"image_name"`
	Registry    string      `json:"registry" yaml:"registry"`
	Credentials Credentials `json:"credentials" yaml:"credentials"`
}

// ffjson
type BuildImageSpecification struct {
	ImageName  string `json:"image_name" yaml:"image_name"`
	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
	Push       *Push  `json:"push" yaml:"push"`
	NoCache    bool   `json:"no_cache" yaml:"no_cache"`
}

// ffjson
type RAIBuildSpecification struct {
	Version        string `json:"version" yaml:"version"`
	ContainerImage string `json:"image" yaml:"image"`
}

// ffjson
type CommandsBuildSpecification struct {
	BuildImage *BuildImageSpecification `json:"build_image,omitempty" yaml:"build_image,omitempty"`
	Build      BuildCommands            `json:"build" yaml:"build"`
}

// ffjson
type BuildSpecification struct {
	RAI       RAIBuildSpecification      `json:"rai" yaml:"rai"`
	Resources Resources                  `json:"resources" yaml:"resources"`
	Commands  CommandsBuildSpecification `json:"commands" yaml:"commands"`
}

// ffjson
type Resources struct {
	CPU     CPUResources  `json:"cpu" yaml:"cpu"`
	GPU     *GPUResources `json:"gpu" yaml:"gpu"`
	Network bool          `json:"network" yaml:"network"`
}

// ffjson
type CPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture"`
}

// ffjson
type GPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture"`
	Count        int    `json:"count" yaml:"count"`
	Memory       int64  `json:"memory" yaml:"memory"`
}

// ffjson
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
