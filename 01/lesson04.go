package main

import (
	"fmt"
)

func main() {
	var a int = 5

	if a > 2 {
		fmt.Println("a > 2")
	}

	var b int = 10

	if b > 4 && b%2 == 0 {
		fmt.Println("b > 4 and b is an even number")
	}

	//-------------------------------------------------------------------

	if a > 2 {
		fmt.Println("a > 2")
	} else {
		fmt.Println("a <= 2")
	}

	if b := -5; b > 3 {
		fmt.Println("b > 3")
	} else {
		fmt.Println("b <= 3")
	}

	//-------------------------------------------------------------------

	var finger int = 3

	if finger == 1 {
		fmt.Println("Thumb")
	} else if finger == 2 {
		fmt.Println("Index")
	} else if finger == 3 {
		fmt.Println("Middle")
	} else if finger == 4 {
		fmt.Println("Ring")
	} else if finger == 5 {
		fmt.Println("Pinky")
	} else {
		fmt.Println("Invalid")
	}
}
