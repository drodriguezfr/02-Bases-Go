package main

var NumeroMes int = 12

func main() {
	switch NumeroMes {
	case 1:
		println("Enero")
	case 2:
		println("Febrero")
	case 3:
		println("Marzo")
	case 4:
		println("Abril")
	case 5:
		println("Mayo")
	case 6:
		println("Junio")
	case 7:
		println("Julio")
	case 8:
		println("Agosto")
	case 9:
		println("Septiembre")
	case 10:
		println("Octubre")
	case 11:
		println("Noviembre")
	case 12:
		println("Diciembre")
	default:
		println("Selecciona un numero de mes valido")
	}
}
