package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}
    numbers = append(numbers, 40)
    part := numbers[1:3]
    fmt.Println("all:", numbers, "slice:", part, "len:", len(numbers), "cap:", cap(numbers))
}
