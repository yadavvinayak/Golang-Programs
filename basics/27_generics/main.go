package main

import "fmt"

func first[T any](items []T) (T, bool) {
    if len(items) == 0 {
        var zero T
        return zero, false
    }
    return items[0], true
}

func main() {
    value, ok := first([]string{"Go", "Python"})
    fmt.Println(value, ok)
}
