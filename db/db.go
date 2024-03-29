package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/util"
)

// TableName is the name of the table in the sqlite database.
const TableName = "activities"

// DuplicateInsertPolicy defines the policy for handling duplicate inserts in a database.
type DuplicateInsertPolicy int

// Enumeration of available duplicate insert policies.
const (
	None    DuplicateInsertPolicy = 0 // Do nothing
	Ignore  DuplicateInsertPolicy = 1 // Ignore duplicate inserts
	Replace DuplicateInsertPolicy = 2 // Replace duplicate inserts
)

func open(path string) (sqldb *sql.DB, err error) {
	sqldb, err = sql.Open("sqlite", path)
	if err != nil {
		return
	}
	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY, description TEXT, duration REAL, predecessorsId TEXT, successorsId TEXT, start INTEGER, finish INTEGER, cost REAL)", TableName)
	_, err = execStmt(sqldb, stmt)
	return
}

func insertActivity(sqldb *sql.DB, act *activity.Activity, duplicateInsertPolicy DuplicateInsertPolicy) (n int64, err error) {
	stmt := "INSERT "
	switch duplicateInsertPolicy {
	case Ignore:
		stmt += "or IGNORE "
	case Replace:
		stmt += "or REPLACE "
	}

	stmt += fmt.Sprintf(
		"INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(%d, %q, %.6f, %q, %q, %d, %d, %.6f)",
		TableName,
		act.Id,
		act.Description,
		act.Duration.Seconds(),
		util.Flat(act.PredecessorsId),
		util.Flat(act.SuccessorsId),
		act.Start.Unix(),
		act.Finish.Unix(),
		act.Cost,
	)

	return execStmt(sqldb, stmt)
}

func insertActivities(sqldb *sql.DB, activities []*activity.Activity, duplicateInsertPolicy DuplicateInsertPolicy) (err error) {
	s := "INSERT "
	switch duplicateInsertPolicy {
	case Ignore:
		s += "or IGNORE "
	case Replace:
		s += "or REPLACE "
	}
	s += fmt.Sprintf("INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", TableName)

	stmt, err := sqldb.Prepare(s)
	if err != nil {
		return
	}
	defer stmt.Close()

	for _, a := range activities {
		_, err = stmt.Exec(
			a.Id,
			a.Description,
			a.Duration.Seconds(),
			util.Flat(a.PredecessorsId),
			util.Flat(a.SuccessorsId),
			a.Start.Unix(),
			a.Finish.Unix(),
			a.Cost,
		)
		if err != nil {
			return
		}
	}

	return
}

func getActivity(sqldb *sql.DB, id int) (act *activity.Activity, err error) {
	stmt, err := sqldb.Prepare(fmt.Sprintf("SELECT description, duration, predecessorsId, successorsId, start, finish, cost FROM %s WHERE id = ?", TableName))
	if err != nil {
		return
	}
	defer stmt.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	err = stmt.QueryRow(id).Scan(&description, &duration, &predecessorsId, &successorsId, &start, &finish, &cost)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	pIds, err := util.Unflat(predecessorsId)
	if err != nil {
		return
	}
	sIds, err := util.Unflat(successorsId)
	if err != nil {
		return
	}

	act = &activity.Activity{
		Id:             id,
		Description:    description,
		Duration:       time.Duration(duration * float64(time.Second)),
		PredecessorsId: pIds,
		SuccessorsId:   sIds,
		Start:          time.Unix(start, 0),
		Finish:         time.Unix(finish, 0),
		Cost:           cost,
	}

	return
}

func getActivities(sqldb *sql.DB, ids []int) (activities []*activity.Activity, err error) {
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", TableName, util.Flat(ids))
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &description, &duration, &predecessorsId, &successorsId, &start, &finish, &cost)
		if err != nil {
			return
		}

		pIds, err := util.Unflat(predecessorsId)
		if err != nil {
			return activities, err
		}
		sIds, err := util.Unflat(successorsId)
		if err != nil {
			return activities, err
		}

		activities = append(activities, &activity.Activity{
			Id:             id,
			Description:    description,
			Duration:       time.Duration(duration * float64(time.Second)),
			PredecessorsId: pIds,
			SuccessorsId:   sIds,
			Start:          time.Unix(start, 0),
			Finish:         time.Unix(finish, 0),
			Cost:           cost,
		})
	}

	return
}

