package pjson

import (
	"database/sql"
	"encoding/json"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
)

// ExportToDb populates the database with activities in json format
func ExportToDb(j []byte, sqldb *sql.DB) error {
	var activities []*activity.Activity
	err := json.Unmarshal(j, &activities)
	if err != nil {
		return err
	}

	err = db.InsertActivities(sqldb, activities)
	if err != nil {
		return err
	}
	return nil
}

// ActivitiesToJson converts a slice of activities to json format
func ActivitiesToJson(activities []*activity.Activity) ([]byte, error) {
	j, err := json.MarshalIndent(activities, "", "\t")
	if err != nil {
		return j, err
	}
	return j, nil
}
