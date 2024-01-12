package activity

import (
	"time"
)

// Activity is a struct representing an activity with various attributes.
type Activity struct {
	Id             int           // Unique identifier of the activity
	Description    string        // description of the activity
	Duration       time.Duration // duration of the activity
	Start          time.Time     // Start time of the activity
	Finish         time.Time     // Finish time of he activity
	PredecessorsId []int         // ID of the activities that precede
	SuccessorsId   []int         // ID of the activities that come after
	Cost           float64       // Cost of the activity
}

/*
func (a *Activity) AddPredecessor(id int, sqldb *sql.DB) {
	a.PredecessorsId = append(a.PredecessorsId, id)
	// update in db too
}

func (a *Activity) AddSuccessor(id int) {
	a.SuccessorsId = append(a.SuccessorsId, id)
	// update in db too
}

func (a *Activity) RemovePredecessor(id int) error {
	idx := slices.Index(a.PredecessorsId, id)
	if idx == -1 {
		return errors.New(fmt.Sprintf("No predecessor with id %d", id))
	}
	a.PredecessorsId = slices.Delete(a.PredecessorsId, idx, idx+1)
	return nil
	// update in db too
}

func (a *Activity) RemoveSuccessor(id int) error {
	idx := slices.Index(a.SuccessorsId, id)
	if idx == -1 {
		return errors.New(fmt.Sprintf("No successor with id %d", id))
	}
	a.SuccessorsId = slices.Delete(a.SuccessorsId, idx, idx+1)
	return nil
	// update in db too
}

func (a *Activity) UpdatePredecessorId(oldId, newId int) error {
	idx := slices.Index(a.PredecessorsId, oldId)
	if idx == -1 {
		return errors.New(fmt.Sprintf("No predecessor with id %d", oldId))
	}
	a.PredecessorsId = slices.Replace(a.PredecessorsId, idx, idx+1, newId)
	return nil
	// update in db too
}

func (a *Activity) UpdateSuccessorId(oldId, newId int) error {
	idx := slices.Index(a.SuccessorsId, oldId)
	if idx == -1 {
		return errors.New(fmt.Sprintf("No successor with id %d", oldId))
	}
	a.SuccessorsId = slices.Replace(a.SuccessorsId, idx, idx+1, newId)
	return nil
	// update in db too
}
*/
