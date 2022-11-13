package graph

import "gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"

type Builder struct {
	graph Graph
}

func NewBuilder() api.Builder {
	return &Builder{graph: New()}
}

func NewBuilderFrom(g Graph) api.Builder {
	return &Builder{graph: g}
}

func (b *Builder) Build() (api.Graph, error) {
	return &b.graph, nil
}

func (b *Builder) AddEdge(edge api.Edge) api.Builder {
	return b.SetEdges(append(b.graph.Edges(), edge))
}

func (b *Builder) AddNode(node api.Node) api.Builder {
	return b.SetNodes(append(b.graph.Nodes(), node))
}

func (b *Builder) SetEdges(edges []api.Edge) api.Builder {
	b.graph.edges = edges
	return b
}

func (b *Builder) SetNodes(nodes []api.Node) api.Builder {
	b.graph.nodes = nodes
	return b
}

// InduceEdges generates Edges from all Nodes of the Graph.
//
// 1. Generates a slice of Edge by looping over Graph.Nodes().
//
// 2. Use the Builder to set the generated []api.Edge.
func (b *Builder) InduceEdges() api.Builder {
	var edges []api.Edge
	for _, node := range b.graph.Nodes() {
		for _, edge := range node.Edges() {
			edges = append(edges, edge)
		}
	}
	b.SetEdges(edges)
	return b
}

// InduceNodes generates Nodes from all Edges of the Graph.
//
// 1. Generate a Set of nodes by looping over Graph.Edges().
//
// 2. Then loop through all nodes of the nodeSet and generates a slice of api.Node.
//
// 3. Use the Builder to set the generated []api.Node.
func (b *Builder) InduceNodes() api.Builder {
	var nodeSet map[api.Node]struct{}
	for _, edge := range b.graph.Edges() {
		nodeSet[edge.Start()] = struct{}{}
		nodeSet[edge.End()] = struct{}{}
	}
	var nodes []api.Node
	for node, _ := range nodeSet {
		nodes = append(nodes, node)
	}
	b.SetNodes(nodes)
	return b
}
