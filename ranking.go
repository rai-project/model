package model

import (
	"hash/fnv"
	"strconv"
	"time"

	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/database/mongodb"
)

// Ranking holds info used to track team rankings
type Ranking struct {
	CreatedAt    time.Time `json:"created_at"  bson:"created_at"`
	Username     string
	Teamname     string
	ProjectURL   string // where the file was uploaded
	IsSubmission bool   `bson:"is_submission"` // is a final submission
}

// Fa2017Ece408Ranking holds fields specific to ECE408
type Fa2017Ece408Ranking struct {
	Ranking            `bson:",inline"`
	Model              string
	Correctness        float64
	OpRuntime          time.Duration // runtime reported by the layer
	UserFullRuntime    time.Duration // user from /usr/bin/time
	SystemFullRuntime  time.Duration // system from /usr/bin/time
	ElapsedFullRuntime time.Duration // elapsed from /usr/bin/time
}

func anonymizeString(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s + ":::" + config.App.Secret))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
}

// Anonymize produces an anonymous Fa2017Ece408Ranking
func (r Fa2017Ece408Ranking) Anonymize() Fa2017Ece408Ranking {
	h := fnv.New32a()
	h.Write([]byte(r.Teamname + ":::" + config.App.Secret))
	ret := r
	ret.Username = anonymizeString(r.Username)
	ret.Teamname = anonymizeString(r.Teamname)
	ret.ProjectURL = "--"
	return ret
}

func (Fa2017Ece408Ranking) TableName() string {
	return "rankings"
}

type Fa2017Ece408RankingCollection struct {
	*mongodb.MongoTable
}

func NewFa2017Ece408RankingCollection(db database.Database) (*Fa2017Ece408RankingCollection, error) {
	tbl, err := mongodb.NewTable(db, Fa2017Ece408Ranking{}.TableName())
	if err != nil {
		return nil, err
	}
	tbl.Create(nil)

	return &Fa2017Ece408RankingCollection{
		MongoTable: tbl.(*mongodb.MongoTable),
	}, nil
}

func (m *Fa2017Ece408RankingCollection) Close() error {
	return nil
}

type Fa2017Ece408Rankings []Fa2017Ece408Ranking

func KeepFirstTeam(rs Fa2017Ece408Rankings) Fa2017Ece408Rankings {
	res := Fa2017Ece408Rankings{}

	seen := map[string]interface{}{}
	for _, r := range rs {
		if _, ok := seen[r.Teamname]; !ok {
			res = append(res, r)
			seen[r.Teamname] = true
		}
	}

	return res
}

// ByOpRuntime allows sorting Fa2017Ece408Rankings by OpRuntime
type ByOpRuntime []Fa2017Ece408Ranking

func (r ByOpRuntime) Len() int           { return len(r) }
func (r ByOpRuntime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByOpRuntime) Less(i, j int) bool { return r[i].OpRuntime < r[j].OpRuntime }
