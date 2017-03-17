package model

type User struct {
	Base        `toml:"-" yaml:"-"`
	Firstname   string `json:"firstname" yaml:"firstname,omitempty" toml:"firstname,omitempty"`
	Lastname    string `json:"lastname" yaml:"lastname,omitempty" toml:"lastname,omitempty"`
	Username    string `json:"username" yaml:"username" toml:"username"`
	Email       string `json:"email" yaml:"email" toml:"email"`
	AccessKey   string `json:"access_key" yaml:"access_key" toml:"access_key"`
	SecretKey   string `json:"secret_key" yaml:"secret_key" toml:"secret_key"`
	Affiliation string `json:"affiliation" yaml:"affiliation" toml:"affiliation"`
	Password    string `json:"password" yaml:"-" toml:"-"`
	// Team      *Team  `json:"team" gorm:"ForeignKey:TeamID;AssociationForeignKey:Refer" toml:"-"`
}
