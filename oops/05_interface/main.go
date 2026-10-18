package main

import "fmt"

type Speaker interface{ Speak() string }

type Dog struct{ Name string }
func (d Dog) Speak() string { return d.Name + " says woof" }

func main() {
    var speaker Speaker = Dog{Name: "Milo"}
    fmt.Println(speaker.Speak())
}
