package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vanillaiice/verano/activity"
)

// TableName is the name of the table in the sqlite database
const TableName = "activities"

func Open(path string) (*sql.DB, error) {
	var sqldb *sql.DB
	sqldb, err := sql.Open("sqlite", path)
	if err != nil {
		return sqldb, err
	}

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY, description TEXT, duration REAL, predecessorsId TEXT, successorsId TEXT, start INTEGER, finish INTEGER, cost REAL)", TableName)
	_, err = sqldb.Exec(stmt)
	if err != nil {
		return sqldb, err
	}

	return sqldb, nil
}

func InsertActivity(sqldb *sql.DB, act *activity.Activity) error {
	stmt := fmt.Sprintf("INSERT INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(%d, %q, %.6f, %q, %q, %d, %d, %.6f)",
		TableName,
		act.Id,
		act.Description,
		act.Duration.Seconds(),
		flat(act.PredecessorsId),
		flat(act.SuccessorsId),
		act.Start.Unix(),
		act.Finish.Unix(),
		act.Cost,
	)
	_, err := sqldb.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func InsertActivities(sqldb *sql.DB, activities []*activity.Activity) error {
	stmt, err := sqldb.Prepare(fmt.Sprintf("INSERT INTO %s(id, description, duration, predecessorsId, successorsId, start, finish, cost) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, a := range activities {
		_, err = stmt.Exec(
			a.Id,
			a.Description,
			a.Duration.Seconds(),
			flat(a.PredecessorsId),
			flat(a.SuccessorsId),
			a.Start.Unix(),
			a.Finish.Unix(),
			a.Cost,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetActivityById(sqldb *sql.DB, id int) (*activity.Activity, error) {
	var act *activity.Activity

	stmt, err := sqldb.Prepare(fmt.Sprintf("SELECT description, duration, predecessorsId, successorsId, start, finish, cost FROM %s WHERE id = ?", TableName))
	if err != nil {
		return act, err
	}
	defer stmt.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	err = stmt.QueryRow(id).Scan(&description, &duration, &predecessorsId, &successorsId, &start, &finish, &cost)
	if err != nil && err != sql.ErrNoRows {
		return act, err
	}

	pIds, err := unflat(predecessorsId)
	if err != nil {
		return act, err
	}
	sIds, err := unflat(successorsId)
	if err != nil {
		return act, err
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

	return act, nil
}

func GetActivitiesById(sqldb *sql.DB, ids []int) ([]*activity.Activity, error) {
	var activities []*activity.Activity

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", TableName, flat(ids))
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return activities, err
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

		pIds, err := unflat(predecessorsId)
		if err != nil {
			return activities, err
		}
		sIds, err := unflat(successorsId)
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

	return activities, nil
}

func GetAllActivities(sqldb *sql.DB) ([]*activity.Activity, error) {
	activities := []*activity.Activity{}
	stmt := fmt.Sprintf("SELECT * FROM %s", TableName)
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return activities, err
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

		pIds, err := unflat(predecessorsId)
		if err != nil {
			return activities, err
		}
		sIds, err := unflat(successorsId)
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

	return activities, nil
}

func GetAllActivitiesMap(sqldb *sql.DB) (map[int]*activity.Activity, error) {
	activities := make(map[int]*activity.Activity)
	stmt := fmt.Sprintf("SELECT * FROM %s", TableName)
	rows, err := sqldb.Query(stmt)
	if err != nil {
		return activities, err
	}
	defer rows.Close()

	var description, predecessorsId, successorsId string
	var duration, cost float64
	var start, finish int64
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &description, &duration, &predecessorsId, &successorsId, &start, &finish)
		if err != nil {
			return activities, err
		}

		pIds, err := unflat(predecessorsId)
		if err != nil {
			return activities, err
		}
		sIds, err := unflat(successorsId)
		if err != nil {
			return activities, err
		}

		activities[id] = &activity.Activity{
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

	return activities, nil
}

func UpdateActivityById(sqldb *sql.DB, act *activity.Activity, id int) (int64, error) {
	var n int64
	stmt := fmt.Sprintf("UPDATE %s SET description = %q, duration = %.6f, predecessorsId=%q, successorsId=%q, start = %d, finish = %d, cost = %.6f WHERE id = %d",
		TableName,
		act.Description,
		act.Duration.Seconds(),
		flat(act.PredecessorsId),
		flat(act.SuccessorsId),
		act.Start.Unix(),
		act.Finish.Unix(),
		act.Cost,
		id,
	)
	res, err := sqldb.Exec(stmt)
	if err != nil {
		return n, err
	}
	n, err = res.RowsAffected()
	if err != nil {
		return n, err
	}
	return n, nil
}

func UpdatePredecessorsById(sqldb *sql.DB, id int, predecessorsId []int) (int64, error) {
	var n int64
	stmt := fmt.Sprintf("UPDATE %s SET predecessorsId=%q WHERE id = %d",
		TableName,
		flat(predecessorsId),
		id,
	)
	res, err := sqldb.Exec(stmt)
	if err != nil {
		return n, err
	}
	n, err = res.RowsAffected()
	if err != nil {
		return n, err
	}
	return n, nil
}

func UpdateSuccessorsById(sqldb *sql.DB, id int, successorsId []int) (int64, error) {
	var n int64
	stmt := fmt.Sprintf("UPDATE %s SET successorsId=%q WHERE id = %d",
		TableName,
		flat(successorsId),
		id,
	)
	res, err := sqldb.Exec(stmt)
	if err != nil {
		return n, err
	}
	n, err = res.RowsAffected()
	if err != nil {
		return n, err
	}
	return n, nil
}

func DeleteActivityById(sqldb *sql.DB, id int) (int64, error) {
	var n int64
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = %d", TableName, id)
	res, err := sqldb.Exec(stmt)
	if err != nil && err != sql.ErrNoRows {
		return n, err
	}
	n, err = res.RowsAffected()
	if err != nil {
		return n, err
	}
	return n, nil
}

func DeleteActivitiesById(sqldb *sql.DB, ids []int) (int64, error) {
	var n int64
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", TableName, flat(ids))
	res, err := sqldb.Exec(stmt)
	if err != nil && err != sql.ErrNoRows {
		return n, err
	}
	n, err = res.RowsAffected()
	if err != nil {
		return n, err
	}
	return n, nil
}
