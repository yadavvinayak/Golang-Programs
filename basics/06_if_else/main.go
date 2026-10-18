package main

import "fmt"

func main() {
    number := 7
    if number%2 == 0 {
        fmt.Println(number, "is even")
    } else {
        fmt.Println(number, "is odd")
    }
}
