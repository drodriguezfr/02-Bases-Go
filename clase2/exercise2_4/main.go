package main

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	println(employees["Benjamin"])

	employees["Federico"] = 25
	delete(employees, "Pedro")

}
