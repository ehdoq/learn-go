package main

import (
	"fmt"
	circle "learn-golang/mypackages/geometry/circle"
	rectangle "learn-golang/mypackages/geometry/rectangle"
	mypackage1 "learn-golang/mypackages/mypackage1"
	mypackage2 "learn-golang/mypackages/mypackage2"
)

func main() {
	mypackage1.Init()
	fmt.Println(mypackage1.Hello())
	fmt.Println("Add: ", mypackage2.Add(1, 2))
	fmt.Println("Sub: ", mypackage2.Sub(10, 4))

	//-------------------------------------------------------------------

	fmt.Println("Circle")
	fmt.Println("\tArea: ", circle.Area(10))
	fmt.Println("\tPerimeter: ", circle.Perimeter(5))
	fmt.Println("Rectangle")
	fmt.Println("\tArea: ", rectangle.Area(4, 6))
	fmt.Println("\tPerimeter: ", rectangle.Perimeter(5, 9))
}
