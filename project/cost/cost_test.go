package cost

import (
	"testing"

	"github.com/vanillaiice/verano/activity"
)

var activities = []*activity.Activity{
	{Cost: 100.25},
	{Cost: 200.35},
	{Cost: 199.40},
}

func TestTotalCost(t *testing.T) {
	totalCost := TotalCost(activities)
	expected := float64(500)
	if totalCost != expected {
		t.Errorf("got %f, want %f", totalCost, expected)
	}
}
