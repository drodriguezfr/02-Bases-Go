package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// open a file
	f, err := os.Open("customers.txt")

	if err != nil {
		defer println("ejecucion finalizada")
		panic("The indicated file was not found or is damaged")

	}

	// read bytes from file
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print bytes
	data := string(bytes)
	fmt.Println(data)
}
