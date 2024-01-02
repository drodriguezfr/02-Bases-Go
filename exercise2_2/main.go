package main

var (
	Edad         int     = 22
	Empleado     bool    = true
	TiempoEmpleo         = false
	Sueldo       float32 = 1000
)

func main() {
	if Edad >= 22 && Empleado && TiempoEmpleo {
		if Sueldo >= 100000 {
			println("Prestamo sin intereses otorgado")
		} else {
			println("Prestamo con intereses otorgado")
		}

	} else {
		println("Prestamo declinado")
	}

}
