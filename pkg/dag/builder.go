package dag

import (
	"fmt"
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/api"
	"gitlab.com/alexandre.mahdhaoui/go-lib-ds-graph/pkg/graph"
)

type Builder struct {
	builder api.Builder
}

func NewBuilder() api.Builder {
	return &Builder{builder: graph.NewBuilder()}
}

func (b *Builder) Build() (api.Graph, error) {
	g, err := b.builder.Build()
	if err != nil {
		return nil, err
	}
	acyclic, err := graph.IsAcyclic(g)
	if err != nil {
		return nil, err
	}
	if !acyclic {
		return nil, fmt.Errorf("graph should be acyclic")
	}
	return g, nil
}

func (b *Builder) AddEdge(edge api.Edge) api.Builder {
	b.builder.AddEdge(edge)
	return b
}

func (b *Builder) AddNode(node api.Node) api.Builder {
	b.builder.AddNode(node)
	return b
}

func (b *Builder) SetEdges(edges []api.Edge) api.Builder {
	b.builder.SetEdges(edges)
	return b
}

func (b *Builder) SetNodes(nodes []api.Node) api.Builder {
	b.builder.SetNodes(nodes)
	return b
}

func (b *Builder) InduceEdges() api.Builder {
	b.builder.InduceEdges()
	return b
}

func (b *Builder) InduceNodes() api.Builder {
	b.builder.InduceNodes()
	return b
}
