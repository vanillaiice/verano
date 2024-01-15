package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/heimdalr/dag"
	"github.com/vanillaiice/verano/activity"
)

// ActivitiesToMap converts a slice of 'activities' into a map with activity ids as keys
// and pointers to activities as values. The resulting map provides quick access to activities by their IDs.
func ActivitiesToMap(activities []*activity.Activity) (activitiesMap map[int]*activity.Activity) {
	activitiesMap = make(map[int]*activity.Activity)
	for _, a := range activities {
		activitiesMap[a.Id] = a
	}
	return
}

// ActivitiesToGraph converts a slice of 'activities' into a directed acyclic graph (DAG).
// It creates a new DAG, adds vertices for each activity, and establishes edges based on the successors' IDs.
func ActivitiesToGraph(activities []*activity.Activity) (g *dag.DAG, err error) {
	g = dag.NewDAG()
	for _, act := range activities {
		err = g.AddVertexByID(fmt.Sprint(act.Id), act)
		if err != nil {
			return g, err
		}
	}
	for _, act := range activities {
		for _, successorId := range act.SuccessorsId {
			err = g.AddEdge(fmt.Sprint(act.Id), fmt.Sprint(successorId))
			if err != nil {
				return g, err
			}
		}
	}
	return
}

// Flat converts a slice of integers into a comma-separated string.
func Flat(vals []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vals)), ","), "[]")
}

// Unflat converts a comma-separated string into a slice of integers.
func Unflat(s string) (intVals []int, err error) {
	if s == "" {
		return intVals, nil
	}
	sVals := strings.Split(s, ",")
	for _, s := range sVals {
		val, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		intVals = append(intVals, val)
	}
	return
}
