package sorter

import (
	"errors"
	"sort"

	"github.com/vanillaiice/verano/activity"
)

// SortActivitiesByDeps performs a topological sort on the activities in the provided 'activitiesMap'
// based on their dependencies and returns a slice representing the sorted order.
// It uses a depth-first search (DFS) approach to traverse the activities and create the topological order.
// The resulting order ensures that activities with dependencies come before their dependent activities.
// It should be noted that the relationship between the activities
// is assumed to be start to finish.
func SortActivitiesByDeps(activitiesMap map[int]*activity.Activity) []int {
	visited := make(map[int]bool)
	order := []int{}

	var dfs func(int)
	dfs = func(id int) {
		visited[id] = true
		for _, successorID := range activitiesMap[id].SuccessorsId {
			if !visited[successorID] {
				dfs(successorID)
			}
		}
		order = append(order, id)
	}

	for id := range activitiesMap {
		if !visited[id] {
			dfs(id)
		}
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	return order
}

// SortActivitiesByOrder sorts activities from the provided 'activities' map based on the specified 'order'.
// It returns a slice of pointers to activities in the sorted order. An error is returned if the length of
// 'activities' does not match the length of 'order'.
func SortActivitiesByOrder(activities map[int]*activity.Activity, order []int) ([]*activity.Activity, error) {
	a := []*activity.Activity{}
	if len(activities) != len(order) {
		return a, errors.New("length of activities do not match length of order")
	}
	for _, id := range order {
		a = append(a, activities[id])
	}
	return a, nil
}

// SortActivitiesById sorts the provided activities
func SortActivitiesById(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Id < activities[j].Id
	})
}

// SortActivitiesById sorts the provided slice of 'activities' based on their IDs in ascending order.
func SortActivitiesByDescription(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Description < activities[j].Description
	})
}

// SortActivitiesByDescription sorts the provided slice of 'activities' based on their descriptions in ascending order.
func SortActivitiesByDuration(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Duration < activities[j].Duration
	})
}

// SortActivitiesByDuration sorts the provided slice of 'activities' based on their durations in ascending order.
func SortActivitiesByStart(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Start.Unix() < activities[j].Start.Unix()
	})
}

// SortActivitiesByFinish sorts the provided slice of 'activities' based on their finish times in ascending order.
func SortActivitiesByFinish(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Finish.Unix() < activities[j].Finish.Unix()
	})
}

// SortActivitiesByCost sorts the provided slice of 'activities' based on their cost in ascending order
func SortActivitiesByCost(activities []*activity.Activity) {
	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Cost < activities[j].Cost
	})
}
