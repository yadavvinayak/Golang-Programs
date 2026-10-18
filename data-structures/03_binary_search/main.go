package main

import (
    "fmt"
    "sort"
)

func binarySearch(values []int, target int) int {
    return sort.Search(len(values), func(i int) bool { return values[i] >= target })
}

func main() {
    values := []int{1, 3, 5, 7, 9}
    index := binarySearch(values, 7)
    if index < len(values) && values[index] == 7 {
        fmt.Println("found at index", index)
    } else {
        fmt.Println("not found")
    }
}
