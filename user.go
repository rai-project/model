package model

type User struct {
	Base
	Firstname string `json:"firstname" toml:"firstname,omitempty"`
	Lastname  string `json:"lastname" toml:"lastname,omitempty"`
	Username  string `json:"username" toml:"username"`
	Email     string `json:"email" toml:"email"`
	AccessKey string `json:"access_key" toml:"access_key"`
	SecretKey string `json:"secret_key" toml:"secret_key"`
	Password  string `json:"password" toml:"-"`
	// Team      *Team  `json:"team" gorm:"ForeignKey:TeamID;AssociationForeignKey:Refer" toml:"-"`
}
