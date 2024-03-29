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

// New Graph is directed by default
func New() Graph {
	return Graph{directed: true}
}

type Graph struct {
	nodes    []api.Node
	edges    []api.Edge
	directed bool
}

func (g *Graph) Nodes() []api.Node {
	return g.nodes
}

func (g *Graph) Edges() []api.Edge {
	return g.edges
}

func (g *Graph) IsDirected() bool {
	return g.directed
}
