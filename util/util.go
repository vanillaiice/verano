package util

import (
	"github.com/vanillaiice/verano/activity"
)

// ActivitiesToMap converts a slice of 'activities' into a map with activity ids as keys
// and pointers to activities as values. The resulting map provides quick access to activities by their IDs.
func ActivitiesToMap(activities []*activity.Activity) map[int]*activity.Activity {
	activitiesMap := make(map[int]*activity.Activity)
	for _, a := range activities {
		activitiesMap[a.Id] = a
	}
	return activitiesMap
}
