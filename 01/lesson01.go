package main

import(
	"fmt"
)

func main(){
	var age int = 20
	var fullName string = "ABC"
	var status bool = true
	var price float64 = 4.5
	var key rune = 'A'

	fmt.Println("age:", age)
	fmt.Println("full name:", fullName)
	fmt.Println("status:", status)
	fmt.Println("price:", price)
	fmt.Printf("key: %c and position in ASCII: %d", key, key)
}