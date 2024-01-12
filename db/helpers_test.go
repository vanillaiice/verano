package db

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFlat(t *testing.T) {
	a := []int{1, 2, 3, 4}
	aflat := flat(a)

	if "1,2,3,4" != aflat {
		t.Error("Unexpected error")
	}
}

func TestUnflat(t *testing.T) {
	s := "1,2,3,4"
	sunflat, err := unflat(s)
	if err != nil {
		t.Error(err)
	}
	want := []int{1, 2, 3, 4}
	if slices.Compare(sunflat, want) != 0 {
		t.Errorf("Error, want %v, got %v", want, sunflat)
	}
}
