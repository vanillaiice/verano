package text

import (
	"fmt"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/gantt"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

var activities = []*activity.Activity{
	{Id: 1, Description: "Cook eggs", Duration: 5 * time.Minute, PredecessorsId: []int{2, 4, 5}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 2, Description: "Buy eggs", Duration: 10 * time.Minute, PredecessorsId: []int{}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 3, Description: "Eat eggs", Duration: 5 * time.Minute, PredecessorsId: []int{1, 6}, SuccessorsId: []int{}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 4, Description: "Buy pan", Duration: 10 * time.Minute, PredecessorsId: []int{}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 5, Description: "Buy salt", Duration: 6 * time.Minute, PredecessorsId: []int{}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 6, Description: "Tip landlord", Duration: 60 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 7, Description: "Get money", Duration: 6 * time.Hour, PredecessorsId: []int{}, SuccessorsId: []int{6}, Start: time.Time{}, Finish: time.Time{}},
}
var activitiesMap = util.ActivitiesToMap(activities)
var order = sorter.SortActivitiesByDeps(activitiesMap)
var activitiesSorted, _ = sorter.SortActivitiesByOrder(activitiesMap, order)

func TestDraw(t *testing.T) {
	activitiesSortedMap := util.ActivitiesToMap(activitiesSorted)
	gantt.UpdateStartFinishTime(activitiesSortedMap, order, time.Now())
	tree, err := MakeTree(activitiesSortedMap, order[len(order)-1])
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tree)
}
