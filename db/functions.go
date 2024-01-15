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
	_, err = sqldb.Exec(stmt)
	if err != nil {
		return
	}

	return
}

func InsertActivity(sqldb *sql.DB, act *activity.Activity) (err error) {
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
	_, err = sqldb.Exec(stmt)
	if err != nil {
		return
	}
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

func GetActivityById(sqldb *sql.DB, id int) (act *activity.Activity, err error) {
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

func GetActivitiesById(sqldb *sql.DB, ids []int) (activities []*activity.Activity, err error) {
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

func GetAllActivities(sqldb *sql.DB) (activities []*activity.Activity, err error) {
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

func GetAllActivitiesMap(sqldb *sql.DB) (activitiesMap map[int]*activity.Activity, err error) {
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

func UpdateActivityById(sqldb *sql.DB, act *activity.Activity, id int) (n int64, err error) {
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

func UpdatePredecessorsById(sqldb *sql.DB, id int, predecessorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf(
		"UPDATE %s SET predecessorsId=%q WHERE id = %d",
		TableName,
		util.Flat(predecessorsId),
		id,
	)
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

func UpdateSuccessorsById(sqldb *sql.DB, id int, successorsId []int) (n int64, err error) {
	stmt := fmt.Sprintf("UPDATE %s SET successorsId=%q WHERE id = %d",
		TableName,
		util.Flat(successorsId),
		id,
	)
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

func DeleteActivityById(sqldb *sql.DB, id int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = %d", TableName, id)
	res, err := sqldb.Exec(stmt)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}

func DeleteActivitiesById(sqldb *sql.DB, ids []int) (n int64, err error) {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", TableName, util.Flat(ids))
	res, err := sqldb.Exec(stmt)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	n, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}

// func UpdateXXXById(sqldb *sqldb, activityId int, newValue XXX) (n int64, err error)
