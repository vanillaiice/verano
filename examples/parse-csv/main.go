package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/vanillaiice/verano/parser/pcsv"
	"github.com/vanillaiice/verano/project/timeline"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

// List of activities in CSV format
var scsv = `Id,Description,Duration,Start,Finish,PredecessorsId,SuccessorsId,Cost
1,Tip landlord,12h0s,-62135596800,-62135596800,2,3,1000000
2,Get money,6h0s,-62135596800,-62135596800,,1,0
3,Edge,12h0s,-62135596800,-62135596800,1,,1000
`

func main() {
	// Parse activities in CSV format to an activity slice
	r := bytes.NewReader([]byte(scsv))
	activities, err := pcsv.CSVToActivities(r)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the activity slice to a graph
	activitiesGraph, err := util.ActivitiesToGraph(activities)
	if err != nil {
		log.Fatal(err)
	}
	// Sort the activities by dependencies using topological sorting on the graph
	order := sorter.SortActivitiesByDeps(activitiesGraph)

	// Define the start of the project
	projectStartDate := time.Now()
	// convert the activity slice to a map with the activity Id as key
	// and the activity as the value
	activitiesMap := util.ActivitiesToMap(activities)
	// Update the Start and Finish time of the activities after getting sorted
	timeline.UpdateStartFinishTime(activitiesMap, order, projectStartDate)

	// Print activities in order
	for i, o := range order {
		fmt.Printf("Activity #%d: %v\n", i+1, activitiesMap[o])
	}
}
