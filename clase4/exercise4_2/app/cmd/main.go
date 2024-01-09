package main

import "app/internal/person"

func main() {
	p := person.Person{ID: 1, Name: "David", DateOfBirth: "15.04.2000"}
	e := person.Employee{ID: 1, Position: "Developer", Person: p}

	e.PrintEmployee(e)
}
