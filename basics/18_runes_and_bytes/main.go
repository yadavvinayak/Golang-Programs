package main

import "fmt"

func main() {
    text := "Go世界"
    fmt.Println("bytes:", len(text))
    fmt.Println("runes:", len([]rune(text)))
    for index, r := range text {
        fmt.Printf("%d: %c\n", index, r)
    }
}
