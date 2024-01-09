package salarycalculator

func Salarycalculator(minutos int, categoria string) (salario float32, err string) {
	horas := minutos / 60
	err = ""
	switch categoria {
	case "a":
		mes := float32(horas * 3000)
		salario = mes * 1.5
	case "b":
		mes := float32(horas * 1500)
		salario = mes * 1.2
	case "c":
		salario = float32(horas * 1000)
	default:
		err = "Ingresa una categoria valida"
	}
	return salario, err
}
