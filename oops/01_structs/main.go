package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Asha", Age: 28}
    fmt.Printf("%s is %d years old.\n", person.Name, person.Age)
}
