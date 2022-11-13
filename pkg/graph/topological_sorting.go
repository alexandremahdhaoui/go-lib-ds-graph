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
	for empty := false; !empty; empty = len(without) == 0 {
		// pop one of these nodes from the list
		node, without, _ = popQueue(without)
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
		return nil, fmt.Errorf("couldn't not compute sorted order; graph has at least one cycle")
	}
	return sortedOrder, nil
}

func popQueue[T any](queue []T) (T, []T, error) {
	switch i := len(queue); {
	case i > 1:
		return queue[0], queue[1:], nil
	case i == 1:
		return queue[0], make([]T, 0), nil
	default:
		return nil, nil, fmt.Errorf("queue should have at least 1 element")
	}
}

func edgeSet(g api.Graph) map[api.Edge]struct{} {
	var set map[api.Edge]struct{}
	for _, edge := range g.Edges() {
		set[edge] = struct{}{}
	}
	return set
}

func edgeSetToSlice(edges map[api.Edge]struct{}) []api.Edge {
	var edgeSlice []api.Edge
	for edge, _ := range edges {
		edgeSlice = append(edgeSlice, edge)
	}
	return edgeSlice
}

func nodeWithoutIncomingEdge(node api.Node, edges map[api.Edge]struct{}) []api.Node {
	return NodesWithoutIncomingEdge([]api.Node{node}, edgeSetToSlice(edges))
}
