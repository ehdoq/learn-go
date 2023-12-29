package main

import (
	"fmt"
)

func main() {
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

	//-------------------------------------------------------------

	var a1, a2, a3 int = 20, -5, 7

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	//-------------------------------------------------------------------

	a := 123

	fmt.Println("a:", a)

	id2, fullName2, age2, status2, price2, key2 := "p01", "name 1", 20, true, 4.5, 'D'

	fmt.Println("id:", id2)
	fmt.Println("age:", age2)
	fmt.Println("full name:", fullName2)
	fmt.Println("status:", status2)
	fmt.Println("price:", price2)
	fmt.Printf("key: %c and position in ASCII: %d", key2, key2)
}
