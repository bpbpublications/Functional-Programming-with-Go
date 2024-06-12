package main

import "fmt"

// Define a structure to model our data
type Employee struct {
    Name string
    Position string
    Age int
}

// A functional factory for creating employees
func NewEmployeeFactory(position string) func(name string, age int) *Employee {
    return func(name string, age int) *Employee {
        return &Employee{
            Name: name,
            Position: position,
            Age: age,
        }
    }
}

func main() {
    // Creating different employee factories based on position
    developerFactory := NewEmployeeFactory("Developer")
    managerFactory := NewEmployeeFactory("Manager")

    // Creating instances using the factories
    john := developerFactory("John Doe", 28)
    jane := managerFactory("Jane Doe", 34)

    // Outputting the created instances
    fmt.Println(john) // &{John Doe Developer 28}
    fmt.Println(jane) // &{Jane Doe Manager 34}
}
