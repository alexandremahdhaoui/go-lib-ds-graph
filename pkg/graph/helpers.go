package graph

import (
	"fmt"
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
)

// IsAcyclic uses KahnTopologicalSortingAlgorithm to determine if a given Graph is acyclic.
//
// Computation complexity is linear: O(|Vertices| + |Edges|) with:
//   - |Vertices|=len(Graph.Nodes())
//   - |Nodes|=len(Graph.Edges())
func IsAcyclic(g api.Graph) (bool, error) {
	if !g.IsDirected() {
		return false, fmt.Errorf("cannot determine if given graph is acyclic; graph should be directed")
	}
	_, err := KahnTopologicalSortingAlgorithm(g)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// NodesWithoutIncomingEdge returns the set of nodes with no incoming edge as []api.Node
//
// 1. Computes the sets of nodes with incoming edges.
//
// 2. Loop over all nodes of the graph, if NOT node in the set of nodes with incoming edges, then append it to the
// list of nodes without edges
func NodesWithoutIncomingEdge(nodes []api.Node, edges []api.Edge) []api.Node {
	var with map[api.Node]struct{}
	for _, edge := range edges {
		with[edge.End()] = struct{}{}
	}

	var without []api.Node
	for _, node := range nodes {
		if _, ok := with[node]; ok {
			continue
		}
		without = append(without, node)
	}
	return without
}
