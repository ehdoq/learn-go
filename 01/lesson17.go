package main

import (
	"fmt"
	entities "learn-golang/entities"
)

func main() {
	product1 := entities.Product2{
		Id:       "p01",
		Name:     "tivi 1",
		Price:    5,
		Quantity: 9,
		Status:   false,
	}

	fmt.Println("Product 1 Info")
	fmt.Println(product1.ToString())
	fmt.Println("Total: ", product1.Total())

	product2 := entities.NewProduct("p02", "name 2", 2, 7, true)
	fmt.Println("Product 2 Info")
	fmt.Println(product2.ToString())
	fmt.Println("Total: ", product2.Total())

	//-------------------------------------------------------------------

	product1.ChangeValue1("abc")
	fmt.Println("Product 1 Info")
	fmt.Println(product1.ToString())

	product3 := entities.NewProduct("p03", "name 3", 2, 7, true)
	product3.ChangeValue2("def")
	fmt.Println("Product 2 Info")
	fmt.Println(product3.ToString())
}
