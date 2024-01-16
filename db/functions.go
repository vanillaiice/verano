package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/util"
)

// TableName is the name of the table in the sqlite database
const TableName = "activities"

func Open(path string) (sqldb *sql.DB, err error) {
	sqldb, err = sql.Open("sqlite", path)
	if err != nil {
		return
	}
	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY, description TEXT, duration REAL, predecessorsId TEXT, successorsId TEXT, start INTEGER, finish INTEGER, cost REAL)", TableName)
	_, err = execStmt(sqldb, stmt)
	return
}

func InsertActivity(sqldb *sql.DB, act *activity.Activity) (n int64, err error) {
	stmt := fmt.Sprintf(
		"INSERT INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(%d, %q, %.6f, %q, %q, %d, %d, %.6f)",
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
	n, err = execStmt(sqldb, stmt)
	return
}

func InsertActivities(sqldb *sql.DB, activities []*activity.Activity) (err error) {
	stmt, err := sqldb.Prepare(fmt.Sprintf(
		"INSERT INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", TableName))
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

func GetActivity(sqldb *sql.DB, id int) (act *activity.Activity, err error) {
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

func GetActivities(sqldb *sql.DB, ids []int) (activities []*activity.Activity, err error) {
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

func GetActivitiesAll(sqldb *sql.DB) (activities []*activity.Activity, err error) {
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

func GetActivitiesAllMap(sqldb *sql.DB) (activitiesMap map[int]*activity.Activity, err error) {
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

func UpdateActivity(sqldb *sql.DB, act *activity.Activity, id int) (n int64, err error) {
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
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateId(sqldb *sql.DB, oldId, newId int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET id=%d WHERE id=%d",
		TableName,
		newId,
		oldId,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateDescription(sqldb *sql.DB, id int, newDescription string) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET description=%s WHERE id=%d",
		TableName,
		newDescription,
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateDuration(sqldb *sql.DB, id int, newDuration time.Duration) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET duration=%.6f WHERE id=%d",
		TableName,
		newDuration.Seconds(),
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateStart(sqldb *sql.DB, id int, newStart time.Time) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET start=%d WHERE id=%d",
		TableName,
		newStart.Unix(),
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateFinish(sqldb *sql.DB, id int, newFinish time.Time) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET finish=%d WHERE id=%d",
		TableName,
		newFinish.Unix(),
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdatePredecessors(sqldb *sql.DB, id int, newPredecessorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET predecessorsId=%q WHERE id = %d",
		TableName,
		util.Flat(newPredecessorsId),
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateSuccessors(sqldb *sql.DB, id int, newSuccessorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET successorsId=%q WHERE id = %d",
		TableName,
		util.Flat(newSuccessorsId),
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func UpdateCost(sqldb *sql.DB, id int, newCost float64) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET cost=%.6f WHERE id=%d",
		TableName,
		newCost,
		id,
	)
	n, err = execStmt(sqldb, stmt)
	return
}

func DeleteActivity(sqldb *sql.DB, id int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = %d", TableName, id)
	n, err = execStmt(sqldb, stmt)
	return
}

func DeleteActivities(sqldb *sql.DB, ids []int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", TableName, util.Flat(ids))
	n, err = execStmt(sqldb, stmt)
	return
}

func execStmt(sqldb *sql.DB, stmt string) (n int64, err error) {
	res, err := sqldb.Exec(stmt)
	if err != nil {
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}
