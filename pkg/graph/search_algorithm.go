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

import "gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"

type NodeSet map[api.Node]struct{}

func isUnvisited(node api.Node, visited map[api.Node]struct{}) bool {
	_, ok := visited[node]
	return !ok
}

// DFS is an implementation of Depth-first Search algorithm for Graph.
// Provides a simplified interface for InternalDFS.
func DFS(
	node api.Node,
	visitor api.Visitor,
) bool {
	visited := make(NodeSet)
	var stack []api.Node
	return InternalDFS(node, visitor, visited, stack)
}

// InternalDFS is a Depth-first Search algorithm for Graph.
// Returns a boolean, that can be returned as true from visitor.Visit().
// This return bool is used to stop searching if the visitor found a satisfying match.
func InternalDFS(
	node api.Node,
	visitor api.Visitor,
	visited NodeSet,
	stack []api.Node,
) bool {
	if isUnvisited(node, visited) {
		if stop := visitor.Visit(node); stop {
			return stop
		}

		visited[node] = struct{}{}
		stack = append(stack, node)
	}

	for _, adj := range node.AdjacentNodes() {
		if isUnvisited(adj, visited) {
			if stop := InternalDFS(adj, visitor, visited, stack); stop {
				return stop
			}
		}
	}

	for len(stack) > 0 {
		stack, node, _ = unstack(stack)
		if stop := InternalDFS(node, visitor, visited, stack); stop {
			return stop
		}
	}
	return false
}

// BFS is an implementation of Breadth-first Search algorithm for Graph.
func BFS(
	node api.Node,
	visitor api.Visitor,
) bool {
	visited := make(NodeSet)
	var queue []api.Node
	return InternalBFS(node, visitor, visited, queue)
}

// InternalBFS is an implementation of Breadth-first Search algorithm for Graph.
func InternalBFS(
	node api.Node,
	visitor api.Visitor,
	visited NodeSet,
	queue []api.Node,
) bool {

	if isUnvisited(node, visited) {
		if stop := visitor.Visit(node); stop {
			return stop
		}
	}

	for _, adj := range node.AdjacentNodes() {
		if isUnvisited(adj, visited) {
			if stop := visitor.Visit(adj); stop {
				return stop
			}
			visited[adj] = struct{}{}
			queue = append(queue, adj)
		}
	}

	for len(queue) > 0 {
		queue, node, _ = dequeue(queue)
		InternalBFS(node, visitor, visited, queue)
	}
	return false
}
