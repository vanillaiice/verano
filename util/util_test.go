package util

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
)

var activities = []*activity.Activity{
	{Id: 1, Description: "Cook eggs", Duration: 5 * time.Minute, PredecessorsId: []int{2, 4, 5}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 2, Description: "Buy eggs", Duration: 10 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 3, Description: "Eat eggs", Duration: 5 * time.Minute, PredecessorsId: []int{1, 6}, SuccessorsId: []int{}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 4, Description: "Buy pan", Duration: 10 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 5, Description: "Buy salt", Duration: 6 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 6, Description: "Tip landlord", Duration: 60 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 7, Description: "Get money", Duration: 6 * time.Hour, PredecessorsId: []int{}, SuccessorsId: []int{6, 2, 3, 5}, Start: time.Time{}, Finish: time.Time{}},
}

func TestActivitiesToMap(t *testing.T) {
	m := ActivitiesToMap(activities)
	expected := map[int]*activity.Activity{
		1: activities[0],
		2: activities[1],
		3: activities[2],
		4: activities[3],
		5: activities[4],
		6: activities[5],
		7: activities[6],
	}
	for k := range m {
		if m[k] != expected[k] {
			t.Errorf("got %+v, want %+v", m[k], expected[k])
		}
	}
}

func TestActivitiesToGraph(t *testing.T) {
	g, err := ActivitiesToGraph(activities)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(g.String())
	/*
		expected := ""
		if g.String() != expected {
			t.Errorf("got %s, want %s", g.String(), expected)
		}
	*/
}

func TestFlat(t *testing.T) {
	a := []int{1, 2, 3, 4}
	aflat := Flat(a)

	if "1,2,3,4" != aflat {
		t.Error("Unexpected error")
	}
}

func TestUnflat(t *testing.T) {
	s := "1,2,3,4"
	sunflat, err := Unflat(s)
	if err != nil {
		t.Error(err)
	}
	want := []int{1, 2, 3, 4}
	if slices.Compare(sunflat, want) != 0 {
		t.Errorf("Error, want %v, got %v", want, sunflat)
	}
}
