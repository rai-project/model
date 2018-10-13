package model

// ffjson: nodecoder
type Credentials struct {
	Username string `json:"username" yaml:"username" toml:"username"`
	Password string `json:"password" yaml:"password" toml:"password"`
}
