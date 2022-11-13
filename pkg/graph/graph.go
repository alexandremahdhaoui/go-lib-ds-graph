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
