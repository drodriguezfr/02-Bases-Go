package alumnos

import "fmt"

type Alumno struct {
	Name     string
	Apellido string
	DNI      string
	Fecha    string
}

func (a *Alumno) Detalle(Alumno) {
	fmt.Println("Nombre: ", a.Name)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
}
