package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Team struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name"`
}
