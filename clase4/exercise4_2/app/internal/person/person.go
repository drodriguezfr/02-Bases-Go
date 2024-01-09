package person

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee(empleado Employee) {
	fmt.Printf("%+v\n", empleado)
}
