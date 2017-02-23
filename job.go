package model

type Job struct {
	Base
	User     *User  `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	Commands []byte `json:"commands"`
}
