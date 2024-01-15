package main

import (
	"fmt"
	"log"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/graph"
	"github.com/vanillaiice/verano/parser/pxlsx"
	"github.com/vanillaiice/verano/project/timeline"
	"github.com/vanillaiice/verano/sorter"
	"github.com/vanillaiice/verano/util"
)

// function to avoid redundancy
func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// open 'activities.xlsx' workbook
	wb, err := xlsx.OpenFile("activities.xlsx")
	die(err)
	// get the first sheet
	sheet := wb.Sheets[0]
	// put the activities in a slice
	activities, err := pxlsx.XLSXToActivities(sheet)
	die(err)
	// convert the activities to a graph in order to sort them
	activitiesGraph, err := util.ActivitiesToGraph(activities)
	die(err)
	// get the order of the activities according to their dependencies
	order := sorter.SortActivitiesByDeps(activitiesGraph)
	// convert the activities to a map to sort them by order
	activitiesMap := util.ActivitiesToMap(activities)
	// define the start date of the project
	projectStartDate := time.Now()
	// update the start and finish times of all activities
	timeline.UpdateStartFinishTime(activitiesMap, order, projectStartDate)
	// init a graphviz graph
	gviz := graphviz.New()
	g, err := gviz.Graph()
	die(err)
	// sort the activities by order
	activitiesSorted, err := sorter.SortActivitiesByOrder(activitiesMap, order)
	die(err)
	// convert the slice of sorted activities to a map
	activitiesSortedMap := util.ActivitiesToMap(activitiesSorted)
	// draw the graph ftom the activities
	err = graph.Draw(g, activitiesSortedMap)
	die(err)
	// render the graph to a PNG file
	err = graph.GraphToImage(gviz, g, graphviz.PNG, "graph.png")
	die(err)

	// Bonus: write the sorted activities to another xlsx file
	wb2 := xlsx.NewFile()
	sheet2, err := wb2.AddSheet("activities")
	die(err)
	pxlsx.ActivitiesToXLSX(activitiesSorted, sheet2)
	wb2.Save("activities-sorted.xlsx")

	// Another bonus: print all the activities in order
	for i, o := range order {
		fmt.Printf("Activity %d: %v\n", i+1, activitiesMap[o])
	}
}
