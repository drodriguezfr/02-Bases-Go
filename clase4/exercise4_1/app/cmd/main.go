package main

import (
	"app/app/internal/product"
)

func main() {
	p := product.Product{ID: 3, Name: "pear", Price: 12.0, Description: "green pear", Category: "fruit"}
	p.Save(p)
	//fmt.Println(product.Products)

	p.GetAll()

	product.GetById(3)

}
