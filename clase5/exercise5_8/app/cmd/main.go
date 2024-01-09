package main

import (
	"app/internal"
	"app/internal/repository"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// dependencies
	customerRepo := repository.NewCustomersMap()

	// open a file
	f, err := os.Open("customers.txt")

	if err != nil {
		defer println("ejecucion finalizada")
		panic("The indicated file was not found or is damaged")

	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		linea := scanner.Text()
		data := string(linea)

		// storing txt into split
		vals := strings.Split(data, ",")

		// converting str into int
		id, err := strconv.Atoi(vals[0])
		if err != nil {
			panic("error, falta un valor")
		}

		cm := internal.Customer{Id: id, File: vals[0], Name: vals[1], Phone: vals[2], Home: vals[3]}

		erro := customerRepo.Save(&cm)

		if erro != nil {
			if err == repository.ErrorInvalidCustomer {
				fmt.Println("value cannot be zero")
			}
		}

	}

	defer println("ejecucion finalizada")

}
