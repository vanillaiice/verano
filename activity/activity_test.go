package activity

import (
	"slices"
	"testing"
	"time"
)

var activity = &Activity{Id: 69, Description: "Testing", Duration: time.Hour, Start: time.Now(), Finish: time.Now(), PredecessorsId: []int{1, 2, 3}, SuccessorsId: []int{4, 5, 6}, Cost: 1000}

func TestAddPredecessor(t *testing.T) {
	temp := *activity
	err := temp.AddPredecessor(7)
	if err != nil {
		t.Error(err)
	}
	expected := []int{1, 2, 3, 7}
	if slices.Compare(temp.PredecessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.PredecessorsId, expected)
	}
	err = temp.AddPredecessor(7)
	if err == nil {
		t.Error("expected AddPredecessor to fail")
	}
}

func TestAddSuccessor(t *testing.T) {
	temp := *activity
	err := temp.AddSuccessor(8)
	if err != nil {
		t.Error(err)
	}
	expected := []int{4, 5, 6, 8}
	if slices.Compare(temp.SuccessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.SuccessorsId, expected)
	}
	err = temp.AddSuccessor(8)
	if err == nil {
		t.Error("expected AddSuccessor to fail")
	}
}

func TestRemovePredecessor(t *testing.T) {
	temp := *activity
	err := temp.RemovePredecessor(3)
	if err != nil {
		t.Error(err)
	}
	expected := []int{1, 2}
	if slices.Compare(temp.PredecessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.PredecessorsId, expected)
	}
	err = temp.RemovePredecessor(3)
	if err == nil {
		t.Error("expected RemovePredecessor to fail")
	}
}

func TestRemoveSuccessor(t *testing.T) {
	temp := *activity
	err := temp.RemoveSuccessor(6)
	if err != nil {
		t.Error(err)
	}
	expected := []int{4, 5}
	if slices.Compare(temp.SuccessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.SuccessorsId, expected)
	}
	err = temp.RemovePredecessor(6)
	if err == nil {
		t.Error("expected RemoveSuccessor to fail")
	}
}

func TestUpdatePredecessorId(t *testing.T) {
	temp := *activity
	err := temp.UpdatePredecessorId(3, 33)
	if err != nil {
		t.Error(err)
	}
	expected := []int{1, 2, 33}
	if slices.Compare(temp.PredecessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.SuccessorsId, expected)
	}
	err = temp.UpdatePredecessorId(3, 33)
	if err == nil {
		t.Error("expected UpdatePredecessorId to fail")
	}
}

func TestUpdateSuccessorId(t *testing.T) {
	temp := *activity
	err := temp.UpdateSuccessorId(6, 66)
	if err != nil {
		t.Error(err)
	}
	expected := []int{4, 5, 66}
	if slices.Compare(temp.SuccessorsId, expected) != 0 {
		t.Errorf("got %v, want %v", temp.SuccessorsId, expected)
	}
	err = temp.UpdateSuccessorId(3, 33)
	if err == nil {
		t.Error("expected UpdateSuccessorId to fail")
	}
}
