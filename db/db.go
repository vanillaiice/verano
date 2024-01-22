package db

import (
	"database/sql"
	"time"

	"github.com/vanillaiice/verano/activity"
	_ "modernc.org/sqlite"
)

// TableName is the name of the table in the sqlite database.
const TableName = "activities"

// DuplicateInsertPolicy defines the policy for handling duplicate inserts in a database.
type DuplicateInsertPolicy int

// Enumeration of available duplicate insert policies.
const (
	Ignore  DuplicateInsertPolicy = 0 // Ignore duplicate inserts
	Replace DuplicateInsertPolicy = 1 // Replace duplicate inserts
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
	db.DB, err = open(path)
	return &db, err
}

// Close closes the db connection
func Close(db *DB) error {
	return db.DB.Close()
}

// InsertActivity inserts the provided activity into the database.
func (db *DB) InsertActivity(act *activity.Activity, duplicateInsertPolicy ...DuplicateInsertPolicy) (n int64, err error) {
	n, err = insertActivity(db.DB, act, duplicateInsertPolicy...)
	return
}

// InsertActivities inserts the provided activities into the database.
func (db *DB) InsertActivities(activities []*activity.Activity, duplicateInsertPolicy ...DuplicateInsertPolicy) (err error) {
	err = insertActivities(db.DB, activities, duplicateInsertPolicy...)
	return
}

// GetActivity retrieves the activity with the specified id from the database.
func (db *DB) GetActivity(id int) (act *activity.Activity, err error) {
	act, err = getActivity(db.DB, id)
	return
}

// GetActivities retrieves the activities with the specified ids from the database.
func (db *DB) GetActivities(ids []int) (activities []*activity.Activity, err error) {
	activities, err = getActivities(db.DB, ids)
	return
}

// GetActivitiesAll retrieves all activities from the database.
// It returns a slice of pointers to activities.
func (db *DB) GetActivitiesAll() (activities []*activity.Activity, err error) {
	activities, err = getActivitiesAll(db.DB)
	return
}

// GetActivitiesAllMap retrieves all activities from the database,
// and returns them as a map with activity ids as keys and pointers to activities as values.
func (db *DB) GetActivitiesAllMap() (activitiesMap map[int]*activity.Activity, err error) {
	activitiesMap, err = getActivitiesAllMap(db.DB)
	return
}

// UpdateActivity updates the activity with the specified id in the database
// using the information provided in the activity.
func (db *DB) UpdateActivity(act *activity.Activity, id int) (n int64, err error) {
	n, err = updateActivity(db.DB, act, id)
	return
}

// UpdateId updates the id of an activity with the specified id in the database
func (db *DB) UpdateId(oldId, newId int) (n int64, err error) {
	n, err = updateId(db.DB, oldId, newId)
	return
}

// UpdateDescription updates the description of an activity with the specified id in the database
func (db *DB) UpdateDescription(id int, newDescription string) (n int64, err error) {
	n, err = updateDescription(db.DB, id, newDescription)
	return
}

// UpdateDuration updates the duration of an activity with the specified id in the database
func (db *DB) UpdateDuration(id int, newDuration time.Duration) (n int64, err error) {
	n, err = updateDuration(db.DB, id, newDuration)
	return
}

// UpdateStart updates the start time of an activity with the specified id in the database
func (db *DB) UpdateStart(id int, newStart time.Time) (n int64, err error) {
	n, err = updateStart(db.DB, id, newStart)
	return
}

// UpdateFinish updates the finish time of an activity with the specified id in the database
func (db *DB) UpdateFinish(id int, newFinish time.Time) (n int64, err error) {
	n, err = updateFinish(db.DB, id, newFinish)
	return
}

// UpdateSuccessors updates the successors of the activity with the specified id in the database.
func (db *DB) UpdateSuccessors(id int, successorsId []int) (n int64, err error) {
	n, err = updateSuccessors(db.DB, id, successorsId)
	return
}

// UpdateCost updates the cost of an activity with the specified id in the database
func (db *DB) UpdateCost(id int, newCost float64) (n int64, err error) {
	n, err = updateCost(db.DB, id, newCost)
	return
}

// UpdatePredecessors updates the predecessors of the activity with the specified id in the database.
func (db *DB) UpdatePredecessors(id int, predecessorsId []int) (n int64, err error) {
	n, err = updatePredecessors(db.DB, id, predecessorsId)
	return
}

// DeleteActivity deletes the activity with the specified id from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivity(id int) (n int64, err error) {
	n, err = deleteActivity(db.DB, id)
	return
}

// DeleteActivities deletes the activities with the specified ids from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivities(ids []int) (n int64, err error) {
	n, err = deleteActivities(db.DB, ids)
	return
}
