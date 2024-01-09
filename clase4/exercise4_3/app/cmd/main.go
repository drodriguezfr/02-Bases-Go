package main

import "app/internal/alumnos"

func main() {
	alumno := alumnos.Alumno{Name: "David", Apellido: "Rodriguez", DNI: "00X1234A", Fecha: "12.02.2019"}
	alumno.Detalle(alumno)
}
