package main

var Edad int

func main() {

	switch {
	case Edad > 18:
		println("Mayor")
	case (Edad < 18 && Edad != 0):
		println("Menor")
	default:
		println("Ingresa una edad")
	}

}
