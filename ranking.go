package model

import (
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

// Ranking holds info used to track team rankings
type Ranking struct {
	Base
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
	BatchSize         int
	Correctness       float64
	OpRuntime         time.Duration // runtime reported by the layer
	UserFullRuntime   time.Duration // runtime captured using time
	SystemFullRuntime time.Duration // full runtime captures using time
}

type Fa2017Ece408Rankings []Fa2017Ece408Ranking
