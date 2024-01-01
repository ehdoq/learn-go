package main

import (
	"fmt"
	"learn-golang/entities"
)

func main() {
	var animal entities.Animal
	dog := entities.Dog{}
	cat := entities.Cat{}

	animal = dog
	fmt.Println(animal.Sound())

	animal = cat
	fmt.Println(animal.Sound())
}
