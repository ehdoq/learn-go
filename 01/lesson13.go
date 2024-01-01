package main

import "fmt"

func main() {
	MapDemo1()
	MapDemo2()
	MapDemo3()
	MapDemo4()

	//-------------------------------------------------------------------

	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	MapDisplay(student)

	//-------------------------------------------------------------------

	fmt.Println("Demo 5")
	NestedMapDemo1()

	fmt.Println("Demo 6")
	NestedMapDemo2()

	fmt.Println("Demo 7")
	NestedMapDemo3()

	fmt.Println("Demo 8")
	NestedMapDemo4()
}

func MapDemo1() {
	student := make(map[string]string)
	student["id"] = "st01"
	student["name"] = "name 1"
	student["address"] = "address 1"

	fmt.Println("Student Info")
	fmt.Println(student)

	fmt.Println("Student Info")
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func MapDemo2() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	student["phone"] = "12345"

	fmt.Println("Student Info")
	fmt.Println(student)

	fmt.Println("Student Info")
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func MapDemo3() {
	student := map[string]string{
		"id":      "st01",
		"name":    "name 1",
		"address": "address 1",
	}
	value, ok := student["name"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}
}

func MapDemo4() {
	student := make(map[string]interface{})
	student["id"] = "st01"
	student["name"] = "name 1"
	student["address"] = "address 1"
	student["age"] = 20
	student["score"] = 7.8
	student["status"] = true

	fmt.Println("Student Info")
	fmt.Println(student)

	fmt.Println("Student Info")
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func MapDisplay(student map[string]string) {
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}

func NestedMapDemo1() {
	categories := map[string][]string{
		"category1": []string{"product 1", "product 2"},
		"category2": []string{"product 3", "product 4", "product 5"},
		"category3": []string{"product 1"},
	}
	fmt.Println(categories)
	fmt.Println("Categories List")
	for key, value := range categories {
		fmt.Println(key)
		for _, v := range value {
			fmt.Println("\t", v)
		}
	}
}

func NestedMapDemo2() {
	categories := make(map[string][]string)

	categories["category1"] = append(categories["category1"], "product 1.1")
	categories["category1"] = append(categories["category1"], "product 1.2")
	categories["category1"] = append(categories["category1"], "product 1.3", "product 1.4")

	categories["category2"] = append(categories["category2"], "product 2.1", "product 2.1")

	fmt.Println(categories)
}

func NestedMapDemo3() {
	categories := map[string][]string{
		"category1": []string{"product 1.1", "product 1.2", "product 1.3"},
		"category2": []string{"product 2.1", "product 2.1"},
	}
	fmt.Println(categories)
	for _, category := range categories {
		fmt.Println(category)
		for _, name := range category {
			fmt.Println("\t", name)
		}
	}
}

func NestedMapDemo4() {
	products := map[string]map[string]interface{}{
		"product1": map[string]interface{}{
			"id":       "p01",
			"name":     "name 1",
			"price":    4.5,
			"quantity": 10,
		},
		"product2": map[string]interface{}{
			"id":       "p02",
			"name":     "name 3",
			"price":    6,
			"quantity": 4,
		},
		"product3": map[string]interface{}{
			"id":       "p03",
			"name":     "name 3",
			"price":    13,
			"quantity": 9,
		},
	}
	fmt.Println(products)
	for _, product := range products {
		fmt.Println("Info of ", product["name"])
		for key, value := range product {
			fmt.Println(key, ":", value)
		}
		fmt.Println("----------------------")
	}
}
