/*
Copyright 2022 Alexandre Mahdhaoui

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	with := make(map[api.Node]struct{})
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

func unstack[T any](stack []T) ([]T, T, error) {
	switch i := len(stack); {
	case i > 1:
		return stack[:i-1], stack[i-1], nil
	case i == 1:
		return make([]T, 0), stack[0], nil
	default:
		panic("len(stack) should be > 0 to unstack")
	}
}

func dequeue[T any](queue []T) ([]T, T, error) {
	switch i := len(queue); {
	case i > 1:
		return queue[1:], queue[0], nil
	case i == 1:
		return make([]T, 0), queue[0], nil
	default:
		panic("len(queue) should be > 0 to dequeue")
	}
}

func edgeSet(g api.Graph) map[api.Edge]struct{} {
	set := make(map[api.Edge]struct{})
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
