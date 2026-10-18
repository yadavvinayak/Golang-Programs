package main

import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
    fmt.Println(Point{X: 3, Y: 5})
}
