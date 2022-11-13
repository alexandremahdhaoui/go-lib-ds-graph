package api

type Graph interface {
	Edges() []Edge
	Nodes() []Node
	IsDirected() bool
}

type Node interface {
	ConnectedNodes() []Node
	Edges() []Edge
}

type Edge interface {
	Start() Node
	End() Node
}

type Builder interface {
	Build() (Graph, error)

	AddEdge(edge Edge) Builder
	AddNode(node Node) Builder

	SetEdges(edges []Edge) Builder
	SetNodes(nodes []Node) Builder

	InduceEdges() Builder
	InduceNodes() Builder
}

type EdgeBuilder interface {
	Build() Edge
	SetStart(node Node) EdgeBuilder
	SetEnd(node Node) EdgeBuilder
}

type NodeBuilder interface {
	Build() []Node
	SetEdges(edges []Node) NodeBuilder
}
