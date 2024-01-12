package db

import (
	"database/sql"

	"github.com/vanillaiice/verano/activity"
	_ "modernc.org/sqlite"
)

// A DB stores a pointer to a sqlite database connection.
type DB struct {
	DB *sql.DB
}

// New creates a new instance of the DB type by initializing and opening a database.
// located at the specified 'path'. It returns a pointer to the created DB and an error, if any.
// The returned DB is ready for use, and the associated database file is opened.
// It should be noted that the DB connection should be closed after use.
func New(path string) (*DB, error) {
	var (
		db  DB
		err error
	)
	db.DB, err = Open(path)
	if err != nil {
		return &db, err
	}
	return &db, nil
}

// InsertActivity inserts the provided activity into the database.
func (db *DB) InsertActivity(act *activity.Activity) (err error) {
	err = InsertActivity(db.DB, act)
	if err != nil {
		return
	}
	return
}

// InsertActivity inserts the provided activities into the database.
func (db *DB) InsertActivities(activities []*activity.Activity) (err error) {
	err = InsertActivities(db.DB, activities)
	if err != nil {
		return
	}
	return
}

// GetActivityById retrieves the activity with the specified id from the database.
func (db *DB) GetActivityById(id int) (act *activity.Activity, err error) {
	act, err = GetActivityById(db.DB, id)
	if err != nil {
		return
	}
	return
}

// GetActivityById retrieves the activities with the specified ids from the database.
func (db *DB) GetActivitiesById(ids []int) (activities []*activity.Activity, err error) {
	activities, err = GetActivitiesById(db.DB, ids)
	if err != nil {
		return
	}
	return
}

// GetAllActivities retrieves all activities from the database.
// It returns a slice of pointers to activities.
func (db *DB) GetAllActivities() (activities []*activity.Activity, err error) {
	activities, err = GetAllActivities(db.DB)
	if err != nil {
		return
	}
	return
}

// GetAllActivitiesMap retrieves all activities from the database.
// and returns them as a map with activity ids as keys and pointers to activities as values.
func (db *DB) GetAllActivitiesMap() (activitiesMap map[int]*activity.Activity, err error) {
	activitiesMap, err = GetAllActivitiesMap(db.DB)
	if err != nil {
		return
	}
	return
}

// UpdateActivityById updates the activity with the specified id in the database
// using the information provided in the activity.
func (db *DB) UpdateActivityById(act *activity.Activity, id int) (n int64, err error) {
	n, err = UpdateActivityById(db.DB, act, id)
	if err != nil {
		return
	}
	return
}

// UpdatePredecessorsById updates the predecessors of the activity with the specified id in the database.
func (db *DB) UpdatePredecessorsById(id int, predecessorsId []int) (n int64, err error) {
	n, err = UpdatePredecessorsById(db.DB, id, predecessorsId)
	if err != nil {
		return
	}
	return
}

// UpdateSuccessorsById updates the successors of the activity with the specified id in the database.
func (db *DB) UpdateSuccessorsById(id int, successorsId []int) (n int64, err error) {
	n, err = UpdateSuccessorsById(db.DB, id, successorsId)
	if err != nil {
		return
	}
	return
}

// DeleteActivityById deletes the activity with the specified id from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivityById(id int) (n int64, err error) {
	n, err = DeleteActivityById(db.DB, id)
	if err != nil {
		return
	}
	return
}

// DeleteActivityById deletes the activities with the specified ids from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivitiesById(ids []int) (n int64, err error) {
	n, err = DeleteActivitiesById(db.DB, ids)
	if err != nil {
		return
	}
	return
}
