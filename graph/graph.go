package graph

import (
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/vanillaiice/verano/activity"
)

// Time format to use when parsing
const timeFormat = "2 Jan 2006 15:04"

// DrawAndRender draws a graphviz graph from a map of activities
// and then renders the graph to an image.
func DrawAndRender(graph *graphviz.Graphviz, activitiesMap map[int]*activity.Activity, format graphviz.Format, filename string) (err error) {
	g, err := graph.Graph()
	if err != nil {
		return
	}
	defer g.Close()
	if err = Draw(g, activitiesMap); err != nil {
		return
	}
	return GraphToImage(graph, g, format, filename)
}

// Draw draws a graphviz graph from a map of activities.
// The graph shows the relationships between activities,
// and the order of activities.
func Draw(graph *cgraph.Graph, activities map[int]*activity.Activity) (err error) {
	graph.SetRankDir(cgraph.LRRank)
	for k, v := range activities {
		node, err := graph.CreateNode(fmt.Sprint(k))
		if err != nil {
			return err
		}
		node.SetLabel(fmt.Sprintf("%s, FOR %s, START @%s, FINISH @%s", v.Description, v.Duration.String(), v.Start.Format(timeFormat), v.Finish.Format(timeFormat)))
	}
	for k, v := range activities {
		node, err := graph.Node(fmt.Sprint(k))
		if err != nil {
			return err
		}
		for _, successorsId := range v.SuccessorsId {
			node2, err := graph.Node(fmt.Sprint(successorsId))
			if err != nil {
				return err
			}
			if _, err = graph.CreateEdge("", node, node2); err != nil {
				return err
			}
		}
	}
	return
}

// GraphToImage renders a graph to an image
func GraphToImage(graph *graphviz.Graphviz, g *cgraph.Graph, format graphviz.Format, filename string) (err error) {
	return graph.RenderFilename(g, format, filename)
}
