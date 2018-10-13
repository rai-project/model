package model

// easyjson:json
type BuildCommands []string

// easyjson:json
type Push struct {
	Push        bool                 `json:"push" yaml:"push"`
	ImageName   string               `json:"image_name" yaml:"image_name"`
	Registry    string               `json:"registry" yaml:"registry"`
	Credentials DockerHubCredentials `json:"credentials" yaml:"credentials"`
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
	CPU           CPUResources   `json:"cpu" yaml:"cpu" validate:"required,dive,required"`
	GPU           *GPUResources  `json:"gpu,omitempty" yaml:"gpu"`
	DataResources *DataResources `json:"dataresources,omitempty" yaml:"dataresources"`
	Network       bool           `json:"network" yaml:"network"`
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

type MountSource struct {
	URL     string `json:"url" yaml:"url" validate:"required"`
	Extract bool   `json:"extract" yaml:"extract"`
}

// easyjson:json
type MountResource struct {
	Mount  string      `json:"mount" yaml:"mount" validate:"required"` // mount location
	Source MountSource `json:"source" yaml:"source"`
	Cache  bool        `json:"cache" yaml:"cache"`
}

// easyjson:json
type DataResources []MountResource
