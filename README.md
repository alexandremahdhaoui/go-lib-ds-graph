[[_TOC_]]

# Repository Structure


| Path                               | Description                                                                                  |
|------------------------------------|----------------------------------------------------------------------------------------------|
| `pkg/`                             | The library.                                                                                 |
| `pkg/api/`                         | Interface definitions for `Graph`, `Builder`, `Node`, `NodeBuilder`, `Edge` & `EdgeBuilder`. |
| `pkg/dag/`                         | DAG Builder.                                                                                 |
| `pkg/graph/`                       | Code related to graph datastructures.                                                        |
| `pkg/graph/builder.go`             | `graph.Builder` struct compliant to the `api.GraphBuilder`.                                  |
| `pkg/graph/graph.go`               | `graph.Graph` struct.                                                                        |
| `pkg/graph/helpers.go`             | Helper functions.                                                                            |
| `pkg/graph/search_algorithms.go`   | contains `DFS` & `BFS` algorithms.                                                           |
| `pkg/graph/topological_sorting.go` | Algorithm to compute a topologically sorted order of a graph.                                |
