package db

import (
	"os"
	"testing"
	"time"

	"github.com/vanillaiice/verano/activity"
)

var id = 1
var duration = 1 * time.Hour
var start = time.Date(2024, 12, 31, 18, 0, 0, 0, time.Local)
var finish = start.Add(duration)

func openDB() (*DB, error) {
	sqldb, err := New("test.db")
	if err != nil {
		return sqldb, err
	}
	return sqldb, err
}

func deleteDB() error {
	err := os.Remove("test.db")
	if err != nil {
		return err
	}
	id = 1
	return nil
}

func insertActivity(sqldb *DB, descr string) error {
	start := time.Date(2024, 12, 31, 18, 0, 0, 0, time.Local)
	duration := 1 * time.Hour
	activity := &activity.Activity{
		Id:          id,
		Description: descr,
		Duration:    duration,
		Start:       start,
		Finish:      finish,
	}
	err := sqldb.InsertActivity(activity)
	if err != nil {
		return err
	}
	id++
	return nil
}

func TestOpen(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	sqldb.DB.Close()
}

func TestInsertActivity(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()
	insertActivity(sqldb, "tip your landlord")
	if err != nil {
		t.Error(err)
	}
}

func TestGetActivityById(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	a, err := sqldb.GetActivityById(1)
	if err != nil {
		t.Error(err)
	}

	descr := "tip your landlord"
	if a.Description != descr {
		t.Errorf("description: want %v, got %v", descr, a.Description)
	}
	if a.Duration != duration {
		t.Errorf("duration: want %v, got %v", duration, a.Duration)
	}
	if a.Start != start {
		t.Errorf("start: want %v, got %v", start, a.Start)
	}
	if a.Finish != finish {
		t.Errorf("finish: want %v, got %v", finish, a.Finish)
	}
}

func TestGetActivitiesById(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	descr := "tip your landlord"
	descr2 := "tip your landlord even more"

	err = insertActivity(sqldb, descr2)
	if err != nil {
		t.Error(err)
	}

	a, err := sqldb.GetActivitiesById([]int{1, 2})
	if err != nil {
		t.Error(t)
	}

	if len(a) != 2 {
		t.Errorf("wrong len for activities, want %d, got %d", 2, len(a))
	}

	if a[0].Description != descr {
		t.Errorf("description: want %v, got %v", descr, a[0].Description)
	}
	if a[0].Duration != duration {
		t.Errorf("duration: want %v, got %v", duration, a[0].Duration)
	}
	if a[0].Start != start {
		t.Errorf("start: want %v, got %v", start, a[0].Start)
	}
	if a[0].Finish != finish {
		t.Errorf("finish: want %v, got %v", finish, a[0].Finish)
	}

	if a[1].Description != descr2 {
		t.Errorf("description: want %v, got %v", descr, a[1].Description)
	}
	if a[1].Duration != duration {
		t.Errorf("duration: want %v, got %v", duration, a[1].Duration)
	}
	if a[1].Start != start {
		t.Errorf("start: want %v, got %v", start, a[1].Start)
	}
	if a[1].Finish != finish {
		t.Errorf("finish: want %v, got %v", finish, a[1].Finish)
	}
}

func TestGetAllActivities(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	a, err := sqldb.GetAllActivities()
	if err != nil {
		t.Error(t)
	}

	if len(a) != 2 {
		t.Errorf("wrong len for activities, want %d, got %d", 2, len(a))
	}
}

func TestUpdateActivityById(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	descr2 := "tip your landlord even more"
	duration2 := 3 * time.Hour
	start2 := time.Date(2024, 1, 1, 10, 0, 0, 0, time.Local)
	finish2 := start2.Add(duration2)

	n, err := sqldb.UpdateActivityById(&activity.Activity{Description: descr2, Duration: duration2, Start: start2, Finish: finish2}, 1)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("Unexpected error, expected 1 row to be affected, got %d", n)
	}

	a, err := sqldb.GetActivityById(1)
	if err != nil {
		t.Error(err)
	}

	if a.Description != descr2 {
		t.Errorf("description: want %v, got %v", descr2, a.Description)
	}
	if a.Duration != duration2 {
		t.Errorf("duration: want %v, got %v", duration2, a.Duration)
	}
	if a.Start != start2 {
		t.Errorf("start: want %v, got %v", start2, a.Start)
	}
	if a.Finish != finish2 {
		t.Errorf("finish: want %v, got %v", finish2, a.Finish)
	}
}

func TestDeleteActivityById(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	n, err := sqldb.DeleteActivityById(2)
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Errorf("Unexpected error, expected 1 row to be affected, got %d", n)
	}
}

func TestDeleteActivitiesById(t *testing.T) {
	sqldb, err := openDB()
	if err != nil {
		t.Error(err)
	}
	defer sqldb.DB.Close()

	err = insertActivity(sqldb, "tip your landlord")
	if err != nil {
		t.Error(err)
	}

	n, err := sqldb.DeleteActivitiesById([]int{1, 3})
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Errorf("Unexpected error, expected 2 rows to be affected, got %d", n)
	}

	err = deleteDB()
	if err != nil {
		t.Error(err)
	}
}

func TestInsertActivities(t *testing.T) {
	sqldb, err := Open("test.db")
	if err != nil {
		t.Error(err)
	}
	defer sqldb.Close()

	predId := []int{1, 2, 3}
	succId := []int{5, 6, 7}

	activities := []*activity.Activity{
		{
			Id:             43,
			Description:    "buy eggs",
			Duration:       duration,
			PredecessorsId: predId,
			SuccessorsId:   succId,
			Start:          start,
			Finish:         finish,
		},
		{
			Id:             35,
			Description:    "cook eggs",
			Duration:       duration,
			PredecessorsId: predId,
			SuccessorsId:   succId,
			Start:          start,
			Finish:         finish,
		},
		{
			Id:             22,
			Description:    "eat eggs",
			Duration:       duration,
			PredecessorsId: predId,
			SuccessorsId:   succId,
			Start:          start,
			Finish:         finish,
		},
	}

	err = InsertActivities(sqldb, activities)
	if err != nil {
		t.Error(err)
	}

	err = deleteDB()
	if err != nil {
		t.Error(err)
	}
}
