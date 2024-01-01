package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%3d", i)
	}

	//-------------------------------------------------------------------

	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%3c", ch)
	}

	//-------------------------------------------------------------------

	var total int = 0

	for i := 1; i <= 10; i++ {
		total += i
	}

	fmt.Println("Total: ", total)

	//-------------------------------------------------------------------

	var counter int = 0

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			counter++
		}
	}

	fmt.Println("Counter: ", counter)

	//-------------------------------------------------------------------

	i := 1

	for {
		fmt.Printf("%3d", i)
		i++
		if i == 10 {
			break
		}
	}

	//-------------------------------------------------------------------

	ch := 'a'

	for {
		fmt.Printf("%3c", ch)
		ch++
		if ch == 'z' {
			break
		}
	}

	//-------------------------------------------------------------------

	var total2 int = 0

	i2 := 1

	for {
		total += i2
		i2++
		if i2 == 10 {
			break
		}
	}

	fmt.Println("Total: ", total2)

	//-------------------------------------------------------------------

	var counter2 int = 0

	i3 := 1

	for {
		if i3%2 == 0 {
			counter++
		}
		i3++
		if i3 == 10 {
			break
		}
	}
	fmt.Println("Counter: ", counter2)

	//-------------------------------------------------------------------

	i4 := 1

	for {
		fmt.Printf("%3d", i4)
		i4++
		if i4 == 10 {
			break
		}
	}

	//-------------------------------------------------------------------

	ch1 := 'a'

	for {
		fmt.Printf("%3c", ch1)

		ch1++

		if ch1 == 'z' {
			break
		}
	}

	//-------------------------------------------------------------------

	var total3 int = 0

	i5 := 1

	for {
		total3 += i5

		i5++

		if i5 == 10 {
			break
		}
	}

	fmt.Println("Total: ", total3)

	//-------------------------------------------------------------------

	var counter1 int = 0

	i6 := 1

	for {
		if i6%2 == 0 {
			counter1++
		}

		i6++

		if i6 == 10 {
			break
		}
	}

	fmt.Println("Counter: ", counter1)

	//-------------------------------------------------------------------

	i7 := 1

	for {
		if i7%2 == 0 {
			continue
		}

		i7++

		fmt.Printf("%3d", i7)

		if i7 == 10 {
			break
		}
	}
}
