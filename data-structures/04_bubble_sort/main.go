package main

import "fmt"

func main() {
    values := []int{5, 1, 4, 2, 8}
    for end := len(values)-1; end > 0; end-- {
        for i := 0; i < end; i++ {
            if values[i] > values[i+1] {
                values[i], values[i+1] = values[i+1], values[i]
            }
        }
    }
    fmt.Println(values)
}
