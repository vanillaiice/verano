package pcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/util"
)

var recordHeader = []string{"Id", "Description", "Duration", "Start", "Finish", "PredecessorsId", "SuccessorsId", "Cost"}

// ExportToDb populates the database with activities in csv format.
func ExportToDb(sqldb *db.DB, reader io.Reader, duplicateInsertPolicy db.DuplicateInsertPolicy) (err error) {
	activities, err := CSVToActivities(reader)
	if err != nil {
		return
	}
	return sqldb.InsertActivities(activities, duplicateInsertPolicy)
}

// ActivitiesToCSV converts a slice of activities to csv format.
func ActivitiesToCSV(activities []*activity.Activity, w io.Writer) (err error) {
	var records [][]string
	records = append(records, recordHeader)
	for _, act := range activities {
		records = append(records, activityToRecord(act))
	}
	writer := csv.NewWriter(w)
	defer writer.Flush()
	return writer.WriteAll(records)
}

// CSVToActivities converts csv format to a slice of activities.
func CSVToActivities(reader io.Reader) (activities []*activity.Activity, err error) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
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

// recordToActivity converts a record to an Activity pointer.
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

// activityToRecord converts an Activity struct to a slice of strings.
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
