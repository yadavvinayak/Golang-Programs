package main

import "fmt"

type Counter struct{ Value int }

func (c *Counter) Increment() { c.Value++ }

func main() {
    c := Counter{}
    c.Increment()
    c.Increment()
    fmt.Println(c.Value)
}
