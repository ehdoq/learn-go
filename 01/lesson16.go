package main

import (
	"fmt"
	entities "learn-golang/entities"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Demo 1")
	StructureDemo1()

	fmt.Println("Demo 2")
	StructureDemo2()

	//-------------------------------------------------------------------

	product := entities.Product{
		Id:       "p03",
		Name:     "name 3",
		Price:    5,
		Quantity: 9,
		Status:   false,
	}
	fmt.Println("Product Info")
	DisplayStructure(product)

	//-------------------------------------------------------------------

	fmt.Println("Product Info")
	product2 := Find()
	DisplayStructure(product2)

	//-------------------------------------------------------------------

	var products = []entities.Product{}
	products = append(products, entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    5,
		Quantity: 9,
		Status:   false,
	})
	products = append(products, entities.Product{
		Id:       "p02",
		Name:     "name 2",
		Price:    2,
		Quantity: 8,
		Status:   true,
	})
	products = append(products, entities.Product{
		Id:       "p03",
		Name:     "name 3",
		Price:    11,
		Quantity: 7,
		Status:   false,
	})

	fmt.Println("Product List")
	DisplayStructure1(products)
	fmt.Println("Product List")
	DisplayStructure2(products)

	//-------------------------------------------------------------------

	var products2 = []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "name 1",
			Price:    5,
			Quantity: 9,
			Status:   false,
		},
		entities.Product{
			Id:       "p02",
			Name:     "name 2",
			Price:    2,
			Quantity: 8,
			Status:   true,
		},
		entities.Product{
			Id:       "p03",
			Name:     "name 3",
			Price:    11,
			Quantity: 7,
			Status:   false,
		},
	}

	fmt.Println("Product List")
	DisplayStructure3(products2)
	fmt.Println("Product List")
	DisplayStructure4(products2)

	//-------------------------------------------------------------------

	fmt.Println("Total", Total(products2))

	//-------------------------------------------------------------------

	var min float64 = 5
	var max float64 = 20
	result := Count(min, max, products2)
	fmt.Println("Result: ", result)

	//-------------------------------------------------------------------

	var id string = "p02"
	result2 := isExists(id, products2)
	fmt.Println("Result: ", result2)

	//-------------------------------------------------------------------

	min2, max2 := FindMaxAndMin(products2)

	fmt.Println("Product Min")
	DisplayStructure(min2)

	fmt.Println("Product Max")
	DisplayStructure(max2)

	//-------------------------------------------------------------------

	fmt.Println("Search product")
	keyword := "vi"
	result3 := Search(keyword, products)
	if len(result3) == 0 {
		fmt.Println("No Result")
	} else {
		fmt.Println("Product List")
		for _, product := range result3 {
			DisplayStructure(product)
			fmt.Println("-------------------------")
		}
	}

	//-------------------------------------------------------------------

	fmt.Println("Sort Product ASC")
	SortASC(products)
	for _, product := range products {
		DisplayStructure(product)
		fmt.Println("-------------------------")
	}

	fmt.Println("Sort Product DESC")
	SortDESC(products)
	for _, product := range products {
		DisplayStructure(product)
		fmt.Println("-------------------------")
	}

	//-------------------------------------------------------------------

	var products3 = []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "tivi 1",
			Price:    5,
			Quantity: 9,
			Status:   false,
			Category: entities.Category{
				Id:   "c1",
				Name: "Category 1",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m1",
				Name:    "Manufacturer 1",
				Address: "Address 1",
			},
		},
		entities.Product{
			Id:       "p02",
			Name:     "tivi 2",
			Price:    2,
			Quantity: 8,
			Status:   true,
			Category: entities.Category{
				Id:   "c1",
				Name: "Category 1",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m1",
				Name:    "Manufacturer 1",
				Address: "Address 1",
			},
		},
		entities.Product{
			Id:       "p03",
			Name:     "laptop 3",
			Price:    11,
			Quantity: 7,
			Status:   false,
			Category: entities.Category{
				Id:   "c2",
				Name: "Category 2",
			},
			Manufacturer: entities.Manufacturer{
				Id:      "m2",
				Name:    "Manufacturer 2",
				Address: "Address 2",
			},
		},
	}

	fmt.Println("Product List")
	DisplayProductList(products3)

	//-------------------------------------------------------------------

	product4 := entities.Product{
		Id:       "p01",
		Name:     "tivi 1",
		Price:    5,
		Quantity: 9,
		Status:   false,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Manufacturer: entities.Manufacturer{
			Id:      "m1",
			Name:    "Manufacturer 1",
			Address: "Address 1",
		},
	}

	p := &product4

	fmt.Println("Product Info")
	fmt.Println("id: ", p.Id)
	fmt.Println("name: ", p.Name)
	fmt.Println("status: ", p.Status)
	fmt.Println("price: ", p.Price)
	fmt.Println("quantity: ", p.Quantity)
	fmt.Println("total: ", p.Price*float64(p.Quantity))
	fmt.Println("Category Info")
	fmt.Println("\tid: ", p.Category.Id)
	fmt.Println("\tname: ", p.Category.Name)
	fmt.Println("Manufacturer Info")
	fmt.Println("\tid: ", p.Manufacturer.Id)
	fmt.Println("\tname: ", p.Manufacturer.Name)
	fmt.Println("\taddress: ", p.Manufacturer.Address)
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

