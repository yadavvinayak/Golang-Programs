package main

import "fmt"

func main() {
    double := func(n int) int { return n * 2 }
    fmt.Println(double(8))
}
