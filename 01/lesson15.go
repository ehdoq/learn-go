package main

import (
	"fmt"
	mypackage1 "learn-golang/mypackages/mypackage1"
	mypackage2 "learn-golang/mypackages/mypackage2"
)

func main() {
	mypackage1.Init()
	fmt.Println(mypackage1.Hello())
	fmt.Println("Add: ", mypackage2.Add(1, 2))
	fmt.Println("Sub: ", mypackage2.Sub(10, 4))
}
