package average

func Average(values ...int) (result float32, err string) {
	var sum int
	for _, value := range values {
		if value < 0 {
			err = "No se permiten numeros negativos"
			return
		}
		sum += value
	}
	result = float32((sum / len(values)))
	return result, ""
}
