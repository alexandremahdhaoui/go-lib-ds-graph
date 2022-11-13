# go-lib-ds-graph

[[_TOC_]]

# Repository Structure


| Path                               | Description                                                                                               |
|------------------------------------|-----------------------------------------------------------------------------------------------------------|
| `pkg/`                             | contains all go code of the library.                                                                      |
| `pkg/api/`                         | contains all interface definitions for `Graph`, `Builder`, `Node`, `NodeBuilder`, `Edge` & `EdgeBuilder`. |
| `pkg/dag/`                         | contains builder for Directed Acyclic Graph datastructures.                                               |
| `pkg/graph/`                       | contains all code related to graph datastructures.                                                        |
| `pkg/graph/builder.go`             | includes `graph.Builder` struct compliant to the `api.GraphBuilder`.                                      |
| `pkg/graph/graph.go`               | includes all code related to the `graph.Graph` struct.                                                    |
| `pkg/graph/helpers.go`             | includes all helpers functions. Uses the API.                                                             |
| `pkg/graph/topological_sorting.go` | includes algorithm to compute a topologically sorted order of a graph.                                    |

# TODO

- implement DFS & BFS.