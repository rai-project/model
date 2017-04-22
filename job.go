package model

import "time"

type JobRequest struct {
	Base
	UploadKey          string             `json:"upload_key"`
	User               User               `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification"`
}

type BuildCommands []string

type BuildImageSpecification struct {
	Title            string `json:"title" yaml:"title"`
	Type             string `json:"type" yaml:"type"`
	ImageName        string `json:"image_name" yaml:"image_name"`
	WorkingDirectory string `json:"working_directory" yaml:"working_directory"`
	Dockerfile       string `json:"dockerfile" yaml:"dockerfile"`
}

type BuildSpecification struct {
	RAI struct {
		Version        string `json:"version" yaml:"version"`
		ContainerImage string `json:"image" yaml:"image"`
	} `json:"rai" yaml:"rai"`
	Resources Resources `json:"resources" yaml:"resources"`
	Commands  struct {
		BuildImage *BuildImageSpecification `json:"build_image,omitempty" yaml:"build_image,omitempty"`
		Build      BuildCommands            `json:"build" yaml:"build"`
	} `json:"commands" yaml:"commands"`
}

type Resources struct {
	CPU     CPUResources `json:"cpu" yaml:"cpu"`
	GPU     GPUResources `json:"gpu" yaml:"gpu"`
	Network bool         `json:"network" yaml:"network"`
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
	Kind      ResponseKind `json:"kind"`
	Body      []byte       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
}
