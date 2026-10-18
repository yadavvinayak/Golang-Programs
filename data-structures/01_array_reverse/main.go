package main

import "fmt"

func main() {
    values := []int{1, 2, 3, 4, 5}
    for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
        values[left], values[right] = values[right], values[left]
    }
    fmt.Println(values)
}
