package timeline

import (
	"fmt"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

func TestUpdateStartFinishTime(t *testing.T) {
	activities := []*activity.Activity{
		{Id: 1, Description: "Cook eggs", Duration: 10 * time.Minute, PredecessorsId: []int{2}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
		{Id: 2, Description: "Buy eggs", Duration: 1 * time.Hour, PredecessorsId: []int{}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
		{Id: 3, Description: "Eat eggs", Duration: 20 * time.Minute, PredecessorsId: []int{1}, SuccessorsId: []int{}, Start: time.Time{}, Finish: time.Time{}},
		{Id: 4, Description: "Scream", Duration: 10 * time.Minute, PredecessorsId: []int{1}, SuccessorsId: []int{}, Start: time.Time{}, Finish: time.Time{}},
	}
	activitiesMap := util.ActivitiesToMap(activities)
	activitiesGraph, err := util.ActivitiesToGraph(activities)
	if err != nil {
		t.Error(err)
	}
	projectStartDate := time.Date(2024, time.January, int(time.Thursday), 10, 0, 0, 0, time.Local)

	sortedOrder := sorter.SortActivitiesByDeps(activitiesGraph)
	UpdateStartFinishTime(activitiesMap, sortedOrder, projectStartDate)

	// TODO: Check that times actually match

	fmt.Println("Sorted Activities:")
	alen := 1
	for _, id := range sortedOrder {
		fmt.Printf("Activity %d: Description=%v, Start=%v, Finish=%v\n", alen, activitiesMap[id].Description, activitiesMap[id].Start, activitiesMap[id].Finish)
		alen++
	}
	fmt.Printf("Total Project Duration is %.3f days\n", activitiesMap[sortedOrder[len(sortedOrder)-1]].Finish.Sub(projectStartDate).Hours()/24)
}
