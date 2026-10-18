package main

import "fmt"

func divide(a, b float64) (float64, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
    if result, ok := divide(10, 2); ok {
        fmt.Println("result:", result)
    }
}
