package main

import "fmt"

func main() {
	apples := [2]string{"red", "green"}

	println("apple 1:", apples[0])

	apples[0] = "yellow"

	fmt.Println("Apples: ", apples)
}
