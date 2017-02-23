package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User     *User         `json:"user"`
	Commands []byte        `json:"commands"`
}
