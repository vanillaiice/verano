package text

import (
	"errors"
	"fmt"

	"github.com/m1gwings/treedrawer/tree"
	"github.com/vanillaiice/verano/activity"
)

const timeFormat = "2 Jan 2006 15:04"

func printNodeContent(act *activity.Activity) string {
	return fmt.Sprintf("%s, %s, %s", act.Description, act.Start.Format(timeFormat), act.Duration.String())
}

func MakeTree(activitiesMap map[int]*activity.Activity, rootId int) (t *tree.Tree, err error) {
	if len(activitiesMap) == 0 {
		return t, errors.New("no activities provided")
	}

	var walk func(*tree.Tree, *activity.Activity)
	walk = func(node *tree.Tree, act *activity.Activity) {
		for i, p := range act.PredecessorsId {
			node.AddChild(tree.NodeString(printNodeContent(activitiesMap[p])))
			child, err := node.Child(i)
			if err != nil {
				return
			}
			walk(child, activitiesMap[p])
		}
	}

	t = tree.NewTree(tree.NodeString(printNodeContent(activitiesMap[rootId])))
	walk(t, activitiesMap[rootId])

	return t, nil
}
