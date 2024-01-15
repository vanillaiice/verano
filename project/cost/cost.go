package cost

import (
	"github.com/vanillaiice/verano/activity"
)

func TotalCost(activities []*activity.Activity) (totalCost float64) {
	for _, act := range activities {
		totalCost += act.Cost
	}
	return
}
