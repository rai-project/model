package model

type Credentials struct {
	Username string `json:"user_name" yaml:"user_name" toml:"user_name"`
	Password string `json:"password" yaml:"password" toml:"user_name"`
}
