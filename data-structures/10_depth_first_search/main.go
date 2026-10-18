package main

import "fmt"

func dfs(graph map[int][]int, node int, visited map[int]bool) {
    if visited[node] { return }
    visited[node] = true
    fmt.Print(node, " ")
    for _, next := range graph[node] { dfs(graph, next, visited) }
}

func main() {
    graph := map[int][]int{1: {2, 3}, 2: {4}, 3: {4}, 4: {}}
    dfs(graph, 1, map[int]bool{})
    fmt.Println()
}
