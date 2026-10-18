package main

import "fmt"

func bfs(graph map[int][]int, start int) {
    visited := map[int]bool{start: true}
    queue := []int{start}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        fmt.Print(node, " ")
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}

func main() {
    graph := map[int][]int{1: {2, 3}, 2: {4}, 3: {4}, 4: {}}
    bfs(graph, 1)
    fmt.Println()
}
