package main

import (
	"encoding/json"
	"fmt"
	entities "learn-golang/entities"
)

func main() {
	product := entities.Product{
		Id:       "p01",
		Name:     "name 1",
		Price:    4.5,
		Quantity: 8,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{Title: "title 1", Content: "content 1"},
			entities.Comment{Title: "title 2", Content: "content 2"},
			entities.Comment{Title: "title 3", Content: "content 3"},
		},
	}
	result, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}

	//-------------------------------------------------------------------

	var str string = `{
		"id":"p01",
		"name":"name 1",
		"price":4.5,
		"quantity":8,
		"status":true,
		"category":{
			"id":"c1",
			"name":"Category 1"
		},
		"comments":[
			{"title":"title 1","content":"content 1"},
			{"title":"title 2","content":"content 2"},
			{"title":"title 3","content":"content 3"}
		]
	}`

	var product2 entities.Product
	json.Unmarshal([]byte(str), &product2)

	fmt.Println("Product Info")
	fmt.Printf("Id: %s\nName: %s\nPrice: %0.2f\nQuantity: %d\nStatus: %t", product2.Id, product2.Name, product2.Price, product2.Quantity, product2.Status)
	fmt.Println("\nCategory Info")
	fmt.Println("\tCategory Id: ", product2.Category.Id)
	fmt.Println("\tCategory Name: ", product2.Category.Name)
	fmt.Println("Comment List")
	for _, comment := range product2.Comments {
		fmt.Println("\t", comment.Title)
		fmt.Println("\t", comment.Content)
		fmt.Println("\t-----------------------")
	}
}
