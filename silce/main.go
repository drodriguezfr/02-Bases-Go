package main

import "fmt"

func main() {
	names := make([]string, 3, 3)
	names[0] = "deivid"
	names[1] = "david"
	names[2] = "eldeivid"
	fmt.Println("names", names)

	names = append(names, "deivis", "deivit")

	fmt.Println("names: ", names)
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))
}
