package main

import (
	"app/internal/average"
	"fmt"
)

func main() {
	result, err := average.Average(1, 2, -3, 4, 5)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Promedio: ", result)

}
