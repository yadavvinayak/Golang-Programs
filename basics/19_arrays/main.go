package main

import "fmt"

func main() {
    numbers := [5]int{4, 2, 8, 1, 5}
    total := 0
    for _, number := range numbers {
        total += number
    }
    fmt.Println("array:", numbers, "sum:", total)
}
