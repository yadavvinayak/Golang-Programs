package main

import "fmt"

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func main() {
    r := Rectangle{Width: 4, Height: 3}
    fmt.Println("area:", r.Area())
}
