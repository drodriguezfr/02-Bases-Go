package main

import (
	"app/internal/salarycalculator"
	"fmt"
)

func main() {
	result, err := salarycalculator.Salarycalculator(60, "c")
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Salario: ", result)
}
