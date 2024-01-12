package gantt

import (
	"time"

	"github.com/vanillaiice/verano/activity"
)

// UpdateStartFinishTime updates the start and finish times of activities in the provided 'activitiesMap'
// based on the order of activities sorted by their dependencies and the 'projectStartDate'.
// It iterates through the 'orderActivitiesSortedByDep' slice, representing activities sorted by their dependencies,
// and calculates the earliest finish time considering the finish times of their predecessors.
// The 'Start' and 'Finish' fields of each activity in 'activitiesMap' are then updated accordingly.
// This function modifies the 'activitiesMap' in-place.
func UpdateStartFinishTime(activitiesMap map[int]*activity.Activity, orderActivitiesSortedByDep []int, projectStartDate time.Time) {
	for _, id := range orderActivitiesSortedByDep {
		a := activitiesMap[id]
		minFinishTime := projectStartDate

		for _, predecessorId := range a.PredecessorsId {
			predecessorFinishTime := activitiesMap[predecessorId].Finish
			if predecessorFinishTime.After(minFinishTime) {
				minFinishTime = predecessorFinishTime
			}
		}

		a.Start = minFinishTime
		a.Finish = minFinishTime.Add(a.Duration)
		activitiesMap[id] = a
	}
}
