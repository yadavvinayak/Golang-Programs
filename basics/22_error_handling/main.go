package main

import (
    "errors"
    "fmt"
)

func positive(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("number must not be negative")
    }
    return n, nil
}

func main() {
    value, err := positive(-2)
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Println(value)
}
