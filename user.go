package model

import (
	"github.com/rai-project/acl"
	"gopkg.in/mgo.v2/bson"
)

// easyjson:json
type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id" gorm:"primary_key" toml:"id,omitempty" validate:"required"`
	Base        `toml:"-" yaml:"-" validate:"required,dive,required"`
	Firstname   string                `json:"firstname" yaml:"firstname,omitempty" toml:"firstname,omitempty" validate:"required"`
	Lastname    string                `json:"lastname" yaml:"lastname,omitempty" toml:"lastname,omitempty" validate:"required"`
	Username    string                `json:"username" yaml:"username" toml:"username"`
	Email       string                `json:"email" yaml:"email" toml:"email" validate:"required,email"`
	AccessKey   string                `json:"access_key" yaml:"access_key" toml:"access_key" validate:"required"`
	SecretKey   string                `json:"secret_key" yaml:"secret_key" toml:"secret_key" validate:"required"`
	Affiliation string                `json:"affiliation" yaml:"affiliation" toml:"affiliation"`
	Password    string                `json:"password" yaml:"-" toml:"-"`
	Team        *Team                 `json:"team" gorm:"ForeignKey:TeamID;AssociationForeignKey:Refer" toml:"-"`
	DockerHub   *DockerHubCredentials `json:"dockerhub,omitempty" yaml:"dockerhub,omitempty" toml:"dockerhub,omitempty"`
	Role        acl.Role              `json:"role" yaml:"role" validate:"required"`
}

func (User) TableName() string {
	return "users"
}

// func UserStructLevelValidation(sl validator.StructLevel) {

// 	user := sl.Current().Interface().(User)

// 	if len(user.Firstname) == 0 && len(user.Lastname) == 0 {
// 		sl.ReportError(user.Firstname, "Firstname", "firstname", "fnameorlname", "")
// 		sl.ReportError(user.Lastname, "Lastname", "lname", "fnameorlname", "")
// 	}

// 	// plus can to more, even with different tag than "fnameorlname"

// }

// func init() {
// 	validate.RegisterStructValidation(UserStructLevelValidation, User{})
// }
