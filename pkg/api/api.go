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

package api

type Graph interface {
	Edges() []Edge
	Nodes() []Node
	IsDirected() bool
}

type Node interface {
	AdjacentNodes() []Node
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

type Visitor interface {
	// Visit returns a bool to signal if the visitor found a match.
	Visit(node Node) bool
}
