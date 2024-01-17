package pjson

import (
	"database/sql"
	"encoding/json"
	"io"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
)

// ExportToDb populates the database with activities in json format.
func ExportToDb(sqldb *sql.DB, reader io.Reader) (err error) {
	activities, err := JSONtoActivities(reader)
	if err != nil {
		return
	}
	err = db.InsertActivities(sqldb, activities)
	return
}

// ActivitiesToJSON converts a slice of activities to json format.
func ActivitiesToJSON(activities []*activity.Activity, writer io.Writer) (err error) {
	j, err := json.MarshalIndent(activities, "", "\t")
	if err != nil {
		return
	}
	_, err = writer.Write(j)
	return
}

// JSONtoActivities converts activities in json format to a slice of activities.
func JSONtoActivities(reader io.Reader) (activities []*activity.Activity, err error) {
	j, err := io.ReadAll(reader)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &activities)
	return
}
