package taxcalculator

func Taxcalculator(salary float32) (tax float32) {
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	} else {
		tax = salary
	}
	return
}
