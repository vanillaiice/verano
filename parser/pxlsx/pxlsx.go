package pxlsx

import (
	"time"

	"github.com/tealeg/xlsx/v3"
	"github.com/vanillaiice/verano/activity"
	"github.com/vanillaiice/verano/db"
	"github.com/vanillaiice/verano/util"
)

// ExportToDb populates the database with activities in xlsx format.
func ExportToDb(sqldb *db.DB, sheet *xlsx.Sheet, duplicateInsertPolicy ...db.DuplicateInsertPolicy) (err error) {
	activities, err := XLSXToActivities(sheet)
	if err != nil {
		return
	}

	err = sqldb.InsertActivities(activities, duplicateInsertPolicy...)
	return
}

// ActivitiesToXLSX converts a slice of activities to xlsx format.
func ActivitiesToXLSX(activities []*activity.Activity, sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	header := []string{"Id", "Description", "Duration", "Start", "Finish", "PredecessorsId", "SuccessorsId", "Cost"}
	for _, h := range header {
		c := row.AddCell()
		c.SetString(h)
		row.PushCell(c)
	}

	for _, activity := range activities {
		row = sheet.AddRow()
		cells := []*xlsx.Cell{}

		id := row.AddCell()
		id.SetInt(activity.Id)
		cells = append(cells, id)

		description := row.AddCell()
		description.SetString(activity.Description)
		cells = append(cells, description)

		duration := row.AddCell()
		duration.SetString(activity.Duration.String())
		cells = append(cells, duration)

		start := row.AddCell()
		start.SetDateTime(activity.Start)
		cells = append(cells, start)

		finish := row.AddCell()
		finish.SetDateTime(activity.Finish)
		cells = append(cells, finish)

		predecessorsId := row.AddCell()
		predecessorsId.SetString(util.Flat(activity.PredecessorsId))
		cells = append(cells, predecessorsId)

		successorsId := row.AddCell()
		successorsId.SetString(util.Flat(activity.SuccessorsId))
		cells = append(cells, successorsId)

		cost := row.AddCell()
		cost.SetFloat(activity.Cost)
		cells = append(cells, cost)

		for _, c := range cells {
			row.PushCell(c)
		}
	}
}

// XLSXToActivities converts activities in xlsx format to a slice of activities.
func XLSXToActivities(sheet *xlsx.Sheet) (activities []*activity.Activity, err error) {
	for i := 0; i < sheet.MaxRow; i++ {
		row, err := sheet.Row(i)
		if err != nil {
			return activities, err
		}

		if row.GetCell(0).String() == "Id" {
			continue
		}
		act := &activity.Activity{}

		id, err := row.GetCell(0).Int()
		if err != nil {
			return activities, err
		}

		description := row.GetCell(1).String()

		duration, err := time.ParseDuration(row.GetCell(2).String())
		if err != nil {
			return activities, err
		}

		start := row.GetCell(3)
		if start.String() == "" || start.String() == "0" {
			act.Start = time.Time{}
		} else {
			startTime, err := start.GetTime(false)
			if err != nil {
				return activities, err
			}
			act.Start = startTime
		}

		finish := row.GetCell(4)
		if finish.String() == "" || finish.String() == "0" {
			act.Finish = time.Time{}
		} else {
			finishTime, err := finish.GetTime(false)
			if err != nil {
				return activities, err
			}
			act.Finish = finishTime
		}

		predecessorsId, err := util.Unflat(row.GetCell(5).String())
		if err != nil {
			return activities, err
		}

		successorsId, err := util.Unflat(row.GetCell(6).String())
		if err != nil {
			return activities, err
		}

		cost := row.GetCell(7)
		if cost.String() == "" {
			act.Cost = 0
		} else {
			costFloat, err := cost.Float()
			if err != nil {
				return activities, err
			}
			act.Cost = costFloat
		}

		act.Id = id
		act.Description = description
		act.Duration = duration
		act.PredecessorsId = predecessorsId
		act.SuccessorsId = successorsId

		activities = append(activities, act)
	}

	return
}