func StructureDemo2() {
	product := entities.Product{
		Id:       "p02",
		Name:     "name 2",
		Price:    5,
		Quantity: 9,
		Status:   false,
	}
	fmt.Println("Product Info")
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("status: ", product.Status)
	fmt.Println("total: ", product.Price*float64(product.Quantity))
}

func DisplayStructure(product entities.Product) {
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("status: ", product.Status)
	fmt.Println("total: ", product.Price*float64(product.Quantity))
}

func Find() entities.Product {
	product := entities.Product{
		Id:       "p02",
		Name:     "name 2",
		Price:    5,
		Quantity: 9,
		Status:   false,
	}
	return product
}

func DisplayStructure1(products []entities.Product) {
	for i := 0; i < len(products); i++ {
		fmt.Println("Id: ", products[i].Id)
		fmt.Println("Name: ", products[i].Name)
		fmt.Println("Price: ", products[i].Price)
		fmt.Println("Quantity: ", products[i].Quantity)
		fmt.Println("Status: ", products[i].Status)
		fmt.Println("Sub Total: ", products[i].Price*float64(products[i].Quantity))
		fmt.Println("==========================")
	}
}

func DisplayStructure2(products []entities.Product) {
	for _, product := range products {
		fmt.Println("Id: ", product.Id)
		fmt.Println("Name: ", product.Name)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Quantity: ", product.Quantity)
		fmt.Println("Status: ", product.Status)
		fmt.Println("Sub Total: ", product.Price*float64(product.Quantity))
		fmt.Println("==========================")
	}
}

func DisplayStructure3(products []entities.Product) {
	for i := 0; i < len(products); i++ {
		fmt.Println("Id: ", products[i].Id)
		fmt.Println("Name: ", products[i].Name)
		fmt.Println("Price: ", products[i].Price)
		fmt.Println("Quantity: ", products[i].Quantity)
		fmt.Println("Status: ", products[i].Status)
		fmt.Println("Sub Total: ", products[i].Price*float64(products[i].Quantity))
		fmt.Println("==========================")
	}
}

func DisplayStructure4(products []entities.Product) {
	for _, product := range products {
		fmt.Println("Id: ", product.Id)
		fmt.Println("Name: ", product.Name)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Quantity: ", product.Quantity)
		fmt.Println("Status: ", product.Status)
		fmt.Println("Sub Total: ", product.Price*float64(product.Quantity))
		fmt.Println("==========================")
	}
}

func Total(products []entities.Product) (result float64) {
	result = 0
	for _, product := range products {
		result += product.Price * float64(product.Quantity)
	}
	return result
}

func Count(min, max float64, products []entities.Product) (counter int) {
	counter = 0
	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			counter++
		}
	}
	return counter
}

func isExists(id string, products []entities.Product) (result bool) {
	result = false
	for _, product := range products {
		if product.Id == id {
			result = true
			break
		}
	}
	return result
}

func FindMaxAndMin(products []entities.Product) (min entities.Product, max entities.Product) {
	min = products[0]
	max = products[0]
	for _, product := range products {
		if product.Price > max.Price {
			max = product
		}
		if product.Price < min.Price {
			min = product
		}
	}
	return min, max
}

func Search(keyword string, products []entities.Product) (result []entities.Product) {
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
			result = append(result, product)
		}
	}
	return result
}

func SortDESC(products []entities.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
}

func SortASC(products []entities.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}

func DisplayProductList(products []entities.Product) {
	for _, product := range products {
		DisplayProduct(product)
		fmt.Println("-------------------------------")
	}
}

func DisplayProduct(product entities.Product) {
	fmt.Println("id: ", product.Id)
	fmt.Println("name: ", product.Name)
	fmt.Println("status: ", product.Status)
	fmt.Println("price: ", product.Price)
	fmt.Println("quantity: ", product.Quantity)
	fmt.Println("total: ", product.Price*float64(product.Quantity))
	fmt.Println("Category Info")
	fmt.Println("\tid: ", product.Category.Id)
	fmt.Println("\tname: ", product.Category.Name)
	fmt.Println("Manufacturer Info")
	fmt.Println("\tid: ", product.Manufacturer.Id)
	fmt.Println("\tname: ", product.Manufacturer.Name)
	fmt.Println("\taddress: ", product.Manufacturer.Address)
}
