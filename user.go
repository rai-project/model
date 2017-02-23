package model

type User struct {
	Base
	Name      string `json:"name"`
	Team      *Team  `json:"team" gorm:"ForeignKey:TeamID;AssociationForeignKey:Refer"`
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
}
