package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	// infinite
	//

	// limited
	for i := 0; 1 <= 10; i = i + 2 {
		fmt.Println("iteration: ", i)
		fmt.Println("this is a random number", rand.Intn(10))
	}
}
