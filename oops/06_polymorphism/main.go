package main

import "fmt"

type Shape interface{ Area() float64 }

type Square struct{ Side float64 }
func (s Square) Area() float64 { return s.Side * s.Side }

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }

func printArea(s Shape) { fmt.Printf("area: %.2f\n", s.Area()) }

func main() {
    printArea(Square{Side: 4})
    printArea(Circle{Radius: 2})
}
