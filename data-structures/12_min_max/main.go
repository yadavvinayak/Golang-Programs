package main

import (
    "fmt"
    "slices"
)

func main() {
    values := []int{12, 4, 19, -2, 7}
    fmt.Println("min:", slices.Min(values), "max:", slices.Max(values))
}
