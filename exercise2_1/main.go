package main

import "fmt"

var Palabra string = "hola"

func main() {
	fmt.Println("Cantidad de letras: ", len(Palabra))
	//for x, character := range Palabra {
	//	fmt.Println(x, character)
	//}

	for x := 0; x < len(Palabra); x++ {
		fmt.Printf("%c \n", Palabra[x])
	}
}
