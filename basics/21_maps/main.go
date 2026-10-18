package main

import "fmt"

func main() {
    scores := map[string]int{"Ana": 90, "Ben": 82}
    scores["Kai"] = 95
    score, exists := scores["Ana"]
    fmt.Println("Ana:", score, "exists:", exists)
    delete(scores, "Ben")
    for name, value := range scores {
        fmt.Println(name, value)
    }
}
