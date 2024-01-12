package graph

import (
	"testing"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/gantt"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

var activities = []*activity.Activity{
	{Id: 1, Description: "Cook eggs", Duration: 5 * time.Minute, PredecessorsId: []int{2, 4, 5}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 2, Description: "Buy eggs", Duration: 10 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 3, Description: "Eat eggs", Duration: 5 * time.Minute, PredecessorsId: []int{1, 6}, SuccessorsId: []int{}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 4, Description: "Buy pan", Duration: 10 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 5, Description: "Buy salt", Duration: 6 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{1}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 6, Description: "Tip landlord", Duration: 60 * time.Minute, PredecessorsId: []int{7}, SuccessorsId: []int{3}, Start: time.Time{}, Finish: time.Time{}},
	{Id: 7, Description: "Get money", Duration: 6 * time.Hour, PredecessorsId: []int{}, SuccessorsId: []int{2, 4, 5, 6}, Start: time.Time{}, Finish: time.Time{}},
}
var activitiesMap = util.ActivitiesToMap(activities)
var activitiesGraph, _ = util.ActivitiesToGraph(activities)
var order = sorter.SortActivitiesByDeps(activitiesGraph)
var activitiesSorted, _ = sorter.SortActivitiesByOrder(activitiesMap, order)
var activitiesSortedMap = util.ActivitiesToMap(activitiesSorted)
var g = graphviz.New()

func TestDraw(t *testing.T) {
	gantt.UpdateStartFinishTime(activitiesSortedMap, order, time.Now())
	graph, err := g.Graph()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = graph.Close(); err != nil {
			return
		}
		g.Close()
	}()
	err = Draw(graph, activitiesSortedMap)
	if err != nil {
		t.Error(err)
	}
	err = GraphToImage(g, graph, graphviz.PNG, "graph.png")
	if err != nil {
		t.Error(err)
	}
}

/*
// FIXME: this is causing a segfault smh
func TestGraphToImage(t *testing.T) {
	graph, err := g.Graph()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = graph.Close(); err != nil {
			return
		}
		g.Close()
	}()
	err = GraphToImage(g, graph, graphviz.PNG, "graph.png")
	if err != nil {
		t.Error(err)
	}
}
*/
