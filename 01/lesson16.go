package main

import ()

func main() {
	fmt.Println("Demo 1")
	StructureDemo1()

	fmt.Println("Demo 2")
	StructureDemo2()
}

func StructureDemo1() {
	var product entities.Product
	product.Id = "p01"
	product.Name = "name 1"
	product.Price = 4.5
	product.Quantity = 20
	product.Status = true
	fmt.Println("Product Info")
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("status: ", product.Status)
	fmt.Println("total: ", product.Price*float64(product.Quantity))
}
