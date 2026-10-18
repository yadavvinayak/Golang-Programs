package main

import (
    "fmt"
    "sort"
)

func main() {
    values := []int{9, 2, 7, 1, 5}
    sort.Ints(values)
    fmt.Println(values)
}
