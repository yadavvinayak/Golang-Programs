package main

import "fmt"

type Address struct{ City string }
type Employee struct {
    Name string
    Address
}

func main() {
    employee := Employee{Name: "Ravi", Address: Address{City: "Pune"}}
    fmt.Println(employee.Name, employee.City)
}
