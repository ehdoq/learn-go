package main

import (
	"fmt"
)

func main() {
	var finger int = 3

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Invalid")
	}

	//-------------------------------------------------------------------

	a := 2
	switch a {
	case 1, 2, 3:
		fmt.Println("a = 1, 2, 3")
	case 4, 5, 6:
		fmt.Println("a = 4, 5, 6")
	default:
		fmt.Println("Invalid")
	}

	//-------------------------------------------------------------------

	switch {
	case a >= 1 && a <= 10:
		fmt.Println("a >= 1 and a <= 10")
	case a >= 10 && a <= 20:
		fmt.Println("a >= 10 and a <= 20")
	default:
		fmt.Println("Invalid")
	}
}
