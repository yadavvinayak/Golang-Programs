package main

import "fmt"

func main() {
    values := []string{"go", "python", "go", "rust", "go", "rust"}
    counts := make(map[string]int)
    for _, value := range values { counts[value]++ }
    fmt.Println(counts)
}
