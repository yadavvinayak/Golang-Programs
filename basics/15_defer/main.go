package main

import "fmt"

func main() {
    defer fmt.Println("This runs when main returns")
    fmt.Println("First")
    fmt.Println("Second")
}
