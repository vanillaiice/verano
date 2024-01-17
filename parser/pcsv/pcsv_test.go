package pcsv

import (
	"bufio"
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
)

var scsv = `Id,Description,Duration,Start,Finish,PredecessorsId,SuccessorsId,Cost
3,Cook eggs,10m0s,-62135596800,-62135596800,2,1,0
2,Buy eggs,30m0s,-62135596800,-62135596800,,3,100
1,Eat eggs,20m0s,-62135596800,-62135596800,3,,0
`
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
	r := bytes.NewReader([]byte(scsv))
	err = ExportToDb(sqldb.DB, r)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove("test.db")
	if err != nil {
		t.Error(err)
	}
}

func TestActivitiesToCSV(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err := ActivitiesToCSV(activities, w)
	if err != nil {
		t.Error(err)
	}
	sb := buf.String()
	if sb != scsv {
		t.Errorf("error parsing csv: want %s, got %s\n", scsv, sb)
	}
}

func TestCSVToActivities(t *testing.T) {
	r := bytes.NewReader([]byte(scsv))
	acts, err := CSVToActivities(r)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(activities); i++ {
		if acts[i].Id != activities[i].Id {
			t.Errorf("got %+v, want %+v", acts[i], activities[i])
		}
	}
}
