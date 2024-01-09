package product

import "fmt"

func GetById(n int) {
	fmt.Println("Producto: ", Products[n-1])
}

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{{1, "apple", 15.0, "red apple", "fruit"}, {2, "tomato", 10.5, "red tomato", "vegetable"}}

func (p Product) Save(Product) []Product {
	Products = append(Products, p)
	return Products
}

func (p Product) GetAll() {
	for _, x := range Products {
		fmt.Println("Producto: ", x)
	}

}
