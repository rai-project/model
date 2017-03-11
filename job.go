package model

type JobRequest struct {
	Base
	User               User               `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	BuildSpecification BuildSpecification `json:"build_specification"`
}

type BuildSpecification struct {
	RAI struct {
		Version        string `json:"version" yaml:"version"`
		ContainerImage string `json:"image" yaml:"image"`
	} `json:"rai" yaml:"rai"`
	Resources struct {
		GPUs int `json:"gpus" yaml:"gpus"`
	} `json:"resources" yaml:"resources"`
	Commands struct {
		Build []string `json:"build" yaml:"build"`
	} `json:"commands" yaml:"commands"`
}
