package main

import (
	"app/internal/taxcalculator"
	"fmt"
)

func main() {
	salary := 160000

	result := taxcalculator.Taxcalculator(float32(salary))
	fmt.Println("Resulting tax: ", result)
}
