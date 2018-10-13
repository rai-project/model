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
	Model              string          `json:"model,omitempty"`
	Correctness        float64         `json:"correctness,omitempty"`
	OpRuntimes         []time.Duration `json:"op_runtimes,omitempty"`          // run times reported by the layer invocations
	UserFullRuntime    time.Duration   `json:"user_full_runtime,omitempty"`    // user from /usr/bin/time
	SystemFullRuntime  time.Duration   `json:"system_full_runtime,omitempty"`  // system from /usr/bin/time
	ElapsedFullRuntime time.Duration   `json:"elapsed_full_runtime,omitempty"` // elapsed from /usr/bin/time
}

// Ranking holds info used to track team rankings
type ECE408Ranking struct {
	Base          `json:",inline"`
	Username      string `json:"username,omitempty"`
	Teamname      string `json:"teamname,omitempty"`
	ProjectURL    string `json:"project_url,omitempty"`                          // where the file was uploaded
	IsSubmission  bool   `bson:"is_submission" json:"is_submission,omitempty"`   // is a final submission
	SubmissionTag string `bson:"submission_tag" json:"submission_tag,omitempty"` // more info about the submission
}

type Ece408JobResponseBody struct {
	ECE408Ranking `bson:",inline"`
	Inferences    []Ece408Inference
}

func (j *Ece408JobResponseBody) MinOpRuntime() time.Duration {
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

// Anonymize produces an anonymous Ece408JobResponseBody
func (r Ece408JobResponseBody) Anonymize() Ece408JobResponseBody {
	h := fnv.New32a()
	h.Write([]byte(r.Teamname + ":::" + config.App.Secret))
	ret := r
	ret.Username = anonymizeString(r.Username)
	ret.Teamname = anonymizeString(r.Teamname)
	ret.ProjectURL = "--"
	return ret
}

func (Ece408JobResponseBody) TableName() string {
	return "ece408_rankings"
}

type Ece408JobResponseBodyCollection struct {
	*mongodb.MongoTable
}

func NewEce408JobResponseBodyCollection(db database.Database) (*Ece408JobResponseBodyCollection, error) {
	tbl, err := mongodb.NewTable(db, Ece408JobResponseBody{}.TableName())
	if err != nil {
		return nil, err
	}
	tbl.Create(nil)

	return &Ece408JobResponseBodyCollection{
		MongoTable: tbl.(*mongodb.MongoTable),
	}, nil
}

func (m *Ece408JobResponseBodyCollection) Close() error {
	return nil
}

type Ece408JobResponseBodys []Ece408JobResponseBody

func KeepFirstTeam(rs Ece408JobResponseBodys) Ece408JobResponseBodys {
	res := Ece408JobResponseBodys{}

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
func FilterNonZeroTimes(js Ece408JobResponseBodys) Ece408JobResponseBodys {
	res := Ece408JobResponseBodys{}

	for _, j := range js {
		inferenceGood := true
		for _, i := range j.Inferences {
			TotalOpRunTime := int64(0)
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

type ByMinOpRuntime []Ece408JobResponseBody

func (r ByMinOpRuntime) Len() int           { return len(r) }
func (r ByMinOpRuntime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByMinOpRuntime) Less(i, j int) bool { return r[i].MinOpRuntime() < r[j].MinOpRuntime() }

func (j *Ece408JobResponseBody) StartNewInference() {
	j.Inferences = append(j.Inferences, Ece408Inference{})
}

func (j *Ece408JobResponseBody) CurrentInference() *Ece408Inference {
	if len(j.Inferences) == 0 {
		j.StartNewInference()
	}
	return &j.Inferences[len(j.Inferences)-1]
}

func (j *Ece408JobResponseBody) RecordOpRuntime(duration time.Duration) {
	j.CurrentInference().OpRuntimes = append(j.CurrentInference().OpRuntimes, duration)
}
