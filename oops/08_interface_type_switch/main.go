package main

import "fmt"

func describe(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println("string:", v)
    case int:
        fmt.Println("integer:", v)
    default:
        fmt.Printf("other: %T\n", v)
    }
}

func main() {
    describe("Go")
    describe(42)
    describe(3.14)
}
