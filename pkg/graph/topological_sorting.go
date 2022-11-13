package graph

import (
	"fmt"
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
)

// KahnTopologicalSortingAlgorithm determine a topologically sorted order of a given directed Graph.
//
// Returns a []api.Node or an error.
// Computation complexity is linear: O(|Vertices| + |Edges|) with:
//   - |Vertices|=len(Graph.Nodes())
//   - |Nodes|=len(Graph.Edges())
func KahnTopologicalSortingAlgorithm(g api.Graph) ([]api.Node, error) {
	var sortedOrder []api.Node
	remainingEdges := edgeSet(g)
	without := NodesWithoutIncomingEdge(g.Nodes(), g.Edges())

	var node api.Node
	// while there is still nodes without incoming edges
	for len(without) > 0 {
		// pop one of these nodes from the list
		without, node, _ = dequeue(without)
		// append the node to the sortedOrder
		sortedOrder = append(sortedOrder, node)
		// for each edges of this node to other connected Node:
		for _, edge := range node.Edges() {
			// remove the edge from the remainingEdges of the graph.
			delete(remainingEdges, edge)
			connectedNode := edge.End()
			// if the connected node has no other incoming edges:
			if len(nodeWithoutIncomingEdge(connectedNode, remainingEdges)) == 1 {
				// add it to the list of nodes without incoming edges
				without = append(without, connectedNode)
			}
		}
	}
	// If the graph has remaining edges, then there is at least one cycle
	if len(remainingEdges) != 0 {
		return nil, fmt.Errorf("sorted order could not be computed; graph has at least one cycle")
	}
	return sortedOrder, nil
}
