package model

// easyjson:json
type DockerHubCredentials struct {
	Username string `json:"username" yaml:"username" toml:"username" validate:"required"`
	Password string `json:"password" yaml:"password" toml:"password" validate:"required"`
}
