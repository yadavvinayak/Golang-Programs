package main

import "fmt"

func linearSearch(values []int, target int) int {
    for i, value := range values {
        if value == target {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(linearSearch([]int{4, 8, 2, 9}, 2))
}
