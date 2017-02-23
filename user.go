package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Team      *Team         `json:"team"`
	SecretKey string        `json:"secret_key"`
	AccessKey string        `json:"access_key"`
}
