package pxlsx

import (
	"os"
	"testing"
	"time"

	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
)

var d1 = time.Minute * 10
var d2 = time.Minute * 30
var d3 = time.Minute * 20
var tt = time.Time{}
var activities = []*activity.Activity{
	{Id: 3, Description: "Cook eggs", Duration: d1, PredecessorsId: []int{2}, SuccessorsId: []int{1}, Start: tt, Finish: tt, Cost: 0},
	{Id: 2, Description: "Buy eggs", Duration: d2, PredecessorsId: []int{}, SuccessorsId: []int{3}, Start: tt, Finish: tt, Cost: 100},
	{Id: 1, Description: "Eat eggs", Duration: d3, PredecessorsId: []int{3}, SuccessorsId: []int{}, Start: tt, Finish: tt, Cost: 0},
}

func TestActivitiesToXLSX(t *testing.T) {
	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("test")
	if err != nil {
		t.Error(err)
	}
	defer sheet.Close()
	ActivitiesToXLSX(activities, sheet)
	err = wb.Save("test.xlsx")
	if err != nil {
		t.Error(err)
	}
}

func TestXLSXtoActivities(t *testing.T) {
	wb, err := xlsx.OpenFile("test.xlsx")
	if err != nil {
		t.Error(err)
	}
	acts, err := XLSXToActivities(wb.Sheet["test"])
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(activities); i++ {
		if acts[i].Id != activities[i].Id {
			t.Errorf("got %+v, want %+v", acts[i], activities[i])
		}
	}
}

func TestExportToDb(t *testing.T) {
	sqldb, err := db.New("test.db")
	if err != nil {
		t.Error(err)
	}
	wb, err := xlsx.OpenFile("test.xlsx")
	if err != nil {
		t.Error(err)
	}
	sheet := wb.Sheet["test"]
	if err != nil {
		t.Error(err)
	}
	err = ExportToDb(sqldb.DB, sheet)
	if err != nil {
		t.Error(err)
	}

	err = os.Remove("test.db")
	if err != nil {
		t.Error(err)
	}
	err = os.Remove("test.xlsx")
	if err != nil {
		t.Error(err)
	}
}
