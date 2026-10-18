package main

import "fmt"

func main() {
    var whole int = 42
    decimal := float64(whole)
    fmt.Printf("int=%d float=%.1f\n", whole, decimal)
}
