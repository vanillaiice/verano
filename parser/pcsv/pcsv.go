package pcsv

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/util"
)

// ExportToDb populates the database with activities in csv format
func ExportToDb(b []byte, sqldb *sql.DB) (err error) {
	var activities []*activity.Activity
	r := bytes.NewReader(b)
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range records {
		if record[0] == "Id" {
			continue
		}
		activity, err := recordToActivity(record)
		if err != nil {
			return err
		}
		activities = append(activities, activity)
	}

	err = db.InsertActivities(sqldb, activities)
	if err != nil {
		return
	}
	return
}

// ActivitiesToCSV converts a slice of activities to csv format
func ActivitiesToCSV(activities []*activity.Activity, w io.Writer) (err error) {
	var records [][]string
	records = append(records, []string{"Id", "Description", "Duration", "Start", "Finish", "PredecessorsId", "SuccessorsId", "Cost"})
	for _, act := range activities {
		records = append(records, activityToRecord(act))
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	err = writer.WriteAll(records)
	if err != nil {
		return
	}

	return
}

func CSVToActivities(b []byte) (activities []*activity.Activity, err error) {
	r := bytes.NewReader(b)
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range records {
		if record[0] == "Id" {
			continue
		}
		activity, err := recordToActivity(record)
		if err != nil {
			return activities, err
		}
		activities = append(activities, activity)
	}
	return
}

// Convert record to *Activity pointer
func recordToActivity(record []string) (act *activity.Activity, err error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return
	}

	duration, err := time.ParseDuration(record[2])
	if err != nil {
		return
	}

	startTimeInt64, err := strconv.ParseInt(record[3], 10, 64)
	if err != nil {
		return
	}
	startTime := time.Unix(startTimeInt64, 0)

	finishTimeInt64, err := strconv.ParseInt(record[4], 10, 64)
	if err != nil {
		return
	}
	finishTime := time.Unix(finishTimeInt64, 0)

	predecessors, err := util.Unflat(record[5])
	if err != nil {
		return
	}

	successors, err := util.Unflat(record[6])
	if err != nil {
		return
	}

	cost, err := strconv.ParseFloat(record[7], 64)
	if err != nil {
		return
	}

	act = &activity.Activity{
		Id:             id,
		Description:    record[1],
		Duration:       duration,
		Start:          startTime,
		Finish:         finishTime,
		PredecessorsId: predecessors,
		SuccessorsId:   successors,
		Cost:           cost,
	}

	return act, nil
}

// Convert Activity struct to slice of strings
func activityToRecord(act *activity.Activity) []string {
	return []string{
		fmt.Sprint(act.Id),
		act.Description,
		act.Duration.String(),
		fmt.Sprint(act.Start.Unix()),
		fmt.Sprint(act.Finish.Unix()),
		util.Flat(act.PredecessorsId),
		util.Flat(act.SuccessorsId),
		fmt.Sprint(act.Cost),
	}
}