func getActivitiesAll(sqldb *sql.DB) (activities []*activity.Activity, err error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", TableName)
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &description, &duration, &predecessorsId, &successorsId, &start, &finish, &cost)
		if err != nil {
			return activities, err
		}

		pIds, err := util.Unflat(predecessorsId)
		if err != nil {
			return activities, err
		}
		sIds, err := util.Unflat(successorsId)
		if err != nil {
			return activities, err
		}

		activities = append(
			activities,
			&activity.Activity{
				Id:             id,
				Description:    description,
				Duration:       time.Duration(duration * float64(time.Second)),
				PredecessorsId: pIds,
				SuccessorsId:   sIds,
				Start:          time.Unix(start, 0),
				Finish:         time.Unix(finish, 0),
				Cost:           cost,
			})
	}

	return
}

func getActivitiesAllMap(sqldb *sql.DB) (activitiesMap map[int]*activity.Activity, err error) {
	stmt := fmt.Sprintf("SELECT * FROM %s", TableName)
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &description, &duration, &predecessorsId, &successorsId, &start, &finish)
		if err != nil {
			return activitiesMap, err
		}

		pIds, err := util.Unflat(predecessorsId)
		if err != nil {
			return activitiesMap, err
		}
		sIds, err := util.Unflat(successorsId)
		if err != nil {
			return activitiesMap, err
		}

		activitiesMap[id] = &activity.Activity{
			Id:             id,
			Description:    description,
			Duration:       time.Duration(duration * float64(time.Second)),
			PredecessorsId: pIds,
			SuccessorsId:   sIds,
			Start:          time.Unix(start, 0),
			Finish:         time.Unix(finish, 0),
			Cost:           cost,
		}
	}

	return
}

func updateActivity(sqldb *sql.DB, act *activity.Activity, id int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET description = %q, duration = %.6f, predecessorsId=%q, successorsId=%q, start = %d, finish = %d, cost = %.6f WHERE id = %d",
		TableName,
		act.Description,
		act.Duration.Seconds(),
		util.Flat(act.PredecessorsId),
		util.Flat(act.SuccessorsId),
		act.Start.Unix(),
		act.Finish.Unix(),
		act.Cost,
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateId(sqldb *sql.DB, oldId, newId int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET id=%d WHERE id=%d",
		TableName,
		newId,
		oldId,
	)
	return execStmt(sqldb, stmt)
}

func updateDescription(sqldb *sql.DB, id int, newDescription string) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET description=%q WHERE id=%d",
		TableName,
		newDescription,
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateDuration(sqldb *sql.DB, id int, newDuration time.Duration) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET duration=%.6f WHERE id=%d",
		TableName,
		newDuration.Seconds(),
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateStart(sqldb *sql.DB, id int, newStart time.Time) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET start=%d WHERE id=%d",
		TableName,
		newStart.Unix(),
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateFinish(sqldb *sql.DB, id int, newFinish time.Time) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET finish=%d WHERE id=%d",
		TableName,
		newFinish.Unix(),
		id,
	)
	return execStmt(sqldb, stmt)
}

func updatePredecessors(sqldb *sql.DB, id int, newPredecessorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET predecessorsId=%q WHERE id = %d",
		TableName,
		util.Flat(newPredecessorsId),
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateSuccessors(sqldb *sql.DB, id int, newSuccessorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET successorsId=%q WHERE id = %d",
		TableName,
		util.Flat(newSuccessorsId),
		id,
	)
	return execStmt(sqldb, stmt)
}

func updateCost(sqldb *sql.DB, id int, newCost float64) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET cost=%.6f WHERE id=%d",
		TableName,
		newCost,
		id,
	)
	return execStmt(sqldb, stmt)
}

func deleteActivity(sqldb *sql.DB, id int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = %d", TableName, id)
	return execStmt(sqldb, stmt)
}

func deleteActivities(sqldb *sql.DB, ids []int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", TableName, util.Flat(ids))
	return execStmt(sqldb, stmt)
}

func execStmt(sqldb *sql.DB, stmt string) (n int64, err error) {
	res, err := sqldb.Exec(stmt)
	if err != nil {
		return
	}
	return res.RowsAffected()
}
