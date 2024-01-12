package pjson

import (
	"os"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
)

var j = `[
	{
		"Id": 3,
		"Description": "Cook eggs",
		"Duration": 600000000000,
		"Start": "0001-01-01T00:00:00Z",
		"Finish": "0001-01-01T00:00:00Z",
		"PredecessorsId": [
			2
		],
		"SuccessorsId": [
			1
		],
		"Cost": 0
	},
	{
		"Id": 2,
		"Description": "Buy eggs",
		"Duration": 1800000000000,
		"Start": "0001-01-01T00:00:00Z",
		"Finish": "0001-01-01T00:00:00Z",
		"PredecessorsId": [],
		"SuccessorsId": [
			3
		],
		"Cost": 100
	},
	{
		"Id": 1,
		"Description": "Eat eggs",
		"Duration": 1200000000000,
		"Start": "0001-01-01T00:00:00Z",
		"Finish": "0001-01-01T00:00:00Z",
		"PredecessorsId": [
			3
		],
		"SuccessorsId": [],
		"Cost": 0
	}
]`
var d1 = time.Minute * 10
var d2 = time.Minute * 30
var d3 = time.Minute * 20
var tt = time.Time{}
var activities = []*activity.Activity{
	{Id: 3, Description: "Cook eggs", Duration: d1, PredecessorsId: []int{2}, SuccessorsId: []int{1}, Start: tt, Finish: tt, Cost: 0},
	{Id: 2, Description: "Buy eggs", Duration: d2, PredecessorsId: []int{}, SuccessorsId: []int{3}, Start: tt, Finish: tt, Cost: 100},
	{Id: 1, Description: "Eat eggs", Duration: d3, PredecessorsId: []int{3}, SuccessorsId: []int{}, Start: tt, Finish: tt, Cost: 0},
}

func TestExportToDb(t *testing.T) {
	sqldb, err := db.New("test.db")
	if err != nil {
		t.Error(err)
	}
	err = ExportToDb([]byte(j), sqldb.DB)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove("test.db")
	if err != nil {
		t.Error(err)
	}
}

func TestActivitiesToJson(t *testing.T) {
	b, err := ActivitiesToJson(activities)
	if err != nil {
		t.Error(err)
	}

	sb := string(b)
	if sb != j {
		t.Errorf("error parsing json: want %s, got %s\n", j, sb)
	}
}
