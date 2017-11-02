package model

import (
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
	"gopkg.in/mgo.v2/bson"
)

// Ranking holds info used to track team rankings
type Ranking struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt  time.Time     `json:"created_at"  bson:"created_at"`
	Username   string
	Teamname   string
	ProjectURL string // where the file was uploaded
}

// Anonymize produces an anonymous Ranking
func (r Ranking) Anonymize() Ranking {
	ret := r
	teamname := strings.Join(strings.Split(namesgenerator.GetRandomName(0), "_"), " ")
	ret.Teamname = strings.Title(teamname)
	ret.ProjectURL = "--"
	return ret
}

// Fa2017Ece408Ranking holds fields specific to ECE408
type Fa2017Ece408Ranking struct {
	Ranking
	Model              string
	BatchSize          int
	Correctness        float64
	OpRuntime          time.Duration // runtime reported by the layer
	UserFullRuntime    time.Duration // user from /usr/bin/time
	SystemFullRuntime  time.Duration // system from /usr/bin/time
	ElapsedFullRuntime time.Duration // elapsed from /usr/bin/time
}

type Fa2017Ece408Rankings []Fa2017Ece408Ranking
