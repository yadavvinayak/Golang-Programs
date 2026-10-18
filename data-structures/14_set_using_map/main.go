package main

import "fmt"

func main() {
    set := map[string]struct{}{}
    for _, item := range []string{"go", "rust", "go", "python"} { set[item] = struct{}{} }
    delete(set, "rust")
    for item := range set { fmt.Println(item) }
}
