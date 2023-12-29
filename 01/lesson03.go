package main

import (
	"fmt"
)

func main() {
	var firstName string = "John"
	var middleName string = "Thomas"
	var lastName string = "Smith"
	var fullName = firstName + " " + middleName + " " + lastName

	fmt.Println("Full Name: ", fullName)

	var age string = "Age:"
	var value int = 20
	var result string = fmt.Sprintf("%s %d", age, value)

	fmt.Println(result)

	//-------------------------------------------------------------------

	var age2 int = 20
	var fullName2 string = "abc"
	var status2 bool = true
	var price2 float32 = 4.5
	var key2 rune = 'B'

	fmt.Printf("age: %d\n", age2)
	fmt.Printf("full name: %s\n", fullName2)
	fmt.Printf("status: %t\n", status2)
	fmt.Printf("price: %1.1f\n", price2)
	fmt.Printf("key: %c and position in ASCII: %d\n", key2, key2)
}
