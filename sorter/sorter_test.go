package sorter

import (
	"slices"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/util"
)

var now = time.Now()
var a1 = &activity.Activity{Id: 1, Description: "Cook eggs", Duration: 10 * time.Minute, PredecessorsId: []int{2}, SuccessorsId: []int{3}, Start: now.Add(time.Hour), Finish: now.Add(time.Hour), Cost: 10}
var a2 = &activity.Activity{Id: 2, Description: "Buy eggs", Duration: 1 * time.Hour, PredecessorsId: []int{}, SuccessorsId: []int{1}, Start: now, Finish: now.Add(time.Minute), Cost: 100}
var a3 = &activity.Activity{Id: 3, Description: "Eat eggs", Duration: 20 * time.Minute, PredecessorsId: []int{1}, SuccessorsId: []int{}, Start: now.Add(time.Minute), Finish: now.Add(time.Minute * 2), Cost: 30}
var activities = []*activity.Activity{a1, a2, a3}
var activitiesMap = util.ActivitiesToMap(activities)

func TestSortActivitiesByDeps(t *testing.T) {
	sortedOrder := SortActivitiesByDeps(activitiesMap)
	sorted := []int{2, 1, 3}
	if slices.Compare(sortedOrder, sorted) != 0 {
		t.Errorf("sorted order wrong, want %v, got %v\n", sorted, sortedOrder)
	}
}

func TestSortActivitiesById(t *testing.T) {
	acts := activities
	SortActivitiesById(acts)
	if acts[0] != a1 || acts[1] != a2 || acts[2] != a3 {
		t.Error("wrong order")
	}
}

func TestSortActivitiesByDuration(t *testing.T) {
	acts := activities
	SortActivitiesByDuration(acts)
	if acts[0] != a1 || acts[1] != a3 || acts[2] != a2 {
		t.Error("wrong order")
	}
}

func TestSortActivitiesByStart(t *testing.T) {
	acts := activities
	SortActivitiesByStart(acts)
	if acts[0] != a2 || acts[1] != a3 || acts[2] != a1 {
		t.Error("wrong order")
	}
}

func TestSortActivitiesByFinish(t *testing.T) {
	acts := activities
	SortActivitiesByFinish(acts)
	if acts[0] != a2 || acts[1] != a3 || acts[2] != a1 {
		t.Error("wrong order")
	}
}

func TestSortActivitiesByCost(t *testing.T) {
	acts := activities
	SortActivitiesByCost(acts)
	if acts[0] != a1 || acts[1] != a3 || acts[2] != a2 {
		t.Error("wrong order")
	}
}
