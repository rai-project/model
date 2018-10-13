// +build ece408ProjectMode

package model

import (
	"hash/fnv"
	"math"
	"strconv"
	"time"

	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/database/mongodb"
)

// Ece408Inference is an inference record for ECE 408 Spring 2018
type Ece408Inference struct {
	Model              string
	Correctness        float64
	OpRuntimes         []time.Duration // run times reported by the layer invocations
	UserFullRuntime    time.Duration   // user from /usr/bin/time
	SystemFullRuntime  time.Duration   // system from /usr/bin/time
	ElapsedFullRuntime time.Duration   // elapsed from /usr/bin/time
}

// Ranking holds info used to track team rankings
type Ranking struct {
	CreatedAt     time.Time `json:"created_at"  bson:"created_at"`
	Username      string
	Teamname      string
	ProjectURL    string // where the file was uploaded
	IsSubmission  bool   `bson:"is_submission"`  // is a final submission
	SubmissionTag string `bson:"submission_tag"` // more info about the submission
}

type Ece408JobResponse struct {
	Ranking    `bson:",inline"`
	Inferences []Ece408Inference
}

func (j *Ece408JobResponse) MinOpRuntime() time.Duration {
	minSeen := int64(math.MaxInt64)

	for _, i := range j.Inferences {
		var TotalOpRunTime int64
		TotalOpRunTime = 0
		for _, h := range i.OpRuntimes {
			TotalOpRunTime += int64(h)
		}
		if TotalOpRunTime < minSeen {
			minSeen = TotalOpRunTime
		}
	}
	return time.Duration(minSeen)
}

func anonymizeString(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s + ":::" + config.App.Secret))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

// Anonymize produces an anonymous Ece408JobResponse
func (r Ece408JobResponse) Anonymize() Ece408JobResponse {
	h := fnv.New32a()
	h.Write([]byte(r.Teamname + ":::" + config.App.Secret))
	ret := r
	ret.Username = anonymizeString(r.Username)
	ret.Teamname = anonymizeString(r.Teamname)
	ret.ProjectURL = "--"
	return ret
}

func (Ece408JobResponse) TableName() string {
	return "rankings"
}

type Ece408JobResponseCollection struct {
	*mongodb.MongoTable
}

func NewEce408JobResponseCollection(db database.Database) (*Ece408JobResponseCollection, error) {
	tbl, err := mongodb.NewTable(db, Ece408JobResponse{}.TableName())
	if err != nil {
		return nil, err
	}
	tbl.Create(nil)

	return &Ece408JobResponseCollection{
		MongoTable: tbl.(*mongodb.MongoTable),
	}, nil
}

func (m *Ece408JobResponseCollection) Close() error {
	return nil
}

type Ece408JobResponses []Ece408JobResponse

func KeepFirstTeam(rs Ece408JobResponses) Ece408JobResponses {
	res := Ece408JobResponses{}

	seen := map[string]interface{}{}
	for _, r := range rs {
		if _, ok := seen[r.Teamname]; !ok {
			res = append(res, r)
			seen[r.Teamname] = true
		}
	}

	return res
}

// FilterNonZeroTimes returns jobs with non-zero runtime
func FilterNonZeroTimes(js Ece408JobResponses) Ece408JobResponses {
	res := Ece408JobResponses{}

	for _, j := range js {
		inferenceGood := true
		for _, i := range j.Inferences {
			var TotalOpRunTime int64
			TotalOpRunTime = 0
			for _, h := range i.OpRuntimes {
				TotalOpRunTime += int64(h)
			}
			if TotalOpRunTime == 0 {
				inferenceGood = false
				break
			}
		}
		if inferenceGood {
			res = append(res, j)
		}
	}

	return res
}

type ByMinOpRuntime []Ece408JobResponse

func (r ByMinOpRuntime) Len() int           { return len(r) }
func (r ByMinOpRuntime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByMinOpRuntime) Less(i, j int) bool { return r[i].MinOpRuntime() < r[j].MinOpRuntime() }

func (j *Ece408JobResponse) StartNewInference() {
	j.Inferences = append(j.Inferences, Ece408Inference{})
}

func (j *Ece408JobResponse) CurrentInference() *Ece408Inference {
	if len(j.Inferences) == 0 {
		j.StartNewInference()
	}
	return &j.Inferences[len(j.Inferences)-1]
}

func (j *Ece408JobResponse) RecordOpRuntime(duration time.Duration) {
	j.CurrentInference().OpRuntimes = append(j.CurrentInference().OpRuntimes, duration)
}
