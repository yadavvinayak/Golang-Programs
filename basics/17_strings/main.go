package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "  go programming  "
    trimmed := strings.TrimSpace(text)
    fmt.Println("trimmed:", trimmed)
    fmt.Println("uppercase:", strings.ToUpper(trimmed))
    fmt.Println("contains Go:", strings.Contains(trimmed, "go"))
    fmt.Println("words:", strings.Split(trimmed, " "))
}
