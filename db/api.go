package db

import (
	"database/sql"
	"time"

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
	db.DB, err = open(path)
	return &db, err
}

// Close closes the db connection
func Close(db *DB) error {
	return db.DB.Close()
}

// InsertActivity inserts the provided activity into the database.
func (db *DB) InsertActivity(act *activity.Activity, duplicateInsertPolicy DuplicateInsertPolicy) (n int64, err error) {
	return insertActivity(db.DB, act, duplicateInsertPolicy)
}

// InsertActivities inserts the provided activities into the database.
func (db *DB) InsertActivities(activities []*activity.Activity, duplicateInsertPolicy DuplicateInsertPolicy) (err error) {
	return insertActivities(db.DB, activities, duplicateInsertPolicy)
}

// GetActivity retrieves the activity with the specified id from the database.
func (db *DB) GetActivity(id int) (act *activity.Activity, err error) {
	return getActivity(db.DB, id)
}

// GetActivities retrieves the activities with the specified ids from the database.
func (db *DB) GetActivities(ids []int) (activities []*activity.Activity, err error) {
	return getActivities(db.DB, ids)
}

// GetActivitiesAll retrieves all activities from the database.
// It returns a slice of pointers to activities.
func (db *DB) GetActivitiesAll() (activities []*activity.Activity, err error) {
	return getActivitiesAll(db.DB)
}

// GetActivitiesAllMap retrieves all activities from the database,
// and returns them as a map with activity ids as keys and pointers to activities as values.
func (db *DB) GetActivitiesAllMap() (activitiesMap map[int]*activity.Activity, err error) {
	return getActivitiesAllMap(db.DB)
}

// UpdateActivity updates the activity with the specified id in the database
// using the information provided in the activity.
func (db *DB) UpdateActivity(act *activity.Activity, id int) (n int64, err error) {
	return updateActivity(db.DB, act, id)
}

// UpdateId updates the id of an activity with the specified id in the database
func (db *DB) UpdateId(oldId, newId int) (n int64, err error) {
	return updateId(db.DB, oldId, newId)
}

// UpdateDescription updates the description of an activity with the specified id in the database
func (db *DB) UpdateDescription(id int, newDescription string) (n int64, err error) {
	return updateDescription(db.DB, id, newDescription)
}

// UpdateDuration updates the duration of an activity with the specified id in the database
func (db *DB) UpdateDuration(id int, newDuration time.Duration) (n int64, err error) {
	return updateDuration(db.DB, id, newDuration)
}

// UpdateStart updates the start time of an activity with the specified id in the database
func (db *DB) UpdateStart(id int, newStart time.Time) (n int64, err error) {
	return updateStart(db.DB, id, newStart)
}

// UpdateFinish updates the finish time of an activity with the specified id in the database
func (db *DB) UpdateFinish(id int, newFinish time.Time) (n int64, err error) {
	return updateFinish(db.DB, id, newFinish)
}

// UpdateSuccessors updates the successors of the activity with the specified id in the database.
func (db *DB) UpdateSuccessors(id int, successorsId []int) (n int64, err error) {
	return updateSuccessors(db.DB, id, successorsId)
}

// UpdateCost updates the cost of an activity with the specified id in the database
func (db *DB) UpdateCost(id int, newCost float64) (n int64, err error) {
	return updateCost(db.DB, id, newCost)
}

// UpdatePredecessors updates the predecessors of the activity with the specified id in the database.
func (db *DB) UpdatePredecessors(id int, predecessorsId []int) (n int64, err error) {
	return updatePredecessors(db.DB, id, predecessorsId)
}

// DeleteActivity deletes the activity with the specified id from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivity(id int) (n int64, err error) {
	return deleteActivity(db.DB, id)
}

// DeleteActivities deletes the activities with the specified ids from the database.
// It returns the number of affected rows and an error if the deletion operation encounters any issues.
func (db *DB) DeleteActivities(ids []int) (n int64, err error) {
	return deleteActivities(db.DB, ids)
}
