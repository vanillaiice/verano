package activity

import (
	"fmt"
	"slices"
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

// AddPredecessor adds a predecessor with the given 'id' to the activity's predecessors list.
// It returns an error if the predecessor already exists in the list.
func (a *Activity) AddPredecessor(id int) (err error) {
	if slices.Index(a.PredecessorsId, id) != -1 {
		return fmt.Errorf("predecessor with id %d already exists", id)
	}
	a.PredecessorsId = append(a.PredecessorsId, id)
	return
}

// AddSuccessor adds a successor with the given 'id' to the activity's successors list.
// It returns an error if the successor already exists in the list.
func (a *Activity) AddSuccessor(id int) (err error) {
	if slices.Index(a.SuccessorsId, id) != -1 {
		return fmt.Errorf("successor with id %d already exists", id)
	}
	a.SuccessorsId = append(a.SuccessorsId, id)
	return
}

// RemovePredecessor removes the predecessor with the given 'id' from the activity's predecessors list.
// It returns an error if the predecessor is not found in the list.
func (a *Activity) RemovePredecessor(id int) (err error) {
	idx := slices.Index(a.PredecessorsId, id)
	if idx == -1 {
		return fmt.Errorf("No predecessor with id %d", id)
	}
	a.PredecessorsId = slices.Delete(a.PredecessorsId, idx, idx+1)
	return
}

// RemoveSuccessor removes the successor with the given 'id' from the activity's successors list.
// It returns an error if the successor is not found in the list.
func (a *Activity) RemoveSuccessor(id int) (err error) {
	idx := slices.Index(a.SuccessorsId, id)
	if idx == -1 {
		return fmt.Errorf(fmt.Sprintf("No successor with id %d", id))
	}
	a.SuccessorsId = slices.Delete(a.SuccessorsId, idx, idx+1)
	return
}

// UpdatePredecessorId updates the predecessor ID from 'oldId' to 'newId' in the activity's predecessors list.
// It returns an error if the predecessor with 'oldId' is not found.
func (a *Activity) UpdatePredecessorId(oldId, newId int) (err error) {
	idx := slices.Index(a.PredecessorsId, oldId)
	if idx == -1 {
		return fmt.Errorf(fmt.Sprintf("No predecessor with id %d", oldId))
	}
	a.PredecessorsId = slices.Replace(a.PredecessorsId, idx, idx+1, newId)
	return
}

// UpdateSuccessorId updates the successor ID from 'oldId' to 'newId' in the activity's successors list.
// It returns an error if the successor with 'oldId' is not found.
func (a *Activity) UpdateSuccessorId(oldId, newId int) (err error) {
	idx := slices.Index(a.SuccessorsId, oldId)
	if idx == -1 {
		return fmt.Errorf(fmt.Sprintf("No successor with id %d", oldId))
	}
	a.SuccessorsId = slices.Replace(a.SuccessorsId, idx, idx+1, newId)
	return
}
