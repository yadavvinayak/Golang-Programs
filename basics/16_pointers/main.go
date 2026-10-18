package main

import "fmt"

func main() {
    value := 10
    pointer := &value
    fmt.Println("value:", value, "pointer points to:", *pointer)
    *pointer = 20
    fmt.Println("updated value:", value)
}
