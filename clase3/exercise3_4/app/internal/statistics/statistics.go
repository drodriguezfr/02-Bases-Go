package statistics

type StatisticsOperation string

const (
	minimum StatisticsOperation = "minimum"

	average = "average"

	maximum = "maximum"
)

type StatisticsFunction func([]int) (float32, string)

func Minimum(values ...int) (result float32, err string) {
	if len(values) < 1 {
		err = "Ingresa al menos un numero"
		return
	}

	result = float32(values[0])
	for _, value := range values {
		if value < result {
			result = value
		}
	}
	return result, ""
}

func Maximum(values ...int) (result int, err string) {
	if len(values) < 1 {
		err = "Ingresa al menos un numero"
		return
	}

	result = values[0]
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result, ""
}

func Average(values ...int) (result float32, err string) {
	var sum int
	for _, value := range values {
		sum += value
	}
	result = float32((sum / len(values)))
	return result, ""
}
