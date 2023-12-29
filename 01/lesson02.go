package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result1 bool
	var result2 float64
	var result3 int64

	result1, _ = strconv.ParseBool("true")
	result2, _ = strconv.ParseFloat("3.1415", 64)
	result3, _ = strconv.ParseInt("-42", 10, 64)

	fmt.Println("result 1: ", result1)
	fmt.Println("result 2: ", result2)
	fmt.Println("result 3: ", result3)

	//-------------------------------------------------------------------

	var a int = 123
	var b float64 = 4.5
	var c bool = true

	var result4 string = fmt.Sprintf("%d", a)
	var result5 string = fmt.Sprintf("%2.1f", b)
	var result6 string = fmt.Sprintf("%t", c)

	fmt.Println("result 4: ", result4)
	fmt.Println("result 5: ", result5)
	fmt.Println("result 6: ", result6)
}
