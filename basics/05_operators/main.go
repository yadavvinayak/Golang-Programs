package main

import "fmt"

func main() {
    a, b := 17, 5
    fmt.Println("sum:", a+b, "difference:", a-b, "product:", a*b)
    fmt.Println("quotient:", a/b, "remainder:", a%b)
    fmt.Println("a > b:", a > b, "both positive:", a > 0 && b > 0)
}
