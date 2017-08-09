package model

import (
	"time"
)

type JobRequest struct {
	Base
	UploadKey          string             `json:"upload_key"`
	User               User               `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification"`
}

type BuildCommands []string

type Credentials struct {
	Username string `json:"user_name" yaml:"user_name"`
	Password string `json:"password" yaml:"password"`
}

type Push struct {
	Push        bool        `json:"push" yaml:"push"`
	ImageName   string      `json:"image_name" yaml:"image_name"`
	Registry    string      `json:"registry" yaml:"registry"`
	Credentials Credentials `json:"credentials" yaml:"credentials"`
}

type BuildImageSpecification struct {
	ImageName  string `json:"image_name" yaml:"image_name"`
	Dockerfile string `json:"dockerfile" yaml:"dockerfile"`
	Push       *Push  `json:"push" yaml:"push"`
	NoCache    bool   `json:"no_cache" yaml:"no_cache"`
}

type RAIBuildSpecification struct {
	Version        string `json:"version" yaml:"version"`
	ContainerImage string `json:"image" yaml:"image"`
}

type CommandsBuildSpecification struct {
	BuildImage *BuildImageSpecification `json:"build_image,omitempty" yaml:"build_image,omitempty"`
	Build      BuildCommands            `json:"build" yaml:"build"`
}

type BuildSpecification struct {
	RAI       RAIBuildSpecification      `json:"rai" yaml:"rai"`
	Resources Resources                  `json:"resources" yaml:"resources"`
	Commands  CommandsBuildSpecification `json:"commands" yaml:"commands"`
}

type Resources struct {
	CPU     CPUResources  `json:"cpu" yaml:"cpu"`
	GPU     *GPUResources `json:"gpu" yaml:"gpu"`
	Network bool          `json:"network" yaml:"network"`
}

type CPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture"`
}

type GPUResources struct {
	Architecture string `json:"architecture" yaml:"architecture"`
	Count        int    `json:"count" yaml:"count"`
	Memory       int64  `json:"memory" yaml:"memory"`
}

type ResponseKind string

const (
	StderrResponse ResponseKind = "Stderr"
	StdoutResponse ResponseKind = "Stdout"
	StdinResponse  ResponseKind = "Stdin"
	EndResponse    ResponseKind = "End"
)

type JobResponse struct {
	ID        string       `json:"id"`
	Kind      ResponseKind `json:"kind"`
	Body      []byte       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
}

func (j *JobRequest) PushQ() bool {
	buildImage := j.BuildSpecification.Commands.BuildImage
	if buildImage == nil {
		return false
	}
	push := buildImage.Push
	if push == nil {
		return false
	}
	if push.Push == false {
		return false
	}
	if push.Credentials.Username == "" {
		return false
	}
	if push.Credentials.Password == "" {
		return false
	}
	return true
}
